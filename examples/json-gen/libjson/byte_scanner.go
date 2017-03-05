package libjson

import "unicode/utf8"

// ByteScanner walks through a []byte, one rune at a time.  Callers can parse the
// results, and save sub-slices of the input data for further work.  Once data
// has been consumed, it is no longer available unless Reset is called.
//
// The typical pattern is to Peek at the data and Advance util some delimiter is
// found, and then Save the slice-in-progress.  Note that if the ByteScanner is
// advanced until it runs out of data, the slice-in-process is still valid and
// saveable.
type ByteScanner struct {
	data    []byte
	seg     []byte
	segSize int
}

// NewByteScanner returns a ByteScanner that will re-slice the provided data
// slice.
func NewByteScanner(data []byte) *ByteScanner {
	bs := &ByteScanner{
		data: data,
		seg:  data,
	}
	return bs
}

// Data returns the unexamined bytes in the input data.
func (bs *ByteScanner) Data() []byte {
	return bs.data
}

// Peek returns the current rune but does not add it to the slice-in-progress.
func (bs *ByteScanner) Peek() rune {
	r, _ := utf8.DecodeRune(bs.data)
	return r
}

// Advance adds the current rune to the slice-in-progress and moves to the next
// rune.
func (bs *ByteScanner) Advance() {
	_, size := utf8.DecodeRune(bs.data)
	bs.segSize += size
	bs.data = bs.data[size:]
}

// Reset moves the scanner back to the start of the slice-in-progress, undoing
// any calls to Advance since the last Save.
func (bs *ByteScanner) Reset() {
	bs.data = bs.seg
	bs.segSize = 0
}

// Save returns the current slice-in-progress, and starts a new slice at
// the current rune.
func (bs *ByteScanner) Save() []byte {
	seg := bs.seg[:bs.segSize]
	bs.seg = bs.data
	bs.segSize = 0
	return seg
}

/* example
func main() {
	raw := []byte("Hello, 世界;This is;an array of bytes;that says \"世界;\".")

	cord := [][]byte{}
	bs := NewByteScanner(raw)
	for len(bs.Data()) > 0 {
		r := bs.Peek()
		if r == ';' {
			// save the current sub-slice without the delimiter
			cord = append(cord, bs.Save())
			// discard the delimiter
			bs.Advance()
			bs.SaveSegment()
		}
		bs.Advance()
	}
	// handle any trailing characters
	if s := bs.Save(); len(s) != 0 {
		cord = append(cord, s)
	}
	fmt.Printf("%q\n", cord)
}
*/
