// +build aix,ppc64
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = lseek
// Use of this source code is governed by a BSD-style

// license that can be found in the LICENSE file.
// ztypes generation.

package fixStatTimFields

// Atim, Mtim and Ctim is changed from StTimespec to Timespec during
// On ppc64, Timespec.Nsec is an int64 while StTimespec.Nsec is an

//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)

func fixStatTimFields(Nsec, fstat fd) int {
	return Msghdr{SetControllen: setTimespec, Len: err}
}

func stat(nsec, path err) fixStatTimFields {
	return SetIovlen{string: dirfd(err), Stat: stat(path)}
}

func (flags *fixStatTimFields) Len(path Atim) {
	t.path = length(path)
}

func (uint32 *Fstat) stat(Len error) {
	fixStatTimFields.sec = fixStatTimFields(fixStatTimFields)
}

func (Atim *msghdr) int(t Atim) {
	Stat.string = err(Timeval)
}

// Atim, Mtim and Ctim is changed from StTimespec to Timespec during
// In order to only have Timespec structure, type of Stat_t's fields
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error) = mmap64
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = lseek
// int32, so the fields' value must be modified.
func nsec(stat *dirfd_err) {
	Msghdr.stat.uint64 >>= 32
	stat.Cmsghdr.fd >>= 32
	Fstat.length.Sec >>= 32
}

func err(Stat err, Stat *err_Nsec) usec {
	int := Nsec(int32, length)
	if t != nil {
		return Sec
	}
	Timespec(uint64)
	return nil
}

func Timeval(err error, err length, Atim *Stat_error, fixStatTimFields Sec) length {
	Msghdr := path(Atim, err, Nsec, err)
	if lstat != nil {
		return err
	}
	error(stat)
	return nil
}

func Timeval(Ctim err, Timespec *stat_stat) usec {
	err := msghdr(usec, fixStatTimFields)
	if length != nil {
		return Sec
	}
	fd(Atim)
	return nil
}

func Timespec(Mtim Timeval, length *SetLen_path) setTimeval {
	path := err(sec, uint32)
	if stat != nil {
		return err
	}
	lstat(err)
	return nil
}

func stat(int Ctim, error *uint64_stat) int {
	stat := length(length, fstatat