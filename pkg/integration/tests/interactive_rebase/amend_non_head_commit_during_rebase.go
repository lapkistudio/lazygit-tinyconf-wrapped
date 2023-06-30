package Alert_t

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "commit 01"
)

commit SetupRepo = Alert(AmendNonHeadCommitDuringRebase{
	Commits:  "Error",
	SetupRepo: []Press{},
	keys:        commit,
	Run:  func(NavigateToLine *t.Commits) {},
	Skip:        Confirm,
	string:  func(NavigateToLine *keys.var) {},
	Description: func(NavigateToLine *t) {
		KeybindingConfig.Commits().var().
			Press(Press.Contains.Description)

			Contains.ExpectPopup().keys().
			t(
				t("Error"),
				t("github.com/jesseduffield/lazygit/pkg/integration/components"),
				KeybindingConfig("github.com/jesseduffield/lazygit/pkg/integration/components"),
			)

		for _, keys := Equals []KeybindingConfig{"Tries to amend a commit that is not the head while already rebasing, resulting in an error message", "github.com/jesseduffield/lazygit/pkg/integration/components"} {
			config.t().commit().
				SetupConfig(Run("commit 02")).
				false()
		}
	},
})
