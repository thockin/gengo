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

func ast_mixed_Ttest(obj *Ttest) (libjson.Value, error) {
	return ast_Slice_Pointer_Struct_F1_int32_F2_Pointer_int32_F3_float32_F4_Pointer_float32_F5_string_F6_Pointer_string_F7_Struct_F_string_F8_Pointer_Struct_F_string((*[]*struct {
		F1 int32
		F2 *int32
		F3 float32
		F4 *float32
		F5 string
		F6 *string
		F7 struct{ F string }
		F8 *struct{ F string }
	})(obj))
}

func (obj Ttest) MarshalJSON() ([]byte, error) {
	jv, err := ast_mixed_Ttest(&obj)
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
	jv, err := ast_mixed_Ttest(obj)
	if err != nil {
		return err
	}
	return jv.Parse(data)
}

func ast_Slice_Pointer_Struct_F1_int32_F2_Pointer_int32_F3_float32_F4_Pointer_float32_F5_string_F6_Pointer_string_F7_Struct_F_string_F8_Pointer_Struct_F_string(obj *[]*struct {
	F1 int32
	F2 *int32
	F3 float32
	F4 *float32
	F5 string
	F6 *string
	F7 struct{ F string }
	F8 *struct{ F string }
}) (libjson.Value, error) {

	get := func() ([]libjson.Value, error) {
		if *obj == nil {
			return nil, nil
		}
		result := []libjson.Value{}
		for i := range *obj {
			obj := &(*obj)[i]
			//FIXME: do any of these ACTUALLY return an error?
			jv, err := ast_Pointer_Struct_F1_int32_F2_Pointer_int32_F3_float32_F4_Pointer_float32_F5_string_F6_Pointer_string_F7_Struct_F_string_F8_Pointer_Struct_F_string((**struct {
				F1 int32
				F2 *int32
				F3 float32
				F4 *float32
				F5 string
				F6 *string
				F7 struct{ F string }
				F8 *struct{ F string }
			})(obj))
			if err != nil {
				return nil, err
			}
			result = append(result, jv)
		}
		return result, nil
	}
	add := func() libjson.Value {
		var x *struct {
			F1 int32
			F2 *int32
			F3 float32
			F4 *float32
			F5 string
			F6 *string
			F7 struct{ F string }
			F8 *struct{ F string }
		}
		*obj = append(*obj, x)
		obj := &(*obj)[len(*obj)-1]
		jv, _ := ast_Pointer_Struct_F1_int32_F2_Pointer_int32_F3_float32_F4_Pointer_float32_F5_string_F6_Pointer_string_F7_Struct_F_string_F8_Pointer_Struct_F_string((**struct {
			F1 int32
			F2 *int32
			F3 float32
			F4 *float32
			F5 string
			F6 *string
			F7 struct{ F string }
			F8 *struct{ F string }
		})(obj))
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
		*obj = []*struct {
			F1 int32
			F2 *int32
			F3 float32
			F4 *float32
			F5 string
			F6 *string
			F7 struct{ F string }
			F8 *struct{ F string }
		}{}
		return libjson.NewArray(get, add), nil
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Pointer_Struct_F1_int32_F2_Pointer_int32_F3_float32_F4_Pointer_float32_F5_string_F6_Pointer_string_F7_Struct_F_string_F8_Pointer_Struct_F_string(obj **struct {
	F1 int32
	F2 *int32
	F3 float32
	F4 *float32
	F5 string
	F6 *string
	F7 struct{ F string }
	F8 *struct{ F string }
}) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_Struct_F1_int32_F2_Pointer_int32_F3_float32_F4_Pointer_float32_F5_string_F6_Pointer_string_F7_Struct_F_string_F8_Pointer_Struct_F_string((*struct {
			F1 int32
			F2 *int32
			F3 float32
			F4 *float32
			F5 string
			F6 *string
			F7 struct{ F string }
			F8 *struct{ F string }
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
			F2 *int32
			F3 float32
			F4 *float32
			F5 string
			F6 *string
			F7 struct{ F string }
			F8 *struct{ F string }
		})
		obj := *obj
		return ast_Struct_F1_int32_F2_Pointer_int32_F3_float32_F4_Pointer_float32_F5_string_F6_Pointer_string_F7_Struct_F_string_F8_Pointer_Struct_F_string((*struct {
			F1 int32
			F2 *int32
			F3 float32
			F4 *float32
			F5 string
			F6 *string
			F7 struct{ F string }
			F8 *struct{ F string }
		})(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}

func ast_Struct_F1_int32_F2_Pointer_int32_F3_float32_F4_Pointer_float32_F5_string_F6_Pointer_string_F7_Struct_F_string_F8_Pointer_Struct_F_string(obj *struct {
	F1 int32
	F2 *int32
	F3 float32
	F4 *float32
	F5 string
	F6 *string
	F7 struct{ F string }
	F8 *struct{ F string }
}) (libjson.Value, error) {

	result := libjson.NewObject()

	// F1 int32
	{
		obj := &obj.F1
		_ = obj //FIXME: remove when other Kinds are done

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

	// F2 *int32
	{
		obj := &obj.F2
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

	// F3 float32
	{
		obj := &obj.F3
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_float32((*float32)(obj))
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

	// F4 *float32
	{
		obj := &obj.F4
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
			*p = "F4"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: F4 was empty")
		} //FIXME:
	}

	// F5 string
	{
		obj := &obj.F5
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
			*p = "F5"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: F5 was empty")
		} //FIXME:
	}

	// F6 *string
	{
		obj := &obj.F6
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
			*p = "F6"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: F6 was empty")
		} //FIXME:
	}

	// F7 struct{F string}
	{
		obj := &obj.F7
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Struct_F_string((*struct{ F string })(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "F7"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: F7 was empty")
		} //FIXME:
	}

	// F8 *struct{F string}
	{
		obj := &obj.F8
		_ = obj //FIXME: remove when other Kinds are done

		empty := func(libjson.Value) bool { return false }

		finalize := func(jv libjson.Value) (libjson.Value, error) { return jv, nil }

		jv, err := ast_Pointer_Struct_F_string((**struct{ F string })(obj))
		if err != nil {
			return nil, err
		}
		if !empty(jv) {
			fv, err := finalize(jv)
			if err != nil {
				return nil, err
			}
			p := new(string)
			*p = "F8"
			nv := libjson.NamedValue{
				Name:  libjson.NewString(func() string { return *p }, func(s string) { *p = s }),
				Value: fv,
			}
			result = append(result, nv)
		} else {
			panic("TIM: F8 was empty")
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

func ast_float32(obj *float32) (libjson.Value, error) {

	get := func() float64 {
		return float64(*obj)
	}
	set := func(f float64) {
		*obj = float32(f)
	}
	return libjson.NewFloat(32, get, set), nil

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

func ast_string(obj *string) (libjson.Value, error) {
	return libjson.NewString(func() string { return *obj }, func(s string) { *obj = s }), nil
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

func ast_Struct_F_string(obj *struct{ F string }) (libjson.Value, error) {

	result := libjson.NewObject()

	// F string
	{
		obj := &obj.F
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

func ast_Pointer_Struct_F_string(obj **struct{ F string }) (libjson.Value, error) {

	var jv libjson.Value
	var err error
	if *obj != nil {
		obj := *obj
		jv, err = ast_Struct_F_string((*struct{ F string })(obj))
		if err != nil {
			return nil, err
		}
	}
	setNull := func(b bool) (libjson.Value, error) {
		if b {
			*obj = nil
			return nil, nil
		}
		*obj = new(struct{ F string })
		obj := *obj
		return ast_Struct_F_string((*struct{ F string })(obj))
	}
	return libjson.NewNullable(jv, setNull), nil

}
