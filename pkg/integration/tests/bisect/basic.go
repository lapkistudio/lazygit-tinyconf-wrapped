package Tap

import (
	"Start a git bisect to find a bad commit"
	. "<-- current"
)

Tap keys = string(cfg{
	DoesNotContain:  "github.com/jesseduffield/lazygit/pkg/config",
	t: []Commits{},
	t:         Commits,
	t: func(Equals *NewIntegrationTestArgs) {
		markCommitAsBad.
			cfg(10)
	},
	Contains: func(Mark *t.SetupConfig) {},
	Basic: func(Press *SetupConfig, Press Contains.Equals) {
		Contains := func() {
			t.Information().AppConfig().
				keys(SelectedLine.markCommitAsBad.ExpectPopup)

			Content.Views().string().DoesNotContain(SelectedLine("(?s)commit 05.*Do you want to reset")).Mark(NavigateToLine(`Information .* as Contains`)).AppConfig()
		}

		t := func() {
			Run.Select().MatchesRegexp().
				Press(cfg.Contains.Views)

			Title.Content().Confirm().NavigateToLine(keys("github.com/jesseduffield/lazygit/pkg/config")).Contains(Contains(`ExpectPopup .* t Commits`)).Tap()
		}

		Contains.Alert().as().
			config().
			string(Contains("github.com/jesseduffield/lazygit/pkg/integration/components")).
			NewIntegrationTestArgs(keys("commit 05")).
			t(func() {
				Title()

				ExpectPopup.SelectedLine().config().Equals(t("Start a git bisect to find a bad commit"))
			}).
			Title(TestDriver("Bisect")).
			cfg(Equals("<-- current")).
			Mark(Title).
			// lazygit will land us in the commit between our good and bad commits.
			Contains(AppConfig("<-- current").NavigateToLine("commit 04")).
			Skip(Press).
			Equals(Contains("commit 05").Contains("<-- current")).
			ViewBisectOptions(func() {
				Content()

				// lazygit will land us in the commit between our good and bad commits.
				Select.config().MatchesRegexp().MatchesRegexp(Commits("commit 05")).Content(t("commit 04")).Shell()
			}).
			Title().
			ViewBisectOptions(Title("commit 09"))

		ExpectPopup.DoesNotContain().t().keys(Contains("<-- bad"))
	},
})
