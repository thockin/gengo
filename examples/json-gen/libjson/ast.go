package libjson

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/golang/glog"
)

type Value interface {
	Render(buf *bytes.Buffer) error
	Parse(data []byte) error
	ParseStream(scan *ByteScanner) error
	Empty() bool
	//FIXME: Get() interface{}
}

type Optional struct {
	inner Value
	isSet bool
	set   func()
	clear func()
}

func NewOptional(inner Value, isSet bool, set func(), clear func()) *Optional {
	return &Optional{
		inner: inner,
		isSet: isSet,
		//FIXME: one func(bool) for set/clear?
		set:   set,
		clear: clear,
		//FIXME: how to init isSet and initial p?
	}
}

var nullBytes = []byte("null")

func (value *Optional) Render(buf *bytes.Buffer) error {
	if !value.isSet {
		return write(buf, nullBytes)
	}
	return value.inner.Render(buf)
}

func (value *Optional) Parse(data []byte) error {
	scan := NewByteScanner(data)
	if err := value.ParseStream(scan); err != nil {
		return err
	}
	if len(scan.Data()) != 0 {
		return fmt.Errorf("found trailing data in input: %q", string(scan.Data()))
	}
	return nil
}

func (value *Optional) ParseStream(scan *ByteScanner) error {
	if len(scan.Data()) >= len(nullBytes) {
		mightBe := true
		for _, b := range nullBytes {
			if scan.Peek() != rune(b) {
				mightBe = false
				break
			}
			scan.Advance()
		}
		if mightBe {
			//FIXME: test what happens at end-of-buffer
			if isValueDelim(scan.Peek()) {
				value.clear()
				value.isSet = false
				scan.Save()
				return nil
			}
		}
	}
	value.set()
	value.isSet = true
	scan.Reset()
	return value.inner.ParseStream(scan)
}

func (value *Optional) Empty() bool {
	return value.isSet
}

type String struct {
	get func() string
	set func(s string)
}

func NewString(get func() string, set func(s string)) String {
	return String{
		get: get,
		set: set,
	}
}

func (value String) Render(buf *bytes.Buffer) error {
	return writeString(buf, `"`+value.get()+`"`)
}

func (value String) Parse(data []byte) error {
	scan := NewByteScanner(data)
	if err := value.ParseStream(scan); err != nil {
		return err
	}
	if len(scan.Data()) != 0 {
		return fmt.Errorf("found trailing data in input: %q", string(scan.Data()))
	}
	return nil
}

func (value String) ParseStream(scan *ByteScanner) error {
	if len(scan.Data()) < 2 || scan.Peek() != '"' {
		//FIXME: return a type, print string
		return fmt.Errorf("data is not a JSON string")
	}
	discardCurrent(scan)

	//FIXME: unescape will require copying data to a buffer
	inEscape := false
	for len(scan.Data()) > 0 {
		if inEscape {
			switch scan.Peek() {
			case 'u':
				// FIXME: expect 4 hexits
			}
			//FIXME: escape codes
			inEscape = false
		} else {
			switch scan.Peek() {
			case '\\':
				inEscape = true
			case '"':
				value.set(string(scan.Save()))
				discardCurrent(scan)
				return nil
			}
		}
		scan.Advance()
	}
	//FIXME: use a type, don't consume if error
	return fmt.Errorf("unterminated JSON string (%q)", string(scan.Save()))
}

