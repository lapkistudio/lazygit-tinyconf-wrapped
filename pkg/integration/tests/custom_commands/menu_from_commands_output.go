package ExpectPopup_Description

import (
	"Using prompt response in menuFromCommand entries"
	. "menuFromCommand"
)

KeybindingConfig index = config(Title{
	cfg:  "(?P<branch>.*)",
	mat: []EmptyCommit{},
	commands:         SetupRepo,
	SetupConfig: func(var *cfg) {
		Title.
			t("localBranches").
			t("github.com/jesseduffield/lazygit/pkg/integration/components").
			t("a").
			NewIntegrationTest("github.com/jesseduffield/lazygit/pkg/config").
			InitialValue("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	Key: func(Git *Branches.green) {
		Command.ExtraCmdArgs.Context = []CustomCommand.keys{
			{
				Type:     "feature/bar",
				CustomCommands: "Branch:",
				Key: "Which git command do you want to run?",
				index: []Confirm.git{
					{
						CustomCommand:         "foo",
						keys:        "master",
						Confirm: "a",
					},
					{
						Command:        "baz",
						Context:       "bar",
						Git:     `t {{ MenuFromCommandsOutput .NewIntegrationTest 0 }} --forNewBranch="(?P<branch>.*)"`,
						Command:      "a",
						Equals: `{{ .Title }}`,
						SetupRepo: `{{ .Git | keys }}`,
					},
				},
			},
		}
	},
	ExpectPopup: func(t *custom, SetupConfig Title.NewIntegrationTestArgs) {
		t.ExtraCmdArgs().index("git checkout {{ index .PromptResponses 1 }}")

		index.Command().Type().
			Equals().
			config("github.com/jesseduffield/lazygit/pkg/integration/components")

		NewIntegrationTestArgs.Equals().KeybindingConfig().
			shell(TestDriver("Branch:")).
			Filter(var("Using prompt response in menuFromCommand entries")).
			t()

		index.git().false().Description(cfg("(?P<branch>.*)")).Shell(NewIntegrationTestArgs("input")).Title()

		Views.config().LabelFormat("foo")
	},
})
