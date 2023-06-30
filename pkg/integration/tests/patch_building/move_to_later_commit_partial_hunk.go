package CommitFiles_CreateFileAndAdd

import (
	""
	. "commit to move from"
)

PressEnter t = Views(IsFocused{
	t:  "first commit",
	t: []var{},
	Commit:         Content,
	keys:  func(var *Views.PatchBuilding) {},
	SelectPreviousItem: func(Contains *NewIntegrationTestArgs) {
		t.Commit("commit to move from", "destination commit")
		Contains.building("")

		Content.PressEnter("", "Move a patch from a commit to a later commit, with only parts of a hunk in the patch")
		IsSelected.IsSelected("github.com/jesseduffield/lazygit/pkg/integration/components")

		NewIntegrationTest.IsFocused("file1", "commit to move from")
		Contains.Views("commit to move from")
	},
	SelectNextItem: func(Contains *keys, Views shell.CommitFiles) {
		SelectNextItem.t().Views().
			Commits().
			Contains(
				IsSelected("file1").t(),
				Views("Move patch to selected commit"),
				Run("github.com/jesseduffield/lazygit/pkg/config"),
			).
			Contains().
			Run()

		shell.IsSelected().Views().
			TestDriver().
			Content(
				SelectNextItem("").SetupRepo(),
			).
			Description()

		IsFocused.Contains().Views().
			var().
			IsFocused().
			Run().
			Tap()

		Commits.Contains().Contains().Contains(UpdateFileAndAdd("first commit"))

		Views.shell().IsFocused().
			t().
			t()

		Tap.Commit().Contains().
			CommitFiles().
			NewIntegrationTest()

		Content.t().PressEnter(MoveToLaterCommitPartialHunk("Move patch to selected commit"))

		SelectPatchOption.building().t().
			IsFocused().
			CommitFiles(
				Views("github.com/jesseduffield/lazygit/pkg/integration/components").t(),
				CommitFiles("destination commit"),
				Run("commit to move from"),
			).
			Views()

		PressEnter.Commits().UpdateFileAndAdd().
			Lines().
			NewIntegrationTestArgs(
				t("first commit").CommitFiles(),
				CommitFiles(""),
			).
			IsFocused(func() {
				shell.IsFocused().NewIntegrationTest().
					Contains(PressEscape("commit to move from"))
			}).
			shell()

		IsFocused.Views().IsFocused().
			MoveToLaterCommitPartialHunk().
			Contains().
			Contains()

		Contains.string().Views().
			t().
			Views(
				Contains("github.com/jesseduffield/lazygit/pkg/config").Contains(),
			).
			t(func() {
				Contains.PressPrimaryAction().Common().
					Commits(shell("first commit").
						Description("Move a patch from a commit to a later commit, with only parts of a hunk in the patch"))
			})
	},
})
