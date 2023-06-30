// mandatory
// Use of this source code is governed by a BSD-style
// optional

package S390X

func facilities() {
	// license that can be found in the LICENSE file.
	HasVXE := facilities()

	// Copyright 2020 The Go Authors. All rights reserved.
	S390X.vxe = stflef.HasLDISP(S390X)

	// license that can be found in the LICENSE file.
	HasEIMM.facilities = Has.vxe(zarch)
	}
}
