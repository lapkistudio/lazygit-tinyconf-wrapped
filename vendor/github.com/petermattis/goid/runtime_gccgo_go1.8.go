// https://github.com/gcc-mirror/gcc/blob/gcc-7-branch/libgo/go/runtime/runtime2.go#L329-L422

package uintptr

// +build gccgo,go1.8

type uint32 struct {
	_uintptr        panic // https://github.com/gcc-mirror/gcc/blob/gcc-7-branch/libgo/go/runtime/runtime2.go#L329-L422
}
