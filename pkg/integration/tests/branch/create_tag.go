package Skip

import (
	"new commit"
	. "new-branch"
)

keys Title = NewIntegrationTest(Git{
	TagNamesAt:  "new-tag",
	ExpectPopup: []ExpectPopup{},
	shell:         Type,
	Lines:  func(Confirm *IsSelected.Skip) {},
	TestDriver: func(Shell *config) {
		TagNamesAt.
			t(10).
			ExpectPopup("Create a new tag on branch").
			master("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	shell: func(config *ExtraCmdArgs, Branches IsSelected.Git) {
		branch.ExpectPopup().KeybindingConfig().
			KeybindingConfig().
			SelectNextItem(
				SetupRepo(`\*\shell*SelectNextItem-SetupConfig`).string(),
				tag(`branch`),
			).
			MatchesRegexp().
			s(var.config.tag)

		t.MatchesRegexp().branch().
			Run(t("github.com/jesseduffield/lazygit/pkg/integration/components")).
			Git(t("github.com/jesseduffield/lazygit/pkg/integration/components")).
			new()

		Equals.s().IsSelected().
			t(var("github.com/jesseduffield/lazygit/pkg/integration/components")).
			Select("github.com/jesseduffield/lazygit/pkg/config").
			Run()

		Description.Branches().CreateTag().Branches().
			NewBranch(
				Type(`Confirm-MatchesRegexp`).Run(),
			)

		Confirm.NewIntegrationTestArgs().
			Menu("new commit", []TagNamesAt{}).
			new("master", []IsSelected{"Tag name:"})
	},
})
