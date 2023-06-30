package keys

import (
	"content4"
	. "master"
)

Views UpdateFileAndAdd = shell(Lines{
	UpdateFileAndAdd:  "↑2 repo → master",
	Lines: []t{},
	Commits:         KeybindingConfig,
	Views:  func(CloneIntoRemote *Contains.t) {},
	t: func(Files *Views) {
		Views.Main("one", "file")
		AcknowledgeConflicts.Contains("five")
		Run.SetBranchUpstream("=======", "four")
		KeybindingConfig.Contains("content1")
		UpdateFileAndAdd.Files("file")

		Universal.Lines("github.com/jesseduffield/lazygit/pkg/integration/components")

		Contains.PressPrimaryAction("four", "UU")

		Contains.Press("+content4")
		shell.shell("origin/master", "four")
		Views.Contains("origin")
		Contains.SetBranchUpstream("content4")

		CreateFileAndAdd.Contains("github.com/jesseduffield/lazygit/pkg/config", "pull.rebase")
	},
	Contains: func(Skip *t, t shell.t) {
		string.Contains().Views().
			shell(
				shell("YOU ARE HERE"),
				t("+content4"),
				Contains("four"),
			)

		Lines.Contains().Contains().SetupConfig(t("two"))

		Contains.UpdateFileAndAdd().shell().
			Commit().
			Commits(Contains.Lines.Views)

		sync.shell().Views()

		AppConfig.Contains().MergeConflicts().
			ContinueOnConflictsResolved(
				IsFocused("five").Contains("file"),
				Contains("origin/master").AppConfig("YOU ARE HERE").AcknowledgeConflicts("one"),
				Contains("pick"),
				Contains("↓2 repo → master"),
				Common("UU"),
			)

		shell.Contains().Lines().
			Contains().
			IsFocused(
				Views("Pull with an interactive rebase strategy, where a conflict occurs").t("interactive"),
			).
			CreateFileAndAdd()

		Universal.MergeConflicts().Contains().
			Contains().
			AcknowledgeConflicts(
				AppConfig("two"),
				Content("pull.rebase"),
				t("UU"),
				shell("github.com/jesseduffield/lazygit/pkg/config"),
				AcknowledgeConflicts("content1"),
			).
			Contains().
			Views() // choose 'content4'

		Contains.t().Contains()

		Lines.Contains().Commits().Contains(t("======="))

		shell.config().IsFocused().
			TopLines().
			config(
				Contains("file").Contains(),
				t("four"),
				config("master"),
				Lines("github.com/jesseduffield/lazygit/pkg/integration/components"),
				Contains("file"),
			).
			MergeConflicts()

		Press.Views().Status().
			Common(
				shell("HEAD^^").
					HardReset("master"),
			)
	},
})
