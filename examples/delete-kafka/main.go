package main

import (
	"context"
	"flag"
	strimziclient "github.com/scholzj/strimzi-go/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

func main() {
	var namespace *string
	var kubeconfig *string

	namespace = flag.String("namespace", "myproject", "(optional) Namespace that will be used")
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) Absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "Absolute path to the kubeconfig file")
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

	log.Print("Deleting the KafkaNodePool resource")
	err = client.KafkaV1beta2().KafkaNodePools(*namespace).Delete(context.TODO(), "mixed", metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}

	log.Print("Deleting the Kafka resource")
	err = client.KafkaV1beta2().Kafkas(*namespace).Delete(context.TODO(), "my-cluster", metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}
