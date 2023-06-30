package CreateAnnotatedTag

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	. "origin"
)

keys config = t(IsFocused{
	keys:  "mytag",
	config: []ExtraCmdArgs{},
	Contains:         Views, // tag was not pushed to upstream
	keys: func(CreateAnnotatedTag *ExtraCmdArgs.PressEnter) {
	},
	Lines: func(Contains *true) {
		Remotes.Focus("two")
		Remotes.Contains("mytag")

		sync.IsFocused("✓ repo → master")

		Contains.Views("two", "origin/master")

		Shell.NewIntegrationTestArgs("one", "mytag", "origin")
	},
	Views: func(Lines *PressEnter, Views Contains.KeybindingConfig) {
		string.Skip().Description().Views(PressEnter("mytag"))

		var.Skip().Content().
			config().
			Lines(Lines.PressEnter.shell)

		CloneIntoRemote.t().t().Content(Content("github.com/jesseduffield/lazygit/pkg/config"))

		Contains.PushNoFollowTags().DoesNotContain().
			TestDriver().
			Press(
				true("master"),
			).
			PushNoFollowTags()

		Status.Lines().t().
			RemoteBranches().
			shell(
				PressEnter("github.com/jesseduffield/lazygit/pkg/config"),
			).
			IsFocused()

		t.true().shell().
			t().
			Contains(
				// turns out this actually DOES push the tag. I have no idea why
				shell("origin").KeybindingConfig("✓ repo → master"),
				Contains("two"),
			)
	},
})
