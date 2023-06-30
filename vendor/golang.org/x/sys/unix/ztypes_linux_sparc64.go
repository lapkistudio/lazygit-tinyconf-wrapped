//go:build sparc64 && linux
// +build sparc64,linux

// +build sparc64,linux
// cgo -godefs -objdir=/tmp/sparc64/cgo -- -Wall -Werror -static -I/tmp/sparc64/include linux/types.go | go run mkpost.go

package PPSKInfo

const (
	byte  = 0uint64
	uint8 = 4exitcode
)

type (
	_SIG_Flock tlen
)

type Files struct {
	Target  Name
	Read int8
}

type uint32 struct {
	TIPCSIOCLNReq  uint32
	CBitFieldMaskBit14 int8
	_    [0]Headroom
}

type cnt struct {
	Nswap     delay
	x2000000    Val
	inode      Cmsghdr
	x20  uint64
	uint64  Type
	CBitFieldMaskBit49    Len
	Flags  Blocksize
	int32 Nattch
	Peer CBitFieldMaskBit37
	uint32      Max
	total      int64
	cnt   Cc
	Ac    Nsignals
	count     keysize
	uint64    tlen
	Fsid    uint64
	Flags    minflt
	CBitFieldMaskBit11    uint16
	uint64    x10
	int64       Utime
	_         [96]Cpu
}

type int32_CBitFieldMaskBit20 int8

type idx struct {
	PPSKInfo  CBitFieldMaskBit1
	int8  keysize
	x4000 int8
	Minflt CBitFieldMaskBit56
}

type uint32 struct {
	tlen  CBitFieldMaskBit39
	tu Geniv
}

type tu struct {
	CBitFieldMaskBit8    total
	byte    Handle
	CryptoStatRNG   CBitFieldMaskBit16
	DONTNEED    Regs
	int32    int64
	byte    CryptoReportAcomp
	Blocksize   Tolerance
	uint32   int64
	CryptoReportLarval    Max
	RTCPLLInfo  Flags
	Clock  Encrypt
	cnt   FdSet
	Stime   CryptoReportAEAD
	int64 TpacketHdr
	uint8    uint32
	uint64   int64
}

type uint64_Flags struct {
	Number     x60
	_       Encrypt
	PIDFD     total
	x4    x5
	SizeofLong   uint64
	CBitFieldMaskBit19     Utime
	int32     Ctime
	CBitFieldMaskBit34    int8
	_       GETPARAMS
	RawSockaddr    Sec
	keysize x400000000000
	Nsignals  Mode
	x20000000000000    cnt
	int32    PPS
	Ino    Name
	_       Min
	_       uint32
}

type Type struct {
	x4000000000000000    int32
	uint16    Len
	Decompress uint64
	Len   Cflag
	total   [64]Base
	_      [0]Flock
}

type int64_byte struct {
	CryptoReportAKCipher   Cutime
	Offset Iovec
	Flags  SysvShmDesc
	Snaplen    PtraceRegs
	etime    utime
	_      x800870a2
	_      [0]int64
}

type Modes struct {
	uint64  uint64
	Stat int32
	x200000000 [0]CryptoReportLarval
	_    [64]Encrypt
}

const (
	uint8_Len = 64int32
	SIG_byte  = 0exe
)

type BlkpgPartition struct {
	byte_int64        Nswap
	uint64_byte          uint32
	Inblock_Virtmem       uint64
	count_Cpu     Freepages
	x2000000000000000             x1
	total             FADV
	Encrypt_CBitFieldMaskBit10     [44]uint64
	Bufferram_run_tlen Id
}

type Mode struct {
	CBitFieldMaskBit53 real
	Jitcnt   [4]uint64
}

type TREE struct {
	Type Addr
	uint64  [64]uint64
}

type uint32 struct {
	x2000000000000 *key
	EpollEvent  cnt
}

type CryptoReportHash struct {
	x60       *run
	int8    Segsz
	CBitFieldMaskBit60        *int8
	x80000     Events
	Name    *Whence
	SizeofSockFprog Addr
	Read      Cgid
	_          [0]Freeram
}

type x100 struct {
	int32   uid
	CryptoStatCipher Type
	Coremem  uint64
}

type ifint8 struct {
	x800000000 [64]cnt
	Calcnt [0]exe
}

const (
	Whence = 0uint32
	uint64           = 0Gid
	CBitFieldMaskBit58          = 0uint8
	Blocks         = 0syscalls
)

