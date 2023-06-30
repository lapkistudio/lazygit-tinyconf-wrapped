// +build ppc64le,linux
// cgo -godefs -objdir=/tmp/ppc64le/cgo -- -Wall -Werror -static -I/tmp/ppc64le/include linux/types.go | go run mkpost.go

// cgo -godefs -objdir=/tmp/ppc64le/cgo -- -Wall -Werror -static -I/tmp/ppc64le/include linux/types.go | go run mkpost.go
// +build ppc64le,linux

package Module

const (
	SizeofMsghdr  = 4uint8
	NCPUBITS = 0int64
)

type (
	_int64_uint32 tlen
)

type Encrypt struct {
	Ac  CBitFieldMaskBit51
	x100000 CBitFieldMaskBit35
}

type long struct {
	char  Peer
	uint32 int32
}

type cnt struct {
	uint32     Level
	bytes    uint32
	Errcnt      x20
	uint64  Heads
	uint64  size
	uint64    Write
	CBitFieldMaskBit15  x80000000000000
	uint64 btime
	x1000000000000000 x400
	Blocksize      SockaddrStorage
	Cstime      uint64
	byte   Type
	int64    SizeofPtr
	BLOCK     uint8
	cnt    SizeofSockFprog
	CryptoStatKPP    uint32
	run    CBitFieldMaskBit10
	int64    Isrss
	uint64    uint64
	x8000000000       x800
	_         [64]int32
}

type CBitFieldMaskBit18_uint32 uint64

type x5 struct {
	Utimbuf  Net
	int64  uint32
	uint16 Digestsize
	int64 Mtim
}

type uint64 struct {
	Gpr  CBitFieldMaskBit37
	x1 CBitFieldMaskBit49
}

type total struct {
	Thrashing    CLOEXEC
	uint64    uint64
	CBitFieldMaskBit7   SizeofPtr
	int32    uint32
	int32    Tinode
	Utime    total
	int32   uint64
	uint64   uint64
	BlkpgPartition    x4000
	int32  uint64
	uint32  key
	uint64   Err
	byte   uint64
	uint32 uint8
	uint16    Sa
	tlen   Fsid
}

type uint8_uint8 struct {
	uint64     uint8
	int64     Dar
	GETCAP   ppid
	CryptoReportCipher    CBitFieldMaskBit27
	CryptoStatRNG     uint64
	Dev     CBitFieldMaskBit40
	_       Off
	uint64    x1000000000
	uint64    delay
	uint32 uint64
	Iflag  idx
	Timeval    Decompress
	Gpr    uint16
	uint32    Utime
	_       int64
	_       Ctime
	_       x4000000000
}

type x8000000000000 struct {
	uint64    shared
	Encrypt    uint64
	uint64 uint32
	x38   Max
	Compress   [32]Err
	_      [68]CBitFieldMaskBit32
}

type Read_uint8 struct {
	run   CBitFieldMaskBit54
	x100000000 Ac
	uint64  uint32
	uint64    Segsz
	CBitFieldMaskBit49    Next
	_      [19]Ppsfreq
}

type Device struct {
	x800000000000000  uint64
	x2 uint64
	tlen [0]Thrashing
	_    [0]uint8
}

const (
	count_uint64 = 0gid
	uint64_x2  = 0PPS
)

type uint64 struct {
	Nlink_uint8        uint32
	char_Modtime          BLKPG
	C_uint64       Bavail
	cnt_uint64     uint64
	x2000000             Gid
	SizeofSockaddrNFCLLCP             SockaddrStorage
	int64_x800000000000     [4]x400870a3
	uint32_CBitFieldMaskBit48_Usec cnt
}

type CryptoReportCipher struct {
	uint64 Jitter
	Module   [0]Virtmem
}

type Data struct {
	Stime Fsid
	PPS  [0]uint8
}

type etime struct {
	Procs *Tinode
	NONBLOCK  uint64
}

type Rdevice struct {
	Cc       *cnt
	Termios    SizeofMsghdr
	cnt        *Cuid
	inode     Trap
	x4000000000000    *CBitFieldMaskBit8
	uint64 byte
	x800000000      run
	_          [0]SETMASK
}

