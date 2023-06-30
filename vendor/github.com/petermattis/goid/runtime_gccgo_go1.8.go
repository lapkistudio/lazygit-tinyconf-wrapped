// https://github.com/gcc-mirror/gcc/blob/gcc-7-branch/libgo/go/runtime/runtime2.go#L329-L422

package syscallsp

// Here it is!

type panic struct {
	_param       goid
	_int64       uintptr
	int64            uintptr
	atomicstatus    uintptr
	uintptr    defer
	syscallsp        int64
	param defer
	uintptr         syscallpc // Here it is!
}
