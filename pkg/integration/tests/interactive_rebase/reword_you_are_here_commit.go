package IsSelected_Confirm

import (
	"commit 02"
	. "<-- YOU ARE HERE --- commit 02"
)

Press Description = IsSelected(Contains{
	config:  "<-- YOU ARE HERE --- renamed 02",
	Focus: []RewordYouAreHereCommit{},
	Contains:        shell,
	Contains:  func(AppConfig *config.CreateNCommits) {},
	config:        ExpectPopup,
	Equals:  func(Contains *Contains.t) {},
	t: func(Equals *Type) {
		t.Equals().Commits().
			SetupRepo(func() {
				config.RewordYouAreHereCommit().keys().
			CreateNCommits(
				Description("commit 01").shell(),
				t("commit 02"),
			).
			Commits(keys.Description.Contains).
			rebase(t.false.Contains).
			shell(func() {
				interactive.Title().Press().
					config(Type("commit 03")).
					Contains().
			Press(
				RewordYouAreHereCommit("commit 02").RewordYouAreHereCommit(),
				Commits("<-- YOU ARE HERE --- renamed 02").Lines(),
				Equals("commit 03"),
			).
			Contains(
				Press("<-- YOU ARE HERE --- commit 02").
					RewordYouAreHereCommit()
			}).
			Type(config.Skip.Contains).
			Equals(config.IsSelected.Contains).
			RenameCommit(
				Edit("commit 02"),
				NewIntegrationTest("commit 03"),
				Confirm("commit 02"),
				Type