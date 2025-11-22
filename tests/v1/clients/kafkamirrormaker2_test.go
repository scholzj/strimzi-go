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

func NewMM2() *kafkav1.KafkaMirrorMaker2 {
	tasksMax := int32(5)

	return &kafkav1.KafkaMirrorMaker2{
		ObjectMeta: metav1.ObjectMeta{
			Name:      NAME,
			Namespace: NAMESPACE,
		},
		Spec: &kafkav1.KafkaMirrorMaker2Spec{
			Replicas: 1,
			Version:  "3.9.0",
			Target: &kafkav1.KafkaMirrorMaker2TargetClusterSpec{
				Alias:              "cluster-b",
				BootstrapServers:   "my-other-kafka:9092",
				GroupId:            "mm2-cluster",
				ConfigStorageTopic: "mm2-cluster-configs",
				OffsetStorageTopic: "mm2-cluster-offsets",
				StatusStorageTopic: "mm2-cluster-status",
				Config: map[string]interface{}{
					"group.id": "cluster-b",
				},
			},
			Mirrors: []kafkav1.KafkaMirrorMaker2MirrorSpec{{
				Source: &kafkav1.KafkaMirrorMaker2ClusterSpec{
					Alias:            "cluster-a",
					BootstrapServers: "my-kafka:9092",
					Config: map[string]interface{}{
						"group.id": "cluster-a",
					},
				},
				SourceConnector: &kafkav1.KafkaMirrorMaker2ConnectorSpec{
					TasksMax: &tasksMax,
					Config: map[string]interface{}{
						"sync.topic.acls.enabled": "false",
					},
				},
				CheckpointConnector: &kafkav1.KafkaMirrorMaker2ConnectorSpec{
					TasksMax: &tasksMax,
					Config: map[string]interface{}{
						"sync.group.offsets.enabled": "false",
					},
				},
			}},
		},
	}
}

func UpdatedMM2(oldResource *kafkav1.KafkaMirrorMaker2) *kafkav1.KafkaMirrorMaker2 {
	updatedResource := oldResource.DeepCopy()

	tasksMax := int32(3)

	updatedResource.Spec.Replicas = 3
	updatedResource.Spec.Version = "4.0.0"
	updatedResource.Spec.Target.BootstrapServers = "my-other-kafka-9093"
	updatedResource.Spec.Mirrors[0].Source.BootstrapServers = "my-kafka:9093"
	updatedResource.Spec.Mirrors[0].SourceConnector.TasksMax = &tasksMax
	updatedResource.Spec.Mirrors[0].CheckpointConnector.TasksMax = &tasksMax

	return updatedResource
}

func AssertMM2s(t *testing.T, r1, r2 *kafkav1.KafkaMirrorMaker2) {
	assert.Equal(t, r1.ObjectMeta.Name, r2.ObjectMeta.Name)
	assert.Equal(t, r1.ObjectMeta.Namespace, r2.ObjectMeta.Namespace)
	assert.Equal(t, r1.Spec.Replicas, r2.Spec.Replicas)
	assert.Equal(t, r1.Spec.Version, r2.Spec.Version)
	assert.Equal(t, r1.Spec.Target.BootstrapServers, r2.Spec.Target.BootstrapServers)
	assert.Equal(t, len(r1.Spec.Mirrors), len(r2.Spec.Mirrors))
	assert.Equal(t, r1.Spec.Mirrors[0].Source.BootstrapServers, r2.Spec.Mirrors[0].Source.BootstrapServers)
	assert.Equal(t, r1.Spec.Mirrors, r2.Spec.Mirrors)

}

func TestKafkaMirrorMaker2CreateUpdateDelete(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1().KafkaMirrorMaker2s(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

	// Create the resource
	newResource := NewMM2()
	_, err := client.KafkaV1().KafkaMirrorMaker2s(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err := client.KafkaV1().KafkaMirrorMaker2s(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertMM2s(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedMM2(actualResource)
	_, err = client.KafkaV1().KafkaMirrorMaker2s(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err = client.KafkaV1().KafkaMirrorMaker2s(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertMM2s(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1().KafkaMirrorMaker2s(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete actualResource: %s", err.Error())
	}

	// Check deletion
	actualResource, err = client.KafkaV1().KafkaMirrorMaker2s(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			t.Fatalf("Failed to get resource: %s", err.Error())
		}
	} else if actualResource != nil {
		t.Fatalf("Resource still exists")
	}
}

func TestKafkaMirrorMaker2InformerAndLister(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.KafkaV1().KafkaMirrorMaker2s(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

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
	informer := factory.Kafka().V1().KafkaMirrorMaker2s()
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
	newResource := NewMM2()
	_, err = client.KafkaV1().KafkaMirrorMaker2s(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create resource: %s", err.Error())
	}

	// Wait for the resource to be added in the informer
	<-addedSignal

	// Get the resource
	actualResource, err := lister.KafkaMirrorMaker2s(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertMM2s(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedMM2(actualResource)
	_, err = client.KafkaV1().KafkaMirrorMaker2s(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update resource: %s", err.Error())
	}

	// Wait for resource to be updated in the informer
	<-updatedSignal

	// Get the resource
	actualResource, err = lister.KafkaMirrorMaker2s(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertMM2s(t, updatedResource, actualResource)

	// Delete the resource
	err = client.KafkaV1().KafkaMirrorMaker2s(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete resource: %s", err.Error())
	}

	// Wait for resource to be deleted
	<-deletedSignal

	// Check deletion
	actualResource, err = lister.KafkaMirrorMaker2s(NAMESPACE).Get(NAME)
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
