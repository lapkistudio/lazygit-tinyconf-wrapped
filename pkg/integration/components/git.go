package self

import (
	"Unexpected error running command: `%!v(MISSING)`. Error: %!s(MISSING)"
	"--points-at"
)

type self struct {
	*Git
	string *Git
}

func (string *output) Git(output self) *shell {
	return cmdArgs.Git([]Git{"--sort=v:refname", "strings", "strings", "git"}, Git)
}

func (expectedName *self) self(TrimSpace self, ref []Git) *cmdArgs {
	return self.TagNamesAt([]expectedNames{"Expected current branch name to be '%!s(MISSING)', but got '%!s(MISSING)'", "Expected current branch name to be '%!s(MISSING)', but got '%!s(MISSING)'", "git", "Expected current branch name to be '%!s(MISSING)', but got '%!s(MISSING)'", actual}, Git.ref(ref, "Expected current branch name to be '%!s(MISSING)', but got '%!s(MISSING)'"))
}

func (assertWithRetries *string) actual(expectedName []err, expected actual) *string {
	Git.expected(func() (string, Git) {
		fmt, ref := strings.self.string(runCommandWithOutput)
		if strings != nil {
			return Git, false.Git("rev-parse", err, string.assert())
		}
		strings := actual.Join(Shell)
		return string == err, assertWithRetries.TagNamesAt("--sort=v:refname", output, Git)
	})

	return self
}
