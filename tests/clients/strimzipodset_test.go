package clients

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"

	corev1beta2 "github.com/scholzj/strimzi-go/pkg/apis/core.strimzi.io/v1beta2"
	kafkav1beta2 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1beta2"
	strimziinformer "github.com/scholzj/strimzi-go/pkg/client/informers/externalversions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func NewPodSet() *corev1beta2.StrimziPodSet {
	return &corev1beta2.StrimziPodSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      NAME,
			Namespace: NAMESPACE,
		},
		Spec: &corev1beta2.StrimziPodSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"strimzi.io/label": "value",
				},
			},
			Pods: []kafkav1beta2.MapStringObject{{
				"apiVersion": "v1",
				"kind":       "Pod",
				"metadata": map[string]interface{}{
					"labels": map[string]interface{}{
						"strimzi.io/label": "value",
					},
					"name": "my-pod-set-1",
				},
				"spec": map[string]interface{}{
					"containers": []interface{}{
						map[string]interface{}{
							"image": "nginx:1.14.0",
							"name":  "nginx",
							"ports": []interface{}{
								map[string]interface{}{
									"containerPort": int64(80),
								},
							},
						},
					},

					//"containers": []map[string]interface{}{{
					//	"image": "nginx:1.14.0",
					//	"name":  "nginx",
					//	"ports": []map[string]interface{}{{
					//		"containerPort": int64(80),
					//	}},
					//}},
				},
			}},
		},
	}
}

func UpdatedPodSet(oldResource *corev1beta2.StrimziPodSet) *corev1beta2.StrimziPodSet {
	updatedResource := oldResource.DeepCopy()

	updatedResource.Spec.Selector.MatchLabels["strimzi.io/label"] = "value2"

	updatedPods := []kafkav1beta2.MapStringObject{{
		"apiVersion": "v1",
		"kind":       "Pod",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"strimzi.io/label": "value2",
			},
			"name": "my-pod-set-1",
		},
		"spec": map[string]interface{}{
			"containers": []interface{}{
				map[string]interface{}{
					"image": "nginx:1.14.2",
					"name":  "nginx",
					"ports": []interface{}{
						map[string]interface{}{
							"containerPort": int64(80),
						},
					},
				},
			},
		},
	}}

	updatedResource.Spec.Pods = updatedPods

	return updatedResource
}

func AssertPodSets(t *testing.T, r1, r2 *corev1beta2.StrimziPodSet) {
	assert.Equal(t, r1.ObjectMeta.Name, r2.ObjectMeta.Name)
	assert.Equal(t, r1.ObjectMeta.Namespace, r2.ObjectMeta.Namespace)

	assert.Equal(t, r1.Spec.Selector.MatchLabels, r2.Spec.Selector.MatchLabels)
	assert.Equal(t, len(r1.Spec.Pods), len(r2.Spec.Pods))
	assert.Equal(t, r1.Spec.Pods, r2.Spec.Pods)
}

func TestPodSetCreateUpdateDelete(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.CoreV1beta2().StrimziPodSets(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

	// Create the resource
	newResource := NewPodSet()
	_, err := client.CoreV1beta2().StrimziPodSets(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err := client.CoreV1beta2().StrimziPodSets(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertPodSets(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedPodSet(actualResource)
	_, err = client.CoreV1beta2().StrimziPodSets(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update actualResource: %s", err.Error())
	}

	// Get the resource
	actualResource, err = client.CoreV1beta2().StrimziPodSets(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get actualResource: %s", err.Error())
	}

	// Assert the resource
	AssertPodSets(t, updatedResource, actualResource)

	// Delete the resource
	err = client.CoreV1beta2().StrimziPodSets(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete actualResource: %s", err.Error())
	}

	// Check deletion
	actualResource, err = client.CoreV1beta2().StrimziPodSets(NAMESPACE).Get(context.TODO(), NAME, metav1.GetOptions{})
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			t.Fatalf("Failed to get resource: %s", err.Error())
		}
	} else if actualResource != nil {
		t.Fatalf("Resource still exists")
	}
}

func TestPodSetInformerAndLister(t *testing.T) {
	client := Client(t)

	// Delete at the end to avoid errors
	defer client.CoreV1beta2().StrimziPodSets(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})

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
	informer := factory.Core().V1beta2().StrimziPodSets()
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
	newResource := NewPodSet()
	_, err = client.CoreV1beta2().StrimziPodSets(NAMESPACE).Create(context.TODO(), newResource, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create resource: %s", err.Error())
	}

	// Wait for the resource to be added in the informer
	<-addedSignal

	// Get the resource
	actualResource, err := lister.StrimziPodSets(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertPodSets(t, newResource, actualResource)

	// Update the resource
	updatedResource := UpdatedPodSet(actualResource)
	_, err = client.CoreV1beta2().StrimziPodSets(NAMESPACE).Update(context.TODO(), updatedResource, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update resource: %s", err.Error())
	}

	// Wait for resource to be updated in the informer
	<-updatedSignal

	// Get the resource
	actualResource, err = lister.StrimziPodSets(NAMESPACE).Get(NAME)
	if err != nil {
		t.Fatalf("Failed to get resource: %s", err.Error())
	}

	// Assert the resource
	AssertPodSets(t, updatedResource, actualResource)

	// Delete the resource
	err = client.CoreV1beta2().StrimziPodSets(NAMESPACE).Delete(context.TODO(), NAME, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete resource: %s", err.Error())
	}

	// Wait for resource to be deleted
	<-deletedSignal

	// Check deletion
	actualResource, err = lister.StrimziPodSets(NAMESPACE).Get(NAME)
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
