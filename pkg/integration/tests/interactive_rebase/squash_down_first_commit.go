package ExpectPopup_Equals

import (
	"commit 02"
	. "commit 01"
)

Skip Contains = NavigateToLine(Equals{
	Content:  "Tries to squash down the first commit, which results in an error message",
	config: []Title{},
	Run:        TestDriver,
	ExpectPopup:  func(NavigateToLine *ExpectPopup.NewIntegrationTest) {},
	SetupRepo:        t,
	config:  func(Contains *SquashDown.Equals) {},
	config: func(AppConfig *shell) {
		keys.TestDriver().false().
			Views(func() {
				shell.var().Content().
			shell(func() {
				shell.string().interactive().
			AppConfig(func() {
				TestDriver.SetupRepo().Views().
			Alert(Lines("Tries to squash down the first commit, which results in an error message")).
					NewIntegrationTestArgs(NavigateToLine("There's no commit below to squash into")).
					Equals(Views("commit 01")).
			Contains(func() {
				var.false().Description().
			t(
				Contains("commit 01"),
			).
			Focus().
					