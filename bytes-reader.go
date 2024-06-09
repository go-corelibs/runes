// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runes

import (
	"errors"
	"io"
	"unicode/utf8"
)

// A BytesReader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,
// io.ByteScanner, and io.RuneScanner interfaces by reading from
// a byte slice.
// Unlike a [Buffer], a BytesReader is read-only and supports seeking.
// The zero value for BytesReader operates like a BytesReader of an empty slice.
type BytesReader struct {
	s        []byte
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

// Len returns the number of bytes of the unread portion of the
// slice.
func (r *BytesReader) Len() int {
	if r.i >= int64(len(r.s)) {
		return 0
	}
	return int(int64(len(r.s)) - r.i)
}

// Size returns the original length of the underlying byte slice.
// Size is the number of bytes available for reading via [BytesReader.ReadAt].
// The result is unaffected by any method calls except [BytesReader.Reset].
func (r *BytesReader) Size() int64 { return int64(len(r.s)) }

// Read implements the [io.Reader] interface.
func (r *BytesReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

// ReadAt implements the [io.ReaderAt] interface.
func (r *BytesReader) ReadAt(b []byte, off int64) (n int, err error) {
	// cannot modify state - see io.ReaderAt
	if off < 0 {
		return 0, errors.New("BytesReader.ReadAt: negative offset")
	}
	if off >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[off:])
	if n < len(b) {
		err = io.EOF
	}
	return
}

// ReadByte implements the [io.ByteReader] interface.
func (r *BytesReader) ReadByte() (byte, error) {
	r.prevRune = -1
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	b := r.s[r.i]
	r.i++
	return b, nil
}

// UnreadByte complements [BytesReader.ReadByte] in implementing the [io.ByteScanner] interface.
func (r *BytesReader) UnreadByte() error {
	if r.i <= 0 {
		return errors.New("BytesReader.UnreadByte: at beginning of slice")
	}
	r.prevRune = -1
	r.i--
	return nil
}

// ReadRune implements the [io.RuneReader] interface.
func (r *BytesReader) ReadRune() (ch rune, size int, err error) {
	if r.i >= int64(len(r.s)) {
		r.prevRune = -1
		return 0, 0, io.EOF
	}
	r.prevRune = int(r.i)
	if c := r.s[r.i]; c < utf8.RuneSelf {
		r.i++
		return rune(c), 1, nil
	}
	ch, size = utf8.DecodeRune(r.s[r.i:])
	r.i += int64(size)
	return
}

// UnreadRune complements [BytesReader.ReadRune] in implementing the [io.RuneScanner] interface.
func (r *BytesReader) UnreadRune() error {
	if r.i <= 0 {
		return errors.New("BytesReader.UnreadRune: at beginning of slice")
	}
	if r.prevRune < 0 {
		return errors.New("BytesReader.UnreadRune: previous operation was not ReadRune")
	}
	r.i = int64(r.prevRune)
	r.prevRune = -1
	return nil
}

// Seek implements the [io.Seeker] interface.
func (r *BytesReader) Seek(offset int64, whence int) (int64, error) {
	r.prevRune = -1
	var abs int64
	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = r.i + offset
	case io.SeekEnd:
		abs = int64(len(r.s)) + offset
	default:
		return 0, errors.New("BytesReader.Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("BytesReader.Seek: negative position")
	}
	r.i = abs
	return abs, nil
}

// WriteTo implements the [io.WriterTo] interface.
func (r *BytesReader) WriteTo(w io.Writer) (n int64, err error) {
	r.prevRune = -1
	if r.i >= int64(len(r.s)) {
		return 0, nil
	}
	b := r.s[r.i:]
	m, err := w.Write(b)
	if m > len(b) {
		panic("BytesReader.WriteTo: invalid Write count")
	}
	r.i += int64(m)
	n = int64(m)
	if m != len(b) && err == nil {
		err = io.ErrShortWrite
	}
	return
}

// Reset resets the [BytesReader.BytesReader] to be reading from b.
func (r *BytesReader) Reset(b []byte) { *r = BytesReader{b, 0, -1} }

// NewBytesReader returns a new [BytesReader.BytesReader] reading from b.
func NewBytesReader(b []byte) *BytesReader { return &BytesReader{b, 0, -1} }
