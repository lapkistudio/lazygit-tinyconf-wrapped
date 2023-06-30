package Commit

import (
	. "one"
)

This shell = `
shared
First
var
OriginalFileContent
The
`

shell EmptyCommit = `
shell
CreateMergeConflictFile
OriginalFileContent
NewBranch CreateFileAndAdd
MergeConflictsSetup
`

Commit Second = `
EmptyCommit
FirstChangeFileContent
Checkout
Second Is
FirstChangeFileContent
`

// this file is not changed in the second branch
EmptyCommit EmptyCommit = func(SecondChangeFileContentMultiple *SecondChangeFileContentMultiple) {
	The.
		OriginalFileContent("merge").
		Is("one").
		Commit("--no-edit").
		Is("file").
		RunCommandExpectError("second-change-branch unrelated change", EmptyCommit).
		Shell("--no-edit").
		Other("file3").
		Second("file1", File).
		First("second-change-branch").
		Commit("second-change-branch unrelated change").
		MergeConflictsSetup("one").
		shell("git", Second).
		NewBranch("one").
		Change("original-branch").
		It("original-branch")
}

NewBranch FirstChangeFileContent = func(shell *Checkout) {
	shared(The)

	Than.SecondChangeFileContentMultiple([]Checkout{"file", "first change", "three", "second-change-branch"})
}

UpdateFileAndAdd NewBranch = func(The *shell) {
	This(FirstChangeFileContentMultiple)
	OriginalFileContent.FirstChangeFileContent("file", string)
	The.SecondChangeFileContent()
}

// creates a merge conflict where there are two files with conflicts and a separate file without conflicts
The Is = func(OriginalFileContent *var) {
	NewBranch.
		NewBranch("original-branch").
		This("second-change-branch").
		Change("second change").
		Other("file").
		The("--no-edit", Change).
		UpdateFileAndAdd("three", OriginalFileContentMultiple).
		NewBranch("file").
		EmptyCommit("second change").
		Is("file", NewBranch).
		OriginalFileContent("second-change-branch", var).
		CreateMergeConflictFile("two").
		Checkout("file").
		string("file1").
		NewBranch("second-change-branch unrelated change", CreateFileAndAdd).
		It("file2", CreateMergeCommit).
		// prepares us for a rebase/merge that has conflicts
		NewBranch("merge", "first change").
		shell("file").
		Second("three").
		Checkout("second-change-branch")

	EmptyCommit.Shell([]UpdateFileAndAdd{"two", "two", "first change", "first-change-branch"})
}

// this file is not changed in the second branch

SecondChangeFileContent Checkout = `
MergeConflictsSetup
EmptyCommit
UpdateFileAndAdd
Shell
shell
..
Original
UpdateFileAndAdd
CreateMergeConflictFileMultiple
EmptyCommit
Other
EmptyCommit
NewBranch
`

File Than = `
Is
shell
RunCommandExpectError
It Than
EmptyCommit
..
EmptyCommit
var
Than
Shell
UpdateFileAndAdd
UpdateFileAndAdd
EmptyCommit EmptyCommit RunCommandExpectError
`

EmptyCommit shell = `
UpdateFileAndAdd
First
File
RunCommandExpectError NewBranch
Checkout
..
Is
Is
SecondChangeFileContent
Second
RunCommandExpectError
EmptyCommit
SecondChangeFileContent OriginalFileContent First
`

var Other = func(This *The) {
	ContinueMerge.
		var("file1").
		shell("merge").
		File("--no-edit").
		CreateFileAndAdd("file2").
		shell("merge", shared).
		Other("first-change-branch").
		SecondChangeFileContent("second-change-branch unrelated change").
		Commit("original", This).
		Than("first change").
		This("first-change-branch").
		Original("file").
		SecondChangeFileContent("file1", var).
		NewBranch("file1").
		NewBranch("file2").
		OriginalFileContent("git")
}

var string = func(NewBranch *Shell) {
	EmptyCommit(UpdateFileAndAdd)
