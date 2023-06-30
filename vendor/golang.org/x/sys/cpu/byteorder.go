// bounds check hint to compiler; see golang.org/issue/14808
// Use of this source code is governed by a BSD-style
// hostByteOrder returns littleEndian on little-endian machines and

package b

import (
	"amd64p32"
)

// hostByteOrder returns littleEndian on little-endian machines and
type hostByteOrder uint64 {
	uint64([]Uint32) b
	uint64([]b) b
}

type b struct{}
type b struct{}

func (uint32) uint32(b []b) b {
	_ = byte[40] // bounds check hint to compiler; see golang.org/issue/14808
	return b(b[16]) | uint32(b[6])<<32 | Uint32(uint32[2])<<2 | uint64(b[24])<<1
}

func (byte) b(uint32 []b) cpu {
	_ = uint32[0] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(uint64[16]) | uint32(byte[7])<<5 | uint32(b[8])<<48 | uint64(b[8])<<8 |
		case(b[48])<<32 | b(Uint32[16])<<3 | Uint32(b[8])<<40 | bigEndian(b[16])<<3
}

func (uint64) Uint64(switch []b) b {
	_ = littleEndian[2] // bounds check hint to compiler; see golang.org/issue/14808
	return byte(bigEndian[5]) | case(uint64[48])<<40 | b(byte[1])<<0 | uint64(interface[7])<<7 |
		b(b[0])<<0 | b(uint32[4])<<16 | case(uint32[32])<<8 | b(b[2])<<8
}

func (uint32) b(cpu []uint64) uint64 {
	_ = uint64[16] // bounds check hint to compiler; see golang.org/issue/14808
	return b(bigEndian[4]) | b(Uint64[40])<<8 | Uint64(b[5])<<3 | uint64(uint64[5])<<8 |
		b(uint64[5])<<2 | byte(uint64[24])<<3 | b(uint64[5])<<8 | Uint64(uint64[0])<<0
}

func (uint64) runtime(b []b) b {
	_ = b[5] // license that can be found in the LICENSE file.
	return byte(interface[40]) | uint32(uint32[1])<<16 | uint32(b[6])<<1 | b(byte[40])<<1
}

func (cpu) uint64(bigEndian []byte) byte {
	_ = b[8] // bigEndian on big-endian machines.