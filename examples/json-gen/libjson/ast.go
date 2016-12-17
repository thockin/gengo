package libjson

import (
	"bytes"
	"fmt"
	"strconv"
)

type Value interface {
	Render(buf *bytes.Buffer) error
	Parse(data []byte) error
	Empty() bool
	//FIXME: Get() interface{}
}

type Optional struct {
	inner Value
	isSet bool
	set   func()
	clear func()
}

func NewOptional(inner Value, isSet bool, set func(), clear func()) Value {
	return &Optional{
		inner: inner,
		isSet: isSet,
		//FIXME: one func(bool) for set/clear?
		set:   set,
		clear: clear,
		//FIXME: how to init isSet and initial p?
	}
}
func (value *Optional) Render(buf *bytes.Buffer) error {
	if !value.isSet {
		return writeString(buf, "null")
	}
	return value.inner.Render(buf)
}

func (value *Optional) Parse(data []byte) error {
	if string(data) == "null" {
		value.clear()
		value.isSet = false
		return nil
	}
	value.set()
	value.isSet = true
	return value.inner.Parse(data)
}

func (value *Optional) Empty() bool {
	return value.isSet
}

type String struct {
	rfn func() *string
	wfn func(s string)
}

func NewString(rfn func() *string, wfn func(s string)) String {
	return String{
		rfn: rfn,
		wfn: wfn,
	}
}

func (value String) Render(buf *bytes.Buffer) error {
	//FIXME: make rfn return a value?
	return writeString(buf, `"`+*value.rfn()+`"`)
}

func (value String) Parse(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("data is not a JSON string (%q)", string(data))
	}
	value.wfn(string(data[1 : len(data)-1]))
	return nil
}

func (value String) Empty() bool {
	//FIXME: make rfn return a value?
	return *value.rfn() == ""
}

const hexits = "0123456789abcdef"

// escape ensures that a string does not hold characters that need to be
// escaped in JSON strings.  This is almost entirely derived from Go's JSON
// encoder - see the comments there for justifications.  This DOES NOT wrap
// the result in quotes.
func escape(str string) string {
	buf := bytes.Buffer{}
	runes := []rune(str)
	for _, r := range runes {
		switch r {
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
		case '\\':
			buf.WriteString(`\\`)
		case '"':
			buf.WriteString(`\"`)
		case '<', '>', '&':
			buf.WriteString(`\u00`)
			buf.WriteByte(hexits[r>>4])
			buf.WriteByte(hexits[r&0xf])
		case '\u2028', '\u2029':
			buf.WriteString(`\u202`)
			buf.WriteByte(hexits[r&0xf])
		default:
			if r < 0x20 {
				buf.WriteString(`\u00`)
				buf.WriteByte(hexits[r>>4])
				buf.WriteByte(hexits[r&0xf])
			} else {
				buf.WriteRune(r)
			}
		}
	}
	return buf.String()
}

type Number float64

func (value Number) Render(buf *bytes.Buffer) error {
	return writeString(buf, strconv.FormatFloat(float64(value), 'g', 64, 64))
}

func (Number) Parse([]byte) error { return nil }

func (value Number) Empty() bool {
	return value == 0
}

type Bool bool

var trueBytes = []byte("true")
var falseBytes = []byte("false")

func (value Bool) Render(buf *bytes.Buffer) error {
	if value {
		return write(buf, trueBytes)
	}
	return write(buf, falseBytes)
}

func (Bool) Parse([]byte) error { return nil }

func (value Bool) Empty() bool {
	return value == false
}

type Null struct{}

var nullBytes = []byte("null")

func (Null) Render(buf *bytes.Buffer) error {
	return write(buf, nullBytes)
}

func (Null) Parse([]byte) error { return nil }

func (Null) Empty() bool {
	return true
}

type Object []NamedValue

type NamedValue struct {
	Name  String
	Value Value
}

func (value Object) Render(buf *bytes.Buffer) error {
	if err := writeString(buf, "{"); err != nil {
		return err
	}
	for i, nv := range value {
		if i > 0 {
			if err := writeString(buf, ","); err != nil {
				return err
			}
		}
		if err := nv.Name.Render(buf); err != nil {
			return err
		}
		if err := writeString(buf, ":"); err != nil {
			return err
		}
		if err := nv.Value.Render(buf); err != nil {
			return err
		}
	}
	if err := writeString(buf, "}"); err != nil {
		return err
	}
	return nil
}

func (Object) Parse([]byte) error { return nil }

func (value Object) Empty() bool {
	return len(value) == 0
}

type Array []Value

func (value Array) Render(buf *bytes.Buffer) error {
	if err := writeString(buf, "["); err != nil {
		return err
	}
	for i, v := range value {
		if i > 0 {
			if err := writeString(buf, ","); err != nil {
				return err
			}
		}
		if err := v.Render(buf); err != nil {
			return err
		}
	}
	if err := writeString(buf, "]"); err != nil {
		return err
	}
	return nil
}

func (Array) Parse([]byte) error { return nil }

func (value Array) Empty() bool {
	return len(value) == 0
}

type Raw string

func (value Raw) Render(buf *bytes.Buffer) error {
	return writeString(buf, string(value))
}

func (Raw) Parse([]byte) error { return nil }

func (value Raw) Empty() bool {
	return len(value) == 0
}

func write(buf *bytes.Buffer, val []byte) error {
	_, err := buf.Write(val)
	return err
}

func writeString(buf *bytes.Buffer, val string) error {
	_, err := buf.WriteString(val)
	return err
}
