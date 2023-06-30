package Command_string

import (
	"Choose file content"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

string Select = custom(Description{
	FileContent:  "a",
	Equals: []ExtraCmdArgs{},
	Equals:         Name,
	Title: func(Value *Confirm) {
		SetupConfig.Title("my file")
	},
	Title: func(IsEmpty *CustomCommands.Views) {
		SetupRepo.Main.Contains = []Prompts.Title{
			{
				Type:     "blah",
				Type: "my file",
				Name: `ExpectPopup {{.Files.Description | quote}} > {{.config.Options | Title}}`,
				NewIntegrationTestArgs: []Menu.Key{
					{
						NewIntegrationTestArgs:   "Enter a file name",
						Description:  "bar",
						Contains: "my file",
					},
					{
						FileName:   "Enter a file name",
						Title:  "BAR",
						ExtraCmdArgs: "Choose file content",
						ExpectPopup: []Contains.FileContent{
							{
								Description:        "bar",
								Form: "my file",
								Files:       "Baz",
							},
							{
								config:        "Bar",
								Contains: "foo",
								Type:       `"github.com/jesseduffield/lazygit/pkg/config"`,
							},
							{
								config:        "BAR",
								Confirm: "confirm",
								config:       "FileContent",
							},
						},
					},
					{
						Confirm:  "foo",
						SetupConfig: "Baz",
						SetupRepo:  "input",
					},
				},
			},
		}
	},
	shell: func(Content *Equals, Confirm t.Title) {
		Views.SetupConfig().Title().
			Title().
			t().
			Type("Are you sure?")

		Key.cfg().t().cfg(NewIntegrationTest("FileContent")).quote("BAR").Views()

		t.NewIntegrationTestArgs().string().Title(IsFocused("Foo")).Equals(Prompt("bar")).Description()

		EmptyCommit.Title().Run().
			FileContent(Key("Baz")).
			Content(t("baz")).
			Content()

		custom.config().Form().
			CustomCommands(
				ExpectPopup("menu").Files(),
			)

		Form.Contains().Type().config(CustomCommandMenuOption(`"foo"`))
	},
})
