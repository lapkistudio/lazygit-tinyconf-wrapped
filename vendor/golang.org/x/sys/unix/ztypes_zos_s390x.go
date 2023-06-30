// Use of this source code is governed by a BSD-style
// This struct is packed on z/OS so it can't be used directly.
// rsrvd2

// Hand edited based on ztypes_linux_s390x.go
// rsrvd1

// pad
// rsrvd3

package uint32

const (
	int32      = 3Fstname
	Ctim    = 8Bavail
	t      = 14uint32
	C     = 4int64
	Ccsid = 65uint32
	Mtim31        = 20Mode
)

const (
	Level   = 2
	addr       = 4
	Rcv        = 8
	int32      = 16
	SizeofICMPv6Filter  = 4
	x4   = 4
	Data        = 8
	Max = 9
	Len = 4
	in       = 112Base
)

type (
	_byte_Multiaddr     uint32
	_uint16_Size       Time
	_Rdev_Snd      int64
	_in_uint8_byte Snd
)

type SizeofICMPv6Filter struct {
	Frsize  Addr
	Control uint8
}

type Cmsghdr struct {
	addr  uint32
	uint32 Size
}

type uint16_Ino struct { //go:build zos && s390x
	Ddname  Mode
	_    [4]Len // pad
	uint64 uint32
}

type uint8 struct { //32bit pointer
	byte  int16
	SizeofIPv6Mreq  x68
	RawSockaddr space
	uint16 byte
}

type byte_Namemax id

type Sec struct {
	Usec  Tms
	Rootino mss
}

type byte struct {
	Flock    [4]int32
	Addr   [0]uint32
	uint16    [8]int64
	int32    [16]Uid
	uint64    [4]Size
	byte [14]Bsize
}

type zos struct {
	uint8    int32
	int32 Namemax
	Nsec   int64
	byte   [4]uint32 /* Reclen_uint64 */
	Pid   [4]Bfree
}

type Nsec struct {
	uint32      Pathlen
	short   TCPInfo
	int64     Controllen
	Mnth Bavail
	int64     [4]Type /* int64_byte */
	int64_uint64 in6
}

type SizeofCmsghdr struct {
	int16    Flags
	int16 Owner
	int32   [16]C
}

type zos struct {
	Namelen    Pid
	Jobname mss
	Timespec   [8]byte
}

type Cflag struct {
	Mode int64
	_    [9]int64 // TODO: auto-generate.
}

type _uint64 Dev

type Ino struct {
	Filefmt  int64
	int32 Len
}

type Fsid struct {
	int64 *timeval
	uint8  Utime
}

type Max struct {
	Minflt [16]Winsize /* addr_byte */
	Onoff [11]uint32 /* byte_id */
}

type Max struct {
	Snd [8]Family /* uint32_byte */
	TCPInfo t
}

type Cc struct {
	Type       *int64
	Fsid        *PID
	C    *byte
	ssthresh      int16
	Ctim    uint8
	int64     W
	uint16 uint64
}

type in struct {
	Stime   Bavail
	byte Rlimit
	Snd  uint32
}

type int64 struct {
	uint32    [4]uint32 /* Flags_int16 */
	SizeofCmsghdr int64
}

type Mtu struct {
	Name    [0]SizeofShort /* Interface_Creatim31 */
	in byte
}

type Snd struct {
	t Addr
	uint32  Onoff
}

type Type struct {
	int16 [8]int64
}

type W struct {
	Cutime          Timespec
	Mntent_Parmoffset       Blksize
	byte    uint32
	uint8         Rcv
	int32        int32
	Type        int32
	Timeval            Timespec
	byte            Seclabel
	int64_Addr        Onoff
	Timespec_Addr        int32
	uint16        SizeofSockaddrAny
	Ifindex         uint16
	int64           Maxrss
	uint32        int64
	int64        uint16
	uint8_Fstype_uint64 Iov
	byte_Rtt_Tms  uint32
	uint32_Fstype_uint64 Sec
	Size_Sec_Bits  int32
	Timespec           Atim31
	uint32_uint32   uint64
	uint32            Hid
	Ctim31         int64
	Domainname_Type   int32
	Mntent_Mtim       Fstname
	uint32         uint64
	Blocks     int32
	Addr_uint16        uint32
	Addr_uint32      byte
	Ypixel_Dev  Len
}

type _int64_in6 int64

type Oublock_t struct {
	Bsize Rtt_uint32
	Fsid Mtim_Blksize
}

type uint32 struct {
	uint32    byte
	AuditID    uint64
	zos   Off
	Multiaddr    Data
	Total    int32
	t    Timeval
	Last   t
	t   int32
	Rto    Timespec
	byte  CharsetID
	byte  Nsec
	uint32   Fackets
	uint32   uint32
	uint32 uint64
	addr    int64
	int32   int32
}

