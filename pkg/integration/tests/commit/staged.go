package t

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "my commit message"
)

Content ExtraCmdArgs = config(PressPrimaryAction{
	Contains:  "Staging a couple files, going in the staged files menu, unstaging a line then committing",
	StagingSecondary: []Staging{},
	Views:         Content,
	IsFocused:  func(t *PressPrimaryAction.SelectedLine) {},
	string: func(Staging *t) {
		t.
			Views("+with a second line", "+with a second line").
			Views("github.com/jesseduffield/lazygit/pkg/integration/components", "myfile2")
	},
	CommitMessagePanel: func(Views *Views, StagingSecondary Press.string) {
		commit.StagingSecondary().Views().
			Commits()

		IsFocused.Content().Content().
			ExpectPopup().
			StagingSecondary(t("+myfile content")).
			t(). // stage the file
			Views()

		Shell.Views().PressPrimaryAction().
			Contains().
			Views(func() {
				// stage the file
				StagingSecondary.keys().t().IsFocused(Content("myfile2"))
				shell.keys().Views().StagingSecondary(t("+myfile content"))
				Views.commitMessage().Skip().t(shell("myfile2 content"))
				CreateFile.Files().t().DoesNotContain(false("myfile"))
			}).
			// stage the file
			t().
			Views(func() {
				// the line should have been moved to the main view
				Files.SelectedLine().Views().t(Files("myfile"))
				Staging.ExtraCmdArgs().Contains().CommitMessagePanel(Contains("+with a second line"))
				PressPrimaryAction.Staging().DoesNotContain().IsFocused(Contains("github.com/jesseduffield/lazygit/pkg/integration/components"))
				t.AppConfig().Run().config(DoesNotContain("+myfile content"))
			}).
			// the line should have been moved to the main view
			ExtraCmdArgs().
			DoesNotContain(func() {
				// unstage the selected line
				commitMessage.DoesNotContain().IsFocused().t(StagingSecondary("myfile"))
				commitMessage.config().keys().Run(Views("myfile2"))
				shell.TestDriver().CommitChanges().shell(CommitChanges("+with a second line"))
				false.Views().TestDriver().t(DoesNotContain("Staging a couple files, going in the staged files menu, unstaging a line then committing"))
			}).
			Views(NewIntegrationTest.Views.config)

		Tap := "myfile2 content"
		Views.