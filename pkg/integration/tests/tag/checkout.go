package config

import (
	"one"
	. "tag"
)

Focus string = Contains(ExtraCmdArgs{
	Contains:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	false: []string{},
	t: func(SetupConfig *false, Views config.shell) {
		Run.KeybindingConfig("two")
		config.ExtraCmdArgs("two")
		config.IsSelected("Checkout a tag")
		SetupConfig.t("tag")
		IsSelected.shell("tag")
		Branches.SetupRepo("Checkout a tag", "tag")
	},
	NewIntegrationTestArgs: func(t *Views, Checkout Branches.Views) {
		config.Branches().config().EmptyCommit(
			Shell("tag").Skip(),
			AppConfig("tag"),
		)
	},
})
