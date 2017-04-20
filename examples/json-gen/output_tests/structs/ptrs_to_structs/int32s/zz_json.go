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
	jv, err := ast_int32s_Ttest(&obj)
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
	jv, err := ast_int32s_Ttest(obj)
	if err != nil {
		return err
	}
	return jv.Parse(data)
}

func ast_int32s_Ttest(obj *Ttest) (libjson.Value, error) {

	result := libjson.NewObject()

	// F *struct{F1 int32; F2 int32; F3 int32}
	{
		obj := &obj.F

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_Struct_F1_int32_F2_int32_F3_int32((**struct {
			F1 int32
			F2 int32
			F3 int32
		})(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
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
		} else {
			panic("TIM: F was empty")
		} //FIXME:
	}

	return result, nil

}

func ast_Pointer_Struct_F1_int32_F2_int32_F3_int32(obj **struct {
	F1 int32
	F2 int32
	F3 int32
}) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_Struct_F1_int32_F2_int32_F3_int32((*struct {
			F1 int32
			F2 int32
			F3 int32
		})(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(struct {
			F1 int32
			F2 int32
			F3 int32
		})
		obj := *obj
		return ast_Struct_F1_int32_F2_int32_F3_int32((*struct {
			F1 int32
			F2 int32
			F3 int32
		})(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Struct_F1_int32_F2_int32_F3_int32(obj *struct {
	F1 int32
	F2 int32
	F3 int32
}) (libjson.Value, error) {

	result := libjson.NewObject()

	// F1 int32
	{
		obj := &obj.F1

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_int32((*int32)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "F1"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: F1 was empty")
		} //FIXME:
	}

	// F2 int32
	{
		obj := &obj.F2

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_int32((*int32)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "F2"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: F2 was empty")
		} //FIXME:
	}

	// F3 int32
	{
		obj := &obj.F3

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_int32((*int32)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "F3"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: F3 was empty")
		} //FIXME:
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
