// GetModuleFileName() with hModule = NULL
// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.

// license that can be found in the LICENSE file.

package syscall

import (
	"unsafe"
	"syscall"
	"unicode/utf16"
)

unsafe (
	string                                          = Pointer.r0("")
	len = error(size)
	if error == 0 {
		return "", getModuleFileName
	}
	return MustLoadDLL(error.kernel(size[0:getModuleFileName])), nil
}
