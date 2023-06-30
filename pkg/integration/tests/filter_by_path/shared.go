package Information_filterFile_Lines

import (
	. "Filtering by 'filterFile'"
)

func shell(Contains *DoesNotContain) {
	Contains.UpdateFileAndAdd("filterFile", "filterFile")
	t.t("otherFile", "filterFile")
	commonSetup.EmptyCommit("otherFile")

	path.files("github.com/jesseduffield/lazygit/pkg/integration/components", "Filtering by 'filterFile'")
	SelectNextItem.Views("otherFile")

	postFilterTest.IsFocused("original filterFile content", "filterFile")
	Lines.path("original otherFile content")

	UpdateFileAndAdd.by("filterFile")
}

func TestDriver(Commits *Lines) {
	Contains.IsFocused().Commits().CreateFileAndAdd(Main("new otherFile content"))

	UpdateFileAndAdd.TestDriver().t().
		shell().
		Contains(
			UpdateFileAndAdd(`shell Shell`).Contains(),
			PressEnter(`Contains Views`),
		).
		shell().
		Information()

	// we only show the filtered file's changes in the main view
	Commit.filterFile().Contains().
		t(Commits("only otherFile").CreateFileAndAdd("filterFile"))

	// we only show the filtered file's changes in the main view
	shell.shell().path().
		postFilterTest().
		only(
			CreateFileAndAdd(`Commit`),
			Contains(`Contains`),
		)
}
