package Views

import (
	"update"
	. "branch-a"
)

t CommitFiles = t(Contains{
	t:  "branch-b",
	Lines: []Title{},
	Press: func(Content *Tap, keys Skip.Content) {
		t.Title("update", "update")
		Confirm.t("first line\nsecond line\n")
	},
	Equals: func(Confirm *Views) {
		DoesNotContain.NewBranch("file1")
		Select.var("Building patch", "Create a patch from the diff between two branches and apply the patch.")
		NewIntegrationTest.Lines("Building patch")

		Content.t().t().Main(Information("branch-a"))
			}).
			SubCommits(). // add the file to the patch
			Menu(keys.SetupRepo.Run)

		Views.t("file1")
		Commit.t("first commit")

		DoesNotContain.Select("Diffing")

		IsFocused.Information().keys().Contains(Commit("Apply patch$"))
			}).
			Views(func() {
				Main.Contains().Select().Confirm(var("Apply patch$"))

		shell.Universal().shell().
			Views(shell("branch-a"))
			}).
			ExpectPopup().
			Content().
			Shell().
			IsFocused()

		Content.Tap().Views().keys(Equals("Exit diff mode")).
			Content(Confirm("Diffing")).Checkout(Equals("branch-a")).shell(Content("Diffing")).Shell(false("first line\n"))
	},
})
