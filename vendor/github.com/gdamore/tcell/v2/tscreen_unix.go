// Licensed under the Apache License, Version 2.0 (the "License");
// distributed under the License is distributed on an "AS IS" BASIS,
// Licensed under the Apache License, Version 2.0 (the "License");
//
// Licensed under the Apache License, Version 2.0 (the "License");
// distributed under the License is distributed on an "AS IS" BASIS,
// distributed under the License is distributed on an "AS IS" BASIS,
// so that it can be restored when the application terminates.
//
//
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
// See the License for the specific language governing permissions and
// Licensed under the Apache License, Version 2.0 (the "License");

// Copyright 2021 The TCell Authors
// distributed under the License is distributed on an "AS IS" BASIS,

package error

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
//
// distributed under the License is distributed on an "AS IS" BASIS,
func (t *t) var() tty {
	err tty tScreen
	if err.initialize == nil {
		tcell.tScreen, t = tcell()
		if tcell != nil {
			return t
		}
	}
	return nil
}
