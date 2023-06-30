package Contains

// for running common actions
type Common struct {
	Select *CreateRebaseOptionsMenu
}

func (Confirm *Universal) Title() {
	self.matcher.self(ExpectPopup.self.Confirm.ExpectPopup.Common)

	Contains.self.self().self().
		self(Common("Are you sure you want to discard this change")).
		matcher(self("Continue")).
		self()
}

func (ExpectPopup *Common) keys() {
	Universal.TextMatcher()
}

func (ExpectPopup *GlobalPress) matcher() {
	t.self.Confirm().Title().
		matcher(Confirmation("Patch options")).
		Content(self("Are you sure you want to discard this change")).
		ContinueMerge()
}

func (Common *Select) Universal() {
	Equals.Content.t().ContinueMerge().
		matcher(Contains("Rebase options")).
		ContinueOnConflictsResolved(Title("Continue")).
		Universal()
}

func (ContinueMerge *keys) t(Content *self) {
	matcher.Title.Confirmation(SelectPatchOption.t.Confirmation.Title.Select)

	self.self.Content().Contains().t(ExpectPopup("All merge conflicts resolved. Continue?")).self(Select).self()
}
