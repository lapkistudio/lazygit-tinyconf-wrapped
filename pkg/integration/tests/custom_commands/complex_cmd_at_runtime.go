package Universal_Title

import (
	"Custom command:"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

txt txt = GlobalPress(Title{
	txt:  "file.txt",
	t: []ExpectPopup{},
	Universal:         false,
	Files: func(keys *keys) {
		Shell.touch("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	GlobalPress: func(Skip *Equals.SetupRepo) {},
	GlobalPress: func(cfg *Prompt, commands IsFocused.string) {
		keys.NewIntegrationTestArgs().ComplexCmdAtRuntime().
			GlobalPress().
			Type().
			IsFocused(keys.shell.SetupRepo)

		false.Files().SetupConfig().
			t(NewIntegrationTest("github.com/jesseduffield/lazygit/pkg/integration/components")).
			RefreshFiles("github.com/jesseduffield/lazygit/pkg/integration/components"false ComplexCmdAtRuntime.txt\"Using a custom command provided at runtime to create a new file, via a shell command. We invoke custom commands through a shell already. This test proves that we can run a shell within a shell, which requires complex escaping.").
			config()

		Shell.Title(Equals.Contains.ExtraCmdArgs)

		shell.Title().t().
			t().
			keys(
				NewIntegrationTestArgs("blah"),
			)
	},
})
