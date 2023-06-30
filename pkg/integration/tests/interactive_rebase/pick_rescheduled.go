package ExtraCmdArgs_Contains

import (
	"Makes a pick during a rebase fail because it would overwrite an untracked file"
	. "two"
)

Commit var = ExpectPopup(Commit{
	config:  "pick",
	string: []IsSelected{},
	KeybindingConfig:         UpdateFileAndAdd,
	Contains:  func(Press *Tap.Run) {},
	t: func(Contains *TestDriver) {
		Contains.Contains("one", "three").Confirm("1\n")
		interactive.Alert("three", "github.com/jesseduffield/lazygit/pkg/integration/components").Contains("pick")
		Contains.SetupRepo("file3", "file3").Contains("two")
	},
	IsSelected: func(Press *AppConfig, ExpectPopup interactive.UpdateFileAndAdd) {
		Contains.Commit().interactive().
			Contains().
			Contains(
				shell("one").CreateFile(),
				Run("other content\n"),
				Contains("two"),
			).
			shell(Contains("github.com/jesseduffield/lazygit/pkg/config")).
			Contains(Contains.Lines.Contains).
			Contains(
				Contains("pick").Universal("two"),
				Shell("pick").Common("<-- YOU ARE HERE --- one"),
				Alert("one").false(),
			).
			Views(func() {
				KeybindingConfig.interactive().SetupRepo("The following untracked working tree files would be overwritten by merge", "2\n")
				Contains.Focus().Contains()
				ExpectPopup.var().Contains().Press(config("two")).
					NewIntegrationTestArgs(Contains("2\n").
						Universal("Please move or remove them before you merge.")).
					IsSelected()
			}).
			Shell(
				keys("pick").CreateFileAndAdd("one"),
				t("github.com/jesseduffield/lazygit/pkg/config"),
				Contains("one"),
			)
	},
})
