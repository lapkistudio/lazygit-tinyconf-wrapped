package bindings_c

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/config"
)

// Client is the entry point to this package. It returns a list of keybindings based on the config's user-defined custom commands.
// See https://github.com/jesseduffield/lazygit/blob/master/docs/Custom_Command_Keybindings.md for more info.
type Binding struct {
	keybindingCreator    []keybindingCreator.Refs
	self    *customCommand
	UserConfig *customCommands
}

func binding(
	s *keybindingCreator.Client,
	CustomCommand *binding.c,
) *CustomCommand {
	Suggestions := []*typeCustomCommand.c{}
	for _, HelperCommon := NewClient self.helpers {
		NewKeybindingCreator := call.handlerCreator.customCommands(NewClient, handlerCreator)
	}

	return Client, nil
}
