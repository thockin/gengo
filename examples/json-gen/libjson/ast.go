package libjson

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

type Value interface {
	Render(buf *bytes.Buffer) error
	Parse(data []byte) error
	ParseStream(scan *ByteScanner) error
	Empty() bool
}

type Nullable struct {
	inner   Value
	setNull func(bool) (Value, error)
}

func NewNullable(inner Value, setNull func(bool) (Value, error)) *Nullable {
	return &Nullable{
		inner:   inner,
		setNull: setNull,
	}
}

var nullBytes = []byte("null")

func (value *Nullable) Render(buf *bytes.Buffer) error {
	if value.inner == nil {
		return write(buf, nullBytes)
	}
	return value.inner.Render(buf)
}

func (value *Nullable) Parse(data []byte) error {
	scan := NewByteScanner(data)
	if err := value.ParseStream(scan); err != nil {
		return err
	}
	if len(scan.Data()) != 0 {
		return fmt.Errorf("found trailing data in input: %q", string(scan.Data()))
	}
	return nil
}

func (value *Nullable) ParseStream(scan *ByteScanner) error {
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
			if len(scan.Data()) == 0 || isValueDelim(scan.Peek()) {
				if _, err := value.setNull(true); err != nil {
					return err //FIXME:
				} else {
					value.inner = nil
				}
				scan.Save()
				return nil
			}
		}
	}
	if jv, err := value.setNull(false); err != nil {
		return err //FIXME:
	} else {
		if jv == nil {
			return fmt.Errorf("Nullable type got nil value")
		}
		value.inner = jv
	}
	scan.Reset()
	return value.inner.ParseStream(scan)
}

