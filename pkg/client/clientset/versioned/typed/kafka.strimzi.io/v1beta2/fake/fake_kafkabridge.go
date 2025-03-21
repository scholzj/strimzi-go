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

// fakeKafkaBridges implements KafkaBridgeInterface
type fakeKafkaBridges struct {
	*gentype.FakeClientWithList[*v1beta2.KafkaBridge, *v1beta2.KafkaBridgeList]
	Fake *FakeKafkaV1beta2
}

func newFakeKafkaBridges(fake *FakeKafkaV1beta2, namespace string) kafkastrimziiov1beta2.KafkaBridgeInterface {
	return &fakeKafkaBridges{
		gentype.NewFakeClientWithList[*v1beta2.KafkaBridge, *v1beta2.KafkaBridgeList](
			fake.Fake,
			namespace,
			v1beta2.SchemeGroupVersion.WithResource("kafkabridges"),
			v1beta2.SchemeGroupVersion.WithKind("KafkaBridge"),
			func() *v1beta2.KafkaBridge { return &v1beta2.KafkaBridge{} },
			func() *v1beta2.KafkaBridgeList { return &v1beta2.KafkaBridgeList{} },
			func(dst, src *v1beta2.KafkaBridgeList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta2.KafkaBridgeList) []*v1beta2.KafkaBridge { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta2.KafkaBridgeList, items []*v1beta2.KafkaBridge) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
