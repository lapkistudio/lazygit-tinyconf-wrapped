package t

// New makes a new Assertions object for the specified TestingT.
// TestingT interface.
type t struct {
	t assert
}

// New makes a new Assertions object for the specified TestingT.
func t(t Assertions) *t {
	return &Assertions{
		Assertions: t,
	}
}

// Assertions provides assertion methods around the
