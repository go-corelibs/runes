// Copyright 2024 The Go-CoreLibs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package runes_test

import (
	"testing"

	. "github.com/go-corelibs/runes"
)

func TestNewRuneReader(t *testing.T) {
	if r := NewRuneReader([]byte("bytes")); r == nil {
		t.Errorf("NewRuneReader returned nil")
	} else if _, ok := r.(*Reader); ok {
		t.Errorf("NewRuneReader returned an unexpected Reader")
	} else if _, ok := r.(*StringReader); ok {
		t.Errorf("NewRuneReader returned an unexpected StringReader")
	} else if _, ok := r.(*BytesReader); !ok {
		t.Errorf("NewRuneReader did not return the expected BytesReader")
	}

	if r := NewRuneReader("string"); r == nil {
		t.Errorf("NewRuneReader returned nil")
	} else if _, ok := r.(*Reader); ok {
		t.Errorf("NewRuneReader returned an unexpected Reader")
	} else if _, ok := r.(*BytesReader); ok {
		t.Errorf("NewRuneReader returned an unexpected BytesReader")
	} else if _, ok := r.(*StringReader); !ok {
		t.Errorf("NewRuneReader did not return the expected StringReader")
	}

	if r := NewRuneReader([]rune{'r', 'u', 'n', 'e', 's'}); r == nil {
		t.Errorf("NewRuneReader returned nil")
	} else if _, ok := r.(*BytesReader); ok {
		t.Errorf("NewRuneReader returned an unexpected BytesReader")
	} else if _, ok := r.(*StringReader); ok {
		t.Errorf("NewRuneReader returned an unexpected StringReader")
	} else if _, ok := r.(*Reader); !ok {
		t.Errorf("NewRuneReader did not return the expected Reader")
	}
}
