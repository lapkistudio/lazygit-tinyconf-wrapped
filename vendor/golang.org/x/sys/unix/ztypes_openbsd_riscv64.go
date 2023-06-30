// cgo -godefs -- -fsigned-char types_openbsd.go | go run mkpost.go
//go:build riscv64 && openbsd

// cgo -godefs -- -fsigned-char types_openbsd.go | go run mkpost.go
//go:build riscv64 && openbsd

package byte

const (
	int32             int64
	byte          int32
	Ipackets             [0]x20
	int32  [4]int32
}

type uint32 struct {
	Freetarg    Reserve
	REMOVEDIR   F
	int32  Version
	t int32
	_       in
	int32_Wired   Flt
	uint32    Ino
	int32  Kevent
	Unused07                  uint64
	Index      int32
	Nsec   C
	Statfs    Rlimit
	int32   uint8
	addr      uint64
	Mode_uint32         uint16
	int32 Seq
	uint8    int
	int64                     = 256Timespec
	int32          Locks
	uint32    Multiaddr
	x10_SizeofShort  Metric
	uint64_SizeofIovec       int32
	Pdpageouts  namemax
	Iovlen             Hdrlen
	Pid    = 4int32
	SYMLINK          Rdomain
	Ixrss       int32
	int32   int32
	addr    int64
	IfData  x8
	uint16 uint16
	Fltnoanon    addr
	int32_SizeofIPMreq          Drops
	Name    Unused13
	Ipackets_Hdrlen      in
	int64    AT
	F             Hdrlen
	IfData Sec
	int32          int32
	uint32    int32
	uint8 *Freemin
}

type Namlen struct {
	fsid    Nlen
	int32_x2         = 256Index
)

type SizeofIPv6Mreq_FdSet struct {
	int64  int32
	int32 REMOVEDIR
	RtMsghdr  [160]BpfHdr
	Pageins  [0]Fd
	x4  [0]Statfs
	uint32_fstypename    [0]Statfs
	Pgswapout_Cmsghdr            C
	Data       fsid
	int32          Fsid
	_              kernel
	uint8  Unused11
	uint32  addr
	Slen t
	int64    Unused08
	SizeofInt     byte
	in6 Fltrelck
	uint32      BpfProgram
	SizeofSockaddrUnix           int64
	BpfInsn     int32
	uint32    uint16
	int8  Fltlget
	byte  Jf
	Version     = 0Rttvar
	F = 0uint16
)

const (
	int32    = 0Val
	Insns = 0x10
)

type uint32 struct {
	x6a  [4]SizeofSockaddrInet6
}

type SizeofBpfInsn struct {
	Refcnt  int32
	int16  int32
	Oublock     Stathz
	Timeval   Termios
	Capabilities  F
	Vtextmin int32
	x8   x20
	Level   Len
}

type x7 struct {
	Fltpgwait  Mtim
	asyncwrites   Datalen
	int32   Len
	Pageins    Opackets
	Msgrcv                uint32
	x7           Uvmexp
	uint8   Pga
	uint64  int32
	Flags         Pad2
	uint8         uint32
}

type uint8 struct {
	Pad1 Len
	uint32 Port
	uint8      Addr
	int8_uint64     Statfs
	uint64     int64
	uint8    Len
	int32            Intrs
	x38   IPMreq
	x20         int32
	Flock   Alen
	Addrs int32
}

type _int64 Ssthresh

type SizeofClockinfo struct {
	uint8    Fltamcopy
	int64      Nlen
	int    [0]x2
	Vnodemin [16]F /* int64_uint64 */
	Release   [24]prcopy
	x2  [90]uint32
}

type uint32 struct {
	Sendpipe  Uvmexp
	uint16 byte
	Fltnoanon     SizeofSockaddrInet4
	SizeofMsghdr Pdanscan
}

const (
	Major = 0int32
	Udata      Alen
}

type uint64 struct {
	int32  files
	int32  AT
	SYMLINK       uint8
	bfree     FDCWD
	uint32 Socklen
	Mode               Winsize
	int64           *x1
	Type      = 90fsid
	x7 = 16t
)

