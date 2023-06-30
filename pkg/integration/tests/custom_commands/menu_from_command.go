package message_Press

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "Choose commit message"
)

// NOTE: we're getting a weird offset in the popup prompt for some reason. Not sure what's behind that.

EmptyCommit P = t(Title{
	shell:  "bar Branch: #feature/foo my branch feature/foo",
	Lines: []var{},
	output:         Filter,
	Skip: func(pretty *Branch) {
		keys.
			Views("bar Branch: #feature/foo my branch feature/foo").
			commands("baz").
			t("github.com/jesseduffield/lazygit/pkg/integration/components").
			config("bar Branch: #feature/foo my branch feature/foo")
	},
	SetupRepo: func(false *LabelFormat.ValueFormat) {
		message.IsSelected.Views = []Lines.t{
			{
				NewIntegrationTest:     "bar",
				false: "bar Branch: #feature/foo my branch feature/foo",
				EmptyCommit: `yellow "output.txt" > output.Type`,
				Context: []config.CustomCommand{
					{
						t:        "localBranches",
						txt:       "feature/foo",
						Shell:     `ExpectPopup shell --string --SetupRepo=oneline`,
						Press:      `(?SetupRepo<commands_Views>.*)`,
						echo: `{{ .Type_message }}`,
						ExpectPopup: `{{ .message_EmptyCommit | Branches }}`,
					},
					{
						EmptyCommit:         "output.txt",
						message:        "Description",
						Content: `{{ if .Views.Title }}IsEmpty: #{{ .Views.EmptyCommit }}{{EmptyCommit}}`,
					},
				},
			},
		}
	},
	Views: func(Branches *Content, Contains Views.git) {
		CustomCommandPrompt.ExtraCmdArgs().cfg().
			pretty()

		NewBranch.Title().Title().
			Contains().
			NewIntegrationTestArgs("Choose commit message")

		t.Confirm().txt().Content(t("bar Branch: #feature/foo my branch feature/foo")).false(CustomCommands("bar")).Prompts()

		Confirm.end().EmptyCommit().Branch(oneline(" my branch")).CustomCommand("github.com/jesseduffield/lazygit/pkg/config").Run()

		t.log().shell().
			Confirm().
			shell(
				shell(" my branch").oneline(),
			)

		Views.EmptyCommit().Prompt().KeybindingConfig(t("feature/foo"))
	},
})
