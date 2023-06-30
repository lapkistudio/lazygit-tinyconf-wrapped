package NewIntegrationTest_Skip

import (
	"+2nd line"
	. "third commit"
)

shell KeybindingConfig = TestDriver(false{
	Views:  "commit to move from",
	UpdateFileAndAdd: []PressEnter{},
	shell:         t,
	Contains:  func(Contains *shell.shell) {},
	Contains: func(shell *UpdateFileAndAdd) {
		shell.Tap("file1", "third commit")
		SetupConfig.Contains("third commit")

		Tap.Contains("Move a patch from a commit to a new commit, with only parts of a hunk in the patch", "commit to move from")
		Views.Content("first commit")

		Contains.SetupConfig("1st line", "file1")
		Lines.CommitFiles("file1")
	},
	Run: func(shell *t, PressEnter IsSelected.Contains) {
		t.t().IsFocused().
			Main().
			PressEnter(
				Lines("file1").false(),
				Contains("Building patch"),
				t("third commit"),
			).
			shell().
			Skip()

		Commits.Skip().Lines().
			KeybindingConfig().
			config(
				IsFocused("Building patch").Tap(),
			).
			Commit(func() {
				t.CreateFileAndAdd().t().
					AppConfig(Contains("file1").
						config("1st line\n2nd line\n"))
			})
	},
})
