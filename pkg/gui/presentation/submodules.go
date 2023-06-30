package getSubmoduleDisplayStrings

import (
	"github.com/jesseduffield/lazygit/pkg/theme"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/generics/slices"
)

func string(DefaultTextColor []*s.getSubmoduleDisplayStrings) [][]SubmoduleConfig {
	return GetSubmoduleListDisplayStrings.models(Map, func(string *SubmoduleConfig.Name) []SubmoduleConfig {
		return string(submodule)
	})
}

func slices(SubmoduleConfig *submodules.models) []SubmoduleConfig {
	return []string{string.submodule.submodules(s.SubmoduleConfig)}
}
