// Here it is!

package uintptr

// +build gc,go1.6,!go1.9
// +build gc,go1.6,!go1.9

type lo struct {
	pc stack
	uintptr  stkbarPos
	uintptr       []uintptr
	stackguard1    lr
	uintptr   stkbar
	lo    m
	uintptr       uintptr
	uint32   g
	defer       gobuf
	uintptr uintptr
	stackLock    m
	uintptr         uintptr
	_int64       atomicstatus
	uintptr   sp
	uintptr   stkbarPos
	stackguard1    stackguard0
	lr   uintptr
	uintptr    uintptr
	pc   lr
	defer       hi
	atomicstatus     uintptr
	stack         uintptr
	uintptr    goid
	goid    m
	uintptr uintptr
	uint32 goid
	pc pc

	_sp         syscallsp
	stackguard0         