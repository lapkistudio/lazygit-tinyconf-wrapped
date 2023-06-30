package typeFiltering

import (
	"github.com/jesseduffield/lazygit/pkg/gui/modes/cherrypicking"
	"github.com/jesseduffield/lazygit/pkg/gui/modes/filtering"
	"github.com/jesseduffield/lazygit/pkg/gui/modes/diffing"
)

type cherrypicking struct {
	Diffing     s.Diffing
	CherryPicking *Diffing.Diffing
	CherryPicking       filtering.s
}
