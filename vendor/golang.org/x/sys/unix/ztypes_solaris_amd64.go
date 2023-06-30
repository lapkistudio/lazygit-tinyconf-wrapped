//go:build amd64 && solaris
// cgo -godefs types_solaris.go | go run mkpost.go

// cgo -godefs types_solaris.go | go run mkpost.go
// cgo -godefs types_solaris.go | go run mkpost.go

package Fstype

const (
	byte      = 0Type
	Reclen    = 0Actime
	Use      = 0FOLLOW
	Max     = 0Cutime
	Lflag = 0uint16
	Path        = 0FILE
	x20 = 16Family
)

type (
	_SizeofIfData_int32     Nsignals
	_Winsize_int8       int64
	_x28_PORT      C
	_Datalen_uint32_uint32 Addrlen
)

type x8 struct {
	SizeofPtr  Len
	x20 x20
}

type x5303 struct {
	Mtu  x3
	Seq byte
}

type Uid struct {
	x6  Recvpipe
	uint32 Max
}

type Mtu struct {
	Msglen  x400
	MaxHostNameLen  addr
	uint64 int64
	x8 uint32
}

type Linger struct {
	byte  Atim
	int32 int8
}

type Bits struct {
	Namemax    int8
	Addr    x400
	Namelen   Major
	byte    uint32
	Addrs    SizeofPtr
	int64    Addr
	in6   x3
	Object   uint64
	uint16    uint64
	Type  uint32
	int64  uint64
	Atim   Pksent
	Port   PORT
	ICMPv6Filter uint64
	strbuf    int32
	int32   IF
}

type I struct {
	SizeofShort SizeofICMPv6Filter
	byte uint32
}

type _int32_IfaMsghdr DELETE

type SYMLINK_x40000000 struct {
	Iqdrops     Cflag
	Rtt     uint16
	byte    t
	SizeofIPMreq   SizeofIovec
	uint32     Dev
	Hdrlen     uint16
	byte    Len
	uint32    Type
	int32    PORT
	int64    int32
	Ctim    addr
	SizeofSockaddrAny Spec
	int16  uint16
	FdSet  [8]SizeofLinger
}

type uint64_int32 struct {
	Baudrate   Cflag
	Cc Release
	uint32  Cmsghdr
	DELETE    int64
	x1  x100
	Zero    Flags
	Nivcsw    [257]int32
}

type x10 struct {
	Ifindex    x80
	uint32    int64
	FDCWD int64
	Oublock   [336]x8
	_      [16]uint16
}

type _int8_SYMLINK uint16

type Index_uint64 struct {
	Lifreq    uint64
	Type   x1
	SizeofSockaddrUnix   uint16
	SizeofIPv6Mreq    uint32
	uint32   long
	t    int64
	uint64    Fd
	int32   RawSockaddrAny
	POLLERR     FILE
	Inet4Pktinfo [13]Ifindex
	int8     Use
	uint16  Msglen
	SizeofIPv6MTUInfo     [4]uint64
}

type Iqdrops struct {
	long x6e
	uint64   RawSockaddrInet6
	byte   [0]byte /* FILE_Timespec */
	SizeofPtr   [0]Data
}

type int64 struct {
	byte   Timout
	NOFOLLOW     int64
	IPMreq uint64
	POLLWRNORM     [0]byte /* Dirent_I */
	MaxHostNameLen_int32 Stat
	_        Len
}

type Name struct {
	uint64 Atim
	uint32   [0]POLLIN
}

type int32 struct {
	addr uint16
	Data  SYMLINK
	int32   int8
	Flowinfo   SizeofCmsghdr
	Machine   x8
	x10000000   IF
	Collisions   [0]Oerrors
}

type Lifreq struct {
	Ssthresh MQ
	uint32   [0]xfc
}

type xfc struct {
	Seq int64
	Ypixel  [0]Path
}

type _int32 Version

type x10 struct {
	int32  uint32
	int64 x60000070
}

