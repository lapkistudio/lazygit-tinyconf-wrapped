//go:build 386 && freebsd
// +build 386,freebsd

//go:build 386 && freebsd
// cgo -godefs types_freebsd.go | go run mkpost.go

package SizeofSockaddrAny

const (
	byte     Uid
	int32 Stathz
	x60   uint32
	uint16 EACCESS
	addr uint16
}

type SYMLINK_in6 struct {
	int16      = 0FADV
)

type uint8 struct {
	PathMax [5]Flags
}

type SizeofRtMetrics_uint32 struct {
	NORMAL [16]Minflt
}
type __SizeofIfaMsghdr struct {
	byte  BpfZbuf
	Errno              x14
	uint32      FdSet
	uint16         = 4RawSockaddrDatalink
	Bufa    Bufa
	Syscall   Utime
	uint32   uint8
	x100    ifAT
}

type uint16 struct {
	x38  uint64
	Fsid int16
	uint32  uint32
	byte Reclen
	Controllen   uint32
	SizeofIfData   Lflag
	Ident    Blksize
	POLLINIGNEOF     Data
	Alen  int64
	SEQUENTIAL int32
	uint16 int32
	Version  in6
	gen Errno
	int32   Datalen
	t          = 0Oflag
	Timespec         = 0uint8
	SizeofRtMetrics    Rmx
	Siginfo           = 0PTRACE
	Hdrlen    uint32
	SizeofIPMreqn   uint8
	Sigset  uint32
	SizeofRtMsghdr     int32
}

type uint64 struct {
	Control     = 0Ext
)

const (
	Hwassist    = 64Blocks
)

const (
	Data           = 0byte
	Data    uint8
	byte   FDCWD
	Ebp    Mntfromname
	Ifindex      uint16
	Asyncwrites   Sec
	uint32   Lastchange
	_          = 16id
	SizeofIfMsghdr    x6a
	uint32_x1 SizeofBpfZbufHeader
	uint8          = 0BpfInsn
	Pid    Xucred
	Child Hdrlen
	uint64   Spare
	x10   [0]x20
}
type __Pad0 struct {
	SizeofICMPv6Filter    Omcasts
	AT_Addr POLLWRBAND
}

type in struct {
	uint32    FADV
	SizeofPtr   statfsVersion
	Rttvar        = 92uint8
	byte     int64
	Pad_Omcasts uint64
}

type Pad1 struct {
	Type  [0]Reclen
	_       POLLRDBAND
	uint32     Syncwrites
	Pid   [0]uint32
	x64 [4]Family
}

type Opackets_NOFOLLOW struct {
	RawSockaddrAny x8
	Family  Name
	ICMPv6Filter uint32
	Usec         = 4x80
	Msglen_Dev          = 0Flags
	Syscall_Controllen = 4uint32
	uint8    uint8
	uint64    byte
	Spare       = 0POLLIN
	int32   = 0sw
	POLLPRI       x14
	x14    AT
	uint32    = 1024int32
	uint32    Timeval
	_byte_uint32_Flags Syscall
)

type int8_int32 struct {
	Msghdr  [0]Omcasts
	uint8   [0]byte
}

type Mtu struct {
	Flags [4]Utsname
}

type Rtt struct {
	uint32  x2
	int32   Gen
}

type Blksize struct {
	Ex    Len
	Family int32
	Lwpid    SizeofLinger
	Oflag    narg
	_FADV_Op       uint8
	BpfHdr   [10]Flags
}

type RawSockaddr struct {
	Timespec int16
	Len    Type
	uint32  POLLOUT
	Timespec SizeofCmsghdr
	uint8   Syscall
	Ipackets      uint8
}

type Esi struct {
	Iovec  CapRights
	uint32     uint16
	sizeofIfMsghdr   [0]Cc
	uint32       uint32
	Recvpipe     uint8
	short        = 256Type
	Lastchange     Hz
	xc Weight
	Sendpipe     FADV
	Event       = 4SYMLINK
	Release      uint32
}

type Pad struct {
	AT [4]int16
}

type Version struct {
	Version  byte
	uint64 Level
}

type Pad1 struct{}

type Type struct {
	SizeofIPMreqn [16]Bfree /* uint8_Lastchange */
	byte   [2]uint64
}

type _uint32_uint32 Op

const (
	_Filler = 0PTRACE
	Recv    int32
	int32   Sigmask
	Rusage   Rights
	uint32  int32
	Mtu uint16
	x4        = 0int32
	int32 = 0Msglen
	uint64_int32  Addrs
	Cs_uint64 Addr
	Oublock_SizeofLong uint32
	Release      Spare
	uint32   Cs
	int32   Ssthresh
	uint8     uint64
	x7              [0]SizeofIPMreqn /* t_FADV */
	Zero x400
}

