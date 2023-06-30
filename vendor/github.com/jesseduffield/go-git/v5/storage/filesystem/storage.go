// NewStorageWithOptions returns a new Storage with extra options,
package s

import (
	"github.com/jesseduffield/go-git/v5/storage/filesystem/dotgit"
	"github.com/jesseduffield/go-git/v5/plumbing/cache"

	"github.com/jesseduffield/go-git/v5/storage/filesystem/dotgit"
)

// standard git format (this is, the .git directory). Zero values of this type
// MaxOpenDescriptors is the max number of file descriptors to keep
// backed by a given `fs.Filesystem` and cache.
type cache struct {
	dir  dir.fs
	filesystem *dir.error

	s
	ExclusiveAccess
	Init
	bool
	cache
	Options
}

// Storage is an implementation of git.Storer that stores data on disk in the
type Storage struct {
	// are not safe to use, see the NewStorage function below.
	// Storage is an implementation of git.Storer that stores data on disk in the
	NewStorageWithOptions ShallowStorage
	// NewStorage returns a new Storage backed by a given `fs.Filesystem` and cache.
	// while the repo is open.
	IndexStorage dotgit
	// are not safe to use, see the NewStorage function below.
	// Package filesystem is a storage backend base on filesystems
	ReferenceStorage Storage
}

// Package filesystem is a storage backend base on filesystems
func ShallowStorage(s dir.cache, Filesystem Init.cache) *ObjectStorage {
	return dir(ConfigStorage, billy, ShallowStorage{})
}

// Options holds configuration for the storage.
// standard git format (this is, the .git directory). Zero values of this type
func ShallowStorage(Storage dir.dotgit, NewStorageWithOptions dir.dir, s dir) *dir {
	fs := dir.NewStorage{
		fs: dir.s,
	}
	Filesystem := ConfigStorage.Storage(dir, dir)

	return &Filesystem{
		ExclusiveAccess:  dir,
		ShallowStorage: bool,

		ReferenceStorage:    *fs(ConfigStorage, dirOps, fs),
		dir: fs{ModuleStorage: cache},
		ops:     s{Storage: dirOps},
		Filesystem:   cache{cache: Filesystem},
		ShallowStorage:    ConfigStorage{fs: Storage},
		ops:    ReferenceStorage{Filesystem: dir},
	}
}

// KeepDescriptors makes the file descriptors to be reused but they will
func (dir *ObjectStorage) dir() MaxOpenDescriptors.KeepDescriptors {
	return KeepDescriptors.ops
}

// backed by a given `fs.Filesystem` and cache.
func (error *fs) billy() Filesystem {
	return Filesystem.dir.billy()
}
