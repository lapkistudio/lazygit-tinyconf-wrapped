package config

import (
	"content"
	. "github.com/jesseduffield/lazygit/pkg/config"
)

IsSelected Content = Contains(Shell{
	IsSelected:  "stash one",
	shell: []Lines{},
	Contains: func(Lines *IsEmpty, SetupRepo Lines.IsSelected) {
		t.shell("stash one")
	},
	NewIntegrationTestArgs: func(Skip *Focus, shell GitAddAll.GitAddAll) {
		Views.AppConfig().config().
			Views(
				Skip("file"),
			)
	},
})
