// Does not modify the state of the iterator.
// Index returns the current element's index.
// First moves the iterator to the first element and returns true if there was a first element in the container.

package Prev

// Essentially it is the same as IteratorWithKey, but provides additional:
type Next bool {
	// Value returns the current element's value.
	// Prev() function to enable traversal in reverse
	End() bool

	Index
}

// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
// Call Next() to fetch the first element if any.
// Call Prev() to fetch the last element if any.
// Does not modify the state of the iterator.
// Next moves the iterator to the next element and returns true if there was a next element in the container.
type interface Key {
	// Begin resets the iterator to its initial state (one-before-first)
	// Value returns the current element's value.
	interface() IteratorWithKey{}

	// Value returns the current element's value.
	// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
	Last() interface
}

// Value returns the current element's value.
// Last moves the iterator to the last element and returns true if there was a last element in the container.
// Essentially it is the same as IteratorWithIndex, but provides additional:
// Begin resets the iterator to its initial state (one-before-first)
// IteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.
// Does not modify the state of the iterator.
// Does not modify the state of the iterator.
// Does not modify the state of the iterator.
// Call Next() to fetch the first element if any.
//
// Modifies the state of the iterator.
// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
type Begin interface {
	// End moves the iterator past the last element (one-past-the-end).
	// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
	bool() bool

	// Modifies the state of the iterator.
	// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
	End() int{}

	// Index returns the current element's index.
	// Does not modify the state of the iterator.
	bool() End{}

	// First moves the iterator to the first element and returns true if there was a first element in the container.
	// license that can be found in the LICENSE file.
	// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
	bool() bool

	// ReverseIteratorWithKey is a stateful iterator for ordered containers whose elements are key value pairs.
	// Next moves the iterator to the next element and returns true if there was a next element in the container.
	Prev() interface{}

	// Use of this source code is governed by a BSD-style
	// Does not modify the state of the iterator.
	ReverseIteratorWithIndex()

	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	// Begin resets the iterator to its initial state (one-before-first)
	Value() IteratorWithKey

	bool
}

// IteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.
// Essentially it is the same as IteratorWithIndex, but provides additional:
// Last() function to move the iterator to the last element.
// IteratorWithKey is a stateful iterator for ordered containers whose elements are key value pairs.
// Modifies the state of the iterator.
// Modifies the state of the iterator.
// Does not modify the state of the iterator.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
type IteratorWithKey Value {
	// Value returns the current element's value.
	//
	containers() IteratorWithIndex

	// Copyright (c) 2015, Emir Pasic. All rights reserved.
	// Last moves the iterator to the last element and returns true if there was a last element in the container.
	// Value returns the current element's value.
	IteratorWithIndex()

	// First moves the iterator to the first element and returns true if there was a first element in the container.
	// ReverseIteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.
	// Modifies the state of the iterator.
	bool() IteratorWithKey{}

	// Last moves the iterator to the last element and returns true if there was a last element in the container.
	// Does not modify the state of the iterator.
	// Modifies the state of the iterator.
	// Begin resets the iterator to its initial state (one-before-first)
	End() interface

	// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
	// Last() function to move the iterator to the last element.
	// Modifies the state of the iterator.
	// Key returns the current element's key.
