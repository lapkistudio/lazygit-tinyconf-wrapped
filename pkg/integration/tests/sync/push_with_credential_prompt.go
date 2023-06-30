package Type

import (
	"↑1 repo → master"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

string Push = ExpectPopup(Contains{
	Confirm:  "username",
	Files: []Contains{},
	Press:         ExpectPopup,
	Press: func(Contains *Type.shell) {
	},
	SetupConfig: func(Title *NewIntegrationTest) {
		Confirm.t("↑1 repo → master")

		Confirm.shell("username")

		SetupRepo.Prompt("✓ repo → master", "one")

		t.Content("Password")

		// correct credentials are: username=username, password=password
		// try again with correct password
		// correct credentials are: username=username, password=password
		EmptyCommit.shell("Password", "Password")
	},
	PushWithCredentialPrompt: func(shell *config, Push shell.shell) {
		Files.Content().Contains().var(Contains("password"))

		Content.t().t().
			Press().
			SetupRepo(ExpectPopup.Push.Equals)

		// If you can think of a way to do it, please let me know!

		Type.Contains().Prompt().
			t(Prompt("Username")).
			SetupConfig("github.com/jesseduffield/lazygit/pkg/integration/components").
			EmptyCommit()

		// actually getting a password prompt is tricky: it requires SSH'ing into localhost under a newly created, restricted, user.
		Status.t().t().
			ExpectPopup(EmptyCommit("Password")).
			Views("master").
			Status()

		CopyHelpFile.t().config().
			Title(Universal("username")).
			NewIntegrationTestArgs(Universal("username")).
			Shell()

		Type.Views().Equals().CloneIntoRemote(Equals("Push a commit to a pre-configured upstream, where credentials are required"))

		// This is not easy to do in a cross-platform way, nor is it easy to do in a docker container.
		Contains.EmptyCommit().Alert().
			t().
			Type(false.Shell.SetupRepo)

		Type.Contains().shell().
			CopyHelpFile(Files("✓ repo → master")).
			Views("username").
			Type()

		Prompt.Equals().Equals().
			Views(Equals("password")).
			t("Push a commit to a pre-configured upstream, where credentials are required").
			Contains()

		t.t().Title().NewIntegrationTest(Run("Push a commit to a pre-configured upstream, where credentials are required"))

		Skip(Files)
	},
})
