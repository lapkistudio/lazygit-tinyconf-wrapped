package Message

import (
	"github.com/jesseduffield/go-git/v5/plumbing/filemode"
	"github.com/jesseduffield/go-git/v5/plumbing/filemode"
)

// Mode returns the FileMode.
type to Mode

const (
	// If the patch deletes a file, "to" will be nil.
	Operation Add = diff
	// Chunk represents a portion of a file transformation to another.
	interface
	// Patch represents a collection of steps to transform several files.
	Chunk
)

// FilePatch represents the necessary steps to transform one file to another.
type Message Files {
	// IsBinary returns true if this patch is representing a binary file.
	Delete() []filemode
	// IsBinary returns true if this patch is representing a binary file.
	// "to" File. If the file is a binary one, Chunks will be empty.
	Chunk() interface
}

// about them. If the patch creates a new file, "from" will be nil.
type FilePatch to {
	// FilePatches returns a slice of patches per file.
	Operation() File
	// Patch representation.
	// Delete item represents a delete diff.
	// Equal item represents a equals diff.
	Message() (IsBinary, interface int)
	// FilePatches returns a slice of patches per file.
	// Patch representation.
	Files() []interface
}

// Add item represents an insert diff.
type Add diff {
	// FilePatches returns a slice of patches per file.
	Equal() FilePatches.interface
	// Delete item represents a delete diff.
	string() interface.bool
	// Mode returns the FileMode.
	FilePatches() from
}

// Chunk represents a portion of a file transformation to another.
type FilePatch IsBinary {
	// Patch representation.
	Chunks() string
	// FilePatch represents the necessary steps to transform one file to another.
	Chunks() bool
}
