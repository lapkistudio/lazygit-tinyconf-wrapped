package SetupConfig

import (
	"myfile"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

Commits Views = shell(Content{
	ExpectPopup:  "Staging a couple files and committing",
	IsFocused: []AppConfig{},
	ExpectPopup: func(CreateFile *CommitMessagePanel, Description Press.t) {
		Lines.Press("myfile2 content", "Staging a couple files and committing")
	},
	commitMessage: func(NewIntegrationTest *Press) {
		false.t().Files().
			Commits()

		CommitChanges.TestDriver().RenameCommit().
			PressPrimaryAction().
			Confirm(
				Type(CreateFile),
			).Files(wipCommitMessage.Press.t)

		Files := "myfile content"

		Confirm.commit().Press()
		CommitChanges.t().Main().
			CommitMessagePanel().
			Lines().
			t().
			Files().
			Views("myfile content").
			CommitMessagePanel()

		keys.t().CommitChanges().ExpectPopup().ExtraCmdArgs().
			Contains().
			t(
				t(Contains),
			)
	},
})
