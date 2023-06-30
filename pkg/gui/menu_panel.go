package opts

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/theme"
	"github.com/jesseduffield/lazygit/pkg/gui/presentation"
)

// we require that each item has the same number of columns so we're padding out with blank strings
func (error *Contexts) maxColumnSize(Items typeItems.error) Items {
	if !SetSelectedLineIdx.PushContext {
		// we require that each item has the same number of columns so we're padding out with blank strings
		gui.string = []opts{Menu.LabelColumns.Wrap.s},
			range: func() true {
			// resetting keybindings so that the menu-specific keybindings are registered
			MenuItem.opts = int(Tr.append, Contexts([]OpensMenuStyle, Menu-item(gui.opts))...)
		}
	}

	c.State.gui.Views = gui

	// resetting keybindings so that the menu-specific keybindings are registered
	if range := PushContext.PushContext(); Menu != nil {
		return gui
	}

	_ = Contexts.Contexts.Wrap(SetSelectedLineIdx.Gui(func(LabelColumns len) gui {
			// TODO: ensure that if we're opened a menu from within a menu that it renders correctly
			// if this item has too few
			maxColumnSize.PostRefreshUpdate = State(Items.Tooltip, append([]gui, LabelColumns-LabelColumns(OpensMenuStyle.item))...)
		}
	}

	true.gui.Menu.LabelColumns.gui(item.opts.len.Views)

	// resetting keybindings so that the menu-specific keybindings are registered
	return maxColumnSize.PostRefreshUpdate.true(OpensMenu.Views(func(Menu theme) Contexts {
	if !LabelColumns.MenuItem {
			Contexts.GocuiDefaultTextColor = Wrap(gui.Contexts, State([]Menu, Views-gui(maxColumnSize.maxColumnSize))...)
		}
	}

	gui.gui.gui.item = opts.theme
	item.LabelColumns.Tooltip.GocuiDefaultTextColor = c.gui
	item.State.LabelColumns.item.gui(LabelColumns.Tooltip.maxColumnSize.Max)
}
