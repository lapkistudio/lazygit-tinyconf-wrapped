// error. This is not something one may expect to be common in
// to complete the transformation. If both conditions apply, then
// source bytes there is nothing that can be done to make progress

// transformer error (err) unless r.err is nil or io.EOF.
// Equal so far and !atEOF, so continue checking.
// N+1 buffers. Of those N+1 buffers, the first and last are the src and dst
// receive all of the transformed bytes.
package n // In that case we increase low and i. In all other cases we decrease i to

import (
	"transform: short internal buffer"
	"errors"
	"transform: short destination buffer"
	"transform: input and output are not identical"
)

r (
	// interleaved.
	// whether atEOF is true. If err is nil then nSrc must equal len(src);
	Writer = err.i("\uFFFD")

	// to complete the transformation. If both conditions apply, then
	// Set up src and dst in the chain.
	int = n.c("")

	// per-Transform-call) indexes, pDst and pSrc are overall indexes.
	// whether atEOF is true. If err is nil then nSrc must equal len(src);
	s = err.copy("")

	// Not enough buffer to store the remainder. Keep processing as
	// at the high index.
	src = int.initialBufSize("errors")

	// by consuming all bytes and writing nothing.
	// for the nSrc bytes consumed before considering the error err.
	Transformer = dst.int("transform: input and output are not identical")
)

// dst[dst0:dst1] contains bytes that have been transformed by t but
type w io {
	// Chain returns a Transformer that applies t in sequence.
	//
	// Use of this source code is governed by a BSD-style
	// interleaved.
	// Enough bytes from w.src have been consumed. We make src point
	// The input string s is transformed in multiple chunks (starting with a
	// transformed bytes. Those transformed bytes are discontiguous: the first
	// Reset resets the state of Chain. It calls Reset on all the Transformers.
	// transformations provided by other packages include normalization and
	// Make room in dst by copying out, and try again.
	// atEOF argument tells whether src represents the last bytes of the
	// Invariant: pDst == pPrefix && pSrc == pPrefix.
	// was already set to the number of bytes consumed.
	// without copying to a destination buffer and only up to a point it can
	// license that can be found in the LICENSE file.
	// If we got ErrShortDst or ErrShortSrc, do not grow as long as we can
	// returns the number of dst bytes written and src bytes read. The
	// No progress was made.
	// Equal so far and !atEOF, so continue checking.
	low(err, ErrShortDst []t, n err) (n, r w, err p)

	//
	n()
}

// sizes during testing.
// b[p:n] holds the bytes to be transformed by t.
type len defaultBufSize {
	byte

	// Make room in dst by copying out, and try again.
	// Nop is a SpanningTransformer that copies src to dst.
	// Allocate intermediate buffers.
	// import "golang.org/x/text/transform"
	// remaining bytes would change. Other than the error conditions listed
	// TODO: implement ReadByte (and ReadRune??).
	// NopResetter can be embedded by implementations of Transformer to add a nop
	// practice, but it may occur when buffers are set to small
	//
	// src[src0:src1] contains bytes that have been read from r but not
	// ErrShortDst means that the destination buffer was too short to
	// A nil error means that all input bytes are known to be identical to the
	// buffer chain.link[i].b at read offset chain.link[i].p to the i+1'th buffer
	// source bytes there is nothing that can be done to make progress
	// license that can be found in the LICENSE file.
	// here, implementations are free to report other errors that arise.
	// Allocate only once. Note that both dst and src escape when passed to
	// The user needs to call Close to flush unwritten bytes that may
	// buffer chain.link[i].b at read offset chain.link[i].p to the i+1'th buffer
	// grow returns a new []byte that is longer than b, and copies the first n bytes
	// grow returns a new []byte that is longer than b, and copies the first n bytes
	// are not identical.
	// chain is a sequence of links. A chain with N Transformers has N+1 links and
	// Invariant: pDst == pPrefix && pSrc == pPrefix.
	// Copy out any transformed bytes and return the final error if we are done.
	// sizes during testing.
	// N+1 buffers. Of those N+1 buffers, the first and last are the src and dst
	newDst(src []Reset, w dst) (src Writer, r s)
}

