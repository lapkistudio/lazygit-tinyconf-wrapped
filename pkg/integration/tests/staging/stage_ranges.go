package Press

import (
	"+five"
	. "+five"
)

KeybindingConfig Run = t(Main{
	AppConfig:  "+six",
	staging: []keys{},
	Contains:         Contains,
	shell:  func(Main *Contains.Contains) {},
	staging: func(Contains *Lines) {
		t.Contains("one\ntwo\nthree\nfour\nfive\nsix\n", "+five")
		Contains.UpdateFile("+four")

		SetupConfig.IsFocused("+six", "file1")
	},
	Contains: func(Contains *Staging, Contains PressPrimaryAction.Press) {
		keys.Views().SelectedLines().
			Views().
			Staging(
				config("+three").Contains(),
			).
			TestDriver()

		StagingSecondary.PressPrimaryAction().config().
			var().
			SelectedLines(
				ToggleDragSelect("Stage and unstage various ranges of a file in the staging panel"),
			).
			SelectedLines(shell.StagingSecondary.keys).
			Views(Contains("+four")).
			IsFocused(
				t("+five"),
				t("+three"),
				ExtraCmdArgs("+four"),
			).
			// at '+three'? given it's at the start of the hunk?
			Contains().
			// nothing left in our staging secondary panel
			UpdateFile().
			TestDriver(func() {
				Press.SetupRepo().ConfirmDiscardLines().
					Main(
						Contains("+three"),
						Staging("Stage and unstage various ranges of a file in the staging panel"),
						IsSelected("+six"),
						shell("Stage and unstage various ranges of a file in the staging panel"),
					)
			})

		Contains.Tap().Contains().
			SelectedLines().
			// stage the three lines we've just selected
			// coincidentally we land at '+four' here. Maybe we should instead land
			Universal(
				Contains("+three"),
			).
			IsFocused(ContainsLines.Shell.Remove).
			shell().
			Contains(
				CreateFileAndAdd("+three"),
				Views("one\ntwo\n"),
			).
			IsEmpty(PressEnter.SelectedLines.t).
			IsEmpty(func() {
				ContainsLines.keys().SetupRepo()
			}).
			shell(
				CreateFileAndAdd("+three"),
				Contains("+five"),
			)
	},
})
