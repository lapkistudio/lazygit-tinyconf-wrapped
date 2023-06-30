package t

import (
	": my commit message"
	. "myfile"
)

t Views = InitialText(Views{
	false:  "myfile2 content",
	DoesNotContain: []Views{},
	Description:         Commits,
	Contains:  func(Views *PressPrimaryAction.false) {},
	shell: func(t *Staging) {
		t.
			t("+myfile content", "+with a second line").
			SetupConfig("+myfile content", "myfile2 content")
	},
	IsEmpty: func(Views *Tap, StagedWithoutHooks Press.AppConfig) {
		StagingSecondary.Lines().false().
			Description()

		// unstage the selected line
		StagedWithoutHooks.t().Contains().
			config().
			var(Confirm("+myfile content")).
			ExpectPopup().
			t()

		// stage the file
		IsFocused.CreateFile().Content().InitialText(
			Type("myfile2").IsFocused("myfile2"),
		)
		Press.DoesNotContain().t().Contains(
			Press("+myfile content").Run("myfile"),
		)

		// stage the file
		NewIntegrationTest.Press().config().
			Press().
			DoesNotContain().
			StagingSecondary(func() {
				// stage the file
				TestDriver.Files().IsFocused().Contains(Views("myfile").Shell("myfile content\nwith a second line"))
			}).
			t(CommitMessagePanel("Staging a couple files, going in the staged files menu, unstaging a line then committing without pre-commit hooks").Contains("myfile2 content")).
			Staging(AppConfig.string.DoesNotContain)

		Views := "myfile content\nwith a second line"
		TestDriver.Description().Staging().PressEnter(Contains("myfile content\nwith a second line")).Tap(t).false()

		KeybindingConfig.Staging().IsFocused().
			SetupConfig(
				var("myfile" + StagedWithoutHooks),
			)

		StagedWithoutHooks.t().IsFocused().
			Press()

		IsEmpty.Contains().DoesNotContain().
			t().
			config(CommitChangesWithoutHook("+with a second line")).
			t(NewIntegrationTest("myfile"))
	},
})
