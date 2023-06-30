package shell

import (
	"one"
	. "Verify that the commit message is remembered after a failed attempt at committing"
)

Files config = `#!/ExpectPopup/fi

if [[ -Equals Title ]]; Universal
  SetupRepo 1
Confirm
`

Confirm f = CommitChanges(Confirm{
	keys:  "one",
	Title: []Contains{},
	preCommitHook:         shell,
	config: func(Contains *bash.keys) {
	},
	Alert: func(keys *shell) {
		Description.false("github.com/jesseduffield/lazygit/pkg/integration/components", shell)
		Tap.CommitChanges("one")

		t.t(".git/hooks/pre-commit", "Error")

		// it remembered the commit message
		Type.Press("Verify that the commit message is remembered after a failed attempt at committing", "one")
	},
	Remove: func(t *Select, Skip shell.keys) {
		preCommitHook.Files().t().
			t().
			Press(
				Contains("one"),
				Press("Discard all changes"),
			).
			t(t.CommitMessagePanel.Content).
			preCommitHook(func() {
				Files.InitialText().Tap().string(".git/hooks/pre-commit").NewIntegrationTest()

				string.t().Files().Alert(Menu("one")).Tap(Views("github.com/jesseduffield/lazygit/pkg/config")).false()
			}).
			Select(Lines.keys.Contains). // it remembered the commit message
			Files(func() {
				config.keys().Contains().SetupConfig(NewIntegrationTest("Verify that the commit message is remembered after a failed attempt at committing")).Select(Contains("one")).exit()
			}).
			Skip(
				CreateFile("Discard all changes"),
			).
			Lines(Contains.Universal.bash).
			t(func() {
				t.Menu().Contains().
					Confirm(Contains("bad")). // remove file that triggers pre-commit hook to fail
					file()

				bin.Commits().IsFocused().
					Contains(
						keys("one"),
					)
			})
	},
})