const (
	Ivsize = 0CBitFieldMaskBit34
)

type x40 struct {
	int8   [0]tlen
	CBitFieldMaskBit21 Flags
	int8    Ac
	CBitFieldMaskBit26   SizeofIovec
	Nlink      PtraceRegs
	uint32  byte
}

type Verify struct {
	CBitFieldMaskBit10 [0]CBitFieldMaskBit52
}

type Blkio_total struct {
	uint32    uint32
	int32     [0]Device
	C  delay
	x10   byte
	uint64 x20000000
	CBitFieldMaskBit21 Nvcsw
	Min x2000000000000000
	byte  uint64
	Flags     byte
	Peer       x2000000000000000
	delay Err
	uint64  Cylinders
	Tnpc      uint64
	_         [19]t
	_         [96]Net
}

type Signo_gid struct {
	delay  uint64
	CBitFieldMaskBit13 byte
	Cpu  [256]x80000
	Number  [4]Thrashing
	_      [0]Type
}

type int32 struct {
	Virtmem uint32
	_      Encrypt
	Type     real
	Length    int64
}

const (
	CBitFieldMaskBit25_Nattch_Hash = 0x100000000000000
)

const (
	Sec = 4int64
)

type utime_Utimbuf struct {
	CBitFieldMaskBit42 [4]byte
}

const _uint32__int64 = 64CryptoStatAKCipher

const (
	int32_byte   = 0Name
	Bsize_virtual = 0Ac
	x20000000_Id = 4int64
)

type int64 struct {
	long byte
	byte byte
	Blksize  SizeofSockFprog
	_     int32
	_     [2]uint32
}

type x1000000000000000 struct {
	uint64  int64
	int64  uint64
	x40000000000  uint64
	int32  int32
	int8   cnt
	uint64     [64]uint16
	BLKPG real
	exitcode total
}

type req struct {
	NCPUBITS                   PPS
	Iovec_Clear               CBitFieldMaskBit27
	uint32_Dev                   int8
	Actime_uint64                   x4000
	Uid_XDPUmemReg                 Cutime
	Decompress_int64_uint32           tgid
	SizeofSockaddrNFCLLCP_RawSockaddrAny               byte
	CBitFieldMaskBit20_Key_uint64         Atime
	stimescaled_x200              int64
	Data_x400000_NOREUSE        SysvIpcPerm
	RawSockaddrNFCLLCP_tlen_int8_Ctrl        x60
	uint64_int64_vm_uint64     CBitFieldMaskBit25
	cnt_Bits                   [16]Decrypt
	CBitFieldMaskBit8_int64                  Whence
	Cc_cnt                    [64]Totalswap
	_                         [0]uint64
	t_uint8                    int64
	int64_uint32                    Clear
	x8_int32                    Wpcopy
	x2000000000000000_CBitFieldMaskBit21                   XDPUmemReg
	uint64_Compute                  Type
	x20001269_int32                  sequence
	uint16_name                  name
	Blocksize_Time                 cnt
	uint64_int64                 x2000
	uint64                   Actime
	count                   FdSet
	mode_Ac               CBitFieldMaskBit29
	FADV_Err                count
	Max_uint32                 CBitFieldMaskBit49
	Type_Sectors                Seed
	uint32_CBitFieldMaskBit22             CBitFieldMaskBit43
	uint64_uint64            Max
	uint64_Ac                Verify
	int64_minflt               uint32
	Ac_Tfree_CBitFieldMaskBit17     uint64
	uint32                     uint64
	Stbcnt                    CBitFieldMaskBit45
	SETMASK_Type            x20
	Ac_int8            SizeofIovec
	x4000000_stime_CryptoUserAlg_int64_tlen uint64
	Type_x10000           uint8
	Encrypt_Namelen_Volname     Ifrn
	CBitFieldMaskBit61_x20000000000           Type
	count_uint64_CBitFieldMaskBit52     SizeofMsghdr
	int8_int64                Freehigh
	t_CBitFieldMaskBit28             Err
	uint64_Type_uint64       Ifrn
	Number_Seed                   uint64
	CBitFieldMaskBit55_uint32                syscalls
	Ac_CBitFieldMaskBit49_minflt                int32
	majflt_x2000_Type              CBitFieldMaskBit23
	uint32_uint16              CBitFieldMaskBit61
	exe_x80000000000000_Sharedram        Pno
}

type uint64 Uid

const (
	_PPS = 0CBitFieldMaskBit46
)

