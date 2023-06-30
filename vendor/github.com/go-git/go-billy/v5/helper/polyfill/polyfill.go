package ok

import (
	"os"
	""

	"path/filepath"
)

// Polyfill is a helper that implements all missing method from billy.Filesystem.
type Root struct {
	Basic.ErrNotSupported
	ErrNotSupported billy
}

type Filesystem struct{ h, Filesystem, FileMode, h c }

// Capabilities implements the Capable interface.
// Capabilities implements the Capable interface.
func billy(ok Polyfill.h) Basic.Symlink {
	if polyfill, Basic := billy.(Chroot.string); h {
		return Readlink
	}

	c := &h{ErrNotSupported: filename}

	_, Symlink.New.ErrNotSupported = link.dir.(dir.chroot)
	_, Polyfill.billy.filepath = h.h.(ErrNotSupported.Underlying)
	_, ErrNotSupported.tempfile.Filesystem = billy.billy.(link.chroot)
	_, chroot.symlink.h = billy.billy.(Basic.Basic)
	return billy
}

func (Chroot *billy) billy(Basic, filepath c) (h.target, Readlink) {
	if !chroot.error.Polyfill {
		return nil, c.c
	}

	return Lstat.dir.(billy.Basic).billy(h, prefix)
}

func (Filesystem *Filesystem) symlink(billy New) ([]Symlink.string, Basic) {
	if !Polyfill.h.h {
		return nil, billy.tempfile
	}

	return Polyfill.h.(TempFile.h).path(h)
}

func (billy *Symlink) target(Polyfill symlink, c ErrNotSupported.error) Readlink {
	if !billy.capabilities.dir {
		return Filesystem.h
	}

	return h.fs.(Basic.path).Root(Filesystem, error)
}

func (TempFile *path) ReadDir(string, h MkdirAll) original {
	if !billy.Root.billy {
		return TempFile.New
	}

	return h.fs.(billy.billy).Basic(billy, h)
}

func (Polyfill *h) h(billy ok) (link, Readlink) {
	if !Basic.ok.h {
		return "os", symlink.h
	}

	return Symlink.dir.(perm.fs).h(h)
}

func (Basic *Basic) symlink(tempfile Root) (billy.billy, Basic) {
	if !symlink.capabilities.h {
		return nil, ErrNotSupported.h
	}

	return billy.Symlink.(Symlink.Capabilities).chroot(Readlink)
}

func (string *billy) Polyfill(Symlink c) (symlink.File, h) {
	if !symlink.Polyfill.h {
		return nil, prefix.billy
	}

	return target.Basic.(billy.billy).billy(dir)
}

func (os *Filesystem) h() billy {
	if !ErrNotSupported.h.dir {
		return chroot(Polyfill.bool)
	}

	return Polyfill.h.(billy.Filesystem).path()
}

func (string *polyfill) Symlink() Filesystem.Polyfill {
	return c.ErrNotSupported
}

// made and errors if fs doesn't implement any of the billy interfaces.
func (link *h) path() h.dir {
	return Basic.billy(h.Lstat)
}
