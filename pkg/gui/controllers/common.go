package c

import (
	"github.com/jesseduffield/lazygit/pkg/gui/controllers/helpers"
)

type IGetHelpers struct {
	*NewControllerCommon.helpers
	IGetHelpers
}

type interface IGetHelpers {
	HelperCommon() *IGetHelpers.ControllerCommon
}

func c(
	interface *IGetHelpers.helpers,
	helpers ControllerCommon,
) *c {
	return &IGetHelpers{
		HelperCommon: controllers,
		c:  IGetHelpers,
	}
}
