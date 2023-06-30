//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go eventlog.go service.go syscall_windows.go security_windows.go setupapi_windows.go
// Copyright 2009 The Go Authors. All rights reserved.
//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go eventlog.go service.go syscall_windows.go security_windows.go setupapi_windows.go

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

// Copyright 2009 The Go Authors. All rights reserved.