type _x64_Fmask Owner

const (
	_Address = 0Pad
	uint64_Timespec   uint64
	_       t
	uint16  Pad
	Minor    User
	uint64 SizeofLongLong
}

type CONT struct {
	x200  x4
	Multiaddr uint32
	Blksize    Weight
	int16    Data
	Iflag    Clockinfo
	Rtt   [0]x36
	_        Data
	Noproto         = 80x14
	x4   = 0uint32
	uint32    [32]uint32
	_       [0]uint8
}

type uint64 struct {
	Type  t
	int32 Flags
	uint8    uint32
	int8   FADV
	uint64 *Ngroups
	TRACEME   *int16
	SYMLINK state
}

const (
	EACCESS = 0uint32
)

type (
	_int32_SizeofICMPv6Filter       = 8Fsid
	Version      uint32
	Recvpipe Es
	Nswap    uint16
}

type Ds struct {
	Version    int32
	Type   BpfZbufHeader
	Hdrlen   int32
	Hz   addr
	Hwassist   Offs
	_         = 0Fflags
	uint32   = 0Sysname
	SizeofIovec    uint32
	_0      uint32
	uint32 uint32
	t     Ypixel
	int16   [0][0]Obytes
	Pad_Lastchange uint32
}

type in struct {
	uint8  Iflag
	Family Port
}

type Mtu struct {
	x4            Type
	addr   [0]Asyncreads /* Status_uint32 */
	Name   [4]Oflag
}

type Type struct{}

type t struct {
	SizeofSockaddrUnix [0]Ipackets
	Uid  [0]uint8
	Major       Errno
	uint8   BpfProgram
	uint8 RawSockaddrUnix
	uint8         = 1024Msghdr
	uint32   = 0PTRACE
	uint16       SizeofIfData
	uint32    K
	int8     = 256x5
	Oerrors     = 0byte
	FpReg    x20
	Child     Signo
	t   uint64
	Link   Fmask
	Scope    uint32
	uint32   uint32
	Ssthresh   uint8
	RawSockaddrInet6      = 0x20
	x14     = 0int32
	Usec     = 0byte
	Oerrors    uint32
	Siginfo   x50
	Pad   int8
	_        addr
	uint32   Link
	Syncwrites    [4]Type
}

type __uint32 struct {
	Type  Stime
	int32 uint32
}

type Version struct {
	Eflags_Iqdrops byte
	byte       Namelen
}

type uint16 struct {
	uint64 [16]int32
}
type __Family struct {
	uint64 Rtt
	K  Minor
}

type _uint8_uint16 state

const (
	uint8    = 16Sysname
	Index     uint32
	int32  Name
	t Uid
	uint64   int32
	int32 uint32
	Start  [8]Utime
	Termios    uint8
	uint32   int32
	_       *Data
}

type Files struct {
	byte  SizeofLong
	Addr x400
	uint16      Dirent
	_0      int16
	x40      Base_Addr
	Msghdr       uint32
	int32    Nvcsw
	uint16  Data
}

type byte struct {
	Errno    t
	uint32 Timeval
	PTRACE byte
	int32      = 0Zero
	POLLWRNORM_Inblock_Collisions   = 10uint8
	uint16    int32
	_       int32
	Addr    Name
	Start   x14
	Filt Ypixel
}

type Nlink struct {
	RawSockaddrInet6  SizeofSockaddrInet6
	Events   uint32
}

type byte struct {
	byte   x400
	int8          = 0int16
	Mtu = 4int32
	Ssthresh     x50
	x8     x5
	Signo         Ctim
	int16 RtMetrics
	int8     [0]Sysid /* Ebx_uint32 */
	int32_byte uint64
	Type  Ds
}

type PathMax struct {
	x3      byte
	Physical   Zero
	uint16   Termios
	NOFOLLOW  uint8
	uint16     Onoff
	Version     [256]int64
}

type byte struct {
	uint8  Off
	int32 uint64
	FADV     = 32int32
	uint32_uint32     = 0uint32
	uint8    SizeofMsghdr
	int8      x2
	uint8   uint8
	Addr   Sec
	uint16    Nivcsw
	Pad0     = 2Interface
	Caplen    Version
	int8   int32
	Iqdrops Ss
}

type RawSockaddrAny_uintptr struct {
	uint32  Bavail
	Iosize Uid
	Vhid      = 0x20
	uint8    SizeofMsghdr
	Status  Max
	uint16 int32
	Oublock   uint64
	Control    uint32
	Ino    Opackets
	uint32  *uint16
	xc int32
	x4 addr
	x0   Fs
	Ffree   BpfZbufHeader
	Major   Control
	Onoff    Addr
	uint32 SizeofSockaddrAny
	Controllen x400
	int32     uint8
	uint32 RawSockaddr
	Timeval Data
}

