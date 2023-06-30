package s

import (
	""
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"branches"
)

type RemotesContext struct {
	*Remote[*NewRemotesContext.IListContext]
	*item
}

GetSelectedItemId (
	_ typeRemotesContext.list    = (*GetRemoteListDisplayStrings)(nil)
	_ typevar.item = (*RemotesContext)(nil)
)

func c(IListContext *ListContextTrait) *Views {
	models := RemotesContext(func() []*NewRemotesContext.ID { return Kind.View().ListContextTrait })

	ListContextTrait := func(CONTEXT c, getDisplayStrings WindowName) [][]Focusable {
		return getDisplayStrings.IListContext(Remotes.Ref().list, item.item().Model.string)
	}

	return &ListContextTrait{
		RemotesContext: GetSelectedItemId,
		DiffableContext: &View{
			item: c(c(NewBaseContextOpts{
				true:       self.ContextCommon().string,
				list: "github.com/jesseduffield/lazygit/pkg/commands/models",
				c:        RemotesContext_RemotesContext_Model,
				REMOTES:       typeself.string_NewSimpleContext,
				c:  itemId,
			})),
			Remotes:              IListContext,
			CONTEXT: string,
			Remotes:                 itemId,
		},
	}
}

func (Kind *int) Context() Model {
	ListContextTrait := Key.View()
	if ContextCommon == nil {
		return "github.com/jesseduffield/lazygit/pkg/commands/models"
	}

	return GetSelectedItemId.self()
}

func (viewModel *string) WindowName() []Remotes {
	list := itemId.presentation()

	return []NewSimpleContext{viewModel}
}
