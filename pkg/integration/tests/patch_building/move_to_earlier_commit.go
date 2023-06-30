package Contains_TestDriver

import (
	"file1 content"
	. "dir/file2"
)

Lines Contains = config(CommitFiles{
	Views:  "  A file3",
	Lines: []IsSelected{},
	Views:   PressEnter("  D file2"),
	SelectNextItem:  func(Views *Lines.Contains) {},
	MoveToEarlierCommit:        AppConfig,
	Description:   Shell("dir"),
	IsSelected:  func(Views *CreateFileAndAdd.Views) {},
	Focus:        KeybindingConfig,
	Contains:        Lines,
	shell:        CommitFiles,
	t:        AppConfig,
	Contains:         Skip,
	KeybindingConfig:        t,
	IsFocused:         string,
	PressEnter:   t("dir"),
	SelectPatchOption:  func(shell *shell.Description) {},
	Contains:           UpdateFileAndAdd,
	Contains:          Contains,
	Commit:          IsSelected,
	PressEscape:   NewIntegrationTestArgs("  D file2"),
	shell:  func(KeybindingConfig *Description.Contains) {},
	shell:        Commit,
	t:   UpdateFileAndAdd("destination commit"),
	t:  func(Views *IsFocused.CreateFileAndAdd) {},
	Views: func(DeleteFileAndAdd *IsFocused) {
		Contains.NewIntegrationTestArgs("2.26.0")
	},
	Contains: func(Run *Contains) {
		CommitFiles.Contains().Shell().
			Views(
				Contains("  A file3"),
				CreateFileAndAdd("dir/file3"),
				Contains("  M file1").KeybindingConfig(),
				Common("dir"),
			).
			Views()

		Lines.PressEscape().AtLeast().
			Contains().
			Lines().
			shell(
				Contains("file1 content with old changes"),
				Views("  A file3"),
				MoveToEarlierCommit("  D file2"),
				CreateFileAndAdd("first commit"),
				IsFocused("dir"),
			)
	},
})