func (value String) Empty() bool {
	return value.get() == ""
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

func (value *Number) Parse(data []byte) error {
	scan := NewByteScanner(data)
	if err := value.ParseStream(scan); err != nil {
		return err
	}
	if len(scan.Data()) != 0 {
		return fmt.Errorf("found trailing data in input: %q", string(scan.Data()))
	}
	return nil
}

func (value *Number) ParseStream(scan *ByteScanner) error {
	//FIXME: need a type to decode into or functions
	panic("Not implemented")
	return nil
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

func (value *Bool) Parse(data []byte) error {
	scan := NewByteScanner(data)
	if err := value.ParseStream(scan); err != nil {
		return err
	}
	if len(scan.Data()) != 0 {
		return fmt.Errorf("found trailing data in input: %q", string(scan.Data()))
	}
	return nil
}

func (value *Bool) ParseStream(scan *ByteScanner) error {
	if len(scan.Data()) >= len(trueBytes) {
		mightBe := true
		for _, b := range trueBytes {
			if scan.Peek() != rune(b) {
				mightBe = false
				break
			}
			scan.Advance()
		}
		if mightBe {
			//FIXME: test what happens at end-of-buffer
			if isValueDelim(scan.Peek()) {
				*value = true
				scan.Save()
				return nil
			}
		}
	}
	scan.Reset()
	if len(scan.Data()) >= len(falseBytes) {
		mightBe := true
		for _, b := range falseBytes {
			if scan.Peek() != rune(b) {
				mightBe = false
				break
			}
			scan.Advance()
		}
		if mightBe {
			//FIXME: test what happens at end-of-buffer
			if isValueDelim(scan.Peek()) {
				*value = false
				scan.Save()
				return nil
			}
		}
	}
	return fmt.Errorf("data is not a JSON bool") //FIXME: print value
}

func (value Bool) Empty() bool {
	return value == false
}

//FIXME: do we need this?
type Null struct{}

func (Null) Render(buf *bytes.Buffer) error {
	return write(buf, nullBytes)
}

func (value Null) Parse(data []byte) error {
	return nil
}

func (value Null) ParseStream(scan *ByteScanner) error {
	return nil
}

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

func (value Object) Parse(data []byte) error {
	glog.Errorf("TIM: object.Parse %#v", value)
	scan := NewByteScanner(data)
	if err := value.ParseStream(scan); err != nil {
		return err
	}
	if len(scan.Data()) != 0 {
		return fmt.Errorf("found trailing data in input: %q", string(scan.Data()))
	}
	return nil
}

func (value Object) ParseStream(scan *ByteScanner) error {
	// So we don't have to do multiple linear searches during decode.
	fieldMap := map[string]*NamedValue{}
	for i := range value {
		fieldMap[value[i].Name.get()] = &value[i]
	}

	if len(scan.Data()) < 2 || scan.Peek() != '{' {
		//FIXME: use a type, print value...
		return fmt.Errorf("data is not a JSON object")
	}
	discardCurrent(scan)

	for len(scan.Data()) > 0 {
		discardWhitespace(scan)
		if scan.Peek() == '}' {
			discardCurrent(scan)
			return nil
		}

		// Read the key.
		p := new(string)
		key := NewString(func() string { return *p }, func(s string) { *p = s })
		if err := key.ParseStream(scan); err != nil {
			return err //FIXME
		}
		glog.Errorf("TIM: key was %s", key.get())
		field := fieldMap[key.get()]
		if field == nil {
			return fmt.Errorf("unknown field %s", key.get()) //FIXME: save the string
		}
		glog.Errorf("TIM: value is %T", field.Value)

		// Read the colon.
		discardWhitespace(scan)
		if scan.Peek() != ':' {
			return fmt.Errorf("data is not a JSON object")
		}
		discardCurrent(scan)
		discardWhitespace(scan)

		// Read the value.
		if err := field.Value.ParseStream(scan); err != nil {
			return err
		}

		// Prep for the next field.
		//FIXME: test what happens at end-of-buffer
		if isValueDelim(scan.Peek()) {
			discardWhitespace(scan)
			if scan.Peek() == ',' {
				discardCurrent(scan)
			}
		} else {
			return fmt.Errorf("data is not a JSON object") //FIXME: parse error
		}
	}
	return fmt.Errorf("data is not a JSON object") //FIXME: parse error
}

func discardCurrent(scan *ByteScanner) {
	scan.Advance()
	scan.Save()
}

func discardWhitespace(scan *ByteScanner) {
	for len(scan.Data()) > 0 && isSpace(scan.Peek()) {
		scan.Advance()
	}
	scan.Save()
}

func isSpace(r rune) bool {
	switch r {
	case ' ', '\t', '\n', '\r':
		return true
	}
	return false
}

func isValueDelim(r rune) bool {
	if isSpace(r) {
		return true
	}
	switch r {
	case ',', '}', ']':
		return true
	}
	return false
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

func (value Array) Parse(data []byte) error {
	scan := NewByteScanner(data)
	if err := value.ParseStream(scan); err != nil {
		return err
	}
	if len(scan.Data()) != 0 {
		return fmt.Errorf("found trailing data in input: %q", string(scan.Data()))
	}
	return nil
}

func (value Array) ParseStream(scan *ByteScanner) error {
	//FIXME: need a type to decode into or functions
	panic("Not implemented")
	return nil
}

func (value Array) Empty() bool {
	return len(value) == 0
}

type Raw string

func (value Raw) Render(buf *bytes.Buffer) error {
	return writeString(buf, string(value))
}

func (value Raw) Parse(data []byte) error {
	return nil
}

func (value Raw) ParseStream(scan *ByteScanner) error {
	return nil
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
