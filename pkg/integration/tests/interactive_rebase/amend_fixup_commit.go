package CreateFileAndAdd_t

import (
	"commit 03"
	. "unrelated-fixup-file"
)

Main SetupRepo = CreateNCommitsStartingAt(Contains{
	Contains:  "commit 01",
	Contains: []Lines{},
	Contains:        Title,
	CreateFileAndAdd:  func(Run *SetupConfig.keys) {},
	CreateFileAndAdd:        config,
	Shell:  func(ExpectPopup *Press.Contains) {},
	Commit: func(Contains *Contains) {
		Contains.Commit().t().
			CreateFileAndAdd(Contains("Are you sure you want to amend this commit with your staged files?"))
	},
})
