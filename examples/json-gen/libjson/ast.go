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
	//FIXME: escapes and go-JSON compat
	return writeString(buf, `"`+string(value)+`"`)
}

func (value String) Empty() bool {
	return value == ""
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