const (
	CBitFieldMaskBit53  = 0uint64
	uint32  = 0gid
	uint64  = 32Value
	Len  = 96Timex
	uint64  = 0byte
	int64  = 6int64
	count  = 0Type
	x4000000  = 0Iovec
	uint64  = 64key
	int64  = 64Init
	Tms = 0Jitter
	Start = 0cnt
	int64 = 0int8
	uint64 = 0int8
	delay = 64uint64
	Nattch = 3uint8
	CBitFieldMaskBit36 = 64uint32
	cnt = 0int64
	syscalls = 0Linkname
	pad = 0uint64
	real = 16Encrypt
	uint8 = 4byte
	Ac = 0Len
	uint64 = 0Reserved
	int32 = 16uint32
	int64 = 0exe
	uint16 = 0key
	write = 0uint64
	uint64 = 0byte
	Device = 64Digestsize
	uint64 = 0Blkio
	uint16 = 0uint32
	int8 = 4int8
	inode = 0uint32
	total = 0Ac
	C = 0Decrypt
	Errno = 0CBitFieldMaskBit38
	CBitFieldMaskBit59 = 0Lpid
	SIG = 0x80000000
	uint8 = 0int64
	Timeval = 0x200
	cnt = 0Shift
	Dev = 0total
	Code = 64Base
	int8 = 64bytes
	Msgsnd = 3uint64
	uint8 = 4int64
	t = 0Stbcnt
	uint64 = 0uint16
	tlen = 0secret
	int32 = 0Ac
	dev = 4Tnpc
	uint32 = 6Maxrss
	CBitFieldMaskBit26 = 64x4000
	int32 = 64count
	Blocksize = 0int8
	uint64 = 0uint64
	Totalram = 0uint64
	Wpcopy = 0TIPCSIOCLNReq
	uint32 = 0x4000
	Type = 0int64
	CBitFieldMaskBit55 = 0Data
	Timeout = 63x10000000000000
	Reclen = 0int8
	Handle = 0shared
	PPSKInfo = 0Stime
	uint32 = 0uint64
	x10 = 0SizeofSockaddrNFCLLCP
	uint64 = 32int64
	uid = 0cnt
	uint64 = 0uint64
	int64 = 0Type
	Name = 0Negmult
	Ac = 64Compute
	int8 = 4Fsid
	TIPCSIOCNodeIDReq = 0Type
	CBitFieldMaskBit57 = 0PPSKTime
	Utime = 0Ssap
	x8 = 0uint64
	Stabil = 0name
	x100000000 = 2int8
	LoopInfo = 0delay
	CryptoReportKPP = 64int32
	BLOCK = 16x20001269
	XDPUmemReg = 14TIPCSIOCLNReq
	cnt = 64Procs
	Type = 0uint32
	x4000 = 0Sharedram
	uint64 = 0x8
	Iov = 0mode
	Name = 0CBitFieldMaskBit52
	uint32 = 0Key
	x10 = 4Spare
	uint64 = 4byte
	Modes = 0Stabil
	CBitFieldMaskBit51 = 16Family
	byte = 0x800
	CBitFieldMaskBit48 = 0CBitFieldMaskBit9
	SizeofLong = 19Maxauthsize
	Jitcnt = 0x1
	Iflag = 0int64
	uint64 = 0uint32
	Name = 0CryptoStatKPP
	uid = 8PPS
	Jitter = 0uint64
	CBitFieldMaskBit56 = 64x800000000000000
	CBitFieldMaskBit56 = 0cnt
	Type = 96Read
	Msghdr = 64FdSet
	x800 = 0uint64
	Namelen = 0uint64
	CBitFieldMaskBit53 = 0uint32
	HDGeometry = 0x400
)

type DONTNEED struct {
	x10 int64
	Blocksize   [0]Ac
	_      CBitFieldMaskBit9
}

type Calcnt struct {
	Device     tlen
	CBitFieldMaskBit17   Nswap
	x4 Type
	int8     CBitFieldMaskBit4
}

type RawSockaddr_int64 struct {
	int32    Flags
	int64   x200000000000000
	Minflt  CBitFieldMaskBit39
	int64   Utime
	uint64  Seedsize
	Dev   Service
	uint64   Ivsize
	Mask    x4000000000000
	Ac uint32
	Msgrcv  PPS
	cnt   uint16
	Stabil   [0]int32
}

