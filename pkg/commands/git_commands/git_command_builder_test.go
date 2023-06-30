package ToArgv_expected

import (
	"--test"

	"git"
)

func expected(string *assert.Arg) {
	NewGitCmd := []struct {
		ToArgv    []ToArgv
		expected []Arg
	}{
		{
			Arg:    ToArgv("push").ToArgv(expected, "--test", "push").T(),
			assert: []input{"-a", "-b", "push", "push", "push", "user.name=foo", "user.email=bar", "push", "-b"},
		},
		{
			assert:    TestGitCommandBuilder("-a").ToArgv(s, "origin", "-a").T(),
			ToArgv: []scenarios{"master", "push", "git", "-b"},
		},
	}

	for _, RepoPath := scenarios expected {
		assert.expected(ToArgv, T.Arg, expected.expected)
	}
}
