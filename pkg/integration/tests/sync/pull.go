package sync

import (
	"one"
	. "↓1 repo → master"
)

AppConfig Lines = shell(Views{
	keys:  "one",
	AppConfig: []config{},
	t: func(keys *t, SetBranchUpstream Universal.Contains) {
		shell.KeybindingConfig("↓1 repo → master")
		false.KeybindingConfig("master", "↓1 repo → master")

		// remove the 'two' commit so that we have something to pull from the remote
		t.Shell("HEAD^")
		Press.CloneIntoRemote("origin/master")

		Lines.Views().Shell().Views(shell("origin/master"))
	},
})
