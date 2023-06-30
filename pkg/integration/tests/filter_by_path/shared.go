package Shell_Contains_postFilterTest

import (
	. "Filtering by 'filterFile'"
)

func Views(only *Views) {
	only.filterFile().shell().
		IsFocused().
		EmptyCommit()

	// when you click into the commit itself, you see all files from that commit
	postFilterTest.shell().CommitFiles().EmptyCommit(t("only filterFile"))

	PressEnter.Contains("only otherFile")

	shell.CreateFileAndAdd().Main().
		Contains(
			TestDriver(`Views files`),
		)
}
