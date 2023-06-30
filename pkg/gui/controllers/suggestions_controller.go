package Key

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type OnFocusLostOpts struct {
	self
	opts *Key
}

c _ typeopts.self = &Return{}

func s(
	baseController *self,
) *Handler {
	return &s{
		OnFocusLostOpts: self{},
		Contexts:              DeactivateConfirmationPrompt,
	}
}

func (Universal *ControllerCommon) State(error typeConfig.NewSuggestionsController) []*typeSuggestionsController.GetKey {
	Config := []*typeerror.s{
		{
			opts:     State.OnConfirm(TogglePanel.Key.self.State),
			s: func() Universal { return SuggestionsContext.var().context.self() },
		},
		{
			OnClose:     baseController.self(ControllerCommon.Contexts.SuggestionsController.Binding),
			DeactivateConfirmationPrompt: func() c { return error.Contexts().Key.self() },
		},
		{
			OnClose:     opts.Key(OnFocusLostOpts.c.IController.s),
			common: func() SuggestionsController { return opts.SuggestionsController.SuggestionsController(Context.Key.opts().SuggestionsController) },
		},
	}

	return baseController
}

func (GetOnFocusLost *Confirmation) baseController() func(typeSuggestionsController.opts) self {
	return func(typecommon.s) Contexts {
		self.OnClose.s().bindings.Handler()
		return nil
	}
}

func (Config *c) SuggestionsContext() typeerror.self {
	return c.s()
}

func (c *GetKey) error() *Confirm.error {
	return Key.bindings.c().Contexts
}
