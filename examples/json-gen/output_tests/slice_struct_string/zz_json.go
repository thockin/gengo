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

func ast_slice_struct_string_Ttest(obj *Ttest) (libjson.Value, error) {

	get := func() ([]libjson.Value, error) {
		if *obj == nil {
			return nil, nil
		}
		result := []libjson.Value{}
		for i := range *obj {
			obj := &(*obj)[i]
			//FIXME: do any of these ACTUALLY return an error?
			val, err := func() (libjson.Value, error) {
				result := libjson.Object{}

				// F string
				{
					obj := &obj.F
					_ = obj //FIXME: remove when other Kinds are done

					empty := func(libjson.Value) bool { return false }

					finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

					val, err := func() (libjson.Value, error) { return ast_string((*string)(obj)) }()

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
					} else {
						panic("TIM: F was empty")
					} //FIXME:
				}

				return result, nil
			}()
			if err != nil {
				return nil, err
			}
			result = append(result, val)
		}
		return result, nil
	}
	add := func() libjson.Value {
		var x struct{ F string }
		*obj = append(*obj, x)
		obj := &(*obj)[len(*obj)-1]
		val, _ := func() (libjson.Value, error) {
			result := libjson.Object{}

			// F string
			{
				obj := &obj.F
				_ = obj //FIXME: remove when other Kinds are done

				empty := func(libjson.Value) bool { return false }

				finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

				val, err := func() (libjson.Value, error) { return ast_string((*string)(obj)) }()

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
				} else {
					panic("TIM: F was empty")
				} //FIXME:
			}

			return result, nil
		}()
		//FIXME: handle error?
		return val
	}
	setNull := func(b bool) {
		if b {
			*obj = nil
		} else {
			*obj = []struct{ F string }{}
		}
	}
	return libjson.NewArray(*obj == nil, get, add, setNull), nil

}

func (obj Ttest) MarshalJSON() ([]byte, error) {
	val, err := ast_slice_struct_string_Ttest(&obj)
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
	val, err := ast_slice_struct_string_Ttest(obj)
	if err != nil {
		return err
	}
	return val.Parse(data)
}

func ast_string(obj *string) (libjson.Value, error) {
	return libjson.NewString(func() string { return *obj }, func(s string) { *obj = s }), nil
}
