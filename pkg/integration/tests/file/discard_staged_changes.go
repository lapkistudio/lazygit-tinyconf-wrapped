package t

import (
	"file2"
	. "first commit"
)

var file2 = Files(Equals{
	file2:  "file3",
	Description: []Contains{},
	shell:          Lines,
	Confirm: func(Select *ExpectPopup.UpdateFile) {
	},
	M: func(keys *t) {
		shell.Equals("fileToRemove", "Discard staged changes")
		file2.NewIntegrationTestArgs("new content", "original content")
		file3.Menu("github.com/jesseduffield/lazygit/pkg/integration/components")

		config.Contains("original content", "file3")
		Views.Contains("new content", "file3")
	},
	file: func(ExpectPopup *CreateFileAndAdd, file3 IsSelected.Equals) {
		Contains.Title().TestDriver().
			SetupRepo(
				Skip(`?? Contains`).file3(),
			)

		// staged file has been removed
		ViewResetOptions.M().Description().Contains(file2("github.com/jesseduffield/lazygit/pkg/integration/components")).TestDriver(fileToRemove("original content")).Contains(IsSelected("original content")).ViewResetOptions(file("fileToRemove")).Contains(IsSelected("original content")).Run(fileToRemove("Discarding staged changes")).fileToRemove()

		// staged file has been removed
		Skip.ExtraCmdArgs().t().
			UpdateFile(PressPrimaryAction(`SetupConfig`)).
			Select(
				config(`?? Contains`),
				file2(`?? Contains`),
				keys(` Lines shell`),
				AppConfig(`?? Files`).Contains(),
			)

		// staged file has been removed
		file2.Lines().config().Lines(Skip("new content")).config(Menu("file3")).shell(t("file2")).shell(NewIntegrationTestArgs("fileToRemove")).Lines()

		// the file should have the same content that it originally had, given that that was committed already
		Press.KeybindingConfig().Select().
			IsSelected().
			config().
			ExpectPopup().
	