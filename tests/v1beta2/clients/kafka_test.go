package clients

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"

	kafkav1beta2 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1beta2"
	strimziinformer "github.com/scholzj/strimzi-go/pkg/client/informers/externalversions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func NewKafka() *kafkav1beta2.Kafka {
	return &kafkav1beta2.Kafka{
		ObjectMeta: metav1.ObjectMeta{
			Name:      NAME,
			Namespace: NAMESPACE,
		},
		Spec: &kafkav1beta2.KafkaSpec{
			Kafka: &kafkav1beta2.KafkaClusterSpec{
				Version: "3.9.0",
				Listeners: []kafkav1beta2.GenericKafkaListener{{
					Name: "internal",
					Type: kafkav1beta2.INTERNAL_KAFKALISTENERTYPE,
					Tls:  true,
					Port: 9092,
				}},
				Config: map[string]interface{}{
					"default.replication.factor": 3,
					"min.insync.replicas":        2,
				},
			},
			EntityOperator: &kafkav1beta2.EntityOperatorSpec{
				TopicOperator: &kafkav1beta2.EntityTopicOperatorSpec{},
				UserOperator:  &kafkav1beta2.EntityUserOperatorSpec{},
			},
		},
	}
}

func UpdatedKafka(oldResource *kafkav1beta2.Kafka) *kafkav1beta2.Kafka {
	updatedResource := oldResource.DeepCopy()

	updatedResource.Spec.Kafka.Version = "4.0.0"
	updatedResource.Spec.Kafka.Listeners = append(updatedResource.Spec.Kafka.Listeners, kafkav1beta2.GenericKafkaListener{
		Name: "external",
		Type: kafkav1beta2.LOADBALANCER_KAFKALISTENERTYPE,
		Tls:  true,
		Port: 9094,
	})
	updatedResource.Spec.Kafka.Config["min.insync.replicas"] = 3

	return updatedResource
}

func AssertKafkas(t *testing.T, r1, r2 *kafkav1beta2.Kafka) {
	assert.Equal(t, r1.ObjectMeta.Name, r2.ObjectMeta.Name)
	assert.Equal(t, r1.ObjectMeta.Namespace, r2.ObjectMeta.Namespace)
	assert.Equal(t, r1.Spec.Kafka.Version, r2.Spec.Kafka.Version)
	assert.Equal(t, len(r1.Spec.Kafka.Listeners), len(r2.Spec.Kafka.Listeners))
	assert.Equal(t, r1.Spec.Kafka.Listeners, r2.Spec.Kafka.Listeners)
	assert.EqualValues(t, r1.Spec.Kafka.Config["default.replication.factor"], r2.Spec.Kafka.Config["default.replication.factor"])
	assert.EqualValues(t, r1.Spec.Kafka.Config["min.insnc.replicas"], r2.Spec.Kafka.Config["min.insnc.replicas"])
}

func TestKafkaCreateUpdateDelete(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1beta2().Kafkas(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

	// Create the resource
	newResource := NewKafka()
	_, err := client.KafkaV1beta2().Kafkas(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err := client.KafkaV1beta2().Kafkas(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertKafkas(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedKafka(actualResource)
	_, err = client.KafkaV1beta2().Kafkas(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err = client.KafkaV1beta2().Kafkas(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertKafkas(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1beta2().Kafkas(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete actualResource: %s", err.Error())
	}

	// Check deletion
	actualResource, err = client.KafkaV1beta2().Kafkas(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			t.Fatalf("Failed to get resource: %s", err.Error())
		}
	} else if actualResource != nil {
		t.Fatalf("Resource still exists")
	}
}

func TestKafkaInformerAndLister(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1beta2().Kafkas(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

	// Setup informer
	added := 0
	addedSignal := make(chan struct{})
	defer close(addedSignal)
	updatedSignal := make(chan struct{})
	defer close(updatedSignal)
	updated := 0
	deletedSignal := make(chan struct{})
	defer close(deletedSignal)
	deleted := 0

	// Create informer and lister
	factory := strimziinformer.NewSharedInformerFactoryWithOptions(client, time.Hour*1)
	informer := factory.Kafka().V1beta2().Kafkas()
	_, err := informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(new interface{}) {
			added++
			addedSignal <- struct{}{}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			updated++
			updatedSignal <- struct{}{}
		},
		DeleteFunc: func(old interface{}) {
			deleted++
			deletedSignal <- struct{}{}
		},
	})
	if err != nil {
		t.Fatalf("Failed to create informer: %s", err.Error())
	}

	stop := make(chan struct{})
	defer close(stop)
	factory.Start(stop)
	if !cache.WaitForCacheSync(stop, informer.Informer().HasSynced) {
		t.Fatalf("Informer failed to sync")
	}

	lister := informer.Lister()

	// Create the resource
	newResource := NewKafka()
	_, err = client.KafkaV1beta2().Kafkas(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create resource: %s", err.Error())
	}

	// Wait for the resource to be added in the informer
	<-addedSignal

	// Get the resource
	actualResource, err := lister.Kafkas(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertKafkas(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedKafka(actualResource)
	_, err = client.KafkaV1beta2().Kafkas(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update resource: %s", err.Error())
	}

	// Wait for resource to be updated in the informer
	<-updatedSignal

	// Get the resource
	actualResource, err = lister.Kafkas(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertKafkas(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1beta2().Kafkas(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete resource: %s", err.Error())
	}

	// Wait for resource to be deleted
	<-deletedSignal

	// Check deletion
	actualResource, err = lister.Kafkas(NAMESPACE).Get(NAME)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			t.Fatalf("Failed to get resource: %s", err.Error())
		}
	} else if actualResource != nil {
		t.Fatalf("Resource still exists")
	}

	// Assert the event handled
	assert.Equal(t, added, 1)
	assert.Equal(t, updated, 1)
	assert.Equal(t, deleted, 1)
	if added != 1 {
		t.Fatalf("Resource was not added once but %d", added)
	}
	if updated != 1 {
		t.Fatalf("Resource was not updated once but %d", updated)
	}
	if deleted != 1 {
		t.Fatalf("Resource was not deleted once but %d", deleted)
	}
}
