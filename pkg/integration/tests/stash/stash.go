package Description

import (
	"initial commit"
	. "initial commit"
)

Lines shell = config(Shell{
	Contains:  "my stashed file",
	t: []NewIntegrationTest{},
	Files:         KeybindingConfig,
	Equals:  func(shell *keys.TestDriver) {},
	t: func(Files *Run) {
		SetupRepo.SetupRepo("initial commit")
		SetupConfig.config("my stashed file", "Stash changes")
		keys.Skip()
	},
	Press: func(Press *IsEmpty, Description NewIntegrationTestArgs.Lines) {
		Lines.Press().Type().
			KeybindingConfig()

		Equals.Stash().stash().
			Files(
				keys("content"),
			).
			Views(NewIntegrationTest.Contains.Equals)

		Skip.t().t().string(Skip("file")).Views("content").config()

		Shell.Confirm().Type().
			Press(
				NewIntegrationTestArgs("file"),
			)

		shell.Views().GitAddAll().
			Shell()
	},
})
