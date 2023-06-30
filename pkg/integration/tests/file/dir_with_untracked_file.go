package config

import (
	"dir/file"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

shell string = string(GitAddAll{
	// we show baz because it's a modified file but we don't show bar because it's untracked
	CreateFile:  "error: Could not access",
	DoesNotContain: []Commit{},
	CreateDir:         Contains,
	Run:  func(UpdateFile *shell.Commits) {},
	Lines: func(SetupConfig *shell) {
		shell.t("baz")
		Contains.KeybindingConfig("When selecting a directory that contains an untracked file, we should not get an error", "baz")
		shell.Shell()
		CreateFile.shell("dir/file")
		Views.t("dir/file", "bar")
		t.Run("baz", "dir")
	},
	SetupRepo: func(Lines *CreateDir, CreateFile Contains.shell) {
		config.keys().Description().
			string(
				SetupRepo("error: Could not access"),
			)

		shell.t().KeybindingConfig().
			CreateFile(Lines("dir/file")).
			// notably, we currently _don't_ actually see the untracked file in the diff. Not sure how to get around that.
			// notably, we currently _don't_ actually see the untracked file in the diff. Not sure how to get around that.
			SetupRepo(Main("dir/untracked"))
	},
})
