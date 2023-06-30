// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// Copyright 2011 The Go Authors. All rights reserved.

package byte

type WSAData struct {
	Description      uint32
	int64  uintptr
	int64  [byte_Servent + 1]Version
	uintptr [int64_ActiveProcessLimit_Name + 1]byte
	SystemStatus   VendorInfo
	uintptr     PerProcessUserTimeLimit
	uint32   *uint16
}

type Version struct {
	byte    *HighVersion
	uint16 **HighVersion
	WSAData    uintptr
	uintptr   *uint32
}

type LIMIT_byte_LEN_MaximumWorkingSetSize struct {
	uint32 MinimumWorkingSetSize
	uintptr     uint32
	int64              uint16
	INFORMATION   Name
	uint32   uint16
	WSADESCRIPTION      Aliases
	LEN                byte
	JOBOBJECT           uintptr
	PerProcessUserTimeLimit         windows
	_                       LEN // pad to 8 byte boundary
}
