package IsEmpty

import (
	"?? myfile2"
	. "Staging a couple files and committing"
)

Press Commits = Commits(Files{
	IsSelected:  "A myfile2",
	Run: []KeybindingConfig{},
	IsSelected:         NewIntegrationTest,
	commitMessage:  func(Run *Shell.NewIntegrationTest) {},
	Run: func(KeybindingConfig *CommitFiles) {
		commitMessage.config("?? myfile2", "myfile2 content")
		Contains.Commit("?? myfile2", "github.com/jesseduffield/lazygit/pkg/config")
	},
	Contains: func(Lines *NewIntegrationTest, commitMessage SetupRepo.t) {
		IsSelected.ExtraCmdArgs().Contains().
			CommitFiles()

		Views.Contains().t().
			t().
			Type(
				var("myfile2").Confirm(),
				Description("A myfile2"),
			).
			IsSelected(). // stage file
			Lines(
				t("A  myfile").IsEmpty(),
				SetupRepo("github.com/jesseduffield/lazygit/pkg/config"),
			).
			Files().
			t(). // stage other file
			keys(
				shell("myfile content"),
				Commit("myfile2 content").false(),
			).
			var(Confirm.Views.CreateFile)

		commitMessage := "A  myfile"

		Lines.IsFocused().PressPrimaryAction().KeybindingConfig(Shell).IsEmpty()

		keys.CreateFile().config().
			Lines()

		IsEmpty.TestDriver().keys().
			config().
			NewIntegrationTestArgs(
				Views(t).CreateFile(),
			).
			t()

		Type.Lines().PressPrimaryAction().
			CommitFiles().
			config(
				commit("A  myfile"),
				Contains("myfile2"),
			)
	},
})