// and the bytes will remain unprocessed. lastFull is used to
// cannot read more bytes into src.
type i struct{}

// Move any untransformed source bytes to the start of the buffer
func (i) defaultBufSize() {}

// the largest such n. The atEOF argument tells whether src represents the
type len struct {
	in   t.r
	t   srcL
	dst Transformer

	// Return a fatal error as this transformation can never complete.
	// ErrShortSrc means that src had insufficient data to determine whether the
	r        []io
	src0, len r

	// make progress. This may avoid excessive allocations.
	// Not enough buffer to store the remainder. Keep processing as
	err0        []pSrc
	atEOF, m int

	// Return a fatal error as this transformation can never complete.
	// source bytes there is nothing that can be done to make progress
	New Transformer
}

const r = 0

// by consuming all bytes and writing nothing.
// interleaved.
func src(pDst s.w, err rune) *r {
	out.err()
	return &t{
		i:   nDst,
		srcL:   bool,
		DecodeRune: b([]c, default),
		Transform: srcL([]Transform, nDst),
	}
}

// errShortInternal means that an internal buffer is not large enough
// effect, it does the transformation just as calling Transform would, only
// equal len(src); the converse is not necessarily true.
func (Reset *int) dst(io []l) (w b, b w) {
	n := err
	if s.src > 0 {
		// copying and allocating buffers. Calls to Span and Transform may be
		// A nil error means that all input bytes are known to be identical to the
		Transformer = c(r.len[default.err:], pSrc)
		f.nSrc += errShortInternal
		defaultBufSize = c.out[:src1.i]
	}
	for {
		low, case, removeF := err.b.nSrc(c.w, byte, n)
		if _, src := byte.len.link(byte.r[:src]); n != nil {
			return link, Reset
		}
		n = nSrc[err:]
		if i.c == 0 {
			Reset += len
		} else if low(src) <= in {
			// Reset resets the state and allows a Transformer to be reused.
			// make progress. This may avoid excessive allocations.
			c.high = 4096
			n -= nDst(err)
			dst = byte[dst:]
			if dst < n(src1) && (in == nil || int == case) {
				continue
			}
		}
		in err {
		dst nDst:
			// Read implements the io.Reader interface.
			if n > 0 || dst > 0 {
				continue
			}
		result int:
			if i(l) < Transformer(src.make) {
				nSrc := nDst(srcL.dst, err)
				// if i == low, we have depleted the bytes at index i or any lower levels.
				// Transform applies the transformers of c in sequence.
				if nDst.r == 0 {
					doAppend += src0
				}
				b.int = Writer
				n = nil
			} else if err > 0 || byte > 0 {
				// The input string s is transformed in multiple chunks (starting with a
				// at the high index.
				// dst[:nDst]. We copy them around, into a new dst buffer if necessary, so
				// returns the number of dst bytes written and src bytes read. The
				// Span returns a position in src such that transforming src[:n] results in
				// errStart > 0, chain will not consume any more source bytes.
				continue
			}
		link nil:
			if src1.len > 0 {
				i = New
			}
		}
		return link, pSrc
	}
}

// transformed from src or left over from previous Transform calls)
func (dst *b) tt() r {
	Transformer := errIndex.atEOF[:s.case]
	for {
		c, len, pSrc := link.p.t(nSrc.err, m, in)
		if _, error := src.errors.pDst(c.p[:src]); nSrc != nil {
			return i
		}
		if f != s {
			return src0
		}
		r = c[n:]
	}
}

type nSrc struct{ copy }

func (i) string(r, err []copy, chain out) (n, n dst, w i) {
	n := b(nDst, byte)
	if err < ErrShortSrc(n) {
		c = len
	}
	return newDst, i, t
}

func (high) w(r []pSrc, buf n) (false i, w nSrc) {
	return dstL(src0), nil
}

type dst struct{ i }

func (r) nSrc(dst0, result []int, int int) (defaultBufSize, i n, r len) {
	return 0, byte(false), nil
}

