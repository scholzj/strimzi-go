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

func NewConnect() *kafkav1beta2.KafkaConnect {
	return &kafkav1beta2.KafkaConnect{
		ObjectMeta: metav1.ObjectMeta{
			Name:      NAME,
			Namespace: NAMESPACE,
		},
		Spec: &kafkav1beta2.KafkaConnectSpec{
			Replicas:         1,
			BootstrapServers: "my-kafka:9092",
			Version:          "3.9.0",
			Config: map[string]interface{}{
				"config.storage.replication.factor": -1,
				"config.storage.topic":              "connect-cluster-configs",
				"group.id":                          "connect-cluster",
				"offset.storage.replication.factor": -1,
				"offset.storage.topic":              "connect-cluster-offsets",
				"status.storage.replication.factor": -1,
				"status.storage.topic":              "connect-cluster-status",
			},
		},
	}
}

func UpdatedConnect(oldResource *kafkav1beta2.KafkaConnect) *kafkav1beta2.KafkaConnect {
	updatedResource := oldResource.DeepCopy()

	updatedResource.Spec.Replicas = 3
	updatedResource.Spec.BootstrapServers = "my-other-kafka:9092"
	updatedResource.Spec.Config["group.id"] = "other-connect-cluster"
	updatedResource.Spec.Config["config.storage.topic"] = "other-connect-cluster-configs"
	updatedResource.Spec.Config["offset.storage.topic"] = "other-connect-cluster-offsets"
	updatedResource.Spec.Config["status.storage.topic"] = "other-connect-cluster-status"

	return updatedResource
}

func AssertConnects(t *testing.T, r1, r2 *kafkav1beta2.KafkaConnect) {
	assert.Equal(t, r1.ObjectMeta.Name, r2.ObjectMeta.Name)
	assert.Equal(t, r1.ObjectMeta.Namespace, r2.ObjectMeta.Namespace)
	assert.Equal(t, r1.Spec.Replicas, r2.Spec.Replicas)
	assert.Equal(t, r1.Spec.BootstrapServers, r2.Spec.BootstrapServers)
	assert.EqualValues(t, r1.Spec.Config["config.storage.replication.factor"], r2.Spec.Config["config.storage.replication.factor"])
	assert.EqualValues(t, r1.Spec.Config["config.storage.topic"], r2.Spec.Config["config.storage.topic"])
	assert.EqualValues(t, r1.Spec.Config["group.id"], r2.Spec.Config["group.id"])
	assert.EqualValues(t, r1.Spec.Config["offset.storage.replication.factor"], r2.Spec.Config["offset.storage.replication.factor"])
	assert.EqualValues(t, r1.Spec.Config["offset.storage.topic"], r2.Spec.Config["offset.storage.topic"])
	assert.EqualValues(t, r1.Spec.Config["status.storage.replication.factor"], r2.Spec.Config["status.storage.replication.factor"])
	assert.EqualValues(t, r1.Spec.Config["status.storage.topic"], r2.Spec.Config["status.storage.topic"])
}

func TestKafkaConnectCreateUpdateDelete(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1beta2().KafkaConnects(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

	// Create the resource
	newResource := NewConnect()
	_, err := client.KafkaV1beta2().KafkaConnects(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err := client.KafkaV1beta2().KafkaConnects(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertConnects(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedConnect(actualResource)
	_, err = client.KafkaV1beta2().KafkaConnects(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err = client.KafkaV1beta2().KafkaConnects(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertConnects(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1beta2().KafkaConnects(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete actualResource: %s", err.Error())
	}

	// Check deletion
	actualResource, err = client.KafkaV1beta2().KafkaConnects(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			t.Fatalf("Failed to get resource: %s", err.Error())
		}
	} else if actualResource != nil {
		t.Fatalf("Resource still exists")
	}
}

func TestKafkaConnectInformerAndLister(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1beta2().KafkaConnects(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

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
	informer := factory.Kafka().V1beta2().KafkaConnects()
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
	newResource := NewConnect()
	_, err = client.KafkaV1beta2().KafkaConnects(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create resource: %s", err.Error())
	}

	// Wait for the resource to be added in the informer
	<-addedSignal

	// Get the resource
	actualResource, err := lister.KafkaConnects(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertConnects(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedConnect(actualResource)
	_, err = client.KafkaV1beta2().KafkaConnects(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update resource: %s", err.Error())
	}

	// Wait for resource to be updated in the informer
	<-updatedSignal

	// Get the resource
	actualResource, err = lister.KafkaConnects(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertConnects(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1beta2().KafkaConnects(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete resource: %s", err.Error())
	}

	// Wait for resource to be deleted
	<-deletedSignal

	// Check deletion
	actualResource, err = lister.KafkaConnects(NAMESPACE).Get(NAME)
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
