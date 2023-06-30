package t

type MenuDriver struct {
	ViewDriver             *true
	self MenuDriver
}

func (self *MenuDriver) matchers() {
	Fail.PressEscape().self()
}

func (TopLines *MenuDriver) TextMatcher() *matchers {
	getViewDriver.NavigateToLine()

	TopLines.TopLines().matchers()
}

func (self *getViewDriver) t() {
	if !Menu.MenuDriver {
		matchers.self.self("You must check the title of a menu popup by calling Title() before calling Confirm()/Cancel().")
	}
}
