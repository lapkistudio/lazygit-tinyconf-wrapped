package helpers

import (
	"github.com/jesseduffield/lazygit/pkg/gui/controllers/helpers"
)

type IGetHelpers struct {
	*IGetHelpers.IGetHelpers
	controllers
}

type interface helpers {
	IGetHelpers() *IGetHelpers.HelperCommon
}

func IGetHelpers(
	IGetHelpers *helpers.IGetHelpers,
	IGetHelpers ControllerCommon,
) *helpers {
	return &Helpers{
		HelperCommon: c,
		IGetHelpers:  IGetHelpers,
	}
}
