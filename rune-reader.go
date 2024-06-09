// Copyright 2024 The Go-CoreLibs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package runes

import (
	"io"
)

// RuneReader defines the interface for additional rune-specific features when
// reading data from a string, bytes or rune slices
type RuneReader interface {
	io.Reader
	io.ReaderAt
	io.WriterTo
	io.Seeker
	io.ByteScanner
	io.RuneScanner

	// Len returns the byte length of the underlying memory
	Len() int

	// Size returns the length of the underlying slice
	Size() int64

	// ReadRuneAt seeks to the index given and reads the rune at that position
	ReadRuneAt(index int64) (ch rune, size int, err error)

	// ReadPrevRuneFrom seeks to the index given and reads the rune previous to
	// that position
	ReadPrevRuneFrom(index int64) (ch rune, size int, err error)

	// ReadNextRuneFrom seeks to the index given and reads the next rune after
	// that position
	ReadNextRuneFrom(index int64) (ch rune, size int, err error)

	// ReadRuneSlice seeks to the index given, and starts accumulating runes,
	// up to the count requested, returning a slice and the total size
	//
	// For byte and string readers, size is the number of bytes in the rune
	// slice. For the rune reader, the size is always 1, which works because the
	// underlying data type is just a slice of runes (no decoding of multi-bytes
	// needed)
	ReadRuneSlice(index, count int64) (slice []rune, size int, err error)
}

// NewRuneReader is a generic wrapper around constructing a NewBytesReader,
// NewStringReader or NewRunesReader depending on the input type given and
// returned as a RuneReader
func NewRuneReader[V []rune | []byte | string](input V) (rb RuneReader) {
	v := &input
	switch t := interface{}(v).(type) {
	case *[]byte:
		return NewBytesReader(*t)
	case *string:
		return NewStringReader(*t)
	case *[]rune:
		return NewRunesReader(*t)
	default:
		panic("the universe is broken")
	}
}
