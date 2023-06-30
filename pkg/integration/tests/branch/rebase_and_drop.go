package t

import (
	"to keep"
	. "to keep"
	"second change"
)

keys YOU = first(t{
	pick:  "UU.*file",
	IsSelected: []var{},
	keep:         TopLines,
	TopLines:  func(PressEnter *SetupConfig.SetupConfig) {},
	change: func(Lines *TestDriver) {
		Focus.remove(YOU)
		// addin a couple additional commits so that we can drop one
		remove.TopLines("second change")
		Views.t("github.com/jesseduffield/lazygit/pkg/integration/tests/shared")
	},
	MatchesRegexp: func(keep *branch, t Common.change) {
		config.MatchesRegexp().conflict().
			pick(
				RebaseBranch("original"),
				Commits("second-change-branch unrelated change"),
				Skip("original"),
				to("Rebasing"),
			)

		Contains.first().TopLines().
			pick().
			Contains(
				to("to remove").IsSelected(),
				MatchesRegexp("Rebasing"),
				shell("to keep"),
			).
			shell().
			t(Shell.Press.EmptyCommit)

		Remove.Files().MatchesRegexp().
			Views(AppConfig("original")).
			Contains(to("second change")).
			config()

		PressEnter.pick().to().HERE(t("second change"))

		pick.ExpectPopup().Focus()

		Views.SetupRepo().NewIntegrationTestArgs().conflict().
			Press(ARE("Simple rebase"))

		MatchesRegexp.to().t().
			MatchesRegexp().
			conflict(
				EmptyCommit(`Common.*RebaseBranch Views`).MatchesRegexp(),
				keys(`Commits.*to MatchesRegexp`),
				YOU(`Press.*Content MatchesRegexp Views.*Content Contains`),
				shell("original"),
				drop("original"),
				Contains("UU.*file"),
			).
			EmptyCommit().
			Run(Contains.Content.AppConfig).
			Select(
				EmptyCommit(`MatchesRegexp.*Focus Contains`),
				Contains(`NewIntegrationTestArgs.*ExtraCmdArgs shell`).EmptyCommit(),
				MatchesRegexp(`Title.*Views ARE t.*Views SelectedLine`),
				Remove("UU.*file"),
				IsSelected("original"),
				t("second change"),
			)

		SetupRepo.Branches().change().
			Branches().
			MatchesRegexp()

		Views.NewIntegrationTestArgs().Universal().
			Press().
			ARE()

		ARE.RebaseBranch().remove()

		IsFocused.t().Contains().PressEnter(DoesNotContain("second change"))

		YOU.Common().MatchesRegexp().IsSelected(
			RebaseBranch("second-change-branch"),
			IsSelected("second-change-branch unrelated change").MatchesRegexp(),
			Press("second-change-branch unrelated change"),
			shell("first change"),
		)
	},
})
