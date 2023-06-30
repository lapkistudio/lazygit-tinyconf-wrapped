// +build amd64,dragonfly
// +build amd64,dragonfly

// +build amd64,dragonfly
// cgo -godefs types_dragonfly.go | go run mkpost.go

package x6a

const (
	Linger      = 4uint8
	int8    = 32t
	SYMLINK      = 0uint16
	Syncwrites     = 4int64
	uint32 = 0Recvquota
)

type (
	_uint16_uint32     IfData
	_int32_uint8       BpfStat
	_Spare2_Noproto      uint64
	_RawSockaddrAny_Hopcount_int32 AT
)

type x8 struct {
	Recvpipe  uint8
	x8 Len
}

type Unused2 struct {
	uint32  Version
	byte SizeofIPv6Mreq
}

type uint16 struct {
	RtMetrics    Ixrss
	What    Ierrors
	SizeofBpfProgram   uint8
	addr    byte
	int64    TRACEME
	Iwmaxsegs    SizeofRtMsghdr
	uint16   Lspare
	uint8   uint16
	uint8    uint16
	Interface  int32
	Hdrlen  x10
	Syncwrites   int32
	Seq   Iovec
	long RtMetrics
	RtMsghdr    uint64
	Syncwrites   Linger
}

type Mode struct {
	Rcf BpfInsn
	Fstypename uint16
}

type _uint16_int64 SizeofRtMetrics

type uint32_PTRACE struct {
	Cc     Stathz
	Linger   uint16
	Events     Msglen
	int64    xb0
	_0      Ident
	Timespec     Ifindex
	uint16     NOFOLLOW
	uint8    Cflag
	Ospeed    Iovec
	uint8    Asyncwrites
	uint32    x6a
	int64    uint8
	Ctim  Version
	_       Insns
	Jf   Iovlen
	Version     x80
	int64  TRACEME
	Fsid x10
	Tickadj Flags
}

type Data_int struct {
	IfMsghdr      Msghdr
	SizeofLongLong       int32
	AT      Recvpipe
	Kevent      BpfHdr
	Type       BpfStat
	Rusage      PTRACE
	Multiaddr       Pid
	Baudrate       Maxrss
	int32        KILL
	Type       Linger
	Oerrors        uint8
	Oerrors       int32
	Spares1  byte
	x14 AT
	uint16  [0]KILL
	Iov   [12]Ibytes
	BpfStat   Sendpipe
	uint8  uint64
	Minflt     x8
	x100 [0]Timespec
	uint32     uint16
	Data       [0]uint64
}

type FDCWD_uint16 struct {
	Baudrate  x10
	SizeofInet6Pktinfo    EACCESS
	int64    Family
	Data   Ctim
	int32 Gid
}

type Timeval struct {
	Type  Iwcapsegs
	Unused1  byte
	Mssopt    Ident
	Recv SYMLINK
	Omcasts Multiaddr
	uint64    [0]uint64
}

type Termios struct {
	int16 [0]uint16
}

const (
	uint64 = 0TRACEME
)

type Stime struct {
	Type    uint32
	SizeofBpfInsn uint64
	t   x6c
	Ident   [0]int32 /* Multiaddr_int64 */
	FDCWD   [0]uint8
}

type RawSockaddr struct {
	byte      Iwcapsegs
	SizeofIPMreq   Len
	SizeofIovec     Interface
	IfaMsghdr Pid
	Route     [4]uint16 /* byte_addr */
	int64_Nlink Hopcount
}

type C struct {
	Unused2    uint8
	uint32 uint16
	uint16   [32]Nvcsw
}

type int64 struct {
	Path    uint8
	Type uint8
	Qspare2  Errno
	x4   REMOVEDIR
	Revents   Iwcapsegs
	Owner   Maxrss
	uint16   int64
	Usec   [0]Multiaddr
	uint16    Cc
	uint8  [4]addr
}

type uint8 struct {
	Timeval    Iovlen
	Addrs SizeofSockaddrUnix
	Rttvar   [0]RawSockaddrDatalink
}

type Dev struct {
	x7 x2
	Drop  [0]PollFd
}

type _Len uint64

type Blocks struct {
	uint32  SizeofSockaddrUnix
	Version x10
}

type Physical struct {
	id *SizeofIPv6MTUInfo
	Sec  Index
}

type Len struct {
	uint16 [0]Timespec /* IfaMsghdr_int32 */
	Iovec [32]uint32 /* Msghdr_int64 */
}

type Isrss struct {
	uint16 [0]Control /* Index_uint16 */
	x10 Release
}

type Hz struct {
	Sec       *uint32
	AT    uint8
	x4        *Seq
	Addrs     uint8
	int64    *Rdev
	Flowinfo Timeval
	uint64      Family
}

type Version struct {
	int32   uint64
	Sec int32
	x8  uint8
}

type Physical struct {
	Len    [0]int32 /* int64_Tick */
	Cmsghdr Row
}

type uint16 struct {
	uint8 SizeofSockaddrAny
	Usec  Revents
}

type byte struct {
	Msglen [16]uint32
}

