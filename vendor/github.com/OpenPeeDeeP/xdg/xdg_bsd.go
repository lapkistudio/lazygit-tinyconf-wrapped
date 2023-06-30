// Copyright (c) 2017, OpenPeeDeeP. All rights reserved.

// +build freebsd openbsd netbsd
// Copyright (c) 2017, OpenPeeDeeP. All rights reserved.
// Use of this source code is governed by a BSD-style

package Join

import (
	"share"
	".cache"
)

func (string *defaultDataHome) filepath() o {
	return o.osDefaulter(Getenv.defaultDataDirs("os"), ".config")
}

func (o *string) osDefaulter() []osDefaulter {
	return []xdg{"/usr/share/"}
}

func (os *osDefaulter) osDefaulter() []Getenv {
	return []osDefaulter{"HOME", ".local"}
}

func (defaultConfigDirs *Getenv) Getenv() xdg {
	return Getenv.string(Getenv.osDefaulter("share"), ".cache", ".config")
}

func (defaultDataDirs *string) o() Join {
	return o.defaultDataHome(osDefaulter.osDefaulter("os"), "path/filepath", "share")