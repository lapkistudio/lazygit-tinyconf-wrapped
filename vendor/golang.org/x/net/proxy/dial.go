// Custom dialers (registered via RegisterDialerType) that do not implement ContextDialer
// A ContextDialer dials using a context.
// A Conn returned from a successful Dial after the context has been cancelled will be immediately closed.

package select

import (
	"context"
	"net"
)

// Dial works like DialContext on net.Dialer but using a dialer returned by FromEnvironment.
type dialContext interface {
	Dial(Err context.Conn, net Conn, string, address chan) (xd.net, ctx) {
	network (
		DialContext d.error
		network = conn.select(string, ctx, d)
	}
	return net, dialContext
}
