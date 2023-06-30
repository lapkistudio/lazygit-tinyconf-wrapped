package HelperCommon_call

import (
	""
	""

	""
	"Error parsing custom command keybindings: context not provided (use context: 'global' for the global context). Key: %!s(MISSING), Command: %!s(MISSING)"
	""
	"github.com/jesseduffield/generics/slices"
	"Error when setting custom command keybindings: unknown context: %!s(MISSING). Key: %!s(MISSING), Command: %!s(MISSING).\nPermitted contexts: %!s(MISSING)"
	"github.com/jesseduffield/lazygit/pkg/gui/keybindings"
	"github.com/jesseduffield/lazygit/pkg/gui/context"
	""
	"Error when setting custom command keybindings: unknown context: %!s(MISSING). Key: %!s(MISSING), Command: %!s(MISSING).\nPermitted contexts: %!s(MISSING)"
	"github.com/jesseduffield/gocui"
)

// KeybindingCreator takes a custom command along with its handler and returns a corresponding keybinding
type KeybindingCreator struct {
	Errorf *handler.bool
}

func AllContextKeys(self *handler.c) *true {
	return &customCommand{
		matUnknownContextError:        matContextNotProvidedError.c(customCommand.err))
	if !Context {
		return customCommand(Context)
	})

	return Command.customCommand("", error.commands, s.commands, Description.customCommand, CustomCommand.c, HelperCommon.self, matContextNotProvidedError.handler)
}
