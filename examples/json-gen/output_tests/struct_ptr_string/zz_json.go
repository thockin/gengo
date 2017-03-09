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

func ast_struct_ptr_string_T(obj *T) (libjson.Value, error) {

	result := libjson.Object{}

	// F
	{
		obj := &obj.F
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		val, err := func() (libjson.Value, error) {
			p := *obj
			if p == nil {
				p = new(string)
			}
			jv, err := ast_string((*string)(p))
			if err != nil {
				return nil, err
			}
			return libjson.NewOptional(jv, *obj != nil, func() { *obj = p }, func() { *obj = nil }), nil
		}()
		if err != nil {
			return nil, err
		}
		if !empty(val) {
			fv, err := finalize(val)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "F"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		}
	}

	return result, nil

}
func Marshal_struct_ptr_string_T(obj T) ([]byte, error) {
	val, err := ast_struct_ptr_string_T(&obj)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := val.Render(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func Unmarshal_struct_ptr_string_T(data []byte, obj *T) error {
	val, err := ast_struct_ptr_string_T(obj)
	if err != nil {
		return err
	}
	return val.Parse(data)
}

func ast_string(obj *string) (libjson.Value, error) {
	return libjson.NewString(func() string { return *obj }, func(s string) { *obj = s }), nil
}
func Marshal_string(obj string) ([]byte, error) {
	val, err := ast_string(&obj)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := val.Render(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func Unmarshal_string(data []byte, obj *string) error {
	val, err := ast_string(obj)
	if err != nil {
		return err
	}
	return val.Parse(data)
}
