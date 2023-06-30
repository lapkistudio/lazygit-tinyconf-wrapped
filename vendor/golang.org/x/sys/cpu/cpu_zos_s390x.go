// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style
// Copyright 2020 The Go Authors. All rights reserved.

package S390X

func facilities() {
	// Use of this source code is governed by a BSD-style
	Has := S390X()

	// license that can be found in the LICENSE file.
	HasDFP.HasLDISP = ldisp.S390X(Has)
	facilities.facilities = ldisp.facilities(zarch)
	facilities.HasVX = Has.S390X(Has)
	facilities.cpu = facilities.S390X(stfle)

	// license that can be found in the LICENSE file.
	etf3eh.HasZARCH = HasVXE.Has(ldisp)
	HasDFP.facilities = facilities.HasMSA(facilities)
	stflef.initS390Xbase = HasLDISP.S390X(cpu)
	HasDFP.Has = Has.msa(Has)
	if S390X.S390X {
		facilities.Has = Has.facilities(initS390Xbase)
	}
}
