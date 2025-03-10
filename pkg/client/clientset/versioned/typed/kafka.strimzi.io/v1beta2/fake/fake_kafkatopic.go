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

// fakeKafkaTopics implements KafkaTopicInterface
type fakeKafkaTopics struct {
	*gentype.FakeClientWithList[*v1beta2.KafkaTopic, *v1beta2.KafkaTopicList]
	Fake *FakeKafkaV1beta2
}

func newFakeKafkaTopics(fake *FakeKafkaV1beta2, namespace string) kafkastrimziiov1beta2.KafkaTopicInterface {
	return &fakeKafkaTopics{
		gentype.NewFakeClientWithList[*v1beta2.KafkaTopic, *v1beta2.KafkaTopicList](
			fake.Fake,
			namespace,
			v1beta2.SchemeGroupVersion.WithResource("kafkatopics"),
			v1beta2.SchemeGroupVersion.WithKind("KafkaTopic"),
			func() *v1beta2.KafkaTopic { return &v1beta2.KafkaTopic{} },
			func() *v1beta2.KafkaTopicList { return &v1beta2.KafkaTopicList{} },
			func(dst, src *v1beta2.KafkaTopicList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta2.KafkaTopicList) []*v1beta2.KafkaTopic { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta2.KafkaTopicList, items []*v1beta2.KafkaTopic) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
