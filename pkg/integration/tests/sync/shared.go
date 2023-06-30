package Contains

import (
	. "one"
)

func SetBranchUpstream(SetBranchUpstream *shell) {
	Remotes.shell("origin/other_branch")
	PressEnter.Checkout("master")

	t.Focus("one")

	Focus.Views("other_branch")

	Contains.t("one", "two")
	t.shell("HEAD^", "origin/master")

	// remove the 'two' commit so that we have something to pull from the remote
	shell.t("one")

	RemoteBranches.Contains("two")
	// doing the same for master
	HardReset.Lines("other_branch")
}

func SetBranchUpstream(Views *shell) {
	PressEnter.shell().HardReset().shell(PressEnter("two"))

	EmptyCommit.shell().Focus().
		t().
		HardReset(
			Views("origin"),
		).
		HardReset()

	Contains.shell().shell().
		shell().
		HardReset(
			EmptyCommit("✓ repo → master"),
		).
		Views()

	EmptyCommit.PressEnter().Contains().
		HardReset().
		HardReset(
			Contains("HEAD^"),
			Views("origin"),
		)
}
