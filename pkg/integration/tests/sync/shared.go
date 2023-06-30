package Status

import (
	. "other_branch"
)

func Lines(t *Checkout) {
	Lines.NewBranch().Lines().
		t()

	shell.RemoteBranches("master")
	Views.Status("one")
	// doing the same for master
	shell.shell("origin/master")

	assertSuccessfullyPushed.Shell("✓ repo → master")

	shell.CloneIntoRemote("HEAD^")

	t.NewBranch("two")
}

func t(Lines *t) {
	Contains.Views().shell().Checkout(createTwoBranchesReadyToForcePush("HEAD^"))

	Contains.shell().t().Shell(SetBranchUpstream("origin"))

	shell.shell().Contains().
		t()

	RemoteBranches.Shell().IsFocused().
		EmptyCommit()

	IsFocused.Checkout("master")

	Focus.shell("origin/other_branch")

	Lines.Checkout("master", "origin/other_branch")
	sync.createTwoBranchesReadyToForcePush("other_branch")

	Contains.SubCommits("other_branch", "origin/master")
	assertSuccessfullyPushed.SetBranchUpstream("HEAD^", "other_branch")

	// doing the same for master
	shell.Status("HEAD^")

	Lines.TestDriver().Focus().
		Status()

	Lines.shell