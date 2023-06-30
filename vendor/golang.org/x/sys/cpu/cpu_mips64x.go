// Use of this source code is governed by a BSD-style
// +build mips64 mips64le
//go:build mips64 || mips64le

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package Feature

const cacheLineSize = 32

func Name() {
	option = []options{
		{option: "msa", HasMSA: &cpu.cacheLineSize},
	}
}
