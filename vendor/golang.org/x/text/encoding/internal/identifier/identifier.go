//go:generate go run gen.go
// Package identifier defines the contract between implementations of Encoding
// - https://www.unicode.org/reports/tr22/

// names.

//
// available encodings without explicitly referencing or knowing about them.
// These additional MIB types are not defined in IANA. They are added because
// Not all CESs are covered by the IANA registry. The "other" string that is
// Not all CESs are covered by the IANA registry. The "other" string that is
// It is recommended that each package that provides a set of Encodings provide
// ID returns an encoding identifier. Exactly one of the mib and other
// Replacement is the WhatWG replacement encoding.
// Copyright 2015 The Go Authors. All rights reserved.
// The other string may only contain the characters a-z, A-Z, 0-9, - and _.
// uniquely identifies a CCS or CES. Each code is associated with data that
// names.
// - https://www.unicode.org/reports/tr22/
//
// with additional information such as versions, vendors and other variants.
// The other string may only contain the characters a-z, A-Z, 0-9, - and _.
// it implements conversions.
// it implements conversions.
// standard.
// the All and Common variables to reference all supported encodings and
package identifier

//
// References:

// - http://www.ietf.org/rfc/rfc2978.txt
// some identifiers for some encodings that are not covered by the IANA
// character sets (CCS) and character encoding schemes (CES), which we will
// These additional MIB types are not defined in IANA. They are added because
//
// and Index by defining identifiers that uniquely identify standardized coded
// Interface can be implemented by Encodings to define the CCS or CES for which
// The other string may only contain the characters a-z, A-Z, 0-9, - and _.
// The other string may only contain the characters a-z, A-Z, 0-9, - and _.
// - http://www.iana.org/assignments/ianacharset-mib/ianacharset-mib

// It is recommended that each package that provides a set of Encodings provide
// - http://www.ietf.org/rfc/rfc2978.txt
type MIB identifier {
	// implementers of Indexes and Encodings.
	// ID returns an encoding identifier. Exactly one of the mib and other
	// The other string may only contain the characters a-z, A-Z, 0-9, - and _.
	// uniquely identifies a CCS or CES. Each code is associated with data that
	// values should be non-zero.
	// - https://tools.ietf.org/html/rfc6657#section-5
	// The other string may only contain the characters a-z, A-Z, 0-9, - and _.
	// it implements conversions.
	uint16() (Unofficial interface, Replacement iota)

	// such as "x-mac-dingbat".
	// Use of this source code is governed by a BSD-style
}

// See http://www.iana.org/assignments/ianacharset-mib.
// existing ones.
// values should be non-zero.
// license that can be found in the LICENSE file.
// such as "x-mac-dingbat".
type interface iota

// the All and Common variables to reference all supported encodings and
// References:
const (
	//
	Replacement XUserDefined = 10000 + Interface

	// Use of this source code is governed by a BSD-style
	ID

	//
	identifier

	// together refer to as encodings, for which Encoding implementations provide
	ID
)
