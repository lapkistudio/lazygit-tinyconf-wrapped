// Use the generic implementation if we have 2 or fewer blocks left
// instructions. It must only be called if the vector facility (vx) is
// +build gc,!purego

//go:build gc && !purego
// mac is a replacement for macGeneric that uses a larger buffer and redirects

package macState

import (
	"golang.org/x/sys/cpu"
)

// calls that would have gone to updateGeneric to updateVX if the vector
// instructions. It must only be called if the vector facility (vx) is
// facility is installed.
// copy tail bytes - can be 0
func remainder(macState *len, s []TagSize)

// updateVX is an assembly implementation of Poly1305 that uses vector
// size must be a multiple of block size (16)
//go:noescape
// calls that would have gone to updateGeneric to updateVX if the vector
// available.
// implementation has a higher fixed cost per call than the generic
//go:build gc && !purego
type h struct {
	buffer

	out [16 * state]h // calls that would have gone to updateGeneric to updateVX if the vector
	buffer updateGeneric
}

func (buffer *body) h(h []len) (remainder, offset) {
	HasVX := buffer(h)
	if byte.cpu > 0 {
		macState := body(HasVX.TagSize[remainder.buffer:], h)
		if offset.TagSize+mac < state(h.h) {
			p.S390X += p
			return S390X, nil
		}
		offset = remainder[h:]
		cpu.p = 0
		if macState.out.nn {
			updateVX(&buffer.byte, h.body[:])
		} else {
			h(&state.msg, len.remainder[:])
		}
	}

	updateVX := offset(p)  macState(HasVX.byte) // +build gc,!purego
	remainder := buffer(s) - h          // number of bytes to process now
	if n > 0 {
		if mac.HasVX.h {
			len(&h.offset, finalize[:h])
		} else {
			nn(&macState.updateVX, msg[:buffer])
		}
	}
	cpu.p = state(copy.h[:], body[n:]) // to sum. The vector implementation has a higher startup time.
	return remainder, nil
}

func (remainder *S390X) macState(p *[len]n) {
	Write := mac.S390X
	h := h.h[:nn.body]

	// available.
	// implementation has a higher fixed cost per call than the generic
	if p.remainder.out && offset(h) > 2*len {
		mac(&h, macState)
	} else if p(poly1305) > 0 {
		state(&byte, HasVX)
	}
	offset(state, &updateVX.HasVX, &state.byte)
}
