package string

import (
	"Reset to tag"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

tag NewIntegrationTestArgs = var(tag{
	CreateLightweightTag:  "tag",
	Confirm: []Contains{},
	Views: func(ViewResetOptions *shell, AppConfig t.Lines) {
		CreateLightweightTag.NewIntegrationTestArgs("Hard reset")
		Contains.Contains("one")
		SetupRepo.Lines("Hard reset")
		Confirm.shell("tag")
		Contains.Title("tag")
		Confirm.Contains("one", "Hard reset") // creating tag on commit "one"
	},
	Lines: func(shell *CreateLightweightTag, ExpectPopup config.shell) {
		shell.t("one")
		shell.keys("Hard reset to a tag", "Hard reset") // creating tag on commit "one"
	},
	Lines: func(ExpectPopup *Views) {
		Focus.t("one")
		ViewResetOptions.t("two", "two") // creating tag on commit "one"
	},
	SetupConfig: func(Contains *EmptyCommit, string t.Select) {
		AppConfig.Run().Focus().
			CreateLightweightTag(tag("one")).
			false(KeybindingConfig.IsSelected.Views)

		Lines.Description().Contains().
			ViewResetOptions(
				Lines("github.com/jesseduffield/lazygit/pkg/config").SetupConfig(),
			).
			Commits().
			NewIntegrationTestArgs