package v1beta2

import (
	"context"
	"log"
	"path/filepath"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	kafkav1beta2 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1beta2"
	strimzi "github.com/scholzj/strimzi-go/pkg/client/clientset/versioned"
	strimziinformer "github.com/scholzj/strimzi-go/pkg/client/informers/externalversions"
)

const NAMESPACE = "default"

func GetConfig() (*rest.Config, error) {
	var kubeconfig string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeconfig = ""
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	return config, err
}

func TestKafkaTopicCreateUpdateDeleteTest(t *testing.T) {
	config, err := GetConfig()
	if err != nil {
		t.Fatal(err)
	}

	clientset, err := strimzi.NewForConfig(config)
	if err != nil {
		t.Fatal(err)
	}

	newTopic := &kafkav1beta2.KafkaTopic{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-test-topic",
		},
		Spec: &kafkav1beta2.KafkaTopicSpec{
			Replicas:   3,
			Partitions: 3,
			Config:     kafkav1beta2.JSONValue{"retention.ms": 7200000, "segment.bytes": 1073741824},
		},
	}

	// Create the topic
	_, err = clientset.KafkaV1beta2().KafkaTopics(NAMESPACE).Create(context.TODO(), newTopic, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create topic: %s", err.Error())
	}

	// Get the topic
	topic, err := clientset.KafkaV1beta2().KafkaTopics(NAMESPACE).Get(context.TODO(), "my-test-topic", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get topic: %s", err.Error())
	}

	// Assert the topic
	if topic.Spec.Partitions != 3 || topic.Spec.Replicas != 3 || topic.Spec.Config["retention.ms"] != int64(7200000) || topic.Spec.Config["segment.bytes"] != int64(1073741824) {
		t.Fatalf("Topic does not look correct: %v", topic)
	}

	// Update the topic
	updatedTopic := topic.DeepCopy()
	updatedTopic.Spec.Replicas = 1
	updatedTopic.Spec.Partitions = 10
	updatedTopic.Spec.Config["segment.bytes"] = 107374182

	_, err = clientset.KafkaV1beta2().KafkaTopics(NAMESPACE).Update(context.TODO(), updatedTopic, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update topic: %s", err.Error())
	}

	// Get the topic
	topic, err = clientset.KafkaV1beta2().KafkaTopics(NAMESPACE).Get(context.TODO(), "my-test-topic", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get topic: %s", err.Error())
	}

	// Assert the topic
	if topic.Spec.Partitions != 10 || topic.Spec.Replicas != 1 || topic.Spec.Config["retention.ms"] != int64(7200000) || topic.Spec.Config["segment.bytes"] != int64(107374182) {
		t.Fatalf("Topic does not look correct: %v", topic)
	}

	// Delete the topic
	err = clientset.KafkaV1beta2().KafkaTopics(NAMESPACE).Delete(context.TODO(), "my-test-topic", metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete topic: %s", err.Error())
	}

	// Check deletion
	topic, err = clientset.KafkaV1beta2().KafkaTopics(NAMESPACE).Get(context.TODO(), "my-test-topic", metav1.GetOptions{})
	if err != nil {
		if err.Error() != "kafkatopics.kafka.strimzi.io \"my-test-topic\" not found" {
			t.Fatalf("Failed to get topic: %s", err.Error())
		}
	} else if topic != nil {
		t.Fatalf("Topic still exists")
	}
}