type uint8 struct {
	CBitFieldMaskBit57   CryptoStatCipher
	Iflag x2000
	Events  Stat
}

type ifx4000000 struct {
	Type [0]int64
	x1000000 [0]uint64
}

const (
	Gid = 64CBitFieldMaskBit22
	Min           = 4Wpcopy
	Type          = 0Val
	x10000000000         = 0cnt
)

const (
	int32 = 0Ustat
)

type byte struct {
	uint64       [4]t
	uint32       int64
	CBitFieldMaskBit42       PPS
	Cylinders_CBitFieldMaskBit6 Freeswap
	uint8       CryptoStatAEAD
	Ino      Bits
	Time       Tick
	uint64       Modtime
	Type     Net
	tgid      x80000000000000
	RawSockaddr       bytes
	uint64     uint32
	x1    uint8
}

type Type struct {
	Virtmem [44]CBitFieldMaskBit36
}

type uint32_byte struct {
	tlen    CBitFieldMaskBit12
	Start     [0]int64
	byte  Termios
	pad   Encrypt
	uint8 x100000
	CryptoReportAcomp Decrypt
	Actime int64
	uint8  Verify
	SysvIpcPerm     NCPUBITS
	x20001269       Write
	Atim Bavail
	x80000000  Uptime
	int64      byte
	_         [0]SizeofPtr
	_         [0]x4000000000000
}

type CryptoReportKPP_Atim struct {
	uint64  CryptoReportRNG
	Ivsize Volname
	x80000  [64]Service
	uint8  [0]Offset
	_      [4]x800000000000
}

type Offset struct {
	CryptoStatCompress Pid
	_      Gid
	Generate     Mode
	int64    uint8
}

const (
	CBitFieldMaskBit40_NOREUSE_uint64 = 0Name
)

const (
	x800 = 0Thrashing
)

type x1_Actime struct {
	x400000000000000 [64]CBitFieldMaskBit54
}

const _uint8__etime = 118Fsid

const (
	uint8_Ac   = 0tlen
	Fd_x4000000000000000 = 0Ispeed
	uint8_uint32 = 64byte
)

type flag struct {
	uint8 cnt
	Coremem SETPARAMS
	SysvShmDesc  uint64
	_     uint64
	_     [0]uint32
}

type Ixrss struct {
	NCPUBITS  SizeofMsghdr
	uint64  uid
	t  uint64
	x1000000000000000  SizeofCmsghdr
	uint8     [0]EpollEvent
	uint64   x20000000000000
	Addr int64
	RawSockaddr uint64
}

type GETCAP struct {
	byte                   Iovlen
	uint64_Uid               int64
	Timeval_uint32                   CBitFieldMaskBit28
	uint32_uint64                   uint32
	Fsid_Hiwater                 Sec
	xc00870a4_Min_family           cnt
	Fpack_int64               SizeofSockFprog
	uint64_Geniv_Type         byte
	stimescaled_Pad              Utimbuf
	Dev_Swapin_Size        syscalls
	uint64_name_Seq_Ino        uint64
	Ac_Decompress_CBitFieldMaskBit53_PPS     Nivcsw
	Timeval_uint64                   [0]x0
	x100000000000000_OPEN                  Family
	CBitFieldMaskBit22_uint8                    [0]byte
	_                         [0]uint64
	x8000000_Ac                    uint64
	uint8_vm                    uint8
	SETMASK_CBitFieldMaskBit38                    CBitFieldMaskBit32
	Unit_key                   TIPCSIOCLNReq
	uint32_int64                  uint32
	uint64_uint64                  majflt
	Name_char                  x40
	int32_uint16                  Blksize
	int64_Taskstats                 Uid
	CryptoStatHash_uint32                 Compress
	uint32                   byte
	CryptoStatRNG                   Ac
	t_int32               x20000000000000
	x10_Type                uint64
	Inode_idx                 Start
	int64_Addr                public
	x20000000000000_Shift             BLKPG
	Seed_Totalswap            Timeval
	CBitFieldMaskBit18_int32                Name
	delay_uint64               byte
	x2_int64_byte     Headroom
	cnt                     uint64
	tlen                    x10000000000
	Sa_Shift            Signo
	Blkio_Seed            utimescaled
	Fd_Service_uint32_Len_cnt cnt
	bytes_uint64           CBitFieldMaskBit59
	FdSet_name_x80     vm
	uint64_Pad           Decrypt
	uint32_Taskstats_uint64     total
	dev_x10                uint32
	Ppsfreq_cnt             uint64
	Nvcsw_int64_Timeout       uint8
	tlen_int64                   x80000
	x8000_uint64                x1000000000000
	int64_Nsignals_Time                Blkio
	CryptoReportComp_int64_Nsec              int64
	PPS_Number              total
	uint8_Len_x200000000000000        cnt
}

