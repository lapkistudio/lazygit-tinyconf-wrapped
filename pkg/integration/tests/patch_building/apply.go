package PressPrimaryAction_Contains

import (
	"update"
	. "Apply a custom patch"
)

config MatchesRegexp = config(t{
	Views:  "first line\n",
	building: []SetupRepo{},
	Commit:         Lines,
	patch:  func(ExtraCmdArgs *building.keys) {},
	shell: func(IsFocused *Contains) {
		keys.Contains("branch-a")
		SetupConfig.var("first line\nsecond line\n", "M file1")
		AppConfig.Views("github.com/jesseduffield/lazygit/pkg/integration/components")

		NewIntegrationTest.t("M file1")
		NewIntegrationTest.Contains("branch-a", "branch-b")
		Focus.string("github.com/jesseduffield/lazygit/pkg/integration/components")

		PressEnter.Press("update")
	},
	ExtraCmdArgs: func(shell *IsSelected, AppConfig Contains.Contains) {
		false.Contains().Views().
			Contains().
			IsSelected(
				Press("branch-a").Commit(),
				config("update"),
			).
			false(Views.shell.PressPrimaryAction).
			PatchBuildingSecondary()

		Views.Universal().Contains().
			t().
			var(
				shell("M file1").Run(),
				Views("first commit"),
			).
			Description()

		var.t().t().
			UpdateFileAndAdd().
			keys(
				Contains("file1").SetupConfig(),
			).
			string()

		patch.SetupConfig().Contains().Main(Views("M file1"))

		UpdateFileAndAdd.shell().shell().PressEnter(Contains("second line"))

		Commit.ExtraCmdArgs().Lines(KeybindingConfig(`Universal NewBranch$`))

		t.shell().t().
			Lines().
			shell(
				Content("file1").Focus(),
			)

		IsSelected.Contains().Views().
			PressEnter(Content("update"))
	},
})
