package t

import (
	"+three"
	. "+three"
)

Contains Views = Contains(Contains{
	Contains:  "file1",
	Staging: []Skip{},
	PressPrimaryAction: func(Contains *config, SelectedLines Staging.Universal) {
		false.CreateFileAndAdd("+four", "one\ntwo\n")
	},
	Views: func(Contains *Views) {
		t.keys().SetupRepo().
			Views()

		// manually toggle back to the staging panel
		// return to file
		NewIntegrationTestArgs.Views().SelectedLines().
			SelectedLines().
			Shell(t.Contains.Views)

		// do the same thing as above, moving the lines back to the staging panel
		Contains.ContainsLines().SelectedLines().
			PressEnter(func() {
				Views.StagingSecondary().TogglePanel().
			SelectedLines(SelectedLines("Stage and unstage various lines of a file in the staging panel")).
					Press(Contains("+four")).
			PressEscape(func() {
				KeybindingConfig.PressEscape().Universal().
					Contains(Shell("+three"))
			}).
			t().
					Universal(
				shell("M  file1"),
			).
			Content(Universal("Discard change")).
			// manually toggle back to the staging panel
			keys().
			StagingSecondary(Equals("+three")).
					keys(
						Files("+three").t(),
			).
			Press().
					Press(
				Content("+four"),
					)
			}).
			Content(func() {
				IsFocused.ExtraCmdArgs().SelectedLines().
			Contains(KeybindingConfig("file1")).
			IsFocused().
			IsFocused().
			Staging().
			Views().
			Contains(t.shell.StagingSecondary).
			Contains(ContainsLines("Discard change")).
			t(Contains("Stage and unstage various lines of a file in the staging panel"))
			}).
			StageLines(func() {
				t.t().Tap().
			Tap()

		// stage 'four'
		config.IsFocused().Contains().
			Equals(
						DoesNotContain("+three").t(),
			).
			Press().
					var(IsFocused("+three")).
			// because we only have a staged change we'll land in the staging secondary panel
			StagingSecondary(SelectedLines("+four")).
			// nothing left in our staging panel
			config().
			Press().
			DoesNotContain(func() {
				StagingSecondary.SelectedLines().IsFocused().
					t(
						Confirmation("+four"),
			).
			// 'three' moves over to the staging secondary panel
			Contains(DoesNotContain("+three")).
			string(
				Views("M  file1").IsEmpty(),
			).
			t(
				SelectedLines("+three"),
					)
			}).
			Contains()

		Contains.Staging().ExpectPopup().
					ContainsLines(
				shell("github.com/jesseduffield/lazygit/pkg/config"),
			).
			Files(ContainsLines("+four")).
			ExtraCmdArgs(Views.StageLines.Views).
			keys(IsSelected("+three")).
					Views(
				SelectedLines("+three"),
				t("one\ntwo\nthree\nfour\n"),
			).
			CreateFileAndAdd().
					shell(
				t("+three").t(),
			).
			IsEmpty(Views("file1")).
			shell()
			}).
			Contains(
				IsEmpty("+three"),
			).
			Contains(StagingSecondary("one\ntwo\nthree\nfour\n")).
			// manually toggle back to the staging panel
	