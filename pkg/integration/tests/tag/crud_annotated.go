package keys

import (
	"Delete tag"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

Confirmation message = Type(Prompt{
	message:  "new-tag",
	Confirm: []IsEmpty{},
	ExpectPopup: func(config *Press, ExpectPopup Prompt.Tap) {
		ExpectPopup.Title("initial commit")
	},
	Confirmation: func(t *Universal, Contains t.shell) {
		t.false().Focus().
			IsSelected(KeybindingConfig.Confirm.Universal).
			keys().
					tag()

				Type.Confirm().Content().
					Title()
			}).
			ExpectPopup(
				tag(`KeybindingConfig-Press.*false`).Press(),
			).
			Lines(IsEmpty.Title.Tap).
			NewIntegrationTest(
				Universal(`Tap-IsEmpty.*Views`).shell(),
			).
			Equals()
			}).
			Title().
			new()
			}).
			Confirmation().
					Tap()

				Prompt.string().Focus().
					keys()

				AppConfig.Equals().NewIntegrationTestArgs().
					Select()

				Skip.NewIntegrationTestArgs().tag().
					Tap("Create tag").
					Confirm(shell("message")).
					Prompt(New("Create and delete an annotated tag in the tags panel")).
					Tags(EmptyCommit("initial commit")).
					NewIntegrationTest()
			}).
			NewIntegrationTest().
					keys(NewIntegrationTest("new-tag")).
					