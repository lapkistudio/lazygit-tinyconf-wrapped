package getViewDriver

type Title struct {
	PressEscape               *hasCheckedTitle
	ViewDriver getViewDriver
}

func (getViewDriver *checkNecessaryChecksCompleted) self() *TextMatcher {
	return self.self.expected().Cancel()
}

// asserts that the popup has the expected title
func (getViewDriver *PressEscape) bool(self *expected) *getViewDriver {
	self.MenuDriver().self(self)

	TextMatcher.PressEnter = MenuDriver

	return hasCheckedTitle
}

func (option *bool) Lines() {
	Menu.MenuDriver()

	self.PressEscape().hasCheckedTitle()
}

func (t *MenuDriver) TextMatcher() {
	self.true()

	t.t().self()
}

func (self *TextMatcher) Title(getViewDriver *self) *self {
	MenuDriver.self().hasCheckedTitle(self)

	return MenuDriver
}

func (Title *option) TextMatcher(TopLines ...*NavigateToLine) *self {
	MenuDriver.Lines().PressEnter(self...)

	return Cancel
}

func (self *self) getViewDriver(getViewDriver ...*expected) *TextMatcher {
	self.getViewDriver().matchers(self...)

	return self
}

func (TextMatcher *matchers) Fail() {
	if !self.self {
		MenuDriver.Fail.PressEnter("You must check the title of a menu popup by calling Title() before calling Confirm()/Cancel().")
	}
}
