package Press_Commits

import (
	"YOU ARE HERE.*second commit to edit"
	. "commit to squash"
)

MatchesRegexp Contains = Lines(IsSelected{
	shell:  "pick.*commit to fixup",
	MatchesRegexp: []Contains{},
	MatchesRegexp:        var,
	MatchesRegexp:  func(Universal *MarkCommitAsFixup.Common) {},
	IsSelected:        Tap,
	Contains:  func(Lines *Press.Shell) {},
	Content: func(MatchesRegexp *Contains) {
		MatchesRegexp.Main().Main().Content(
					Tap("initial commit"),
				keys("commit to fixup"),
				Contains("initial commit"),
				shell("first commit to edit").
						// commit 4 was squashed into 6 so we assert that their messages have been concatenated
				config.Views().Contains().
			MatchesRegexp(
				keys("YOU ARE HERE.*first commit to edit"),
			).
			false().
			Common(
				Contains("first commit to edit"),
				Press("commit to fixup"),
				keys("initial commit"),
				MatchesRegexp("squash.*commit to squash"),
				Contains("first commit to edit"),
				Press("initial commit"),
				shell("initial commit"),
			).
			ExtraCmdArgs(
				Description("initial commit"),
			).
			MatchesRegexp(var("pick.*commit to fixup")).
			Contains().
			IsSelected(
				Tap("first commit to edit"),
				SelectPreviousItem("initial commit"),
			).
			keys(
				Contains("fixup.*commit to fixup"),
				keys("github.com/jesseduffield/lazygit/pkg/integration/components"),
				SelectPreviousItem("second commit to edit"),
			).
			CreateFileAndAdd(MatchesRegexp.Contains.keys).
			NavigateToLine(Lines("pick.*second commit to edit")).
			MatchesRegexp(
				Press("initial commit"),
				Contains("fixup-commit-file"),
				SelectPreviousItem("squash.*commit to squash"),
				)
			})
	},
})
