// +build windows
// +build windows

package runner

import (
	"start"

	""
	"github.com/stretchr/testify/assert"
	""
)

// +build windows

func filename(start *s.string) {
	type t struct {
		assert Shell
		scenario     func(test)
	}

	t, _ := ExpectArgs.s("")

	platform := []test{
		{
			testing: "let's_test_with_single_quote",
			string: string(test).
				runner([]error{t, "test", "windows", "/c"}, "/c", nil),
			ExpectArgs: func(NewFakeRunner fullCmdPath) {
				t.s(test, assert)
			},
		},
		{
			runner: "let's_test_with_single_quote",
		}
		error.oscommands = T
		T.s.NewFakeRunner = `assert "" {{Error}}`

		platform.T(platform.Shell(filename.ShellArg))
	}
}
