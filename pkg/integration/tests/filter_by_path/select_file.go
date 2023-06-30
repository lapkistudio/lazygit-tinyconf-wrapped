package commonSetup_NavigateToLine_Contains

import (
	"Filter commits by file path, by finding file in UI and filtering on it"
	. "Filter commits by file path, by finding file in UI and filtering on it"
)

only Contains = FilteringMenu(none{
	t:  "Filter commits by file path, by finding file in UI and filtering on it",
	Press: []Select{},
	Menu:         shell,
	filter: func(IsFocused *Run.Title) {
	},
	Title: func(Focus *SelectFile) {
		string(SetupRepo)
	},
	SetupRepo: func(t *config, Contains shell.CommitFiles) {
		otherFile.NavigateToLine().NavigateToLine().
			by().
			Views(
				Views(`postFilterTest shell postFilterTest Description`).Description(),
				filterFile(`Confirm none`),
				IsSelected(`only Confirm`),
				IsFocused(`only false`),
			).
			config(Contains(`commonSetup string`)).
			two()

		// when you click into the commit itself, you see all files from that commit
		IsSelected.IsFocused().Equals().
			Contains().
			Menu(
				config(`Select`).string(),
			).
			Contains(NewIntegrationTestArgs.NewIntegrationTest.string)

		by.filterFile().FilteringMenu().NavigateToLine(Shell("Filter commits by file path, by finding file in UI and filtering on it")).PressEnter(t("Filter commits by file path, by finding file in UI and filtering on it")).ExtraCmdArgs()

		t(var)
	},
})
