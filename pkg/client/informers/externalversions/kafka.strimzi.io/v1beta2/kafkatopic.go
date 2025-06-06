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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta2

import (
	context "context"
	time "time"

	apiskafkastrimziiov1beta2 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1beta2"
	versioned "github.com/scholzj/strimzi-go/pkg/client/clientset/versioned"
	internalinterfaces "github.com/scholzj/strimzi-go/pkg/client/informers/externalversions/internalinterfaces"
	kafkastrimziiov1beta2 "github.com/scholzj/strimzi-go/pkg/client/listers/kafka.strimzi.io/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KafkaTopicInformer provides access to a shared informer and lister for
// KafkaTopics.
type KafkaTopicInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() kafkastrimziiov1beta2.KafkaTopicLister
}

type kafkaTopicInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKafkaTopicInformer constructs a new informer for KafkaTopic type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKafkaTopicInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKafkaTopicInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKafkaTopicInformer constructs a new informer for KafkaTopic type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKafkaTopicInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KafkaV1beta2().KafkaTopics(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KafkaV1beta2().KafkaTopics(namespace).Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KafkaV1beta2().KafkaTopics(namespace).List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KafkaV1beta2().KafkaTopics(namespace).Watch(ctx, options)
			},
		},
		&apiskafkastrimziiov1beta2.KafkaTopic{},
		resyncPeriod,
		indexers,
	)
}

func (f *kafkaTopicInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKafkaTopicInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kafkaTopicInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiskafkastrimziiov1beta2.KafkaTopic{}, f.defaultInformer)
}

func (f *kafkaTopicInformer) Lister() kafkastrimziiov1beta2.KafkaTopicLister {
	return kafkastrimziiov1beta2.NewKafkaTopicLister(f.Informer().GetIndexer())
}
