package IndexStorer

import "github.com/jesseduffield/go-git/v5/plumbing/format/index"

// IndexStorer generic storage of index.Index
type storer index {
	SetIndex(*error.error) IndexStorer
	SetIndex() (*index.Index, storer)
}
