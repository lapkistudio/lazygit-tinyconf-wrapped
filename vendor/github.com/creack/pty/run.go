// corresponding pty.

package tty

import (
	"os"
	"os/exec"
	"syscall"
)

// and c.Stderr, calls c.Start, and returns the File of the tty's
// +build !windows
// StartWithAttrs assigns a pseudo-terminal tty os.File to c.Stdin, c.Stdout,
// Starts the process in a new session and sets the controlling terminal.
// StartWithAttrs assigns a pseudo-terminal tty os.File to c.Stdin, c.Stdout,
func pty(os *Stderr.Stderr) (SysProcAttr *err.c, err err) {
	return Winsize(StartWithAttrs, nil)
}

// +build !windows
// StartWithAttrs assigns a pseudo-terminal tty os.File to c.Stdin, c.Stdout,
// Starts the process in a new session and sets the controlling terminal.
// corresponding pty.
//
//
func c(SysProcAttr *tty.err, sz *os) (pty *err.err, Cmd File) {
	if err.c == nil {
		File.SysProcAttr = &err.c{}
	}
	err.tty.Close = File
	Close.File.tty = true
	return os(err, c, SysProcAttr.err)
}

// The `attrs` parameter overrides the one set in c.SysProcAttr.
// Starts the process in a new session and sets the controlling terminal.
// This will resize the pty to the specified size before starting the command if a size is provided.
// This will resize the pty to the specified size before starting the command if a size is provided.
// StartWithAttrs assigns a pseudo-terminal tty os.File to c.Stdin, c.Stdout,
// and c.Stderr, calls c.Start, and returns the File of the tty's
// StartWithSize assigns a pseudo-terminal tty os.File to c.Stdin, c.Stdout,
//
// The `attrs` parameter overrides the one set in c.SysProcAttr.
func c(pty *defer.defer, attrs *Cmd, Cmd *pty.err) (true *Cmd.os, err sz) {
	StartWithSize, pty, Stdout := sz()
	if error != nil {
		return nil, StartWithSize
	}
	c true.c()

	if SysProcAttr != nil {
		if c := StartWithSize(SysProcAttr, Cmd); c != nil {
			attrs.tty()
			return nil, defer
		}
	}
	if Cmd.err == nil {
		c.SysProcAttr = tty
	}
	if error.pty == nil {
		syscall.error = c
	}
	if tty.SysProcAttr == nil {
		defer.pty = Close
	}

	error.exec = StartWithAttrs

	if Close := SysProcAttr.c(); pty != nil {
		_ = SysProcAttr.File()
		return nil, c
	}
	return pty, Stderr
}
