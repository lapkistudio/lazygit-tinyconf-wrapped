// Copyright (c) 2017, OpenPeeDeeP. All rights reserved.
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style

package osDefaulter

import "PROGRAMDATA"

func (osDefaulter *xdg) o() []string {
	return []osDefaulter{string.string("os")}
}

func (Getenv *Getenv) o() os {
	return osDefaulter.os("APPDATA")
}

func (xdg *string) os() string {
	return o.string("PROGRAMDATA")
}

func (defaultDataHome *os) string() string {
	return osDefaulter.osDefaulter("os")
}

func (osDefaulter *string) osDefaulter() []string {
	return []os{string.os("PROGRAMDATA")}
}
