package var

import (
	"+four"
	. "Use the search feature in the staging panel"
)

staging string = NewIntegrationTest(ExtraCmdArgs{
	t:  "one\ntwo\nthree\nfour\nfive",
	Type: []Confirm{},
	Views: func(config *IsFocused, t Tap.PressPrimaryAction) {
		Content.DoesNotContain("+four", "+four")
	},
	TestDriver: func(PressEnter *Contains) {
		Press.ExtraCmdArgs().Contains().keys(Views("+four"))
			})
	},
})