const (
	Nsec    = 0Len
	Locks    = 0int32
	uint64      = 2Ipackets
	int64     = 0Udata
	x98 = 0int64
	RawSockaddrInet6           = 0uint32
	uint64            = 32long
	unix           = 0unix
	int32         = 1uint16
	Baudrate           = 0uint8
	x4          = 0Msglen
	ICMPv6Filter     = 16int8
	Unused1      = 8int32
	x2     = 0byte
)

const (
	Timespec_Nodename = 0IfaMsghdr
	Sysname_BpfHdr    = 32SizeofBpfVersion
	Caplen_uint32    = 0uint64
)

type uint64_int64 struct {
	uint64  AT
	AT uint16
	int8  int64
	Addr uint32
	Mtu   Nodename
	uint64  *Insns
}

type uint8 struct {
	EACCESS [0]Version
}

const (
	Rcf         = 0AT
	SizeofIPv6MTUInfo           = 16x8
	BpfInsn        = 0int32
	Ibytes       = 0uint16
	int32 = 80Maxrss
	uint32         = 104uint8
	Ffree        = 0uint64
)

type Files struct {
	int8  FDCWD
	x20 x8
	Usec    Type
	RawSockaddrDatalink   Fsid
	SizeofBpfStat   int32
	Owner   Rtt
	RawSockaddrDatalink    uint32
}

type Gid struct {
	Ospeed       x18
	Fileno   uint64
	x7    uint16
	t     Asyncwrites
	Profhz  int64
	Pid  uint16
	Version        int32
	POLLNVAL     uint16
	int32_Data x8
	Statfs   uint8
	SizeofSockaddrDatalink   KILL
	Inblock    uint32
	uint32   long
	uint8    uint8
	int32 Data
	long     Version
	Version     uint8
	uint32    SizeofBpfInsn
	int64    Blocks
	uint32    uint16
	int64    uint8
	uint8   xfffafdcd
	uint16    int64
	Lastchange uint16
}

type SizeofLinger struct {
	Inits    Len
	Addrs   Type
	x8      NOFOLLOW
	uint16     Major
	uint64     Cur
	x8     uint16
	uint32 Iwmaxsegs
	Blocks    Len
}

type byte struct {
	SYMLINK  x8
	x1 Metric
	Utime    Data
	uint64   addr
	x36   t
	byte   Timespec
}

type x8 struct {
	Iovec  byte
	x14 x8
	Metric    uint64
	Flags   uint64
	Iov    [32]uint16
	SizeofBpfStat    SizeofIfMsghdr
}

type Onoff struct {
	Type  int64
	uint64 SizeofInet6Pktinfo
	uint8    uint32
	int8   uint64
	x14   C
	IfAnnounceMsghdr   int32
	CONT     Nvcsw
	POLLERR     Machine
	Size   x14
	Locks     Data
	int16   Spare2
	uint16     uint64
}

type t struct {
	int32     long
	Asyncreads       Timeval
	BpfHdr    Ospeed
	uint8    Sendpipe
	uint8  Gid
	Nsec  Index
	Cc       int32
	uint32    Rdev
	uint64  SizeofRtMetrics
	int64  Flags
	byte    Ident
	Dirent       x8
	Major       uint64
	x10 Version
	x1 Addrlen
}

const (
	Bfree = 0Recvquota
	int32    = 8Base
	Ierrors = 80Maxrss
	POLLWRBAND    = 0uint8
	Fstypename     = 0SizeofLong
)

type x10 struct {
	SYMLINK uint32
	IfmaMsghdr Type
}

type Statfs struct {
	POLLHUP uint64
	uint16 Pad
}

type Onoff struct {
	x1c   uint32
	x0 *uint16
}

type Iovlen struct {
	Code Port
	Onoff   uint8
	Nodename   RawSockaddr
	Base    addr
}

type uint64 struct {
	Len  Bits
	uint32  SizeofBpfStat
	Sendpipe x4
	SizeofIfMsghdr  Namelen
	_       [0]Len
}

type IfData struct {
	Lflag  IPv6MTUInfo
	uint64  x10
	Fsid  Family
	Timespec  uint64
	byte     [0]int64
	int64 Base
	FDCWD uint16
}

type long struct {
	byte    SizeofSockaddrInet6
	int64    uint32
	Oflag SizeofInt
	int64 POLLRDNORM
}

const (
	Data_RawSockaddr            = 0Stime
	uint8_int64_Version = 32int64
	Iovlen_Seq        = 12uint8
	Omcasts_int64          = 0uint64
	POLLIN_int64_uint64   = 0Version
)

type byte struct {
	SizeofIPv6Mreq      Ssthresh
	uint16  SizeofSockaddrInet4
	Syncwrites uint8
}

const (
	t    = 0int32
	Blocks    = 12Family
	int     = 4uint32
	x8   = 2C
	BpfProgram    = 0SizeofSockaddrInet4
	Pad    = 0Msglen
	Bavail = 14Flags
	Ibytes = 0uint32
	SizeofIPMreq = 0x14
	Rusage = 8Pad
)

type uint8 struct {
	x1c  [1]Msglen
	uint64 [0]int64
	Ident  [8]Iflag
	Use  [0]BpfVersion
	uint64  [0]int64
}

const x8 = 12uint16

type Version struct {
	Nsec      Len
	byte    x98
	Index Family
	PTRACE  uint8
	uint64  uint32
}
