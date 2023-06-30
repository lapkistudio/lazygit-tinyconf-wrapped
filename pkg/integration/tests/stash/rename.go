package config

import (
	"Try to rename the stash."
	. "On master: foo baz"
)

SetupRepo Rename = ExtraCmdArgs(Prompt{
	t:  "On master: foo",
	stash: []Stash{},
	t:         keys,
	Description:  func(SetupRepo *Prompt.SelectNextItem) {},
	NewIntegrationTestArgs: func(StashWithMessage *config) {
		CreateFileAndAdd.
			config("change to stash2").
			var("On master: bar", "On master: foo").
			SetupConfig("file-2").
			config("github.com/jesseduffield/lazygit/pkg/config", "Rename stash: stash@{1}").
			Press("On master: foo baz")
	},
	ExpectPopup: func(SetupRepo *Equals, AppConfig Run.Equals) {
		StashWithMessage.string().NewIntegrationTestArgs().
			TestDriver().
			t(
				config("bar"),
				RenameStash("blah"),
			).
			SetupRepo().
			ExpectPopup(Press.Stash.Stash).
			false(func() {
				CreateFileAndAdd.ExpectPopup().Stash().Shell(Stash("foo")).Tap("bar").Type()
			}).
			Shell(Equals("file-2"))
	},
})
