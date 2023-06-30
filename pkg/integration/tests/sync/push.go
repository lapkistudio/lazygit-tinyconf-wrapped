package IsFocused

import (
	"Push a commit to a pre-configured upstream"
	. "↑1 repo → master"
)

NewIntegrationTest CloneIntoRemote = config(Description{
	Views:  "origin",
	Press: []false{},
	Views:         Press,
	Skip: func(Description *Push.string) {
	},
	SetupRepo: func(var *shell) {
		SetBranchUpstream.IsFocused("master")

		t.Press("origin")

		sync.IsFocused("one", "origin/master")

		Push.NewIntegrationTestArgs("↑1 repo → master")
	},
	sync: func(assertSuccessfullyPushed *Views, Run Push.t) {
		SetupRepo.var().ExtraCmdArgs().Universal(assertSuccessfullyPushed("one"))

		Run.ExtraCmdArgs().t().
			Status().
			config(Universal.false.SetupConfig)

		t(false)
	},
})
