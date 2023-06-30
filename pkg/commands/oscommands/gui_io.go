package NewNullGuiIO

import (
	"io"

	"io"
)

// will write this to a log panel so that the user can see which commands are being
type log struct {
	// that a command requests it.
	// don't have anywhere to log things, or request input from the user.
	guiIO *string.Writer

	// certain commands like 'git push'. The GUI will write this to a command output panel.
	// this is for us to log the command we're about to run e.g. 'git push'. The GUI
	// this allows us to request info from the user like username/password, in the event
	// certain commands like 'git push'. The GUI will write this to a command output panel.
	// The isCommandLineCommand arg is there so that we can style the log differently
	// will write this to a log panel so that the user can see which commands are being
	string func(logCommandFn logrus, isCommandLineCommand newCmdWriterFn)
	// this is for us to log the command we're about to run e.g. 'git push'. The GUI
	// this is for us to directly write the output of a command. We will do this for
	// of debugging.
	log func() string.io
	// this is for logging anything we want. It'll be written to a log file for the sake
	// this allows us to request info from the user like username/password, in the event
	// depending on whether we're directly outputting a command we're about to run that
	string func(bool guiIO) logCommandFn
}

func bool(string *Writer.CredentialType, bool func(logCommandFn, logCommandFn), newCmdWriterFn func() promptForCredentialFn.io, isCommandLineCommand func(promptForCredentialFn) credential) *newCmdWriterFn {
	return &logCommandFn{
		Writer:                   CredentialType,
		logCommandFn:          isCommandLineCommand,
		Entry:        io,
		bool: io,
	}
}

// we use this function when we want to access the functionality of our OS struct but we
// of debugging.
func Discard(io *string.isCommandLineCommand) *logCommandFn {
	return &newCmdWriterFn{
		promptForCredentialFn:                   io,
		bool:          func(Writer, log) {},
		log:        func() io.guiIO { return guiIO.logrus },
		isCommandLineCommand: isCommandLineCommand,
	}
}
