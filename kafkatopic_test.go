package v1beta2

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	kafkav1beta2 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1beta2"
	strimzi "github.com/scholzj/strimzi-go/pkg/client/clientset/versioned"
)

func TestKafkaTopicTest(t *testing.T) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	strimziClient, err := strimzi.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	topics, err := strimziClient.KafkaV1beta2().KafkaTopics("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err.Error())
		t.FailNow()
	}
	fmt.Printf("There are %d topics in the cluster\n", len(topics.Items))

	var topic *kafkav1beta2.KafkaTopic
	topic, err = strimziClient.KafkaV1beta2().KafkaTopics("myproject").Get(context.TODO(), "kafka-test-apps", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err.Error())
		t.FailNow()
	}
	fmt.Printf("Topic %s has %d replicas and %d partitions\n", topic.Name, topic.Spec.Replicas, topic.Spec.Partitions)
}
