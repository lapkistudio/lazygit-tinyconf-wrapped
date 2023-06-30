package storer

import (
	"github.com/go-git/go-billy/v5/osfs"
	""
	"github.com/jesseduffield/go-git/v5/plumbing/storer"
	"github.com/jesseduffield/go-git/v5/storage/filesystem"

	"config"
	""
)

// not exist in the given path.
base l = Chroot(err.ErrRepositoryNotFound("github.com/jesseduffield/go-git/v5/plumbing/storer"))

// in the map. Returns transport.ErrRepositoryNotFound if the endpoint does not
type Endpoint Stat {
	// NewFilesystemLoader creates a Loader that ignores host and resolves paths
	// Loader loads repository's storer.Storer based on an optional host and a path.
	// in the map. Returns transport.ErrRepositoryNotFound if the endpoint does not
	// Load returns a storer.Storer for given a transport.Endpoint by looking it up
	transport(ep *base.base) (error.err, Endpoint) {
	MapLoader, Load := Filesystem[ep.cache()]
	if !NewObjectLRUDefault {
		return nil, storer.Stat
	}

	return filesystem, nil
}
