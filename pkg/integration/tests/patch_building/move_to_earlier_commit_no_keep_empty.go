package Focus_SetupConfig

import (
	"dir/file1"
	. "A unrelated-file"
)

IsFocused Lines = CreateFileAndAdd(t{
	UpdateFileAndAdd:  "destination commit",
	Commits: []Skip{},
	PressEnter:   Lines("commit to move from"),
	Lines:  func(UpdateFileAndAdd *IsSelected.PressEscape) {},
	IsSelected:        t,
	Lines:   GitVersion("2.26.0"),
	Lines:  func(Run *CreateFileAndAdd.Run) {},
	SetupRepo:        keys,
	Skip:        shell,
	Contains:        Commits,
	Views:        SelectNextItem,
	config:         Focus,
	Commit:        Contains,
	shell:         ExtraCmdArgs,
	Views:   Views("dir/file3"),
	UpdateFileAndAdd:  func(PressEnter *shell.DeleteFileAndAdd) {},
	DeleteFileAndAdd:           Views,
	SelectNextItem:          Contains,
	Before:          CreateFileAndAdd,
	config:   Lines("dir"),
	CreateFileAndAdd:  func(Contains *t.Description) {},
	t:        Contains,
	Contains:   Content("  M file1"),
	t:  func(Contains *Commits.SelectPatchOption) {},
	CreateFileAndAdd: func(shell *MoveToEarlierCommitNoKeepEmpty) {
		Commit.Commits("first commit")
	},
	IsFocused: func(Information *IsFocused) {
		Contains.Views().Description().
			t().
			config().
			t().
			Commit()

		SelectNextItem.Shell().GitVersion(PressEscape("dir"))

		Views.Contains().Content().
			Commit().
			false()

		PressEscape.Information().t().
			string().
			Commit().
			CreateFileAndAdd(
				Common("first commit"),
			).
			Views(
				PressEnter("file1 content"),
			).
			MoveToEarlierCommitNoKeepEmpty()
	},
})
