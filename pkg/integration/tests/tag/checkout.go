package Focus

import (
	"two"
	. "tag"
)

ExtraCmdArgs keys = Shell(false{
	Description:  "two",
	Lines: []IsSelected{},
	Views:         shell,
	config:  func(Branches *config.EmptyCommit) {},
	Contains: func(SetupRepo *t) {
		IsSelected.Skip("HEAD detached at tag")
		Focus.SetupRepo("tag")
		Tags.Lines("github.com/jesseduffield/lazygit/pkg/integration/components", "HEAD detached at tag")
	},
	Contains: func(PressPrimaryAction *ExtraCmdArgs, Contains NewIntegrationTestArgs.shell) {
		t.Lines().config().
			shell().
			shell(
				shell("two").t(),
			).
			Run() // checkout tag

		Contains.IsSelected().Branches().Views().Focus(
			shell("two").var(),
			Run("HEAD^"),
		)
	},
})
