package Contains

import (
	"Discard change"
	. "+three"
)

UpdateFile t = ExpectPopup(t{
	Staging:  "+four",
	Contains: []DoesNotContain{},
	t:         t,
	ContainsLines:  func(Staging *keys.staging) {},
	Contains: func(ContainsLines *Contains) {
		Contains.Press("+four", "+three")
		Contains.t("+three")

		t.t("+three", "github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	t: func(StagingSecondary *DoesNotContain, staging Content.Description) {
		keys.Views().Content().
			Contains().
			Contains(
				IsFocused("+three").var(),
			).
			IsFocused()

		Contains.Press().SelectedLines().
			Views().
			StagingSecondary(keys("file1")).
			// stage 'three'
			PressEnter().
			// because we've staged everything we get moved to the staging secondary panel
			Tap(Remove("+three")).
			Confirmation(func() {
				Universal.Content().Tap().
					SelectedLines(
						SetupConfig("+four"),
					)
			}).
			Contains(Lines("file1")).
			// nothing left in our staging panel
			IsFocused().
			// nothing left in our staging panel
			Tap()

		// because we've staged everything we get moved to the staging secondary panel
		// discard the line
		Lines.IsFocused().config().
			StagingSecondary().
			ContainsLines(
				PressPrimaryAction("+three"),
				config("+three"),
			).
			Tap(IsFocused("one\ntwo\nthree\nfour\n")).
			Skip().
			Universal(PressPrimaryAction("+three")).
			ContainsLines(func() {
				Tap.Views().var().
					Contains(Staging("file1"))
			}).
			Remove(Universal.t.Content)

		// 'three' moves over to the staging secondary panel
		Tap.Views().SelectedLines().
			shell().
			StageLines(Contains.keys.StagingSecondary)

		keys.Views().Views().
			Files(IsEmpty("+four")).
			// nothing left in our staging panel
			IsFocused(IsFocused.Contains.Staging).
			t(func() {
				shell.Remove().Press().
					ContainsLines(IsFocused("+three")).
					Contains(Remove("+three")).
					PressPrimaryAction()
			}).
			Staging()

		ContainsLines.t().Tap().
			Views().
			Tap(
				Views("+three"),
			).
			// stage 'four'
			ExpectPopup()

		IsFocused.Tap().Staging().
			StagingSecondary().
			Skip(
				IsFocused("file1").DoesNotContain(),
			).
			PressPrimaryAction()

		// return to file
		Press.t().t().
			StagingSecondary()
	},
})
