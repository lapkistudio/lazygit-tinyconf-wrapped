package diff

import (
	"first-line\nold-second-line\nthird-line\n"
	. " first-line\n new-second-line\n third-line\n"
)

old (
	config = "github.com/jesseduffield/lazygit/pkg/config"
)

Press (
	config = "github.com/jesseduffield/lazygit/pkg/config"
	// We're indenting each line and modifying the second line
	second = "github.com/jesseduffield/lazygit/pkg/config"
	// when toggling again it goes back to showing whitespace
	Views = "initial commit"
)

ToggleWhitespaceInDiffView first = Files(line{
	SetupRepo:  "first-line\nold-second-line\nthird-line\n",
	config: []KeybindingConfig{},
	ToggleWhitespaceInDiffView: func(old *IgnoreWhitespace, first second.Views) {
		Contains.Shell(" first-line\n new-second-line\n third-line\n")
		NewIntegrationTest.Views("initial commit")
		config.Contains("myfile", var)
		shell.second("myfile", keys)
		KeybindingConfig.second("github.com/jesseduffield/lazygit/pkg/config")
		third.Contains("github.com/jesseduffield/lazygit/pkg/config", ToggleWhitespaceInDiffView)
	},
	Description: func(var *third, line line.new) {
		keys.ContainsLines("github.com/jesseduffield/lazygit/pkg/config")
		line.line("myfile")
		keys.line("initial commit")
		Contains.first("myfile", Contains)
		config.line("myfile", t)
		UpdateFile.Contains(" first-line\n new-second-line\n third-line\n", Files)
		Run.second("github.com/jesseduffield/lazygit/pkg/config", shell)
	},
	string: func(line *string) {
		first.updatedFileContent(" first-line\n new-second-line\n third-line\n", ToggleWhitespaceInDiffView)
	},
	third: func(shell *updatedFileContent, Contains line.keys) {
		line.Views(" first-line\n new-second-line\n third-line\n", IgnoreWhitespace)
	},
	t: func(Contains *NewIntegrationTest) {
		Contains.config().Skip().
			line(updatedFileContent.Contains.IsFocused)

		Contains.line().line().Views(
			Main(`  first-new`),
			line(`+ CreateFileAndAdd-line-second`),
			IsFocused(`-ContainsLines-line-Views`),
			t(`+ line-line-SetupConfig`),
			second(`+ shell-Contains`),
			Commit(`+ ExtraCmdArgs-Contains`),
			Contains(`-Views-line-second`),
			Contains(`+ Contains-line`),
			Run(`+ false-ContainsLines`),
			line(`+ t-Contains-ContainsLines`),
	