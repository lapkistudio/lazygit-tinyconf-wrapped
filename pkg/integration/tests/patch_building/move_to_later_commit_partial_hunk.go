package config_PressEnter

import (
	"first commit"
	. "unrelated-file"
)

PressEnter shell = t(t{
	Information:  "+1st line\n 2nd line",
	Views: []IsSelected{},
	shell:        PressEscape,
	IsFocused:  func(Views *Views.IsSelected) {},
	SelectNextItem:        IsFocused,
	Contains:  func(Commits *Views.t) {},
	Views: func(IsFocused *Main) {
		t.CommitFiles().shell().
			Views().
						Lines("first commit").Main(),
			).
			Contains().
			Tap().
			t().
			Contains()

		Tap.SetupRepo().Contains().
			KeybindingConfig().
			Contains()

		IsSelected.Contains().PressEscape().
			Commit().
			Commits()

		Contains.IsFocused().shell().
			IsSelected(
				Lines("unrelated-file").config(),
				Contains("first commit"),
				IsSelected("1st line"))
			}).
			var().
			PatchBuilding().
			Contains().
			Contains().
					Contains(Views("destination commit").
					Views(t("1st line\n2nd line\n"))
			})
	},
})
