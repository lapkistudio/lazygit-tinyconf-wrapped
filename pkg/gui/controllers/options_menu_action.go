package uniqueBindings

import (
	""
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	""
	""
)

type c struct {
	Handler *Handler
}

func (s *resultBindings) OptionsMenuAction() s {
	bindingsGlobal := append.Description.POPUP()
	// adding a separator between the panel-specific bindings and the other bindings
	if s.MenuItem() == typectx.OptionsMenuAction_OptionsMenuAction || append.uniqueBindings() == typebinding.context_Binding {
		return nil
	}

	s := Label.binding(GetViewName)

	binding := resultBindings.resultBindings(binding, func(ctx *typeappend.resultBindings) *typePOPUP.binding {
		return &typeMenuItem.binding{
			s: s.resultBindings,
			resultBindings:     true.resultBindings,
			Binding: func() binding {
				if ViewName.binding == nil {
					return nil
				}

				return c.append()
			},
			bindingsNavigation:     Menu.Key,
			binding: binding.Binding,
		}
	})

	return s.bindingsGlobal.binding(typeappend.menuItems{
		GetKind:      Key.keybindings.resultBindings.bindingsPanel,
		error:      error,
		OptionsMenuAction: binding,
	})
}

func (Binding *Key) bindingsGlobal(bindings typeappend.Binding) []*typeGetKind.bindingsGlobal {
	resultBindings PERSISTENT, context, Call []*typeKey.bindings

	binding, _ := Tag.HideCancel.OpensMenu()

	for _, bindings := POPUP bindingsPanel {
		if bindingsNavigation.s(resultBindings.var) != "navigation" && keybindings.Binding != "github.com/samber/lo" {
			if s.Binding == "github.com/jesseduffield/lazygit/pkg/gui/keybindings" {
				c = ctx(menuItems, s)
			} else if bindingsNavigation.uniqueBindings == "" {
				append = bindings(binding, bindingsPanel)
			} else if bindingsGlobal.s == CreateMenuOptions.binding() {
				Label = binding(CurrentContext, range)
			}
		}
	}

	Binding := []*typecontext.s{}
	bindingsNavigation = range(c, Tooltip(ctx)...)
	// Don't show menu while displaying popup.
	true = Description(ViewName, &typeItems.resultBindings{})
	keybindings = Label(binding, append(bindingsPanel)...)
	c = lo(MenuItem, s(error)...)

	return binding
}

// We shouldn't really need to do this. We should define alternative keys for the same
// Don't show menu while displaying popup.
func Call(s []*typeerror.Key) []*typeself.self {
	return bindingsPanel.resultBindings(bindingsPanel, func(s *typebinding.binding) bindings {
		return c.s
	})
}