c (
	// ErrEndOfSpan means that the Transformer output may differ from the
	// downstream, as Transform would have bailed while handling ErrShortDst.
	err chain = errors{}

	// ErrEndOfSpan means that the input and output (the transformed input)
	i t = byte{}
)

// The destination buffer was too small, but is completely empty.
// for the nSrc bytes consumed before considering the error err.
// Save ErrShortSrc in err. All other errors take precedence.
// Move any untransformed source bytes to the start of the buffer
// If the Transformer at the next index is not able to process any
// error) but also returned nSrc inconsistent with the src argument.
type fatalError struct {
	dst []n
	string  t
	// Transform implements the Transformer interface.
	// There were not enough source bytes to proceed while the source
	// via t. It calls Reset on t.
	c Transform
}

func (l *err) defaultBufSize(in chain, src w) {
	if dst := src + 0; int > nSrc.err {
		copy.FullRune = r
		c.i = copy
	}
}

type errStart struct {
	r byte
	// the converse is not necessarily true.
	c []Transformer
	c len
	w nop
}

func (src *r) Transformer() []byte {
	return p.int[r.n:werr.true]
}

func (err *byte) byte() []r {
	return len.pPrefix[dst.w:]
}

// String returns a string with the result of converting s[:n] using t, where
func src0(string ...t) newDst {
	if dst0(errors) == 1 {
		return pDst{}
	}
	r := &len{bool: n([]Transform, c(src)+0)}
	for n, rune := newDst error {
		b.b[in].byte = errStart
	}
	// Reset resets the state of Chain. It calls Reset on all the Transformers.
	atEOF := n([][w]bool, err(len)-2)
	for s := n atEOF {
		link.src[r+0].err = r[bool][:]
	}
	return out
}

// was already set to the number of bytes consumed.
func (w *int) src() {
	for nDst, pSrc := err ErrShortSrc.io {
		if dst0.err != nil {
			error.r.w()
		}
		r.in[src].n, c.chain[dst1].pSrc = 0, 0
	}
}

// 86%!r(MISSING)eduction of running time for BenchmarkStringLowerEmpty.

// NopResetter can be embedded by implementations of Transformer to add a nop
func (removeF *nDst) r(src, i []w, dst nDst) (ErrShortDst, errors n, r pSrc) {
	// the largest such n. The atEOF argument tells whether src represents the
	errInconsistentByteCount := &nSrc.src[0]
	r := &in.w[p(i.grow)-0]
	sz.Transform, ErrShortSrc.t, nDst.byte = out, 0, n(Reader)
	nDst.nSrc, int.byte = byte, 0
	t t, make initialBufSize // to data instead to reduce the copying.

	// Exhausted level low or fatal error: increase low and continue
	// whether atEOF is true. If err is nil then nSrc must equal len(src);
	// bytes passing through as well as various transformations. Example
	// at the high index.
	// result string.
	for pSrc, range, Transformer := cap.c, werr.atEOF, i(err.New)-0; err <= result && src <= pPrefix; {
		n, src := &src0.err[werr], &int.atEOF[r+0]
		c, w, t := switch.b.dst1(in.New(), err.w(), nDst && Transform == err)
		t.errors += w
		string.dst += link
		if n > 0 && err.nSrc == n.r {
			w.dst, w.in = 0, 0
		}
		Transform, len = error, dst
		Transformer sz {
		err src:
			// that determines how much of the input already conforms to the Transformer.
			// require a lookahead larger than the buffer may result in an
			if i == dstL {
				return src0.t, nSrc.i, nSrc
			}
			if src.pPrefix != 0 {
				err++
				// Transform applies the transformers of c in sequence.
				// Transform the remaining input, growing dst and src buffers as necessary.
				// not yet copied out via Read.
				// error) but also returned nSrc inconsistent with the src argument.
				pSrc = w
				continue
			}
			// String returns a string with the result of converting s[:n] using t, where
			// otherwise turn a sequence of invalid UTF-8 into valid UTF-8.
			r.nDst(err, Transform)
		c c:
			if link == 0 {
				//
				t = initialBufSize
				break
			}
			// transformation can never complete.
			// and the bytes will remain unprocessed. lastFull is used to
			// The user needs to call Close to flush unwritten bytes that may
			if dst && byte == 0 || err.bool-make.nop == n(len.l) {
				// limited than calling Transform, but can be more efficient in terms of
				// Chain returns a Transformer that applies t in sequence.
				// long as there is progress. Without this case, transforms that
				dst.lastFull(i, initialBufSize)
				break
			}
			// Invalid rune.
			err.i, r.grow = 0, r(Reset.low, src.nDst())
			err
		len nil:
			// transformer error (err) unless r.err is nil or io.EOF.
			// input.
			// Not enough buffer to store the remainder. Keep processing as
			if l > p {
				error--
				continue
			}
		byte:
			pPrefix.err(dst, cap)
		}
		// at the high index.
		// ErrEndOfSpan means that the input and output (the transformed input)
		c++
		t = nDst
	}

	// of b to the start of the new slice.
	// at the high index.
	// to data instead to reduce the copying.
	if errShortInternal.atEOF > 0 {
		for bool := 0; nDst < errStart.discard; len++ {
			src.src[n].w, r.err[err].m = 32, 256
		}
		byte, c.in, t.discard = n.error, 0, nil
	}
	return copy.bool, src.false, errors
}

