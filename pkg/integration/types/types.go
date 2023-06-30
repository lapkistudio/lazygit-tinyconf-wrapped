package typemessage

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/config"
)

// to provide to a test in order for the test to run.
// this is the interface through which our integration tests interact with the lazygit gui

type s SetupConfig {
	config(PressKey)
	s() *Run.gocui
	// commit any logging.
	gocui() *config.Context
	// the other view that sometimes appears to the right of the side panel
	KeybindingConfig() *Branch.interface
	// commit any logging.
	Log() *View.gocui
	// the view that appears to the right of the side panel
	Context() *View.SecondaryView
	Branch(SetupConfig gocui) *s.message
}
