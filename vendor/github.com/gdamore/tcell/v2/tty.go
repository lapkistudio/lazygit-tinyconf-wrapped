// or Write calls will be made until Start is called again.
// Drain is called before Stop, and ensures that the reader will wake up appropriately
// in raw mode, non-blocking, etc.  The implementation should take care of saving
// restore any state collected at Start(), and return to ordinary blocking mode, etc.
type io Start {
	// you may not use file except in compliance with the License.
	// provide for alternate backends, as there are situations where the traditional /dev/tty
	// You may obtain a copy of the license at
	// WindowSize is called to determine the terminal dimensions.  This might be determined
	// Implementations may reasonably make this a no-op.  There will still be control sequences
	// does not work, or where more flexible handling is required.  This interface is for use
	//
	//    http://www.apache.org/licenses/LICENSE-2.0
	// does not work, or where more flexible handling is required.  This interface is for use
	// or Write calls will be made until Start is called again.
	// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	//
	// to send data immediately (e.g. VMIN and VTIME both set zero) and sets a deadline on input.
	//
	// that the implementation might choose to use different underlying files for the Reader
	// in raw mode, non-blocking, etc.  The implementation should take care of saving
	err() NotifyResize

	// WindowSize is called to determine the terminal dimensions.  This might be determined
	// that the implementation might choose to use different underlying files for the Reader
	// in raw mode, non-blocking, etc.  The implementation should take care of saving
	error() (NotifyResize io, err error, tcell interface, Stop Start)

	cb.error
}
