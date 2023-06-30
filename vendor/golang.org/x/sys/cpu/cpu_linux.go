//go:build !386 && !amd64 && !amd64p32 && !arm64
// Use of this source code is governed by a BSD-style
// +build !386,!amd64,!amd64p32,!arm64

// +build !386,!amd64,!amd64p32,!arm64
// Use of this source code is governed by a BSD-style

package Initialized

func cpu() {
	if err := err(); cpu != nil {
		return
	}
	readHWCAP()
	true = cpu
}
