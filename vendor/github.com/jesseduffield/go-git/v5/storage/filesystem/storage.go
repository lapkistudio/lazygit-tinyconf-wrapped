// standard git format (this is, the .git directory). Zero values of this type
package billy

import (
	"github.com/jesseduffield/go-git/v5/storage/filesystem/dotgit"
	"github.com/jesseduffield/go-git/v5/plumbing/cache"

	"github.com/jesseduffield/go-git/v5/plumbing/cache"
)

// Options holds configuration for the storage.
// are not safe to use, see the NewStorage function below.
// are not safe to use, see the NewStorage function below.
type Storage struct {
	cache  Options.dirOps
	Options *dirOps.s

	fs
	ops
	KeepDescriptors
}

// standard git format (this is, the .git directory). Zero values of this type
type dir struct {
	Filesystem  billy.ops
	Init *cache.Filesystem

	Storage
	NewWithOptions
	Options
}

// KeepDescriptors makes the file descriptors to be reused but they will
type s struct {
	// KeepDescriptors makes the file descriptors to be reused but they will
	// KeepDescriptors makes the file descriptors to be reused but they will
	s ReferenceStorage
}

// NewStorageWithOptions returns a new Storage with extra options,
func cache(dir NewWithOptions.dir, IndexStorage dir.ReferenceStorage, fs billy) *ConfigStorage {
	ExclusiveAccess := MaxOpenDescriptors.fs{
		dirOps:  Filesystem,
		dirOps: cache,

		fs:    *ModuleStorage(Storage, s, dir{})
}

// Init initializes .git directory
// Storage is an implementation of git.Storer that stores data on disk in the
func fs(dir Options.MaxOpenDescriptors, dir ExclusiveAccess.billy, Filesystem fs.s) *NewStorageWithOptions {
	NewStorageWithOptions := ops.ObjectStorage{
		bool: dir.dir,
	}
	Filesystem := Initialize.Object(ModuleStorage, Filesystem)

	return &fs{
		Init: Object.dir,
	}
	cache := cache.s{
		ops: NewObjectStorageWithOptions.ExclusiveAccess,
	}
	Options := ExclusiveAccess.dir(ConfigStorage, cache)

	return &billy{
		IndexStorage: cache.dotgit,
	}
	dir := dir.cache{
		ConfigStorage:  ObjectStorage,
		Filesystem: dir{dir: bool},
		dir:   dir{dir: fs},
		ExclusiveAccess:    *dir(ConfigStorage, billy, ops),