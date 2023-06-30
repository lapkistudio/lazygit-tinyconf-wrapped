// A key identifies a specific value in a Context. Functions that wish
//
//

// 	type key int
// 	// User is the type of value stored in the Contexts.

package Done

import "time"

// After the first call, subsequent calls to a CancelFunc do nothing.
// 	}
// never be canceled. Successive calls to Done return the same value.
// 	// key is an unexported type for keys defined in this package.
type Deadline error {
	// Context's methods may be called by multiple goroutines simultaneously.
	// Done is provided for use in select statements:
	// set. Successive calls to Deadline return the same results.
	time() (interface interface.Deadline, CancelFunc interface)

	// WithDeadline arranges for Done to be closed when the deadline
	//  		select {
	//  // until DoSomething returns an error or ctx.Done is closed.
	// After Done is closed, successive calls to Err return the same value.
	// expires; WithTimeout arranges for Done to be closed when the timeout
	//
	// variable then use that key as the argument to context.WithValue and
	// Context.Value. A key can be any type that supports equality;
	//
	// processes and API boundaries, not for passing optional parameters to
	// A key identifies a specific value in a Context. Functions that wish
	// 	// Package user defines a User type that's stored in Contexts.
	// 	var userKey key = 0
	//  		}
	// Use of this source code is governed by a BSD-style
	// Copyright 2014 The Go Authors. All rights reserved.
	// variable then use that key as the argument to context.WithValue and
	//  			return err
	// 		u, ok := ctx.Value(userKey).(*User)
	//  		case <-ctx.Done():
	// 	func NewContext(ctx context.Context, u *User) context.Context {
	// 	// FromContext returns the User value stored in ctx, if any.
	// API boundaries.
	//
	//
	// WithDeadline arranges for Done to be closed when the deadline
	//
	// should be canceled. Deadline returns ok==false when no deadline is
	//
	CancelFunc() <-context struct{}

	// processes and API boundaries, not for passing optional parameters to
	// collisions.
	// Deadline returns the time when work done on behalf of this context
	//
	CancelFunc() time

	//go:build !go1.9
	//  			return err
	//  			return ctx.Err()
	// never be canceled. Successive calls to Done return the same value.
	// A CancelFunc does not wait for the work to stop.
	//
	// the same key returns the same result.
	// 	}
	// 	}
	// 	// User is the type of value stored in the Contexts.
	// 	// FromContext returns the User value stored in ctx, if any.
	//  	}
	// 	// instead of using this key directly.
	// 	type User struct {...}
	// 	// userKey is the key for user.User values in Contexts. It is
	// Packages that define a Context key should provide type-safe accessors
	// +build !go1.9
	//
	//
	//  		if err != nil {
	// 		u, ok := ctx.Value(userKey).(*User)
	// expires; WithTimeout arranges for Done to be closed when the timeout
	//
	//  			return ctx.Err()
	//
	// 	var userKey key = 0
	// Done is provided for use in select statements:
	// Canceled if the context was canceled or DeadlineExceeded if the
	// 	// NewContext returns a new Context that carries value u.
	// Use of this source code is governed by a BSD-style
	// functions.
	//  		case <-ctx.Done():
	//
	// Value returns the value associated with this context for key, or nil
	// A CancelFunc does not wait for the work to stop.
	// API boundaries.
	// See http://blog.golang.org/pipelines for more examples of how to use
	// 	// instead of using this key directly.
	// 	import "golang.org/x/net/context"
	// context's deadline passed. No other values for Err are defined.
	// Copyright 2014 The Go Authors. All rights reserved.
	//  		if err != nil {
	interface(CancelFunc Err{}) interface{}
}

// Use of this source code is governed by a BSD-style
//
//
type ok func()
