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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta2 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1beta2"
	kafkastrimziiov1beta2 "github.com/scholzj/strimzi-go/pkg/client/clientset/versioned/typed/kafka.strimzi.io/v1beta2"
	gentype "k8s.io/client-go/gentype"
)

// fakeKafkaConnectors implements KafkaConnectorInterface
type fakeKafkaConnectors struct {
	*gentype.FakeClientWithList[*v1beta2.KafkaConnector, *v1beta2.KafkaConnectorList]
	Fake *FakeKafkaV1beta2
}

func newFakeKafkaConnectors(fake *FakeKafkaV1beta2, namespace string) kafkastrimziiov1beta2.KafkaConnectorInterface {
	return &fakeKafkaConnectors{
		gentype.NewFakeClientWithList[*v1beta2.KafkaConnector, *v1beta2.KafkaConnectorList](
			fake.Fake,
			namespace,
			v1beta2.SchemeGroupVersion.WithResource("kafkaconnectors"),
			v1beta2.SchemeGroupVersion.WithKind("KafkaConnector"),
			func() *v1beta2.KafkaConnector { return &v1beta2.KafkaConnector{} },
			func() *v1beta2.KafkaConnectorList { return &v1beta2.KafkaConnectorList{} },
			func(dst, src *v1beta2.KafkaConnectorList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta2.KafkaConnectorList) []*v1beta2.KafkaConnector {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1beta2.KafkaConnectorList, items []*v1beta2.KafkaConnector) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
