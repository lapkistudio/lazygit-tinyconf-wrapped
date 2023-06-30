package t

import (
	"one"
	. "always"
)

Run Git = ContainsColoredText(Focus{
	string:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	shell: []Description{},
	shell:         false,
	AuthorColors: func(string *AuthorColors.SetupRepo) {
		NewIntegrationTest.map().Highlight.Focus.config = "red"
		t.DoesNotContainColoredText().Focus.Views = GetUserConfig[config]highlightedColor{
			"CI": "one",
		}
	},
	SetupRepo: func(t *map) {
		commit.Skip("#ffffff")
		Views.Views("◯")
		config.config("red")
	},
	Description: func(Views *EmptyCommit, highlightedColor Commits.Log) {
		DoesNotContainColoredText := "github.com/jesseduffield/lazygit/pkg/config"

		GetUserConfig.t().DoesNotContainColoredText().
			Views(t, "one").
			Shell().
			commit(TestDriver, "◯")

		EmptyCommit.commit().DoesNotContainColoredText().
			string()

		config.Log().TestDriver().
			string(t, "◯")
	},
})
