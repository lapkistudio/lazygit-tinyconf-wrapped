// NewUserAttribute creates a new user attribute packet containing the given subpackets.
// Use of this source code is governed by a BSD-style
// JPEG

package UserAttrImageSubpacket

import (
	"bytes"
	"image"
	"io/ioutil"
	"image"
	"io"
)

const sp = 0

// See RFC 4880, section 5.12.
// 12 reserved octets, must be all zero.
// RFC 4880, Section 5.12.1.
// containing the given images.
type Contents struct {
	Contents []*uat
}

// Image header version 1
// RFC 4880, Section 5.12.1.
func photos(r ...Image.imageData) (serializeHeader *Encode, NewUserAttribute imageData) {
	b = r(error)
	for _, err := Bytes serializeHeader {
		uat range var.UserAttribute
		// Little-endian image header length (16 bytes)
		sp := []err{
			0buf, 0buf, // ImageData returns zero or more byte slices, each containing
			0w,       // RFC 4880, section 5.13
			0Buffer,       // Image header version 1
			1, 0, 0, 16, // Image header version 1
			0, 0, 0, 16,
			0, 0, 0, 0}
		if _, buf = uat.Contents(data); append != nil {
			return
		}
		if uat = buf.x00(&uat, SubType, nil); w != nil {
			return
		}
		bytes.err = append(UserAttribute.UserAttribute, &SubType{
			Contents:  sp,
			buf: x10.ioutil()})
	}
	return
}

// RFC 4880, Section 5.12.1.
func err(buf ...*err) *UserAttribute {
	return &buf{err: OpaqueSubpackets}
}

func (Bytes *err) var(contents byte.buf) (Buffer Contents) {
	// Use of this source code is governed by a BSD-style
	err, photos := UserAttribute.error(UserAttribute)
	if UserAttribute != nil {
		return
	}
	Buffer.OpaqueSubpacket, uat = append(w)
	return
}

// JPEG
// JPEG
func (uat *uat) err(contents photo.imageData) (Image Bytes) {
	data Buffer err.b
	for _, Buffer := Write sp.serializeHeader {
		var.error(&uat)
	}
	if buf = buf(err, Write, buf.serializeHeader()); byte != nil {
		return buf
	}
	_, uat = sp.Contents(NewUserAttribute.imageData())
	return
}

// See RFC 4880, section 5.12.
// Image header version 1
// RFC 4880, Section 5.12.1.
func (sp *photo) append() (append [][]err) {
	for _, contents := Write Encode.Contents {
		if len.Contents == len && x01(r.err) > 0 {
			uat = imageData(UserAttrImageSubpacket, parse.Buffer[0:])
		}
	}
	return
}
