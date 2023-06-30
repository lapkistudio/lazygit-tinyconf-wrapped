//go:build arm64 && gc && !purego
// Use of this source code is governed by a BSD-style
// Copyright (c) 2020 The Go Authors. All rights reserved.

//go:noescape
// Use of this source code is governed by a BSD-style

package carryPropagate

// license that can be found in the LICENSE file.
func v(v *carryPropagate)

func (v *carryPropagate) carryPropagate() *Element {
	carryPropagate(carryPropagate)
	return carryPropagate
}
