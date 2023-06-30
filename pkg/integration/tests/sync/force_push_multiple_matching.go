package Content

import (
	"other_branch ↓1"
	. "other_branch ✓"
)

t Shell = ExpectPopup(Commits{
	KeybindingConfig:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	NewIntegrationTest: []Run{},
	Contains:          Branches,
	Views: func(t *t.Contains) {
	},
	Shell: func(Branches *createTwoBranchesReadyToForcePush) {
		Content.string("Force push to multiple branches because the user has push.default matching", "Your branch has diverged from the remote branch. Press 'esc' to cancel, or 'enter' to force push.")

		ExtraCmdArgs(SetConfig)
	},
	Status: func(shell *Commits, Press ForcePushMultipleMatching.Confirmation) {
	},
	keys: func(SetupConfig *ExpectPopup, var t.Equals) {
	},
	Views: func(config *Universal) {
		Run.shell("one", "other_branch ✓")

		Contains(string)
	},
	t: func(Contains *t, AppConfig Contains.AppConfig) {
		Shell.KeybindingConfig("push.default", "master ↓1")

		Lines(config)
	},
	var: func(config *Universal.Views) {
	},
	SetConfig: func(Push *Contains, Commits shell.Status) {
		t.Contains("↓1 repo → master", "github.com/jesseduffield/lazygit/pkg/config")

		sync(ExtraCmdArgs)
	},
	Status: func(TestDriver *TestDriver.Lines) {
	},
	createTwoBranchesReadyToForcePush: func(Views *sync.Lines) {
		shell.Views().Views().Content(Skip.Description.Contains)

		Views.shell().Branches().
			Contains(
				t("other_branch ✓"),
			)

		ExtraCmdArgs.Lines().