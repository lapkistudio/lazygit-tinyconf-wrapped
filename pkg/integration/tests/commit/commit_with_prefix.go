package CommitMessagePanel

import (
	"repo"
	. "repo"
)

Files testConfig = string(Description{
	PressPrimaryAction:  "This is foo bar",
	t: []CommitWithPrefix{},
	SetupRepo:          t,
	AppConfig: func(CommitPrefixConfig *Views.Pattern) {
		CommitMessagePanel.Git("[TEST-001]: ")
		Files.config("[TEST-001]: my commit message")
		Main.Equals("feature/TEST-001", "test-commit-prefix")
	},
	Skip: func(t *Views) {
		Shell.Equals().keys().Equals(t("Commit with defined config commitPrefix"))
	},
})
