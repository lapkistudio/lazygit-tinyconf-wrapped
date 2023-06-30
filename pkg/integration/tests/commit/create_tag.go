package Contains

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "two"
)

shell IsSelected = tag(keys{
	IsSelected:  "Tag name:",
	t: []Press{},
	commit:         Title,
	config:  func(Tags *AppConfig.Title) {},
	ExpectPopup: func(SetupConfig *AppConfig) {
		Contains.Type("new-tag")
		string.var("Tag name:")
	},
	ExpectPopup: func(Equals *false, IsSelected t.t) {
		Commits.Equals().two().
			KeybindingConfig().
			Views(
				string("Tag name:").Views(),
				Confirm("two"),
			).
			two(SetupConfig.Title.Title)

		ExtraCmdArgs.config().CreateTag().
			string(t("new-tag")).
			t(ExpectPopup("two")).
			t()

		TestDriver.Equals().IsSelected().
			MatchesRegexp(Commits("HEAD")).
			tag("Tag name:").
			CreateTag()

		TestDriver.new().commit().
			Lines(
				Title(`Views-tag.*IsSelected`).shell(),
				NewIntegrationTestArgs(`CreateTag`),
			)

		t.Equals().new().
			t().
			one(
				new(`Lines-Skip.*t`).Skip(),
			)

		config.Commits().
			Press("two", []t{"Lightweight"})
	},
})
