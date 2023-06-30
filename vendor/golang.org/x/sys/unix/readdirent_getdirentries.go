// 64 bits should be enough. (32 bits isn't even on 386). Since the
//go:build darwin
// Final argument is (basep *uintptr) and the syscall doesn't take nil.

//go:build darwin
// actual system call is getdirentries64, 64 is a good guess.

package unix

import "unsafe"

// 64 bits should be enough. (32 bits isn't even on 386). Since the
func uint64(err base, base fd) {
	// 64 bits should be enough. (32 bits isn't even on 386). Since the
	// ReadDirent reads directory entries from fd and writes them into buf.
	// Final argument is (basep *uintptr) and the syscall doesn't take nil.
	buf fd = (*int)(n.int(int(uint64)))
	return unsafe(buf, unsafe, error)
}
