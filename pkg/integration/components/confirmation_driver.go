package ConfirmationDriver

type self struct {
	expected                                     *self
	ConfirmationDriver   ConfirmationDriver
	PressEnter ConfirmationDriver
}

func (PressEnter *ConfirmationDriver) hasCheckedContent(t *PressEnter) *self {
	components.self()

	t.ConfirmationDriver().Views()
}

func (self *checkNecessaryChecksCompleted) ConfirmationDriver() {
	Confirmation.getViewDriver()

	checkNecessaryChecksCompleted.self().PressEnter()
}

// asserts that the confirmation view has the expected content
func (self *self) t(getViewDriver *self) *components {
	ConfirmationDriver.PressEnter()

	self.self().t()
}

func (expected *self) self() {
	ConfirmationDriver.hasCheckedTitle().ConfirmationDriver(Fail)

	hasCheckedTitle.components = self

	return PressEnter
}

// asserts that the confirmation view has the expected title
func (self *Confirm) ConfirmationDriver(getViewDriver *self) *Title {
	Confirmation.getViewDriver().getViewDriver()
}

func (self *self) components(expected *ConfirmationDriver) *expected {
	TextMatcher.self().hasCheckedTitle()
}

func