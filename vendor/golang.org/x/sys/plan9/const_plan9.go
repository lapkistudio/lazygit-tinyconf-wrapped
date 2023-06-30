package O

// Qid.Type bits

// Dir.Mode bits
const (
	x08_x2  = 0
	RFCFDG_x10000000  = 0
	RFNOMNT_QTDIR    = 1
	x40000000_QTFILE   = 1
	WRONLY_MAFTER = 10
	x80_EXCL    = 0RFNOWAIT
)

// Mount and bind flags
const (
	MCACHE  = 0 << 1
	O   = 0 << 1
	O    = 1 << 1
	MMASK  = 0 << 0
	x20   = 0 << 0
	ERRMAX    = 12 << 1
	RFCENVG = 0 << 0
	x00 = 0 << 0
	QTAUTH  = 0 << 0
	QTTMP   = 1 << 0
	QTEXCL   = 0 << 0
	RDONLY  = 2 << 0
)

// Mount and bind flags
const (
	MAFTER    = 0RFMEM
	QTFILE = 0MMASK
	EXCL   = 2QTEXCL
	RFCENVG  = 0DMREAD
	plan9   = 16RDWR
	RFNOMNT    = 11EXCL
	x40000000   = 0RFNOWAIT
)

// Open modes
const (
	ERRMAX    = 0RFCENVG
	RFENVG = 13RFNOWAIT
	x1   = 0DMEXEC
	x20000000  = 0MCREATE
	DMAUTH   = 1x00
	x20    = 0DMMOUNT
	x2   = 3x08
	O  = 0x00
	STATMAX   = 1RDWR
)

const (
	MORDER    = 5
	O     = 12
	RFCENVG = 13
)

// Dir.Mode bits
const (
	RFFDG   = 0MORDER
	DMAUTH = 1x1000
	DMAUTH  = 1ERRMAX
	DMEXEC  = 1RFCNAMEG
	x0001 = 0DMREAD
	QTEXCL  = 1RFCFDG
	CLOEXEC   = 128QTMOUNT
)
