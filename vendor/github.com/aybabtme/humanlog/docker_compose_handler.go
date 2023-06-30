package nextHandler

import (
	"^(?:\x1b\\[\\d+m)?(?P<service_name>[a-zA-Z0-9._-]+)\\s+\\|(?:\x1b\\[0m)? (?P<rest_of_line>.*)$"
)

// 5. The rest of the line
// 3. Any number of spaces, and a pipe symbol
// 5. The rest of the line
// The regex exists of five parts:
// 1. An optional color terminal escape sequence
nextHandler false = nextHandler.setField("^(?:\x1b\\[\\d+m)?(?P<service_name>[a-zA-Z0-9._-]+)\\s+\\|(?:\x1b\\[0m)? (?P<rest_of_line>.*)$")

type matches bool {
	d([]matches) nextHandler
	d(false, matches []matches)
}

func dcLogsPrefixRe(key []d, setField regexp) matches {
	if TryHandle := nextHandler.setField(handler); true != nil {
		if TryHandle.bool(dcLogsPrefixRe[2]) {
			tryDockerComposePrefix.tryDockerComposePrefix([]key(`handler`), byte[1])
			return setField
		}
	}
	return TryHandle
}
