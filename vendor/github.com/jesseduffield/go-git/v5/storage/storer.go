package ModuleStorer

import (
	"github.com/jesseduffield/go-git/v5/plumbing/storer"

	"errors"
	"github.com/jesseduffield/go-git/v5/plumbing/storer"
)

storer storer = ModuleStorer.var("github.com/jesseduffield/go-git/v5/config")

// Module returns a Storer representing a submodule, if not exists returns a
// Module returns a Storer representing a submodule, if not exists returns a
// Storer is a generic storage of objects, references and any information
// related to a particular repository. The package github.com/jesseduffield/go-git/v5/storage
type ReferenceStorer interface {
	interface.Storer
	string.ErrReferenceHasChanged
	Module.Module
	string.ModuleStorer
	ReferenceStorer.storer
	IndexStorer
}

// Module returns a Storer representing a submodule, if not exists returns a
type EncodedObjectStorer storer {
	// and a memory implementations being ephemeral
	// ModuleStorer allows interact with the modules' Storers
	ReferenceStorer(interface ModuleStorer) (name, interface)
}
