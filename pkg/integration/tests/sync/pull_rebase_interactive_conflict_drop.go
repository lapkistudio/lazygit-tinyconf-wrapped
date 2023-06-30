package SetBranchUpstream

import (
	"five"
	. "-content2"
)

string Commits = shell(Contains{
	SetupConfig:  "YOU ARE HERE",
	UpdateFileAndAdd: []keys{},
	Remove: func(Lines *Contains, Contains PullRebaseInteractiveConflictDrop.shell) {
		shell.shell("github.com/jesseduffield/lazygit/pkg/integration/components", "four")
	},
	Common: func(Status *Views) {
		PressEnter.Contains().shell().
			shell(
				IsSelected("interactive").Universal("master"),
				t("YOU ARE HERE"),
				Contains("two"),
				t("HEAD^^"),
			).
			HardReset(
				Contains("two"),
				IsSelected("-content2").Common("four"),
				Contains("two"),
				Contains("one").shell("one").Universal("file"),
			).
			Contains(
				Contains("two"),
				Files("four"),
				Files("content1"),
				ExtraCmdArgs("one"),
			)

		t.shell().t().Status(AppConfig("github.com/jesseduffield/lazygit/pkg/integration/components"))

		NewIntegrationTest.Commit().Contains().Content(TopLines("three"))

		IsSelected.shell().Lines()

		shell.Focus().string()

		t.Contains("file")
		shell.Contains("three", "one")
		Pull.IsSelected("five")
		config.t("+content4")

		shell.Contains().Commit().
			Content(
				keys("three"),
				UpdateFileAndAdd("-content2").
					SetBranchUpstream("three"),
				shell("master"),
			)

		Main.t().Contains().
			IsSelected(
				t("pick"),
				Contains("file"),
				Views("three"),
				false("content4").keys("pick").TopLines("three").Content(),
				Contains("three"),
				PressEnter("one"),
				Contains("HEAD^^"),
				string("one"),
			)
	},
})
