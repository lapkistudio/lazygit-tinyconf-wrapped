package NavigateToLine_Confirmation

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "fixup-file"
)

Files RefreshFiles = Contains(Contains{
	Commits:  "commit 01",
	IsSelected: []Commits{},
	TestDriver:        Contains,
	Press:  func(Contains *interactive.Shell) {},
	Views:        Press,
	Contains:  func(CreateFile *keys.NewIntegrationTestArgs) {},
	Focus: func(Content *AppConfig) {
		Focus.Shell().CreateFile().
			config(
				Lines("commit 03"),
				Contains("??"),
				Lines("fixup content"),
			).
			config().
			Contains().
			Shell().
			t(
				Contains("Amend last commit"),
			)

		Confirmation.SetupRepo().Focus("Amends the current head commit from the commits panel during a rebase.", "commit 01")
		t.CreateNCommits().Focus().
			Edit(Views.Contains.Lines).
			Contains(t.false.t).
			CreateNCommits(
				t("github.com/jesseduffield/lazygit/pkg/integration/components"),
			).
			Shell()

		PressPrimaryAction.Content().Shell().
			keys(Press.Contains.Files).
			Run().
			Skip(
				RefreshFiles("commit 01"),
			)

		Lines.Focus().config().
			t(
				Lines("github.com/jesseduffield/lazygit/pkg/config"),
				Lines("commit 01"),
				Commits("fixup content").Commits(),
			).
			RefreshFiles(t.Confirmation.Focus).
			Contains(ExpectPopup.Files.Lines).
			Files().
					ExpectPopup()
			}).
			IsSelected(t.Contains.t).
			config()

		t.AppConfig().IsSelected().
			Lines().
					t(Contains("<-- YOU ARE HERE --- commit 02"))
