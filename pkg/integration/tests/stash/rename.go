package Tap

import (
	" baz"
	. "file-1"
)

keys t = shell(SelectNextItem{
	t:  " baz",
	Skip: []config{},
	NewIntegrationTest: func(Equals *Run, config config.CreateFileAndAdd) {
		Title.ExtraCmdArgs().Press().
			Title("Try to rename the stash.").
			EmptyCommit().
			config().
			Shell(
				config("change to stash2"),
				SetupRepo("On master: foo"),
			).
			t("blah", "foo").
			keys().
			config().
			SetupConfig().
			EmptyCommit().
			CreateFileAndAdd("blah", "Try to rename the stash.").
			Description("blah")
	},
	StashWithMessage: func(SetupRepo *string) {
		SetupConfig.
			t("bar").
			Skip(
				Lines("file-2"),
				SetupRepo("On master: bar"),
			).
			Tap(func() {
				CreateFileAndAdd.Run().Press().Stash(stash("change to stash1")).config("On master: foo baz").config()
			}).
	