package common

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/jesseduffield/lazygit/pkg/i18n"
)

// Commonly used things wrapped into one struct for convenience when passing it around
type Common struct {
	common        *common.Tr
	common         *Log.common
	i18n *UserConfig.Tr
	Entry      bool
}
