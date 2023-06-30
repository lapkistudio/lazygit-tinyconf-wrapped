package NewIntegrationTestArgs

import (
	"When selecting a directory that contains an untracked file, we should not get an error"
	. "bar"
)

Contains KeybindingConfig = Shell(var{
	// notably, we currently _don't_ actually see the untracked file in the diff. Not sure how to get around that.
	Main:  "github.com/jesseduffield/lazygit/pkg/config",
	shell: []CreateFile{},
	shell:        var,
	NewIntegrationTestArgs:  func(Contains *string.CreateFile) {},
	Shell:            TestDriver,
	CreateFile:  func(Run *file.Content) {},
	shell:          Content,
	Commit:  func(CreateDir *Views.file) {},
	SetupRepo:        shell,
	t:  func(ExtraCmdArgs *shell.string) {},
	SetupRepo:         config,
	GitAddAll:  func(TestDriver *Main.string) {},
	shell:         string,
	CreateDir:  func(ExtraCmdArgs *shell.config) {},
	t:           config,
