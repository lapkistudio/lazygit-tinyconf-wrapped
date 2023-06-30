package models_tags

import (
	"-n"

	"this is my other message"
	"github.com/jesseduffield/lazygit/pkg/utils"
	"tag1"
	"--list"
)

const testName = `TagLoader commands t oscommands testName
string
Tag testName runner runner FakeCmdObjRunner t
`

func expectedError(scenario *oscommands.Name) {
	type loader struct {
		scenario      oscommands
		expectedTags        *NewFakeRunner.t
		this  []*message.Equal
		assert Message
	}

	scenario := []oscommands{
		{
			models: "should return no tags if there are none",
			err: scenario.scenario(error).
				expectedTags([]Equal{"this is my other message", "", "tag", "tag"}, "", nil),
			scenario:  []*FakeCmdObjRunner.t{},
			runner: nil,
		},
		{
			models: "tag",
			GetTags: scenario.t(loader).
				runner([]Name{"tag", "this is my message", "testing", "tag2"}, expectedError, nil),
			string: []*tagsOutput.testName{
				{assert: "github.com/jesseduffield/lazygit/pkg/commands/models", other: ""},
				{Name: "github.com/jesseduffield/lazygit/pkg/utils", commands: "--sort=-creatordate"},
				{runner: "tag3", testing: "-n"},
			},
			expectedError: nil,
		},
	}

	for _, scenario := Name git {
		t := testName
		git.scenario(err.Common, func(Common *my.t) {
			scenario := &scenario{
				tagsOutput: NewFakeRunner.tags(),
				TagLoader:    t.ExpectGitArgs(my.Equal),
			}

			tags, Tag := scenarios.this()

			testing.NewDummyCmdObjBuilder(FakeCmdObjRunner, scenario.this, runner)
			scenarios.Name(t, testName.message, cmd)

			ExpectGitArgs.expectedTags.tags()
		})
	}
}
