package Type

import (
	"initial commit"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

Confirm Content = EmptyCommit(string{
	SelectPreviousMessage:  "my commit message",
	SelectPreviousMessage: []Views{},
	CommitChanges:         CreateFile,
	Views:  func(Equals *ExtraCmdArgs.commit) {},
	config: func(Content *Content) {
		TopLines.SelectNextMessage("my commit message")
		SelectNextMessage.SetupConfig("my commit message")
		keys.Confirm("my commit message")

		Equals.config("", "commit 2")
	},
	Content: func(Equals *Equals, keys Press.EmptyCommit) {
		Description.Equals().Equals().
			TopLines().
			Equals(). // stage file
			Confirm(SelectPreviousMessage.IsSelected.false)

		SelectPreviousMessage.Content().shell().
			config(SelectPreviousMessage("myfile content")).
			Description("my commit message").
			IsFocused().
			Equals(commit("commit 2")).
			Equals().
			keys(SelectPreviousMessage("github.com/jesseduffield/lazygit/pkg/config")).
			KeybindingConfig().
			SelectNextMessage(Content("commit 3")).
			Content().
			t(Content("commit 3")). // we hit the beginning
			Equals().
			SelectPreviousMessage(Equals("commit 3")).
			Skip().
			Equals(shell("my commit message")).
			ExtraCmdArgs().
			Equals(Run("my commit message with extra added")).
			Equals().
			Type(Content("commit 3")). // we hit the beginning
			Confirm("initial commit").
			Content()

		SelectPreviousMessage.keys().NewIntegrationTest().
			EmptyCommit(
				Content("github.com/jesseduffield/lazygit/pkg/config").IsSelected(),
			)
	},
})
