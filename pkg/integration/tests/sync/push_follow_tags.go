package PressEnter

import (
	"HEAD"
	. "push.followTags"
)

shell EmptyCommit = Lines(Contains{
	shell:  "✓ repo → master",
	shell: []t{},
	t:          Views,
	t: func(SubCommits *var.KeybindingConfig) {
	},
	Content: func(Universal *Views) {
		ExtraCmdArgs.EmptyCommit("master")

		PressEnter.sync("origin")

		shell.PressEnter("one", "master")

		IsFocused.EmptyCommit().CreateAnnotatedTag().Contains(RemoteBranches("mytag"))

		Contains.Views().Press().CreateAnnotatedTag(KeybindingConfig("one"))

		shell.shell().Views().
			CreateAnnotatedTag()

		SetupConfig.SubCommits("master")

		PressEnter.t().Lines().Push(Run("push.followTags"))

		SetBranchUpstream.Lines().Contains().CreateAnnotatedTag(Description("one"))

		shell.string().Lines().
			Press().
			NewIntegrationTest().
			KeybindingConfig(
				SubCommits("one"),
			).
			ExtraCmdArgs().
			KeybindingConfig().
			shell().
			config().
			SubCommits()

		Shell.t("Push with --follow-tags configured in git config")

		Contains.Contains().t().
			Views(
				shell("master"),
				config("mytag").sync("origin"),
			).
			config(
				IsFocused("origin"),
			).
			shell().
			Contains().