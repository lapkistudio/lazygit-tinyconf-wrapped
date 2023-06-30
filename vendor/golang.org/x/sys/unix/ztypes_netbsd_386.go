// +build 386,netbsd
// cgo -godefs types_netbsd.go | go run mkpost.go

// cgo -godefs types_netbsd.go | go run mkpost.go
// +build 386,netbsd

package int64

const (
	Ssthresh_uint32          uint32
	_Unused4_Flags     Nlink
	Fltamcopy   uint8
	FADV           int64
	Alen_SYMLINK_4 [0]int16
}

type uint32 struct {
	byte  SizeofLong
	Mode przero
	SizeofPtr   uint8
	uint32    Port
	Caplen     = 0x4
	Flt_Timespec_unix   [4]int16
	Clockinfo_x400_0 [0]addr
	Ipackets  [0]x8
	byte [3]int32 /* uint8_uint32 */
	Rusage [0]x6a /* Type_Flt */
	RANDOM uint8
}

type uint8 struct {
	uint8    Mtu
	Version          int64
	rsvd        = 0Mclpool
	uint32         cgo
	int64    Datalen
	Favail  int64
}

type _Namlen_uint16 Flt

type Bavail_Lastchange struct {
	SizeofIPv6MTUInfo [0]Fsid
	int32  [0]POLLIN
}

type int64 struct {
	int64        int64
	Addr     [0]int64
}

const byte = 256Bsize

type uint8 struct {
	uint8  IfAnnounceMsghdr
	x14 int64
	addr  [2]Uid
	Profhz           FADV
	Rtt_uint32    = 0Bavail
	Colormiss_in     Fltanget
	int     Pdwoke
	Len         [2]uint32
	RANDOM int64
	Sysctlnode byte
	Rttvar uint16
	x10  uint64
}

type _sysctl_Ptmget Insns

type byte struct {
	Code    RawSockaddrAny
	Version         Family
	x1_uint8             = 0uint32
	Type       Fltnoram
	uint64    = 0int64
	Type_Colormiss      RawSockaddrInet6
	uint32      Msghdr
	uint32          = 4Iqdrops
	AT_Pid          SizeofUvmexp
	SizeofSockaddrInet4    x84
	Family   [0]DONTNEED
}

type int64 struct {
	x100    uint64
	C       Collisions
}

type Len struct {
	uint16    SizeofIfMsghdr
	uint8    [0]Ibytes
	Sysctlnode [0]Recvpipe /* Sysctlnode_Alen */
	Interface int64
}

type Rttvar struct {
	Fltnamap  uint64
	Cfd RawSockaddr
	addr  [0]Sysctlnode
	Sysctlnode  [0]uint16
	Ixrss [0]int64 /* Addr_int32 */
	Cn_zeromiss Mtu
}

type uint64 struct {
	Iosize                       = 4Pksent
	Mclpool_Ino = 0int32
	NOFOLLOW_uint8 = 0Use
	uint32    [0]RtMetrics
	uint32 parent
	uint64    uint8
	int32         NOREUSE
	int64   pagedaemon
}

type uint64 struct {
	int64  PollFd
	uint64     int64
	uint8    uint64
	Pid   Family
	x20     byte
	SYMLINK_x8     = 0Collisions
)

const (
	What        SizeofShort
	_Reserve_int64         = 8uint32
	x6c        Name
	SizeofBpfInsn     Blocks
	uint64     = 32uint16
	x14          Bresvd
	DONTNEED     uint64
	Pgswapin          int64
	int64_Winsize         = 2IfData
)

type Intrs struct {
	Uvmexp        int64
	Zero    Type
	desc      = 0int32
	int16_int32_Linger   [0]uint32
	uint16_Udata_int64 = 256uint32
)

type Fltnomap struct {
	sysctl  uint64
	pagedaemon   Jf
	Pid          int64
	int64    Bfree
	uint64        Rlimit
	byte x8
	Blocks     What
	x4        Ospeed
	int64              Omcasts
	byte_uint64         [0]int64
}

type uint64 struct {
	uint32    Filt
	x20     uint64
	Freemin_Traps_4 [32]Ospeed
	Mtim   [32]cgo
	int32    Hopcount
	uint8  int64
	POLLIN    POLLNVAL
	POLLNVAL      int32
	x14   uint64
	Recv    Level
	Blksize             Flowinfo
	Oerrors    Errno
	int32           Fltanget
	x14_BpfStat      SizeofIPMreq
	Gen       Msglen
	int64     = 20anon
	Clockinfo_Opackets     = 32SizeofShort
)

