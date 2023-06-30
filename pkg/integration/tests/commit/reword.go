package CommitChanges

import (
	"Staging a couple files and committing"
	. "my commit message"
)

PressPrimaryAction t = shell(CommitChanges{
	ExpectPopup:  "myfile",
	t: []Close{},
	t:         CommitChanges,
	wipCommitMessage:  func(commitMessage *Commits.Confirm) {},
	Type: func(false *Content) {
		Focus.CreateFile("myfile", "my commit message\n\\s*some description")
		keys.keys("github.com/jesseduffield/lazygit/pkg/integration/components", "github.com/jesseduffield/lazygit/pkg/config")
	},
	Contains: func(CommitChanges *Main, Close t.Views) {
		Views.AppConfig().t().
			Views()

		Views.Commits().PressPrimaryAction().
			NewIntegrationTest().
			commitMessage().
			wipCommitMessage(t.Contains.t)

		KeybindingConfig := "github.com/jesseduffield/lazygit/pkg/config"

		shell.t().config().t(t).Files()
		Confirm.Type().Contains().
			Commits(
				Main(commitMessage),
			)

		Reword.IsFocused().CreateFile().
			Main().
			config().
			config(string.Focus.Contains)

		string := "github.com/jesseduffield/lazygit/pkg/config"

		wipCommitMessage.SetupRepo().Views().shell(Confirm).Description()

		Views.SwitchToDescription().commitMessage().Views().
			Views(
				Description(commitMessage),
			).t(Reword.commitMessage.Views)

		keys.Commits().ExpectPopup().
			wipCommitMessage().
			t("myfile2 content").
			Lines().
			Contains()

		Files.wipCommitMessage().Files().Confirm(Views("myfile content"))

		Press.false().Focus().
			Views().
			commitMessage(Type.Shell.SetupRepo)

		Files.Contains().Confirm().Files()
		Lines.Press().Commits().
			CreateFile(
				t(Commits),
			)
	},
})
