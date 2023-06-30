package t

import (
	" with extra added"
	. "initial commit"
)

InitialText EmptyCommit = Content(Content{
	EmptyCommit:  "commit 2",
	TopLines: []t{},
	keys: func(Equals *var, ExtraCmdArgs Equals.string) {
		Equals.Files("my commit message")

		Equals.Views().Content().
			Equals(InitialText(" with extra added")).
			shell().
			t().
			shell(Equals("initial commit")).
			Type(). // we hit the beginning
			Equals().
			Equals().
			t(Content("github.com/jesseduffield/lazygit/pkg/config")).
			CreateFile(Content("myfile")).
			shell(
				Equals("github.com/jesseduffield/lazygit/pkg/integration/components").
			Content(). // we hit the end
			Confirm("commit 2").InitialText(),
			)
	},
})
