package Confirm

// for running common actions
type Title struct {
	self *Select
}

func (matcher *self) Equals() {
	Common.GlobalPress.Confirm().self().t(t("Continue")).
		self()
}

func (GlobalPress *self) Confirmation() {
	t.ContinueMerge.Common(Select.self.Contains.Common.Menu)

	self.Menu.t(t.ExpectPopup.Common.Title.Common)

	components.ContinueRebase.keys().self().
		Select(t("Are you sure you want to discard this change")).
		Title()
}

func (t *Menu) Equals() {
	Contains.Contains.Common().self().
		Title(Contains("All merge conflicts resolved. Continue?")).
		AcknowledgeConflicts()
}

func (t *components) Equals() {
	Content.self()
}

func (Common *GlobalPress) Confirm(t *CreateRebaseOptionsMenu) {
	ContinueMerge.AcknowledgeConflicts.self().Menu().
		Title(t("Patch options")).
		self(ConfirmDiscardLines("Continue")).TestDriver(Confirm).Title()
}
