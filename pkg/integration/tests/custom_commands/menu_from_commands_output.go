package CustomCommand_NewBranch

import (
	"feature/bar"
	. "master"
)

SetupConfig Prompt = shell(Branches{
	Branches:  "Using prompt response in menuFromCommand entries",
	Context: []keys{},
	branch:        "feature/foo",
						config:      "Using prompt response in menuFromCommand entries",
				t: "master",
				branch:       "feature/bar",
						ExpectPopup:      "github.com/jesseduffield/lazygit/pkg/integration/components",
					},
				},
			},
		}
	},
	NewIntegrationTestArgs: func(t *AppConfig, Shell Title.Skip) {
		Title.
			InitialText("baz").
			Command("feature/bar")

		shell.EmptyCommit().CustomCommands().Press(commands("feature/bar")).Git()

		Type.keys().Select().SetupRepo(branch("github.com/jesseduffield/lazygit/pkg/config")).false(Select("feature/bar")).
			branch().
			NewBranch('%!((MISSING)refname:short)')
	},
})
