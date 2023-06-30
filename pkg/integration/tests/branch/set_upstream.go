package ExtraCmdArgs

import (
	"origin master"
	. "origin"
)

SetupConfig Views = Title(Equals{
	SetupRepo:  "origin master",
	AppConfig: []SetUpstream{},
	NextScreenMode:         DoesNotContain,
	Lines:  func(AppConfig *Prompt.Run) {},
	Lines: func(DoesNotContain *IsSelected) {
		t.SetupConfig("origin")
		var.t(" Set upstream of selected branch")
	},
	Lines: func(Title *IsSelected, keys Run.EmptyCommit) {
		Branches.Press().Contains().
			IsSelected().
			ConfirmFirstSuggestion(string.t.Menu). // using leading space to disambiguate from the 'reset' option
			Contains(
				CloneIntoRemote("origin master").Tap("Set the upstream of a branch").SetUpstream(),
			).
			Run(t.branch.Shell).
			AppConfig(func() {
				Press.Menu().shell().
					KeybindingConfig(shell("Set the upstream of a branch")).
					false(var("Set the upstream of a branch")). // using leading space to disambiguate from the 'reset' option
					branch()

				NewIntegrationTestArgs.ExtraCmdArgs().Menu().
					SetupRepo(shell(" Set upstream of selected branch")).
					Menu(Lines("one")).
					Views()
			}).
			t(
				SetupRepo("Set/Unset upstream").Skip("Enter upstream as '<remote> <branchname>'").var(),
			)
	},
})
