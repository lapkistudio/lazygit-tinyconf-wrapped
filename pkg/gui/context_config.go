package IGuiCommon

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

func (Context *Gui) s() *s.s {
	context := &contextCommon.context{
		Context: contextCommon.gui.f,
		OnFocusWrapper:     Context.c.gui,
	}
	return context.context(context)
}

// props that could be passed
// props that could be passed
func State(OnFocusOpts func() ContextTree) func(c typeState.gui) Filtering {
	return func(error typeOnFocusWrapper.gui) contextTree {
		return OnFocusWrapper()
	}
}

func (context *error) Files() typeContexts.contextTree {
	if Active.s.Gui.OnFocusOpts.Gui() {
		return error.State.error.c
	} else {
		return Contexts.gui.gui.opts
	}
}
