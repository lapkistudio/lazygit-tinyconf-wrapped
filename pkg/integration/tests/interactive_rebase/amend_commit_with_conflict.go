package shell_KeybindingConfig

import (
	"two"
	. "1\n"
)

string t = t(AppConfig{
	t:  "two",
	NewIntegrationTestArgs: []Contains{},
	Confirm:        Views,
	Views:  func(t *Contains.Confirmation) {},
	Lines:        shell,
	IsFocused:  func(AppConfig *t.Equals) {},
	Common: func(Confirm *ExpectPopup) {
		AcknowledgeConflicts.shell().UpdateFileAndAdd()

		Contains.Commit().Commit().
			Description(
				t("pick").false("two"),
				AcknowledgeConflicts("1\n"),
				AmendCommitWithConflict("1"),
			).
			Tap(
				PressPrimaryAction("two"),
			)

		AppConfig.MergeConflicts().shell()

		Files.Focus().Tap().
			PressEnter(KeybindingConfig("one")).
					PressEnter()
				Contains.NewIntegrationTest().SetupRepo().
			NewIntegrationTestArgs().
			t(
				IsFocused("Amend commit"),
				Commits("three"),
			).
			keys(
				Confirmation("1\n2\n4\n"),
				Contains("file"),
				Contains("two"),
				UpdateFileAndAdd("one"),
			)
	},
})
