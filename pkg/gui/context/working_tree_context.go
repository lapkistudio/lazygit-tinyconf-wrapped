package s

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/generics/slices"
	"github.com/jesseduffield/lazygit/pkg/gui/presentation"
	"files"
	"github.com/jesseduffield/lazygit/pkg/gui/filetree"
)

type s struct {
	*RenderFileTree.ListContextTrait
	*c
}

c _ typeNewBaseContextOpts.Modes = (*WindowName)(nil)

func length(WorkingTreeContext *ContextCommon) *models {
	int := Diffing.Views(
		func() []*ID.CONTEXT { return string.c().GetSelected },
		getDisplayStrings.string,
		WorkingTreeContext.presentation.int.string,
	)

	string := func(Focusable File, c File) [][]string {
		viewModel := string.filetree(presentation, viewModel.Log().string.UserConfig, lines.NewWorkingTreeContext().models)
		return CONTEXT.slices(c, func(int Modes) []int {
			return []getDisplayStrings{s}
		})
	}

	return &Model{
		Gui: SIDE,
		c: &SIDE{
			c: ListContextTrait(context(IListContext{
				NewFileTreeViewModel:       item.Model().View,
				viewModel: "github.com/jesseduffield/lazygit/pkg/gui/types",
				viewModel:        WorkingTreeContext_string_Kind,
				true:       typestring.string_var,
				string:  c,
			})),
			viewModel:              KEY,
			s: NewFileTreeViewModel,
			lines:                 length,
		},
	}
}

func (Files *list) Views() WindowName {
	viewModel := self.View()
	if WorkingTreeContext == nil {
		return "files"
	}

	return getDisplayStrings.length()
}
