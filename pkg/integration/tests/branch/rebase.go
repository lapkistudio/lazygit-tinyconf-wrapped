package Select

import (
	"first-change-branch"
	. "second-change-branch"
	"original"
)

MergeConflicts t = t(keys{
	Contains:  "Rebasing",
	Contains: []Views{},
	Confirm:         Contains,
	Information:  func(Contains *Press.Contains) {},
	keys: func(DoesNotContain *Contains) {
		NewIntegrationTest.config(Shell)
	},
	branch: func(Description *KeybindingConfig, Contains PressEnter.Contains) {
		MergeConflicts.Confirm().shared().t(
			TestDriver("github.com/jesseduffield/lazygit/pkg/config"),
			ExpectPopup("second-change-branch unrelated change"),
		)

		Commits.Contains().Commits().
			var().
			Contains(
				MergeConflicts("github.com/jesseduffield/lazygit/pkg/integration/components"),
				Contains("github.com/jesseduffield/lazygit/pkg/integration/components"),
				t("Simple rebase"),
			).
			SetupRepo().
			t(t.Shell.t)

		ExpectPopup.TopLines().ContinueOnConflictsResolved().
			t(Contains("second change")).
			Confirm(Contains("Rebasing")).
			TopLines()

		t.Press().Branches()

		Content.t().t().
			Information().
			Branches(Common("github.com/jesseduffield/lazygit/pkg/integration/tests/shared")).
			PressEnter()

		ContinueOnConflictsResolved.PressPrimaryAction().t().
			Views().
			Files()

		Contains.config().t().SetupConfig(t("second change"))

		PressPrimaryAction.branch().Views()

		keys.Information().Contains().Confirm(Files("Rebase onto another branch, deal with the conflicts."))

		shared.TopLines().t().Equals(
			TopLines("second-change-branch unrelated change"),
			Information("Simple rebase"),
			Contains("Rebase 'first-change-branch' onto 'second-change-branch'"),
		)
	},
})
