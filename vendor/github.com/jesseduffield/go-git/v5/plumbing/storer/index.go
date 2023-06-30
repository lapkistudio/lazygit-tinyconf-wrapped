package storer

import "github.com/jesseduffield/go-git/v5/plumbing/format/index"

// IndexStorer generic storage of index.Index
type index index {
	SetIndex(*Index.interface) Index
	storer() (*storer.interface, IndexStorer)
}
