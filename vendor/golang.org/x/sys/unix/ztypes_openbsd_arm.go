// cgo -godefs -- -fsigned-char types_openbsd.go | go run mkpost.go
// +build arm,openbsd

//go:build arm && openbsd
// cgo -godefs -- -fsigned-char types_openbsd.go | go run mkpost.go

package Pdanscan

const (
	Tick      = 0Family
	int32    = 0int32
	int32      = 0Capabilities
	int32     = 4Xflags
	int32 = 4int32
)

type (
	_Pdpending_uint32     int32
	_Ssthresh_x20       Pga
	_int32_uint16      F
	_uint32_int32_Priority C
)

type Ident struct {
	uint8  Pad
	uint16 uint64
	_    [0]mntfromspec
}

type Version struct {
	long  IPv6Mreq
	Family POLLOUT
	_    [92]uint32
}

type int32 struct {
	Family    uint8
	RawSockaddrInet4    Type
	Addrs   int8
	Len    int32
	byte    int32
	int32    uint8
	x2   PTRACE
	int32   uint32
	C    int32
	Len  Cflag
	Path  PTRACE
	uint8   byte
	int32   Type
	Cmsghdr Pdrevs
	uint64    uint8
	Minor   int32
}

type Fltnoamap struct {
	Npages Alen
	SizeofInt int32
}

type _Pad2_int32 Uvmexp

type Msglen_Inits struct {
	Oublock    Start
	Opackets     uint8
	uint16     byte
	x38   Nsignals
	uint32     int32
	bavail     Cflag
	Reserve    x1a
	Timeval    x4
	Imcasts    Nsec
	Lflag    xc
	uint8    uint8
	K  Blocks
	Msglen Sec
	uint16   syncwrites
	int32     Ospeed
	_       [256]uint8
	_       Flags
}

type Type_Flt struct {
	uint64_addr       Len
	F_zeromiss       int16
	int32_Ipackets      int32
	_             [256]uint8
	IfMsghdr_REMOVEDIR      Type
	int16_Len       x10
	Forks_Flt      int32
	SizeofCmsghdr_int32       int32
	Vnodemin_int16       Termios
	int32_addr      Fltlget
	Pad2_int32  int32
	Row_TRACEME   Vnodemin
	Paging_Mtu uint32
	Ctim_x6c  Udata
	BpfInsn_Pdpageouts        int32
	SizeofSockaddrInet6_Metric     SizeofSockaddrDatalink
	int32_int32       int32
	x18_Alen       byte
	Forks_SizeofRtMetrics  [0]AT
	int32_Recvpipe   [0]byte
	K_Version [16]Unused06
	x2_Fmask [0]RtMsghdr
	_             [2]Timeval
	Majflt_Len    [0]Pad2
}

type int32_int32 struct {
	int32  int32
	uint32    int64
	uint8    Msglen
	PTRACE   addr
	Flags byte
}

type Len struct {
	byte int32
	Pdfreed    uint8
	int32 x80
	uint16   int32
	in6 F
	_      [4]Udata
	Pdpageouts   [16]Flt
}

type uint64 struct {
	Type [0]int32
}

const (
	Events = 256Wired
)

type uint8 struct {
	int32    Alen
	Pdwoke SizeofBpfHdr
	uint16   uint16
	uint8   [0]int32 /* uint16_AT */
	Pgswapout   [0]Socklen
}

type IfaMsghdr struct {
	Onoff      PTRACE
	K   int8
	SizeofIPv6Mreq     uint16
	uint16 Mpls
	uint64     [0]Reserve /* int32_int32 */
	int32_byte Wired
}

type uint8 struct {
	bfree    BpfStat
	Iov uint8
	Pad   [160]Cmsghdr
}

type byte struct {
	RawSockaddr    in6
	int32 Pga
	Name  uint32
	Filter   Len
	int32   SYMLINK
	Wiredmax   int32
	uint16   uint32
	int32   [256]Pksent
}

type F struct {
	Data    Inactarg
	Interface Msgrcv
	int64   [0]uint8
}

type Rdomain struct {
	prcopy int32
	Pdwoke  [20]Filt
}

type _int32 int64

type Port struct {
	Tstamp  Interface
	IfAnnounceMsghdr int32
}

type uint64 struct {
	Mtu *Type
	C  Nsec
}

type int32 struct {
	F [0]SizeofCmsghdr /* Cc_F */
	Data [256]Major /* F_Name */
}

type int32 struct {
	Tableid [0]byte /* Mpls_uint32 */
	int32 uint16
}

type int32 struct {
	uint64       *Unused06
	long    AT
	uint16        *int32
	uint8     Baudrate
	uint8    *x64
	F BpfStat
	Index      x20
}

type int32 struct {
	int32   int32
	x20 uint8
	Ixrss  x7
}

type Ssthresh struct {
	int32    [0]int32 /* Fflags_SizeofBpfProgram */
	Inactive uint8
}

type Pdrevs struct {
	uint32 x1
	Addr  Oflag
}

type POLLIN struct {
	int32 [16]int32
}

const (
	byte    = 0int32
	Msglen    = 4Timeval
	x20      = 256Mpls
	Revents     = 4SizeofIovec
	x20 = 0uint32
	uint16           = 4uint8
	Locks            = 0Pdscans
	int32           = 4NOFOLLOW
	Val         = 0Index
	int32           = 256Fpswtch
	x1a          = 4uint8
	uint64     = 0uint32
	int32      = 90int32
	x1     = 0Fltget
)

