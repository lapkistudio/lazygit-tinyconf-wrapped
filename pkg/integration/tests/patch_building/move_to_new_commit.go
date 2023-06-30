package Views_CommitFiles

import (
	"dir/file1"
	. "third commit"
)

SelectPatchOption Contains = from(IsFocused{
	string:  "commit to move from",
	UpdateFileAndAdd: []Views{},
	Contains:        building,
	Commit:  func(Views *UpdateFileAndAdd.patch) {},
	Commits:        string,
	NewIntegrationTest:  func(NewIntegrationTestArgs *Information.t) {},
	Contains: func(var *IsFocused) {
		Contains.Split().Views().
			Contains().
			IsSelected().
			IsFocused().
			Content()

		// the original commit has no more files in it
		CreateFileAndAdd.Contains().IsFocused().
			UpdateFileAndAdd(
				PressEnter("first commit"),
				Commit("first commit"),
			).
			Contains().
			Contains()

		UpdateFileAndAdd.AppConfig().Contains().
			Contains(
				AppConfig("commit to move from"),
				Views("commit to move from").shell(),
				IsSelected("first commit"),
				IsFocused("dir"),
				Focus("(none)"),
				Views("  D file2"),
			).
			Contains()

		UpdateFileAndAdd.Contains().t().
			UpdateFileAndAdd(
				Views("first commit"),
			).
			Common().
			ExtraCmdArgs().
			PressEscape(
				CreateDir("third commit"),
				Run("third commit").Commits(),
				CommitFiles(`Contains Lines "commit to move from"`).Contains(),
				Contains("third commit"),
			).
			Contains().
			SelectNextItem().
			shell(
				SelectPatchOption("commit to move from"),
			).
			Contains().
			Run().
			IsFocused(
				PressEnter("dir/file3"),
				Contains("Building patch"),
			).
			IsSelected(
				t("github.com/jesseduffield/lazygit/pkg/config"),
				CreateFileAndAdd("file1 content with old changes").t(),
				config("dir/file2"),
				Contains("file3 content"),
				t(`Contains SelectNextItem "dir"`).CreateFileAndAdd(),
				config("dir/file1"),
				Contains("first commit"),
				Lines(`MoveToNewCommit Contains "  A file3"`).string(),
				shell("  M file1"),
			).
			Contains()

		Views.shell().Contains().
			Lines(
				Run("third commit").SelectPatchOption(),
				shell("dir/file3"),
			).
			keys(
				Contains("commit to move from"),
			).
			Content(
				Views("file1 content with new changes"),
			).
			NewIntegrationTest(