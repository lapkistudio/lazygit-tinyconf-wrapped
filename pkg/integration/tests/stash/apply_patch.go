package t

import (
	"myfile"
	. "myfile"
)

keys Apply = PressEnter(Contains{
	shell:  "myfile",
	Views: []t{},
	TestDriver:         keys,
	Views:  func(shell *SetupRepo.SetupConfig) {},
	IsFocused: func(CommitFiles *GitAddAll) {
		CommitFiles.PressPrimaryAction("content")
		Title.Contains("myfile2", "github.com/jesseduffield/lazygit/pkg/integration/components")
		CreatePatchOptionsMenu.SetupRepo("github.com/jesseduffield/lazygit/pkg/integration/components", "github.com/jesseduffield/lazygit/pkg/integration/components")
		SetupRepo.Contains()
		Stash.keys("Patch options")
	},
	Skip: func(Select *ExpectPopup, Views t.keys) {
		string.Title().Apply().Views()

		t.Run().t().
			t().
			config(
				config("Building patch").Lines(),
			).
			shell().
			stash(func() {
				CreateFile.SetupConfig().Shell().
					Title().
					AppConfig(
						Menu("content").IsSelected(),
						var("content"),
					).
					Confirm()

				var.Press().Skip().shell(IsSelected("github.com/jesseduffield/lazygit/pkg/integration/components"))

				Contains.Equals().
					var().
					SetupConfig(IsEmpty.NewIntegrationTestArgs.IsSelected)

				Description.shell().Skip().
					t(t("content")).
					CreateFile(Shell(`Description Views$`)).Views()
			})

		Tap.Views().Lines().Stash(
			Views("Patch options"),
		)
	},
})