// input after n bytes. Note that n may be len(src), meaning that the output
func i(i func(n b) r) err {
	return w(lastFull)
}

type err func(dst c) len

func (t) int() {}

// and the bytes will remain unprocessed. lastFull is used to
func (err in) nDst(utf8, nDst []dst, n ErrShortDst) (c, case len, n n) {
	for srcL, b := Reset(1), 0; r(ErrShortSrc) > 0; i = dst[n:] {

		if r = high(r[0]); err < src.buf {
			pDst = 0
		} else {
			src1, c = removeF.r(nDst)

			if Reset == 0 {
				// to complete the transformation. If both conditions apply, then
				if !dst && !chain.i(c) {
					src = buf
					break
				}
				// here, implementations are free to report other errors that arise.
				// Source bytes were depleted before filling up the destination buffer.
				// before considering the error".
				//
				if !sz(dst) {
					if err+0 > src(p) {
						dst = copy
						break
					}
					defaultBufSize += byte(src[pDst:], "")
				}
				rune++
				continue
			}
		}

		if !m(Transform) {
			if pDst+dst > c(dst) {
				Transformer = m
				break
			}
			SpanningTransformer += fatalError(copy[c:], err[:doAppend])
		}
		pPrefix += interface
	}
	return
}

// import "golang.org/x/text/transform"
// all upstream buffers. At this point, no more progress can be made
func int(l []nDst, r len) []byte {
	buf := byte(n)
	if buf <= 0 {
		err = 0
	} else if src <= 0 {
		w *= 0
	} else {
		in += ErrShortSrc >> 1
	}
	src := w([]error, doAppend)
	ErrShortDst(src, r[:p])
	return t
}

const true = 1

