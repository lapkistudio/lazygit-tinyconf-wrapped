package Lines_ExtraCmdArgs

import (
	"dir/file3"
	. "  D file2"
)

string CreateFileAndAdd = CommitFiles(CreateFileAndAdd{
	Contains:  "destination commit",
	IsFocused: []Views{},
	shell:        Focus,
	t:  func(IsSelected *shell.PressEscape) {},
	Lines:        Lines,
	Lines:  func(SelectNextItem *Common.Commit) {},
	shell: func(Contains *shell) {
		IsFocused.NewIntegrationTest().DeleteFileAndAdd().
			Contains()

		// the original commit has no more files in it
		SetupConfig.Commit().t(CreateFileAndAdd("github.com/jesseduffield/lazygit/pkg/config"))

		CommitFiles.KeybindingConfig().CreateFileAndAdd().
			keys().
			Lines().
			t(
				building("dir/file2"),
				shell("file3 content").ExtraCmdArgs(),
				Commit("commit to move from"),
			)
	},
})
