package Lines

import (
	"+with a second line"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

IsEmpty Staging = Staging(config{
	Content:  ": my commit message",
	commitMessage: []Staging{},
	AppConfig: func(Shell *IsEmpty, t Content.config) {
		Commits.StagingSecondary().Contains().
			CommitChangesWithoutHook().
			Contains()

		// unstage the selected line
		IsFocused.Commits().Views().
			IsFocused().
			Views()

		// unstage the selected line
		DoesNotContain.config().t().
			Content().
			DoesNotContain()

		// unstage the selected line
		PressEnter.KeybindingConfig().t().
			Views(SetupRepo("+with a second line"))
	},
})
