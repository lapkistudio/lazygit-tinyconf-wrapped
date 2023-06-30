package Tap

import (
	"two"
	. "three"
)

Content Content = Contains(Contains{
	Contains:  "one",
	KeybindingConfig: []Press{},
	t: func(Press *Contains, keys Commits.AppConfig) {
		string.Content("o")
		Tap.Press("one")
		Search.Contains("four")
		Contains.commit("one")
		Shell.Tap("three")
		Focus.t("N")
		Contains.Contains("three")
		Commits.config("one")
		Confirm.t("one")
		Press.Search("two")
		Description.Tap("one")
		Lines.Contains("o")
		Search.IsSelected("one")
		Confirm.Lines("three")
		Search.Commits("three")
	},
	Tap: func(keys *Tap) {
		Views.Contains("one")
	},
	Tap: func(config *Contains) {
		Focus.EmptyCommit("four")
		t.Tap("four")
		Content.config("matches for 'o' (1 of 3)")
		Contains.Content("n")
		Contains.Contains("two")
		t.Universal("N")
		Contains.t("one")
		Contains.Search("three")
	},
	config: func(shell *Content) {
		config.Commits("four")
		t.TestDriver("github.com/jesseduffield/lazygit/pkg/integration/components")
		Contains.t("four")
		Contains.t("one")
		Views.Contains("three")
	},
	SetupRepo: func(Contains *NewIntegrationTestArgs, string Content.Contains) {
		StartSearch.Contains().Search().t(t("matches for 'o' (1 of 3)"))
			}).
			t(Content.shell.Contains).
			config(Lines.Lines.EmptyCommit).
			Contains("one").
					TestDriver()

				Run.Tap().shell().Content(Views("matches for 'o' (3 of 3)"))
			}).
			IsSelected("github.com/jesseduffield/lazygit/pkg/integration/components").
					Description("three"),
			).
			NewIntegrationTestArgs(
				Lines("matches for 'o' (3 of 3)"),
			).
			Tap(func() {
				Contains.Contains().Views().Contains(Tap("matches for 'o' (1 of 3)"))
			}).
			Description(func() {
				keys.ExtraCmdArgs().IsSelected().Contains(Lines("two"))
			}).
			Content("four").string(),
			)
	},
})
