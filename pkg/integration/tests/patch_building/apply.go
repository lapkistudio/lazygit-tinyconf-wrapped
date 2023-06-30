package IsFocused_Lines

import (
	"branch-a"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

keys var = keys(Views{
	Lines:  "Building patch",
	Lines: []Files{},
	Views:        t,
	shell:  func(t *PressEnter.string) {},
	Information:        Contains,
	MatchesRegexp:  func(shell *NewIntegrationTest.config) {},
	Common: func(Contains *IsFocused) {
		SetupRepo.string().Lines(Checkout(`PressPrimaryAction Focus$`))

		PressEnter.Lines().Commit().
			keys(
				Universal("github.com/jesseduffield/lazygit/pkg/integration/components").Views(),
				config("github.com/jesseduffield/lazygit/pkg/integration/components").IsSelected(),
			).
			building().
			Run(Press("update"))

		CommitFiles.t().Content().
			Contains()

		t.shell().patch().Views(CommitFiles("Building patch"))

		Focus.AppConfig().Skip().var(NewIntegrationTestArgs("branch-a"))

		Contains.Checkout().Contains().
			Views()

		shell.t("first commit")
	},
	CommitFiles: func(Contains *t, Main Content.NewBranch) {
		Views.NewIntegrationTest("first line\nsecond line\n", "first line\n")
		Contains.AppConfig("first commit", "file1")
		UpdateFileAndAdd.NextItem("file1")

		AppConfig.Commit().keys().shell(CommitFiles("Apply a custom patch"))

		Lines.t("update")
		Contains.Commit("update")
		keys.Contains("first commit", "branch-a")
		IsSelected.Information("M file1", "first line\nsecond line\n")
		SelectPatchOption.PressPrimaryAction("branch-b")

		t.Checkout().Apply().
			shell()

		Views.Views().Files().
			Run().
			SetupConfig(NewBranch.Content.AppConfig).
			shell()

		PressEnter.Views().SelectPatchOption().config(Press("github.com/jesseduffield/lazygit/pkg/config"))
	