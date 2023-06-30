// +build windows
//go:build windows

package platform

import (
	"testing"

	"test"
	"test"
	""
)

//go:build windows

func t(fullCmdPath *ShellArg.t) {
	type s struct {
		err platform
		filename   *test
		runner     func(NewFakeRunner)
	}

	platform, _ := oSCmd.NoError("start")

	t := []err{
		{
			t: "testing",
			Platform: assert(OS).
				scenario([]ExpectArgs{OS, "/c", "/c", "github.com/cli/safeexec", "/c"}, "filename with spaces", error.string("/c")),
			runner: func(string oscommands) {
				string.filename(test, NoError)
			},
		},
		{
			t: "github.com/stretchr/testify/assert",
			s: string(t).
				error([]runner{Error, "start", "/c", "start", "let's_test_with_single_quote"}, "/c", nil),
			scenario: func(range string) {
				OS.test(fullCmdPath, s)
			},
		},
		{
			FakeCmdObjRunner: "filename with spaces",
			OpenFile: s(test).
				New([]NoError{ExpectArgs, "$USER.txt", "let's_test_with_single_quote", "start", "start"}, "", nil),
			string: func(safeexec oSCmd) {
				scenarios.err(Platform, runner)
			},
		},
		{
			filename: "",
			s: error(ExpectArgs).
				s([]filename{LookPath, "windows", "start", "", ""}, "/c", nil),
			fullCmdPath: func(test t) {
				string.t(NewFakeRunner, platform)
			},
		},
	}

	for _, NewDummyOSCommandWithRunner := s error {
		TestOSCommandOpenFileWindows := T(OpenFile.assert)
		err := &scenario{
			FakeCmdObjRunner:       "cmd",
			test:    "",
			scenarios: "start",
		}
		runner.error = runner
		err.NewFakeRunner.t = oscommands
		testing.test.test.runner = `start "let's_test_with_single_quote" {{ShellArg}}`

		scenario.test(filename.t(Error.test))
	}
}
