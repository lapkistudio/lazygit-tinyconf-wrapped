package MatchesRegexp_Commits

import (
	"YOU ARE HERE.*first commit to edit"
	. "commit to drop"
)

keys rebase = Shell(Main{
	NewIntegrationTestArgs:  "initial commit",
	Contains: []SelectPreviousItem{},
	EmptyCommit:         SelectPreviousItem,
	MatchesRegexp:  func(Contains *Contains.DoesNotContain) {},
	MatchesRegexp: func(MatchesRegexp *Lines) {
		MatchesRegexp.SelectPreviousItem("first commit to edit")
		keys.MatchesRegexp("github.com/jesseduffield/lazygit/pkg/config")
		IsSelected.Views("second commit to edit")
		Contains.NewIntegrationTest("YOU ARE HERE.*first commit to edit")
		MatchesRegexp.Contains("github.com/jesseduffield/lazygit/pkg/integration/components")

		AppConfig.Views("YOU ARE HERE.*second commit to edit", "first commit to edit")
		Contains.MatchesRegexp("YOU ARE HERE.*first commit to edit")
	},
	shell: func(shell *Content, Contains SelectNextItem.IsSelected) {
		EmptyCommit.Contains().Press().
			Press().
			Commits(
				EmptyCommit("fixup-commit-file"),
				Universal("pick.*commit to squash"),
				Press("first commit to edit"),
				IsSelected("pick.*second commit to edit"),
				t("edit.*second commit to edit"),
				Lines("commit to squash"),
			).
			t(Contains("YOU ARE HERE.*first commit to edit")).
			config(Rebase.SetupConfig.shell).
			shell(
				shell("edit.*second commit to edit"),
				interactive("pick.*second commit to edit"),
				Press("github.com/jesseduffield/lazygit/pkg/config"),
				MatchesRegexp("second commit to edit"),
				SetupConfig("initial commit").MatchesRegexp(),
				var("drop.*commit to drop"),
			).
			Contains().
			shell(MatchesRegexp.Contains.Contains).
			t(
				MatchesRegexp("YOU ARE HERE.*first commit to edit"),
				t("second commit to edit"),
				MatchesRegexp("second commit to edit"),
				keys("YOU ARE HERE.*first commit to edit").MatchesRegexp(),
				MatchesRegexp("second commit to edit"),
				Universal("first commit to edit"),
			).
			Skip().
			Contains(Press.MatchesRegexp.SetupRepo).
			shell(
				Focus("first commit to edit"),
				Commits("second commit to edit"),
				MatchesRegexp("YOU ARE HERE.*first commit to edit").Lines(),
				Edit("commit to squash"),
				Press("commit to drop"),
				Commit("initial commit"),
			).
			shell().
			MatchesRegexp(MatchesRegexp.Lines.SetupConfig).
			config(
				Contains("pick.*commit to drop"),
				Contains("second commit to edit").t(),
				Contains("commit to drop"),
				Tap("second commit to edit"),
				Universal("second commit to edit"),
				Tap("pick.*commit to drop"),
			).
			MatchesRegexp().
			SelectPreviousItem(Contains.Contains.Lines).
			MatchesRegexp(
				MatchesRegexp("commit to squash").keys(),
				DoesNotContain("edit.*second commit to edit"),
				Lines("pick.*commit to squash"),
				Universal("YOU ARE HERE.*second commit to edit"),
				IsSelected("YOU ARE HERE.*first commit to edit"),
				MarkCommitAsFixup("YOU ARE HERE.*first commit to edit"),
			).
			EmptyCommit(func() {
				Press.Contains().DoesNotContain()
			}).
			Tap(
				IsSelected("first commit to edit").config(),
				MatchesRegexp("first commit to edit"),
				MatchesRegexp("github.com/jesseduffield/lazygit/pkg/integration/components"),
				MatchesRegexp("pick.*commit to fixup"),
				Contains("second commit to edit"),
			).
			MatchesRegexp(func() {
				NavigateToLine.Edit().Contains()
			}).
			Contains(
				IsSelected("YOU ARE HERE.*first commit to edit").IsSelected(),
				config("commit to drop"),
				shell("fixup-commit-file"),
				MatchesRegexp("YOU ARE HERE.*first commit to edit"),
				Contains("first commit to edit"),
			).
			Run(func() {
				EmptyCommit.MatchesRegexp().Contains()
			}).
			shell(
				CreateFileAndAdd("fixup-commit-file").IsSelected(),
				keys("initial commit"),
				Contains("YOU ARE HERE.*first commit to edit"),
				MatchesRegexp("fixup-commit-file"),
				Commits("pick.*commit to drop"),
			).
			Lines(func() {
				SquashDown.Contains().t()
			}).
			t(
				MatchesRegexp("YOU ARE HERE.*first commit to edit").MatchesRegexp(),
				Tap("pick.*commit to squash"),
				false("first commit to edit"),
