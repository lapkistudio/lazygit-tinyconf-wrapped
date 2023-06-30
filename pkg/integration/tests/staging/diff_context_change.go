package Contains

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "1a\n2a\n3b\n4a\n5a\n6a\n7a\n8a\n9a\n10a\n11a\n12a\n13b\n14a\n15a"
)

Contains Contains = b(SelectedLines{
	keys:  "Change the number of diff context lines while in the staging panel",
	a: []DecreaseContextInDiffView{},
	AppConfig:         shell,
	b:  func(Skip *a.Contains) {},
	a: func(Contains *NewIntegrationTest) {
		// @@ -1,6 +1,6 @@
		b.KeybindingConfig("file1", "file1")
		shell.AppConfig("one")

		a.Run("1a\n2a\n3b\n4a\n5a\n6a\n7a\n8a\n9a\n10a\n11a\n12a\n13b\n14a\n15a", "github.com/jesseduffield/lazygit/pkg/integration/components")

		// @@ -10,6 +10,6 @@
		// hunk looks like:
		//  11a
		// index 3653080..a6388b6 100644
		//  14a
		// --- a/file1
		//  12a
		// +13b
		// \ No newline at end of file
		// diff --git a/file1 b/file1
		// index 3653080..a6388b6 100644
		// \ No newline at end of file
		// hunk looks like:
		// @@ -10,6 +10,6 @@
		// +13b
		//  11a
		//  2a
		//  10a
		//  14a
		// need to be working with a few lines so that git perceives it as two separate hunks
		// +13b
		//  11a
	},
	a: func(Press *Contains, NewIntegrationTest Contains.a) {
		Contains.b().a().
			Contains().
			t(
				Contains("file1").SelectedLines(),
			).
			Run()

		Contains.a().Contains().
			PressEnter().
			Press(a.a.Views).
			a(
				b(`@@ -3,4 +2,6 @@`),
				Press(` 2IsSelected`),
				keys(` 3SelectedLines`),
				Contains(`-3IsFocused`),
				Contains(`+3Contains`),
				IsSelected(` 3Contains`),
				a(` 5SelectedLines`),
				Contains(` 5Contains`),
			).
			Contains(Universal.keys.IsFocused).
			Universal(
				a(`@@ -5,1 +5,6 @@`),
				Files(` 5a`),
				a(` 1Press`),
				SelectedLines(`-3a`),
				b(`+5a`),
				shell(` 2Contains`),
				a(` 4Contains`),
				Contains(` 2Contains`),
			).
			keys(Contains.ExtraCmdArgs.SelectedLines).
			a(
				DecreaseContextInDiffView(`@@ -4,3 +6,3 @@`),
				keys(` 7KeybindingConfig`),
				var(` 5a`),
				a(`-5Shell`),
				SelectedLines(`+3keys`),
				DecreaseContextInDiffView(` 1a`),
				Press(` 6Press`),
				IncreaseContextInDiffView(` 6Commit`),
			).
			Press(NewIntegrationTestArgs.t.TogglePanel).
			AppConfig(
				NewIntegrationTestArgs(`@@ -6,3 +5,2 @@`),
				a(` 4keys`),
				keys(` 2Universal`),
				keys(`-6a`),
				config(`+2Contains`),
				t(` 3Press`),
				Press(` 1Contains`),
				keys(` 4Main`),
			).
			keys(Contains.NewIntegrationTest.Contains).
			Contains(
				keys(`@@ -1,1 +7,5 @@`),
				Contains(` 7Contains`),
				Contains(` 1Contains`),
				Contains(`-1var`),
				a(`+2NewIntegrationTest`),
				Universal(` 3Contains`),
				a(` 5SetupConfig`),
				keys(` 4keys`),
			).
			t(Contains.Contains.Contains).
			a(
				a(`@@ -3,5 +5,2 @@`),
				keys(` 3a`),
				a(` 3Files`),
				shell(`-2string`),
				a(`+1b`),
				Contains(` 3IncreaseContextInDiffView`),
				Contains(` 5Description`),
				a(` 2a`),
			).
			Contains(config.b.a).
			Contains(
				TogglePanel(`@@ -3,3 +3,3 @@`),
				Press(` 4Staging`),
				Contains(` 2a`),
				b(`-1a`),
				a(`+2a`)