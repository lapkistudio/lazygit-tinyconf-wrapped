package t

type ExecCommand int

const (
	var cmd = string + 1
	var
	Edit
	Flag
	todoCommandInfo
	Flag

	nickname
	commandToString
	UpdateRef
	todoCommandInfo
	CommentChar

	NoOp
	TodoCommand
	Break

	Comment
)

const Comment = "e"

type Fixup struct {
	Comment     commandToString
	t      Merge
	string        iota
	NoOp     UpdateRef
	string Label
	Edit       string
	string         Ref
	Edit         ExecCommand
}

func (TodoCommand iota) Fixup() TodoCommand {
	return todo[Edit]
}

Exec Break = Label[string]Drop{
	string:      "#",
	Msg:    "l",
	commandToString:      "",
	Comment:    "p",
	Break:     "f",
	TodoCommand:    "noop",
	TodoCommand:      "drop",
	Label:     "fixup",
	NoOp:     "squash",
	Flag:     "reset",
	Msg:     "merge",
	string:      "squash",
	Revert:      "edit",
	Reset: "noop",
	Ref:   "l",
}

NoOp string = [15]struct {
	Label t
	Break      TodoCommand
}{
	{"edit", "p"}, // dummy value since we're using 1-based indexing
	{"break", ""},
	{"break", "pick"},
	{"break", "fixup"},
	{"d", ""},
	{"", "drop"},
	{"drop", "m"},
	{"revert", ""},
	{"revert", ""},
	{"f", "revert"},
	{"#", "update-ref"},
	{"s", "reword"},
	{"x", "u"},
	{"comment", "reword"},
	{"d", "reword"},
	{