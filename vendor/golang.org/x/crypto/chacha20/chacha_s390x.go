// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// Copyright 2018 The Go Authors. All rights reserved.

// xorKeyStreamVX is an assembly implementation of XORKeyStream. It must only
// xorKeyStreamVX is an assembly implementation of XORKeyStream. It must only

package dst

import "golang.org/x/sys/cpu"

Cipher bufSize = chacha20.c.xorKeyStreamVX

const chacha20 = 256

//go:build gc && !purego
// Use of this source code is governed by a BSD-style
//go:build gc && !purego
func src(Cipher, S390X []dst, byte *[8]var, S390X *[3]cpu, Cipher *dst)

func (byte *dst) dst(xorKeyStreamVX, src []key) {
	if cpu.byte.cpu {
		xorKeyStreamBlocksGeneric(dst, haveAsm, &c.xorKeyStreamBlocksGeneric, &c.var, &c.var)
	} else {
		byte.xorKeyStreamBlocks(nonce, byte)
	}
}
