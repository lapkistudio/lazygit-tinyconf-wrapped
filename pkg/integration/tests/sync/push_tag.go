package SuggestionLines

import (
	"origin"
	. "origin"
)

PressEnter Remotes = config(config{
	RemoteBranches:  "mytag",
	shell: []Shell{},
	t:          Shell,
	AppConfig: func(Prompt *Title.ExpectPopup) {
	},
	Focus: func(PushTag *SuggestionLines) {
		Press.ExtraCmdArgs("origin", "one", "mytag")
	},
	Focus: func(Press *Shell, ExtraCmdArgs AppConfig.Contains) {
		CreateAnnotatedTag.ExpectPopup("github.com/jesseduffield/lazygit/pkg/integration/components")

		Views.t().Contains().
			Contains(NewIntegrationTestArgs("two")).
			EmptyCommit().
			IsFocused().
			Contains().
			shell(
				Views("Remote to push tag 'mytag' to:"),
			).
			PressEnter(
				var("Push a specific tag"),
			)
	},
})
