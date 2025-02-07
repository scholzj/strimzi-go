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
    corev1 "k8s.io/api/core/v1"
    networkingv1 "k8s.io/api/networking/v1"
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
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    
    Spec *KafkaTopicSpec `json:"spec,omitempty"`
    Status *KafkaTopicStatus `json:"status,omitempty"`
}

type KafkaTopicStatus struct {
    Conditions []Condition `json:"conditions,omitempty"`
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`
    TopicName string `json:"topicName,omitempty"`
    TopicId string `json:"topicId,omitempty"`
    ReplicasChange *ReplicasChangeStatus `json:"replicasChange,omitempty"`
}

type ReplicasChangeStatus struct {
    TargetReplicas int32 `json:"targetReplicas,omitempty"`
    State ReplicasChangeState `json:"state,omitempty"`
    Message string `json:"message,omitempty"`
    SessionId string `json:"sessionId,omitempty"`
}

type ReplicasChangeState string
const (
    PENDING_REPLICASCHANGESTATE ReplicasChangeState = "pending"
    ONGOING_REPLICASCHANGESTATE ReplicasChangeState = "ongoing"
)

type Condition struct {
    Type string `json:"type,omitempty"`
    Status string `json:"status,omitempty"`
    LastTransitionTime string `json:"lastTransitionTime,omitempty"`
    Reason string `json:"reason,omitempty"`
    Message string `json:"message,omitempty"`
}

type KafkaTopicSpec struct {
    TopicName string `json:"topicName,omitempty"`
    Partitions int32 `json:"partitions,omitempty"`
    Replicas int32 `json:"replicas,omitempty"`
    Config JSONValue `json:"config,omitempty"`
}
