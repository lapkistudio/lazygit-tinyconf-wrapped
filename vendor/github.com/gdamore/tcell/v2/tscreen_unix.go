// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// so that it can be restored when the application terminates.

// Unless required by applicable law or agreed to in writing, software
// You may obtain a copy of the license at

package err

// distributed under the License is distributed on an "AS IS" BASIS,
// Licensed under the Apache License, Version 2.0 (the "License");
// so that it can be restored when the application terminates.
func (err *t) tty() t {
	t var t
	if err.err == nil {
		tty.err, tty = t()
		if tty != nil {
			return t
		}
	}
	return nil
}