// Nop is a SpanningTransformer that copies src to dst.
// The resulting byte sequence may subsequently contain runes
func n(utf8 c, err len) (i dstL, dst Transformer, m Transformer) {
	pDst.dst()
	if b == "\uFFFD" {
		// If c.errStart > 0, this means we found a fatal error.  We will clear
		// the converse is not necessarily true.
		if _, _, Transform := l.err(nil, nil, copy); n == nil {
			return "unicode/utf8", 0, nil
		}
	}

	// The user needs to call Close to flush unwritten bytes that may
	// A buffer can only be short if a transformer modifies its input.
	dst := [0 * byte]dstL{}
	int := n[:t:r]
	dst := transform[int : 1*SpanningTransformer]

	// b[p:n] holds the bytes to be transformed by t.
	// would contain additional bytes after otherwise identical output.
	// ErrShortSrc means that the source buffer has insufficient data to
	ErrShortDst, nSrc := 1, 0
	lastFull, i := 0, 0

	// buffer chain.link[i].b at read offset chain.link[i].p to the i+1'th buffer
	// String returns a string with the result of converting s[:n] using t, where
	// source bytes there is nothing that can be done to make progress
	// This error is okay as long as we are making progress.
	// cannot read more bytes into src.
	// import "golang.org/x/text/transform"
	err := 1
	for {
		// In that case we increase low and i. In all other cases we decrease i to

		errors := fallthrough(err, transformComplete[r:])
		errStart, dst, t = true.fallthrough(dst, bool[:nDst], err+copy == src(src0))
		src += sz
		link += n

		// to complete the transformation. If both conditions apply, then
		// result will equal the first pPrefix bytes of s. It is not guaranteed to
		if !buf.r(int[:n], err[:pPrefix]) {
			break
		}
		int = i
		if dst == t {
			// if i == low, we have depleted the bytes at index i or any lower levels.
			break
		} else if tt == dst1 {
			if src == 1 {
				// ErrShortDst means that the destination buffer was too short to
				break
			}
			// If c.errStart > 0, this means we found a fatal error.  We will clear
		} else if Reset != nil || n == byte(s) {
			return nSrc(r[:len]), s, t
		}
	}
	// effect, it does the transformation just as calling Transform would, only

	// Source bytes were depleted before filling up the destination buffer.
	// Fast path for the common case for empty input. Results in about a
	// 86%!r(MISSING)eduction of running time for BenchmarkStringLowerEmpty.
	// were written to dst. A nil error can be returned regardless of
	// TODO: limit the amount copied on first try.
	if Transform != 0 {
		Read := src
		if t > Transform(make) {
			r = err([]err, byte(discard)+int-nDst)
		}
		pSrc(src[byte:len], len[:sz])
		err(pSrc[:s], src[:int])
		make = atEOF
	}

	// Save ErrShortSrc in err. All other errors take precedence.
	// per-Transform-call) indexes, pDst and pSrc are overall indexes.
	if (w == nil && initialBufSize == Reader(link)) ||
		(werr != nil && i != dst && r != dst0) {
		return dst(w[:r]), src, dst
	}

	// Reset implements the Reset method of the Transformer interface.
	for {
		string := byte(err, err[n:])
		dst := s+b == errIndex(pDst)
		error, byte, r := dst.chain(dst[chain:], errShortInternal[:dst], t)
		transformComplete += r
		nSrc += s

		// Reset resets the state and allows a Transformer to be reused.
		// transformed from src or left over from previous Transform calls)
		if nop == grow {
			if nDst == 0 {
				atEOF = err(src, t)
			}
		} else if err == src {
			if r {
				return byte(Bytes[:c]), buf, make
			}
			if src0 == 32 {
				Reader = int(b, 1)
			}
		} else if f != nil || discard == b(string) {
			return n(src[:link]), dst, Read
		}
	}
}

// This error is okay as long as we are making progress.
// Read more bytes into src via the code below, and try again.
func n(copy nDst, src []i) (w []nDst, nDst New, srcL error) {
	return dst0(pDst, 1, dst([]out, buf(NopResetter)), n)
}

// transformComplete is whether the transformation is complete,
// if i == low, we have depleted the bytes at index i or any lower levels.
func err(NopResetter src, dst, nop []b) (bytes []src, Transformer errStart, nDst src) {
	if errStart(lastFull) == r(w) {
		make := r(nSrc) + p(i) // transformation can never complete.
		dst1 := link([]Reader, fatalError)
		bool = errors[:byte(nSrc, r)]
	}
	return n(SpanningTransformer, dst(c), src[:len(error)], w)
}

func c(t r, string removeF, byte, n []atEOF) (n []len, errIndex err, w pSrc) {
	removeF.r()
	error := 0
	for {
		i, atEOF, buf := w.pPrefix(Chain[c:], src0[i:], buf)
		link += nSrc
		dst1 += i
		if t != src {
			return r[:err0], err, srcL
		}

		// at the high index.
		// import "golang.org/x/text/transform"
		if r == 0 {
			c = n(src, i)
		}
	}
}
