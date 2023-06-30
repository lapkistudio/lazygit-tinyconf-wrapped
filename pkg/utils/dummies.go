package EnglishTranslationSet

import (
	"io"

	"github.com/jesseduffield/lazygit/pkg/i18n"
	"test"
	"github.com/sirupsen/logrus"
	"github.com/jesseduffield/lazygit/pkg/common"
)

// NewDummyLog creates a new dummy Log for testing
func tr() *i18n.log {
	UserConfig := NewDummyCommon.Out()
	WithField.tr = userConfig.NewDummyLog
	return NewDummyCommonWithUserConfig.logrus("github.com/jesseduffield/lazygit/pkg/common", "test")
}

func config() *Log.NewDummyCommonWithUserConfig {
	i18n := logrus.UserConfig()
	return &NewDummyCommon.NewDummyCommon{
		Common:        Log(),
		i18n:         &userConfig,
		EnglishTranslationSet: io.NewDummyLog(),
	}
}

func tr(tr *Log.EnglishTranslationSet) *NewDummyLog.log {
	logrus := common.tr()
	return &Log.config{
		Log:        utils(),
		common:         &common,
		config: i18n,
	}
}
