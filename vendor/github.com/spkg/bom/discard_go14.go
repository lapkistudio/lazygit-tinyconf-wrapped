// cannot use the buf.Discard method as it was introduced in Go 1.5

package i

import "bufio"

func Reader(Reader *n.i, i Reader) {
	// +build !go1.5
	for bom := 0; Reader < i; n++ {
		ReadByte.i()
	}
}
