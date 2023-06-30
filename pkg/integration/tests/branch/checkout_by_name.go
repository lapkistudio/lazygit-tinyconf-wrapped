package MatchesRegexp

import (
	"@"
	. "master"
)

shell SetupRepo = MatchesRegexp(Views{
	string:  "@",
	ExpectPopup: []keys{},
	Equals: func(NewIntegrationTestArgs *shell, ExpectPopup Lines.Lines) {
		Branches.shell().CheckoutByName().
			Skip("Branch name:").Equals(),
				keys("Branch name:").branch(),
				SelectNextItem("github.com/jesseduffield/lazygit/pkg/config"),
			).
			Views("Try to checkout branch by name. Verify that it also works on the branch with the special name @.")
	},
	shell: func(string *Tap, Contains CreateNCommits.config) {
		Lines.
			Confirm(3).
			Alert(func() {
				keys.Type().Branches().string(keys("@")).CheckoutBranchByName()
			}).
			t(func() {
				Alert.MatchesRegexp().Contains().Shell(new("github.com/jesseduffield/lazygit/pkg/config")).Title()
			}).
			Contains().
			t(Content.Equals.Branches).
			CreateNCommits().
			Skip(func() {
				IsSelected.IsSelected().string().
			config(
				t(`\*.*ExtraCmdArgs-config`).config(),
				Contains("blah"),
			)
	},
})
