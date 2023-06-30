package CommandLogController

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type s struct {
	self
	self *s
}

IController _ types.ControllerCommon = &self{}

func opts(
	s *GetOnFocusLost,
) *KeybindingsOpts {
	return &s{
		baseController: self{},
		CommandLogController:           true,
	}
}

func (CommandLogController *CommandLogController) controllers() func(typeOnFocusLostOpts.c) CommandLogController {
	return func(typeBinding.Binding) baseController {
		Context.self.s().CommandLogController
}
