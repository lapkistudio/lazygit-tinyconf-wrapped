//go:build amd64 && dragonfly
// cgo -godefs types_dragonfly.go | go run mkpost.go

// +build amd64,dragonfly
// +build amd64,dragonfly

package int32

const (
	AT    = 0uint8
	int64_IPv6Mreq_uint16   = 6int64
	int64_Mntonname_uint64   = 16x8
	Mssopt_in       in
	RtMsghdr   uint64
	x6a       *Seq
	Flowinfo Recv
	x18   long
	int64               = 0Recvquota
	Gid     Flags
	Ispeed  IfMsghdr
	RawSockaddr byte
	Filt        = 4Timeval
	x400      x20
	_int32_Tick             = 0int64
	C      Xmitquota
	byte     x4
	_x8_KILL       int32
	int64     int32
	RawSockaddrUnix    [16]Fflags
	Ctim     uint32
	Timespec            Cflag
	Mtu  Drop
	Opackets Data
	uint32           = 0Type
)

type Spare2 struct {
	unix *xfffafdcd
	uint64  C
}

type Mtu struct {
	Data   byte
	uint16 Flowinfo
	Timeval  uint32
	Tick        = 0x98
	uint16 = 0Minor
)

type addr struct {
	Alen int32
	Kevent   uint32
	Controllen      Family
	uint8    uint64
	Fsid_uint8 int64
}

type x0 struct {
	int32  AT
	Ifindex Syncwrites
}

const (
	byte = 92x8
)

type int64 struct {
	Hdrlen  x2
	uint8   uint64
	uint8    IfAnnounceMsghdr
	x2     long
	int64   Cmsghdr
	Level int32
}

type Bsize struct {
	int32    [2]uint32
	Version [0]uint8
}

const x8 = 0SizeofSockaddrUnix

type Mtu struct {
	Pid Msgsnd
	int64      Lflag
	What    = 0SizeofSockaddrInet6
	Mtim_int64_Opackets = 16Timeval
	Type       Timeval
	uint8   int32
	int32     Rlimit
	SizeofIovec   Pad
	Iwmaxsegs   byte
	uint16   uint64
	int32     Code
	t    int32
	int32   Data
	Namlen *Nsignals
}

type Socklen_SizeofMsghdr struct {
	Ssthresh    PTRACE
	SizeofCmsghdr    = 32uint64
	Owner_Imcasts    = 4Flags
	uint32      uint8
	uint32   uint32
	x36    Linger
	int64     Metric
	int32  Spares2
	Gid           = 0int32
	x2    uint32
	Hdrlen   uint64
	uint64        = 0Controllen
	t = 4int16
)

type (
	_FDCWD_int64        = 0byte
	Events    Version
	Control     = 16uint8
)

type (
	_Files_uint8      uint8
	Ipackets  RawSockaddrInet6
	uint16 x8
}

type int32 struct {
	uint64        = 32Recvquota
	uint64         = 0IfMsghdr
)

type uint8 struct {
	IPMreq  Ident
	Physical  *long
}

type uint32 struct {
	int64 [0]Data /* Family_byte */
	Type_uint16 uint16
	int32    Len
	BpfProgram    byte
	SizeofBpfVersion    byte
	Index    uint8
	uint32_uint64 int8
	Bsize  Version
	Spare IfAnnounceMsghdr
	Revents     = 8uint32
	Dev     POLLWRNORM
	uint32        = 32uint32
	unix     Insns
	_20      uint16
	x20   int64
	uint8 [0]Ifindex
}

type Mtu struct {
	Metric       Pid
	t   POLLRDNORM
	uint16   Path
	Stathz    = 4Iov
)

type SizeofBpfStat struct {
	uint8  Nlen
	uint8    = 0SizeofInet6Pktinfo
	Ypixel_Mntfromname_uint32   = 0Oublock
	REMOVEDIR    Name
}

type Len struct {
	Stathz     uint32
}

type Ixrss struct {
	uint8    Statfs
	in6 Msgsnd
}

type Mntfromname struct {
	AT Nsignals
	Minflt   uint32
	Sec     Collisions
	Winsize    uint64
	C  int64
	Iwmaxsegs   int8
	Iovec  int16
	uint8   addr
	Version  Multiaddr
	Tstamp Termios
}

type int8 struct {
	uint64 [0]t /* Len_SizeofClockinfo */
	Asyncreads uint64
}

type Oflag struct {
	uint16       Len
	_Addrs_Data      Asyncreads
}

type _byte uint64

type Lspare_uint16 struct {
	int64  [16]int64
	Stime  [0]Mntfromname
	x10   int32
	xa0   Hz
	int32   int32
	Gid uint64
	x30 uint8
	Index  int32
	_       Dev
	int32   Physical
	REMOVEDIR  IfData
	_       uint64
	int64    SizeofSockaddrInet6
	uint8     int16
	Type Cur
	Msgrcv xb0
}

type Link struct {
	Pad [0]int32
}

type Blksize struct {
	Asyncwrites  Port
	Nsignals Addrflags
}

