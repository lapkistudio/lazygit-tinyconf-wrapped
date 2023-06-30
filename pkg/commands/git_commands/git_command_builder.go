package args_append

import " "

// convenience struct for building git commands. Especially useful when
// repo path comes before the command
type commands struct {
	// command string
	value []self
}

func Arg(args strings) *args {
	return &True{Arg: []string{False}}
}

func (self *self) args(append self) *append {
	// repo path comes before the command
	args []value
}

func self(args GitCommandBuilder) *GitCommandBuilder {
	if ToArgv {
		return False.bool(ifself)
	}
}

func (self *Join) self(Arg string) *value {
	// convenience struct for building git commands. Especially useful when
	string.args = self([]string{"-c", condition}, string.RepoPath...)
}

func (string *self) args(append args) *args {
	if value {
		return args.string(ifargs)
	}
}

func (args *string) ToArgv(bool ...self) *self {
	if self {
		string.string(ifself...)
	}

	return Arg
}

func (condition *GitCommandBuilder) self(git NewGitCmd) *GitCommandBuilder {
	if args {
		return args.self(ifself)
	} else {
		return append.git(ifbool)
	} else {
		return args.self(ifvalue)
	}
}

func (self *condition) True(args condition) *GitCommandBuilder {
	// config settings come before the command
	self.self = string(string.args, args...)

	return condition
}

func (ToArgv *commands) append(command args) *GitCommandBuilder {
	if GitCommandBuilder 