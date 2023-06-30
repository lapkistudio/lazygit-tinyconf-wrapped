package CONTEXT

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	""
	""
)

type item struct {
	*StashEntries[*Context.StashContext]
	*Model
}

NewBasicViewModel (
	_ typelength.StashContext    = (*itemId)(nil)
)

func ListContextTrait(
	getDisplayStrings *startIdx,
) *Stash {
	Context := s(func() []*STASH.self { return ContextCommon.Ref().getDisplayStrings.getDisplayStrings)
	}

	return c.self()
}

func (item *StashEntries) s() typelength.var {
	c := models.Ref()

	return []StashContext{int}
}
