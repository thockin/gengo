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
	"testing"
)

func Test_advanceWhitespace(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
	}{{
		input:  "abc",
		expect: "abc",
	}, {
		input:  " abc",
		expect: "abc",
	}, {
		input:  "  abc",
		expect: "abc",
	}, {
		input:  "  a b c ",
		expect: "a b c ",
	}, {
		input:  "\t\tabc",
		expect: "abc",
	}, {
		input:  "\n\nabc",
		expect: "abc",
	}, {
		input:  "    ",
		expect: "",
	}}

	for i, tc := range testCases {
		scan := NewByteScanner([]byte(tc.input))
		advanceWhitespace(scan)
		if string(scan.Data()) != tc.expect {
			t.Errorf("[%d] expected %q, got %q", i, tc.expect, string(scan.Data()))
		}
	}
}

func Test_advanceNull(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
		err    bool
	}{{
		input:  "null",
		expect: "",
	}, {
		input:  "null ",
		expect: " ",
	}, {
		input: "nullnull ",
		err:   true,
	}, {
		input: "Null",
		err:   true,
	}, {
		input: "nULL",
		err:   true,
	}}

	for i, tc := range testCases {
		scan := NewByteScanner([]byte(tc.input))
		if err := advanceToken(scan, advanceNull); err != nil {
			if tc.err == false {
				t.Errorf("[%d] expected success, got %v", i, err)
			}
		} else {
			if tc.err == true {
				t.Errorf("[%d] expected error, got success", i)
			} else if err == nil && string(scan.Data()) != tc.expect {
				t.Errorf("[%d] expected %q, got %q", i, tc.expect, string(scan.Data()))
			}
		}
	}
}

func Test_advanceBool(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
		err    bool
	}{{
		input:  "true",
		expect: "",
	}, {
		input:  "true ",
		expect: " ",
	}, {
		input: "truefalse",
		err:   true,
	}, {
		input: "True",
		err:   true,
	}, {
		input: "tRUE",
		err:   true,
	}, {
		input:  "false",
		expect: "",
	}, {
		input:  "false ",
		expect: " ",
	}, {
		input: "falsetrue",
		err:   true,
	}, {
		input: "False",
		err:   true,
	}, {
		input: "fALSE",
		err:   true,
	}}

	for i, tc := range testCases {
		scan := NewByteScanner([]byte(tc.input))
		if err := advanceToken(scan, advanceBool); err != nil {
			if tc.err == false {
				t.Errorf("[%d] expected success, got %v", i, err)
			}
		} else {
			if tc.err == true {
				t.Errorf("[%d] expected error, got success", i)
			} else if err == nil && string(scan.Data()) != tc.expect {
				t.Errorf("[%d] expected %q, got %q", i, tc.expect, string(scan.Data()))
			}
		}
	}
}

func Test_advanceNumber(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
		err    bool
	}{{
		input:  "0",
		expect: "",
	}, {
		input:  "-0",
		expect: "",
	}, {
		input:  "0 ",
		expect: " ",
	}, {
		input:  "-0 ",
		expect: " ",
	}, {
		input:  "123",
		expect: "",
	}, {
		input:  "-123",
		expect: "",
	}, {
		input:  "123.456",
		expect: "",
	}, {
		input:  "123.456 ",
		expect: " ",
	}, {
		input:  "123.000",
		expect: "",
	}, {
		input:  "-123.456",
		expect: "",
	}, {
		input:  "-123.456 ",
		expect: " ",
	}, {
		input: "123.-456",
		err:   true,
	}, {
		input:  "123e456",
		expect: "",
	}, {
		input:  "123e456 ",
		expect: " ",
	}, {
		input:  "123e000",
		expect: "",
	}, {
		input:  "123e001",
		expect: "",
	}, {
		input:  "123.456e789",
		expect: "",
	}, {
		input:  "123.456e789 ",
		expect: " ",
	}, {
		input:  "123.456e000",
		expect: "",
	}, {
		input:  "123.456E789",
		expect: "",
	}, {
		input:  "123.456E000",
		expect: "",
	}, {
		input:  "123.456e+789",
		expect: "",
	}, {
		input:  "123.456e-789",
		expect: "",
	}, {
		input:  "123.456e+000",
		expect: "",
	}, {
		input:  "123.456e-000",
		expect: "",
	}, {
		input:  "123.456E+000",
		expect: "",
	}, {
		input:  "123.456E-000",
		expect: "",
	}, {
		input: "123xyz",
		err:   true,
	}, {
		input: "0x12345678",
		err:   true,
	}, {
		input: "00",
		err:   true,
	}, {
		input: "0755",
		err:   true,
	}, {
		input: ".007",
		err:   true,
	}, {
		input: "0.",
		err:   true,
	}, {
		input: "0e",
		err:   true,
	}}

	for i, tc := range testCases {
		scan := NewByteScanner([]byte(tc.input))
		if err := advanceToken(scan, advanceNumber); err != nil {
			if tc.err == false {
				t.Errorf("[%d] expected success, got %v", i, err)
			}
		} else {
			if tc.err == true {
				t.Errorf("[%d] expected error, got success", i)
			} else if err == nil && string(scan.Data()) != tc.expect {
				t.Errorf("[%d] expected %q, got %q", i, tc.expect, string(scan.Data()))
			}
		}
	}
}

