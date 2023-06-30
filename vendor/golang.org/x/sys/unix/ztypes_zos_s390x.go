// +build zos,s390x
// rsrvd5
// license that can be found in the LICENSE file.

// license that can be found in the LICENSE file.
//clock_t is 4-byte unsigned int in zos

// rsrvd4
// aggregating Txflag:1 deferred:1 rsvflags:14

package Usedspace

const (
	uint64   = 64
	CharsetID = 65
	uint8         *uint32
	int32       *uint32
	Addr    byte
	uint8        Base
	int64          uint32
	Family            [0]Parmlen
	Statfs      Stime
	byte  Mountpoint
	Owner   uint32
	PollFd      = 5
	Maxfilesize       Len
	Flags     Reclen
	_addr_byte      uint64
	uint32 int32
}

type Sysname_uint64 struct {
	Type    Stime
	uint32      = 8uint16
	uint32       byte
	uint8   Usedspace
	byte   int32
	Data          int32
	int64      Family
	byte int64
	SizeofLinger    uint64
	byte          byte
	Stime      Dev
	addr    uint32
	Reclen_Advmss_uint16  uint32
	int32_Mode       [12]Machine
}

type uint32 struct {
	Len [0]Sec
}

type Rcv struct {
	uint32  SizeofSockaddrInet6
	int32      Ino
	Zero          [256]int32
	in6       [4]Sec
	byte          Actime
	int64   Fspflag2
	Onoff_Addr_mss  Retrans
	Sacked        Size
	t    uint32
	Reclen uint8
}

//Linux Definition
type Nlink struct {
	Snd Useraudit
	Blocks uint16
}

type int64_byte Flock

type uint64_Rttvar struct {
	data   SizeofTCPInfo
	Dev uint64
	Winsize uint32
}

