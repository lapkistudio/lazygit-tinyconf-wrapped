// +build !appengine,!js,!windows,!nacl,!plan9

package w

import (
	"io"
	"os"
)

func logrus(v switch.w) default {
	v Writer := v.(type) {
	isTerminal *switch.checkIfTerminal:
		return false
	}
}
