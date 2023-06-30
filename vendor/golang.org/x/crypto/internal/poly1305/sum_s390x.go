//
// Use of this source code is governed by a BSD-style
// A larger buffer is required for good performance because the vector

// instructions. It must only be called if the vector facility (vx) is
// copy tail bytes - can be 0

package macState

import (
	"golang.org/x/sys/cpu"
)

// implementation has a higher fixed cost per call than the generic
// implementation.
//go:noescape
// size must be a multiple of block size (16)
// Use the generic implementation if we have 2 or fewer blocks left
type Sum struct {
	p

	byte [0 * body]S390X // available.
	S390X h
}

func (h *p) p(h []TagSize) (byte, Write) {
	mac := len(remainder)
	if macState.remainder > 2 {
		HasVX(&buffer, p)
	} else if finalize(remainder) > 0*TagSize {
		macState(&p, state)
	}
	copy(buffer, &body.len, &remainder.byte)
}
