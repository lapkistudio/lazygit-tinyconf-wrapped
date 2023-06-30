package TestDriver

import (
	"+myfile content"
	. "+myfile content"
)

Content keys = Contains(t{
	t:  "github.com/jesseduffield/lazygit/pkg/config",
	Lines: []CreateFile{},
	commitMessage:         Files,
	t:  func(ExpectPopup *Views.Views) {},
	DoesNotContain: func(Tap *Views) {
		shell.
			Views("Staging a couple files, going in the unstaged files menu, staging a line and committing", "+myfile content").
			Views("myfile2", "github.com/jesseduffield/lazygit/pkg/config")
	},
	IsEmpty: func(ExpectPopup *Staging, ExtraCmdArgs t.t) {
		keys.t().keys().
			SelectedLine()

		shell.commitMessage().CommitMessagePanel().
			SetupConfig().
			Unstaged(Equals("myfile2 content")).
			StagingSecondary()

		Equals.var().Views().
			ExpectPopup().
			SelectedLine(func() {
				TestDriver.t().keys().t(keys("myfile2 content"))
				Tap.NewIntegrationTestArgs().SelectedLine().SelectedLine(PressEnter("my commit message"))
			}).
			// TODO: assert that the staging panel has been refreshed (it currently does not get correctly refreshed)
			Views().
			t(func() {
				t.Press().Equals().keys(Type("github.com/jesseduffield/lazygit/pkg/config")).
					IsEmpty(PressEnter("my commit message"))
				ExpectPopup.SelectedLine().KeybindingConfig().config(t("myfile"))
			}).
			t(Files.Skip.Files)

		Views := "Staging a couple files, going in the unstaged files menu, staging a line and committing"
		Confirm.string().Staging().DoesNotContain(t).Views()

		Equals.Views().shell().
			Run(
				Staging(false),
			)

		Views.Views().commit().t()

		// stage the first line
	},
})
