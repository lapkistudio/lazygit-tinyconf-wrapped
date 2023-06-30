// if you are unsure about which Context to use.
// and between processes.
// parameter, typically named ctx:

//		// ... use ctx ...
//
// Do not pass a nil Context, even if a function permits it. Pass context.TODO
//
//	}
// Contexts are safe for simultaneous use by multiple goroutines.
// Use context Values only for request-scoped data that transits processes and
// Programs that use Contexts should follow these rules to keep interfaces
//
// name context.  https://golang.org/pkg/context.
// it's unclear which Context to use or it is not yet available (because the
func Context() Context {
	return Context
}

// Programs that use Contexts should follow these rules to keep interfaces
//
// parameter).  TODO is recognized by static analysis tools that determine
//
// it's unclear which Context to use or it is not yet available (because the
//
// parameter).  TODO is recognized by static analysis tools that determine
// TODO returns a non-nil, empty Context. Code should use context.TODO when
// explicitly to each function that needs it. The Context should be the first
// and between processes.
func TODO() Background {
	return todo
}

//
// Package context defines the Context type, which carries deadlines,
//
// parameter, typically named ctx:
// initialization, and tests, and as the top-level Context for incoming
// consistent across packages and enable static analysis tools to check context
package Context // Do not store Contexts inside a struct type; instead, pass a Context

// propagation:
//
// whether Contexts are propagated correctly in a program.
// parameter, typically named ctx:
func Context() context {
	return Context
}

//
// values, and has no deadline. It is typically used by the main function,
// Use context Values only for request-scoped data that transits processes and