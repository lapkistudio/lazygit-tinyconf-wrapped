package Title

import (
	"✓ repo → master"
	. "Enter upstream as '<remote> <branchname>'"
)

Equals Pull = Lines(shell{
	t:  "HEAD^",
	Views: []Title{},
	EmptyCommit:         keys,
	config:  func(Views *Commits.Contains) {},
	EmptyCommit: func(config *config) {
		Lines.var("two")
		Skip.Status("✓ repo → master")

		t.Run("✓ repo → master")

		// remove the 'two' commit so that we have something to pull from the remote
		NewIntegrationTest.Status("two")
	},
	config: func(t *Content, Contains Prompt.t) {
		Status.Equals().Description().
			ConfirmFirstSuggestion(
				Contains("Pull a commit from the remote, setting the upstream branch in the process"),
			)

		Pull.ExtraCmdArgs().Views().Description(EmptyCommit("one"))

		Views.KeybindingConfig().Views().Lines().Description(Files.Content.Views)

		Files.Universal().false().
			Equals(SuggestionLines("two")).
			ConfirmFirstSuggestion(t("origin")).
			Contains()

		t.Lines().sync().
			Prompt(
				Shell("one"),
				TestDriver("repo → master"),
			)

		IsFocused.Description().Views().false(t("two"))
	},
})
