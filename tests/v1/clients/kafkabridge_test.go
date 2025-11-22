package clients

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	kafkav1 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1"
	strimziinformer "github.com/scholzj/strimzi-go/pkg/client/informers/externalversions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func NewBridge() *kafkav1.KafkaBridge {
	return &kafkav1.KafkaBridge{
		ObjectMeta: metav1.ObjectMeta{
			Name:      NAME,
			Namespace: NAMESPACE,
		},
		Spec: &kafkav1.KafkaBridgeSpec{
			Replicas:         1,
			BootstrapServers: "my-kafka:9092",
			Http:             &kafkav1.KafkaBridgeHttpConfig{Port: 8080},
		},
	}
}

func UpdatedBridge(oldResource *kafkav1.KafkaBridge) *kafkav1.KafkaBridge {
	updatedResource := oldResource.DeepCopy()

	updatedResource.Spec.Replicas = 3
	updatedResource.Spec.BootstrapServers = "my-other-kafka:9092"
	updatedResource.Spec.Http.Port = 8443

	return updatedResource
}

func AssertBridge(t *testing.T, r1, r2 *kafkav1.KafkaBridge) {
	assert.Equal(t, r1.ObjectMeta.Name, r2.ObjectMeta.Name)
	assert.Equal(t, r1.ObjectMeta.Namespace, r2.ObjectMeta.Namespace)
	assert.Equal(t, r1.Spec.Replicas, r2.Spec.Replicas)
	assert.Equal(t, r1.Spec.BootstrapServers, r2.Spec.BootstrapServers)
	assert.Equal(t, r1.Spec.Http.Port, r2.Spec.Http.Port)
}

func TestKafkaBridgeCreateUpdateDelete(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1().KafkaBridges(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

	// Create the resource
	newResource := NewBridge()
	_, err := client.KafkaV1().KafkaBridges(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err := client.KafkaV1().KafkaBridges(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertBridge(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedBridge(actualResource)
	_, err = client.KafkaV1().KafkaBridges(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err = client.KafkaV1().KafkaBridges(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertBridge(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1().KafkaBridges(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete actualResource: %s", err.Error())
	}

	// Check deletion
	actualResource, err = client.KafkaV1().KafkaBridges(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			t.Fatalf("Failed to get resource: %s", err.Error())
		}
	} else if actualResource != nil {
		t.Fatalf("Resource still exists")
	}
}

func TestKafkaBridgeInformerAndLister(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1().KafkaBridges(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

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
	informer := factory.Kafka().V1().KafkaBridges()
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
	newResource := NewBridge()
	_, err = client.KafkaV1().KafkaBridges(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create resource: %s", err.Error())
	}

	// Wait for the resource to be added in the informer
	<-addedSignal

	// Get the resource
	actualResource, err := lister.KafkaBridges(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertBridge(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedBridge(actualResource)
	_, err = client.KafkaV1().KafkaBridges(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update resource: %s", err.Error())
	}

	// Wait for resource to be updated in the informer
	<-updatedSignal

	// Get the resource
	actualResource, err = lister.KafkaBridges(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertBridge(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1().KafkaBridges(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete resource: %s", err.Error())
	}

	// Wait for resource to be deleted
	<-deletedSignal

	// Check deletion
	actualResource, err = lister.KafkaBridges(NAMESPACE).Get(NAME)
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
