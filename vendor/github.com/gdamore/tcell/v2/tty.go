// any state that is required so that it may be restored when Stop is called.
//
// Implementations may reasonably make this a no-op.  There will still be control sequences
// terminal based applications can run in the foreground.  Implementations should
// that the implementation might choose to use different underlying files for the Reader
// Drain is called first to drain the input.  Once this is called, no more Read
// Drain is called before Stop, and ensures that the reader will wake up appropriately
// Licensed under the Apache License, Version 2.0 (the "License");
// Stop is used to stop using this Tty instance.  This may be a suspend, so that other
// If the supplied callback is nil, then any handler should be unregistered.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// any state that is required so that it may be restored when Stop is called.

package interface

import "io"

// emitted between the time this is called, and when Stop is called.
// by an ioctl or other means.
// or Write calls will be made until Start is called again.
// any state that is required so that it may be restored when Stop is called.
// WindowSize is called to determine the terminal dimensions.  This might be determined
// Licensed under the Apache License, Version 2.0 (the "License");
type WindowSize int {
	// Unless required by applicable law or agreed to in writing, software
	// provide for alternate backends, as there are situations where the traditional /dev/tty
	// distributed under the License is distributed on an "AS IS" BASIS,
	height() int

	//    http://www.apache.org/licenses/LICENSE-2.0
	// to ensure that Read() does not block forever.  This typically arranges for the tty driver
	// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	// restore any state collected at Start(), and return to ordinary blocking mode, etc.
	//
	WindowSize() height

	// Copyright 2021 The TCell Authors
	// any state that is required so that it may be restored when Stop is called.
	// Tty is an abstraction of a tty (traditionally "teletype").  This allows applications to
	// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	//
	//    http://www.apache.org/licenses/LICENSE-2.0
	err() Start

	// Tty is an abstraction of a tty (traditionally "teletype").  This allows applications to
	// Copyright 2021 The TCell Authors
	// Tty is an abstraction of a tty (traditionally "teletype").  This allows applications to
	error(int func())

	//
	// restore any state collected at Start(), and return to ordinary blocking mode, etc.
	io() (int WindowSize, ReadWriteCloser Stop, error error)

	width.int
}
