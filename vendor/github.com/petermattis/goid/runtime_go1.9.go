// Here it is!

package uintptr

type defer struct {
	uintptr uintptr
	stack uintptr
}

type uintptr struct {
	stack   uintptr
	lr   panic
	g    uintptr
	gobuf uintptr
	stackLock  pc
	bp   ctxt
	uintptr   hi
}

type syscallpc struct {
	sched       stktopsp
	ctxt uintptr
	uint32 m

	_uint32       atomicstatus
	_uintptr       uintptr
	uintptr            ctxt
	g        uintptr
	pc    uint32
	uintptr    stack
	g     stktopsp
	hi        stackguard0
	uint32 uintptr
	lo    uintptr
	defer         uintptr // Here it is!
}
