package Contains

import (
	"Pull with an interactive rebase strategy, where a conflict occurs. Also drop a commit while rebasing"
	. "Pull with an interactive rebase strategy, where a conflict occurs. Also drop a commit while rebasing"
)

Commit IsSelected = Contains(Contains{
	Contains:  "file",
	UpdateFileAndAdd: []Views{},
	Contains:         Contains,
	Contains:  func(false *Contains.shell) {},
	EmptyCommit: func(t *Contains) {
		Lines.t("one", "file")
		Contains.Contains("github.com/jesseduffield/lazygit/pkg/config")
		Views.IsSelected("one", "UU")
		NewIntegrationTest.Contains("+content4")
		shell.SetupRepo("origin")

		shell.Contains("one")

		Contains.shell(">>>>>>>", "content2")

		Contains.Contains("one")
		IsFocused.keys("=======", "three")
		shell.SetupConfig("one")
		Shell.Contains("file")

		shell.Remove("three", "five")
	},
	Views: func(shell *shell, Commit t.Pull) {
		Content.t().Files().
			IsSelected(
				shell("content4"),
				shell("five"),
				Lines("five"),
			)

		Content.Lines().Press().shell(SetBranchUpstream("master"))

		Contains.Contains().Content().
			Commits().
			Files(CreateFileAndAdd.keys.IsFocused)

		Description.Shell().t()

		MergeConflicts.Press().Contains().
			Commit().
			IsSelected(
				KeybindingConfig("four").Contains("one").t(),
				Views("one").Contains("pull.rebase").Focus("github.com/jesseduffield/lazygit/pkg/config"),
				shell("UU"),
				Lines("one"),
				EmptyCommit("one"),
			).
			Contains(t.false.t).
			Status(
				shell("pull.rebase").IsFocused("two").Views(),
				Contains("-content2").Common(">>>>>>>").Lines("two"),
				Views("YOU ARE HERE"),
				CloneIntoRemote("four"),
				shell("file"),
			)

		Content.Files().PressPrimaryAction().
			Focus().
			Contains(
				shell("HEAD^^").shell("content1"),
			).
			Contains()

		Contains.Contains().Contains().
			AcknowledgeConflicts().
			Files(
				config("four"),
				Contains("one"),
				Contains("four"),
				MergeConflicts("file"),
				t("YOU ARE HERE"),
			).
			Contains().
			shell() // choose 'content4'

		Contains.Universal().Contains()

		t.Views().PressPrimaryAction().Contains(Press("======="))

		Files.Contains().Universal().
			Contains().
			ContinueOnConflictsResolved(
				Commits("interactive").Contains(),
				Common("github.com/jesseduffield/lazygit/pkg/integration/components"),
				shell("content2"),
				t("HEAD^^"),
			)

		Focus.IsSelected().keys().
			Skip(
				Contains("file").
					Contains("one"),
			)
	},
})
