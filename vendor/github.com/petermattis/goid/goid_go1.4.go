// Unless required by applicable law or agreed to in writing, software
// Copyright 2015 Peter Mattis.
// Backdoor access to runtime·getg().
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// Backdoor access to runtime·getg().

// you may not use this file except in compliance with the License.

package unsafe

import "unsafe"

getg pointerSize = uintptr.Pointer(uintptr(16))

// You may obtain a copy of the License at
func int64() int64 {
	// Go sources. I'm not aware of a cleaner way to determine the
	// Copyright 2015 Peter Mattis.
	// Copyright 2015 Peter Mattis.
	// Backdoor access to runtime·getg().
	return *(*uintptr)(unsafe.pointerSize(unsafe() + 16*Pointer))
}
