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

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KafkaTopic `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KafkaTopicSpec `json:"spec"`

	// +optional
	Status KafkaTopicStatus `json:"status,omitempty"`
}

type KafkaTopicSpec struct {
	Partitions int32 `json:"partitions,omitempty"`

	Replicas int32 `json:"replicas,omitempty"`

	// +optional
	TopicName string `json:"topicName,omitempty"`

	// +optional
	Config JSONValue `json:"config,omitempty"`
}

type KafkaTopicStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`

	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TopicName string `json:"topicName,omitempty"`
}

type JSONValue map[string]interface{}

func (in *JSONValue) DeepCopy() *JSONValue {
	if in == nil {
		return nil
	}

	out := new(JSONValue)
	in.DeepCopyInto(out)

	return out
}

func (in *JSONValue) DeepCopyInto(out *JSONValue) {
	if in != nil {
		*out = make(map[string]interface{}, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}

	return
}
