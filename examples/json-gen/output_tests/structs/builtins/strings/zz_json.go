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
	jv, err := ast_strings_Ttest(&obj)
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
	jv, err := ast_strings_Ttest(obj)
	if err != nil {
		return err
	}
	return jv.Parse(data)
}

func ast_strings_Ttest(obj *Ttest) (libjson.Value, error) {

	result := libjson.NewObject()

	// F1 string
	{
		obj := &obj.F1

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_string((*string)(obj))
		if err != nil {
			return nil, err
		}
		fv, err := finalize(jv)
		if err != nil {
			return nil, err
		}
		p := new(string)
		*p = "F1"
		nv := libjson.NamedValue{
			Name:      libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
			Value:     fv,
			OmitEmpty: false,
		}
		result = append(result, nv)
	}

	// F2 string
	{
		obj := &obj.F2

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_string((*string)(obj))
		if err != nil {
			return nil, err
		}
		fv, err := finalize(jv)
		if err != nil {
			return nil, err
		}
		p := new(string)
		*p = "F2"
		nv := libjson.NamedValue{
			Name:      libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
			Value:     fv,
			OmitEmpty: false,
		}
		result = append(result, nv)
	}

	// F3 string
	{
		obj := &obj.F3

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_string((*string)(obj))
		if err != nil {
			return nil, err
		}
		fv, err := finalize(jv)
		if err != nil {
			return nil, err
		}
		p := new(string)
		*p = "F3"
		nv := libjson.NamedValue{
			Name:      libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
			Value:     fv,
			OmitEmpty: false,
		}
		result = append(result, nv)
	}

	return result, nil

}

func ast_string(obj *string) (libjson.Value, error) {
	return libjson.NewString(func() string { return *obj }, func(s string) { *obj = s }), nil
}
