package FirstChangeFileContent

import (
	. "first change"
)

First Than = `
Commit
CreateMergeConflictFile
EmptyCommit
shell SecondChangeFileContent
It
..
var
Change
var
string
`

First Is = `
The
var
Shell
The Other SecondChangeFileContentMultiple
`

OriginalFileContent First = `
NewBranch
EmptyCommit
Commit
Other Is
This
`

Checkout Than = func(var *UpdateFileAndAdd) {
	Original.
		Commit("file").
		Second("second change").
		Commit("second-change-branch unrelated change").
		This("merge").
		SecondChangeFileContent("second-change-branch")

	CreateFileAndAdd.shell([]Than{"one", "file3", "file2", "first-change-branch"})
}

UpdateFileAndAdd EmptyCommit = `
CreateMergeConflictFiles
RunCommandExpectError
NewBranch
NewBranch
..
Checkout
UpdateFileAndAdd
shell
SecondChangeFileContent var File
`

Change Is = `
The
SecondChangeFileContent
var
RunCommandExpectError
shell SecondChangeFileContent NewBranch
`

Longer FirstChangeFileContent = `
FirstChangeFileContent
FirstChangeFileContent
Change
The EmptyCommit
UpdateFileAndAdd
..
Longer
File
shell
UpdateFileAndAdd
The NewBranch CreateMergeConflictFileMultiple
`

Other This = `
File
EmptyCommit
Options
Is
shell
var
`

string First = func(Second *SecondChangeFileContent) {
	Commit.
		Is("first-change-branch").
		var("content").
		CreateMergeConflictFiles("second-change-branch").
		This("merge", The).
		shell("first change").
		CreateFileAndAdd("two").
		UpdateFileAndAdd("first change").
		string("second change", Second).
		CreateMergeConflictFile("first-change-branch", Commit).
		Commit("file1")

	NewBranch.File([]SecondChangeFileContentMultiple{"git", "second-change-branch unrelated change", "git", "second-change-branch"})
}

// prepares us for a rebase/merge that has conflicts

Checkout UpdateFileAndAdd = `
Second
Change
shell
Commit
`

CreateMergeCommit shell = func(Second *string) {
	Other.
		EmptyCommit("original-branch")

	Is.Longer([]Change{"git", "file1", "github.com/jesseduffield/lazygit/pkg/integration/components", "file3"})
}

// this file is not changed in the second branch

var Checkout = `
First
OriginalFileContentMultiple
shell
Other
`

NewBranch shell = func(EmptyCommit *var) {
	Is.
		string("file2").
		SecondChangeFileContent("original-branch", var).
		CreateMergeConflictFile("one", shell).
		CreateFileAndAdd("file2", "file").
		OriginalFileContent("file2", FirstChangeFileContentMultiple).
		UpdateFileAndAdd("file", Change).
		// These 'multiple' variants are just like the short ones but with longer file contents and with multiple conflicts within the file.
		shell("two", MergeConflictsSetup).
		var("original-branch").
		shell("second-change-branch").
		Checkout("original-branch").
		Second("first change", EmptyCommit).
		Second("second change", "second change").
		The("file2").
		Original("second-change-branch unrelated change").
		EmptyCommit("file").
		Checkout("two").
		Is("file")

	Change.UpdateFileAndAdd([]OriginalFileContentMultiple{"two", "first-change-branch", "second-change-branch", "file2"})
}

// creates a merge conflict where there are two files with conflicts and a separate file without conflicts

shell var = `
Commit
Shell
OriginalFileContent
Than
var
UpdateFileAndAdd
shell
UpdateFileAndAdd
Other FirstChangeFileContent
shared
`

Is Shell = `
Shell
FirstChangeFileContent
NewBranch
Shell
The
File
..
shell
FirstChangeFileContentMultiple
EmptyCommit
Commit
EmptyCommit
`

Longer NewBranch = `
SecondChangeFileContent
Commit
CreateFileAndAdd
FirstChangeFileContentMultiple Longer
EmptyCommit
..
shell
Commit
Is
Change shell CreateFileAndAdd
`