func Test_advanceString(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
		err    bool
	}{{
		input:  `"abc"`,
		expect: ``,
	}, {
		input:  `"abc" `,
		expect: ` `,
	}, {
		input:  `"abc\n123\tdef\u1234\"\"" `,
		expect: ` `,
	}, {
		input: `"abc"123`,
		err:   true,
	}}

	for i, tc := range testCases {
		scan := NewByteScanner([]byte(tc.input))
		if err := advanceToken(scan, advanceString); err != nil {
			if tc.err == false {
				t.Errorf("[%d] expected success, got %v", i, err)
			}
		} else {
			if tc.err == true {
				t.Errorf("[%d] expected error, got success", i)
			} else if err == nil && string(scan.Data()) != tc.expect {
				t.Errorf("[%d] expected %q, got %q", i, tc.expect, string(scan.Data()))
			}
		}
	}
}

func Test_advanceArray(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
		err    bool
	}{{
		input:  `[]`,
		expect: ``,
	}, {
		input:  `[] `,
		expect: ` `,
	}, {
		input:  `[ ]`,
		expect: ``,
	}, {
		input:  `["abc"]`,
		expect: ``,
	}, {
		input:  `[ "abc"]`,
		expect: ``,
	}, {
		input:  `["abc" ]`,
		expect: ``,
	}, {
		input:  `[ "abc" ]`,
		expect: ``,
	}, {
		input: `["abc",]`,
		err:   true,
	}, {
		input: `["abc" ,]`,
		err:   true,
	}, {
		input: `["abc", ]`,
		err:   true,
	}, {
		input:  `["abc", "def","ghi"]`,
		expect: ``,
	}, {
		input:  "[ \n \"abc\" \n , \n \"def\" \n , \n \"ghi\" \n ]",
		expect: ``,
	}, {
		input: `["abc"]123`,
		err:   true,
	}, {
		input:  `[[]]`,
		expect: ``,
	}, {
		input:  `[[ ]]`,
		expect: ``,
	}, {
		input:  `[ [ ] ]`,
		expect: ``,
	}, {
		input:  `[["abc"]]`,
		expect: ``,
	}, {
		input:  `[ [ "abc" ] ]`,
		expect: ``,
	}, {
		input: `[["abc",]]`,
		err:   true,
	}, {
		input: `[["abc"],]`,
		err:   true,
	}, {
		input:  `[ ["abc", "def"], ["ghi", "jkl"] ]`,
		expect: ``,
	}, {
		input:  `[ ["abc", 123], ["def", false] ]`,
		expect: ``,
	}, {
		input: `[ "abc", abc, "def", def ]`,
		err:   true,
	}, {
		input: `[["abc"]["def"]]`,
		err:   true,
	}, {
		input: `["abc"]["def"]`,
		err:   true,
	}}

	for i, tc := range testCases {
		scan := NewByteScanner([]byte(tc.input))
		if err := advanceToken(scan, advanceArray); err != nil {
			if tc.err == false {
				t.Errorf("[%d] expected success, got %v", i, err)
			}
		} else {
			if tc.err == true {
				t.Errorf("[%d] expected error, got success", i)
			} else if err == nil && string(scan.Data()) != tc.expect {
				t.Errorf("[%d] expected %q, got %q", i, tc.expect, string(scan.Data()))
			}
		}
	}
}

func Test_advanceObject(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
		err    bool
	}{{
		input:  `{}`,
		expect: ``,
	}, {
		input:  `{} `,
		expect: ` `,
	}, {
		input:  `{ }`,
		expect: ``,
	}, {
		input:  `{"abc":123}`,
		expect: ``,
	}, {
		input:  `{ "abc":123}`,
		expect: ``,
	}, {
		input:  `{"abc" :123}`,
		expect: ``,
	}, {
		input:  `{"abc": 123}`,
		expect: ``,
	}, {
		input:  `{"abc":123 }`,
		expect: ``,
	}, {
		input:  `{ "abc" : 123 }`,
		expect: ``,
	}, {
		input:  "{\n\t\"abc\"\n\t:\n\t123\n\t}",
		expect: ``,
	}, {
		input: `{"abc":123,}`,
		err:   true,
	}, {
		input: `{"abc":123 ,}`,
		err:   true,
	}, {
		input: `{"abc":123, }`,
		err:   true,
	}, {
		input:  `{"abc": "def"}`,
		expect: ``,
	}, {
		input:  `{"abc": true}`,
		expect: ``,
	}, {
		input:  `{"abc": false}`,
		expect: ``,
	}, {
		input:  `{"abc": null}`,
		expect: ``,
	}, {
		input:  `{"abc": []}`,
		expect: ``,
	}, {
		input:  `{"abc": [[[]]]}`,
		expect: ``,
	}, {
		input:  `{"abc": [1,2,3]}`,
		expect: ``,
	}, {
		input:  `{"abc": {"abc":{"abc":{}}}}`,
		expect: ``,
	}, {
		input:  `{"abc": {"1":23}}`,
		expect: ``,
	}, {
		input:  `{"a":1, "b":2, "c":3}`,
		expect: ``,
	}, {
		input:  `{ "a" : 1 , "b" : 2 , "c" : 3 }`,
		expect: ``,
	}, {
		input: `{"abc":123}xyz`,
		err:   true,
	}, {
		input: `{"abc"}`,
		err:   true,
	}, {
		input: `{"abc":123}{"def":456}`,
		err:   true,
	}, {
		input: `{"abc":123 "def":456}`,
		err:   true,
	}}

	for i, tc := range testCases {
		scan := NewByteScanner([]byte(tc.input))
		if err := advanceToken(scan, advanceObject); err != nil {
			if tc.err == false {
				t.Errorf("[%d] expected success, got %v", i, err)
			}
		} else {
			if tc.err == true {
				t.Errorf("[%d] expected error, got success", i)
			} else if err == nil && string(scan.Data()) != tc.expect {
				t.Errorf("[%d] expected %q, got %q", i, tc.expect, string(scan.Data()))
			}
		}
	}
}
