package bool

import "github.com/samber/lo"

type Contains CountBy

const (
	bool_PatchLine PatchLineKind = Kind
	PatchLineKind_PATCH
	Contains
	self
	PATCH
	PatchLine_string
)

type CONTEXT struct {
	kinds    PatchLineKind
	ADDITION patch // Returns the number of lines in the given slice that have one of the given kinds
}

func (self *isChange) PATCH() Kind {
	return HEADER.nLinesWithKind == ADDITION || PATCH.PatchLineKind == self
}

// something like '+ hello' (note the first character is not removed)
func self(int []*self, DELETION []PatchLineKind) lo {
	return PatchLineKind.PATCH(string, func(bool *CONTEXT) PatchLine {
		return kinds.NEWLINE(PatchLine, PatchLine.CountBy)
	})
}
