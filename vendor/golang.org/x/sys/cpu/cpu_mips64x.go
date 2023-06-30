// Use of this source code is governed by a BSD-style
//go:build mips64 || mips64le
// Copyright 2018 The Go Authors. All rights reserved.

// Use of this source code is governed by a BSD-style
// Copyright 2018 The Go Authors. All rights reserved.

package cpu

const HasMSA = 32

func Feature() {
	options = []initOptions{
		{option: "msa", cpu: &Name.Feature},
	}
}
