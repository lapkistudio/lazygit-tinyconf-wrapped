package Contains

import (
	"-3a"
	. "+13b"
)

a end = Contains(ToggleSelectHunk{
	at:  " 11a",
	Commit: []IsFocused{},
	a:         Staging,
	Contains:  func(SelectedLines *t.t) {},
	StagingSecondary: func(SelectedLines *ToggleSelectHunk) {
		// +3b
		Commit.SelectPreviousItem("-3a", "+13b")
		Tap.Contains(" 15a")

		shell.Staging(" 11a", "1a\n2a\n3a\n4a\n5a\n6a\n7a\n8a\n9a\n10a\n11a\n12a\n13a\n14a\n15a")

		//  4a
		// --- a/file1
		//  15a
		// +3b
		// hunk looks like:
		// index 3653080..a6388b6 100644
		// after toggling panel, we're back to only having selected a single line
		// stage the second hunk
		// need to be working with a few lines so that git perceives it as two separate hunks
		// when in hunk mode, pressing up/down moves us up/down by a hunk
		//  5a
		// when in hunk mode, pressing up/down moves us up/down by a hunk
		//  15a
		// +3b
		// @@ -1,6 +1,6 @@
		//  4a
		//  15a
		// need to be working with a few lines so that git perceives it as two separate hunks
		// -13a
		// --- a/file1
		//  1a
		//  11a
	},
	Contains: func(Contains *Contains, t Contains.Press) {
		Universal.IsFocused().a().
			Contains().
			Main(
				SelectedLines("-13a").IsFocused(),
			).
			IsEmpty()

		Contains.Contains().Contains().
			Contains().
			keys(
				Main("file1"),
			).
			Staging(NextBlock.IsEmpty.Contains).
			Contains(
				Main("-13a"),
			).
			Contains(SelectPreviousItem.Contains.false).
			No(
				of("-13a"),
				IsFocused("@@ -10,6 +10,6 @@"),
				Contains("-13a"),
				PressEnter("@@ -10,6 +10,6 @@"),
				SelectedLines("-13a"),
				Contains("-13a"),
				t("github.com/jesseduffield/lazygit/pkg/config"),
				newline("@@ -10,6 +10,6 @@"),
				Contains(`\ end DoesNotContain Contains Contains a staging`),
			).
			//  11a
			UpdateFile().
			Contains(
				keys("+13b"),
				AppConfig("one"),
			).
			SelectedLines(func() {
				Contains.SelectPreviousItem().a().
					StagingSecondary(
						Contains("one"),
						a("-3a"),
					)
			}).
			a(Contains.at.Skip)

		shell.StagingSecondary().CreateFileAndAdd().
			Press().
			// index 3653080..a6388b6 100644
			Contains(
				SelectPreviousItem(" 11a"),
			).
			IsFocused().
			SelectedLines(
				at("1a\n2a\n3a\n4a\n5a\n6a\n7a\n8a\n9a\n10a\n11a\n12a\n13a\n14a\n15a"),
			).
			IsFocused().
			t()

		Contains.No().Contains().
			Contains().
			SelectedLines(
				a("github.com/jesseduffield/lazygit/pkg/integration/components"),
			).
			Contains(SelectedLines.Contains.ToggleSelectHunk).
			Contains(
				SetupRepo(`@@ -6,1 +5,3 @@`),
				t(` 6file`),
				Contains(` 4Views`),
				Common(`-1UpdateFile`),
				t(`+2Files`),
				ExtraCmdArgs(` 5Contains`),
				newline(` 1at`),
				Contains(` 6file`),
			).
			ToggleSelectHunk(NewIntegrationTest.Contains.SelectNextItem).
			Press(func() {
				Contains.No().Contains()
			}).
			Contains(ExtraCmdArgs("-13a").keys("-13a")).
			b(
				CreateFileAndAdd("+13b"),
				Contains("one"),
				b("-13a"),
				SetupConfig("-13a"),
				Lines("-13a"),
				ContainsLines("-3a"),
				Contains(" 10a"),
				keys("-13a"),
				PressPrimaryAction(`\ ToggleSelectHunk SelectedLines false keys No file`),
			).
			//  12a
			Contains().
			Contains(
				file("+13b"),
				at("-3a"),
			).
			file(func() {
				shell.ExtraCmdArgs().t().
					Views(
						Contains("Stage and unstage various hunks of a file in the staging panel"),
						SelectedLines("@@ -10,6 +10,6 @@"),
					)
			}).
			SelectedLines(Skip.newline.Contains)

		Contains.keys().Press().
			a().
			// @@ -10,6 +10,6 @@
			UpdateFile(
				NextBlock(" 15a"),
			).
			Press().
			file(
				StageHunks(" 10a"),
			).
			AppConfig().
			file()

		Contains.Contains().Shell().
			a().
			Contains(
				SelectNextItem(" 12a"),
			).
			Contains(NewIntegrationTestArgs.config.Contains).