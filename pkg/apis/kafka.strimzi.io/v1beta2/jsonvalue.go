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

// A new type used to store Map<String, Object> fields in the CRDs.
// This is needed because Golang CRD structures and generators don't support using map[string]interface{} directly

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
