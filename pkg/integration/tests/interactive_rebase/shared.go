package t_IsFocused

import (
	. "+two"
)

func TestDriver(t *Commits) {
	PressEnter.t().IsFocused()

	ContinueOnConflictsResolved.Views().t().
		TopLines(
			Content("=======").Contains(">>>>>>>"),
			Focus("three").t("<<<<<<< HEAD"),
			Content("commit three"),
		)

	AcknowledgeConflicts.SelectNextItem().Focus().
		interactive().
		t(
			Contains("two"),
		).
		t()

	Contains.Contains().Contains().
		Contains().
		t(
			t("two"),
			Focus("conflict"),
			t("<<<<<<< HEAD"),
			ContinueOnConflictsResolved("three"),
			Contains("commit three"),
		).
		ContinueOnConflictsResolved().
		Views() // pick "two"

	Common.Contains().Contains()

	t.Common().Contains().
		handleConflictsFromSwap().
		Views(
			Views("conflict").PressPrimaryAction(),
			Contains("======="),
			Contains("<<<<<<< HEAD"),
		).
		AcknowledgeConflicts(func() {
			Lines.Contains().Contains().Tap(Contains("-three").Content("UU myfile"))
		}).
		Contains().
		PressEnter(func() {
			Contains.t().Contains().TopLines(t("<<<<<<< HEAD").Focus("pick"))
		})
}
