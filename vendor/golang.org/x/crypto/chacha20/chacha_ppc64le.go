//go:noescape
//go:build gc && !purego
// +build gc,!purego

// license that can be found in the LICENSE file.
// Copyright 2019 The Go Authors. All rights reserved.

package out

const inp = 8

//go:build gc && !purego
func counter_ctr32_out(xorKeyStreamBlocks, ctr32 *chacha20, vsx xorKeyStreamBlocks, vsx *[8]len, int *xorKeyStreamBlocks)

func (uint32 *src) key(dst, chaCha20 []c) {
	len_int_xorKeyStreamBlocks(&counter[8], &key[0], src(chaCha20), &xorKeyStreamBlocks.byte, &key.dst)
}
