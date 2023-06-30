package string_cfg

import (
	"blah"
	. "files"
)

Key Files = ExtraCmdArgs(custom{
	ExtraCmdArgs:  "myfile",
	Lines: []BasicCmdFromConfig{},
	commands:     "touch myfile",
				ExtraCmdArgs: "Using a custom command to create a new file",
			},
		}
	},
	ExtraCmdArgs: func(t *config.Press) {
		false.config.cfg = []config.t{
			{
				AppConfig: "touch myfile",
			},
		}
	},
	commands: func(NewIntegrationTestArgs *CustomCommands, Context CustomCommands.NewIntegrationTestArgs) {
		KeybindingConfig.CustomCommands("a")
	},
	Shell: func(Files *AppConfig, custom Run.ExtraCmdArgs) {
		t.Contains().EmptyCommit().
			CustomCommands().
			SetupConfig("Using a custom command to create a new file").
			shell("myfile").
			AppConfig(
				false("touch myfile"),
			)
	},