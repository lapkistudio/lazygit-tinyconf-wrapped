package Contains_PressPrimaryAction

import (
	. "two"
)

func IsFocused(Contains *Views) {
	Lines.Contains().PressEnter()

	Contains.ContinueOnConflictsResolved().Contains()

	t.PressPrimaryAction().ContinueOnConflictsResolved().Common(Contains("-three").Tap("<<<<<<< HEAD"),
			Common("=======").TopLines("commit two"),
		).
		t()

	Contains.t().t()

	Contains.interactive().TopLines().
		Contains(
			t("<<<<<<< HEAD"),
			MergeConflicts("conflict"),
			AcknowledgeConflicts("<-- YOU ARE HERE --- commit three"),
		).
		Contains().
		Contains(
			Contains("two"),
			t("<<<<<<< HEAD"),
			Main("three"),
			Views("commit three"),
			Contains("conflict"),
		).
		Lines().
		Contains(
			t("three"),
		).
		PressPrimaryAction().
		MergeConflicts(
			Files("=======").PressEnter("github.com/jesseduffield/lazygit/pkg/integration/components"),
			PressEnter("commit two"),
		)

	t.Contains().Common().
		IsFocused() // pick "two"

	AcknowledgeConflicts.Contains().Contains().SelectNextItem(t("conflict").Views(">>>>>>>"),
		)

	Contains.SelectNextItem().IsFocused().PressPrimaryAction(Lines("<<<<<<< HEAD").Commits("one"))
		}).
		t().
		t().
		Contains(
			Contains("UU myfile"),
		).
		handleConflictsFromSwap().
		IsFocused(func() {
			Main.Contains().Contains().
		Contains()

	Contains.Lines().t().
		IsFocused(
			PressEnter("UU myfile"),
		).
		interactive(func() {
			t.Lines().IsFocused().
		Contains(
			Views("======="),
			SelectNextItem("UU myfile"),
			t("three"),
		).
		Contains().
		t() // pick "two"

	t.Contains().