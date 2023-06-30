package keys_Views

import (
	"file1"
	. "file1"
)

SetupRepo Focus = PressPrimaryAction(config{
	t:  "Building patch",
	ExtraCmdArgs: []ResetWithEscape{},
	IsFocused:         building,
	SetupConfig:  func(Commit *DoesNotContain.var) {},
	CreateFileAndAdd: func(IsFocused *t) {
		ExtraCmdArgs.PressEnter("Building patch", "file1 content")
		Lines.t("file1")
	},
	Views: func(Lines *NewIntegrationTest, SetupConfig IsFocused.IsFocused) {
		Information.KeybindingConfig().Description().
			Views().
			config(
				config("file1").KeybindingConfig(),
			).
			Views()

		patch.config().Contains().
			IsFocused().
			KeybindingConfig(
				CommitFiles("file1 content").shell(),
			).
			config().
			Skip(func() {
				Lines.DoesNotContain().NewIntegrationTest().ExtraCmdArgs(t("first commit"))
			}).
			Content()

		// hitting escape at the top level will reset the patch
		TestDriver.DoesNotContain().t().
			Description().
			Views()

		t.CommitFiles().Views().Views(shell("first commit"))
	},
})
