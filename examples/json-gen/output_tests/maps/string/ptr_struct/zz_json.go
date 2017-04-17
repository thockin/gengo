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

func ast_ptr_struct_Ttest(obj *Ttest) (libjson.Value, error) {
	return ast_Map_string_To_Pointer_ptr_struct_Struct((*map[string]*Struct)(obj))
}

func (obj Ttest) MarshalJSON() ([]byte, error) {
	jv, err := ast_ptr_struct_Ttest(&obj)
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
	jv, err := ast_ptr_struct_Ttest(obj)
	if err != nil {
		return err
	}
	return jv.Parse(data)
}

func ast_Map_string_To_Pointer_ptr_struct_Struct(obj *map[string]*Struct) (libjson.Value, error) {

	keyToString := func(k string) (string, error) {
		return string(k), nil
	}
	keyFromString := func(s string) (string, error) {
		return string(s), nil
	}

	var keys []string
	var vals []*Struct
	add := func(ks string) libjson.Value {
		if k, err := keyFromString(ks); err != nil {
			panic(err) //FIXME
		} else {
			keys = append(keys, k)
		}
		var x *Struct
		vals = append(vals, x)
		obj := &vals[len(vals)-1]
		jv, _ := ast_Pointer_ptr_struct_Struct((**Struct)(obj))
		//FIXME: handle error?
		return jv
	}
	get := func() (map[string]libjson.Value, error) {
		if *obj == nil && keys == nil {
			return nil, nil
		}
		result := map[string]libjson.Value{}
		for k, v := range *obj {
			obj := new(*Struct)
			*obj = v
			//FIXME: do any of these ACTUALLY return an error?
			jv, err := ast_Pointer_ptr_struct_Struct((**Struct)(obj))
			if err != nil {
				return nil, err
			}
			if ks, err := keyToString(k); err != nil {
				panic(err)
			} else {
				result[ks] = jv
			}
		}
		for i := range keys {
			obj := &vals[i]
			//FIXME: do any of these ACTUALLY return an error?
			jv, err := ast_Pointer_ptr_struct_Struct((**Struct)(obj))
			if err != nil {
				return nil, err
			}
			if ks, err := keyToString(keys[i]); err != nil {
				panic(err)
			} else {
				result[ks] = jv
			}
		}
		return result, nil
	}
	finishParse := func() {
		for i := range keys {
			(*obj)[keys[i]] = vals[i]
		}
	}
	var jv libjson.Value
	if *obj != nil {
		jv = libjson.NewMap(add, get, finishParse)
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = map[string]*Struct{}
		return libjson.NewMap(add, get, finishParse), nil
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Pointer_ptr_struct_Struct(obj **Struct) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_ptr_struct_Struct((*Struct)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(Struct)
		obj := *obj
		return ast_ptr_struct_Struct((*Struct)(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_ptr_struct_Struct(obj *Struct) (libjson.Value, error) {

	result := libjson.NewObject()

	// String string
	{
		obj := &obj.String

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
			*p = "String"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: String was empty")
		} //FIXME:
	}

	// Int int32
	{
		obj := &obj.Int

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
			*p = "Int"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Int was empty")
		} //FIXME:
	}

	// Float float64
	{
		obj := &obj.Float

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_float64((*float64)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Float"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Float was empty")
		} //FIXME:
	}

	// Struct struct{X string}
	{
		obj := &obj.Struct

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Struct_X_string((*struct{ X string })(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Struct"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Struct was empty")
		} //FIXME:
	}

	// Slice []string
	{
		obj := &obj.Slice

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Slice_string((*[]string)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Slice"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Slice was empty")
		} //FIXME:
	}

	// Map map[string]string
	{
		obj := &obj.Map

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Map_string_To_string((*map[string]string)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Map"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Map was empty")
		} //FIXME:
	}

	return result, nil

}

func ast_string(obj *string) (libjson.Value, error) {
	return libjson.NewString(func() string { return *obj }, func(s string) { *obj = s }), nil
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

func ast_float64(obj *float64) (libjson.Value, error) {

	get := func() float64 {
		return float64(*obj)
	}
	set := func(f float64) {
		*obj = float64(f)
	}
	return libjson.NewFloat(64, get, set), nil

}

func ast_Struct_X_string(obj *struct{ X string }) (libjson.Value, error) {

	result := libjson.NewObject()

	// X string
	{
		obj := &obj.X

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
			*p = "X"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: X was empty")
		} //FIXME:
	}

	return result, nil

}

func ast_Slice_string(obj *[]string) (libjson.Value, error) {

	get := func() ([]libjson.Value, error) {
		if *obj == nil {
			return nil, nil
		}
		result := []libjson.Value{}
		for i := range *obj {
			obj := &(*obj)[i]
			//FIXME: do any of these ACTUALLY return an error?
			jv, err := ast_string((*string)(obj))
			if err != nil {
				return nil, err
			}
			result = append(result, jv)
		}
		return result, nil
	}
	add := func() libjson.Value {
		var x string
		*obj = append(*obj, x)
		obj := &(*obj)[len(*obj)-1]
		jv, _ := ast_string((*string)(obj))
		//FIXME: handle error?
		return jv
	}
	var jv libjson.Value
	if *obj != nil {
		jv = libjson.NewArray(get, add)
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = []string{}
		return libjson.NewArray(get, add), nil
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Map_string_To_string(obj *map[string]string) (libjson.Value, error) {

	keyToString := func(k string) (string, error) {
		return string(k), nil
	}
	keyFromString := func(s string) (string, error) {
		return string(s), nil
	}

	var keys []string
	var vals []string
	add := func(ks string) libjson.Value {
		if k, err := keyFromString(ks); err != nil {
			panic(err) //FIXME
		} else {
			keys = append(keys, k)
		}
		var x string
		vals = append(vals, x)
		obj := &vals[len(vals)-1]
		jv, _ := ast_string((*string)(obj))
		//FIXME: handle error?
		return jv
	}
	get := func() (map[string]libjson.Value, error) {
		if *obj == nil && keys == nil {
			return nil, nil
		}
		result := map[string]libjson.Value{}
		for k, v := range *obj {
			obj := new(string)
			*obj = v
			//FIXME: do any of these ACTUALLY return an error?
			jv, err := ast_string((*string)(obj))
			if err != nil {
				return nil, err
			}
			if ks, err := keyToString(k); err != nil {
				panic(err)
			} else {
				result[ks] = jv
			}
		}
		for i := range keys {
			obj := &vals[i]
			//FIXME: do any of these ACTUALLY return an error?
			jv, err := ast_string((*string)(obj))
			if err != nil {
				return nil, err
			}
			if ks, err := keyToString(keys[i]); err != nil {
				panic(err)
			} else {
				result[ks] = jv
			}
		}
		return result, nil
	}
	finishParse := func() {
		for i := range keys {
			(*obj)[keys[i]] = vals[i]
		}
	}
	var jv libjson.Value
	if *obj != nil {
		jv = libjson.NewMap(add, get, finishParse)
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = map[string]string{}
		return libjson.NewMap(add, get, finishParse), nil
	}
	return libjson.NewNullable(jv, setNull), nil

}