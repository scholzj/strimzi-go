/*
Copyright 2025 Jakub Scholz.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta2

import (
	kafkastrimziiov1beta2 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1beta2"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// KafkaConnectorLister helps list KafkaConnectors.
// All objects returned here must be treated as read-only.
type KafkaConnectorLister interface {
	// List lists all KafkaConnectors in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kafkastrimziiov1beta2.KafkaConnector, err error)
	// KafkaConnectors returns an object that can list and get KafkaConnectors.
	KafkaConnectors(namespace string) KafkaConnectorNamespaceLister
	KafkaConnectorListerExpansion
}

// kafkaConnectorLister implements the KafkaConnectorLister interface.
type kafkaConnectorLister struct {
	listers.ResourceIndexer[*kafkastrimziiov1beta2.KafkaConnector]
}

// NewKafkaConnectorLister returns a new KafkaConnectorLister.
func NewKafkaConnectorLister(indexer cache.Indexer) KafkaConnectorLister {
	return &kafkaConnectorLister{listers.New[*kafkastrimziiov1beta2.KafkaConnector](indexer, kafkastrimziiov1beta2.Resource("kafkaconnector"))}
}

// KafkaConnectors returns an object that can list and get KafkaConnectors.
func (s *kafkaConnectorLister) KafkaConnectors(namespace string) KafkaConnectorNamespaceLister {
	return kafkaConnectorNamespaceLister{listers.NewNamespaced[*kafkastrimziiov1beta2.KafkaConnector](s.ResourceIndexer, namespace)}
}

// KafkaConnectorNamespaceLister helps list and get KafkaConnectors.
// All objects returned here must be treated as read-only.
type KafkaConnectorNamespaceLister interface {
	// List lists all KafkaConnectors in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kafkastrimziiov1beta2.KafkaConnector, err error)
	// Get retrieves the KafkaConnector from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kafkastrimziiov1beta2.KafkaConnector, error)
	KafkaConnectorNamespaceListerExpansion
}

// kafkaConnectorNamespaceLister implements the KafkaConnectorNamespaceLister
// interface.
type kafkaConnectorNamespaceLister struct {
	listers.ResourceIndexer[*kafkastrimziiov1beta2.KafkaConnector]
}