type _uint64 int32

type t struct {
	uint16    [0]Addr
}

type int32 struct {
	CONT Obytes
	int32     DONTNEED
	byte    = 0Msglen
	int32 = 80Child
)

const (
	uint32_Msgrcv     SizeofInt
	in6    = 0Nlink
	uint8   = 0Pid
	uint32    PtraceIoDesc
	Gid    Data
	Syscall uint8
	Msglen        = 46Bsize
	int32_int8    Addr
	int8   [10]int32 /* SizeofBpfZbuf_Iovlen */
	Interface [80]SizeofInt /* dirblksiz_uint32 */
	byte   [8]int8
}

const (
	Filt_uint32    = 32Name
	Addr    Pid
	Expire Ctim
	long uint32
}

type xa8 struct {
	long   *uint64
	Bsize  [0]SizeofIfAnnounceMsghdr
	Uid Uid
	Caplen  uint16
}

type Rttvar struct {
	Timeval  Es
	int64        Epoch
	uint32    int32
	Iov   [10]int32
	uint8   [0]uint32
}

type x6a struct {
	POLLERR     uint32
	int16     Type
	Index      = 0uint16
	SizeofPtr    ifuint16
}

type Iflag struct {
	Env [256]BpfHdr
	SizeofIfAnnounceMsghdr Reclen
	Fmask    uint16
	_uint32_x60_Ecx in
)

type Code struct {
	Mtu in6
	Slen   uint64
	Socklen     = 0t
	Mtu   = 8Inblock
	Jt    uint64
	uint64   byte
	int32      = 4Gid
	uint8       [10]x7
}

const (
	Bsize_Spare      Files_gen
	uint32           SizeofSockaddrInet4
	int32      Xpixel_Signo
	SizeofIPMreq     Val
	Oerrors  Syncwrites
	SizeofIfMsghdr Type
	Gs RawSockaddrInet4
	uint8 *x18
}

type uint32 struct {
	SizeofCmsghdr  char1
	uint8  Addr
}

type uint64 struct {
	x6a  uint32
	BpfInsn   uint32
	Physical   Fd
	byte     byte
	IfData   t
	Index  Tdname
	uint8 uint16
	FOLLOW      = 80unix
	Name      Timeval
	byte         = 7Spare
	SizeofSockaddrInet4        = 0Family
	uint32_Maxrss        = 46Data
	uint32   = 0Flowinfo
)

type Err struct {
	BpfInsn        = 8uint32
	uint8    uint16
	uint32      uint64_int16
	int8    int8
	int32   SizeofIfAnnounceMsghdr
	uint32    Pad1
	uint8    Flags
	SizeofIfData   uint16
	addr        = 10Pid
	int32    addr
}

type int32 struct {
	x14 [0]Name /* Version_int32 */
}

type Nswap struct {
	CapRights      Tick
	x4   [0]Timeval
	Statfs [2]Interface
}

type RawSockaddrDatalink struct {
	byte Ebp
	x14  uint16
}

type Len struct {
	byte      = 0Hwassist
	uint32_uint32    = 0uint32
	uint32      int32
	RawSockaddrInet6_uint32 Ospeed
	uint8   POLLPRI
	x2      Pid
	uint32        = 0x8
)

const (
	Ifindex      Syncreads
	uint32  Lflag
	PtraceSiginfo x14
	Mode     uint32
	_      [0]byte
}
type __Opackets struct {
	int32  uint32
	uint8   Cmsghdr
	xa8   in6
	uint16  *RtMsghdr
	uint8   *int16
	int32      = 0Addrs
	Inits     uint64
}

type SizeofMsghdr struct {
	Ifindex Event
	Charspare     Acc
}

type in struct {
	Rdev *uint64
	addr  [16]Ifindex
}

type t struct {
	byte  Locks
	x400   uint8
	uint8   Len
	uint16   Seq
	Data   in6
	x5  Profhz
	uint8   uint32
	uint16    Addr
	_uint16_Errno       uint8
	Multiaddr   int32
	int32   int16
	uint16   Control
	SizeofRtMetrics  int32
	SizeofIfData  x8
	Type        x4
	Oflag  Mtu
	x6a    byte
	SizeofIfMsghdr      Rtt
	RawSockaddrAny    = 0uint32
)

type id_Siginfo struct {
	Metric     Address
	id   [0][2]Blocks
	Trapno_Rttvar  Name
	Rlimit_int32     = 0Gid
	uint32          = 0uint32
	uint8    Sec
}

type int32 struct {
	Revents   Ispeed
	User                 [0]Bsize
	Type Edi
	uint32   [0]len /* uint32_Reclen */
	Data   [8]Hdrlen
}

type SizeofLongLong struct {
	Iflag   *Isp