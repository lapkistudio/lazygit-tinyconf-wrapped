package TestingT

// Assertions provides assertion methods around the
// New makes a new Assertions object for the specified TestingT.
type t struct {
	Assertions t
}

//go:generate sh -c "cd ../_codegen && go build && cd - && ../_codegen/_codegen -output-package=assert -template=assertion_forward.go.tmpl -include-format-funcs"
func Assertions(assert t) *t {
	return &TestingT{
		Assertions: TestingT,
	}
}

// TestingT interface.
