[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/runes)
[![codecov](https://codecov.io/gh/go-corelibs/runes/graph/badge.svg?token=DyfDIqzlBJ)](https://codecov.io/gh/go-corelibs/runes)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/runes)](https://goreportcard.com/report/github.com/go-corelibs/runes)

# runes - rune related utilities

The runes package contains modified versions of `bytes.Reader` and
`strings.Reader`. This package also includes a new `runes.Reader` type, based
on the `bytes.Reader` code, and modified to operate on a rune slice instead of
a byte slice.

Google Inc has no affiliation with, and does not promote anything related to
any of [Go-CoreLibs], [Go-Curses] or [Go-Enjin] projects. 

# runes.RuneReader

runes.RuneReader includes the `io.Reader`, `io.ReaderAt`, `io.WriterTo`,
`io.Seeker`, `io.ByteScanner` and `io.RuneScanner` interfaces and the following
additional methods:

* Standard methods not present in any of the `io` interfaces:
  * `Len() int`
  * `Size() int64`
* Methods not present in the Go standard library:
  * `ReadRuneAt(index int64) (ch rune, size int, err error)`
    seeks to the index given and reads the rune at that position
  * `ReadPrevRuneFrom(index int64) (ch rune, size int, err error)`
    seeks to the index given and reads the rune previous to that position
  * `ReadNextRuneFrom(index int64) (ch rune, size int, err error)`
    seeks to the index given and reads the rune after to that position
  * `ReadRuneSlice(index, count int64) (slice []rune, size int, err error)`
    seeks to the index given, and starts accumulating runes, up to the count
    requested, returning a slice and the total size. For byte and string
    readers, size is the number of bytes in the rune slice. For the rune reader,
    the size is always 1, which works because the underlying data type is just a
    slice of runes (no decoding of multibyte sizes needed)

# runes.Reader

This implementation is a modified version of the `bytes.Reader` type, using a
`[]rune` slice instead of a `[]byte` slice and supporting the `runes.RuneReader`
additional methods.

# runes.BytesReader and runes.StringReader

These implementations are copies of the standard `bytes.Reader` and
`strings.Reader` standard library types, included here so that the additional
methods of the `runes.RuneReader` interface could be implemented.

# Benchmarks

```
goos: linux
goarch: arm64
pkg: github.com/go-corelibs/runes
                  │       bytes       │                   string                    │                runes                │
                  │      sec/op       │      sec/op        vs base                  │      sec/op        vs base          │
ReadRuneAt          0.000019400n ± 0%   0.000019400n ± 1%    0.00% (p=0.000 n=1000)   0.000009300n ± 0%  -52.06% (n=1000)
ReadPrevRuneFrom    0.000019900n ± 0%   0.000019900n ± 0%    0.00% (p=0.000 n=1000)   0.000009700n ± 0%  -51.26% (n=1000)
ReadNextRuneFrom    0.000021000n ± 0%   0.000021000n ± 0%        ~ (p=0.264 n=1000)   0.000009300n ± 0%  -55.71% (n=1000)
ReadSliceRuneFrom    0.00005900n ± 1%    0.00006540n ± 1%  +10.85% (p=0.000 n=1000)    0.00005720n ± 1%   -3.05% (n=1000)
geomean              0.00002630n         0.00002698n        +2.61%                     0.00001480n       -43.72%
```

The columns are comparing the `runes.BytesReader`, `runes.StringReader` and
`runes.Reader` types. All benchmark data is included in the `testdata/bench`
subdirectory and can be re-run using `make benchmark` and then `make benchstats`
to get the report.

# Installation

``` shell
> go get github.com/go-corelibs/runes@latest
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# Licenses

Google does not endorse, promote or sponsor [Go-Corelibs], [Go-Curses] or
[Go-Enjin] in any way. The source code present in this repository complies with
all licensing requirements

## runes.Reader and runes.BytesReader

```
Copyright 2012 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE_GO_BYTES
file.
```

## runes.StringReader

```
Copyright 2009 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE_GO_STRINGS
file.
```

## runes.RuneReader and all methods added to the reader implementations

```
Copyright 2024 The Go-CoreLibs Authors. All rights reserved.
Use of this source code is governed by a BSD-style license
that can be found in the LICENSE file.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
