// Copyright 2024 The Go-CoreLibs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package runes_test

import (
	"math/rand"
	"testing"

	. "github.com/go-corelibs/runes"
)

const N = 10000       // make this bigger for a larger (and slower) test
var testString string // test data for write tests
var testBytes []byte  // test data; same as testString but as a slice.
var testRunes []rune  // test data; same as testString but as a slice.
var testIndex []int64

func init() {
	testBytes = make([]byte, N)
	testRunes = make([]rune, N)
	testIndex = make([]int64, N)
	for i := 0; i < N; i++ {
		testBytes[i] = 'a' + byte(i%26)
		testRunes[i] = 'a' + rune(i%26)
		testIndex[i] = int64(rand.Intn(N-10) + 1) // don't test read prev from 0
	}
	testString = string(testBytes)
}

// Byte Reader

func BenchmarkBytesReaderReadRuneAt(b *testing.B) {
	r := NewBytesReader(testBytes)
	for i := 0; i < N; i++ {
		_, _, _ = r.ReadRuneAt(testIndex[i])
	}
}

func BenchmarkBytesReaderReadPrevRuneFrom(b *testing.B) {
	r := NewBytesReader(testBytes)
	for i := 0; i < N; i++ {
		_, _, _ = r.ReadPrevRuneFrom(testIndex[i])
	}
}

func BenchmarkBytesReaderReadNextRuneFrom(b *testing.B) {
	r := NewBytesReader(testBytes)
	for i := 0; i < N; i++ {
		_, _, _ = r.ReadNextRuneFrom(testIndex[i])
	}
}

func BenchmarkBytesReaderReadSliceRuneFrom(b *testing.B) {
	r := NewBytesReader(testBytes)
	for i := 0; i < N/4; i++ {
		_, _, _ = r.ReadRuneSlice(testIndex[i], 10)
	}
}

// String Reader

func BenchmarkStringReaderReadRuneAt(b *testing.B) {
	r := NewStringReader(testString)
	for i := 0; i < N; i++ {
		_, _, _ = r.ReadRuneAt(testIndex[i])
	}
}

func BenchmarkStringReaderReadPrevRuneFrom(b *testing.B) {
	r := NewStringReader(testString)
	for i := 0; i < N; i++ {
		_, _, _ = r.ReadPrevRuneFrom(testIndex[i])
	}
}

func BenchmarkStringReaderReadNextRuneFrom(b *testing.B) {
	r := NewStringReader(testString)
	for i := 0; i < N; i++ {
		_, _, _ = r.ReadNextRuneFrom(testIndex[i])
	}
}

func BenchmarkStringReaderReadSliceRuneFrom(b *testing.B) {
	r := NewStringReader(testString)
	for i := 0; i < N/4; i++ {
		_, _, _ = r.ReadRuneSlice(testIndex[i], 10)
	}
}

// Runes Reader

func BenchmarkRunesReaderReadRuneAt(b *testing.B) {
	r := NewRunesReader(testRunes)
	for i := 0; i < N; i++ {
		_, _, _ = r.ReadRuneAt(testIndex[i])
	}
}

func BenchmarkRunesReaderReadPrevRuneFrom(b *testing.B) {
	r := NewRunesReader(testRunes)
	for i := 0; i < N; i++ {
		_, _, _ = r.ReadPrevRuneFrom(testIndex[i])
	}
}

func BenchmarkRunesReaderReadNextRuneFrom(b *testing.B) {
	r := NewRunesReader(testRunes)
	for i := 0; i < N; i++ {
		_, _, _ = r.ReadNextRuneFrom(testIndex[i])
	}
}

func BenchmarkRunesReaderReadSliceRuneFrom(b *testing.B) {
	r := NewRunesReader(testRunes)
	for i := 0; i < N/4; i++ {
		_, _, _ = r.ReadRuneSlice(testIndex[i], 10)
	}
}
