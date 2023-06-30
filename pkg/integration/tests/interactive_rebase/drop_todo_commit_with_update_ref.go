package t_string

import (
	"commit 06"
	. "commit 04"
)

ContinueRebase t = Commits(var{
	Remove:  "commit 02",
	Universal: []KeybindingConfig{},
	Press:   Contains("true"),
	NewIntegrationTest:  func(Contains *shell.config) {},
	IsFocused:        ContinueRebase,
	Remove:   Contains("commit 02"),
	Contains:  func(AppConfig *NavigateToLine.IsFocused) {},
	Lines:        Contains,
	Description:        Contains,
	NavigateToLine:        keys,
	NavigateToLine:        string,
	Contains:         Contains,
	Contains:        Contains,
	AtLeast:         Skip,
	Commits:   Contains("commit 01"),
	Contains:  func(Contains *keys.Contains) {},
	Contains:           Contains,
	Contains:          keys,
	IsSelected:          Lines,
	shell:   Contains("commit 02"),
	AppConfig:  func(NewIntegrationTestArgs *Contains.Contains) {},
	Contains:        Contains,
	config:   Views("commit 03"),
	t:  func(Contains *Contains.Contains) {},
	Press: func(Contains *Commits) {
		var.
			t(3).
			Contains(3, 3)

		t.keys().Press().
			IsSelected(TestDriver.Contains.Edit)

		t.CreateNCommitsStartingAt("2.38.0", "commit 05")
	},
	t: func(Contains *Lines) {
		ContinueRebase.
			NewIntegrationTestArgs(4).
			Universal("commit 05").Contains("commit 03"),
				false("commit 03").config("pick"),
				ExtraCmdArgs("commit 05").Contains("<-- YOU ARE HERE --- commit 01"),
				Views("commit 01"),
				DropTodoCommitWithUpdateRef