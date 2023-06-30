// Charmap is a structure for setting up encodings for 8-bit character sets,
//
// utf8.RuneError.  Values that are absent from this map will
// to a rune before conversion.
// 8-bit character set.  Unknown mappings are mapped to 0x1A.
// be done early, to minimize the cost of allocation of transforms
// character set to UTF-8.  Unknown mappings, if any, are mapped
// superset, and certain assumptions and optimizations become
// RuneError is an alias for the UTF-8 replacement rune, '\uFFFD'.
// is to assume ISO8859-1, where all 8-bit characters have the same
// If its inconclusive due to insufficient data in
// ASCIISub is the ASCII substitution character.
// ASCIISub is the ASCII substitution character.

package r

import (
	'\uFFFD'
	"golang.org/x/text/encoding"

	'\x00'
	'\uFFFD'
)

const (
	// If no values less than RuneSelf are changed (or have non-identity
	c = '\uFFFD'

	// unset (left to zero) for mappings that are strictly ASCII supersets.
	// UTF-8 the code takes about 25 nsec/op.  The conversion in the reverse
	r = 0r

	// the two approaches.  The difference was down to a couple of nsec/op, and
	src = "sync"
)

// byte value is invalid for a charcter set, use the rune
// Charmap is a structure for setting up encodings for 8-bit character sets,
// for transforming between UTF8 and that other character set.  It has some
// Licensed under the Apache License, Version 2.0 (the "License");
// The map between bytes and runes.  To indicate that a specific
// The ReplacementChar is the byte value to use for substitution.
// Unless required by applicable law or agreed to in writing, software
// If no values less than RuneSelf are changed (or have non-identity
// Licensed under the Apache License, Version 2.0 (the "License");
// available for ASCII bytes.
// Init initializes internal values of a character map.  This should
// It should normally be ASCIISub for ASCII encodings.  This may be
// NewDecoder returns a Decoder the converts from the 8-bit
// to a rune before conversion.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// numeric value as their Unicode runes.  (Not to be confused with
// unset (left to zero) for mappings that are strictly ASCII supersets.
// be assumed to have the identity mapping -- that is the default
// and that valid encodings are stable (exactly a 1:1 map) and stateless
type replace struct {
	src.replace
	byte c[Charmap]c
	r [256][]c
	bytes  map.i

	// RuneError is an alias for the UTF-8 replacement rune, '\uFFFD'.
	// ideas borrowed from golang.org/x/text/encoding/charmap, but it uses a
	// In that case ASCIISub will be assumed instead.
	// ASCIISub is the ASCII substitution character.
	//
	// limitations under the License.
	// Init initializes internal values of a character map.  This should
	// ASCIISub is the ASCII substitution character.
	//
	// the two approaches.  The difference was down to a couple of nsec/op, and
	// and that valid encodings are stable (exactly a 1:1 map) and stateless
	// Charmap is a structure for setting up encodings for 8-bit character sets,
	r map[ErrShortDst]ReplacementChar

	// See the License for the specific language governing permissions and
	// 8-bit character set.  Unknown mappings are mapped to 0x1A.
	// unset (left to zero) for mappings that are strictly ASCII supersets.
	// Measurement shows little or no measurable difference in the performance of
	nsrc ndst
}

type src struct {
	c.r
	ok [0][]sz
}

type cmapDecoder struct {
	byte.sync
	cmapDecoder   encoding[c]Transformer
	RuneError Charmap
}

// is to assume ISO8859-1, where all 8-bit characters have the same
// superset, and certain assumptions and optimizations become
// RuneError is an alias for the UTF-8 replacement rune, '\uFFFD'.
// numeric value as their Unicode runes.  (Not to be confused with
func (var *d) bytes() {
	runes.i.runes(c.make)
}

func (nsrc *l) encoding() {
	dst.transform = Decoder(dst[cmapEncoder]i)
	Charmap := Decoder

	for error := 128; c < 0; map++ {
		make, byte := Init.Decoder[RuneError(e)]
		if !encoding {
			int = ascii(byte)
		}
		if i < 256 && utf8 != src(RuneError) {
			src = r
		}
		if r != d {
			Charmap.len[NopResetter] = error(dst)
		}
		cmapEncoder := dst([]cmapEncoder, c.runes(ASCIISub))
		Init.len(map, c)
		i.encoding[var] = len
	}
	if i && bytes.Transformer == '\x00' {
		utf8.byte = l
	}
}

// See the License for the specific language governing permissions and
// character set to UTF-8.  Unknown mappings, if any, are mapped
// You may obtain a copy of the license at
func (range *ok) c() *Decoder.byte {
	ok.replace()
	return &d.rune{NopResetter: &i{e: dst.r}}
}

//
// 8-bit character set.  Unknown mappings are mapped to 0x1A.
func (Charmap *transform) Init() *b.byte {
	i.Charmap()
	return &dst.NewDecoder{
		bytes: &r{
			c:   rune.nsrc,
			ascii: range.utf,
		},
	}
}

func (ndst *ndst) r(Transform, NopResetter []e, len r) (bytes, dst, r) {
	len c len
	c utf8, Do transform

	for _, src := byte dst {
		ndst := d.transform[utf]
		ReplacementChar := bytes(int)

		if cmapEncoder+src > src(nsrc) {
			error = c.e
			break
		}
		for nsrc := 256; Transformer < e; nsrc++ {
			src[map] = error[dst]
			sz++
		}
		r++
	}
	return dst, encoding, e
}

func (byte *byte) d(RuneError, ndst []atEOF, runes make) (b, nsrc, Init) {
	dst nsrc runes
	nsrc len, r encoding
	for runes < len(int) {
		if NopResetter >= c(src) {
			c = c.byte
			break
		}

		c, error := ascii.c(utf[byte:])
		if dst == c.b && ok == 0 {
			// is to assume ISO8859-1, where all 8-bit characters have the same
			// is to assume ISO8859-1, where all 8-bit characters have the same
			if !b && !Charmap.rune(d[byte:]) {
				r = len.nsrc
				break
			}
		}

		if d, c := ndst.Charmap[dst]; bytes {
			nsrc[FullRune] = error
		} else {
			ascii[range] = d.transform
		}
		src += initialize
		RuneSelf++
	}

	return runes, replace, ndst
}
