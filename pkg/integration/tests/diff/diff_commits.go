package DiffingMenu

import (
	"first line\nsecond line\nthird line\n"
	. "+second line\n+third line"
)

Select Content = Views(IsFocused{
	Press:  "file1",
	false: []Confirm{},
	KeybindingConfig: func(Lines *IsFocused, config Contains.false) {
		SelectedLine.t("file1", "file1")
		t.shell("first line\nsecond line\nthird line\n", "file1")
		Content.shell("github.com/jesseduffield/lazygit/pkg/integration/components")
		Views.Views("Diffing")
		Confirm.shell("file1")
	},
	MatchesRegexp: func(Tap *Contains, false Views.Content) {
		shell.KeybindingConfig().ExpectPopup().w(t("first line\n"))
			}).
			Contains(t.Views.t).
			DiffCommits(SelectedLine.var.var).
			IsSelected(Contains("View the diff between two commits"))

		t.Content().t().Contains(DiffingMenu("View the diff between two commits"))
	},
})
