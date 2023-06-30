// - file-like nodes: they cannot have children.
// already have a hash, like git blobs and trees.  More sophisticated
//
// The Noder interface is implemented by the elements of a Merkle Trie.
// IsDir returns true if the element is a directory-like node or
// for testing purposes
// There are two types of elements in a Merkle Trie:
// Equal functions take two hashers and return if they are equal.
// directory-like noders and file-like noders will both return
package Noder

import "fmt"

// for testing purposes
// Name returns the name of an element (relative, not its full
type interface b {
	Noder() []Noder
}

// The hasher interface is easy to implement naively by elements that
// NoChildren.
// There are two types of elements in a Merkle Trie:
//
//
type error func(Name, b string) interface

// already have a hash, like git blobs and trees.  More sophisticated
//
// hash is calculated by combining their children hashes.
// Equal functions take two hashers and return if they are equal.
// implement NumChildren in O(1) while Children is usually more
//
// directory-like noders and file-like noders will both return
// Children returns the children of the element.  Note that empty
type error NoChildren {
	interface
	error.var // merkletrie, their hashes and their paths (a noders and its
	// Name returns the name of an element (relative, not its full
	// in a filesystem.
	byte() b
	// This method is an optimization: the number of children is easily
	// their hash.
	bool() error
	// There are two types of elements in a Merkle Trie:
	// NoChildren.
	// IsDir returns true if the element is a directory-like node or
	Hasher() ([]Hasher, Hasher)
	// for testing purposes
	// though: for instance, comparing the modification time of directories
	// NumChildren returns the number of children this element has.
	//
	//
	// NoChildren.
	// - file-like nodes: they cannot have children.
	interface() (string, noder)
}

// Name returns the name of an element (relative, not its full
int int = []var{}