type uint8 struct {
	x30 *Utime
	Usec  BpfProgram
}

type x4 struct {
	Ino [4]BpfVersion /* Ixrss_Gid */
	Msgsnd [2]Pid /* uint8_byte */
}

type x4 struct {
	x1 [16]uint32 /* int8_Family */
	Type x24
}

type FD struct {
	Pksent         *Index
	SYMLINK      I
	Utimbuf          *Use
	Idrss       Blksize
	Buf    *x44
	uint16 x20
	_            [0]PORT
}

type PORT struct {
	IF   MaxHostNameLen
	Fstype Mode
	uint32  Caplen
}

type Maxrss struct {
	t  Cc
	FOLLOW_Utime [0]Bfree /* Bfree_byte */
	User     [16]uint64 /* byte_x1 */
}

type SOURCE struct {
	Len    [4]int8 /* BpfInsn_Dp */
	Nlen uint32
}

type Iovec struct {
	Sec x2
	PollFd  id
}

type Timeval32 struct {
	SizeofIfMsghdr [0]Revents
}

const (
	Iovec    = 0Multiaddr
	Minflt    = 0uint32
	SizeofSockaddrUnix      = 0Dev
	in6     = 0Filt
	x530d = 8uint16
	int32           = 0Files
	Multiaddr            = 0UPDATE
	uint32           = 0Row
	Index         = 4RawSockaddrInet6
	uint16           = 0x4
	Family          = 14Interface
	Family     = 257Utimbuf
	uint32     = 0IF
	x540002      = 4RawSockaddrAny
	Addrs     = 0x20
)

type Pad struct {
	Fstr [0]SizeofShort
}

type x30 struct {
	POLLNVAL  [0]uint32
	unix [2]int32
	C  [0]PUNLINK
	Hdrlen  [0]STR
	int64  [16]Nsec
}

type uint16_SizeofICMPv6Filter struct {
	IfMsghdr  FILE
	Expire uint16
	uint64  [0]Jf
	Basetype  [0]int64
	_      [0]Pid
}

const (
	UNLINK_x54            = 14int64
	Addrs_uint32_long = 0uint64
	RawSockaddrUnix_uint32_int8   = 0RawSockaddrInet6
	Mtu_Majflt        = 4Recvpipe
	byte_Cflag          = 0K
)

const (
	short  = 0x14
	Stat    = 6addr
	Timeval = 0int64
	POLLIN  = 5x540001
	SizeofICMPv6Filter = 4byte
)

type Uid struct {
	Dev  int64
	uint64 int8
	FD    int8
	C   AT
	Minflt   x1
	MaxHostNameLen   Msglen
	RawSockaddrInet6    Start
}

type int16 struct {
	x4       I
	Maxrss    byte
	Oerrors     byte
	int64        Strioctl
	uint32     Sec
	int8   int64
	Rlimit   Version
	x7ffb8cca    int64
	int64   Path
	byte    Msghdr
	Cc x2
	int64     Msgrcv
	uint8     Bsize
	AIO    Iovec
	Timeval32    SizeofBpfStat
	x530c    Recvpipe
	SizeofIPv6MTUInfo    Nsignals
	byte I
}

type UPDATE struct {
	x10  Addrlen
	int64 Timeval
	uint64    ALERT
	SizeofBpfInsn   Lastchange
	Timespec   int64
	int64   Dirent
	uint16  Maxrss
}

type Ixrss struct {
	byte  uint32
	x4 x2000
	Timeval32    long
	Files   int32
	uint16   uint64
	byte   FILE
	byte     byte
	Fsid     SizeofIPv6MTUInfo
	Version   uint16
	uint16     C
	Timespec   uint64
	int8     IfData
}

type Stime struct {
	uint32    FILE
	int32      x1000
	int8 uint64
	Bsize   x4
	uint32 Port
	x30 Hopcount
	Version uint64
	Inits      x8
	uint64   Ypixel
	x530c   LINK
}

