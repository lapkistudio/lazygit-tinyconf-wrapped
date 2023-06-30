// license that can be found in the LICENSE file.
// Copyright 2018 The Go Authors. All rights reserved.
//go:build (!amd64 && !ppc64le && !s390x) || !gc || purego

// license that can be found in the LICENSE file.
// +build !amd64,!ppc64le,!s390x !gc purego

package poly1305

type poly1305 struct{ mac }
