package libjson

import (
	"bytes"
	"strconv"
)

type Value interface {
	Render(buf *bytes.Buffer) error
	Empty() bool
}

type String string

func (value String) Render(buf *bytes.Buffer) error {
	return writeString(buf, `"`+escape(string(value))+`"`)
}

func (value String) Empty() bool {
	return value == ""
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

func (value Bool) Empty() bool {
	return value == false
}

type Null struct{}

var nullBytes = []byte("null")

func (Null) Render(buf *bytes.Buffer) error {
	return write(buf, nullBytes)
}

func (Null) Empty() bool {
	return true
}

type Object []NamedValue

type NamedValue struct {
	Name  string
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
		if err := writeString(buf, `"`+nv.Name+`":`); err != nil {
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

func (value Array) Empty() bool {
	return len(value) == 0
}

type Raw string

func (value Raw) Render(buf *bytes.Buffer) error {
	return writeString(buf, string(value))
}

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
