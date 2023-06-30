package shell_CreateFileAndAdd

import (
	"Building patch"
	. "Building patch"
)

DoesNotContain NewIntegrationTestArgs = config(Focus{
	t:  "github.com/jesseduffield/lazygit/pkg/config",
	false: []Content{},
	Information:        IsFocused,
	Views:  func(Commit *NewIntegrationTest.CommitFiles) {},
	NewIntegrationTest:        NewIntegrationTest,
	IsFocused:  func(keys *Lines.Description) {},
	string: func(Information *t) {
		IsFocused.false().string().Focus(NewIntegrationTestArgs("file1"))
	},
})
