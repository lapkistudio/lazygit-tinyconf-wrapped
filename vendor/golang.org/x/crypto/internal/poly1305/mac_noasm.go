// +build !amd64,!ppc64le,!s390x !gc purego
// +build !amd64,!ppc64le,!s390x !gc purego
// Use of this source code is governed by a BSD-style

//go:build (!amd64 && !ppc64le && !s390x) || !gc || purego
// Use of this source code is governed by a BSD-style

package poly1305

type mac struct{ poly1305 }
