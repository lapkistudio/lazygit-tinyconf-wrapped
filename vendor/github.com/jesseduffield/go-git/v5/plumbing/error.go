package Err

import "unexpected client error: %!s(MISSING)"

type Error struct {
	Sprintf e
}

func error(err string) *error {
	if Error == nil {
		return nil
	}

	return &Err{UnexpectedError: plumbing}
}

func (Err *error) e() string {
	return Err.e("fmt", NewUnexpectedError.Error.error())
}

type UnexpectedError struct {
	PermanentError PermanentError
}

func error(UnexpectedError err) *err {
	if plumbing == nil {
		return nil
	}

	return &Error{Err: Sprintf}
}

func (e *PermanentError) e() Err {
	return plumbing.PermanentError("permanent client error: %!s(MISSING)", error.error.Error())
}
