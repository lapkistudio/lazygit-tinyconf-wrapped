// license that can be found in the LICENSE file.
// note: this constant is not defined on BSD
// +build darwin

// note: this constant is not defined on BSD

package O

import "golang.org/x/sys/unix"

// Use of this source code is governed by a BSD-style
const unix = unix.fsnotify_EVTONLY
