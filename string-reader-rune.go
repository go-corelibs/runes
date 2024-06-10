// Copyright 2024 The Go-CoreLibs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package runes

import (
	"errors"
	"io"
	"unicode/utf8"
)

// ReadRuneAt is a convenience method combining Seek and ReadRune into one
// operation. The index argument is always relative to the start of the
// slice, equivalent to Seek(index, io.SeekStart)
//
// ReadRuneAt was added by go-corelibs
func (r *StringReader) ReadRuneAt(index int64) (ch rune, size int, err error) {
	r.prevRune = -1
	if index < 0 {
		return 0, 0, errors.New("StringReader.ReadRuneAt: negative position")
	} else if index >= int64(len(r.s)) {
		return 0, 0, io.EOF
	}
	r.i = index
	r.prevRune = int(r.i)
	if c := r.s[r.i]; c < utf8.RuneSelf {
		r.i++
		return rune(c), 1, nil
	}
	ch, size = utf8.DecodeRuneInString(r.s[r.i:])
	r.i += int64(size)
	return
}

// ReadPrevRuneFrom is a convenience method combining Seek and ReadRune into one
// operation. The index argument is always relative to the start of the
// slice, equivalent to Seek(index, io.SeekStart)
//
// ReadPrevRuneFrom was added by go-corelibs
func (r *StringReader) ReadPrevRuneFrom(index int64) (ch rune, size int, err error) {
	r.prevRune = -1
	if index <= 0 {
		return 0, 0, errors.New("StringReader.ReadPrevRuneFrom: zero or negative position")
	} else if index >= int64(len(r.s)) {
		return 0, 0, io.EOF
	}
	if r.s[index-1] < utf8.RuneSelf {
		r.i -= 1
		return rune(r.s[index-1]), 1, nil
	}
	ch, size = utf8.DecodeLastRuneInString(r.s[:index])
	r.i -= int64(size)
	return
}

// ReadNextRuneFrom is a convenience method combining Seek and ReadRune into one
// operation. The index argument is always relative to the start of the
// slice, equivalent to Seek(index, io.SeekStart)
//
// ReadNextRuneFrom was added by go-corelibs
func (r *StringReader) ReadNextRuneFrom(index int64) (ch rune, size int, err error) {
	r.prevRune = -1
	if index < 0 {
		return 0, 0, errors.New("StringReader.ReadNextRuneFrom: negative position")
	} else if index >= int64(len(r.s)) {
		return 0, 0, io.EOF
	}
	r.i = index
	r.prevRune = int(r.i)
	// move this forward once
	if c := r.s[r.i]; c < utf8.RuneSelf {
		r.i++ // got this one
	} else {
		_, sz := utf8.DecodeRuneInString(r.s[r.i:])
		r.i += int64(sz)
	}
	// find next one
	if next := r.s[r.i]; next < utf8.RuneSelf {
		// returning the next one
		return rune(next), 1, nil
	}
	ch, size = utf8.DecodeRuneInString(r.s[r.i:])
	r.i += int64(size)
	return
}

// ReadRuneSlice is a convenience method combining Seek and then ReadRune
// operations accumulating the requested count of runes, starting at the
// index given. The index argument is always relative to the start of the
// slice, equivalent to Seek(index, io.SeekStart). The count argument is
// exclusive, meaning start at the index and stop at index+count, equivalent
// to the slice index syntax of bytes[index:index+count]
//
// ReadRuneSlice was added by go-corelibs
func (r *StringReader) ReadRuneSlice(index, count int64) (slice []rune, size int, err error) {
	r.prevRune = -1
	if index < 0 {
		return nil, 0, errors.New("StringReader.ReadRuneSlice: negative position")
	} else if count < 1 {
		return nil, 0, errors.New("StringReader.ReadRuneSlice: zero or negative count")
	} else if index >= int64(len(r.s)) {
		return nil, 0, io.EOF
	}
	r.i = index

	length := int64(len(r.s))

	slice = make([]rune, count)
	track := int64(0)
	for idx := index; idx < length; {
		if track == count {
			break
		} else if r.i >= length {
			return nil, 0, io.EOF
		}
		r.prevRune = int(r.i)
		if c := r.s[r.i]; c < utf8.RuneSelf {
			slice[track] = rune(c)
			track += 1
			size += 1
			r.i++
		} else {
			ch, sz := utf8.DecodeRuneInString(r.s[r.i:])
			slice[track] = ch
			track += 1
			size += sz
			r.i += int64(sz)
		}
	}
	return
}
