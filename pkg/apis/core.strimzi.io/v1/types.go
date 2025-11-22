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

package v1

import (
	kafkav1 "github.com/scholzj/strimzi-go/pkg/apis/kafka.strimzi.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type StrimziPodSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []StrimziPodSet `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type StrimziPodSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   *StrimziPodSetSpec   `json:"spec,omitempty"`
	Status *StrimziPodSetStatus `json:"status,omitempty"`
}

type StrimziPodSetStatus struct {
	Conditions         []Condition `json:"conditions,omitempty"`
	ObservedGeneration int64       `json:"observedGeneration,omitempty"`
	Pods               int32       `json:"pods,omitempty"`
	ReadyPods          int32       `json:"readyPods,omitempty"`
	CurrentPods        int32       `json:"currentPods,omitempty"`
}

type Condition struct {
	Type               string `json:"type,omitempty"`
	Status             string `json:"status,omitempty"`
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	Reason             string `json:"reason,omitempty"`
	Message            string `json:"message,omitempty"`
}

type StrimziPodSetSpec struct {
	Selector *metav1.LabelSelector     `json:"selector,omitempty"`
	Pods     []kafkav1.MapStringObject `json:"pods,omitempty"`
}
