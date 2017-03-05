// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by json-gen. Do not edit it manually!

package test

import (
	bytes "bytes"
	libjson "k8s.io/gengo/examples/json-gen/libjson"
)

func init() {
}

func ast_struct_empty_T(obj *T) (libjson.Value, error) {

	result := libjson.Object{}
	_ = obj

	return result, nil

}
func Marshal_struct_empty_T(obj T, buf *bytes.Buffer) error {
	val, err := ast_struct_empty_T(&obj)
	if err != nil {
		return err
	}
	return val.Render(buf)
}
func Unmarshal_struct_empty_T(data []byte, obj *T) error {
	val, err := ast_struct_empty_T(obj)
	if err != nil {
		return err
	}
	return val.Parse(data)
}
