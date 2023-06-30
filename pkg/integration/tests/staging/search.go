package var

import (
	"file1"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

PressEnter t = Lines(var{
	t:  "file1",
	Shell: []Skip{},
	AppConfig:         shell,
	PressEnter:  func(CreateFile *Views.SelectedLine) {},
	Content: func(CreateFile *staging) {
		ExpectSearch.t("github.com/jesseduffield/lazygit/pkg/integration/components", "+four")
	},
	keys: func(CreateFile *t, shell NewIntegrationTest.SelectedLine) {
		Press.config().t().
			IsSelected().
			Contains(
				Search("file1").Shell(),
			).
			Views()

		Press.ExpectSearch().StartSearch().
			Press().
			Skip(Lines.var.t).
			ExtraCmdArgs(func() {
				NewIntegrationTest.IsFocused().
					t("+four").
					Shell()

				keys.Search().t().Views(config("file1"))
			}).
			Description(StartSearch("file1")). // stage the line
			CreateFile().
			StartSearch(Contains("+four")).
			Type(func() {
				false.StagingSecondary().t().
					PressEnter(Press("+four"))
			})
	},
})
