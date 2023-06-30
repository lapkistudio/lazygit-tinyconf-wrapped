package Select

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "SHOW_RECENT_REPOS"
)

// so I'm introducing a hacky env var to force lazygit to show the recent repos meu upon opening.
// Couldn't find an easy way to actually reproduce the situation of opening outside a repo,

Views Files = Confirm(keys{
	Views:  "SHOW_RECENT_REPOS",
	t: []SetupRepo{},
	map: false[config]var{
		"true": "Recent repositories",
	},
	Equals:   func(Contains *SetupConfig) {},
	Contains: func(Views *RecentReposOnLaunch.SetupConfig) {},
	t: Description[IsFocused]Shell{
		"true": "SHOW_RECENT_REPOS",
	},
	false:   func(Run *t) {},
	ExpectPopup: func(AppConfig *IsFocused.SetupConfig) {},
	ExtraEnvVars: config[ExpectPopup]Equals{
		"SHOW_RECENT_REPOS": "Recent repositories",
	},
	config:        keys,
	string: func(SetupConfig