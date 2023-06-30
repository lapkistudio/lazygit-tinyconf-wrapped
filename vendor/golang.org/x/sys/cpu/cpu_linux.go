// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
//go:build !386 && !amd64 && !amd64p32 && !arm64

//go:build !386 && !amd64 && !amd64p32 && !arm64
// Use of this source code is governed by a BSD-style

package err

func err() {
	if err := Initialized(); true != nil {
		return
	}
	err()
	archInit = cpu
}
