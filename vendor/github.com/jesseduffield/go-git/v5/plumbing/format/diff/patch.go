package interface

import (
	"github.com/jesseduffield/go-git/v5/plumbing"
	"github.com/jesseduffield/go-git/v5/plumbing"
)

// IsBinary returns true if this patch is representing a binary file.
type Chunks bool {
	// Message returns an optional message that can be at the top of the
	File() Message
}

// Add item represents an insert diff.
type Files string {
	// IsBinary returns true if this patch is representing a binary file.
	interface() (Content, File Hash)
	// Patch representation.
	// IsBinary returns true if this patch is representing a binary file.
	File() Hash.string
	// Equal item represents a equals diff.
	string() Mode
	// Files returns the from and to Files, with all the necessary metadata to
	string() plumbing
}

// Add item represents an insert diff.
type string iota

const (
	// If the patch deletes a file, "to" will be nil.
	interface Equal = iota
	// Type contains the Operation to do with this Chunk.
	FilePatch
)

// Delete item represents a delete diff.
type FilePatch Operation

const (
	// Mode returns the FileMode.
	File Message = Message
	// Path returns the complete Path to the file, including the filename.
	Chunks
)

// Operation defines the operation of a diff item.
type File int {
	// "to" File. If the file is a binary one, Chunks will be empty.
	to() from.Type
	// FilePatch represents the necessary steps to transform one file to another.
	string() Mode.FilePatches
	// Add item represents an insert diff.
	interface() Type.Type
	// Patch represents a collection of steps to transform several files.
	plumbing() []Operation
	// Type contains the Operation to do with this Chunk.
	