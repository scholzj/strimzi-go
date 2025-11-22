package clients

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	kafkav1beta2 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1beta2"
	strimziinformer "github.com/scholzj/strimzi-go/pkg/client/informers/externalversions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func NewRebalance() *kafkav1beta2.KafkaRebalance {
	return &kafkav1beta2.KafkaRebalance{
		ObjectMeta: metav1.ObjectMeta{
			Name:      NAME,
			Namespace: NAMESPACE,
		},
		Spec: &kafkav1beta2.KafkaRebalanceSpec{
			Goals: []string{"CpuCapacityGoal", "DiskCapacityGoal"},
		},
	}
}

func UpdatedRebalance(oldResource *kafkav1beta2.KafkaRebalance) *kafkav1beta2.KafkaRebalance {
	updatedResource := oldResource.DeepCopy()

	updatedResource.Spec.Goals = []string{"RackAwareGoal", "ReplicaCapacityGoal"}

	return updatedResource
}

func AssertRebalances(t *testing.T, r1, r2 *kafkav1beta2.KafkaRebalance) {
	assert.Equal(t, r1.ObjectMeta.Name, r2.ObjectMeta.Name)
	assert.Equal(t, r1.ObjectMeta.Namespace, r2.ObjectMeta.Namespace)
	assert.Equal(t, r1.Spec.Goals, r2.Spec.Goals)
}

func TestKafkaRebalanceCreateUpdateDelete(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1beta2().KafkaRebalances(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

	// Create the resource
	newResource := NewRebalance()
	_, err := client.KafkaV1beta2().KafkaRebalances(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err := client.KafkaV1beta2().KafkaRebalances(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertRebalances(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedRebalance(actualResource)
	_, err = client.KafkaV1beta2().KafkaRebalances(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err = client.KafkaV1beta2().KafkaRebalances(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertRebalances(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1beta2().KafkaRebalances(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete actualResource: %s", err.Error())
	}

	// Check deletion
	actualResource, err = client.KafkaV1beta2().KafkaRebalances(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			t.Fatalf("Failed to get resource: %s", err.Error())
		}
	} else if actualResource != nil {
		t.Fatalf("Resource still exists")
	}
}

func TestKafkaRebalanceInformerAndLister(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1beta2().KafkaRebalances(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

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
	informer := factory.Kafka().V1beta2().KafkaRebalances()
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
	newResource := NewRebalance()
	_, err = client.KafkaV1beta2().KafkaRebalances(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create resource: %s", err.Error())
	}

	// Wait for the resource to be added in the informer
	<-addedSignal

	// Get the resource
	actualResource, err := lister.KafkaRebalances(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertRebalances(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedRebalance(actualResource)
	_, err = client.KafkaV1beta2().KafkaRebalances(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update resource: %s", err.Error())
	}

	// Wait for resource to be updated in the informer
	<-updatedSignal

	// Get the resource
	actualResource, err = lister.KafkaRebalances(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertRebalances(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1beta2().KafkaRebalances(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete resource: %s", err.Error())
	}

	// Wait for resource to be deleted
	<-deletedSignal

	// Check deletion
	actualResource, err = lister.KafkaRebalances(NAMESPACE).Get(NAME)
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
