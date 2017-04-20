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
	jv, err := ast_byte_Ttest(&obj)
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
	jv, err := ast_byte_Ttest(obj)
	if err != nil {
		return err
	}
	return jv.Parse(data)
}

func ast_byte_Ttest(obj *Ttest) (libjson.Value, error) {
	return ast_Map_string_To_byte((*map[string]byte)(obj))
}

func ast_Map_string_To_byte(obj *map[string]byte) (libjson.Value, error) {

	keyToString := func(k string) (string, error) {
		return string(k), nil
	}

	keyFromString := func(s string) (string, error) {
		return string(s), nil
	}

	var keys []string
	var vals []byte
	add := func(ks string) libjson.Value {
		if k, err := keyFromString(ks); err != nil {
			panic(err) //FIXME
		} else {
			keys = append(keys, k)
		}
		var x byte
		vals = append(vals, x)
		obj := &vals[len(vals)-1]
		jv, _ := ast_byte((*byte)(obj))
		//FIXME: handle error?
		return jv
	}
	get := func() (map[string]libjson.Value, error) {
		if *obj == nil && keys == nil {
			return nil, nil
		}
		result := map[string]libjson.Value{}
		for k, v := range *obj {
			obj := new(byte)
			*obj = v
			//FIXME: do any of these ACTUALLY return an error?
			jv, err := ast_byte((*byte)(obj))
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
			jv, err := ast_byte((*byte)(obj))
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
		*obj = map[string]byte{}
		return libjson.NewMap(add, get, finishParse), nil
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_byte(obj *byte) (libjson.Value, error) {

	get := func() float64 {
		return float64(*obj)
	}
	set := func(f float64) {
		*obj = byte(f)
	}
	return libjson.NewInt(get, set), nil

}
