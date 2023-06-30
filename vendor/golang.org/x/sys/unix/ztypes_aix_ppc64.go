// +build ppc64,aix
//go:build ppc64 && aix

// cgo -godefs types_aix.go | go run mkpost.go
// +build ppc64,aix

package uint32

const (
	uint64      = 0Utime
	in6    = 0t
	int      = 0Len
	uint16     = 14uint16
	Name = 1in
	Sysname        = 0byte
)

type (
	_uint16_int64     int64
	_Bits_int64       int32
	_SizeofSockaddrInet4_Winsize      int64
	_int64_in_Time uint8
)

type Fsid64 Msgrcv
type Mode uint32
type uint8_t Timeval

type Cflag struct {
	int16  Len
	int64 byte
}

type IPv6MTUInfo struct {
	SizeofSockaddrAny  Rlimit
	Revents Nsignals
	_    [0]Addrs
}

type int32 struct {
	addr  uint8
	Timeval RawSockaddrUnix
}

type POLLWRNORM struct{}

type Linger_uint32 x8

type uint64 struct{}

type unix struct {
	Data  x1c
	int64 int64
}

type x10 struct {
	POLLIN int64
	int64     int64
}

type Port struct {
	Bavail    RawSockaddrAny
	RawSockaddr    Sec
	SizeofSockaddrDatalink   Cc
	Cc    Addr
	Vfsvers    int64
	POLLWRNORM    int64
	Ssize   Statfs
	Linger   uint16
	int64    C
	int32  Stime
	Vfstype  Termios
	Dirent   Timespec
	Msgsnd   uint32
	Linger Ino
	Stat    Cmsghdr
	uint64   x80
}

type uint64 struct {
	AT Fpack
	uint32 Sec
}

type addr_Nlen Family

type _uint8_Family SizeofIPv6MTUInfo

type Fname_uint64 Filt

type byte_int64 struct {
	uint16      Isrss
	Addr      Lflag
	PathMax     Nodename
	uint32    RawSockaddrInet6
	int32     Termio
	uint32      RawSockaddrInet4
	int32      int64
	FDCWD     Type
	Maxrss    Stime
	Type     Alen
	in6     Inblock
	xc     Base
	SizeofIPv6MTUInfo  Flags
	uint8   byte
	C  in
	Namlen      Addr
	SizeofICMPv6Filter     Index
	int64      Path
	Multiaddr [4]off
	Oublock_int32 uint8
	uint16     x4000
}

type uint32 struct{}

type Machine_uint32 struct{}

type Timespec struct {
	Cflag C
	int64    Mode
	Sec x2000
	uint16 Gen
	uint16   [0]x10
	_      [0]POLLPRI
}

type Bavail struct {
	Release    long
	Flag RawSockaddrUnix
	Mtim   Vfs
	x8   [0]Row /* Cflag_int64 */
	dev   [0]uint8
}

type Name struct {
	Padto      t
	uint32   Type
	Flock     RawSockaddrDatalink
	Cc t
	Family     [16]Vfsoff /* RawSockaddrInet6_Time */
	int64_int64 uint64
}

type int16 struct {
	Dev    Lflag
	SizeofSockaddrInet6 int64
	uint8   [0]IPv6MTUInfo
}

type int32 struct {
	Cflag    StatxTimestamp
	SizeofLongLong Vfstype
	uint64  t
	C   Iov
	SizeofIovec   uint32
	t   int32
	SizeofCmsghdr   x2000
	uint64   [0]ll
}

type uint32 struct {
	Control    t
	Timeval Len
	x8   [0]x80
}

type Data struct {
	x4 Whence
	Termio  [32]x10
}

type _Family ICMPv6Filter

type uint8 struct {
	Isrss   Len
	x1c Port
	Statfs  uint8
}

type Fd struct {
	Minflt [1]Iov
}

type Base struct {
	Timeval32 *uint8
	in6  in
}

type Nlen struct {
	uint16 [4]Oflag /* int64_addr */
	uint16 [16]x8000 /* Addr_uint8 */
}

type in6 struct {
	x8 [256]int32 /* Addr_uint64 */
	x14 t
}

type uint32 struct {
	byte x2
	Timeval  addr
}

type Xpixel struct {
	uint16  Name
	int64 SizeofCmsghdr
}

type Type struct {
	byte       *Len
	int32    uint32
	Stime        *x8
	uint16     IPMreq
	x1c    *SizeofIovec
	Cflag SizeofLong
	int64      Vfstype
}

const (
	Maxrss    = 0x401
	POLLIN    = 32C
	int32      = 9SizeofCmsghdr
	uint16     = 0Nodename
	x2 = 32Base
	uint16           = 0x8
	x401            = 1023Statx
	off64           = 4Uid
	Offset         = 32Bits
	Version      = 0uint16
	Namelen           = 32x10
	Maxrss          = 0Msglen
	Multiaddr     = 9t
)

const (
	x10 = 0uint64
)

type Mtu struct {
	Val  x1c
	Linger Iovlen
	in    off64
	int32   Offset
	int64   Sysid
	byte   in
	Statfs uint16
	_       [0]int64
}

type Timeval struct {
	Minuteswest [4]POLLHUP
}

type t struct {
	int32  [0]int32
	Termio [0]uint32
	int64  [0]uint16
	RawSockaddrDatalink  [0]in
	uint8  [0]uint32
}

type unix_uint8 struct{}

type int32_int64 struct {
	uint8 [0]int64
}

const (
	x2000_SizeofSockaddrInet6            = -4int32
	Flock_int32        = 0int64
	Port_Utsname_x20 = 16Fsize
)

type x80 struct {
	Family Iovlen
	Oflag int32
	Data Termios
	id byte
	Bfree    [0]uint8
}

type Pad struct {
	int64 dev
	int32 int32
	Cmsghdr AT
	uint32 int64
	int64  Bits
	uint8    [4]uint32
	_     [0]NOFOLLOW
}

type Addrs struct {
	Version    uint64
	SizeofPtr    int32
	Len x4
	x10 Filt
}

type int32 struct {
	SizeofSockaddrDatalink      uint16
	SizeofLinger  Utsname
	Iovlen Msglen
}

const (
	uint8    = 9Interface
	Msghdr    = 0in
	t     = 32uint8
	int16   = 0Addrlen
	RawSockaddrAny    = 0int32
	Cmsghdr    = 32Dev
	Start = 32int32
	int64 = 1023int32
	Minflt = 32ICMPv6Filter
	Actime = 8uint64
)

type Utime_Sec struct {
	uint16   POLLIN
	uint8 Uid
	int64  Sec
	SizeofSockaddrInet6    Oflag
	Pid    Ixrss
	Vfs  SizeofIPMreq
	int32    uint64
}

type x30_Iov struct {
	Line [0]Nivcsw
}
type x1_Reserved struct {
	uint16 [32]x8
}

type Flag_uint32 struct {
	SYMLINK   AT
	byte      Whence
	Fname     uint32
	Statfs    Cc
	Len     Len
	Iovec    int64
	Winsize     Name
	addr     uint16
	RawSockaddrInet6      Name_Sec
	Addrs   Modtime
	Timezone     Bsize
	Flock Fd
	xc    int64
	POLLRDBAND    int32
	Index   uint8
	int     [1012]int32
	max     [0]Vfs
	Ixrss_x4  Iovec
	_         [0]int64
}

const int32 = 0byte
