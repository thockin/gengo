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

func ast_everything_Ttest(obj *Ttest) (libjson.Value, error) {

	result := libjson.NewObject()

	// Byte1 *byte
	{
		obj := &obj.Byte1
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_byte((**byte)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Byte1"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Byte1 was empty")
		} //FIXME:
	}

	// Byte2 *byte
	{
		obj := &obj.Byte2
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_byte((**byte)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Byte2"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Byte2 was empty")
		} //FIXME:
	}

	// Bool1 *bool
	{
		obj := &obj.Bool1
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_bool((**bool)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Bool1"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Bool1 was empty")
		} //FIXME:
	}

	// Bool2 *bool
	{
		obj := &obj.Bool2
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_bool((**bool)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Bool2"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Bool2 was empty")
		} //FIXME:
	}

	// Int8 *int8
	{
		obj := &obj.Int8
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_int8((**int8)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Int8"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Int8 was empty")
		} //FIXME:
	}

	// Int16 *int16
	{
		obj := &obj.Int16
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_int16((**int16)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Int16"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Int16 was empty")
		} //FIXME:
	}

	// Int32 *int32
	{
		obj := &obj.Int32
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_int32((**int32)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Int32"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Int32 was empty")
		} //FIXME:
	}

	// Uint8 *uint8
	{
		obj := &obj.Uint8
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_uint8((**uint8)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Uint8"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Uint8 was empty")
		} //FIXME:
	}

	// Uint16 *uint16
	{
		obj := &obj.Uint16
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_uint16((**uint16)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Uint16"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Uint16 was empty")
		} //FIXME:
	}

	// Uint32 *uint32
	{
		obj := &obj.Uint32
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_uint32((**uint32)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Uint32"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Uint32 was empty")
		} //FIXME:
	}

	// Float32 *float32
	{
		obj := &obj.Float32
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_float32((**float32)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Float32"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Float32 was empty")
		} //FIXME:
	}

	// Float64 *float64
	{
		obj := &obj.Float64
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_float64((**float64)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "Float64"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: Float64 was empty")
		} //FIXME:
	}

	// String1 *string
	{
		obj := &obj.String1
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_string((**string)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "String1"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: String1 was empty")
		} //FIXME:
	}

	// String2 *string
	{
		obj := &obj.String2
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_string((**string)(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "String2"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: String2 was empty")
		} //FIXME:
	}

	return result, nil

}

func (obj Ttest) MarshalJSON() ([]byte, error) {
	jv, err := ast_everything_Ttest(&obj)
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
	jv, err := ast_everything_Ttest(obj)
	if err != nil {
		return err
	}
	return jv.Parse(data)
}

func ast_Pointer_byte(obj **byte) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_byte((*byte)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(byte)
		obj := *obj
		return ast_byte((*byte)(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Pointer_bool(obj **bool) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_bool((*bool)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(bool)
		obj := *obj
		return ast_bool((*bool)(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Pointer_int8(obj **int8) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_int8((*int8)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(int8)
		obj := *obj
		return ast_int8((*int8)(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Pointer_int16(obj **int16) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_int16((*int16)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(int16)
		obj := *obj
		return ast_int16((*int16)(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Pointer_int32(obj **int32) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_int32((*int32)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(int32)
		obj := *obj
		return ast_int32((*int32)(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Pointer_uint8(obj **uint8) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_uint8((*uint8)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(uint8)
		obj := *obj
		return ast_uint8((*uint8)(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Pointer_uint16(obj **uint16) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_uint16((*uint16)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(uint16)
		obj := *obj
		return ast_uint16((*uint16)(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Pointer_uint32(obj **uint32) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_uint32((*uint32)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(uint32)
		obj := *obj
		return ast_uint32((*uint32)(obj))
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

func ast_Pointer_float64(obj **float64) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_float64((*float64)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(float64)
		obj := *obj
		return ast_float64((*float64)(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Pointer_string(obj **string) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_string((*string)(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(string)
		obj := *obj
		return ast_string((*string)(obj))
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

func ast_bool(obj *bool) (libjson.Value, error) {
	return libjson.NewBool(func() bool { return *obj }, func(b bool) { *obj = b }), nil
}

func ast_int8(obj *int8) (libjson.Value, error) {

	get := func() float64 {
		return float64(*obj)
	}
	set := func(f float64) {
		*obj = int8(f)
	}
	return libjson.NewInt(get, set), nil

}

func ast_int16(obj *int16) (libjson.Value, error) {

	get := func() float64 {
		return float64(*obj)
	}
	set := func(f float64) {
		*obj = int16(f)
	}
	return libjson.NewInt(get, set), nil

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

func ast_uint8(obj *uint8) (libjson.Value, error) {

	get := func() float64 {
		return float64(*obj)
	}
	set := func(f float64) {
		*obj = uint8(f)
	}
	return libjson.NewInt(get, set), nil

}

func ast_uint16(obj *uint16) (libjson.Value, error) {

	get := func() float64 {
		return float64(*obj)
	}
	set := func(f float64) {
		*obj = uint16(f)
	}
	return libjson.NewInt(get, set), nil

}

func ast_uint32(obj *uint32) (libjson.Value, error) {

	get := func() float64 {
		return float64(*obj)
	}
	set := func(f float64) {
		*obj = uint32(f)
	}
	return libjson.NewInt(get, set), nil

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

func ast_float64(obj *float64) (libjson.Value, error) {

	get := func() float64 {
		return float64(*obj)
	}
	set := func(f float64) {
		*obj = float64(f)
	}
	return libjson.NewFloat(64, get, set), nil

}

func ast_string(obj *string) (libjson.Value, error) {
	return libjson.NewString(func() string { return *obj }, func(s string) { *obj = s }), nil
}
