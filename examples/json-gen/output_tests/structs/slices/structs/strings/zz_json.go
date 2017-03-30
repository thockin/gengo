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

func ast_strings_Ttest(obj *Ttest) (libjson.Value, error) {

	result := libjson.Object{}

	// F []struct{F1 string; F2 string; F3 string}
	{
		obj := &obj.F
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Slice_Struct_F1_string_F2_string_F3_string((*[]struct {
			F1 string
			F2 string
			F3 string
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

func ast_Slice_Struct_F1_string_F2_string_F3_string(obj *[]struct {
	F1 string
	F2 string
	F3 string
}) (libjson.Value, error) {

	get := func() ([]libjson.Value, error) {
		if *obj == nil {
			return nil, nil
		}
		result := []libjson.Value{}
		for i := range *obj {
			obj := &(*obj)[i]
			//FIXME: do any of these ACTUALLY return an error?
			jv, err := ast_Struct_F1_string_F2_string_F3_string((*struct {
				F1 string
				F2 string
				F3 string
			})(obj))
			if err != nil {
				return nil, err
			}
			result = append(result, jv)
		}
		return result, nil
	}
	add := func() libjson.Value {
		var x struct {
			F1 string
			F2 string
			F3 string
		}
		*obj = append(*obj, x)
		obj := &(*obj)[len(*obj)-1]
		jv, _ := ast_Struct_F1_string_F2_string_F3_string((*struct {
			F1 string
			F2 string
			F3 string
		})(obj))
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
		*obj = []struct {
			F1 string
			F2 string
			F3 string
		}{}
		return libjson.NewArray(get, add), nil
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Struct_F1_string_F2_string_F3_string(obj *struct {
	F1 string
	F2 string
	F3 string
}) (libjson.Value, error) {

	result := libjson.Object{}

	// F1 string
	{
		obj := &obj.F1
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_string((*string)(obj))
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

	// F2 string
	{
		obj := &obj.F2
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_string((*string)(obj))
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

	// F3 string
	{
		obj := &obj.F3
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_string((*string)(obj))
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

func ast_string(obj *string) (libjson.Value, error) {
	return libjson.NewString(func() string { return *obj }, func(s string) { *obj = s }), nil
}
