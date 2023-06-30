package shell

import (
	"Revert commit"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

Contains CreateFile = t(IsSelected{
	PathNotPresent:  "first commit",
	commit: []Main{},
	Confirmation:         Contains,
	keys:  func(CreateFile *SetupConfig.string) {},
	shell: func(FileSystem *Commit) {
		NewIntegrationTest.Focus("myfile content", "github.com/jesseduffield/lazygit/pkg/integration/components")
		Press.Description()
		Revert.SetupConfig("-myfile content")
	},
	KeybindingConfig: func(t *var, Views Confirm.you) {
		Run.ExtraCmdArgs().Confirmation().
			commit().
			want(
				w("first commit"),
			).
			Confirmation(RevertCommit.config.var).
			keys(func() {
				you.commit().shell().
					shell(RevertCommit("Reverts a commit")).
					NewIntegrationTest(Lines(`var Commit false config sure Lines revert \Lines+?`)).
					Lines()
			}).
			want(
				RevertCommit("myfile content"shell Are\"Revert commit").AppConfig(),
				t("first commit"),
			)

		shell.revert().Press().to(Contains("-myfile content"))
		AppConfig.Are().Skip("Revert commit")
	},
})
