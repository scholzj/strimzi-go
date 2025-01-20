package v1beta2

import (
	"context"
	"log"
	"path/filepath"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	kafkav1beta2 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1beta2"
	strimzi "github.com/scholzj/strimzi-go/pkg/client/clientset/versioned"
)

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
		log.Fatal(err)
		t.FailNow()
	}

	clientset, err := strimzi.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	newTopic := &kafkav1beta2.KafkaTopic{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-test-topic",
		},
		Spec: kafkav1beta2.KafkaTopicSpec{
			Replicas:   3,
			Partitions: 3,
			Config:     map[string]string{"retention.ms": "7200000", "segment.bytes": "1073741824"},
		},
	}

	// Create the topic
	_, err = clientset.KafkaV1beta2().KafkaTopics("myproject").Create(context.TODO(), newTopic, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create topic: %s", err.Error())
	}

	// Get the topic
	topic, err := clientset.KafkaV1beta2().KafkaTopics("myproject").Get(context.TODO(), "my-test-topic", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get topic: %s", err.Error())
	}

	// Assert the topic
	if topic.Spec.Partitions != 3 || topic.Spec.Replicas != 3 || topic.Spec.Config["retention.ms"] != "7200000" || topic.Spec.Config["segment.bytes"] != "1073741824" {
		t.Fatalf("Topic does not look correct: %v", topic)
	}

	// Update the topic
	updatedTopic := topic.DeepCopy()
	updatedTopic.Spec.Replicas = 1
	updatedTopic.Spec.Partitions = 10
	updatedTopic.Spec.Config["segment.bytes"] = "107374182"

	_, err = clientset.KafkaV1beta2().KafkaTopics("myproject").Update(context.TODO(), updatedTopic, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("Failed to update topic: %s", err.Error())
	}

	// Get the topic
	topic, err = clientset.KafkaV1beta2().KafkaTopics("myproject").Get(context.TODO(), "my-test-topic", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to get topic: %s", err.Error())
	}

	// Assert the topic
	if topic.Spec.Partitions != 10 || topic.Spec.Replicas != 1 || topic.Spec.Config["retention.ms"] != "7200000" || topic.Spec.Config["segment.bytes"] != "107374182" {
		t.Fatalf("Topic does not look correct: %v", topic)
	}

	// Delete the topic
	err = clientset.KafkaV1beta2().KafkaTopics("myproject").Delete(context.TODO(), "my-test-topic", metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete topic: %s", err.Error())
	}

	// Check deletion
	topic, err = clientset.KafkaV1beta2().KafkaTopics("myproject").Get(context.TODO(), "my-test-topic", metav1.GetOptions{})
	if err != nil {
		if err.Error() != "kafkatopics.kafka.strimzi.io \"my-test-topic\" not found" {
			t.Fatalf("Failed to get topic: %s", err.Error())
		}
	} else if topic != nil {
		t.Fatalf("Topic still exists")
	}
}
