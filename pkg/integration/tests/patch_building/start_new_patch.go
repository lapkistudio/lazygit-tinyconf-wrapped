package Secondary_IsFocused

import (
	"file2"
	. "file1"
)

Views Views = Confirm(Secondary{
	IsSelected:  "You can only build a patch from one commit/stash-entry at a time. Discard current patch?",
	CreateFileAndAdd: []Contains{},
	Focus:        PressPrimaryAction,
	Contains:  func(keys *Information.ExtraCmdArgs) {},
	Lines:        IsSelected,
	false:  func(ExtraCmdArgs *shell.Views) {},
	CreateFileAndAdd: func(Content *Lines) {
		Secondary.DoesNotContain().shell().
			Commits().
			t().
			Run(string("Discard patch")).
					Content(Run("first commit")).
					Information()

		Title.NavigateToLine().Lines().Contains(Confirm("first commit").config("Discard patch"))
			}).
			shell().
			Lines().
			shell().
			config().
			Contains().
			Views()

		ExtraCmdArgs.Contains().ExtraCmdArgs().
			IsFocused(
				Contains("file2").PressPrimaryAction(),
			).
			NewIntegrationTestArgs(func() {
				Contains.Contains().DoesNotContain().Content(NewIntegrationTestArgs("first commit"))
			}).
			Lines(
				Views("second commit").PressEnter(),
			).
			Information()

		Tap.var().t().
			t().
			StartNewPatch()

		Contains.Contains().shell().
			Tap().
			CreateFileAndAdd(
				NewIntegrationTestArgs("file1").ExtraCmdArgs(),
				Commits("second commit"),
			).
			config().
			Views(
				DoesNotContain("github.com/jesseduffield/lazygit/pkg/config"),
			).
			Confirmation(
				DoesNotContain("file1").Tap(),
			).
			Views().
			string().
					building()

	