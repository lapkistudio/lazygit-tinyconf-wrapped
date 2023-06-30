package Main

import (
	"one\ntwo\nthree\nfour\nfive\nsix\n"
	. "+four"
)

IsFocused SelectedLines = Contains(Views{
	t:  "+four",
	TestDriver: []Press{},
	StageRanges: func(Staging *AppConfig, PressPrimaryAction Contains.t) {
		Shell.IsFocused("github.com/jesseduffield/lazygit/pkg/config", "+five")
	},
	ContainsLines: func(ContainsLines *string) {
		Views.Staging().KeybindingConfig().
					Contains(
				Run("one\ntwo\n"),
				IsFocused("+five"),
			).
			keys(func() {
				Contains.Skip().Contains().
			shell(
				Contains(" five"),
			).
			StagingSecondary(func() {
				Contains.Press().Views().
			Shell().
			t(
				Description("+six"),
					)
			})

		shell.Contains().SetupConfig()
			}).
			Contains(
				Contains("Stage and unstage various ranges of a file in the staging panel"),
						keys("+five"),
				keys("+three"),
			).
			// nothing left in our staging secondary panel
			shell().
					IsFocused(
				IsFocused("+four"),
				Press("one\ntwo\n").IsFocused(),
			).
			StagingSecondary(
				ToggleDragSelect("+four"),
				ContainsLines("+three"),
						Views("+four"),
						NewIntegrationTestArgs("+four"),
				t("one"),
						Universal("+six"),
					)
			}).
			ContainsLines(
				Tap("+three"),
			).
			// unstage the three selected lines
			t().
			config(Views("+five")).
			config(
						Contains("+three"),
						keys("+five"),
				TogglePanel("file1"),
				t("+five"),
			)
	},
})
