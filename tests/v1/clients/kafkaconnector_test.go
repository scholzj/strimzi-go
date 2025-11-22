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

func NewConnector() *kafkav1.KafkaConnector {
	tasksMax := int32(3)
	return &kafkav1.KafkaConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name:      NAME,
			Namespace: NAMESPACE,
		},
		Spec: &kafkav1.KafkaConnectorSpec{
			TasksMax: &tasksMax,
			Class:    "org.apache.kafka.connect.file.FileStreamSourceConnector",
			Config:   map[string]interface{}{"file": "/opt/kafka/LICENSE", "topic": "my-topic"},
		},
	}
}

func UpdatedConnector(oldResource *kafkav1.KafkaConnector) *kafkav1.KafkaConnector {
	updatedResource := oldResource.DeepCopy()

	tasksMax := int32(5)
	updatedResource.Spec.TasksMax = &tasksMax
	updatedResource.Spec.Class = "org.apache.kafka.connect.file.FileStreamSinkConnector"
	updatedResource.Spec.Config["file"] = "/opt/kafka/README"

	return updatedResource
}

func AssertConnectors(t *testing.T, r1, r2 *kafkav1.KafkaConnector) {
	assert.Equal(t, r1.ObjectMeta.Name, r2.ObjectMeta.Name)
	assert.Equal(t, r1.ObjectMeta.Namespace, r2.ObjectMeta.Namespace)
	assert.Equal(t, r1.Spec.TasksMax, r2.Spec.TasksMax)
	assert.Equal(t, r1.Spec.Class, r2.Spec.Class)
	assert.Equal(t, r1.Spec.Config["file"], r2.Spec.Config["file"])
	assert.Equal(t, r1.Spec.Config["topic"], r2.Spec.Config["topic"])
}

func TestKafkaConnectorCreateUpdateDelete(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1().KafkaConnectors(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

	// Create the resource
	newResource := NewConnector()
	_, err := client.KafkaV1().KafkaConnectors(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err := client.KafkaV1().KafkaConnectors(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertConnectors(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedConnector(actualResource)
	_, err = client.KafkaV1().KafkaConnectors(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err = client.KafkaV1().KafkaConnectors(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertConnectors(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1().KafkaConnectors(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete actualResource: %s", err.Error())
	}

	// Check deletion
	actualResource, err = client.KafkaV1().KafkaConnectors(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			t.Fatalf("Failed to get resource: %s", err.Error())
		}
	} else if actualResource != nil {
		t.Fatalf("Resource still exists")
	}
}

func TestKafkaConnectorInformerAndLister(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1().KafkaConnectors(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

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
	informer := factory.Kafka().V1().KafkaConnectors()
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
	newResource := NewConnector()
	_, err = client.KafkaV1().KafkaConnectors(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create resource: %s", err.Error())
	}

	// Wait for the resource to be added in the informer
	<-addedSignal

	// Get the resource
	actualResource, err := lister.KafkaConnectors(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertConnectors(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedConnector(actualResource)
	_, err = client.KafkaV1().KafkaConnectors(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update resource: %s", err.Error())
	}

	// Wait for resource to be updated in the informer
	<-updatedSignal

	// Get the resource
	actualResource, err = lister.KafkaConnectors(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertConnectors(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1().KafkaConnectors(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete resource: %s", err.Error())
	}

	// Wait for resource to be deleted
	<-deletedSignal

	// Check deletion
	actualResource, err = lister.KafkaConnectors(NAMESPACE).Get(NAME)
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
