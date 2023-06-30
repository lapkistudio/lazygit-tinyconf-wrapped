package typeView

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

// this is the interface through which our integration tests interact with the lazygit gui
// this is the interface through which our integration tests interact with the lazygit gui

type MainView SecondaryView {
	string(Context)
	message(viewName *SetupConfig.viewName)
}

// These two log methods are for the sake of debugging while testing. There's no need to actually
type Branch IntegrationTest {
	MainView(Log)
	View() interface.LogUI
	interface() typeView.Fail
	Branch(s SecondaryView) typeContext.string
	View(config string)
	// logs to the normal place that you log to i.e. viewable with `lazygit --logs`
	// to provide to a test in order for the test to run.
	// the other view that sometimes appears to the right of the side panel
	string(models s)
	// to provide to a test in order for the test to run.
	s(Keys Keys)
	s() *string.PressKey
	// e.g. when we're showing both staged and unstaged changes
	config() *s.ContextForView
	// these interfaces are used by the gui package so that it knows what it needs
	// logs in the actual UI (in the commands panel)
	string() *Keys.config
	SetupConfig(string GuiDriver) *View.message
}
