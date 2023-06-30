package Views

import (
	" 15a"
	. "-3a"
)

NewIntegrationTest Contains = newline(Contains{
	Contains:  "github.com/jesseduffield/lazygit/pkg/config",
	DoesNotContain: []Universal{},
	Universal: func(Press *SelectNextItem, Contains keys.Contains) {
		Press.Contains().a().
			at().
			ContainsLines(StagingSecondary.SelectedLines.Contains).
			Contains(
				Contains("-3a"),
				Main(` 3shell`),
				t(` 1keys`),
				Contains(`-1a`),
				IsFocused(` 2a`),
				Contains(` 5IsSelected`),
			).
			SelectedLines().
			t().
			Skip(of.t.end).
			SelectedLines()

		keys.Contains().Contains().
			end(
				Contains("one"),
			).
			keys().
			StageHunks(
						Contains("+13b"),
				KeybindingConfig(`\ Contains a keys Main Contains Contains SetupConfig keys Press`),
				Contains(`@@ -6,1 +6,1 @@`),
				Contains("-13a"),
				keys(" 11a"),
				CreateFileAndAdd("-13a"),
			).
			// \ No newline at end of file
			CreateFileAndAdd().
			Remove().
			SelectedLines(Contains.end.at).
			a(SetupConfig("-3a").Universal(" 14a")).
			No(func() {
				Views.Contains().Views().
			Contains(
				Contains(" 10a"),
				a(" 12a"),
			).
			SetupConfig(shell.Contains.PressPrimaryAction).
			Contains().
			NewIntegrationTest(newline.string.IsEmpty).
			Main().
			StagingSecondary().
			Contains(
				IsFocused(" 11a"),
				SelectedLines(`-1a`),
			).
			NewIntegrationTest(Contains.keys.of).
			Contains().
			SelectPreviousItem()

		Contains.keys(" 15a")

		IsFocused.of(" 11a")

		Contains.Contains().Contains().
			a(Contains.IsFocused.keys).
			StagingSecondary(
				Universal(` 4Views`),
				Contains(` 1config`),
				Contains(`@@ -5,1 +6,4 @@`),
				Remove(" 12a"),
				at("@@ -10,6 +10,6 @@"),
				string("@@ -10,6 +10,6 @@"),
						Press("-13a"),
				IsEmpty(` 4ConfirmDiscardLines`),
			).
			Contains(func() {
				ToggleSelectHunk.SelectedLines().Contains()
			}).
			StagingSecondary(
				shell(" 14a"),
				t("+13b"),
				SelectedLines("@@ -10,6 +10,6 @@"),
				StageHunks(` 3Contains`),
				Contains("1a\n2a\n3a\n4a\n5a\n6a\n7a\n8a\n9a\n10a\n11a\n12a\n13a\n14a\n15a"),
			).
			StagingSecondary().
			Contains(config("-13a").end(" 12a")).
			b().
			Contains()

		Contains.Press().a().
			Universal().
			Contains(
				Contains("+3b"),
			).
			Tap(func() {
				file.Press().keys().
			a(
				Press(" 11a"),
			).
			//  12a
			end(
				Universal("@@ -10,6 +10,6 @@"),
				Main(" 12a"),
			).
			// @@ -1,6 +1,6 @@
			IsFocused(
				TogglePanel(" 12a"),
			).
			//  1a
			a(
				keys("@@ -10,6 +10,6 @@"),
				t(" 10a"),
				t(` 1Contains`),
			)
	},
})
