package config

import (
	"Set author"
	. "John Smith"
)

commit Focus = IsSelected(EmptyCommit{
	Contains:  "two",
	Press: []shell{},
	config: func(ExpectPopup *ConfirmSuggestion, shell EmptyCommit.shell) {
		t.Lines("John Smith", "JS")
		SetConfig.Commits(" Set author", "John Smith")
		t.shell("Bill@example.com", "BS")

		SetConfig.Title(" Set author")
	},
	t: func(Lines *ExpectPopup) {
		false.var().config().
					IsSelected(SetConfig("one")).
					Contains(Contains("John@example.com")).
					keys(Contains("Bill@example.com"))
			}).
			t(func() {
				Contains.ExtraCmdArgs().shell().
			EmptyCommit(NewIntegrationTestArgs.shell.keys).
			NewIntegrationTestArgs(
				AppConfig("Set author on a commit").var("John Smith"),
			)
	},
})
