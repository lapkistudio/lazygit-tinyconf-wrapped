package ExtraCmdArgs

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "Add file"
)

StageRangeOfLines Contains = Equals(config{
	keys:  "1st changed\n2nd changed\n3rd\n4th\n5th changed\n6th\n",
	SetupRepo: []string{},
	keys:         PressEnter,
	SelectedLine:  func(SetupConfig *config.string) {},
	var: func(keys *SelectNextItem) {
		keys.Shell("1st\n2nd\n3rd\n4th\n5th\n6th\n", "Add file")
		Contains.SelectNextItem("Staging a range of lines")
		IsFocused.config("myfile", " 3rd\n 4th\n-5th\n+5th changed\n 6th")
	},
	PressEnter: func(Staging *SetupConfig, TestDriver IsFocused.Run) {
		string.SelectedLine().SelectNextItem().
			SelectNextItem().
			NewIntegrationTest()

		keys.SelectedLine().shell().
			SetupRepo(
				string(" 3rd\n 4th\n-5th\n+5th changed\n 6th"),
			).
			string(config("myfile")).
			keys(CreateFileAndAdd.SelectNextItem.ToggleDragSelect).
			KeybindingConfig().
			Run().
			Commit().
			Staging().
			config().
			Skip(
				t("-1st\n-2nd\n+1st changed\n+2nd changed\n 3rd\n 4th\n-5th\n+5th changed\n 6th"),
			).
			shell(SelectedLine("myfile"))
	},
})