func TestKafkaTopicInformerAndLister(t *testing.T) {
	config, err := GetConfig()
	if err != nil {
		t.Fatal(err)
	}

	clientset, err := strimzi.NewForConfig(config)
	if err != nil {
		t.Fatal(err)
	}

	added := 0
	addedSignal := make(chan struct{})
	defer close(addedSignal)
	onAdd := func(new interface{}) {
		newKt := new.(*kafkav1beta2.KafkaTopic)
		log.Printf("New topic %s in namespace %s", newKt.Name, newKt.Namespace)
		added++
		addedSignal <- struct{}{}
	}
	updatedSignal := make(chan struct{})
	defer close(updatedSignal)
	updated := 0
	onUpdate := func(old interface{}, new interface{}) {
		newKt := new.(*kafkav1beta2.KafkaTopic)
		oldKt := old.(*kafkav1beta2.KafkaTopic)
		log.Printf("Updated topic %s in namespace %s", newKt.Name, oldKt.Namespace)
		updated++
		updatedSignal <- struct{}{}
	}
	deletedSignal := make(chan struct{})
	defer close(deletedSignal)
	deleted := 0
	onDelete := func(old interface{}) {
		oldKt := old.(*kafkav1beta2.KafkaTopic)
		log.Printf("Deleted topic %s in namespace %s", oldKt.Name, oldKt.Namespace)
		deleted++
		deletedSignal <- struct{}{}
	}

	// Create informer and lister
	factory := strimziinformer.NewSharedInformerFactoryWithOptions(clientset, time.Hour*1)
	informer := factory.Kafka().V1beta2().KafkaTopics()
	_, err = informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: onUpdate,
		DeleteFunc: onDelete,
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

	// Create the topic
	newTopic := &kafkav1beta2.KafkaTopic{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-test-topic2",
		},
		Spec: &kafkav1beta2.KafkaTopicSpec{
			Replicas:   3,
			Partitions: 3,
			Config:     kafkav1beta2.JSONValue{"retention.ms": 7200000, "segment.bytes": 1073741824},
		},
	}

	_, err = clientset.KafkaV1beta2().KafkaTopics(NAMESPACE).Create(context.TODO(), newTopic, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create topic: %s", err.Error())
	}

	<-addedSignal

	// Get the topic
	topic, err := lister.KafkaTopics(NAMESPACE).Get("my-test-topic2")
	if err != nil {
		t.Fatalf("Failed to get topic: %s", err.Error())
	}

	// Assert the topic
	if topic.Spec.Partitions != 3 || topic.Spec.Replicas != 3 || topic.Spec.Config["retention.ms"] != int64(7200000) || topic.Spec.Config["segment.bytes"] != int64(1073741824) {
		t.Fatalf("Topic does not look correct: %v", topic)
	}

	// Update the topic
	updatedTopic := topic.DeepCopy()
	updatedTopic.Spec.Replicas = 1
	updatedTopic.Spec.Partitions = 10
	updatedTopic.Spec.Config["segment.bytes"] = 107374182

	_, err = clientset.KafkaV1beta2().KafkaTopics(NAMESPACE).Update(context.TODO(), updatedTopic, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update topic: %s", err.Error())
	}

	<-updatedSignal

	// Get the topic
	topic, err = lister.KafkaTopics(NAMESPACE).Get("my-test-topic2")
	if err != nil {
		t.Fatalf("Failed to get topic: %s", err.Error())
	}

	// Assert the topic
	if topic.Spec.Partitions != 10 || topic.Spec.Replicas != 1 || topic.Spec.Config["retention.ms"] != int64(7200000) || topic.Spec.Config["segment.bytes"] != int64(107374182) {
		t.Fatalf("Updated topic does not look correct: %v", topic)
	}

	// Delete the topic
	err = clientset.KafkaV1beta2().KafkaTopics(NAMESPACE).Delete(context.TODO(), "my-test-topic2", metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete topic: %s", err.Error())
	}

	<-deletedSignal

	// Check deletion
	topic, err = lister.KafkaTopics(NAMESPACE).Get("my-test-topic2")
	if err != nil {
		if err.Error() != "kafkatopics.kafka.strimzi.io \"my-test-topic2\" not found" {
			// pass
		} else {
			t.Fatalf("Failed to get topic: %s", err.Error())
		}
	} else if topic != nil {
		t.Fatalf("Topic still exists")
	}

	// Assert the event handled
	if added != 1 {
		t.Fatalf("Topic was not added once but %d", added)
	}
	if updated != 1 {
		t.Fatalf("Topic was not updated once but %d", updated)
	}
	if deleted != 1 {
		t.Fatalf("Topic was not deleted once but %d", deleted)
	}
}