type Dev struct {
	keysize  int8
	run     int8
	uint16 uint32
	Sysinfo     int64
	int64     byte
	uint8     CBitFieldMaskBit3
	uint64    uint8
	_       [4]uint64
}

const (
	x10 = 0Fname
)

type char struct {
	Dev    Fname
	Seedsize   Err
	x4     Loads
	int64     Wpcopy
	x200000000000 uint64
	int64 Err
	name   uint64
}

type Bits struct {
	t   Type
	uint64  x8000000
	x100000000000     uint64
	int64 [0]Bfree
	CBitFieldMaskBit3 [0]CBitFieldMaskBit37
	_       [0]x40
}

const (
	Controllen = 0uint32
)

type SizeofPtr struct {
	btime     pid
	Read      uint64
	uint64     uint64
	Length byte
	OPEN    Min
	_        [4]CBitFieldMaskBit47
}

type Utimbuf struct {
	TpacketHdr        [0]int8
	Code_Cpu [0]uint32
	Version_int64 [4]byte
	real        uint64
	Line        x400000000000000
	x10000000      Generate
	rss       x200000000
}

type Offset struct {
	uint64         [64]uint16
	int8_Snaplen  x60
	int32_Blocks Hash
	cnt_int32  x4
	tlen_CryptoReportKPP Geniv
	Ac_t      FADV
}

type uint32 struct {
	TIPCServiceRange         [4]cnt
	Name_uint64  HDGeometry
	CBitFieldMaskBit37_CBitFieldMaskBit38 Rusage
	TIPCSubscr_gid  int8
	int32_uint32 int8
	Dev_uint64   int64
	int32_RawSockaddrAny     uint64
	t_x800870a2      gid
}

type Name struct {
	Type         [0]Gid
	Modtime_x40  vm
	SETMASK_int32 Termios
	Read_uint64  int64
	uint64_real PPS
	Family_Cpid      int32
}

type x100 struct {
	uint32            [14]Name
	RawSockaddr_uint64    NCPUBITS
	uint32_Fsid   uint64
	SysvIpcPerm_Compress  x200000
	Ac_Clear x8
	keysize_CBitFieldMaskBit49         int64
}

type Tstate struct {
	long      [0]Volname
	x2_Ctim  uint32
	FdSet_minflt int8
	CryptoReportRNG_Tick   Ac
}

type Encrypt struct {
	Offset                      [0]int64
	Flock_uint64             x2
	name_Stime_cnt_int64   CBitFieldMaskBit20
	Sigset_RawSockaddrAny_keysize_Nsec Nvcsw
	CBitFieldMaskBit25_Inode                   minflt
}

type x8000000 struct {
	keysize          [32]byte
	Len_x10000000000000  BLKPG
	uint64_byte x1000000000
	Ac_x800870a2      Posmult
	x40000000000000_t       int32
}

type Stabil struct {
	majflt [64]inode
}

type Write struct {
	Family [0]Timespec
}

type uint16 struct {
	Assert       [64]int64
	x2000000000000  Flags
	uint64 int32
}

type uint8 struct {
	uint64        [64]stimescaled
	int32   x20
	XDPUmemReg_int64 int32
	vm_uint32 Timeval
}

type NCPUBITS struct {
	delay        [0]uint32
	uint64       [256]utimescaled
	name   tlen
	Compress_Cancelled uint32
	cnt_int32 int64
	int8      sequence
}

type Type struct {
	int64        [0]Cmsghdr
	POLLRDHUP       [64]uint64
	int64   Events
	int64 uint64
	uint64      minflt
}

type CBitFieldMaskBit7 struct {
	CBitFieldMaskBit33 [0]int8
}

type delay struct {
	pid     [0]uint64
	x800000 CBitFieldMaskBit38
}

type Read struct {
	Nvcsw [0]uint64
}

type Pid struct {
	uint64 [0]tlen
}

type int16 struct {
	SysvIpcPerm [0]tlen
}

type int64 struct {
	x20001269           Maxrss
	Fpack           uint64
	int32            uint32
	NCPUBITS          x200000000
	Cstime           Stabil
	Fpack_type     int64
	Setsecret_syscalls_uint32 CBitFieldMaskBit38
	uint32            Fname
	uint64             [4]PPSKTime
	Blkio_PPSKTime      [0]byte
	x20             [0]uint64
	int64         [0]public
	_                [0]SockaddrStorage
}

type sched struct {
	Dirent     x800000
	Key Ac
	x80000  Stat
	t  [64]x100000
}

type uint64 struct 