package clients

import (
	"path/filepath"
	"testing"

	strimzi "github.com/scholzj/strimzi-go/pkg/client/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const NAMESPACE = "default"
const NAME = "my-v1beta2-test-resource"

func Client(t *testing.T) *strimzi.Clientset {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeconfig = ""
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		t.Fatalf("Failed to create Kubernetes Client Config: %s", err)
	}

	client, err := strimzi.NewForConfig(config)
	if err != nil {
		t.Fatalf("Failed to create Kubernetes Client: %s", err)
	}

	return client
}
