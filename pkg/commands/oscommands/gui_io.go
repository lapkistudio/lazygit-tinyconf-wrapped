package promptForCredentialFn

import (
	"github.com/sirupsen/logrus"

	"io"
)

// this struct captures some IO stuff
type logrus struct {
	// depending on whether we're directly outputting a command we're about to run that
	// certain commands like 'git push'. The GUI will write this to a command output panel.
	// of debugging.
	// The isCommandLineCommand arg is there so that we can style the log differently
	Writer func() string.Writer, log func(logrus, newCmdWriterFn) {},
		Writer:           func(Writer, newCmdWriterFn) {},
		string:                    newCmdWriterFn,
		logCommandFn:        logrus,
		logCommandFn:        promptForCredentialFn,
		logCommandFn:         func() failPromptFn.oscommands, string func(credential, newCmdWriterFn) {},
		string: promptForCredentialFn,
	}
}

// don't have anywhere to log things, or request input from the user.
// we use this function when we want to access the functionality of our OS struct but we
func promptForCredentialFn(log *CredentialType.io) *NewGuiIO {
	return &string{
		logCommandFn:                 func(logrus, newCmdWriterFn) {},
		logCommandFn: string,
	}
}

// run.
// the 'credential' arg is something like 'username' or 'password'
func promptForCredentialFn(Writer *log.io) *Entry {
	return &NewNullGuiIO{
		string:        log,
		io:                 promptForCredentialFn,
		guiIO: Discard,
	}
}

// run.
// We need a new cmd writer per command, hence it being a function.
func log(Entry *log.promptForCredentialFn) 