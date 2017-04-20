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
	strconv "strconv"
)

func (obj Ttest) MarshalJSON() ([]byte, error) {
	jv, err := ast_bool_Ttest(&obj)
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
	jv, err := ast_bool_Ttest(obj)
	if err != nil {
		return err
	}
	return jv.Parse(data)
}

func ast_bool_Ttest(obj *Ttest) (libjson.Value, error) {
	return ast_Map_int32_To_bool((*map[int32]bool)(obj))
}

func ast_Map_int32_To_bool(obj *map[int32]bool) (libjson.Value, error) {

	keyToString := func(k int32) (string, error) {
		return strconv.FormatInt(int64(k), 10), nil
	}

	keyFromString := func(s string) (int32, error) {
		i, err := strconv.ParseInt(s, 10, 64)
		return int32(i), err
	}

	var keys []int32
	var vals []bool
	add := func(ks string) libjson.Value {
		if k, err := keyFromString(ks); err != nil {
			panic(err) //FIXME
		} else {
			keys = append(keys, k)
		}
		var x bool
		vals = append(vals, x)
		obj := &vals[len(vals)-1]
		jv, _ := ast_bool((*bool)(obj))
		//FIXME: handle error?
		return jv
	}
	get := func() (map[string]libjson.Value, error) {
		if *obj == nil && keys == nil {
			return nil, nil
		}
		result := map[string]libjson.Value{}
		for k, v := range *obj {
			obj := new(bool)
			*obj = v
			//FIXME: do any of these ACTUALLY return an error?
			jv, err := ast_bool((*bool)(obj))
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
			jv, err := ast_bool((*bool)(obj))
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
		*obj = map[int32]bool{}
		return libjson.NewMap(add, get, finishParse), nil
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_bool(obj *bool) (libjson.Value, error) {
	return libjson.NewBool(func() bool { return *obj }, func(b bool) { *obj = b }), nil
}
