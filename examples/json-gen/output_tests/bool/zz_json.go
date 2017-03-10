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

func ast_bool_Ttest(obj *Ttest) (libjson.Value, error) {
	return ast_bool((*bool)(obj))
}

func (obj Ttest) MarshalJSON() ([]byte, error) {
	val, err := ast_bool_Ttest(&obj)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := val.Render(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (obj *Ttest) UnmarshalJSON(data []byte) error {
	val, err := ast_bool_Ttest(obj)
	if err != nil {
		return err
	}
	return val.Parse(data)
}

func ast_bool(obj *bool) (libjson.Value, error) {
	return libjson.NewBool(func() bool { return *obj }, func(b bool) { *obj = b }), nil
}