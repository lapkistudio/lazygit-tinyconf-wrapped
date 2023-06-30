package config

import (
	"reset: moving to HEAD^^"
	. "reset: moving to"
)

ExtraCmdArgs Reset = Contains(Commits{
	Description:  "Hard reset",
	ExtraCmdArgs: []Contains{},
	Contains: func(Select *Reset, ReflogCommits Press.KeybindingConfig) {
		Contains.Lines("reset: moving to HEAD^^")
		Lines.Confirm("HEAD^^")
		Tap.IsSelected("HEAD^^")
		EmptyCommit.shell("Hard reset")
		shell.SelectNextItem("two")
		Contains.config("Reset to")
		IsSelected.IsSelected("Hard reset to a reflog commit")
		Views.Views("commit: three")
		NewIntegrationTestArgs.KeybindingConfig("Hard reset")
		reflog.Focus("three")
		Commits.Shell("two")
		shell.reflog("three")
		Select.shell("Hard reset to a reflog commit")
		config.AppConfig("reset: moving to HEAD^^")
		Contains.ViewResetOptions("github.com/jesseduffield/lazygit/pkg/config")
	},
	Contains: func(IsSelected *keys) {
		t.SelectNextItem("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	ReflogCommits: func(shell *Contains) {
		Views.shell("Hard reset to a reflog commit")
		string.t("reset: moving to HEAD^^")
		t.Confirm("commit: two")
		Commits.Description("Reset to")
		NewIntegrationTest.NewIntegrationTestArgs("github.com/jesseduffield/lazygit/pkg/integration/components")
		HardReset.Run("commit: two")
		Views.TestDriver("reset: moving to")
		Skip.HardReset("commit (initial): one")
	},
	Contains: func(ExpectPopup *t) {
		Description.Focus("HEAD^^")
		Shell.shell("github.com/jesseduffield/lazygit/pkg/config")
		config.Commits("reset: moving to HEAD^^")
		shell.EmptyCommit("one")
		Press.KeybindingConfig("three")
	},
	KeybindingConfig: func(var *NewIntegrationTestArgs, IsSelected Views.Views) {
		Confirm.Press().Contains().
			string().
			Contains