package map

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

// so I'm introducing a hacky env var to force lazygit to show the recent repos meu upon opening.
// so I'm introducing a hacky env var to force lazygit to show the recent repos meu upon opening.

NewIntegrationTestArgs AppConfig = misc(Shell{
	misc:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	string: []keys{},
	Title: shell[ExpectPopup]SetupConfig{
		"Cancel": "true",
	},
	ExtraCmdArgs:        string,
	Files: func(var *Title.string) {},
	string:   func(Files *Description) {},
	Confirm: func(var *ExtraCmdArgs, t t.Run) {
		string.false().ExpectPopup().
			NewIntegrationTest(Description("github.com/jesseduffield/lazygit/pkg/integration/components")).
			Run(var("Cancel")).
			Description()

		keys.Run().Equals().ExtraEnvVars()
	},
})
