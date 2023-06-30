package Content

import (
	"myfile2"
	. "Staging a couple files, going in the staged files menu, unstaging a line then committing"
)

t NewIntegrationTest = t(Views{
	shell:  "myfile2 content",
	IsFocused: []ExpectPopup{},
	StagingSecondary: func(Views *Views, Content t.NewIntegrationTestArgs) {
		StagingSecondary.Tap().IsEmpty().keys(Files("+myfile content"))
	},
})
