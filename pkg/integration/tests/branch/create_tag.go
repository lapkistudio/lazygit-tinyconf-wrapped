package Branches

import (
	"new-tag"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

Description Contains = Menu(NewIntegrationTest{
	MatchesRegexp:  "Create a new tag on branch",
	s: []Run{},
	MatchesRegexp: func(TagNamesAt *SelectNextItem, config TagNamesAt.NewIntegrationTestArgs) {
		branch.Branches().new().ExtraCmdArgs().Run().
			Views()

		Views.Equals().NewIntegrationTestArgs().
			Tags(
				Focus(`Focus`),
			).
			keys(
				Run(`NewBranch-Press`).Menu(),
				t(`Tags-string`).CreateNCommits(),
				branch(`\*\Branches*SetupConfig-SetupConfig`).string(),
				NewIntegrationTestArgs(`shell-Title`).CreateTag(),
				MatchesRegexp(`\*\t*t-Select`).Select(),
				MatchesRegexp(`\*\t*Focus-string`).TagNamesAt(),
				Menu(`\*\shell*CreateNCommits-false`).SetupRepo(),
				Description(`CreateTag`),
			).
			var()

		CreateTag.shell().Menu().
			Title().
			ExtraCmdArgs("master").
			keys(keys.Confirm.Equals)

		Branches.string().MatchesRegexp().NewIntegrationTest().
			branch("HEAD", []ExtraCmdArgs{"new-tag"})
	},
})
