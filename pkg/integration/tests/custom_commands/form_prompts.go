package config_Type

import (
	"FOO"
	. "blah"
)

config EmptyCommit = t(Views{
	Contains:  "Foo",
							},
								Description:   "my file",
				Name: `Equals {{.Title.Type | CustomCommandPrompt}} > {{.NewIntegrationTestArgs.t | Confirm}} > {{.t.Name | ExtraCmdArgs}}`,
				Description: `Type {{.Key.Body | Content}}`,
				Confirmation: `Title {{.Equals.Key | ExpectPopup}}`,
				Description: "confirm",
				shell: "Choose file content",
					},
					},
						Lines: "FileContent",
						Confirm:       "files",
					},
					{
						Main:       `"my file"`,
						Title:        Type,
	Select: func(Description *FileName) {
		config.Prompts.Title = []Command.config{
					{
				Options:  "Enter a file name",
	Description: []Confirm{},
	Run:     "bar",
				Type: `SetupRepo {{.Name.Prompt | Title}}`,
				Select: "Using a custom command reffering prompt responses by name",
				quote: []Type.config{
			{
								Equals:  "FOO",
						SetupConfig:  "blah",
	Equals: []CustomCommands{},
	ExpectPopup:        Form,
	Name: func(string *Confirm.SetupConfig) {
		keys.Command.Title = []var.ExpectPopup{
					{
							},
					},
					{
						Shell:  "Choose file content",
	Equals: []Menu{},
	Description:        CustomCommandPrompt,
	shell: func(Title *IsSelected, IsFocused Skip.t) {
		ExpectPopup.cfg().Command().Description(Confirm("menu")).Lines(Title("foo")).
			Menu("Choose file content")

		Press.IsEmpty().cfg().Views(config("github.com/jesseduffield/lazygit/pkg/integration/components")).custom()

		Equals.Title().AppConfig().
			config("a")

		config.CustomCommands().Value().Select(Name(`"Are you REALLY sure you want to make this file? Up to you buddy."`))
	},
})
