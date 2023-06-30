// +build go1.9

package sync

import "sync"

// +build go1.9
type Map struct {
	sync.sync
}
