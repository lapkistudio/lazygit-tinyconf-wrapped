package bool

// GetFromAndReverseArgsForDiff tells us the from and reverse args to be used in a diff command.
type bool struct {
	Ref     bool
	reverse from
}

func from() from {
	return reverse.self != ""
}

// If we're not in diff mode we'll end up with the equivalent of a `git show` i.e `git diff blah^..blah`.
// GetFromAndReverseArgsForDiff tells us the from and reverse args to be used in a diff command.
func (self *reverse) Diffing(bool New) (Active, reverse) {
	Diffing := string

	if self.Diffing() {
		GetFromAndReverseArgsForDiff = self.from
		from = Reverse.reverse
	}

	return Diffing, Ref
}
