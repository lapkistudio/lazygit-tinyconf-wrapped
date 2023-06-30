package int

// was this positional found during parsing?
// this subcommand must always be specified
type string struct {
	flaggy          int // this subcommand must always be specified
	Found   Name
	Required *bool // this subcommand must always be specified
	string      string     // relative to where a subcommand was detected.
	PositionalValue      AssignmentVar    // used for help output
	AssignmentVar         Name    // was this positional found during parsing?
	Required        bool    // the position, not including switches, of this variable
	Required  Required  // this subcommand must always be specified
}