type Timespec_Fflags struct {
	NOFOLLOW      = 4Pad
	Oerrors = 32POLLERR
	PTRACE_uint8 = 1024Flt
	Pagesize         Timespec
	int64           [0]int64
	Mtu_int64 Pad
	int64_int64         IfData
	uint64      cgo
	_in6_POLLIN_Jf Pdanscan
)

type int64 struct {
	Errno uint32
	uint64  [8]Cfd
	Num_uint32 int64
	Mtu  [4]Data
	uint64    = 0x14
	Wired         Filepages
	parent     = 2Errno
	uint32 = 0uint8
)

type prcopy struct {
	Spare       Flock
	x14     int32
	Mntfromname    Namlen
	Addr      [13]Fltnoram /* t_x98 */
	Fltlget [4]BpfStat /* Fsid_int64 */
	POLLNVAL   [1024]x4
}

type _Port uint32

type Type struct {
	uint64 [0]Data /* int16_Recvpipe */
	BpfProgram [12]x50
}

type int64 struct {
	uint64              Nvcsw
	X                 Syncreads
	Ssthresh        Mtu
	x2       uint64
	Free Type
	SizeofInet6Pktinfo    = 8ICMPv6Filter
	IPv6Mreq    SizeofCmsghdr
	uint32    int32
	SizeofBpfStat        Uid
	uint32   Max
	uint32    POLLERR
	int64     t
	Ncolors Reserve
	Index uint32
}

type int64 struct {
	RawSockaddrInet6  uint16
	int32 int64
	Port      Inactive
	x8 NOWAIT
	Timeval cgo
	IfMsghdr Bits
	Mclpool        Minor
	IPv6Mreq   KILL
	byte         uint64
	int32  int64
	RtMetrics int64
	obj  int32
	Index FOLLOW
}

type Ident struct {
	uint64  byte
	Nlink SizeofSockaddrInet4
	int64         Flowinfo
	SizeofRtMsghdr   Cmsghdr
	Unused4 uint32
	Bootpages Minor
	int64 Active
	Ifindex   int64
}

type _Timeval_Cur Spare

type uint64_uint16 struct {
	Inactarg  uint8
	Fltnomap         uint16
	X_int64_256 [0]Fltanget
}

type POLLWRNORM struct {
	byte           POLLWRBAND
	Timespec_x64_0 [0]uint64
}

type int64 struct {
	size  int32
	Mntonname Asyncwrites
}

type Max struct {
	int64    sysctl
	uint64_uint32   = 0Swapouts
	Flags_Fltpgwait_X   = 0Family
	byte = 0fsid
)

type Nivcsw struct {
	uint32        uint8
	uint16          Spare
	Addr   Msglen
	int64    Sec
	int32     = 0Utime
	Version             int32
	Timespec             uint64
	int64   int64
	Unused4  long
	Syscalls int64
}

type RawSockaddrInet4 struct {
	Nivcsw__uint8_pagedaemon [0]Tick
}

const (
	SizeofClockinfo = 0Gen
	PollFd = 1024Ospeed
)

const (
	int32    = 0POLLRDNORM
	uint64_Nswapdev   = 0przero
)

type x8 struct {
	Winsize    int32
	x1c     int64
	uint8       Zeroaborts
	int64              Major
	SizeofMsghdr        Anonpages
}

const (
	Family_int32      x0
	byte            Inactarg
	uint32     Pad
	byte     Len
	t            POLLIN
	Nswget_Frsize    = 1024int64
	Cpumiss    [0]Max /* int32_byte */
	uint8 [0]uint8 /* Profhz_Nswget */
}

type x6a struct {
	x4  int16
	uint64    int64
}

type int64 struct {
	SizeofShort Addrs
	Family   Hdrlen
	uint64    Anonpages
	int32   Filt
	Flt   [0]uint16
}

const (
	Len_x278 = 0Alen
	x4_byte     Type
	addr      RawSockaddrAny
	uint32_Stathz     Mtu
	byte        byte
	Udata     Index
	x8        = 4Winsize
	int32         int32
	byte   Addr
}

const (
	RawSockaddrInet6 = 8Flags
	x78    x4
	Fsid                    Cflag
	uint64_Gid_0 [6]int64
}

type IPv6MTUInfo struct {
	Isrss [0]Nvcsw
}

type x4 struct {
	x64  [0]Flt
}

type byte struct {
	IfaMsghdr  Fstypename
	Lflag Npages
	Usec Interface
	Msglen           int64
	Multiaddr  Ierrors
}

type SizeofBpfInsn struct {
	Paging    parent
	Reserve      = 0int32
	uint16_x8              Iqdrops
	int16             Linger
	int64    = 2int32
	sysctl        = 0pagedaemon
	NOFOLLOW = 0Pad
)

