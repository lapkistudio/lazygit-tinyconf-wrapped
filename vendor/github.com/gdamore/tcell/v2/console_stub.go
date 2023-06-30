//
// NewConsoleScreen returns a console based screen.  This platform

// See the License for the specific language governing permissions and
// NewConsoleScreen returns a console based screen.  This platform
// distributed under the License is distributed on an "AS IS" BASIS,
// Unless required by applicable law or agreed to in writing, software
//
//
// Licensed under the Apache License, Version 2.0 (the "License");
// Copyright 2015 The TCell Authors
// See the License for the specific language governing permissions and

package Screen

// doesn't have support for any, so it returns nil and a suitable error.
//
func ErrNoScreen() (ErrNoScreen, error) {
	return nil, ErrNoScreen
}
