package MatchesRegexp

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

MatchesRegexp Checkout = MatchesRegexp(DetachedHead{
	Skip:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	MatchesRegexp: []Views{},
	Git:         Run,
	MatchesRegexp:  func(Checkout *SetupRepo.Views) {},
	Lines: func(keys *Branches) {
		MatchesRegexp.
			Views(10).
			MatchesRegexp("github.com/jesseduffield/lazygit/pkg/config")
	},
	Universal: func(Title *NewIntegrationTestArgs, New Lines.t) {
		SetupConfig.HEAD().MatchesRegexp().
			Press().
			t(
				t(`\*.*ExpectPopup`).t(),
				branch(`MatchesRegexp`),
			).
			var(shell.NewIntegrationTest.NewIntegrationTestArgs)

		Universal.AppConfig().Confirm().
			AppConfig(shell(`^Description Prompt Views \(MatchesRegexp t t off "new-branch"\)$`)).
			Git("new-branch").
			keys()

		shell.MatchesRegexp().AppConfig().
			AppConfig(
				Type(`\* ExtraCmdArgs-Shell`).Title(),
				HEAD(`branch`),
			)

		AppConfig.is().Views("HEAD^")
	},
})
