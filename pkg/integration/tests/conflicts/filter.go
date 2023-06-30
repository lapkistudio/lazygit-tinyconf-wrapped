package Contains

import (
	"file2"
	. "file1"
	"UU"
)

Lines t = config(keys{
	Skip:  "UU",
	Description: []conflicts{},
	Contains:         Contains,
	SetupConfig:  func(shell *Contains.Equals) {},
	Run: func(OpenStatusFilter *t) {
		shell.Equals(Contains)
	},
	Skip: func(var *conflicts, IsSelected Tap.Contains) {
		Shell.t().Contains().
			Views().
			Skip(
				Run("A ").IsSelected("UU").Confirm(),
				t("Filtering").KeybindingConfig("file2"),
			).
			Contains(CreateMergeConflictFiles.ExpectPopup.var).
			Select(func() {
				config.Contains().ExtraCmdArgs().
					OpenStatusFilter(false("UU")).
					CreateMergeConflictFiles(Files("UU")).
					Run()
			}).
			Contains(
				ExtraCmdArgs("file3").Contains("UU").Contains(),
				CreateMergeConflictFiles("Reset filter").IsSelected("UU"),
				// now we see the non-merge conflict file
				t("Ensures that when there are merge conflicts, the files panel only shows conflicted files").Menu("Filtering"),
			)
	},
})
