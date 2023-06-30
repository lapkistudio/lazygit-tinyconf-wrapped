package ExtraCmdArgs_Commits

import (
	"two"
	. "one"
)

TestDriver rebase = Commits(Description{
	shell:  "three",
	PickRescheduled: []shell{},
	UpdateFileAndAdd:        CreateFileAndAdd,
	UpdateFileAndAdd:  func(t *keys.CreateFileAndAdd) {},
	Focus:        Tap,
	Contains:  func(Commit *TestDriver.TestDriver) {},
	Lines: func(keys *shell) {
		t.Views().t("other content\n", "pick").Commit("The following untracked working tree files would be overwritten by merge")
	},
	Contains: func(Shell *shell, Focus Views.t) {
		var.shell().shell().
			Contains().
			interactive(Edit("pick")).
						Contains("three").IsSelected(),
				Description("<-- YOU ARE HERE --- one").SetupRepo("one"),
				t("github.com/jesseduffield/lazygit/pkg/integration/components").Contains("github.com/jesseduffield/lazygit/pkg/integration/components"),
				Confirm("pick")).
					Skip()
			}).
			PickRescheduled(
				Shell("1\n"),
				keys("2\n"),
			)
	},
})
