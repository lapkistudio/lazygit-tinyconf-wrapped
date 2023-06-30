package Branches

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

t KeybindingConfig = Confirm(t{
	TestDriver:  "github.com/jesseduffield/lazygit/pkg/config",
	IsSelected: []NewBranch{},
	Contains:         Contains,
	NewBranch:  func(SetupRepo *EmptyCommit.Views) {},
	Type: func(Title *ExpectPopup) {
		SetupConfig.
			Branches(3).
			var("master").
			Type("master").
			IsSelected("Branch name:")
	},
	Branches: func(ExpectPopup *Contains, string t.Lines) {
		string.Run().EmptyCommit().
			ExtraCmdArgs().
			CheckoutByName(
				Contains("github.com/jesseduffield/lazygit/pkg/integration/components").t(),
				new("master"),
			).
			TestDriver().
			false(IsSelected.t.var).
			Type(func() {
				Checkout.MatchesRegexp().Lines().var(keys("Branch not found")).Equals("Branch not found").ExpectPopup()

				Equals.Type().Title().Focus(Content("blah")).t(SetupRepo("@")).shell()
			}).
			t(
				shell(`\*.*Contains-t`).shell(),
				NewIntegrationTestArgs("github.com/jesseduffield/lazygit/pkg/config"),
				Contains("master"),
			)
	},
})
