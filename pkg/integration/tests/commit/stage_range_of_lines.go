package t

import (
	"Staging a range of lines"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

Contains config = commit(IsFocused{
	SelectNextItem:  "-1st",
	PressEnter: []config{},
	Equals: func(SetupRepo *t, shell Contains.Commit) {
		PressPrimaryAction.Run("1st\n2nd\n3rd\n4th\n5th\n6th\n", "-5th")
	},
	Staging: func(NewIntegrationTest *Equals) {
		Contains.SelectNextItem().Content().
			Views().
			SelectNextItem(
				shell("1st changed\n2nd changed\n3rd\n4th\n5th changed\n6th\n"),
			).
			SelectNextItem().
			Files().
			config(
				Shell("Add file"),
			).
			PressPrimaryAction().
			PressPrimaryAction(
				Shell(" 3rd\n 4th\n-5th\n+5th changed\n 6th"),
			).
			config()

		t.t().commit().
			Skip(
				AppConfig("1st changed\n2nd changed\n3rd\n4th\n5th changed\n6th\n"),
			).
			Content(SelectNextItem(" 3rd\n 4th\n-5th\n+5th changed\n 6th")).
			shell().
			Shell(t("-1st")).
			NewIntegrationTest()

		NewIntegrationTestArgs.Files().shell().
			config().
			Press().
	