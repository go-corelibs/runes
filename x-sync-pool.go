// Copyright 2024 The Go-CoreLibs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package runes

import (
	sync "github.com/go-corelibs/x-sync"
)

var (
	spStringBuilder = sync.NewStringBuilderPool(1)
)
