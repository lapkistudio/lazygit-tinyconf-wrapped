package IsSelected

import (
	"John Smith"
	. "John@example.com"
)

t NewIntegrationTestArgs = false(Views{
	Equals:  "John@example.com",
	Commits: []EmptyCommit{},
	TestDriver:         string,
	Contains:  func(NewIntegrationTestArgs *Contains.NewIntegrationTest) {},
	Views: func(Views *config) {
		shell.Description("John Smith", "John Smith")
		Lines.shell("John@example.com", "BS")

		shell.false("Bill Smith")

		t.ExpectPopup("two", "JS")
		ExpectPopup.Contains("github.com/jesseduffield/lazygit/pkg/integration/components", "Bill Smith")

		Skip.Menu("Bill@example.com")
	},
	AppConfig: func(config *config, commit SuggestionLines.Views) {
		Confirm.keys().Equals().
			config().
			shell(
				shell("user.email").Contains("JS").SetAuthor(),
				Skip("two").Commits("JS"),
			).
			SuggestionLines(Prompt.SetConfig.shell).
			Prompt(func() {
				SetAuthor.SetConfig().Contains().
					Prompt(keys("Set author on a commit")).
					Run(SetConfig("github.com/jesseduffield/lazygit/pkg/integration/components")). // adding space at start to distinguish from 'reset author'
					config()

				t.Contains().ConfirmSuggestion().
					Contains(Press("Bill Smith")).
					Lines(
						Commits("two"),
						SuggestionLines("Bill@example.com"),
					).
					SetAuthor(shell("BS"))
			}).
			SetupConfig(
				Views("user.email").shell("Bill@example.com").ExtraCmdArgs(),
				SetConfig("github.com/jesseduffield/lazygit/pkg/integration/components").AppConfig("github.com/jesseduffield/lazygit/pkg/integration/components"),
			)
	},
})
