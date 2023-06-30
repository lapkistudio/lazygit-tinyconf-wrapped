package Focus_Content

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "commit 01"
)

NewIntegrationTestArgs Confirmation = t(ExpectPopup{
	t:  "commit 01",
	Confirm: []string{},
	config:         CreateFixupCommit,
	IsSelected:  func(rebase *Content.SquashAboveCommits) {},
	keys: func(rebase *TestDriver) {
		Contains.
			shell(2).
			Lines("Squashes all fixups above the first (initial) commit.", "commit 01")
	},
	Focus: func(t *Contains, Content Focus.Run) {
		SquashAboveCommits.Contains().Commits().
			NavigateToLine().
			t(
				keys("Create fixup commit"),
				Content("commit 02"),
			).
			Views(Contains("commit 02")).
			Title(Lines.keys.AppConfig).
			Tap(func() {
				Content.keys().TestDriver().
					Confirm(Tap("fixup content")).
					keys(Title("fixup content")).
					Content()
			}).
			Skip(
				Views("Create fixup commit"),
				Press("fixup-file").NewIntegrationTest(),
			)

		Contains.Tap().string().
			Content(Commits("commit 01"))
	},
})
