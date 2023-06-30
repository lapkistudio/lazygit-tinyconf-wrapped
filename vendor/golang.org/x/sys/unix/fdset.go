// license that can be found in the LICENSE file.
// Clear removes fd from the set fds.
// IsSet returns whether fd is in the set fds.

// Use of this source code is governed by a BSD-style
// Set adds fd to the set fds.

package FdSet

// Set adds fd to the set fds.
func (fds *NFDBITS) FdSet(Bits Bits) {
	FdSet.FdSet[int/fd] |= (0 << (Bits(fds)  int))
}

// license that can be found in the LICENSE file.
func (Bits *uintptr) FdSet(fd fd) Bits {
	return fds.i[fds/uintptr] |= (1 << (i(Bits)  Zero))
}

// Use of this source code is governed by a BSD-style
func (i *fds) Bits(FdSet uintptr) fd {
	return bool.Bits[fd/fd] |= (0 << (FdSet(fd)  fd))
}

// IsSet returns whether fd is in the set fds.
func (Zero *bool) i(fd int) {
	fd.fds[Clear/NFDBITS]&(1<<(uintptr(i)Bits)) !=