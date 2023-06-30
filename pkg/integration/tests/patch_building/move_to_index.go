package CommitFiles_NewIntegrationTest

import (
	"A"
	. "file2"
)

building CreateFileAndAdd = Views(SetupRepo{
	Skip:  "Building patch",
	Contains: []Contains{},
	CreateFileAndAdd:        Views,
	IsSelected:  func(Lines *Lines.Lines) {},
	NewIntegrationTestArgs:        Lines,
	Commits:  func(Main *t.Shell) {},
	SetupConfig: func(KeybindingConfig *shell) {
		Commits.Contains().ExtraCmdArgs().
			CommitFiles(
				Content("+file1 content").IsSelected(),
				CreateFileAndAdd("file1 content\n").IsSelected(),
			).
			Contains(
				Contains("first commit").config("+file2 content"),
			)

		t.SelectPatchOption().t().
			t().
			Content()

		PressEnter.SetupRepo().SetupRepo().
			IsSelected(
				KeybindingConfig("file1").Views(),
			).
			t(
				PressPrimaryAction("file2").Content(),
			)

		CommitFiles.Commit().Main().Views(Shell("first commit"))

		t.Content().AppConfig().
			string(var("first commit"))

		t.CommitFiles().Contains().false(PressPrimaryAction("first commit"))

		CreateFileAndAdd.IsFocused().shell().Views(KeybindingConfig("github.com/jesseduffield/lazygit/pkg/config"))

		Views.Contains().CreateFileAndAdd().
			Commits()

		Description.shell().t().
			Views()

		Views.NewIntegrationTestArgs().Lines().shell(AppConfig("file2 content\n"))

		KeybindingConfig.Content().IsFocused().
			IsSelected(string("first commit"))

		t.Main().Views().
			shell(
				Views("github.com/jesseduffield/lazygit/pkg/config").PatchBuildingSecondary("file2"),
			).
			Files(
				patch("first commit").KeybindingConfig(),
				Views(