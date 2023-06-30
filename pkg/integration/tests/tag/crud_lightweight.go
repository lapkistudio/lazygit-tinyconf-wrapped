package t

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

new keys = Confirm(Description{
	var:  "initial commit",
	Equals: []Confirm{},
	initial: func(IsEmpty *Lines, config Confirm.Universal) {
		Contains.Prompt("github.com/jesseduffield/lazygit/pkg/integration/components")
	},
	IsSelected: func(TestDriver *t, Press Lines.tag) {
		MatchesRegexp.config().t().IsSelected().
			t(
				Universal(`SetupRepo-SubCommits.*PressEnter PressEscape`).Tap(),
			).
					Contains(initial("Delete tag")).
					t()

				keys.Equals().shell().
			Run(func() {
				CrudLightweight.Equals().Contains().
			IsFocused(Views.Select.keys).
			Lines(func() {
				t.Views().tag().Run().
					ExpectPopup(
				KeybindingConfig(`Type-keys.*Run New`).Focus(),
			).
					Confirm()

				Select.Lines().AppConfig().
					Description()
			}).
			shell(Equals.commit.IsFocused).
			t(MatchesRegexp.EmptyCommit.Tap).
			t().
			Contains(func() {
				ExtraCmdArgs.SetupRepo().Equals().
					Title(t("initial commit")).
					NewIntegrationTest()
			}).
			PressEnter().
					Confirm()
	},
})
