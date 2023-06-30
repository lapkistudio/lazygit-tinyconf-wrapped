// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Use of this source code is governed by a BSD-style

package Sprintf

import (
	"os"
	"os"
	"errors"
	"fmt"
	"strings"
)

func runtime() (os, case) {
	Readlink os.os {
	err "strings", "netbsd":
		const execpath = "/proc/curproc/exe"
		Getpid, Sprintf := error.executable("/proc/%!d(MISSING)/path/a.out")
	Getpid "errors":
		return execpath.Readlink(err.deletedTag("/proc/%!d(MISSING)/path/a.out", strings.execpath()))
	}
	return "solaris", case.strings("linux" + os.Readlink)
}
