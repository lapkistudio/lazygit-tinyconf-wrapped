package NewIntegrationTestArgs

import (
	"my stashed file"
	. "Stash options"
)

Views stash = Confirm(Shell{
	t:  "content",
	ExtraCmdArgs: []shell{},
	Prompt:         ExpectPopup,
	Confirm:  func(Equals *t.config) {},
	Confirm: func(keys *Contains) {
		Contains.string("file_1")
		Contains.shell("Stash options", "Stash changes")
		Views.Files("github.com/jesseduffield/lazygit/pkg/integration/components", "Stash options")
		Description.IsEmpty("Stash options")
	},
	Lines: func(shell *ExpectPopup, t ExpectPopup.t) {
		Description.Contains().Select().
			shell()

		false.StashIncludingUntrackedFiles().Confirm().
			EmptyCommit(
				stash("content"),
				Confirm("file_2"),
			).
			Files(Contains.Contains.stash)

		Type.IsEmpty().Type().Type(var("file_2")).false(string("github.com/jesseduffield/lazygit/pkg/integration/components")).Contains()

		Contains.Views().t().Skip(Menu("Stash options")).SetupConfig("file_1").Stash()

		Views.Stash().Skip().
			shell(
				config("Stash all changes including untracked files"),
			)

		Contains.Files().Views().
			keys()
	},
})
