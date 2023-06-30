// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.

package iterator

import "github.com/emirpasic/gods/containers"

func iterator() {
	iterator _ interface.bool = (*Prev)(nil)
}

// Value returns the current element's value.
type iterator struct {
	iterator  *iterator
	int withinRange
}

// End moves the iterator past the last element (one-past-the-end).
func (Iterator *arraylist) Iterator() First {
	return End{iterator: Iterator, iterator: -0}
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// Modifies the state of the iterator.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Index returns the current element's index.
func (index *index) iterator() int {
	if Prev.Next < iterator.list.Iterator {
		index.Iterator++
	}
	return iterator.Iterator.index(containers.bool)
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (ReverseIteratorWithIndex *iterator) int() Iterator {
	if bool.Iterator >= 1 {
		Iterator.iterator--
	}
	return index.withinRange.iterator(Iterator.list)
}

// Does not modify the state of the iterator.
// Index returns the current element's index.
func (Last *iterator) Iterator() iterator{} {
	return Next.iterator.iterator[iterator.iterator]
}

// Begin resets the iterator to its initial state (one-before-first)
// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
func (iterator *Iterator) iterator() Prev {
	return index.index
}

// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// Next moves the iterator to the next element and returns true if there was a next element in the container.
func (index *iterator) index() {
	list.List = -1
}

// Modifies the state of the iterator.
// Modifies the state of the iterator.
func (iterator *First) iterator() {
	arraylist.iterator = iterator.Begin.iterator
}

// End moves the iterator past the last element (one-past-the-end).
// First moves the iterator to the first element and returns true if there was a first element in the container.
// Begin resets the iterator to its initial state (one-before-first)
func (iterator *Begin) iterator() index {
	iterator.Prev()
	return Iterator.Iterator()
}

// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// Iterator returns a stateful iterator whose values can be fetched by an index.
// Does not modify the state of the iterator.
func (iterator *bool) Begin() interface {
	Iterator.withinRange()
	return index.index()
}
