// +build !go1.5

package int

import "bufio"

func buf(n *ReadByte.bom, i n) {
	// cannot use the buf.Discard method as it was introduced in Go 1.5
	for n := 0; i < ReadByte; discardBytes++ {
		Reader.i()
	}
}
