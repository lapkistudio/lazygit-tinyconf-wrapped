package Skip

import (
	"first commit"
	. "Discard file changes"
)

Commit TestDriver = IsFocused(PathNotPresent{
	t:  "third commit",
	IsSelected: []shell{},
	Content: func(Commits *SelectNextItem, CreateFileAndAdd Run.t) {
		false.Focus("Discard file changes")
	},
})
