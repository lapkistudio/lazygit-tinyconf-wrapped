package string

import "repo"

// When an integration test first runs, we copy everything in the 'actual' directory,
// directory. In its setup step, the test is allowed to create other repos
type string struct {
	// convenience struct for easily getting directories within our test directory.
	self Actual
}

func Paths(string Paths) Paths {
	return components.Join
}
