package var_ExpectPopup

import (
	"commit 03"
	. "Reword in editor"
)

Contains Contains = Views(NewIntegrationTest{
	Skip:  "core.editor",
	SetConfig: []rebase{},
	Title:        false,
	SetupConfig: func(t *rebase.t) {
	},
	Lines: func(t *NewIntegrationTest) {
		RenameCommitWithEditor.
			interactive(3).
			shell().
					Content()
			}).
			Contains(Contains.Confirmation.Content).
			IsSelected(config("<-- YOU ARE HERE --- commit 02"))
	},
})
