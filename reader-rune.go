// Copyright 2024 The Go-CoreLibs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package runes

import (
	"errors"
	"io"
	"slices"
)

// ReadRuneAt is a convenience method combining Seek and ReadRune into one
// operation. The index argument is always relative to the start of the
// slice, equivalent to Seek(index, io.SeekStart)
//
// ReadRuneAt was added by go-corelibs
func (r *Reader) ReadRuneAt(index int64) (ch rune, size int, err error) {
	r.prevRune = -1
	if index < 0 {
		return 0, 0, errors.New("Reader.ReadRuneAt: negative position")
	} else if index >= int64(len(r.s)) {
		return 0, 0, io.EOF
	}
	r.prevRune = int(r.i)
	r.i = index // seek
	ch = r.s[r.i]
	size = 1
	r.i += 1 // move
	return
}

// ReadPrevRuneFrom is a convenience method combining Seek and ReadRune into one
// operation. The index argument is always relative to the start of the
// slice, equivalent to Seek(index, io.SeekStart)
//
// ReadPrevRuneFrom was added by go-corelibs
func (r *Reader) ReadPrevRuneFrom(index int64) (ch rune, size int, err error) {
	r.prevRune = -1
	if index <= 0 {
		return 0, 0, errors.New("Reader.ReadPrevRuneFrom: zero or negative position")
	} else if index >= int64(len(r.s)) || index == 0 {
		return 0, 0, io.EOF
	}
	r.prevRune = int(r.i)
	r.i = index // seek
	ch = r.s[r.i-1]
	size = 1
	r.i -= 1 // move
	return
}

// ReadNextRuneFrom is a convenience method combining Seek and ReadRune into one
// operation. The index argument is always relative to the start of the
// slice, equivalent to Seek(index, io.SeekStart)
//
// ReadNextRuneFrom was added by go-corelibs
func (r *Reader) ReadNextRuneFrom(index int64) (ch rune, size int, err error) {
	r.prevRune = -1
	if index < 0 {
		return 0, 0, errors.New("Reader.ReadNextRuneFrom: negative position")
	} else if index+1 >= int64(len(r.s)) {
		return 0, 0, io.EOF
	}
	r.prevRune = int(r.i)
	r.i = index
	ch = r.s[r.i+1]
	size = 1
	r.i += 1
	return
}

// ReadRuneSlice is a convenience method combining Seek and then ReadRune
// operations accumulating the requested count of runes, starting at the
// index given. The index argument is always relative to the start of the
// slice, equivalent to Seek(index, io.SeekStart). The count argument is
// exclusive, meaning start at the index and stop at index+count, equivalent
// to the slice index syntax of bytes[index:index+count]
//
// The size returned is the number of runes returned, not the number of bytes.
//
// ReadRuneSlice was added by go-corelibs
func (r *Reader) ReadRuneSlice(index, count int64) (slice []rune, size int, err error) {
	r.prevRune = -1
	if index < 0 {
		return nil, 0, errors.New("Reader.ReadRuneSlice: negative position")
	} else if count < 1 {
		return nil, 0, errors.New("Reader.ReadRuneSlice: zero or negative count")
	} else if index >= int64(len(r.s)) {
		return nil, 0, io.EOF
	}
	r.i = index

	length := int64(r.Len())

	slice = make([]rune, count)
	track := int64(0)
	for idx := index; idx < length; {
		if track == count {
			break
		}
		r.prevRune = int(r.i)
		ch := r.s[r.i]
		slice[track] = ch
		track += 1
		size += 1
		r.i += 1 // rune slice, not bytes
	}
	slices.Clip(slice)
	return
}