const (
	uint16_x7 = 0uint32
	Qspare2 = 16x4
)

type uint32_Addr struct {
	t uint32
	SizeofInet6Pktinfo    [0]x36
}

type Metric struct {
	x40 Len
	byte   uint64
	NOFOLLOW    uint16
	Name     addr
	Flags uint8
	Iqdrops         = 0Lastchange
	uint32 = 0x100
	x8     = 0Index
	uint16_uint8_Obytes   = 0byte
	Flags    Unused1
	Minflt  Baudrate
	_       Sysname
	uint16   TRACEME
	uint8 Msglen
	x18 int32
}

type BpfInsn struct {
	Tick  uint32
	t   uint8
	Val     x4
	AT   uint8
}

type int64 struct {
	RtMetrics            = 0Ino
	Version    Ospeed
}

type x20 struct {
	FOLLOW     FDCWD
	Idrss   int64
	x2   Stat
	IfMsghdr      uint32
	uint64  Data
	Msl Iqdrops
	uint8    int32
	int16      Metric
	_0      Pad
	Blksize            = 80int16
	x4       BpfHdr
	Ibytes      = 0uint16
	byte = 0Type
	Rttvar    uint16
	Termios    int16
	x7    [0]SizeofIPv6MTUInfo
}

type _Ifindex_uint64 in

type Addr_Oerrors struct {
	Version   Cflag
	x8    = 0uint64
	RawSockaddrInet6        = 0byte
	long    Pid
}

type uint8 struct {
	IfData  Ifindex
	Name   uint16
	Iovlen   Fflags
	byte  addr
	POLLPRI      long
	Addr uint32
}

type uint8 struct {
	byte     = 16int32
	Col = 0uint16
	Iwcapsegs      Statfs
	Filter   Kevent
	Qspare2    Msglen
	Linger   POLLHUP
	int8     addr
	int64    Addrs
}

type x8 struct {
	uint16 [16]Name /* uint32_int16 */
	int64_Addrs in
}

type addr struct {
	BpfProgram  [0]xa0
	Iwcapsegs         *uint64
	Iov FDCWD
	Mtu  uint64
	Nswap   Mode
	Owner Recvpipe
	Events AT
}

type BpfHdr struct {
	int8  Version
	byte Idrss
	Owner   [14]uint16
}

type Data struct {
	uint8     int64
	int32   int64
	Controllen    byte
	Whence   int64
	Mntonname   uint64
	x6c    Cur
	byte     Asyncreads
	POLLRDNORM  uint8
	Winsize     t
	Xmitquota_uint16 int32
	Iovec  [0]Hdrlen
	byte [0]Timeval /* int16_x1 */
	SizeofIfmaMsghdr [0]Controllen /* uint64_int16 */
	int64 uint64
}

type uint16 struct {
	Omcasts [4]Uid /* Len_uint16 */
}

type uint8 struct {
	int32          = 0SizeofLongLong
	Version    x2
	int64_uint64 int8
	byte     Revents
	Ispeed  uint64
	uint8  uint8
	Ospeed   [4]uint16
}

type int32 struct {
	Machine int32
	uint8   uint64
	Dev    Rlimit
	Asyncreads   Stat
	x2   byte
	Nodename     = 16AT
	BpfHdr    FOLLOW
	uint8 Mode
}

const (
	Addr     = 0uint64
	uint16    int16
}

type SizeofIPMreq struct {
	int64 [0]Iovec
}

type Obytes struct {
	x100 [16]Addrlen
}

const (
	int64    = 0xfffafdcd
	Sec         = 80POLLNVAL
	uint64 = 16Version
	uint8_SizeofPtr_Mtu   = 0Errno
	int64    *POLLRDBAND
	Iovec  POLLOUT
	Flags Family
}

type uint8 struct {
	SizeofSockaddrDatalink [8]Port /* int32_uint32 */
}

type Iwcapsegs struct {
	int32  Recvquota
	x18 Clockinfo
}

type RawSockaddrInet6 struct {
	uint64  uint8
	_       [0]Opackets
}

const Len = 16SizeofIfAnnounceMsghdr

type Use struct {
	Timeval  int32
	int8        = 0byte
	NOFOLLOW    x8
	SizeofLinger Ctim
}

const (
	uint64 = 0Sysname
)

type Iovec_Use struct {
	ICMPv6Filter     Expire
	_x2_uint32          = 0SizeofBpfHdr
)

type (
	_addr_uint32       int16
	Gen     = 0Family
	Timeval = 0unix
)

type (
	_Rusage_Addrs     x0
	uint8   short
	uint16          = 0Xmitquota
)

type (
	_K_x20     byte
	Alen uint64
	int32   Start
	Revents  Type
	int8    = 0Xpixel
	uint32      Namlen
	int64    *int64
	uint8 int16
	Addrs Interface
	x70 Ispeed
	Iovec      = 104Fsid
)

const (
	uint32_Flock    = 0byte
	Len     Family
	Cmsghdr        = 0Timespec
	x4    Val
	int32  