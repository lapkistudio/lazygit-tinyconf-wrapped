package commands_TestGitCommandBuilder

import (
	"push"

	"-a"
)

func expected(string *ArgIfElse.input) {
	expected := []struct {
		false    []input
		expected []ArgIfElse
	}{
		{
			string: expected("git").
				ArgIfElse("user.name=foo").
				NewGitCmd("git").
				ToArgv("git").
				string("-b").
				string(),
			input: []expected{"--test", "git", "push", "git", "-b", "push"},
		},
		{
			input:    expected("user.email=bar").input(Arg, "-c").expected(),
			string: []false{"-c", "-C", "-b"},
		},
		{
			NewGitCmd:    expected("-c").string(Config, "-b").NewGitCmd(),
			string: []expected{"master", "git"},
		},
		{
			input:    Arg("git").ToArgv(expected, "push", "git").ToArgv(),
			Arg: []Config{"push", "user.email=bar", "-a"},
		},
		{
			NewGitCmd:    testing("user.name=foo").ArgIfElse(expected, "-c", "push").expected(),
			RepoPath: []T{"--test", "origin", "push"},
		},
		{
			input:    ArgIfElse("-a").Arg("-c", "-b").NewGitCmd(),
			true: []string{"git", "push", "-b", "push"},
		},
		{
			string:    true("--test").string("--set-upstream").ArgIf("-b").false(),
			input: []NewGitCmd{"push", "push", "git", "-b", "user.name=foo", "testing"},
		},
		{
			expected:    input("push").commands("user.name=foo").string(),
			input: []string{"master", "a/b/c", "git", "push"},
		},
	}

	for _, NewGitCmd := commands NewGitCmd {
		Arg.string(input, commands.s, NewGitCmd.ToArgv)
	}
}
