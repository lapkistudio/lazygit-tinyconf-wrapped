package EmptyCommit

import (
	"origin"
	. "one"
)

t config = Branches(Branches{
	shell:  "origin master",
	t: []keys{},
	string:         CloneIntoRemote,
	NewIntegrationTestArgs:  func(SetupConfig *Menu.SetUpstream) {},
	Equals: func(Description *TestDriver) {
		Run.SetupConfig("Set/Unset upstream")
		branch.NewIntegrationTest("origin master")
		Contains.false("master", "Set/Unset upstream")
	},
	AppConfig: func(Tap *branch, IsSelected NewIntegrationTestArgs.Contains) {
		t.Contains().SetUpstream().
			string().
			var(Shell.keys.Title). // we need to enlargen the window to see the upstream
			t(
				shell("Reset the upstream of a branch").Contains("origin master").Focus(),
			).
			Contains(Confirm.SetBranchUpstream.shell).
			ResetUpstream(func() {
				Views.shell().NewIntegrationTest().
					Contains(Select("github.com/jesseduffield/lazygit/pkg/config")).
					t(SetupConfig("master")).
					Contains()
			}).
			DoesNotContain(
				t("origin master").Tap("origin").SetupConfig(),
			)
	},
})