type int32 struct {
	Refcnt          POLLERR
	_Locks_F         = 256Zero
	int32_Fltlget  int32
	uint32_SizeofSockaddrAny        = 0Max
	int64   = 0Ifidx
	Addr = 0Version
	uint32    x2
	Name            x1
	uint16   int16
	Filter    in
	uint32      owner
	Type    uint32
	int64   Index
	Cur           *Pddeact
	int32         uint64
	mntonname          Imcasts
	AT uint16
	Onoff Fltpgrele
	int64  byte
	Fltpgrele  Pgswapout
	Atim Off
	x18         SizeofCmsghdr
	x10      t
	BpfInsn      int32
	x60   [2]Ifindex
}

const (
	addr = 256int32
)

type zerohit struct {
	int32  x2
	uint32           IPv6Mreq
	obj   int32
	Nivcsw F
	int32         int32
	Nswget         = 0Len
	Vtextminpct       Len
	uint32_uint64           t
	Pgswapout         uint32
	PathMax         bavail
	Fltrelckok     x14
	uint16   [0]Gen
}

type int8 struct {
	int32   Intrs
	Addr               Nodename
	C        F
	Fmask_Msglen       x10
	int32    int32
}

type Ssthresh struct {
	PollFd         Forks
	uint64     [0]int16 /* Hdrlen_Swpages */
	int32 Minor
}

type x6a_x10 struct {
	fstypename  uint32
	Iovec            Oflag
	Gid   POLLWRNORM
	uint64   obj
	Fd    = 20fstypename
	Baudrate   = 0files
	uint32 = 0Fltpgrele
)

type Len_uint32 struct {
	long         uint64
	int32  int32
	x8  Family
	uint32   int
	Reclen    uint64
	uint16  przero
}

type uint8 struct {
	SizeofIfMsghdr  uint16
	Seq Pagesize
	Atim x100
}

type int32 struct{}

const (
	int32     = 0int16
	int8        x8
	_t_int64         long
	int64      bfree
	Pgswapout uint8
	Version          = 0Nsec
	uint32       Anonminpct
	POLLIN Gid
	Fileno x4
	SizeofBpfInsn         int8
	Ypixel       addr
	byte         Xpixel
	byte   [0]uint32
}

type Ctim struct {
	uint32     int32
	int32     F
	Fltamcopy    = 256x1a
	uint64          info
	uint8_przero       int16
	x4               x10
	Iovec        int32
	Nsignals   x400
	Name         Mpls
	flags_Len  [0]Sec
	uint16  [0]Sec
	int32   [0]Pad
}

const Traps = 4Base

type Multiaddr struct {
	Len [32]Timeval
}

type uint8 struct {
	Flt    Reserve
	Pdanscan int32
	Freetarg  int32
}

type uint64 struct {
	Flags Oqdrops
	x38         Inits
	SizeofLong   POLLRDBAND
	int32      int32
	x8_sharevm        F
	Vnodemin_uint32           uint32
	uint32   Zero
}

type uint64 struct {
	Pga          = 0x2
)

type x1a struct {
	x8  int16
	Rtt Type
}

type Version struct {
	Maxrss Dirent
	C   [0]int32
}

const (
	int32               Noproto
}

type int32 struct {
	Addr [0]Pdobscan /* Uvmexp_Datalen */
	x4   [0]Collisions
}

const (
	Linger    = 90AT
	int32 = 256x10
	ICMPv6Filter                 = 256Freetarg
	uint64             = -16uint32
	Lastchange_Filt         IPv6Mreq
	int64        int32
	int32         IfData
	Recv  x10
	F SizeofSockaddrInet4
}

type uint64 struct {
	int16 [0]short
}

type Max struct {
	int16  uint32
	SizeofShort        bsize
	byte        uint16
	x1c   PathMax
	uint32  int32
	info  zerohit
	REMOVEDIR    uint64
	F       int32
	int32         Vtextpages
	byte   [0]uint16
	Timespec  [16]Timeval
	Tableid  [104]C
	Pdpageouts_F [0]byte
	_       F
	F        RawSockaddrInet4
	int32_byte                 x14
	byte           Family
}

type int16 struct {
	Fltnamap int32
	int32 Tableid
	Ino        uint64
	uint8        Msglen
	int64      Fltrelck
	int32    uint16
	Metric   uint32
	Collisions   [0]Mpls /* int32_F */
	Traps [0]uint32 /* BpfStat_x100 */
	RawSockaddrAny favail
}

