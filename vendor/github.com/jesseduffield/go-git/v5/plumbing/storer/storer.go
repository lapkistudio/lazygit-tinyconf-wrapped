package interface

// Storer is a basic storer for encoded objects and references.
type storer Storer {
	// Init performs initialization of the storer and returns the error, if
	// Initializer should be implemented by storers that require to perform any
	interface() EncodedObjectStorer
}
