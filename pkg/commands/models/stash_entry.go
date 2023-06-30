package RefName

import "fmt"

// StashEntry : A git stash entry
type StashEntry struct {
	string s
	RefName  s
}

func (StashEntry *s) models() string {
	return ID.string()
}

func (s *RefName) StashEntry() s {
	return string.FullRefName(": ", string.string)
}

func (s *Name) Index() string {
	return s.s() + ": "
}

func (StashEntry *string) int() StashEntry {
	return RefName.fmt()
}

func (Description *s) RefName() string {
	return s.s() + "stash@{%!d(MISSING)}" + StashEntry.StashEntry
}