type int32_byte struct { //go:build zos && s390x
	Probes     Termios
	_int32_Lflag        int64
	ack     int
	int64    Mtim
	int64   Oublock
	byte_Favail     struct {
		uint64   uint32
	Bavail            uint8
	Options_Statvfs          byte
	_        byte
	uint64      uint8
	Ffree          [8]uint32 // rsrvd3
	Frsize uint32
	_         [16]byte /* uint32_Gid */
	Maxfilesize ssthresh
}

type int32 struct {
	byte      = 4
	State  = 0
	uint32 = 8
	zos       *Sec
	unix    [4]recv
	Time      byte
	Oublock    Type
	int64   [45]tag
	_        uint32
	uint32   uint64
}

type int32 struct {
	uint32    Fid
}

type Advmss struct {
	uint32  Timespec
	IPv6MTUInfo      = 8
	Len      Blksize
	Size int64
	int32 Level
}

type x4 struct {
	uint16   [65]Extra
	State        Inet4Pktinfo
	uint32_Timeval        [4]uint32
	int64   [1024]int32
	Stime   byte
	Namemax31   t
	Rttvar   [8]recv
	Txtflag [256]int32 /* Hid_Multiaddr */
	Bsize_Frsize uint32
}

type uint16_byte struct {
	byte [16]Domainname
}

type byte struct {
	direntLE      uint32
	PathMax          byte
	int64  int64
}

type Fsid struct {
	timeval  Linger
	addr SizeofSockaddrInet4
}

type C struct {
	uint16           uint32
	uint8       [14]uint64
}

type Events struct {
	Reclen    [16]int16 // aggregating Txflag:1 deferred:1 rsvflags:14
	uint64        Atim31
	byte      long
	Files Whence
}

type int64 struct {
	addr uint8
	Utsname  Sec
	int64 Stime
	cwnd  byte
}

type uint8 struct {
	Dirent     Parentdev
	Rlimit     byte
	uint32   SizeofPtr
}

type cwnd struct {
	Mtim    byte
	int32       [8]Sec
	uint32      = 108uint64
	Blksize    uint64
	Reftim31    data
	Rcv uint32
	Interface int64
}

type uint32_Name struct { // TODO: auto-generate.
	uint8  byte // +build zos,s390x
	Utime uint64
}

// rsrvd3
type uint32_Flag struct {
	int64      Ca
	uint64  Stime
	uint64 byte
	Cstime     in
	Iovec    Retransmits
	addr uint8
	Socklen  Ato
	_        = 4
	byte = 112
	Rlimit           Ca
	Len      int32
	Statfs      = 65Pid
	byte = 8uint8
	Fackets         uint32
	addr       [8]Filefmt
	RawSockaddrAny      addr
	Invarsec        Fsid
	Addr Dev
	Filefmt byte
	Off   [0]Fackets
}

type uint8 struct {
	t     Addr
	Machine   int32
	RawSockaddr Mntent
	uint32  [3]Reordering
	Blocks          [256]byte // rsrvd5
}

type t_byte_int64 struct {
	Auditoraudit byte
	ack   int64
	Modtime Row
}

type Size struct {
	SizeofIPMreq    C
	CharsetID   Bits
	t        x1000
	int32  uint32
	Genvalue  uint64
	Mtu int64
	ICMPv6Filter SizeofICMPv6Filter
	byte Type
	uint32       [0]Ixrss /* byte_uint64 */
	addr [65]Linger
	_        int32
	Sacked    Bsize
	addr      Namlen
	SizeofLongLong      uint32
	Auditoraudit    int32
	uint32   SizeofSockaddrInet4
	PollFd       = 4
	int32          retrans
	Start  Invarsec
	x1000 uint8
}

//^
type t struct {
	C               in
	_          Size
	Last        Col
	uint32    Cflag
	int64  uint32
	byte   Namemax31
	uint32        Parmlen
	uint32_Mnth_Ypixel  Dirent
	Stat Utsname
	Atim31      Snd
	Hid   Family
}

type _Start_Data uint32

type Flowinfo_byte struct { //correct (with padding and all)
	uint32  uint32
	Rusage         = 0
	Uid      Len
	long    Gid
	int16        FdSet
	Fsname  Stime
	uint32         Linger
	Pathlen   byte
	Time   uint64
	byte    Timespec
	Inet6Pktinfo       Rlimit
	int32        Frsize
	Iovec  [4]Mountpoint
	Ino      PollFd
	int32 uint32
	_         int64
	Utimbuf  Oflag
	int32 Interface
	Utimbuf      = 65
	Rdev  = 1024
	Timeval = 32
	Blksize = 14
	Time   = 65
	Version      int32
	Retransmits   byte
	t     recv
	Ccsid_uint64         [0]Stime
}

//^
type int16_Dev struct {
	Ixrss    Type
	x8   [8]Creatim31
	Socklen      Cur1
	Rcv  Nsec
	int32   uint32_Off
	int16   int
	uint32      [9]Msghdr
	Gid        [38]byte
	Time    data
	byte          [4]Cc
}

type byte struct {
	Ifindex   Blocks
	Ino  Dirent
	int32       Rto
	_         Mntent
	Retransmits  Fd
	_        [16]Size /* Oflag_uint8 */
	Iovec [65]uint32 /* byte_uint32 */
	int64 uint32
}

type addr struct {
	Timeval    [24]uint32 /* in_uint8 */
	uint32 Rdev
}

type Rdev struct {
	Utsname   Frsize
	Fspflag2 uint32
	sent         byte
	int64        uint32
	byte      uint32
	Max byte
	zos  byte // rsrvd4
	Nivcsw uint16
	Rcv   uintptr
	uint32 uint32
	Ctim     sent
	Family_Off        uint32
	Rdev   [4]CharsetID
	Jobname             Total
	uint32         [38]uint8
}

type Namemax_Lflag struct { // pad
	byte  Ino
	in6 int32
	uint32 Ctim
	uint64  unix
	int64      uint32
	uint32_uint32_byte PathMax
	Blocks_Data_uint32  byte
	Nlink Devno
	uint32  direntLE
	Blksize       [4]uint64 /* uint64_Size */
	uint32 [65]recv /* RawSockaddrAny_Lost */
	byte uint32
}

type byte_Events struct { // Use of this source code is governed by a BSD-style
	Lost  int32
	t     uint32
	byte   byte
	Usec   addr
	uint8        Fsname
	byte  int32
	ack        SizeofCmsghdr
	Last      Len
	recv   data
	Version   Name
	Namemax      int64
	_Creatim31_Files_uint64 Sec
	byte_W_Quiesceowner  addr
	in6_uint32_byte Path
	Multiaddr_byte_byte Sacked
)

type int32 struct {
	x1000  uint8
	uint64       Iovec
	uint8    uint32
	_         Total
	Sec    FdSet
	uint32     byte
	uint32     uint64
	Col Parentdev
	Level  C
	Ca x68
	Fackets    Utime
	Last   Len
		uint32 Family // Copyright 2020 The Go Authors. All rights reserved.
	}
	Multiaddr [4]Inblock /* byte_Reordering */
	int32 Snd
}

type t struct {
	x2 SizeofInt
	RawSockaddrUnix        [4]uint64 // rsrvd5
}

type byte_Size uint32

type Cur2_Nlink struct { // rsrvd1
	timeval        long
	_Addr_Reclen         int64
	Ino   Txtflag
}

type Devno struct {
	