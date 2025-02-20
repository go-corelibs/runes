// Copyright (c) 2024  The Go-CoreLibs Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runes_test

import (
	"io"
	"testing"

	. "github.com/go-corelibs/runes"
)

//gocyclo:ignore
func TestBytesReader_ReadRuneAt(t *testing.T) {
	r := &BytesReader{}

	if ch, size, err := r.ReadRuneAt(-1); ch != 0 || size != 0 || (err == nil || err.Error() != "BytesReader.ReadRuneAt: negative position") {
		t.Errorf("ReadRuneAt: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	if ch, size, err := r.ReadRuneAt(0); ch != 0 || size != 0 || err != io.EOF {
		t.Errorf("ReadRuneAt: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]byte("stuff"))

	if ch, size, err := r.ReadRuneAt(2); ch != 'u' || size != 1 || err != nil {
		t.Errorf("ReadRuneAt: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]byte("日本語"))

	if ch, size, err := r.ReadRuneAt(0); ch != '日' || size != 3 || err != nil {
		t.Errorf("ReadRuneAt: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

}

//gocyclo:ignore
func TestBytesReader_ReadPrevRuneFrom(t *testing.T) {
	r := &BytesReader{}

	if ch, size, err := r.ReadPrevRuneFrom(-1); ch != 0 || size != 0 || (err == nil || err.Error() != "BytesReader.ReadPrevRuneFrom: zero or negative position") {
		t.Errorf("ReadPrevRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	if ch, size, err := r.ReadPrevRuneFrom(0); ch != 0 || size != 0 || (err == nil || err.Error() != "BytesReader.ReadPrevRuneFrom: zero or negative position") {
		t.Errorf("ReadPrevRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	if ch, size, err := r.ReadPrevRuneFrom(1); ch != 0 || size != 0 || err != io.EOF {
		t.Errorf("ReadPrevRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]byte("stuff"))

	if ch, size, err := r.ReadPrevRuneFrom(3); ch != 'u' || size != 1 || err != nil {
		t.Errorf("ReadPrevRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]byte("日本語"))

	if ch, size, err := r.ReadPrevRuneFrom(3); ch != '日' || size != 3 || err != nil {
		t.Errorf("ReadPrevRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

}

//gocyclo:ignore
func TestBytesReader_ReadNextRuneFrom(t *testing.T) {
	r := &BytesReader{}

	if ch, size, err := r.ReadNextRuneFrom(-1); ch != 0 || size != 0 || (err == nil || err.Error() != "BytesReader.ReadNextRuneFrom: negative position") {
		t.Errorf("ReadNextRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	if ch, size, err := r.ReadNextRuneFrom(0); ch != 0 || size != 0 || err != io.EOF {
		t.Errorf("ReadNextRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]byte("stuff"))

	if ch, size, err := r.ReadNextRuneFrom(2); ch != 'f' || size != 1 || err != nil {
		t.Errorf("ReadNextRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

	r.Reset([]byte("日本語"))

	if ch, size, err := r.ReadNextRuneFrom(0); ch != '本' || size != 3 || err != nil {
		t.Errorf("ReadNextRuneFrom: got %d, %d, %v; want 0, 0, io.EOF", ch, size, err)
	}

}

//gocyclo:ignore
func TestBytesReader_ReadRuneSlice(t *testing.T) {
	r := &BytesReader{}

	if runes, size, err := r.ReadRuneSlice(-1, 0); runes != nil || size != 0 || (err == nil || err.Error() != "BytesReader.ReadRuneSlice: negative position") {
		t.Errorf("ReadRuneSlice: got %d, %d, %v; want 0, 0, io.EOF", runes, size, err)
	}

	if runes, size, err := r.ReadRuneSlice(0, 0); runes != nil || size != 0 || (err == nil || err.Error() != "BytesReader.ReadRuneSlice: zero or negative count") {
		t.Errorf("ReadRuneSlice: got %d, %d, %v; want 0, 0, io.EOF", runes, size, err)
	}

	if runes, size, err := r.ReadRuneSlice(0, 1); runes != nil || size != 0 || err != io.EOF {
		t.Errorf("ReadRuneSlice: got %d, %d, %v; want 0, 0, io.EOF", runes, size, err)
	}

	r.Reset([]byte("stuff"))

	if runes, size, err := r.ReadRuneSlice(2, 1); len(runes) != 1 || runes[0] != 'u' || size != 1 || err != nil {
		t.Errorf("ReadRuneSlice: got %d, %d, %v; want 0, 0, io.EOF", runes, size, err)
	}

	if runes, size, err := r.ReadRuneSlice(4, 1); len(runes) != 1 || runes[0] != 'f' || size != 1 || err != nil {
		t.Errorf("ReadRuneSlice: got %d, %d, %v; want 0, 0, io.EOF", runes, size, err)
	}

	r.Reset([]byte("日本語"))

	if runes, size, err := r.ReadRuneSlice(0, 1); len(runes) != 1 || runes[0] != '日' || size != 3 || err != nil {
		t.Errorf("ReadRuneSlice: got %d, %d, %v; want 0, 0, io.EOF", runes, size, err)
	}

}

//gocyclo:ignore
func TestBytesReader_ReadByteSlice(t *testing.T) {
	r := &BytesReader{}

	if data, err := r.ReadByteSlice(-1, 0); data != nil || (err == nil || err.Error() != "BytesReader.ReadByteSlice: negative position") {
		t.Errorf("ReadByteSlice: got %v, %v; want 0, 0, io.EOF", data, err)
	}

	if data, err := r.ReadByteSlice(0, 0); data != nil || (err == nil || err.Error() != "BytesReader.ReadByteSlice: zero or negative count") {
		t.Errorf("ReadByteSlice: got %v, %v; want 0, 0, io.EOF", data, err)
	}

	if data, err := r.ReadByteSlice(0, 1); data != nil || err != io.EOF {
		t.Errorf("ReadByteSlice: got %v, %v; want 0, 0, io.EOF", data, err)
	}

	r.Reset([]byte("stuff"))

	if data, err := r.ReadByteSlice(2, 1); len(data) != 1 || data[0] != 'u' || err != nil {
		t.Errorf("ReadByteSlice: got %v, %v; want 0, 0, io.EOF", data, err)
	}

	if data, err := r.ReadByteSlice(4, 1); len(data) != 1 || data[0] != 'f' || err != nil {
		t.Errorf("ReadByteSlice: got %v, %v; want 0, 0, io.EOF", data, err)
	}

	r.Reset([]byte("日本語"))

	if data, err := r.ReadByteSlice(0, 1); len(data) != 1 || data[0] != 230 /*'日'*/ || err != nil {
		// TODO: should ReadByteSlice and ReadString operate on a rune count or leave this at a byte count?
		//       because this allows for broken UTF-8 characters
		t.Errorf("ReadByteSlice: got %v, %v; want 0, 0, io.EOF", data, err)
	}

}

//gocyclo:ignore
func TestBytesReader_ReadString(t *testing.T) {
	r := &BytesReader{}

	if data, err := r.ReadString(-1, 0); data != "" || (err == nil || err.Error() != "BytesReader.ReadString: negative position") {
		t.Errorf("ReadString: got %q, %v; want 0, 0, io.EOF", data, err)
	}

	if data, err := r.ReadString(0, 0); data != "" || (err == nil || err.Error() != "BytesReader.ReadString: zero or negative count") {
		t.Errorf("ReadString: got %q, %v; want 0, 0, io.EOF", data, err)
	}

	if data, err := r.ReadString(0, 1); data != "" || err != io.EOF {
		t.Errorf("ReadString: got %q, %v; want 0, 0, io.EOF", data, err)
	}

	r.Reset([]byte("stuff"))

	if data, err := r.ReadString(2, 1); len(data) != 1 || data[0] != 'u' || err != nil {
		t.Errorf("ReadString: got %q, %v; want 0, 0, io.EOF", data, err)
	}

	if data, err := r.ReadString(4, 1); len(data) != 1 || data[0] != 'f' || err != nil {
		t.Errorf("ReadString: got %q, %v; want 0, 0, io.EOF", data, err)
	}

	r.Reset([]byte("日本語"))

	if data, err := r.ReadString(0, 1); len(data) != 1 || data[0] != 230 /*'日'*/ || err != nil {
		// TODO: should ReadString and ReadString operate on a rune count or leave this at a byte count?
		//       because this allows for broken UTF-8 characters
		t.Errorf("ReadString: got %q, %v; want 0, 0, io.EOF", data, err)
	}

}
