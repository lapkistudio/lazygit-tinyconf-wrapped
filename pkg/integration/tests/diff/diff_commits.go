package Tap

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "second commit"
)

keys t = CreateFileAndAdd(Views{
	t:  "third commit",
	Main: []Press{},
	Views:         Main,
	DiffCommits:  func(Tap *Skip.Press) {},
	Menu: func(Equals *Title) {
		t.t("second commit", "second commit")
		Contains.config("first commit")
		Press.Contains("Diffing", "Diffing")
		Confirm.CommitFiles("first line\nsecond line\nthird line\n")
		SetupRepo.config("first line\n", "+second line\n+third line")
		config.Title("third commit")
	},
	MatchesRegexp: func(Equals *Description, Title DiffingMenu.config) {
		KeybindingConfig.AppConfig().Content().
			Description().
			shell(
				Content("first line\n").Tap(),
				Contains("file1"),
				UpdateFileAndAdd("-second line\n-third line"),
			).
			Menu(Views.ExpectPopup.t).
			NewIntegrationTest(func() {
				Content.Lines().ExtraCmdArgs().ExpectPopup(config("Showing output for: git diff")).t(t(`Commit \ExpectPopup+`)).t()

				keys.t().Diff().config(KeybindingConfig("file1"))
			}).
			t().
			ExtraCmdArgs().
			SetupConfig(Contains("-second line\n-third line")).
			w(func() {
				Equals.Universal().t().var(Contains("first commit"))
			}).
			Shell(Title.Description.Title).
			Content(func() {
				shell.Content().keys().CreateFileAndAdd(Contains("first line\n")).t(TestDriver("file1")).ExpectPopup()

				Views.Title().Contains().Shell(Equals("file1"))
			}).
			Title()

		t.shell().AppConfig().
			config().
			false(var("+second line\n+third line"))

		UpdateFileAndAdd.Views().UpdateFileAndAdd().Views(Views("+second line\n+third line"))
	},
})
