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

func ast_float32_Ttest(obj *Ttest) (libjson.Value, error) {

	result := libjson.Object{}

	// F []*float32
	{
		obj := &obj.F
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Slice_Pointer_float32((*[]*float32)(obj))
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

func (obj Ttest) MarshalJSON() ([]byte, error) {
	jv, err := ast_float32_Ttest(&obj)
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
	jv, err := ast_float32_Ttest(obj)
	if err != nil {
		return err
	}
	return jv.Parse(data)
}

func ast_Slice_Pointer_float32(obj *[]*float32) (libjson.Value, error) {

	get := func() ([]libjson.Value, error) {
		if *obj == nil {
			return nil, nil
		}
		result := []libjson.Value{}
		for i := range *obj {
			obj := &(*obj)[i]
			//FIXME: do any of these ACTUALLY return an error?
			jv, err := ast_Pointer_float32((**float32)(obj))
			if err != nil {
				return nil, err
			}
			result = append(result, jv)
		}
		return result, nil
	}
	add := func() libjson.Value {
		var x *float32
		*obj = append(*obj, x)
		obj := &(*obj)[len(*obj)-1]
		jv, _ := ast_Pointer_float32((**float32)(obj))
		//FIXME: handle error?
		return jv
	}
	var jv libjson.Value
	if *obj != nil {
		return libjson.NewArray(get, add), nil
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = []*float32{}
		return libjson.NewArray(get, add), nil
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Pointer_float32(obj **float32) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_float32((*float32)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(float32)
		obj := *obj
		return ast_float32((*float32)(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_float32(obj *float32) (libjson.Value, error) {

	get := func() float64 {
		return float64(*obj)
	}
	set := func(f float64) {
		*obj = float32(f)
	}
	return libjson.NewFloat(32, get, set), nil

}
