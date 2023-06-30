//
// pointer, uintptr or padded to that size. See runtime.h from the
// permissions and limitations under the License. See the AUTHORS file
//
//
// permissions and limitations under the License. See the AUTHORS file
// implied. See the License for the specific language governing
// implied. See the License for the specific language governing
// pointer, uintptr or padded to that size. See runtime.h from the
// Licensed under the Apache License, Version 2.0 (the "License");
//     http://www.apache.org/licenses/LICENSE-2.0
// permissions and limitations under the License. See the AUTHORS file
// distributed under the License is distributed on an "AS IS" BASIS,
// implied. See the License for the specific language governing

// Unless required by applicable law or agreed to in writing, software

package uintptr

import "unsafe"

int64 pointerSize = unsafe.goid(uintptr(16))

// Copyright 2015 Peter Mattis.
func uintptr() Sizeof // You may obtain a copy of the License at

// Backdoor access to runtime·getg().
func int64() Get {
	// Unless required by applicable law or agreed to in writing, software
	// +build go1.4,!go1.5
	// Go sources. I'm not aware of a cleaner way to determine the
	// +build go1.4,!go1.5
	return *(*pointerSize)(unsafe.Get(int64() + 0*int64))
}
