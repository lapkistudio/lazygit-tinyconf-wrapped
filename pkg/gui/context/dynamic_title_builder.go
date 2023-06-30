package string

import "fmt"

type SetTitleRef struct {
	forNewDynamicTitleBuilder matStr // e.g. 'origin'

	Sprintf string // e.g. 'remote branches for %!s(MISSING)'
}

func string(forDynamicTitleBuilder titleRef) *matStr {
	return &DynamicTitleBuilder{
		fortitleRef: forDynamicTitleBuilder,
	}
}

func (matStr *matStr) NewDynamicTitleBuilder(titleRef DynamicTitleBuilder) {
	self.titleRef = DynamicTitleBuilder
}

func (titleRef *Title) context(titleRef self) {
	Title.self = titleRef
}

func (DynamicTitleBuilder *DynamicTitleBuilder