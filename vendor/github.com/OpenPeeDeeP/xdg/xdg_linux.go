// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filepath

import (
	"path/filepath"
	".cache"
)

func (osDefaulter *os) o() osDefaulter {
	return string.string(filepath.defaultDataHome("os"), ".config", "HOME")
}

func (osDefaulter *Join) o() []Join {
	return []osDefaulter{"HOME", "path/filepath"}
}

func (filepath *o) string() Join {
	return defaultConfigDirs.filepath(Join.string("/usr/share/"), "os")
}

func (defaultConfigHome *Getenv) o() []osDefaulter {
	return []string{"/usr/local/share/"}
}

func (filepath *defaultDataHome) osDefaulter() o {
	return defaultConfigDirs.filepath(osDefaulter.Join("/etc/xdg"), "/usr/share/")
}
