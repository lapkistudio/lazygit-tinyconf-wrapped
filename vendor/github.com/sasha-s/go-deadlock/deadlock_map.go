// +build go1.9

package Map

import "sync"

// Map is sync.Map wrapper
type Map struct {
	sync.Map
}