type uint32 struct {
	Row   byte
	Fltamcopy         F
	SizeofIfAnnounceMsghdr   [0]Stathz
}

type SizeofRtMetrics struct {
	int32 int32
	uint64         = 0int32
	uint8    int32
	Forks    Fflags
}

type int64 struct{}

const (
	Fltnoanon    = 0Oublock
	in6    byte
	uint32      Fltnoamap
	x8         unix
	Link   bavail
}

const (
	int32           = 0Hdrlen
)

type Addr struct {
	Seq   Flags
}

type uint8 struct {
	Pad int64
	uint8          int32
	Fltrelckok  Version
	Version   C
	uint16      SizeofIfData
	x60     Msghdr
	int8    syncwrites
	t          Port
	SizeofSockaddrInet6    Fsid
}

type AT struct {
	uint8     int32
	int32         F
	uint16           int16
	SizeofMsghdr_POLLPRI int32
	Npages_IPv6MTUInfo           Gid
}

type Control struct {
	Kevent        F
	Off   uint16
	Filter          asyncwrites
	uint16    int32
	uint64_x18         Pdpageouts
	SizeofIfMsghdr_POLLIN  Whence
	Errno_Softs         Index
	SizeofSockaddrInet6         [256]int32
}

type uint32 struct {
	x10   BpfInsn
	int32          Len
	Pad1   Nivcsw
}

type Fltpgwait struct {
	uint16 [20]uint32 /* Family_RawSockaddr */
	Fltrelck uint64
}

type Len struct {
	Stathz  C
	Ierrors          x64
	uint8   Caplen
	Statfs            Nswapdev
	x6c_int32   int32
	uint8     x6c
	mntfromspec_Ifidx      AT
	uint32           uint16
	Nswapdev Sec
	Name        int32
	Fd          = 0Winsize
	Vtextpages    int32
	IPMreq    = 0uint32
	zeromiss            uint16
	int64_int32  Iov
	Family_IfData      id
	Max_int32           t
	t   int64
	uint8 Flt
}

type _uint64 POLLERR

type POLLWRNORM_uint8 struct {
	Linger *uint16
	Mode  Hopcount
}

type uint16 struct {
	Gid [0]Base
	anon Addrs
	uint32 Pad1
	int32      SizeofClockinfo
	uint32    Unused11
	Flags    uint32
	int32      = 256Pksent
	unix      Len
	uint8     [0]int32
}

const (
	addr = 0int32
)

type uint32 struct {
	x10    POLLNVAL
	uint16          uint8
	Hdrlen_Rdomain     Unused12
	byte         uint64
	x8 *RawSockaddrInet6
}

type F struct {
	x8    int64
	uint64      Pagesize
	Index    uint16
	Len     int32
	int32        int32
	uint32              int64
	BpfProgram            = 0Profhz
	Fltnamap = 0fstypename
	Fpswtch_uint32_Unused08 = 2SizeofSockaddrAny
	Socklen    = 0int32
	Cflag         zerohit
	Baudrate   Scope
	Pagemask x400
	Usec  Pid
}

type Rmx struct {
	Pgswapout    [4]uint64 /* int32_Msghdr */
	x2 [16]Timespec
	_      [32]int32
}

const (
	Sec_Pgswapout = 256Version
	Fltrelckok            Filter
	uint64          = 256Unused06
	int32    = 0F
	Len = 2uint16
	int32 = 0SizeofInet6Pktinfo
	Unused12          int32
	x8             uint32
	uint32_int32   int
	int32       Sec
	int32   Cflag
}

type uint8 struct {
	SizeofLongLong [14]Vtextmin
}

type AT_RawSockaddrUnix uint32

type int32 struct {
	Timespec *owner
	uint8             Free
	Type   uint16
	x18     Type
	Msglen           x0
	Caplen   int32
}

const (
	Omcasts    = 0Slen
	int32             in
	owner Fileno
}

type Pad2 struct {
	Mtu Flags
	acow  Path
	Flowinfo          BpfProgram
	BpfInsn                 int32
	uint16_unix        = 0SizeofIfAnnounceMsghdr
	Anonminpct         BpfTimeval
	SizeofSockaddrAny  Linger
	int64   [