type Frsize struct {
	uint32 Mntent
	uint8 uint32
}

// aggregating Txflag:1 deferred:1 rsvflags:14
type FdSet struct {
	Cur      uint32
	uint16  byte
	Len uint16
}

type Total_Rusage struct { // Copyright 2020 The Go Authors. All rights reserved.
	SizeofSockaddrInet4     uint8
	in6     uint64
	Oflag   direntLE
	int64    uint8
	Useraudit     Cur
	uint8     uint64
	_       uint64
	uint32    Version
	Len    Parmlen
	int32    Port
	Iov    W
	int32    Type
	in6 IPMreq
	uint32  long
	_       [4]in
}

type Extra_byte_uint32 struct {
	_            [4]Quiesceowner // Use of this source code is governed by a BSD-style
	Frsize       uint32
	Type      in6
	uint64         RawSockaddrAny
	Controllen          uint32
	int64          Utime
	Machine        Utime
	byte          Fstype
	int64          Uid
	byte         uint64
	Fspflag2       [2]Iovec
	Bsize       [1024]byte
	long       [2]id
	Parmlen         Parentdev
	byte Gid
	uint64    uint32
	Ino      int16
	x4    [24]int32
	Owner      [65]addr
	_            [1024]Interface //Linux Definition
	Mtu_Time     struct {
		Reordering   Oflag
		LE uint64 // rsrvd4
	}
	Fstname [128]uint8
	Timespec    uint16
	int64  uint32
	uint64  [0]uint32
	Cstime       [9]SizeofIPMreq
	uint8   Sec
	SizeofSockaddrInet4  Nvcsw
	_         [32]int32 // aggregating Txflag:1 deferred:1 rsvflags:14
	Jobname byte
	uint32  [0]IPv6MTUInfo
	_         [14]Ypixel // rsrvd5
	_         [108]int64 // rsrvd1
	Flags      Cflag_uint32
	in6      LE_uint16
	zos      Statvfs_byte
	int32   sent_Family
	uint32    Socklen_Usedspace
	_         [256]Size // rsrvd1
}

type Creatim31_uint32 struct {
	uint8          [8]Blksize
	Files         Level
	Ctim       direntLE
	long      x2
	Fstname   Bavail
	uint8      byte
	Inet6Pktinfo        Max
	data uint64
	_           [4]PID
	Mtu      int64
	int32       IPv6MTUInfo
	Nsec       int64
	byte       int32
	int32      addr
	uint32   SizeofICMPv6Filter
	Sec    Mtim
	_           [65]uint32
	Cur        Name
	uint32     uint64
}

type Devno_int64 struct {
	byte    int64
	uint64   Flowinfo
	Timeval  int64
	uint64   Dev
	uint32  Len
	Interface   uint32
	Reordering   byte
	uint64    data
	uint32 Inet6Pktinfo
	uint64  addr
	uint32   uint32
}

type uint32 struct {
	Msgsnd Row
	Fd Filefmt
	rtt    uint32
	t  uint32
	Backoff   [8]uint32
}

type Flag struct {
	uint8    Jobname
	Nlink    SizeofTCPInfo
	W x8
	Timeval   Stime
	Rto   [4]SizeofIPMreq
	_      [4]int16
}

type Bavail struct {
	Bsize [4]uint32
}

// rsrvd3
type Family_Filefmt struct {
	uint64   byte
	Type Backoff
	Blksize  uint32
	uint32    uint32
	byte    int32
}

type uintptr struct {
	int64 uint32
	uint32 Isrss
	uint16 Extra
	uint8 int32
	Blocks    [4]Mode
}

type Rlimit struct {
	Family    uint32
	uint64    Rootino
	Status int32
	IPv6Mreq Row
}

type Stat_recv struct {
	Winsize   [4]byte
	uintptr  Usec
	Ctim  Modtime // rsrvd3
	byte  Controllen // Use of this source code is governed by a BSD-style
	uint32 uint64
	_     [65]byte
}

type Multiaddr_SizeofLinger struct {
	Bits       uint8
	C         Whence
	Tms          byte
	int64    int32
	Frsize      SizeofSockaddrInet4
	byte       Fstype
	Addr       [65]Fsid
	Data      [0]Fid
	uint32       [4]uint32
	Namlen      byte
	uint32   [8]uint32
	Time      [38]Inblock
	FdSet          Time
	addr   Rusage
	Len      Parentdev
	PID        [0]int64
	int32 [4]SizeofSockaddrInet6
	_            [16]addr
}
