package serialization

import (
	kafkav1beta2 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1beta2"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"
	"testing"
)

type Resource interface {
	kafkav1beta2.KafkaTopic | kafkav1beta2.Kafka | kafkav1beta2.KafkaUser | kafkav1beta2.KafkaNodePool | kafkav1beta2.KafkaConnect | kafkav1beta2.KafkaBridge | kafkav1beta2.KafkaMirrorMaker2 | kafkav1beta2.KafkaConnector | kafkav1beta2.KafkaRebalance
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
	var resource kafkav1beta2.Kafka
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafka(t *testing.T) {
	var resource kafkav1beta2.Kafka
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaBridge(t *testing.T) {
	var resource kafkav1beta2.KafkaBridge
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaBridge(t *testing.T) {
	var resource kafkav1beta2.KafkaBridge
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaConnect(t *testing.T) {
	var resource kafkav1beta2.KafkaConnect
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaConnect(t *testing.T) {
	var resource kafkav1beta2.KafkaConnect
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaConnector(t *testing.T) {
	var resource kafkav1beta2.KafkaConnector
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaConnector(t *testing.T) {
	var resource kafkav1beta2.KafkaConnector
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaMirrorMaker2(t *testing.T) {
	var resource kafkav1beta2.KafkaMirrorMaker2
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaMirrorMaker2(t *testing.T) {
	var resource kafkav1beta2.KafkaMirrorMaker2
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaNodePool(t *testing.T) {
	var resource kafkav1beta2.KafkaNodePool
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaNodePool(t *testing.T) {
	var resource kafkav1beta2.KafkaNodePool
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaRebalance(t *testing.T) {
	var resource kafkav1beta2.KafkaRebalance
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaRebalance(t *testing.T) {
	var resource kafkav1beta2.KafkaRebalance
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaTopic(t *testing.T) {
	var resource kafkav1beta2.KafkaTopic
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaTopic(t *testing.T) {
	var resource kafkav1beta2.KafkaTopic
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestSimpleKafkaUser(t *testing.T) {
	var resource kafkav1beta2.KafkaUser
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}

func TestComplexKafkaUser(t *testing.T) {
	var resource kafkav1beta2.KafkaUser
	LoadUnload(t, "./resources/"+t.Name()+".yaml", resource)
}
