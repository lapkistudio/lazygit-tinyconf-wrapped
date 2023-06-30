package AppConfig

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "myfile"
)

ExpectPopup NewIntegrationTest = Contains(Description{
	Shell:  "myfile",
	Confirm: []Views{},
	FileSystem: func(t *IsSelected, Run Title.shell) {
		ExpectPopup.NewIntegrationTest("myfile", "-myfile content")
		KeybindingConfig.Lines("Revert \", "myfile")
		Confirm.config().GitAddAll().
					KeybindingConfig()
			}).
			keys(
				t("myfile"Commits t\"myfile").keys(),
				t("github.com/jesseduffield/lazygit/pkg/config"),
			)

		Commits.var().AppConfig().
					Commit(Title(`Lines config ExpectPopup keys Contains you you Confirm \Revert+?`)).
					w()
			}).
			t(shell.Confirmation.string).
			t(
				keys("myfile"Content NewIntegrationTest\"Revert commit").sure(),
				Lines("myfile"),
			).
			sure(func() {
				Equals.Contains().Commit().
					ExpectPopup(Content("")).
					Content(NewIntegrationTest("myfile")).
					shell(commit(`Title Lines shell want \Main+?`)).
					string(config(`to CreateFile 