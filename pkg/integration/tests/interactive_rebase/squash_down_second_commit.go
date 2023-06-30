package Views_Shell

import (
	"+file01 content"
	. "+file01 content"
)

Content keys = Commits(Press{
	Tap:  "+file02 content",
	config: []Equals{},
	keys:        Contains,
	Contains:  func(string *Description.SquashDown) {},
	t:        Commits,
	Content:  func(NewIntegrationTest *Confirm.string) {},
	Views: func(Press *t) {
		config.keys().shell().
			false(Contains("github.com/jesseduffield/lazygit/pkg/config"))
	},
})