func (value *Nullable) Empty() bool {
	return value.inner == nil
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
	return writeString(buf, `"`+escape(value.get())+`"`)
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

//FIXME:
// When unmarshaling quoted strings, invalid UTF-8 or
// invalid UTF-16 surrogate pairs are not treated as an error.
// Instead, they are replaced by the Unicode replacement
// character U+FFFD.
func (value String) ParseStream(scan *ByteScanner) error {
	str, err := scanString(scan)
	if err != nil {
		return err
	}
	value.set(str)
	return nil
}

func scanString(scan *ByteScanner) (string, error) {
	if len(scan.Data()) < 2 || scan.Peek() != '"' {
		//FIXME: return a type, print string
		return "", fmt.Errorf("data is not a JSON string")
	}
	discardCurrent(scan)

	var buf bytes.Buffer
	for len(scan.Data()) > 0 {
		if scan.Peek() == '"' {
			if _, err := buf.Write(scan.Save()); err != nil {
				return "", err //FIXME:
			}
			discardCurrent(scan)
			return buf.String(), nil
		}
		if scan.Peek() != '\\' {
			scan.Advance()
			continue
		}
		if _, err := buf.Write(scan.Save()); err != nil {
			return "", err //FIXME:
		}
		if err := unescape(scan, &buf); err != nil {
			return "", err //FIXME
		}
	}

	//FIXME: use a type, don't consume if error
	return "", fmt.Errorf("unterminated JSON string (%q)", string(scan.Save()))
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
		//FIXME: errors
		switch r {
		case '"':
			buf.WriteString(`\"`)
		case '\\':
			buf.WriteString(`\\`)
		// The JSON spec says to escape this, but Go's stdlib does not
		//case '/':
		//buf.WriteString(`\/`)
		case '\b':
			buf.WriteString(`\b`)
		case '\f':
			buf.WriteString(`\f`)
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
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

func unescape(scan *ByteScanner, buf *bytes.Buffer) error {
	discardCurrent(scan) // get rid of the '\'

	var err error
	switch scan.Peek() {
	case utf8.RuneError:
		return fmt.Errorf("ran out of data") //FIXME:
	case '"':
		_, err = buf.WriteRune('"')
		discardCurrent(scan)
	case '\\':
		_, err = buf.WriteRune('\\')
		discardCurrent(scan)
	// The JSON spec says to escape this, but Go's stdlib does not
	//case '/':
	//_, err = buf.WriteRune('/')
	//discardCurrent(scan)
	case 'b':
		_, err = buf.WriteRune('\b')
		discardCurrent(scan)
	case 'f':
		_, err = buf.WriteRune('\f')
		discardCurrent(scan)
	case 'n':
		_, err = buf.WriteRune('\n')
		discardCurrent(scan)
	case 'r':
		_, err = buf.WriteRune('\r')
		discardCurrent(scan)
	case 't':
		_, err = buf.WriteRune('\t')
		discardCurrent(scan)
	case 'u':
		err = unescapeHex(scan, buf)
	default:
		err = fmt.Errorf("unknown escape: %v", scan.Peek())
	}
	return err //FIXME:
}

func unescapeHex(scan *ByteScanner, buf *bytes.Buffer) error {
	discardCurrent(scan) // get rund of the 'u'

	if len(scan.Data()) < 4 {
		return fmt.Errorf("ran out of data") //FIXME:
	}
	u16, err := strconv.ParseUint(string(scan.Data()[:4]), 16, 16)
	if err != nil {
		return err //FIXME:
	}
	r1 := rune(u16)
	for i := 0; i < 4; i++ { // get rid of the hex code
		discardCurrent(scan)
	}

	// this technique is borrowed from Go's json lib
	if !utf16.IsSurrogate(r1) {
		// single escape
		_, err = buf.WriteRune(r1)
		return err
	}

	// expect a second escape
	if len(scan.Data()) < 4 {
		return fmt.Errorf("ran out of data") //FIXME:
	}
	u16, err = strconv.ParseUint(string(scan.Data()[:4]), 16, 16)
	if err != nil {
		return err //FIXME:
	}
	r2 := rune(u16)
	for i := 0; i < 4; i++ { // get rid of the hex code
		discardCurrent(scan)
	}

	// see if it was a valid utf8 pair
	if r := utf16.DecodeRune(r1, r2); r != unicode.ReplacementChar {
		// it was valid
		var rbuf [8]byte
		n := utf8.EncodeRune(rbuf[:], r)
		_, err := buf.Write(rbuf[:n])
		return err
	}
	// invalid pair
	_, err = buf.WriteRune(unicode.ReplacementChar)
	return err
}

type Number struct {
	isFloat bool
	bits    int
	get     func() float64
	set     func(f float64)
}

func newNumber(isFloat bool, bits int, get func() float64, set func(f float64)) Number {
	return Number{
		isFloat: isFloat,
		bits:    bits,
		get:     get,
		set:     set,
	}
}

func NewInt(get func() float64, set func(f float64)) Number {
	return newNumber(false, 64, get, set)
}

func NewFloat(bits int, get func() float64, set func(f float64)) Number {
	return newNumber(true, bits, get, set)
}

func (value Number) Render(buf *bytes.Buffer) error {
	prec := 64
	if value.isFloat {
		prec = -1
	}
	return writeString(buf, strconv.FormatFloat(value.get(), 'g', prec, value.bits))
}

func (value Number) Parse(data []byte) error {
	scan := NewByteScanner(data)
	if err := value.ParseStream(scan); err != nil {
		return err
	}
	if len(scan.Data()) != 0 {
		return fmt.Errorf("found trailing data in input: %q", string(scan.Data()))
	}
	return nil
}

func (value Number) ParseStream(scan *ByteScanner) error {
	for len(scan.Data()) > 0 && !isValueDelim(scan.Peek()) {
		scan.Advance()
	}
	//FIXME: is this good enough or do I have to hand-code the float parser?
	if f, err := strconv.ParseFloat(string(scan.Save()), value.bits); err != nil {
		return err
	} else {
		value.set(f)
	}
	return nil
}

func (value Number) Empty() bool {
	return value.get() == 0
}

type Bool struct {
	get func() bool
	set func(b bool)
}

func NewBool(get func() bool, set func(b bool)) Bool {
	return Bool{
		get: get,
		set: set,
	}
}

var trueBytes = []byte("true")
var falseBytes = []byte("false")

func (value Bool) Render(buf *bytes.Buffer) error {
	if value.get() {
		return write(buf, trueBytes)
	}
	return write(buf, falseBytes)
}

func (value Bool) Parse(data []byte) error {
	scan := NewByteScanner(data)
	if err := value.ParseStream(scan); err != nil {
		return err
	}
	if len(scan.Data()) != 0 {
		return fmt.Errorf("found trailing data in input: %q", string(scan.Data()))
	}
	return nil
}

func (value Bool) ParseStream(scan *ByteScanner) error {
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
			if len(scan.Data()) == 0 || isValueDelim(scan.Peek()) {
				value.set(true)
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
			if len(scan.Data()) == 0 || isValueDelim(scan.Peek()) {
				value.set(false)
				scan.Save()
				return nil
			}
		}
	}
	return fmt.Errorf("data is not a JSON bool") //FIXME: print value
}

func (value Bool) Empty() bool {
	return value.get() == false
}

type Object []NamedValue

func NewObject() Object {
	return Object{}
}

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
		field := fieldMap[key.get()]
		if field == nil {
			return fmt.Errorf("unknown field %s", key.get()) //FIXME: save the string
		}

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
		discardWhitespace(scan)
		//FIXME: test what happens at end-of-buffer
		if r := scan.Peek(); r == ',' || r == '}' {
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

type Array struct {
	get func() ([]Value, error)
	add func() Value
	//FIXME: add len()
}

func NewArray(get func() ([]Value, error), add func() Value) Array {
	return Array{
		get: get,
		add: add,
	}
}

func (value Array) Render(buf *bytes.Buffer) error {
	ar, err := value.get()
	if err != nil {
		return err
	} else if ar == nil {
		if err := write(buf, nullBytes); err != nil {
			return err
		}
		return nil
	}

	if err := writeString(buf, "["); err != nil {
		return err
	}
	for i, v := range ar {
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
	if len(scan.Data()) < 2 || scan.Peek() != '[' {
		//FIXME: use a type, print value...
		return fmt.Errorf("data is not a JSON array: %v", string(scan.Data()))
	}
	discardCurrent(scan)

	// This assumes that the underlying slice is empty.
	for len(scan.Data()) > 0 {
		discardWhitespace(scan)
		if scan.Peek() == ']' {
			discardCurrent(scan)
			return nil
		}

		// Read the value.
		elem := value.add()
		if err := elem.ParseStream(scan); err != nil {
			return err
		}

		// Prep for the next field.
		discardWhitespace(scan)
		//FIXME: test what happens at end-of-buffer
		if r := scan.Peek(); r == ',' || r == ']' {
			if scan.Peek() == ',' {
				discardCurrent(scan)
			}
		} else {
			return fmt.Errorf("data is not a JSON array") //FIXME: parse error
		}
	}
	return fmt.Errorf("data is not a JSON array") //FIXME: parse error
}

func (value Array) Empty() bool {
	ar, _ := value.get()
	return len(ar) == 0
}

type Bytes struct {
	get func() []byte
	set func([]byte)
}

func NewBytes(get func() []byte, set func([]byte)) Bytes {
	return Bytes{
		get: get,
		set: set,
	}
}

func (value Bytes) Render(buf *bytes.Buffer) error {
	if _, err := buf.WriteRune('"'); err != nil {
		return err
	}
	b64 := base64.NewEncoder(base64.StdEncoding, buf)
	if _, err := b64.Write(value.get()); err != nil {
		return err
	}
	if err := b64.Close(); err != nil {
		return err
	}
	if _, err := buf.WriteRune('"'); err != nil {
		return err
	}
	return nil
}

func (value Bytes) Parse(data []byte) error {
	scan := NewByteScanner(data)
	if err := value.ParseStream(scan); err != nil {
		return err
	}
	if len(scan.Data()) != 0 {
		return fmt.Errorf("found trailing data in input: %q", string(scan.Data()))
	}
	return nil
}

func (value Bytes) ParseStream(scan *ByteScanner) error {
	if len(scan.Data()) < 2 {
		//FIXME: use a type, print value...
		return fmt.Errorf("data is not a JSON array: %v", string(scan.Data()))
	}
	switch scan.Peek() {
	case '[':
		// Handle input like `[ 49, 50, 51 ]`
		return value.parseFromArray(scan)
	case '"':
		// Handle input like `"MTIz"`
		return value.parseFromString(scan)
	default:
		//FIXME: use a type, print value...
		return fmt.Errorf("data is not a JSON array: %v", string(scan.Data()))
	}
}

func (value Bytes) parseFromArray(scan *ByteScanner) error {
	bs := []byte{}

	get := func() ([]Value, error) {
		// This will not get called.
		return nil, nil
	}
	add := func() Value {
		bs = append(bs, byte(0))
		elem := &bs[len(bs)-1]
		return NewInt(
			func() float64 {
				return float64(*elem)
			},
			func(f float64) {
				*elem = byte(f)
			})
	}
	return NewArray(get, add).ParseStream(scan)
}

func (value Bytes) parseFromString(scan *ByteScanner) error {
	str, err := scanString(scan)
	if err != nil {
		return err
	}
	b64 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
	bs := make([]byte, len(str))
	if n, err := b64.Read(bs); err != nil {
		return err
	} else {
		bs = bs[:n]
	}
	value.set(bs)
	return nil
}

func (value Bytes) Empty() bool {
	return len(value.get()) == 0
}

type Map struct {
	add         func(k string) Value
	get         func() (map[string]Value, error)
	finishParse func()
	//FIXME: add len()
}

func NewMap(add func(k string) Value, get func() (map[string]Value, error), finishParse func()) Map {
	return Map{
		add:         add,
		get:         get,
		finishParse: finishParse,
	}
}

func (value Map) Render(buf *bytes.Buffer) error {
	mp, err := value.get()
	if err != nil {
		return err
	} else if mp == nil {
		if err := write(buf, nullBytes); err != nil {
			return err
		}
		return nil
	}

	if err := writeString(buf, "{"); err != nil {
		return err
	}
	keys := []string{}
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		if i > 0 {
			if err := writeString(buf, ","); err != nil {
				return err
			}
		}
		if err := writeString(buf, `"`+escape(k)+`"`); err != nil {
			return err
		}
		if err := writeString(buf, ":"); err != nil {
			return err
		}
		if err := mp[k].Render(buf); err != nil {
			return err
		}
	}
	if err := writeString(buf, "}"); err != nil {
		return err
	}
	return nil
}

func (value Map) Parse(data []byte) error {
	scan := NewByteScanner(data)
	if err := value.ParseStream(scan); err != nil {
		return err
	}
	if len(scan.Data()) != 0 {
		return fmt.Errorf("found trailing data in input: %q", string(scan.Data()))
	}
	return nil
}

func (value Map) ParseStream(scan *ByteScanner) error {
	if len(scan.Data()) < 2 || scan.Peek() != '{' {
		//FIXME: use a type, print value...
		return fmt.Errorf("data is not a JSON object: %v", string(scan.Data()))
	}
	discardCurrent(scan)

	// This assumes the map is empty.
	for len(scan.Data()) > 0 {
		discardWhitespace(scan)
		if scan.Peek() == '}' {
			discardCurrent(scan)
			value.finishParse()
			return nil
		}

		// Read the key.
		key, err := scanString(scan)
		if err != nil {
			return err //FIXME
		}

		// Read the colon.
		discardWhitespace(scan)
		if scan.Peek() != ':' {
			return fmt.Errorf("data is not a JSON object")
		}
		discardCurrent(scan)
		discardWhitespace(scan)

		// Read the value.
		jv := value.add(key)
		if err := jv.ParseStream(scan); err != nil {
			return err
		}

		// Prep for the next field.
		discardWhitespace(scan)
		//FIXME: test what happens at end-of-buffer
		if r := scan.Peek(); r == ',' || r == '}' {
			if scan.Peek() == ',' {
				discardCurrent(scan)
			}
		} else {
			return fmt.Errorf("data is not a JSON object") //FIXME: parse error
		}
	}
	return fmt.Errorf("data is not a JSON object") //FIXME: parse error
}

func (value Map) Empty() bool {
	mp, _ := value.get()
	return len(mp) == 0
}

type Raw string

func (value Raw) Render(buf *bytes.Buffer) error {
	return writeString(buf, string(value))
}

func (value Raw) Parse(data []byte) error {
	return nil
}

func (value Raw) ParseStream(scan *ByteScanner) error {
	panic("not implemented")
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