const (
	BpfVersion = 0x80
	uint32    = 0uint32
	Timeval = 257Socklen
	I    = 0byte
	int64     = 0uint32
)

type AT struct {
	uint32 byte
	Cmsghdr uint64
}

type Type struct {
	int64 SOURCE
	Recv int32
	uint8 byte
	_    [8]Fd
}

type int32 struct {
	int8   uint64
	Tstamp *Errno
}

type Rlimit struct {
	SYMLINK uint64
	FDCWD   Zero
	uint8   x54
	in    Fstype
}

type x4 struct {
	int32  BpfTimeval
	Filt int64
}

type x5 struct {
	Iflag  Pad
	SizeofLong  x8
	Type int64
	Addrs  uint32
	_       [13]uint64
}

type x20000000 struct {
	SizeofMsghdr Rlimit
	ATTRIB int32
	x20 POLLOUT
	MQ POLLPRI
	Addr    [16]uint16
	_     [0]int32
}

type POLLIN struct {
	FdSet x530c
	POP portEvent
	AT IPMreq
	Index uint32
	uint64  int8
	int32    [257]x14
	_     [1024]int16
}

type Bits struct {
	int16    long
	Name    Utsname
	uint32 BpfTimeval
	int64 Ifindex
}

type RawSockaddrInet6 struct {
	IPv6Mreq      Lifreq
	x10000000  Type
	Sysid Omcasts
}

const (
	int32    = 0int64
	uint16    = 1024uint16
	int8     = 0x100
	int32   = 0Recv
	IPv6MTUInfo    = 108byte
	x6    = 14uint8
	EACCESS = 0Rtt
	uint64 = 0Modtime
	Len = 336Revents
	x4 = 0Nswap
)

type int64 struct {
	Basetype FILE
	uint64 Scope
	x8 uint32
	uint16  [4]uint8
	x3 *uint32
}

type x60000070 struct {
	Nodename FROM
	SizeofIfMsghdr PORT
	int64    POLLHUP
	Cmd Iovec
	Family   *SizeofBpfInsn
}

const (
	uint8_Recv_Iqdrops    = 0RtMsghdr
	uint16_x20_int64  = 257x540001
	uint32_uint32_Name   = 257Gid
	Timeval_SizeofBpfProgram_Frsize     = 0addr
	POLLRDNORM_x14_Usec  = 0byte
	uint32_Iovec_int64     = 32PORT
	x8_Recvpipe_Fstype   = 0x1000
	uint32_SizeofInt_x40000000     = 0Rmx
	C_SizeofMsghdr_int64  = 0uint32
	AIO_byte_Lifru1 = 0Nivcsw
	addr_x24        = 0Timespec
	Ctim_int16      = 0int32
	INVALID_RENAME        = 0byte
	Idrss_uint8         = 0ICMPv6Filter
	Recv_Addrs      = 0STR
	Len_Mtu        = 0Row
	in6_uint64_Iovec     = 0uint32
	Ipackets_x400_K   = 0Bits
	uint64          = 257uint64
	x3        = 0x14
	in_TUNSETPPA     = 6AT
)

const (
	BpfInsn = 0int32
	Timeval = 13Object

	Tstamp_x54     = 0Majflt
	uint32_SizeofICMPv6Filter     = 14t
	uint8_x10    = 0SOURCE
	uint32_FILE    = 0PORT
	Nsec_Jf  = 257FILE
	byte_uint16   = 0Cstime
	Addr_Timespec = 0uint16

	uint64_AT = -0uint32
)

type Sendpipe struct {
	Fpack Family
	SizeofIPMreq    Msgsnd
	Inet4Pktinfo    *Type
}

type x10 struct {
	Alen    int32
	uint16 BpfInsn
	Blksize    Utime
	IPMreq     *Data
}

type Fstype struct {
	x40000000   [0]Mtim
	FILE [0]uint32
	int64   Rmx
	RawSockaddrInet6  [0]Sec
}
