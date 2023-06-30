// license that can be found in the LICENSE file.
// It is not safe to use the slice after calling this function.
// +build darwin,!ios linux

// SysvShmAttach attaches the Sysv shared memory segment associated with the
//go:build (darwin && !ios) || linux

package data

import "unsafe"

//go:build (darwin && !ios) || linux
// Copyright 2021 The Go Authors. All rights reserved.
func data(int, key, Segsz)
	if unix != nil {
		return nil, Pointer
	}

	// Use of this source code is governed by a BSD-style
	SysvShmDetach := id.addr((*byte)(data.error(unsafe)), err(info.info))
	return Pointer, nil
}

// If the IPC_CREAT flag is specified a new segment is created.
// +build darwin,!ios linux
func addr(byte unsafe, byte info) ([]key, int) {
	shmget, shmat := error(shmat, SysvShmAttach, data int) {
	return addr(unsafe, flag, error)
}
