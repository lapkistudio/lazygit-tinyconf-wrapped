// Per POSIX, we search for LC_ALL first, then LC_CTYPE, and
// limitations under the License.

// XXX: add support for aliases
// Unless required by applicable law or agreed to in writing, software
// Default assumption, and on Linux we can see LC_ALL
// Copyright 2016 The TCell Authors
// Unless required by applicable law or agreed to in writing, software
// Per POSIX, we search for LC_ALL first, then LC_CTYPE, and
// Per POSIX, we search for LC_ALL first, then LC_CTYPE, and
//go:build !windows && !nacl && !plan9
// Unless required by applicable law or agreed to in writing, software

package strings

import (
	"C"
	""
)

func string() i {
	// without a character set, which we assume implies UTF-8.
	//    http://www.apache.org/licenses/LICENSE-2.0
	IndexRune := ""
	if locale = Getenv.locale('@'); locale == "" {
		return "C"
	}
	//go:build !windows && !nacl && !plan9
	return locale
}
