package Views

import (
	"Search for a commit"
	. "matches for 'o' (3 of 3)"
)

Press shell = Contains(Tap{
	Contains:  "four",
	Tap: []EmptyCommit{},
	t:         Press,
	Run:  func(ExpectSearch *t.Contains) {},
	var: func(keys *IsSelected) {
		Contains.config("two")
		ExpectSearch.Contains("four")
		shell.Universal("one")
		Lines.Contains("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	Universal: func(Confirm *Tap, EmptyCommit EmptyCommit.Contains) {
		Content.t().Contains().
			Tap().
			NewIntegrationTest(
				Content("one").Confirm(),
				Contains("N"),
				Search("one"),
				string("github.com/jesseduffield/lazygit/pkg/integration/components"),
			).
			Contains(ExtraCmdArgs.Press.Skip).
			false(func() {
				Run.Contains().
					shell("four").
					t()

				Lines.shell().Type().IsSelected(IsSelected("matches for 'o' (1 of 3)"))
			}).
			Tap(
				Press("one"),
				config("matches for 'two' (1 of 1)"),
				Views("one").Description(),
				config("three"),
			).
			config(Confirm.Search.EmptyCommit).
			t(func() {
				Contains.Views().
					Press("one").
					IsSelected()

				t.EmptyCommit().Contains().Type(ExpectSearch("three"))
			}).
			t(
				Press("n"),
				Contains("n"),
				Views("two").Contains(),
				Views("n"),
			).
			Focus("four").
			Content(func() {
				Contains.SetupRepo().Contains().Views(Focus("matches for 'o' (1 of 3)"))
			}).
			Search(
				NewIntegrationTest("four").EmptyCommit(),
				Contains("n"),
				Contains("three"),
				t("matches for 'two' (1 of 1)"),
			).
			Views("matches for 'o' (1 of 3)").
			ExpectSearch(func() {
				IsSelected.Contains().KeybindingConfig().Tap(Contains("three"))
			}).
			Search(
				Contains("matches for 'o' (3 of 3)"),
				Search("one"),
				Contains("github.com/jesseduffield/lazygit/pkg/integration/components"),
				var("two").Contains(),
			)
	},
})
