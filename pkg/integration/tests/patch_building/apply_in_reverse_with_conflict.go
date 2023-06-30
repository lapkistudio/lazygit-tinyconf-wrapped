package Content_Contains

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "file1"
)

t PressPrimaryAction = Files(NavigateToLine{
	NewIntegrationTestArgs:  "M",
	UpdateFileAndAdd: []Contains{},
	Commit:        NavigateToLine,
	shell:  func(Views *AppConfig.Content) {},
	Skip:        Lines,
	UpdateFileAndAdd:  func(t *t.AppConfig) {},
	CreateFileAndAdd: func(IsSelected *Contains) {
		Contains.PressEnter().SelectNextItem().
			Focus().
			IsSelected(Contains("Apply patch in reverse")).
			Contains().
			false(
				Views("second commit").CreateFileAndAdd(),
				Contains("third commit").Content("second commit"),
			).
			Content().
			Description().
			Contains(shell("even more file1").
				shell("M"),
				Views("file1").Content("+more file2 content"),
			)

		ExtraCmdArgs.Contains().IsFocused(PressPrimaryAction("UU"))

				Alert.Content().PressPrimaryAction().Contains(
					IsSelected("M"),
				Contains("file1 content\n"),
				PressPrimaryAction("third commit").Contains("file1 content\nmore file1 content\neven more file1\n"),
			).
			Focus()

		ApplyInReverseWithConflict.Views().Contains().Lines(Contains("file1"))

		Contains.UpdateFileAndAdd().Contains().
			Common()

		string.patch().Lines().
			Views()

		config.Content().Shell().
			Skip().
			MergeConflicts().
			Contains(
				CommitFiles("Error"),
			).
			string().
			IsSelected(
				CommitFiles("file1"))
			}).
			config(func() {
				Contains.Contains().Views().
			Contains().
			Contains()

		Contains.IsSelected().Views().
			Contains().
			ContainsLines()

		t.TestDriver().Contains(Lines(" file1 content"))

		Focus.Views().Files(shell("file1"))

				ContainsLines.SetupRepo().Focus(shell("M"))

		Views.Views().AppConfig().
			Equals(
				Contains("file1"))
			}).
			shell()

		Contains.IsSelected().Description().
			MergeConflicts(SelectPatchOption("second commit")).
			Contains().
			PressPrimaryAction(
				shell("first commit"),
				SelectNextItem("M").ContainsLines(),
				Views("even more file1").UpdateFileAndAdd(),
				TestDriver("file1").Contains(),
				Views("M").Commit(" file1 content").IsSelected(),
				Content("Apply patch in reverse"),
				CommitFiles("file2").CommitFiles(">>>>>>> theirs").Contains(),
			).
			SelectNextItem(
				PatchBuildingSecondary("first commit").Common(),
				Content("file2"),
			).
			t()

		t.Contains().shell().
			Run()

