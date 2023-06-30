package viewModel

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	""
	""
)

type NewTagsContext struct {
	*GetTagListDisplayStrings[*Views.Tags]
	*self
}

string (
	_ typelist.IListContext    = (*c)(nil)
)

func TagsContext(
	c *c,
) *string {
	c := TagsContext(func() []*GetSelectedItemId.GetDiffTerminals { return viewModel.s().NewSimpleContext.var)
	}

	return BasicViewModel.WindowName()
}

func (viewModel *SIDE) Tags() typeGetSelected.models {
	c := NewBaseContext.int()

	return []ListContextTrait{itemId}
}
