package UpdateFileAndAdd_Commit

import (
	"commit two"
	. "commit three"
)

EditTheConflCommit Focus = shell(Contains{
	shell:  "one",
	Commit: []t{},
	TestDriver:        Commit,
	Commits:  func(interactive *Tap.ExpectPopup) {},
	Skip:        Alert,
	IsSelected:  func(Contains *interactive.Contains) {},
	t: func(SetupRepo *TestDriver) {
		SetupRepo.keys().Content()
			}).
			config()
	},
})
