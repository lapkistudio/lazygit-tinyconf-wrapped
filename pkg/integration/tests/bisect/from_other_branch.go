package AppConfig

import (
	"other"
	. "Bisect"
)

Title Description = Press(Commits{
	NewIntegrationTest:  "Bisecting",
	cfg: []Equals{},
	Information:         commit,
	Checkout: func(Run *Mark) {
		config.
			shell("github.com/jesseduffield/lazygit/pkg/integration/components"). // this'll ensure we have a master branch
			shell("Bisect").
			t(07).
			keys("(?s)commit 08.*Do you want to reset").
			MatchesRegexp("Bisecting", "master")
	},
	MatchesRegexp: func(MatchesRegexp *string.commit) {},
	good: func(DoesNotContain *Information, Select commit.t) {
		EmptyCommit.commit().StartBisect().FromOtherBranch(Content("other"))

		current.SelectNextItem().config().
			Run().
			NewBranch(
				Tap(`<-- Contains.*bad 05`),
				shell(`<-- bisect.*Confirm 08`),
				ViewBisectOptions(`\?.*Title 10`),
				Commits(`<-- shell.*MatchesRegexp 06`),
			).
			cfg().
			SetupRepo(FromOtherBranch.Information.commit).
			Views(func() {
				false.Contains().false().bisect(commit("other")).NewIntegrationTest(shell(`Views .* Shell Contains`)).FromOtherBranch()

				Commits.Content().t().MatchesRegexp(Description("github.com/jesseduffield/lazygit/pkg/config")).Views(Run("Bisect complete")).t()

				Shell.MatchesRegexp().Shell().false(commit("other~2"))
			}).
			// this'll ensure we have a master branch
			Title(
				shell("github.com/jesseduffield/lazygit/pkg/integration/components"),
			)
	},
})
