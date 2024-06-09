// Copyright 2024 The Go-CoreLibs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package runes_test

import (
	"io"
	"testing"

	. "github.com/go-corelibs/runes"
)

//gocyclo:ignore
func TestRunesReader_ReadRuneAt(t *testing.T) {
	r := &Reader{}

	if ch, size, err := r.ReadRuneAt(-1); ch != 0 || size != 0 || (err == nil || err.Error() != "Reader.ReadRuneAt: negative position") {
		t.Errorf("ReadRuneAt: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	if ch, size, err := r.ReadRuneAt(0); ch != 0 || size != 0 || err != io.EOF {
		t.Errorf("ReadRuneAt: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]rune{'s', 't', 'u', 'f', 'f'})

	if ch, size, err := r.ReadRuneAt(0); ch != 's' || size != 1 || err != nil {
		t.Errorf("ReadRuneAt: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]rune{'日', '本', '語'})

	if ch, size, err := r.ReadRuneAt(0); ch != '日' || size != 1 || err != nil {
		t.Errorf("ReadRuneAt: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

}

//gocyclo:ignore
func TestRunesReader_ReadPrevRuneFrom(t *testing.T) {
	r := &Reader{}

	if ch, size, err := r.ReadPrevRuneFrom(-1); ch != 0 || size != 0 || (err == nil || err.Error() != "Reader.ReadPrevRuneFrom: zero or negative position") {
		t.Errorf("ReadPrevRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	if ch, size, err := r.ReadPrevRuneFrom(0); ch != 0 || size != 0 || (err == nil || err.Error() != "Reader.ReadPrevRuneFrom: zero or negative position") {
		t.Errorf("ReadPrevRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	if ch, size, err := r.ReadPrevRuneFrom(1); ch != 0 || size != 0 || err != io.EOF {
		t.Errorf("ReadPrevRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]rune{'s', 't', 'u', 'f', 'f'})

	if ch, size, err := r.ReadPrevRuneFrom(3); ch != 'u' || size != 1 || err != nil {
		t.Errorf("ReadPrevRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]rune{'日', '本', '語'})

	if ch, size, err := r.ReadPrevRuneFrom(1); ch != '日' || size != 1 || err != nil {
		t.Errorf("ReadPrevRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

}

//gocyclo:ignore
func TestRunesReader_ReadNextRuneFrom(t *testing.T) {
	r := &Reader{}

	if ch, size, err := r.ReadNextRuneFrom(-1); ch != 0 || size != 0 || (err == nil || err.Error() != "Reader.ReadNextRuneFrom: negative position") {
		t.Errorf("ReadNextRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	if ch, size, err := r.ReadNextRuneFrom(0); ch != 0 || size != 0 || err != io.EOF {
		t.Errorf("ReadNextRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]rune{'s', 't', 'u', 'f', 'f'})

	if ch, size, err := r.ReadNextRuneFrom(2); ch != 'f' || size != 1 || err != nil {
		t.Errorf("ReadNextRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]rune{'日', '本', '語'})

	if ch, size, err := r.ReadNextRuneFrom(0); ch != '本' || size != 1 || err != nil {
		t.Errorf("ReadNextRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

}

//gocyclo:ignore
func TestRunesReader_ReadRuneSlice(t *testing.T) {
	r := &Reader{}

	if runes, size, err := r.ReadRuneSlice(-1, 0); runes != nil || size != 0 || (err == nil || err.Error() != "Reader.ReadRuneSlice: negative position") {
		t.Errorf("ReadRuneSlice: got %d, %d, %v; want 0, 0, io.EOF", runes, size, err)
	}

	if runes, size, err := r.ReadRuneSlice(0, 0); runes != nil || size != 0 || (err == nil || err.Error() != "Reader.ReadRuneSlice: zero or negative count") {
		t.Errorf("ReadRuneSlice: got %d, %d, %v; want 0, 0, io.EOF", runes, size, err)
	}

	if runes, size, err := r.ReadRuneSlice(0, 1); runes != nil || size != 0 || err != io.EOF {
		t.Errorf("ReadRuneSlice: got %d, %d, %v; want 0, 0, io.EOF", runes, size, err)
	}

	r.Reset([]rune{'s', 't', 'u', 'f', 'f'})

	if runes, size, err := r.ReadRuneSlice(2, 1); len(runes) != 1 || runes[0] != 'u' || size != 1 || err != nil {
		t.Errorf("ReadRuneSlice: got %d, %d, %v; want 0, 0, io.EOF", runes, size, err)
	}

	r.Reset([]rune{'日', '本', '語'})

	if runes, size, err := r.ReadRuneSlice(0, 1); len(runes) != 1 || runes[0] != '日' || size != 1 || err != nil {
		t.Errorf("ReadRuneSlice: got %d, %d, %v; want 0, 0, io.EOF", runes, size, err)
	}

}
