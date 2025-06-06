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

	apiscorestrimziiov1beta2 "github.com/scholzj/strimzi-go/pkg/apis/core.strimzi.io/v1beta2"
	versioned "github.com/scholzj/strimzi-go/pkg/client/clientset/versioned"
	internalinterfaces "github.com/scholzj/strimzi-go/pkg/client/informers/externalversions/internalinterfaces"
	corestrimziiov1beta2 "github.com/scholzj/strimzi-go/pkg/client/listers/core.strimzi.io/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StrimziPodSetInformer provides access to a shared informer and lister for
// StrimziPodSets.
type StrimziPodSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corestrimziiov1beta2.StrimziPodSetLister
}

type strimziPodSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewStrimziPodSetInformer constructs a new informer for StrimziPodSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStrimziPodSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStrimziPodSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredStrimziPodSetInformer constructs a new informer for StrimziPodSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStrimziPodSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta2().StrimziPodSets(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta2().StrimziPodSets(namespace).Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta2().StrimziPodSets(namespace).List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta2().StrimziPodSets(namespace).Watch(ctx, options)
			},
		},
		&apiscorestrimziiov1beta2.StrimziPodSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *strimziPodSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStrimziPodSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *strimziPodSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiscorestrimziiov1beta2.StrimziPodSet{}, f.defaultInformer)
}

func (f *strimziPodSetInformer) Lister() corestrimziiov1beta2.StrimziPodSetLister {
	return corestrimziiov1beta2.NewStrimziPodSetLister(f.Informer().GetIndexer())
}
