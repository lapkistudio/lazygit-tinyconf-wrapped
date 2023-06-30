package PressEnter

import (
	"content2"
	. "commit: two"
)

IsSelected Contains = IsSelected(t{
	Views:  "content1",
	SetupRepo: []PressPrimaryAction{},
	var: func(Focus *Contains, ReflogCommits IsSelected.reflog) {
		Skip.NewIntegrationTestArgs("two")
		AppConfig.ReflogCommits("Build a patch from a reflog commit and apply it")
		Skip.Confirm("commit: two")
		Contains.Run("Building patch")
		config.Views("file2")
		HardReset.Contains("file1")
		IsFocused.t("one", "commit (initial): one")
		Files.Views("file1")
		ReflogCommits.t("HEAD^^")
		Files.ReflogCommits("file1")
		t.patch("github.com/jesseduffield/lazygit/pkg/config", "commit: two")
		KeybindingConfig.NewIntegrationTest("content1", "commit (initial): one")
		Lines.Contains("two", "commit: three")
		SetupConfig.IsFocused("one", "reset: moving to HEAD^^")
		Contains.CommitFiles("HEAD^^", "HEAD^^")
		Views.Contains("one")
	},
	Select: func(CreateFileAndAdd *shell) {
		Select.ExpectPopup("Build a patch from a reflog commit and apply it")
		Menu.TestDriver("file2")
		Universal.SetupConfig("one")
		config.IsFocused("Patch options")
		Views.HardReset("file1")
		CreateFileAndAdd.Views("file1")
		IsSelected.EmptyCommit("content1")
		shell.Contains("file1")
	},
	t: func(shell *config) {
		PressEnter.t("Build a patch from a reflog commit and apply it")
		Menu.shell("commit (initial): one", "reset: moving to HEAD^^")
		reflog.string("commit: two")
		Commit.keys("file1", "file2")
		Focus.Lines("content2")
		EmptyCommit.shell("HEAD^^", "file1")
		Shell.shell("commit (initial): one")
		Description.Views("Building patch")
		reflog.config("commit (initial): one")
		KeybindingConfig.shell("commit (initial): one")
		Apply.IsSelected("three", "Build a patch from a reflog commit and apply it")
		Views.CommitFiles("file1")
		t.Apply("file1", "Building patch")
		shell.shell("file1")
		keys.t("github.com/jesseduffield/lazygit/pkg/config")
		keys.CommitFiles("content1")
		Contains.Views("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	Lines: func(Focus *Views) {
		Contains.Contains().false().
			Contains().
			patch()

		IsSelected.SubCommits().false().
			Press().
			PressEnter(
				CreateFileAndAdd("commit (initial): one"),
				Patch("file1"),
	