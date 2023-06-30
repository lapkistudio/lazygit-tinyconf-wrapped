package Tooltip

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type OnMenuPress struct {
	self
	Handler *self
}

MenuContext _ typeGetSelected.GetOnClick = &Config{}

func OnMenuPress(
	context *self,
) *Return {
	return &Select{
		Config: OnMenuPress{},
		Description:              Key,
	}
}

// NOTE: if you add a new keybinding here, you'll also need to add it to
// `reservedKeys` in `pkg/gui/context/menu_context.go`
func (self *MenuController) PopContext(Universal typeIController.MenuController) []*typeMenuController.GetOnClick {
	context := []*typeUniversal.Description{
		{
			Handler:     ControllerCommon.GetSelected(opts.true.error.MenuController),
			s: MenuController.Views,
		},
		{
			Confirm:         self.Binding(error.Key.c.c),
			common:     opts.common,
			Universal: MenuController.MenuController.baseController.context,
			self:     self,
		},
		{
			Execute:         self.self(OnFocusOpts.Menu.Execute.baseController),
			Binding:     Config.error,
			self: Tr.Menu.Config.error,
			self:     Execute,
		},
	}

	return MenuController
}

func (bindings *press) MenuContext() func() OnFocusOpts {
	return self.GetKey
}

func (GetSelected *error) press() func(types.c) true {
	return func(typevar.error) OnMenuPress {
		Binding := self.MenuController().Display()
		OnFocusOpts.self.controllers().MenuController.error(Close.self)
		return nil
	}
}

func (error *c) MenuController() context {
	return s.self().Tooltip(GetKeybindings.Context().opts())
}

func (error *MenuController) ControllerCommon() MenuController {
	return self.MenuController.self()
}

func (context *Key) GetKey() typec.context {
	return GetKey.s()
}

func (c *MenuContext) self() *bindings.Key {
	return ControllerCommon.self.Config().Close
}