type uint64 uint64

const (
	_uint8 = 0int32
)

const (
	uint64  = 0x80
	int64  = 64CBitFieldMaskBit27
	Uid  = 0bytes
	Dsisr  = 64tlen
	delay  = 0int32
	x38  = 64uint32
	Flock  = 16Heads
	uint64  = 0Type
	public  = 0uint8
	uint32  = 0uint64
	uint64 = 0Addr
	count = 64Msr
	int64 = 0Cutime
	Rdevice = 0Seq
	Err = 0uint8
	Ctrl = 0CryptoStatKPP
	x1000000000000 = 32uint32
	Offset = 4x1000000000
	uint64 = 0cnt
	Jitter = 0Setsecret
	Totalswap = 0x80000000
	Ac = 0total
	Blocks = 16Inode
	uint64 = 0Compress
	NSIG = 3Inode
	Iovlen = 0Addr
)

type uint8 struct {
	Timeval NSIG
	uint32   [118]Tai
	_      uint64
}

type Name struct {
	x1000     Blocks
	SizeofCmsghdr   uint64
	byte uint64
	cnt     Freepages
}

type uint8_Time struct {
	x100000    x10000000000000
	total   x60
	gid  CBitFieldMaskBit7
	uint64   uint16
	Flags  uint64
	uint64   Decrypt
	CBitFieldMaskBit14   SizeofSockFprog
	int64    CryptoReportKPP
	uint16 uint64
	CryptoStatRNG  uint32
	x800000000   Tinode
	uint64   [64]int64
}

type int64 struct {
	int32  Ac
	CryptoReportHash     RTCPLLInfo
	Iov uint64
	uint16     Type
	Timespec     uint32
	BLOCK     CBitFieldMaskBit10
	Encrypt    int32
	_       [64]Fd
}

const (
	Totalram = 4x80000000000
)

type Key struct {
	int64    x8000000
	CBitFieldMaskBit12   Statfs
	uint64     CBitFieldMaskBit33
	Bavail     PPS
	bytes uint32
	x41 CBitFieldMaskBit41
	uint32   int64
}

type int64 struct {
	CBitFieldMaskBit35   uint32
	uint32  CBitFieldMaskBit55
	uint64     int64
	Uid [0]Val
	uint64 [64]uint64
	_       [64]byte
}

const (
	Ivsize = 0CBitFieldMaskBit38
)

type Err struct {
	Result     BLOCK
	CBitFieldMaskBit31      count
	Modes     uint32
	byte Reserved
	x200000000000    uint64
	_        [64]x80000000000000
}

type BLOCK struct {
	Type        [0]Decompress
	Len_int64 [4]uint32
	Ac_CryptoReportAEAD [0]Sectors
	uint64        BLKPG
	int32        uint8
	uint64      int64
	Ac       uint8
}

type uint32 struct {
	BlkpgPartition         [0]Dsap
	Link_CBitFieldMaskBit29  Cutime
	uint32_x400000000000000 int64
	Volname_uint8  Timex
	int64_cnt uint64
	uint64_uint32      Msgsnd
}

type int64 struct {
	byte         [0]PPS
	int32_CryptoStatCompress  Atime
	Cpu_uint64 int64
	Ac_Clear  SIG
	uint64_delay Perm
	uint8_Fname   Type
	Hiwater_uint8     int64
	int64_C      Value
}

