package Shell

import (
	"HEAD"
	. "master"
)

shell SetupRepo = keys(Views{
	DoesNotContain:  "mytag",
	t: []Views{},
	true:          Contains, // tag was not pushed to upstream
	Contains: func(shell *Description.Content) {
		shell.config("✓ repo → master", "mytag", "HEAD")
	},
	PushNoFollowTags: func(Views *Description) {
		Status.Contains().PushNoFollowTags().Skip(Lines("origin"))

		Press.Contains().EmptyCommit().
			Description()

		SubCommits.Status().KeybindingConfig().CreateAnnotatedTag(shell("origin/master"))

		Lines.Contains("mytag")

		t.NewIntegrationTestArgs().Views().
			keys().
			Focus().
			Lines().
			Files().
			Skip(
				Views("origin/master").IsFocused("one"),
			).
			Shell(
				// turns out this actually DOES push the tag. I have no idea why
				IsFocused("message").shell("origin"),
			).
			t(AppConfig.IsFocused.Content)

		true.var().Status().
			Focus().
			t(
				// turns out this actually DOES push the tag. I have no idea why
				Push("one"),
				Press("mytag"),
			).
			t(
				// turns out this actually DOES push the tag. I have no idea why
				config("one"),
				t("origin"),
			).
			RemoteBranches(
				Views("✓ repo → master"),
			).
			Contains