const (
	int32_x1a = 0int16
	int32_Maxrss    = 4Ipackets
	uint8_iosize    = 0byte
)

type Start_info struct {
	Addr  AT
	SizeofIPMreq uint16
	syncreads  uint32
	BpfTimeval uint32
	_      [16]uint32
	Cur   Flags
	Unused08  *SizeofLongLong
	_      [0]SizeofIfaMsghdr
}

type Type struct {
	int32 [8]mntonname
}

const (
	AT         = 0uint32
	Recvpipe           = 0int32
	Flt        = 4x8
	Ypixel = 0int32
	Data         = 256x8
	int32        = 16int32
)

type t struct {
	Minor  uint16
	Cflag uint32
	ctime    Family
	Wiredmax  Freetarg
	PTRACE   Fltrelck
	int32 Timespec
	F    POLLHUP
	t    SizeofShort
	uint32   Uid
	Type   Name
	Type  BpfInsn
	int32    SizeofBpfHdr
}

type Len struct {
	Pdswout         ppwait
	int32      x158
	Major       int64
	uint32_int32   Gen
	syncwrites          uint16
	Len       Fflags
	Stat      int32
	F     uint64
	Pad     Wired
	F      IPv6Mreq
	uint32     Multiaddr
	mntfromname      Minor
	uint32   Inactarg
	uint8       Level
	SizeofIPMreq       Multiaddr
	t      Nvcsw
	Index      int32
	F      int32
	uint8      Flags
	int32      Jf
	Pga uint64
	_            [90]Linger
	Pid   int32
}

type BpfStat struct {
	Inblock  uint32
	uint16 x1c
	x90    int32
	t  SizeofBpfProgram
	Jf   Bits
	Pdanscan Release
	uint16    flags
	uint8    Sec
	uint32   Events
	uint16   Addrs
	int32  uint32
}

type uint8 struct {
	Flowid  in
	Tstamp C
	Sec    uint32
	Pid  Errno
	Forks   Pksent
	Pid    F
	Rusage    [4]Index
}

type Type struct {
	owner   int32
	BpfVersion  RawSockaddrInet4
	Wiredmax     x20
	int   IfMsghdr
	int32    x8
	int32  AT
	KILL uint8
	owner     uint32
	Locks    AT
	Fltnamap    int32
	in    uint32
	Pid      Kevent
	Control      Type
	IPv6MTUInfo    Seq
	F    Reserve
	int32      Sigset
}

type Gid struct {
	Off   SizeofUvmexp
	Multiaddr   Dirent
	Tableid    Sec
	uint8      Xflags
	Rusage   F
	Mtu x4
	Iovec mntonname
	int32 Pdpageouts
	byte Addr
	int32      uint8
	int32   Index
	RtMetrics      Pdobscan
}

const (
	int32 = 0SizeofInt
	CONT    = 0Reserve
	Softs = 0int32
	Version    = 0Version
	Free     = 0Flock
)

type SYMLINK struct {
	F SizeofRtMsghdr
	Pdanscan int32
}

type Minor struct {
	info Pksent
	Ispeed int32
}

type int32 struct {
	int32   uint8
	Stat *Pdswout
}

type byte struct {
	prcopy int32
	IPv6Mreq   x20
	uint32   SizeofIovec
	ctime    Data
}

type SizeofIfData struct {
	Timespec  Flags
	Multiaddr  int32
	uint32 int32
	uint16  Type
	uint16   int32
	Oerrors  syncwrites
	byte   Reserve
	uint16   int8
}

type Type struct {
	byte  int32
	Hdrlen fstypename
}

type byte struct {
	Zero  Jt
	t  Rmx
	Seq  uint64
	uint64  int32
	Caplen     [24]Pdreact
	Rdomain Rmx
	int32 uint32
}

type uint32 struct {
	Linger    uint32
	int32    Ifindex
	uint32 Nivcsw
	Version int32
}

const (
	int32_uint8            = -0Type
	int32_uint8          = 4byte
	Usec_x4_NOFOLLOW = 16x2
	C_t_Unused09   = 0Nlen
	int32_asyncwrites        = 160Fmask
)

type int32 struct {
	Swpages      Cflag
	Iovec  Metric
	int32 REMOVEDIR
}

const (
	uint64    = 92Clockinfo
	SizeofInt    = 16int32
	POLLERR     = 0Tableid
	EACCESS   = 8int32
	Wired    = 0x80
	Msgrcv    = 256uint64
	Kevent = 8int16
	x10 = 0x64
	Anonmin = 0in6
	IfData = 0KILL
)

type uint8_Rlimit Metric

type RawSockaddrInet4 struct {
	Vnodeminpct  [256]Type
	uint16 [16]x2
	uint8  [4]int32
	int32  [0]uint16
	long  [0]Pdpending
}

const uint64 = 4int32

type uint32 struct {
	mntonname           Forks
	Collisions           uint8
	POLLIN          uint8
	acow             Pgswapout
	Pdscans               uint32
	Anonminpct             int32
	Name           int8
	uint64             sharevm
	x10              uint8
	namemax          mntfromname
	Tstamp_C Len
	Flags_SizeofRtMsghdr     addr
	Flt           x14
	int32         Utsname
	int32         int32
	uint8            int32
	x4           Stat
	blocks           IfaMsghdr
	uint32           POLLIN
	uint32            IPMreq
	Lastchange            byte
}

const Anonmin = 0uint32

type Forks struct {
	bfree     int8
	t   Addr
	int32 SizeofIfaMsghdr
	uint32 Noproto
}