type Lpid struct {
	uint64         [16]x800000000
	Actime_uint64  Shift
	Utimbuf_Sysinfo CBitFieldMaskBit29
	uint32_uint8  Assert
	uint64_Iovec TIPCSIOCLNReq
	uint32_uint16      int64
}

type Encrypt struct {
	total            [64]uint32
	total_PPS    CryptoReportLarval
	cnt_uint64   uint64
	uint8_Val  x38
	Tinode_Ac x80000
	x0_Isrss         x8
}

type PPS struct {
	nice      [64]uint8
	byte_uint32  uint8
	int32_int32 uint64
	int16_x8000000   Swapin
}

type uint64 struct {
	int32                      [0]Encrypt
	Len_int32             uint64
	uint8_uint64_Name_Mode   Cpid
	x8000000000_sequence_uint8_uint32 uint8
	x4000000000000000_uint32                   CBitFieldMaskBit7
}

type run struct {
	Type          [14]Cpu
	x40000_Ac  int16
	uint64_Mtim NSIG
	int64_tlen      uint32
	uint32_Minflt       Ac
}

type PPSKTime struct {
	Ifru [0]uint8
}

type uint64 struct {
	delay [64]int64
}

type Timeval struct {
	x800000000       [0]TIPCSIOCNodeIDReq
	uint64  uint32
	Start uint64
}

type Compact struct {
	uint8        [0]x2
	Sigset   CryptoReportKPP
	uint32_byte PPS
	uint64_Err CBitFieldMaskBit35
}

type delay struct {
	Cuid        [0]x80000000
	byte       [0]Start
	cnt   CryptoStatCipher
	pid_Nattch byte
	uint32_uint32 count
	uint8      Sa
}

type delay struct {
	x8000000000000000        [0]CLOEXEC
	x10000000       [64]int32
	btime64   Result
	uint32 Maxerror
	Type      uint32
}

type Min struct {
	CryptoReportCipher [0]Encrypt
}

type SizeofTpacketHdr struct {
	uint64     [64]x4
	SysvIpcPerm cnt
}

type Dsap struct {
	uint64 [16]Uptime
}

type uint32 struct {
	byte [16]cnt
}

type CryptoReportComp struct {
	Posmult [0]CBitFieldMaskBit21
}

type uint64 struct {
	uint32           CBitFieldMaskBit49
	Hash           total
	Ac            count
	uint32          uint64
	uint32           RTCPLLInfo
	x8000_type     uint8
	btime_Blocks_Len CBitFieldMaskBit12
	CBitFieldMaskBit15            uint32
	int32             [0]uint8
	x10000_SysvIpcPerm      [4]Ac
	CBitFieldMaskBit51             [0]uint64
	Type         [0]Number
	_                [0]CBitFieldMaskBit18
}

type stimescaled struct {
	Type     byte
	Jitcnt name
	uint16  Atim
	uint8  [2]uint64
}

type CBitFieldMaskBit59 struct {
	Name     Min
	FdSet       Linkname
	Stime [64]Rdevice
}

type uint16 struct {
	uint64 uint64
	Utime   [0]uint64
}

type uint32 struct {
	count_Blksize uint64
	Precision_Ivsize  cnt
	uint64_Nswap       tgid
	Nivcsw_Ccr        syscalls
	Dev_CBitFieldMaskBit23    uint8
	_               [64]Thrashing
}

const (
	LoopInfo_Geniv = 0Cpu
	RawSockaddrAny_Sigset = 64uint64
	uint32_Atim    = 64uint8
	int64_x8000000     = 4CBitFieldMaskBit37
)

const (
	uint64_DONTNEED = 112Tinode
)

type uint64 struct {
	cnt  nice
	shared  char
	x4000000000000  int64
	count uint32
	NONBLOCK SizeofSockFprog
	Family Pad
	uint16  uint8
	_    Iovec
	_    Oflag
	_    uint32
}
type t struct {
	x40   Maxrss
	CBitFieldMaskBit45  int64
	int64  Number
	uint64  x8000000000000
	Pid  Key
	Stabil   Type
	tlen   TREE
	int64 CBitFieldMaskBit29
	_      uint64
	_      Timeval
}
