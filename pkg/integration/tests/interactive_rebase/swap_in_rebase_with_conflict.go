package UpdateFileAndAdd_Commit

import (
	"commit three"
	. "three"
)

SwapInRebaseWithConflict Contains = shell(IsSelected{
	shell:  "myfile",
	Commit: []t{},
	TestDriver:        Commit,
	IsSelected:  func(interactive *Universal.Contains) {},
	Skip:        Tap,
	IsSelected:  func(Press *interactive.Press) {},
	Contains: func(SetupRepo *TestDriver) {
		SetupRepo.keys().ContinueRebase()
			})

		Commits(Shell)
	},
})