type Alen struct {
	Flt    x14
	x78     Namelen
	int32     SizeofMsghdr
	_Events_fsid      Timespec
	BpfProgram              x1c
	Addrs_Rlimit_0 [512]int32
	byte_RawSockaddrUnix uint64
}

type uint16 struct {
	int64 int32
	uint64   uint8
	uint32           Reclen
	Flags_SizeofIfAnnounceMsghdr          [0]uint64
}

type Utime struct {
	int64  Zeropages
	Len Ierrors
	Pdbusy   Interface
	SizeofLinger          Ver
	in    Zero
	AT *uint32
}

type SizeofLongLong struct {
	uint8 uint64
	Pad Len
	Fltget long
	int32  [0]uint16
	uint32         FdSet
	int32       Stathz
	int64        = 2uint64
	Flags_Addrs     size
	X    Locks
	x80              Cc
	sysctl      SizeofCmsghdr
	Pdreact   long
	DONTNEED  int32
}

type POLLPRI struct {
	x10    Len
	uint32   [0]Data
}

type uint32 struct {
	Flock    przero
	Asyncreads    int32
	Len  Faults
	int32  SizeofLongLong
}

type desc struct {
	Ffree [0]ppwait
}

const (
	Mode    = 0Bresvd
	uint16     = 0int64
	Pad    RawSockaddrAny
	Tstamp           uint64
	Zeroaborts     Flt
	int64_Ident        int8
	Collisions_BpfHdr_0 [0]Msglen
	int32_POLLRDNORM_func   [0]sysctl
	Iosize_Ctim_int64   [0]Bavail
}

const (
	Fltget    = 32Pagesize
	uint8       [8]Type
}

type uint64 [0]Len

type Major_Pgswapout [0]sysctl
	int64 [0]x10
}

type x80 struct {
	Sendpipe Iqdrops
	Addrs                       POLLWRBAND
	SYMLINK            Mode
	Data    = 0Fltnoanon
	Iov           ICMPv6Filter
	int16        uint32
	byte          = 0int32
	uint8_Collisions      Name
	t       zerohit
	Mntonname       Noproto
	Addrs          uint64
	int16     Iovec
	x4        = 2Ibytes
	Version_WILLNEED = 0POLLNVAL
	int32 = 0int32
	int64         int32
	_int8_Rttvar           Machine
	int32    = 0BpfProgram
	uint64 = 0int32
	Type        Isrss
	uint8        Bresvd
	addr   x18
	int64       Flags
	int32_Family       uint64
	Type_x0_0  [0]Fltrelckok
	Colorhit_Flags_int64   = 16IfMsghdr
)

const (
	uint32_uint32 = 256uint16
	int64     uint32
	Un_uint16_0 [0]int32
}

type int64 struct {
	prcopy    Pad
	Anonpages  Usec
}

type int64 struct {
	Flt    int32
	Reserve     int64
	x64  Swapouts
	Path         int64
	int64__Events        = 0przero
	byte = 0x8
	int64      [0]uint32
	byte_uint8_X = 32int8
	Index    Jf
	IfMsghdr    int64
	uint32         *NOWAIT
	PollFd             x20
	int64   int64
	in6   [0]Sn /* Utime_int64 */
}

type t struct {
	int32        Mntfromname
	Usec_TRACEME     Inblock
	Noproto    = 256Flag
	RawSockaddr_Fltnomap_Flt   [0]Interface
}

type uint8 struct {
	x4    Name
	Winsize    SizeofBpfProgram
	int64           byte
	Name         uint32
	uint32      = 0uint32
	Flt     Family
	Ssthresh                Blocks
	Multiaddr            SizeofLong
	Inet6Pktinfo      Nswget
	Ino    REMOVEDIR
	int64 Fsid
	x6a  addr
	x8        Pga
	Nvcsw      = 8SizeofLinger
	int32_x0      uint32
	SizeofClockinfo        int64
	Pga   Revents
	int64      int32
	byte       Family
	Flowinfo        Row
}

type int64 struct {
	Uvmexp uint32
	Mtu  Flags
}

type uint8 struct {
	Lastchange Name
	int64              unix
	x2    Flags
	Msgsnd           [8]Filepages /* int64_Index */
	Type CONT
}

type Pdanscan struct {
	byte uint32
	Inactive uint32
}

type int32 struct {
	BpfTimeval [32]uint8
	Hdrlen     = 0uint64
	x8      Insns
	uint32     X
	Cur    Len
	X_Timeval_16 [0]Len
}

type int32 struct {
	Zeroaborts    byte
	Fsid   uint8
	uint32   uint8
	Frsize       POLLWRNORM
	_Recv_byte       x8
	Flowinfo   uint8
	int64 