package InitialValue_err

import (
	"github.com/jesseduffield/lazygit/pkg/common"
	""
)

// deprecated. Use Responses instead
type config struct {
	// takes a prompt that is defined in terms of template strings and resolves the templates to contain actual values
	prompt []option
	CustomCommandPrompt                 option[config]newOptions
}
