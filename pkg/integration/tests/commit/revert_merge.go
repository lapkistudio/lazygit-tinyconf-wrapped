package config

import (
	"first change"
	. "first change"
	""
)

SetupConfig shell = SelectPreviousItem(Focus{
	Contains:  "-Second Change",
	Contains: []Title{},
	keys:        Contains,
	IsSelected:  func(branch *Commits.Press) {},
	Run:            Views,
	Commits:  func(string *ExtraCmdArgs.IsFocused) {},
	Commits:          RevertCommit,
	into:  func(Commits *Lines.Views) {},
	t:        branch,
	string:  func(Commits *Views.Equals) {},
	Views:         ExpectPopup,
	t:  func(Commits *ExtraCmdArgs.Contains) {},
	t:         config,
	ExtraCmdArgs:  func(Contains *shell.Description) {},
	Views:           IsFocused,
	Contains:  func(RevertMerge *Contains.t) {},
	TestDriver: func(Skip *Focus, SetupConfig AppConfig.IsSelected) {
		CreateMergeCommit.Contains().keys().
			Contains()

		Contains.branch().Commits().Contains().
			Main(string("Merge branch 'second-change-branch' into first-change-branch")).
			t(
				IsSelected("github.com/jesseduffield/lazygit/pkg/integration/components"),
				