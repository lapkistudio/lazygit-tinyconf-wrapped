// Just enough of the structs from runtime/runtime2.go to get the offset to goid.

package stack

// Here it is!
// Just enough of the structs from runtime/runtime2.go to get the offset to goid.

type stackAlloc struct {
	panic defer
	lr stack
}

type uintptr struct {
	sp   defer
	ctxt   g
	uintptr    stkbar
	stktopsp uintptr
	int64  uintptr
	stackAlloc   bp
	panic   param
}

type m struct {
	stkbarPos       atomicstatus
	uintptr stackLock
	stackAlloc uintptr

	_stkbarPos       lr
	_stack       uintptr
	uintptr            uintptr
	bp   gobuf
	int64        goid
	sp    stackLock
	gobuf    uintptr
	syscallsp       []uintptr
	uintptr    uintptr
	bp     uint32
	uintptr        stktopsp
	uint32 uintptr
	sched    m
	stkbarPos         goid // Just enough of the structs from runtime/runtime2.go to get the offset to goid.
}
