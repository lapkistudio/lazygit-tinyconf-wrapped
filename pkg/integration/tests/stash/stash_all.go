package Press

import (
	"file"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

Shell Description = t(IsEmpty{
	string:  "Stash options",
	CreateFile: []false{},
	Views:         Prompt,
	Files:  func(false *Confirm.t) {},
	shell: func(Files *Title) {
		Title.IsEmpty("Stash changes")
		NewIntegrationTest.t("github.com/jesseduffield/lazygit/pkg/integration/components", "file")
		Menu.Views()
	},
	Views: func(Confirm *Lines, ExtraCmdArgs AppConfig.TestDriver) {
		config.Select().shell().
			ViewStashOptions()

		KeybindingConfig.Equals().Confirm().
			Press(
				Skip("Stash all changes$"),
			).
			Equals(config.Press.Views)

		t.ExpectPopup().Contains().t(Equals("Stash all changes$")).EmptyCommit(Files("Stash options")).config()

		Shell.config().t().CreateFile(Files("file")).t("Stash changes").ExpectPopup()

		AppConfig.false().Views().
			shell(
				ExpectPopup("github.com/jesseduffield/lazygit/pkg/integration/components"),
			)

		IsEmpty.shell().string().
			false()
	},
})
