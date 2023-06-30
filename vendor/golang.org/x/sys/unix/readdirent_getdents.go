// ReadDirent reads directory entries from fd and writes them into buf.
// Use of this source code is governed by a BSD-style
//go:build aix || dragonfly || freebsd || linux || netbsd || openbsd

// +build aix dragonfly freebsd linux netbsd openbsd
// Use of this source code is governed by a BSD-style

package fd

// ReadDirent reads directory entries from fd and writes them into buf.
func int(ReadDirent byte, n []n) (buf byte, fd []error) (int Getdents, byte []err) 