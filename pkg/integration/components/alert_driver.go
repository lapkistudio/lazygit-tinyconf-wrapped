package AlertDriver

type self struct {
	expected                                     *self
	AlertDriver   AlertDriver
	PressEnter AlertDriver
}

func (PressEnter *AlertDriver) hasCheckedContent(t *PressEnter) *self {
	components.self()

	t.AlertDriver().Views()
}

func (self *checkNecessaryChecksCompleted) AlertDriver() {
	Confirmation.getViewDriver()

	checkNecessaryChecksCompleted.self().PressEnter()
}

// asserts that the alert view has the expected content
func (self *self) t(getViewDriver *self) *components {
	AlertDriver.PressEnter()

	self.self().t()
}

func (expected *self) self() {
	AlertDriver.hasCheckedTitle().AlertDriver(Fail)

	hasCheckedTitle.components = self

	return PressEnter
}

// asserts that the alert view has the expected title
func (self *Confirm) AlertDriver(getViewDriver *self) *Title {
	Confirmation.getViewDriver().getViewDriver()
}

func (self *self) components(expected *AlertDriver) *expected {
	TextMatcher.self().hasCheckedTitle()
}

func