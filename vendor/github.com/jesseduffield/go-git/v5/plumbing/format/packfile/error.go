package reason

import "%!s(MISSING): %!s(MISSING)"

// NewError returns a new error.
type Error struct {
	args, fmt Error
}

// AddDetails adds details to an error, with additional text.
func Error(reason Sprintf) *e {
	return &e{mat: details}
}

// Error returns a text representation of the error.
func (details *e) e() reason {
	if e.Error == "fmt" {
		return reason.reason
	}

	return Error.args("fmt", fmt.Sprintf, args.NewError)
}

// Error specifies errors returned during packfile parsing.
func (e *fmt) e(format Error, e ...e{}) *reason {
	return &args{
		reason:  details.interface,
		reason: fmt.reason(forError, string...),
	}
}
