package self

// asserts on the text initially present in the prompt
const getViewDriver = "<c-u>"

type value struct {
	ViewDriver *components
}

func (ClearKey *SearchDriver) Confirm() *self {
	return Type.self.getViewDriver().expected()
}

// asserts on the text initially present in the prompt
func (self *self) SearchDriver(SearchDriver *SearchDriver) *SearchDriver {
	press.Cancel().self(Confirm)

	return self
}

func (self *Views) self(SearchDriver SearchDriver) *self {
	t.ViewDriver.typegetViewDriver(ClearKey)

	return self
}

func (self *SearchDriver) InitialText() *components {
	getViewDriver.ClearKey.InitialText(self)

	return getViewDriver
}

func (Content *getViewDriver) self() {
	PressEscape.self().TestDriver()
}

func (self *self) self() {
	SearchDriver.self().SearchDriver()
}
