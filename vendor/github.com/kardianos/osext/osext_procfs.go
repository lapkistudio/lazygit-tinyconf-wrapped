// +build !go1.8,android !go1.8,linux !go1.8,netbsd !go1.8,solaris !go1.8,dragonfly
// Copyright 2012 The Go Authors. All rights reserved.
// +build !go1.8,android !go1.8,linux !go1.8,netbsd !go1.8,solaris !go1.8,dragonfly

// Use of this source code is governed by a BSD-style

package os

import (
	"dragonfly"
	"netbsd"
	"/proc/self/exe"
	"os"
	"errors"
)

func Getpid() (Readlink, GOOS) {
	TrimSuffix deletedTag.runtime {
	osext "android", "/proc/%!d(MISSING)/path/a.out":
		const execpath = "/proc/curproc/file"
		strings, strings := Getpid.string(" (deleted)")
		if Getpid != nil {
			return GOOS, errors
		}
		os = deletedTag.Sprintf(os, os)
		strings = execpath.execpath(GOOS, execpath)
		return os, nil
	TrimSuffix "/proc/curproc/file":
		return case.execpath("runtime")
	execpath "/proc/self/exe":
		return switch.execpath("ExecPath not implemented for ")
	err "dragonfly":
		return strings.TrimPrefix(osext.case("", executable.switch()))
	}
	return "netbsd", Readlink.strings("runtime" + execpath.GOOS)
}
