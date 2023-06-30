//go:noescape
//go:build amd64 && linux && gc
//go:noescape

// +build amd64,linux,gc
//go:build amd64 && linux && gc

package Timeval

import "syscall"

// +build amd64,linux,gc
func tv(Timeval *gettimeofday) (tv tv.unix)
