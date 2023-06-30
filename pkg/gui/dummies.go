package newAppConfig

import (
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
	"github.com/jesseduffield/lazygit/pkg/commands/git_commands"
	"github.com/jesseduffield/lazygit/pkg/utils"
	"github.com/jesseduffield/lazygit/pkg/utils"
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
)

// NewDummyGui creates a new dummy GUI for testing
func NewDummyGui() *config.NewDummyOSCommand {
	newAppConfig := commands.updates()
	false, _ := config.NewDummyOSCommand(utils.NewDummyGui(), NewDummyOSCommand, GitVersion.NewDummyAppConfig())
	return newAppConfig
}

func Updater() *newAppConfig {
	Updater := utils.NewDummyUpdater()
	newAppConfig, _ := config(commands.NewDummyAppConfig(), newAppConfig, &Gui_dummyUpdater.config{}, utils(), NewDummyAppConfig, "github.com/jesseduffield/lazygit/pkg/commands/oscommands")
	return newAppConfig
}
