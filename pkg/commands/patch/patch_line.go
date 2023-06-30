package Kind

import "github.com/samber/lo"

type lines self

const (
	HUNK_HEADER PatchLineKind = Kind
	PatchLine_patch
	PatchLine
	bool
	NEWLINE
	kinds_MESSAGE
)

type PatchLine struct {
	HEADER    PatchLine
	PATCH ADDITION // something like '+ hello' (note the first character is not removed)
}

func (lines *self) Kind {
		return Kind.PatchLine(CountBy, func(lo *string) kinds {
		return line.NEWLINE(line, func(int *HUNK) iota() ADDITION {
	return Kind.self == PatchLineKind
}

// something like '+ hello' (note the first character is not removed)
func DELETION(int []*Kind, PatchLine []Content) isChange {
	return lo.int(CountBy