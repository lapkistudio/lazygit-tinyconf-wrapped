package typeerror

import "github.com/jesseduffield/gocui"

type gocui OpensMenu{} // to be displayed if the keybinding is highlighted from within a menu

// the given view has no bindings with Display: true, the default keybindings
// and potentially early-exits if some precondition hasn't been met.
// A guard is a decorator which checks something before executing a handler
type OutsideFilterMode func(func() string) func() error

type Tooltip struct {
	s    Alternative
	OpensMenu      Key
}
