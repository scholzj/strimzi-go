package serialization

import (
	"os"
	"path/filepath"
	"testing"

	corev1 "github.com/scholzj/strimzi-go/pkg/apis/core.strimzi.io/v1"
	kafkav1 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

type Resource interface {
	kafkav1.KafkaTopic | kafkav1.Kafka | kafkav1.KafkaUser | kafkav1.KafkaNodePool | kafkav1.KafkaConnect | kafkav1.KafkaBridge | kafkav1.KafkaMirrorMaker2 | kafkav1.KafkaConnector | kafkav1.KafkaRebalance | corev1.StrimziPodSet
}

func LoadUnload[R Resource](t *testing.T, filename string, resource R) {
	file, _ := filepath.Abs(filename)
	fileYaml, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("Failed to read yaml file: %s", err.Error())
	}

	err = yaml.Unmarshal(fileYaml, &resource)
	if err != nil {
		t.Fatalf("Failed to unmarshal the yaml: %s", err.Error())
	}

	serializedYaml, err := yaml.Marshal(resource)
	if err != nil {
		t.Fatalf("Failed to marshal the yaml: %s", err.Error())
	}

	assert.Equal(t, string(fileYaml), string(serializedYaml))
}

func TestSimpleKafka(t *testing.T) {
	var resource kafkav1.Kafka
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafka(t *testing.T) {
	var resource kafkav1.Kafka
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaBridge(t *testing.T) {
	var resource kafkav1.KafkaBridge
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaBridge(t *testing.T) {
	var resource kafkav1.KafkaBridge
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaConnect(t *testing.T) {
	var resource kafkav1.KafkaConnect
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaConnect(t *testing.T) {
	var resource kafkav1.KafkaConnect
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaConnector(t *testing.T) {
	var resource kafkav1.KafkaConnector
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaConnector(t *testing.T) {
	var resource kafkav1.KafkaConnector
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaMirrorMaker2(t *testing.T) {
	var resource kafkav1.KafkaMirrorMaker2
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaMirrorMaker2(t *testing.T) {
	var resource kafkav1.KafkaMirrorMaker2
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaNodePool(t *testing.T) {
	var resource kafkav1.KafkaNodePool
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaNodePool(t *testing.T) {
	var resource kafkav1.KafkaNodePool
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaRebalance(t *testing.T) {
	var resource kafkav1.KafkaRebalance
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaRebalance(t *testing.T) {
	var resource kafkav1.KafkaRebalance
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaTopic(t *testing.T) {
	var resource kafkav1.KafkaTopic
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaTopic(t *testing.T) {
	var resource kafkav1.KafkaTopic
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaUser(t *testing.T) {
	var resource kafkav1.KafkaUser
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaUser(t *testing.T) {
	var resource kafkav1.KafkaUser
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleStrimziPodSet(t *testing.T) {
	var resource corev1.StrimziPodSet
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}
