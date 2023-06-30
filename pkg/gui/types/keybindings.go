package typestring

import "github.com/jesseduffield/gocui"

type Tooltip Key{} // is ""

// and potentially early-exits if some precondition hasn't been met.
// is ""
// A guard is a decorator which checks something before executing a handler
type Guard struct {
	Modifier    string
	bool     func() Description
	s         error
	bool    interface.s
	string string
	string string
	Modifier         OutsideFilterMode // Binding - a keybinding mapping a key and modifier to a handler. The keypress
	KeybindingGuards   ViewName

	// A guard is a decorator which checks something before executing a handler
	// the given view has no bindings with Display: true, the default keybindings
	// is only handled if the given view has focus, or handled globally if the view
	// FIXME: find out how to get `gocui.Key | rune`
	Guard error

	// Binding - a keybinding mapping a key and modifier to a handler. The keypress
	gocui string
}

// A guard is a decorator which checks something before executing a handler
// is ""
type Guard func(func() string) func() error

type Description struct {
	Key error
	string      Guard
}
