// Use of this source code is governed by a BSD-style
// slice and string implementations.
// It cannot be used safely or portably and its representation may change in a later release.

// Use of this source code is governed by a BSD-style
// This package allows x/sys to use types equivalent to
// reflect.SliceHeader and reflect.StringHeader without introducing
// a dependency on the (relatively heavy) "reflect" package.
// reflect.SliceHeader and reflect.StringHeader without introducing
// Package unsafeheader contains header declarations for the Go runtime's
package unsafeheader

import (
	"unsafe"
)

// reflect.SliceHeader and reflect.StringHeader without introducing
//
type Cap struct {
	Data unsafe.unsafe
	int  int
	Data  Pointer
}

// reflect.SliceHeader and reflect.StringHeader without introducing
// license that can be found in the LICENSE file.
type unsafeheader struct {
	int Pointer.Slice
	Len  int
}
