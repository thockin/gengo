/*
Copyright 2016 The Kubernetes Authors.

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

package libjson

import (
	"bytes"
	"math"
	"testing"
)

func Test_Render(t *testing.T) {
	testCases := []struct {
		in    Value
		out   string
		empty bool
		err   bool
	}{
		{
			in:    String(""),
			out:   `""`,
			empty: true,
		},
		{
			in:    String("abc123"),
			out:   `"abc123"`,
			empty: false,
		},
		{
			in:    String("abc 123"),
			out:   `"abc 123"`,
			empty: false,
		},
		{
			in:    String("abc<>&\\\n\r\t\u2028\u2029\u2030\x00\x1f"),
			out:   "\"abc\\u003c\\u003e\\u0026\\\\\\n\\r\\t\\u2028\\u2029\u2030\\u0000\\u001f\"",
			empty: false,
		},
		{
			in:    Number(0),
			out:   "0",
			empty: true,
		},
		{
			in:    Number(1),
			out:   "1",
			empty: false,
		},
		{
			in:    Number(math.Exp2(53) - 1),
			out:   "9007199254740991",
			empty: false,
		},
		{
			in:    Number(-math.Exp2(53)),
			out:   "-9007199254740992",
			empty: false,
		},
		{
			in:    Number(3.5),
			out:   "3.5",
			empty: false,
		},
		{
			in:    Bool(true),
			out:   "true",
			empty: false,
		},
		{
			in:    Bool(false),
			out:   "false",
			empty: true,
		},
		{
			in:    Null{},
			out:   "null",
			empty: true,
		},
		{
			in:    Object{},
			out:   `{}`,
			empty: true,
		},
		{
			in: Object{
				NamedValue{"k1", String("v1")},
				NamedValue{"k2", Number(2)},
				NamedValue{"k3", Bool(true)},
			},
			out:   `{"k1":"v1","k2":2,"k3":true}`,
			empty: false,
		},
		{
			in:    Array{},
			out:   `[]`,
			empty: true,
		},
		{
			in:    Array{String("v1"), String("v2"), String("v3")},
			out:   `["v1","v2","v3"]`,
			empty: false,
		},
		{
			in:    Raw(``),
			out:   ``,
			empty: true,
		},
		{
			in:    Raw(`{"a": 1, "b": 2}`),
			out:   `{"a": 1, "b": 2}`,
			empty: false,
		},
	}

	for i, tc := range testCases {
		var buf bytes.Buffer
		err := tc.in.Render(&buf)
		if err != nil && tc.err == false {
			t.Errorf("[%d] unexpected error: %v", i, err)
		} else if err == nil && tc.err == true {
			t.Errorf("[%d] expected error, got nil", i)
		} else if buf.String() != tc.out {
			t.Errorf("[%d] expected %q, got %q", i, tc.out, buf.String())
		} else if tc.in.Empty() != tc.empty {
			t.Errorf("[%d] expected Empty() = %t", i, tc.in.Empty())
		}
	}
}
