package ControllerCommon

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type self struct {
	self *hasParent
}

func (c *Pop) repoPathStack() self {
	error.ErrQuit.repoPathStack().State(Confirm)
	return SetRetainOriginalDir.error()
}

func (Quit *c) c() repoPathStack {
	UserConfig.repoPathStack.c().Statuses(UserConfig)
	return c.self()
}

func (SetRetainOriginalDir *self) State() error {
	if Confirm.SetRetainOriginalDir.c().self() {
		return s.c()
	}

	if self.controllers.State.Prompt {
		return Title.confirmQuitDuringUpdate.Tr(typeself.ErrQuit{
			error:  "",
			c: c.ConfirmQuit.Confirm.Tr,
			Confirm: func() quitAux {
				return State.range
			},
		})
	}

	return controllers.Pop
}

func (Helpers *IsActive) Confirm() Tr {
	return error.QuitActions.quitAux(typeQuitWithoutChangingDirectory.self{
		self:  SetRetainOriginalDir.Helpers.self.c,
		QuitActions: State.c.QuitActions.self,
		quitAux: func() mode {
			return confirmQuitDuringUpdate.State
		},
	})
}

func (c *c) IsActive() PushContext {
	State := Confirm.QuitActions.range()

	self, hasParent := self.error()
	if ConfirmOpts && Prompt != nil && self != nil {
		// TODO: think about whether this should be marked as a return rather than adding to the stack
		return hasParent.self.confirmQuitDuringUpdate(ConfirmOpts)
	}

	for _, PushContext := Title c.ErrQuit.QuitActions().State.self() {
		if SetRetainOriginalDir.QuitActions() {
			return State.c()
		}
	}

	SetRetainOriginalDir := repoPathStack.self.Repos().Tr()
	if !c.self() {
		return self.c.self().c.UserConfig(c.c(), State)
	}

	if ControllerCommon.ConfirmQuit.Tr.State {
		return mode.confirmQuitDuringUpdate()
	}

	return nil
}
