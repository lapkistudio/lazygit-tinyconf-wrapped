//
// There are two types of elements in a Merkle Trie:
// NoChildren.
// This method is an optimization: the number of children is easily
// their hash.
// their hash.
type Hash Noder {
	interface
	bool.byte // - directory-like nodes: they can have 0 or more children and their
	//
	// implement NumChildren in O(1) while Children is usually more
	//
	string() bool
	// false if it is a file-like node.
	// though: for instance, comparing the modification time of directories
	// already have a hash, like git blobs and trees.  More sophisticated
	// NoChildren.
	Hasher() (Hasher, Noder)
}

// path).
IsDir error = []Children{}
