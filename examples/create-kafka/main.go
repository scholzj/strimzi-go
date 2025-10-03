package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"
	"time"

	kafkav1beta2 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1beta2"
	strimziclient "github.com/scholzj/strimzi-go/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var namespace *string
	var kubeconfig *string

	namespace = flag.String("namespace", "myproject", "(optional) Namespace that will be used")
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	client, err := strimziclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	nodepool := &kafkav1beta2.KafkaNodePool{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "mixed",
			Labels: map[string]string{"strimzi.io/cluster": "my-cluster"},
		},
		Spec: &kafkav1beta2.KafkaNodePoolSpec{
			Replicas: 3,
			Roles:    []kafkav1beta2.ProcessRoles{kafkav1beta2.BROKER_PROCESSROLES, kafkav1beta2.CONTROLLER_PROCESSROLES},
			Storage: &kafkav1beta2.Storage{
				Type: kafkav1beta2.EPHEMERAL_STORAGETYPE,
			},
		},
	}

	log.Print("Creating the KafkaNodePool resource")
	nodepool, err = client.KafkaV1beta2().KafkaNodePools(*namespace).Create(context.TODO(), nodepool, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	kafka := &kafkav1beta2.Kafka{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "my-cluster",
			Annotations: map[string]string{"strimzi.io/kraft": "enabled", "strimzi.io/node-pools": "enabled"},
		},
		Spec: &kafkav1beta2.KafkaSpec{
			Kafka: &kafkav1beta2.KafkaClusterSpec{
				Listeners: []kafkav1beta2.GenericKafkaListener{{
					Name: "internal",
					Type: kafkav1beta2.INTERNAL_KAFKALISTENERTYPE,
					Tls:  false,
					Port: 9092,
				}},
			},
			EntityOperator: &kafkav1beta2.EntityOperatorSpec{
				TopicOperator: &kafkav1beta2.EntityTopicOperatorSpec{},
				UserOperator:  &kafkav1beta2.EntityUserOperatorSpec{},
			},
		},
	}

	log.Print("Creating the Kafka resource")
	kafka, err = client.KafkaV1beta2().Kafkas(*namespace).Create(context.TODO(), kafka, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	log.Print("Waiting for the Kafka cluster to get ready")
	ready := waitUntilReady(client, namespace)
	if ready {
		log.Print("The Kafka cluster is ready")
	} else {
		log.Fatal("The Kafka cluster is NOT ready")
	}
}

func waitUntilReady(client *strimziclient.Clientset, namespace *string) bool {
	watchContext, watchContextCancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer watchContextCancel()

	watcher, err := client.KafkaV1beta2().Kafkas(*namespace).Watch(watchContext, metav1.ListOptions{FieldSelector: fields.OneTermEqualSelector(metav1.ObjectNameField, "my-cluster").String()})
	if err != nil {
		panic(err)
	}

	defer watcher.Stop()

	for {
		select {
		case event := <-watcher.ResultChan():
			if isReady(event) {
				return true
			}
		case <-watchContext.Done():
			log.Print("Timed out waiting for the Kafka cluster to be ready")
			return false
		}
	}
}

func isReady(event watch.Event) bool {
	k := event.Object.(*kafkav1beta2.Kafka)
	if k.Status != nil && k.Status.Conditions != nil && len(k.Status.Conditions) > 0 {
		for _, condition := range k.Status.Conditions {
			if condition.Type == "Ready" && condition.Status == "True" {
				//log.Print("The Kafka cluster is ready")
				return true
			}
		}

		//log.Print("The Kafka cluster has conditions but is not ready")
		return false
	} else {
		//log.Print("The Kafka cluster has no conditions")
		return false
	}
}
