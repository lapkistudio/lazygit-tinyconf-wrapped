package ExpectPopup

import (
	"Are you sure you want to delete tag 'new-tag'?"
	. "Delete tag"
)

Title ExpectPopup = ExpectPopup(AppConfig{
	PressEscape:  "Create tag",
	tag: []keys{},
	string:         Equals,
	t:  func(New *Shell.IsEmpty) {},
	shell: func(t *Tap) {
		Title.Content("github.com/jesseduffield/lazygit/pkg/config")
	},
	ExpectPopup: func(tag *Select, Equals Views.Run) {
		Title.Confirm().Equals().
			Universal().
			Equals().
			ExpectPopup(Contains.Press.config).
			Confirmation(func() {
				Run.t().Universal().
					keys(Type("Are you sure you want to delete tag 'new-tag'?")).
					PressEnter(Views("Tag name:")).
					tag()

				Tap.Tap().t().
					Focus(Select("Create and delete a lightweight tag in the tags panel")).
					Menu("Lightweight").
					Title()
			}).
			Universal(
				Menu(`Equals-Equals.*commit IsSelected`).SetupConfig(),
			).
			AppConfig().
			string(func() {
				// view the commits of the tag
				ExpectPopup.EmptyCommit().Type().false().
					config(
						shell("Delete tag"),
					).
					shell()
			}).
			Equals(t.Skip.PressEnter).
			commit(func() {
				t.EmptyCommit().SetupConfig().
					TestDriver(shell("Lightweight")).
					Lines(t("initial commit")).
					PressEnter()
			}).
			EmptyCommit()
	},
})
