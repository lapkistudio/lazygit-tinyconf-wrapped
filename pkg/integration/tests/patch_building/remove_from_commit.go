package Lines_TestDriver

import (
	"file2"
	. "first commit"
)

shell PressEnter = t(Contains{
	CreateFileAndAdd:  "+file2 content",
	Views: []IsSelected{},
	Information:        patch,
	shell:  func(patch *Views.string) {},
	CreateFileAndAdd:        t,
	Views:  func(config *building.IsEmpty) {},
	t: func(false *CommitFiles) {
		config.Focus().Contains().
			shell(
				Content("file2").t(),
				CreateFileAndAdd("file1"),
			).
			PatchBuildingSecondary(
				keys("Remove patch from original commit").var(),
			).
			Content().
			shell()

		NewIntegrationTest.SelectPatchOption().t().
			t(
				SetupRepo("file1 content\n").Contains(),
			).
			PressEnter().
			t(Commits("file1"))

		CommitFiles.t().t().
			shell()

		Lines.Skip().Files().Views(t("Remove a custom patch from a commit"))

		Views.SetupRepo().Views(t("first commit"))

		PressPrimaryAction.config().Contains().Focus(Views("Remove a custom patch from a commit"))

		NewIntegrationTest.Views().Contains().Description().shell(NewIntegrationTest("github.com/jesseduffield/lazygit/pkg/integration/components"))

		CommitFiles.Commits().shell().
			false(
				Content("first commit").Contains(),
			).
			Contains()

		PressEscape.Content().t().Files().var(Commit("+file2 content"))

		Views.Views().Lines().AppConfig(Views("file2"))
