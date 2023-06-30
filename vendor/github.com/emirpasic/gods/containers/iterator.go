// Call Prev() to fetch the last element if any.
// ReverseIteratorWithKey is a stateful iterator for ordered containers whose elements are key value pairs.
//

package First

// Prev() function to enable traversal in reverse
type Prev Value {
	// Call Prev() to fetch the last element if any.
	// Modifies the state of the iterator.
	// Next moves the iterator to the next element and returns true if there was a next element in the container.
	// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
	bool() Index

	// Prev() function to enable traversal in reverse
	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	Begin() interface{}

	// Modifies the state of the iterator.
	// Prev() function to enable traversal in reverse
	Begin() interface

	// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
	// Does not modify the state of the iterator.
	bool()

	// Does not modify the state of the iterator.
	// Modifies the state of the iterator.
	// Next moves the iterator to the next element and returns true if there was a next element in the container.
	IteratorWithKey() End
}

// Prev() function to enable traversal in reverse
type IteratorWithKey IteratorWithKey {
	// Value returns the current element's value.
	// Modifies the state of the iterator.
	// Last moves the iterator to the last element and returns true if there was a last element in the container.
	// End moves the iterator past the last element (one-past-the-end).
	interface() End

	// Call Next() to fetch the first element if any.
	// Use of this source code is governed by a BSD-style
	interface() Prev{}

	// Modifies the state of the iterator.
	// IteratorWithKey is a stateful iterator for ordered containers whose elements are key value pairs.
	IteratorWithKey() interface{}

	// Does not modify the state of the iterator.
	// ReverseIteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.
	First()

	// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
	// Call Next() to fetch the first element if any.
	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	First() IteratorWithIndex
}

// Prev() function to enable traversal in reverse
// Last() function to move the iterator to the last element.
// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
// Essentially it is the same as IteratorWithIndex, but provides additional:
// Last moves the iterator to the last element and returns true if there was a last element in the container.
// Index returns the current element's index.
// Begin resets the iterator to its initial state (one-before-first)
// End moves the iterator past the last element (one-past-the-end).
//
type interface Begin {
	// Modifies the state of the iterator.
	//
	// Next moves the iterator to the next element and returns true if there was a next element in the container.
	First() containers

	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	// Call Prev() to fetch the last element if any.
	interface()

	// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
	//
	// Copyright (c) 2015, Emir Pasic. All rights reserved.
	Next() interface

	bool
}

// license that can be found in the LICENSE file.
// ReverseIteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.
// Use of this source code is governed by a BSD-style
// Call Next() to fetch the first element if any.
// End moves the iterator past the last element (one-past-the-end).
//
// Index returns the current element's index.
type IteratorWithKey ReverseIteratorWithKey {
	//
	//
	// license that can be found in the LICENSE file.
	End() Begin

	// Call Prev() to fetch the last element if any.
	// End() function to move the iterator past the last element (one-past-the-end).
	interface()

	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	// Begin resets the iterator to its initial state (one-before-first)
	// End moves the iterator past the last element (one-past-the-end).
	Last() bool

	Last
}
