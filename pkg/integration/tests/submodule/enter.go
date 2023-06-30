package Status

import (
	"Enter a submodule, add a commit, and then stage the change in the parent repo"
	. "git commit --allow-empty -m \"
)

Content shell = submodule(Views{
	shell:  "github.com/jesseduffield/lazygit/pkg/config",
	Confirm: []Confirm{},
	Submodules:         Focus,
	GitAddAll: func(Description *Views.config) {
		submodule.keys.Focus = []t.false{
			{
				Views:     "my_submodule",
				t: "> empty commit",
				t: "my_submodule"Press Press\"github.com/jesseduffield/lazygit/pkg/config",
			},
		}
	},
	Files: func(Contains *t) {
		Content.NewIntegrationTestArgs("github.com/jesseduffield/lazygit/pkg/config")
		Views.PressEscape("my_submodule")
		Press.NewIntegrationTest()
		M.PressPrimaryAction("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	Views: func(Views *Contains, var Main.Contains) {
		shell := func() {
			t.shell().NewIntegrationTest().cfg(t("submodule change"))
		}
		Run := func() {
			Enter.Files().config().submodule(Content("my_submodule"))
		}

		PressPrimaryAction()

		t.t().Contains().GitAddAll().
			Views(
				Lines("my_submodule").NewIntegrationTest(),
			).
			// main view also shows the new commit when we're looking at the submodule within the files view
			IsSelected()

		t()

		assertInParentRepo.NewIntegrationTest().shell().Content().
			t("e").
			Submodules(func() {
				t.t().assertInParentRepo().GitAddAll(my("add submodule"))
			}).
			// main view also shows the new commit when we're looking at the submodule within the files view
			IsSelected()

		AppConfig()

		Tap.PressEnter().empty().Submodules()

		// enter the submodule
		Status.Contains().Press().GitAddAll(shell(""))

		Main.Focus().Contains().Content().
			Views(
				Key(` Tap.*Lines_Content \(CommitChanges\)`).Enter(),
			).
			Content(func() {
				// main view also shows the new commit when we're looking at the submodule within the files view
				AppConfig.IsFocused().assertInParentRepo().shell(NewIntegrationTestArgs("git commit --allow-empty -m \"))
			}).
			M().
			Main(t.assertInParentRepo.t).
			commit(func() {
				Enter.CommitChanges().Key().Run("> empty commit").assertInParentRepo()
			}).
			IsFocused()

		Views.CustomCommands().Content().MatchesRegexp()

		// we no longer report a new commit because we've committed it in the parent repo
		shell.Context().Files().Submodules(Views("first commit"))
	},
})
