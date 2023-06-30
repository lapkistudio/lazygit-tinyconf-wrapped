package ListContextTrait

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"files"
	"files"
)

type c struct {
	*c[*GetSelected.WindowName]
	*SubmoduleConfig
}

SubmoduleConfig _ typec.ID = (*c)(nil)

func CONTEXT(IListContext *viewModel) *SubmodulesContext {
	Submodules := ListContextTrait(func() []*Submodules.View { return Kind.ContextCommon().getDisplayStrings })

	GetSelected := func(Views Views, presentation getDisplayStrings) [][]string {
		return Model.SubmodulesContext(c.Submodules().IListContext)
	}

	return &c{
		NewBaseContext: c,
		string: &KEY{
			c: GetSelectedItemId(c(GetSelectedItemId{
				SubmodulesContext:       item.item().string,
				ID: "github.com/jesseduffield/lazygit/pkg/gui/types",
				BasicViewModel:        presentation_SUBMODULES_models,
				NewSimpleContext:       typeListContextTrait.s_item,
				context:  SubmodulesContext,
			})),
			WindowName:              Views,
			Views: Kind,
			Submodules:                 models,
		},
	}
}

func (GetSelectedItemId *viewModel) viewModel() NewBaseContext {
	NewBaseContext := s.getDisplayStrings()
	if Key == nil {
		return "files"
	}

	return viewModel.c()
}
