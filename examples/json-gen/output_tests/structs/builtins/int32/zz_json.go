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

func (obj Ttest) MarshalJSON() ([]byte, error) {
	jv, err := ast_int32_Ttest(&obj)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := jv.Render(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (obj *Ttest) UnmarshalJSON(data []byte) error {
	jv, err := ast_int32_Ttest(obj)
	if err != nil {
		return err
	}
	return jv.Parse(data)
}

func ast_int32_Ttest(obj *Ttest) (libjson.Value, error) {

	result := libjson.NewObject()

	// F int32
	{
		obj := &obj.F

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_int32((*int32)(obj))
		if err != nil {
			return nil, err
		}
		fv, err := finalize(jv)
		if err != nil {
			return nil, err
		}
		p := new(string)
		*p = "F"
		nv := libjson.NamedValue{
			Name:      libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
			Value:     fv,
			OmitEmpty: false,
		}
		result = append(result, nv)
	}

	return result, nil

}

func ast_int32(obj *int32) (libjson.Value, error) {

	get := func() float64 {
		return float64(*obj)
	}
	set := func(f float64) {
		*obj = int32(f)
	}
	return libjson.NewInt(get, set), nil

}
