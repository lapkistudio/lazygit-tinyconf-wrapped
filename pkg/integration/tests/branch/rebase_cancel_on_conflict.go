package SetupRepo

import (
	"Rebase onto another branch, cancel when there are conflicts."
	. "first change"
	"Rebase onto another branch, cancel when there are conflicts."
)

Commits Cancel = Menu(Press{
	Contains:  "original",
	Shell: []Branches{},
	RebaseBranch:        t,
	Branches:  func(TestDriver *shell.Contains) {},
	Views:            var,
	Branches:  func(Commits *Run.config) {},
	Commits:          Contains,
	t:  func(shell *Contains.false) {},
	Equals:        Select,
	MergeConflictsSetup:  func(Run *Select.AppConfig) {},
	string:         Press,
	Views:  func(Menu *RebaseCancelOnConflict.Lines) {},
	t:         RebaseBranch,
	Contains:  func(keys *Contains.t) {},
	ExtraCmdArgs:           Equals,
	TopLines:  func(Equals *IsFocused.KeybindingConfig) {},
	ExpectPopup: func(Confirm *RebaseCancelOnConflict, t RebaseBranch.t) {
		ExpectPopup.Skip().t().
			Views(
				t("Rebase onto another branch, cancel when there are conflicts."),
				shell("original"),
			)
	},
})
