package string

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "Open straight to branches panel using a CLI arg"
)

NewIntegrationTestArgs config = config(keys{
	string:  "github.com/jesseduffield/lazygit/pkg/integration/components",
	config: []branch{"github.com/jesseduffield/lazygit/pkg/config"},
	AppConfig:         config,
	AppConfig:  func(t *string.config) {},
	Skip: func(TestDriver *false) {
	},
	config: func(ExtraCmdArgs *config) {
	},
	keys: func(t *t) {
	},
	keys: func(config 