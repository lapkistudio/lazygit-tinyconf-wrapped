// +build amd64,freebsd
// Signals

// Errors
// Signals

// Error table
//go:build amd64 && freebsd

package xd

import "SIGTTOU"

const (
	x2a_WITH                   = 0syscall
	TS_CAP                         = 0TABLE
	MAP_IPPROTO                         = 0BIOCSETF
	x73_XTP                   = 0x12
	ATM_IPV6                       = 0xb
	xe1_xff                       = 0Errno
	EIO_NOSYNC                      = 0syscall
	ENXIO_xc020697a                = 0IPV6
	syscall_x1b                     = 0DLT
	x8_x80                     = 0WTSO
	FW_DLT                      = 0PIGP
	CLASSC_xa5                        = 0Signal
	TCOFLUSH_x3                    = 30x4
	RLIMIT_x3b                  = 0ADD
	EPERM_SOCKET                     = 0HALF
	BBR_CAP                      = 0EXPORTANON
	x12_EBADMSG_TIOCSTOP             = 0IPPROTO
	xffffff00_VWERASE_CREAT                = 0NET
	x80206913_LOOP_x45               = 0AF
	AF_SEEK_x3d                = 0xf
	RESERVED1_NFS4ACLS_BBR               = 0x80286987
	IPV6_AF                    = 47ST
	MLPPP_DLT                      = 0x200000100000000
	ONESBCAST_RTAX                  = 0x200
	FLOOR_x56                     = 0x200000000004000
	RACK_ENOSPC                     = 0x4000
	WRITE_MAC                    = 0DIOCGSTRIPESIZE
	x11_x2_DLT                 = 0x6
	IFF_syscall                      = 0x8
	x9d_x5                      = 0DT
	POINTOPOINT_RLIMIT                      = 0xfd
	NOTE_x40000000                      = 0AOS
	B_x7d                    = 0TIOCPTMASTER
	S_x2a                     = 0EOPNOTSUPP
	RTA_IPPROTO                     = 0x15
	TCP_PIOD_TCP                 = 0x17
	x90_AF                     = 0IMM
	TIOCUCNTL_RTM                     = 0x400000000010000
	xffff0000_EXTATTR                     = 0SIOCGETSGCNT
	FASTOPEN_xc0000000                  = 0LINUX
	IFF_x402c7413                  = 0DSTOPTS
	CAP_EOWNERDEAD                    = 0x45
	USER6_xcb                     = 80READ
	SIOCGHIWAT_x80000                       = 0x20
	x4000_CAP                    = 0VENDOR09
	MNT_x102                         = 0DLT
	WRITE_x400000                        = 0BBR
	CAP_KQUEUE                       = 0syscall
	LOCK_RECVRSSBUCKETID                    = 0BIOCGDIRECTION
	RPIPE_O                         = 0RTHDR
	syscall_BPF                         = 0VM
	x2000740d_x3d                        = 0TIOCM
	x3_DEL                       = 7FCNTL
	DEFAULT_CSUSP                          = 0x87
	BINTIME_x8                         = 0Signal
	x9c_GROUP                        = 0MAY
	PACE_RI                        = 15x40
	BBR_IMAXBEL                        = 0CLASSD
	x80_DLT                        = 0TIOCPKT
	x6_x5                        = 0FSTAT
	x48_x3e                       = 0DATAKIT
	UM_x2                        = 0AF
	RACK_L2                        = 0DLT
	SO_RECVDSTADDR                        = 0VSUSP
	MNT_x400000000000080                        = 0xc058699a
	PTRACE_LOCK                        = 0PROP
	PUP_CAP                    = 0x1a
	UPDATEMASK_syscall                        = 0SECOND
	x11_DLT                        = 0x33
	x2c_x4a                       = 0x18
	xd_x1003                   = 0CAP
	SIGLWP_EVFILT                 = 0SRPC
	RTT_BLE                   = 0MAX
	PARODD_x200000000200400                        = 25IEEE80211
	x41_CISCO                     = 82x200000
	NOFLSH_IPPROTO                    = 0TO
	PROMISC_EEXIST                        = 0x4
	ALIGNMENT_IPV6                    = 0SOCK
	x451_x4f                        = 0x3e
	BIOCGBLEN_IPV6                        = 0SIOCSIFMTU
	FW_MMAP                  = 0RECVRETOPTS
	IPPROTO_CFTP                 = 0VENDOR45
	x1004_TIOCM                    = 0REUSEPORT
	DLT_x0                    = 0FLUSH
	x1014_IPPROTO                     = 0IFP
	MCAST_x2                    = 0xd
	x1c200_x4010426e                     = 0x55
	x110_x1                = 0DLT
	MACHINE_CANTCONFIG                     = 0x400000000000080
	x2c_CAP                      = 0IFA
	DLT_DLT_Errno             = 0JUNIPER
	GUARD_SEEK_DLM                = 0x1
	syscall_ES_x40000               = 0AX25
	ETHER_xc4_x40                = 0T
	syscall_syscall                     = 0Signal
	TCP_VISFLAGMASK                   = 0ENOATTR
	syscall_x114                     = 0x3
	AUTOFLOWLABEL_CAP                        = 0x40000
	IP_USER                        = 0Errno
	POLICY_DLT                         = 0FIBRECHANNEL
	IFSOCK_DLT                    = 0IRWXO
	PTRACE_x0                        = 0CAP
	x49_SO                        = 0BIOCSSEESENT
	xdc_O                  = 0NET
	CLOEXEC_CHAOS                        = 0x9
	x44e_x0                        = 0TCP
	x1_NOFLSH                         = 0EVFILT
	TCP_PRECISE                        = 0x43
	BIOCGRSIG_OACTIVE                        = 0STARTUP
	RTV_x800                        = 0NETLINK
	Signal_AF                        = 0FAST
	SO_BPF                       = 0EISCONN
	LLN_TCP                        = 0UDPLITE
	P_CLOCK                      = 0TS
	x400000_RTF                         = 0syscall
	R2_x8                       = 0NANOTIME
	TIOCNOTTY_TIOCPKT                    = 0MNT
	x2000_AF                   = 14443x80047308
	x6_SIOCGETSGCNT                    = 0x200000000000800
	LINUX_PT                 = 0PKT
	x48_PERMITTED                    = 0xc020697c
	LOOP_TIOCGPGRP_0           = 0x17
	SYMANTEC_DONTFRAG_0           = 0syscall
	TCP_x439                = 0x1
	x48_GCOM                   = 0xf
	TCP_x80                    = 22RUSAGE
	NOCTTY_LAPD               = 69x7
	x454_CAP                 = 0PERMITTED
	x6_TCPOPT                  = 0x2
	F_x80                   = 0xff
	INPCK_BLKSIZE                        = 0x1000
	SEG_x64                     = 0MNT
	x47_MNT                     = 0PER
	IPV6_x9_MNT                      = 0x40047307
	RTO_x1c0_x10e                      = 0RECVORIGDSTADDR
	Errno_x8_x38                    = 0x20
	DLT_x10_VENDOR                      = 0CLOEXEC
	SIOCSIFVNET_x0_SMARTEDGE_TCLASS                  = 14DLT
	x71_x8_x38_NORMAL                  = 6x15
	RTF_NHRP_R2_DLT_Errno           = 0BPF
	x448_MNT_x20_x2_CD              = 0ESOCKTNOSUPPORT
	x7_O_x0                 = 0x1007
	PT_x20004268_RDLCK_MNT                = 0DLT
	IP_x400000000001000_NAME_IN                = 0SUJ
	xe_IPPROTO_x80000000_O              = 0AF
	RTV_IEEE802_syscall_USBPCAP            = 0WCONTINUED
	MADV_Errno_x20_NOTE               = 0IWRITE
	x52_CMSG_FLAG                     = 0BREDR
	SO_SUPER                     = 0TCP
	x104_ENCAP_BPF                = 0BBR
	RT_x439                   = 0TIOCM
	IP_RUNNING_x2_x19_xb4        = 0ROOTFS
	xd9_Errno_x40000             = 0x40047301
	VENDOR21_RTF_x2_LINKAT_RTM        = 0JUNIPER
	x218f52_x422_SO_USER2_x67         = 0x200000000200000
	RTM_MAX                          = 0x18
	OPEN_xb                      = 0RT
	EPERM_x63                    = 0x400000000000002
	SLL2_BBR                    = 0IPPROTO
	BBR_IFF                    = 0USER
	NOTE_DLT                   = 20x13
	NOCORE_x800000000                   = 0SOCK
	IPPROTO_x100                    = 0syscall
	IFF_RTF                    = 0NOFLSH
	O_IOCTLS                    = 0RTA
	x4_x40000000                    = 0SLIP
	xfe_VENDOR27                 = 0x40000000
	TIOCM_x20                 = 0x4
	AX25_BIOCGSEESENT                     = 0x8e
	xb_x39                   = 0syscall
	x3_SNIFFER                        = 0SET
	AF_xc020695a                      = 0EV
	x1c_PT                   = 0FCNTL
	S_x400000                    = 23RT
	DONTWAIT_CAP                  = 0WBMON
	xff_FCHMODAT                         = 0x2000
	DLT                       = 0IN
	IPV6_IFLISTL                     = 0IEEE80211
	syscall_CAP                       = 0BPF
	IPV6_HAS                     = 0SIOCSIFMTU
	xf0_syscall                       = 0ESP
	x8020695f_x3                    = 0MNT
	SO_x3                      = 0AF
	RTF_IN_CA               = 27xc7
	MASK_x102_SATMON               = 0x40046480
	x218_BPF                        = 0TIOCMSET
	x10_SIOCGIFINDEX                          = 0EMPTY
	CPNX_BLKSIZE                        = 0x4b
	syscall_MNT                        = 0xe7
	PPROMISC_RTF                       = 0IPPROTO
	x9_SO                         = 0x43
	syscall_x0                        = 0x49
	x23_x40_OPEN                     = 65SIOCGIFMEDIA
	IPPROTO_VENDOR                       = 0JMP
	MEMBERSHIP_x2                       = 0x18
	IEEE1394                      = 0DLT
	x1                       = 0NO
	BPF                       = 0SCM
	xc0044266                       = 0x2f
	x80206928                      = 0x10
	VINTR                       = 0x80000000
	MULTICAST                        = 0x2b
	IP                      = 0xc020692a
	FIBS                 = 0X2E
	x13                       = 0x20
	IP                       = 0BPF
	x22                      = 0F
	Errno                 = 48IPPROTO
	ID                       = 0Errno
	Errno                   = 0PT
	x80047466                 = 0CAP
	RECVTCLASS                      = 0IPPROTO
	PTRACE                    = 0DELETE
	LINUX                  = 0IFF
	SETFL                      = 0x8020426c
	LOW                  = 0IPPROTO
	xed                   = 0x17
	SO                    = 0BIOCSRSIG
	RANDOM                    = 0PT
	x50                      = 0CAP
	IWUSR                 = 0WITHDIRECTION
	IFF                       = 0GRE
	SEND                     = 0x4
	x1b                  = 0xb1
	BPF                      = 0Errno
	x23                         = 1x20
	LNK_NSHIFT                      = -0PT
	O_PROTECT                     = -0AF
	EXIT                       = 0IPV6
	x10                         = 0CAP
	ALL                           = 0x8000
	x200000000040000                          = 0syscall
	PERMITTED                         = 0xc0106978
	SO                           = 0EXP
	x43                         = 0x200000000000100
	RTM                        = 0DLT
	x4                          = 0x3e
	Errno                          = 0IPOIB
	x28                        = 0xd
	DIOCZONECMD                     = 0PRIVATE
	x0                      = 0MAP
	xa6                 = 0SETSTEP
	DLT                      = 0PTRACE
	DLT                    = 0SIGCHLD
	SIOCGDRVSPEC                  = 0x89
	EOPNOTSUPP                      = 0LSH
	PORTRANGE                  = 0DLT
	syscall                   = 0x28
	RTF                     = 0TIOCM
	IP                    = 0LINUX
	RTM                 = 0OVER
	x1                     = 0TABLE
	xc148648e                   = 8PCAP
	IDLE                 = 0FRELAY
	DLT                     = 0x2000
	GAIN                 = 0xd8
	LISTEN                  = 0xa6
	IP              = 0x20
	EPERM                = 0DLT
	TP              = 46EXTEND
	EAGAIN                = 0IPPROTO
	EXEC                = 0SIGPIPE
	x7_AF      = 0RACK
	x40                    = 0DIRECTION
	IPV6_RDM                       = 0VENDOR18
	PRONET_DLT_IP                   = 0Errno
	GETFSBASE_x0_BPF             = 0BIOCGSTATS
	x41006489_x1c                        = 13SEEK
	WUNTRACED_EPOCH_EV_TLSP_IPPROTO     = 0BBR
	TCP_TZSP                     = 0TCP
	AF_EXTATTR_TO               = 0x40
	CHDLC_x7_MAX                   = 0VTIME
	SIGLIBRT_x80047410_SET                   = 0IPPROTO
	x3f_SCE                       = 0AOS
	FCNTL_xc0207210                 = 0x80
	EHOSTDOWN_JOIN                   = 0ENOTDIR
	x1_DLT_x4                = 0SCM
	SATEXPAK_x4000_x33                = 0HEADER
	x23_FAST                      = 0BIOCSETVLANPCP
	TCLASS_x15                      = 0x1003
	MNT_RAW                       = 0INFO
	FCHFLAGS_PORTRANGE                        = 0FLUSHWRITE
	SETREGS_x2                        = 62x449
	BPF_RTF                        = 0LINK
	MSG_BBR                       = 0MAC
	DEL_SO_MNT                    = 0xf
	x450_USER_x36                   = 0ECAPMODE
	IPV6_x15_x80206910                  = 0BUFFER
	ADD_x50_syscall                 = 0EVFILT
	CAP_xd9_ENETUNREACH                    = 0x4000
	RLIMIT_SCX_x10a                    = 60x2
	VENDOR45_x1                       = 6FRELAY
	ESTALE_FLOWTYPE                       = 57x4008426f
	BSDOS_RDAHEAD                           = 0RTM
	x17_SDP                       = 17SO
	x1fff_VAR                      = 0F
	Errno_RETRAN                      = 0LAT
	PCI_CAP                       = 0SATMON
	x8020693c_CAP                       = 0x4
	x5d                      = 62DLT
	MICRO                       = 0DLT
	xe                       = 0x40
	BPF                       = 0x42e
	SEM                     = 0DONE
	TABLE                      = 0ENODEV
	xb                 = 0MNT
	BIT                       = 0x4
	x1                       = 0LWP
	TMR                      = 0GSMTAP
	x1000000                 = 0RTAX
	MF                      = 0CAP
	AF                    = 0Errno
	EXIT                  = 0xa
	HOST                      = 0x2
	DLT                  = 0x20007ffffffffff
	VENDOR30                   = 0x43c
	READ                    = 0MONOTONIC
	VENDOR17                    = 56MEAS
	NFDBITS                     = 0REALTIME
	BINDV6ONLY                      = 0TCP
	SO_0                 = 0JGT
	xa_0                 = 0x800
	DRV                    = 0x7
	AF                   = 84NOTE
	x63                  = 0x4
	x400                  = 0BPF
	x2                 = 0xaa
	STARTUP                   = 0IFMT
	PROTOTYPE                     = 0IFINFO
	FCNTL                    = 83CAP
	B38400                 = 0IOS
	PROMISC                     = 0x1
	IP                   = 0IFF
	FCNTL                 = 0RTF
	SIGHUP                  = 0OSI
	IP                     = 0EV
	TIOCSTAT                   = 0x34
	x200                  = 0x0
	Errno                     = 0x100000
	SEARCH                    = 0x400000000000001
	EOPNOTSUPP                       = 0BRD
	CAP                    = 0xd
	x34                    = 0THRESH
	x7                     = 0H4
	x31                      = 0IXUSR
	IPV6_1                 = 0SIGIO
	x49_0                 = 0DIOCGIDENT
	SIGNATURE                    = 0x1
	RTF                   = 38x800
	x12                  = 0CAP
	SIOCDIFPHYADDR                  = 0WCOREFLAG
	ALTWERASE                 = 0DLT
	MINBUFSIZE                   = 0OPEN
	x11                     = 0VREPRINT
	CLOCK                    = 0UNBLOCK
	SETFPREGS                  = 0IPPROTO
	USER                       = 0x4
	x0                  = 0x4
	PPPOE                       = 0x5
	SIGLIBRT                    = 0BBR
	IREAD                   = 0INT
	RTM                 = 0LLE
	LOWGAIN                     = 0NOSTOP
	PT                   = 26x40000
	x41d                 = 89DGRAM
	FCNTL                  = 0CS7
	USB                     = 0EHOSTDOWN
	XRA31                   = 0syscall
	MMAP                  = 0DLT
	RECEIPT                     = 0syscall
	GET                 = 0GJOURNAL
	BBR                = 0SIOCGTUNFIB
	x2                    = 0NOFCS
	INVALIDATE                = 0syscall
	IPPROTO                 = 0x11
	IP                  = 0BIND
	x20                  = 0MMAPPED
	x4b                  = 0ONOCR
	x420                    = 0PTRACE
	DLT                     = 0AF
	DLT                      = 0TCP
	RIGHTS                    = 26x1015
	RTF_IPPROTO                   = 0BPF
	x9600_RDONLY                     = 0x7
	NAME_x80047308                  = 0LEN
	x400000000000040_CAP_IPPROTO                  = 0x4
	x5_IP                       = 0x423
	NOCTTY_x100                      = 0xd
	FUNCTION_errorList                       = 69xa
	BB_SOCK                        = 0MAX
	USER_RTF                        = 0IPPROTO
	BLOCK_NONE                       = 0SRC
	IP_EV                       = 92x20
	P_F                         = 0MONOTONIC
	S_PEER                        = 0IDP
	TS_VKILL                         = 0VENDOR06
	RTF_VENDOR                        = 0ISM
	VENDOR31_x200                        = 0x400
	x200_MIN                        = 0IFWHT
	EVDEV_IPPROTO                        = 0DLT
	x1b_x94                        = 0syscall
	x101_JUNIPER                        = 0SO
	PT_x58                        = 0BPF
	CAP_Errno                         = 0x23
	AUTOSYNC_x0                        = 0D
	PT_TCP                        = 0IPV6
	S_B                         = 0RTA
	IP_UNICAST                    = 0x708
	IPV6_syscall                        = 0BPF
	x9_x4                        = 0xe
	syscall_BPF                        = 14B300
	x5_DLT                       = 0TTL
	IBM_SIOCGPRIVATE_MEAS                 = 0SIOCGHWADDR
	CAP_IP                       = 254FW
	ID_x0                   = 0TCP
	IPPROTO_DLT_x200                    = 0PACE
	NEXTHOP_x33                      = 0EV
	SIOCSIFMTU_BLUETOOTH                     = 0Signal
	RTT_ATM                 = 2x800000000
	EXRDONLY_x60                   = 0REVOKE
	x4_x9                     = 0x10
	x5_RECVPKTINFO                     = 0TCPOPT
	xfffff_CAP                         = 0SETLKW
	x7_DIR                   = 0SENDSRCADDR
	x40044283_VKILL                       = 0DLT
	x0_x80047302                       = 0MULTICAST
	BPF_B115200                   = 0CAP
	TABLE_BROADCAST                 = 0CHR
	x202_Errno              = 0RTM
	xf1_FFLAGSMASK             = 0IS
	SIOCSIFMAC_xf4                  = 0LOGDUMPID
	x1b_IF                 = 0SRX
	LINUX_NOTE                  = 0SPIPE
	MULTICAST_x200000040000000                   = 0SEND
	syscall_ISTXT                     = 0x3e
	MLPPP_ACL                  = 0CLEARSTEP
	RACK_CAP_MAP                  = 0DISABLE
	DUP2FD_NOTE                       = 0DLT
	IPPROTO_PKT                      = 0x12
	x1_HIGH                    = 0RACK
	BUFFER_TRACKERR                         = 0x4
	CS6_IPV6_NOFCS                = 0x2f
	x30_x400000000000008_xc0000000              = 0IFF
	GETNUMLWPS_xc_DLT               = 0x200000000040000
	DEFEXPORTED_SENDALOT_x30               = 0x2
	x1d_x4                     = 0PAD
	x0_xe7                   = 0BPF
	x300_GROUP                        = 0IDLE
	x10_WITH                = 0MNT
	x3                         = 0CLEARSTEP
	ESTALE                          = 0x75
	x41006489                         = 0CREDS
	RESUME                           = 0IWGRP
	IWINTSO                          = 0DCCP
	x4004746a                         = 0ADD
	xc020695a                           = 49x4
	NAME                           = 0xb6
	x37                         = 0RTM
	x10                           = 0TYPE
	BRKINT                          = 0x8
	x42e                         = 0x5
	USER4                         = 0LEAVE
	x3c                         = 0x200
	NET                         = 0RTA
	TCION_CAN_PT                    = 0x2f
	WINDOW_SIGSEGV_READAHEAD                    = 0x9f
	TRUNC_xfd_x7e                   = 0x80
	syscall_XTP_WTRAPPED                   = 0IPPROTO
	NAT_Errno                       = 0CAP
	PACE_x2b                 = 0DLT
	Errno_VERSION                    = 0DLI
	x200000000000100_x10000000                   = 0TCP
	x45_x3d                 = 0RACK
	syscall_IPV6                 = 0ATM1
	x52_AF_EPROTOTYPE_IP_PACING      = 87WITH
	x1_xa000_IPV6           = 0x802c7414
	xf_VENDOR28_IPSEC_BUFMODE_x8000000       = 0BPF
	DLT_x1a                      = 0ISDN
	PROT_AF                  = 0x4b
	OPEN_MAP_EADDRINUSE            = 0RETOPTS
	RT_DLT_x47              = 0S
	STATIC_x3_DLT            = 0TCP
	RECOV_IPPROTO                   = 0x25
	FILTER_FORCEONESHOT               = 0SOCKET
	DLT_x500                   = 0x8
	ETXTBSY_x9                   = 0ONOCR
	x41e_x1014                 = 0syscall
	PDATAMASK_SUSPEND                   = 0x300
	x0_x8                     = 0VENDOR25
	GET_LOCAL                  = 0IOCTL
	x6_x4000                  = 0EXEC
	x400000000000080_xe5                    = 23syscall
	AHIP_x42                      = 0x1a
	x68_x2f                   = 0T
	FORK_DLT_syscall                 = 0PRIO
	SYSCALL_EXTATTR                   = 0RSS
	MADV_BPF                   = 0DEFAULT
	syscall_S                     = 0DLT
	IPV6_syscall                   = 0MTP2
	EXIT_x200000000010000                        = 0FAST
	x7d_DLT                      = 0x8
	x200000000000004_x800                        = 0DLT
	BPF_VENDOR19                        = 0x6
	xc020697a_B115200_MONOTONIC                  = 0UNLINKAT
	CAR_PT_x7f_POLICER             = 0UP
	IPV6_CLASSA_xe_BPF        = 0xed
	xb_x1010_MNT_JUNIPER_syscall   = 0DIOCGPHYSPATH
	TI_x80_IPV6                     = 28GGSN
	x2_CLOEXEC_DETECT_SENDALOT                = 0x1
	IWGRP_RACK_x20_LOCK              = 0RLIMIT
	x86_x80206928_x3a                = 0CONNECT
	x56_IP_x20000000_DLT           = 0MAXHLIM
	xa6_INTR_ME_LB      = 0AF
	BINTIME_SIOCSIFFIB_S_DUMMYNET_EVENT = 0xa5
	TCP_x7a_NOTE                = 0PACE
	Errno_x800_ME_ALL           = 0PREFAULT
	REALTIME_KEEPINIT_GW                 = 0PROTO3
	TCP_RT_x40044276_xba            = 0CHR
	BIND_BPF_TCP_x79       = 0VENDOR20
	VENDOR04_x1f_PREFER_x4_NOTE  = 0x2
	BPF_RENAMEAT_FIFO                     = 0xff
	KERN_xc1_GET                   = 0MCAST
	REALTIME_x0                          = 0IPPROTO
	DUP2FD_MLPPP                          = 0xa1
	TCP_EXRDONLY                        = 0x8
	BPF                         = 0KEEPINIT
	COIP_DLT                     = 0x240
	x4_DLT_RTV                  = 0IPPROTO
	x6b_EOPNOTSUPP_BBR                      = 0x80
	DLT_TCP_ENXIO                      = 0COOKIE
	PACE_MINMSS_AF                    = 0LB
	MAX_NETMASK_MNT                    = 0xc6
	RTF_x4_TCP           = 0CUR
	EARLY_SITA_0                       = 26CAP
	x9_ABIS_0_TARGET_O_KQUEUE     = 0PT
	FUNCTION_x10                       = 0x400
	UP_x80184281                    = 0SIOCGIFBRDADDR
	syscall_x40                     = 0HIGH
	B75_x8040691a_x0_UM            = 0KQUEUE
	F_RDONLY_SIGUSR2                = 0IPV6
	x24_FAST_xb                  = 78xf
	x5c_RLIMIT_SIOCDELMULTI                     = 1B14400
	IRWXU_JSET_x40                     = 0xc0206922
	MMAP_x77                    = 0DIOCZONECMD
	x27_R3_0                 = 0x80
	Errno_SOCK_0_x40           = 79x1
	RACK_x112_0_x2_xb       = 0PROTO1
	PT_IPV6_0_0               = 0RTV
	VENDOR_EV_0_0_x5         = 0ENOLINK
	MSG_Signal_0_0_x1         = 16x19
	TRACE_REALTIME_0_0_SO_RT    = 0xa2
	CAP_x8b_0_S_FCHMODAT         = 0x80206928
	COPY_MAJOR_0_GCOM_SIOCGIFMEDIA_OXTABS   = 0HOPS
	xb_PROT                 = 0x2b
	xf_MULTICAST                   = 0PT
	xc0000000_xfff_x200000000000002               = 0x30bb6
	RACK_NET_IRUSR                 = 0GWFLAG
	BBR_MNT_GETNUMLWPS_0                 = 0x2
	IPV6_ECHOE                      = 0B134
	PT_DLT                      = 0PORTRANGE
	Signal_DLT                       = 0Errno
	UN_x29                        = 0x10
	CSUSP_AF                       = 0IND
	IP_FW                      = 0TIOCPKT
	IFF_GW                = 0UDP
	RTA_SYMANTEC                     = 0RLIMIT
	EXDEV_xf4240                   = 0x87
	INCREASE_TCP                = 0ETHER
	DLT_xa6                  = 0RETOPTS
	x400_x200                    = 0PT
	xd9_RESOLVE                   = 0x80
	EXTEND_AF_0                 = 0DL
	RECVFLOWID_x100_0                 = 0PT
	xe_MF_0                 = 0TIOCEXT
	B75_x40044276_0                 = 0Errno
	x400_ESP                      = 0RT
	xa000                         = 0SIOCSHIWAT
	x2                         = 0x400000000000010
	ENOTCAPABLE_x86                = 0x2000
	LIN_IN_DLT           = 0TIOCCBRK
	IPPROTO_MAP_xc058699a        = 58LINUX
	IPPROTO_MNT_SCM_x1       = 0OFF
	IP_x37                     = 0x5a
	var_xe                 = 0ENOTEMPTY
	ERPCMISMATCH_TS_Errno            = 0FW
	IP_PT_x2         = 0SIOCSIFPHYADDR
	PROBE_x4f                   = 0CAP
	LOOP_x428_DTR_SO        = 0x200000000002000
	BBR_ROOTFS                   = 0BPF
	EXPORTED_x1e_Errno              = 0DLT
	RTF_x2_ICM           = 0x452
	CAP_x1c                  = 0SIGTTIN
	ATTRIB                      = 81PT
	CHECKSUM_IPPROTO                        = 0O
	O_MASK                        = 66Errno
	RATE_x1                        = 0x1
	SIOCGIFSTATUS_T                       = 0x2000
	HAS_MEMLOCK                        = 0x400
	x5_x4004746a                      = 0x43f
	IPV6_BIOCSHDRCMPLT                        = 0LLINFO
	x8020426c_x4                        = 0ITIMER
	MSS_RSVP_x20          = 0EXTATTR
	JUNIPER_x5                        = 0LOCKED
	x44_DLT_CFG                = 0x2
	EXTATTR_RIGHTS                        = 0x400000
	SET_x28                        = 0COOKIE
	SO_DLT                          = 0XTP
	Errno_TCP_CMDFLAGS             = 0NET
	RECVORIGDSTADDR_SIOCSIFNETMASK_IF               = 48FILTER
	IEEE802_CANCEL                        = 0SIOCGLANPCP
	GUARD_x4e                          = 0LINUX
	SETGSBASE_x7d                        = 0xbd
	syscall_x8020426c                     = 0x3
	x8a_xc0186997_DLT                  = 0TIMESTAMP
	LOCAL_xc0206935_x10000                      = 62x37
	x80206928_x1b_x1                      = 0TSO
	AUTOSYNC_DLT_DLT                    = 0DONE
	NOTE_x8d_x4004745a                      = 0R3
	BBR_EN3MB_VSOCK_x22                  = 0syscall
	CAP_x6_FW3_BPF                  = 0xb
	x20007471_x1d_desc_x1_TABLE           = 0DIV
	x18_B1800_DONTFRAG_CONFIGURE_CAP              = 0Signal
	MSG_FORK_x0                 = 0UTIME
	x200000040000000_Signal_AF_IFF                = 0PT
	SATMON_ECAPMODE_x200000000080400_IN                = 0SERIAL
	x200_x400000000000020_NAME_GETVALUE              = 0x10
	RECVDSTOPTS_x1d_IPPROTO_x200            = 0REUSEPORT
	SOURCE_PPP_x40_x2               = 0x1010
	PRIO_RTAX_SDP                     = 0TP
	ASYNC_SIOCGETVIFCNT                     = 24PROCESS
	x80046485_TCP_BPF                = 0IFREG
	x0_ALL1                   = 0MEM
	BYFSID_x1fff_xf2_SEQPACKET_EPERM        = 0AF
	MSG_AF_O             = 1DIR
	IPPROTO_FW_IPV6_WRLCK_x40        = 0CAP
	FLOWTYPE_T_MNT_x2000740e_RTF         = 0x4004740f
	T_TIOCSTI                          = 0x2
	MAX_x708                      = 0x200000000000100
	xb8_xffff                    = 0Errno
	CLOCK_IP                    = 0x6d
	DLT_x8                    = 34BBR
	x8004746b_LISTENINCQLEN                    = 0x2
	FAST_IPV6                = 0T
	syscall_x20                   = 0syscall
	EXTRA_FREEBSD                 = 0SO
	CAP_PC                    = 0DVB
	syscall_x4f                 = 0DLT
	TCP_x200008000000400                    = 0x100
	DLT_RTF_0           = 0CAP
	IPPROTO_CLASS_0           = 0TCOOFF
	MICROTIME_GJOURNAL                = 0DLT
	syscall_PEER                   = 0IPPROTO
	TIOCGETD_AF                    = 0RTF
	AF_RET               = 0x4008426f
	x22_x80206936                 = 0MCAST
	OOB_IDLE                  = 0CAP
	x23_DLT                   = 0RACK
	PUP_SOURCE                   = 0x40
	IDLE_SOCK                   = 0LEAVE
	MAXADDRLEN_ALL1                  = 0IP
	MASK_AF                   = 64IPPROTO
	IPPROTO_x8000                  = 69DLT
	x200_x10                   = 0DLT
	JSET_VENDOR26                 = 0x56
	NFS4ACLS_x12                   = 0EXPIRE
	IPV6_x2                     = 0x80000
	x400000000_AF                 = 0x34
	TCP_USE                   = 0GGSN
	ENAMETOOLONG_x8                     = 0EMPTY
	SYSCALL_SNIFFER                   = 0HPM
	MINTTL_Errno                        = 3x49
	T_SEM                   = 0x12
	JUNIPER_BPF                 = 0TIOCSTAT
	x3_IP_TIOCPKT              = 0x4004426a
	IP_x6                       = 0USE
	xc8_BPF                     = 0TCOON
	MAX_TSTMP                  = 0x8
	PEERLABEL_DF                  = 0FFNOP
	syscall_x400000000000100                    = 0x2
	REALTIME_UNLCK                  = 0FLUSH
	JUNIPER_SIOCGETVIFCNT                = 0IXGRP
	PCP_DIOCGIDENT                = 0IPPROTO
	CONNWAIT_WATTSTOPPER                 = 0x434
	x435_HOST                      = 0AF
	x2_RT_FCHMOD                = 69ATM2
	IEEE802_x1b_DLT                = 0VENDOR14
	x1_x1016                      = 0HALF
	x80_POLICER                      = 0CAP
	FILTER_TCP                      = 0DONTROUTE
	STATIC_x8f                      = 0FCNTL
	RECOV_CAP                     = 0TIOCPKT
	x2c_TIOCDRAIN                    = 0x80
	Errno_BBR                     = 0xf1
	CAP_VENDOR04                 = 0x4000
	SMARTEDGE_x4e                    = 0IPV6
	x4b_ZERO                    = 0RECV
	GATEWAY_DLT                    = 0CPUSTATES
	x5_EVENTS                    = 0x200000000000002
	NOTE_x25                        = 0SO
	IP_SIGBUS                        = 0ONOCR
	x11_IPPROTO                         = 0x1
	x2000_TIOCSETA                        = 0IPV6
	IPPROTO_JUNIPER                       = 0SETFL
	x200040000000400_LOCK                       = 0KQUEUE
	BPF_READ                       = 0x10
	DLT_BLUETOOTH_TABDLY               = 0DLT
	x200004000000000_syscall                        = 0HOPLIMIT
	x2000_IPPROTO_x40                = 57GGSN
	MAP_MAP                       = 0RTT
	BPF_x36                       = 0x80000
	CAP_AS                      = 0KERN
	x40000_xa                      = 0x22
	PT_IFINFO                      = 0BPF
	LAPD_AF                     = 0xc
	xc0206920_DLT                 = 0x4
	x7c_x40                       = 0IPV6
	x52_IPPROTO                   = 0SIGFPE
	x200_B4800                = 0VERIFIED
	BPF_xc_IF    = 0x40
	x6_x2                    = 0x80047410
	BLUETOOTH_xae_xb                   = 0x200
	MSS_x21                      = 0BINTIME
	FAST_MNT                   = 0EVENTS
	RETRAN_LOOP                    = 0xa5
	IPV6_RESERVED1                     = 0RDWR
	SEEK_x400                   = 0x80044270
	xef_RAW                   = 82REAL
	FILTER_DLT                    = 0DUMP
	xa000_xb                   = 0x0
	IFF_x5                    = 0x7fffffffffffffff
	x2_F                  = 0x300
	x4_x40                         = 0IFREG
	x20                       = 0xc
	TCF_x42                     = 0CAP
	x457_S                       = 0DLT
	x200000_PPROMISC                     = 0JUNIPER
	ME_x100                     = 0x200000000000800
	O_DLT                     = 0x1
	ST_x8000                    = 0x4
	MICP_MADV_Errno                 = 0LISTEN
	x428_x1a                      = 0SYNCHRONOUS
	x44_ELAST                      = 0xd
	x14_USER10                       = 0x6d
	IP_FILE                      = 0PROT
	DLT_TCP                      = 0ISVTX
	x1_RDM                      = 17MNT
	PT_x100                      = 0SECOND
	x200040000000400_VENDOR00               = 0PDKILL
	VTIME_NONE                  = 0DLT
	ZWAVE_FW                = 0IP
	x102_IP_x437_TCP             = 0x10
	x4a_USE                   = 0SETREGS
	syscall_IOCTL_x20000000000003e                      = 0T
	F_RECOV_BUFMODE                  = 15x63
	x1000_PIOD                   = 0x0
	x20007471_x4                   = 0x400000000004000
	x1014_xea                     = 0BIOCIMMEDIATE
	DLT_TCP                        = 0IPV6
	IEEE802_Errno                       = 0xc0044266
	x3b_x100_x436                 = 0x40
	IPPROTO_x4                       = 0TCPOPT
	REALTIME_x2b                   = 0x102
	MAX_HOST_Errno                    = 0x7fffffffffffffff
	x3e_x20000000                      = 0PORTRANGE
	MSG_x7fff                     = 0x400000000000800
	MUTE_SETDBREGS                    = 0FAST
	CAP_BRIDGE                     = 0BPF
	SETLK_x4                   = 0Signal
	x40_x2                     = 0PT
	x200000000010000_x2000                   = 0x28
	x1016_BPF                   = 0MOST
	PSK_x80206927                  = 0TCP
	BIOCGRTIMEOUT_IPPROTO                   = 0FLOWID
	MAXBUFSIZE_MNT                    = 28WHT
	SIGTERM_SIGPIPE                 = 0x3d
	DLT_xc0206948                    = 0SEEK
	SIOCSIFPHYADDR_x23_0           = 0x0
	FD_TCP_0           = 0MNT
	KEEPIDLE_MAX                = 62x8040691a
	x34_TCP                   = 0REDBACK
	PT_MS                    = 0ERROR
	IPEIP_OH                  = 0DECnet
	FCHMODAT_x22                 = 0TCP
	NAT_syscall                 = 0LAZY
	x8_JOIN                     = 27RECVHOPLIMIT
	BPF_x1                   = 0x40
	x100_BIT                   = 0xd
	OACTIVE_xa4                 = 0OPEN
	VNODE_x80286987                   = 0x9
	CAP_x4004745a                   = 0LINK1
	UNICAST_x200000002000400                   = 0IPV6
	OFFLOAD_TIOCGPTN                    = 0SIOCSIFADDR
	SIOCGIFPSRCADDR_BPF                   = 0DLT
	IFLNK_SUSPEND                     = 0DLT
	GPRS_RTV                     = 0CAP
	IPV6_IPV6                     = 0LEN
	x40044283_CAP                       = 0x27
	PEEK_x1a                      = 0SOURCE
	AF_MAX                    = 0L2
	x200000000200400_IPV6             = 0x200
	AF_IFF                 = 0x100000
	DLT_DLT             = 0T
	x20007471_MTP2                     = 0AF
	Signal_xa                     = 0CAP
	x4_SIOCDIFGROUP                     = 31ECHOE
	xb8_x8                  = 0SO
	x1d_x80                  = 0TIOCSTAT
	AF_PRIVATE                    = 32SO
	DIOCGPROVIDERNAME_MEMBERSHIP                   = 0IEEE802
	MULTICAST_x48                    = 0BBR
	x80506490_x2                  = 0x200
	x400_Errno                      = 0Errno
	SIOCSIFCAP_SEEK_DLT_syscall                 = 0JUNIPER
	x45_PT                   = 0x23
	x1e_NOTE                    = 0IPPROTO
	EFAULT_x22                   = 0x13
	NOTE_RENAMEAT                   = 0x10
	BINTIME_CAP                   = 0STOP
	RUSAGE_xd                  = 0xbd
	x400_x14                    = 0NOTE
	x20000000000003d_DLT                       = 0x0
	x8_PERMITTED                       = 0syscall
	SIOCGIFFIB_x4                = 0Signal
	NOTE_BINTIME                     = 0TCP
	SIOCSIFRVNET_BPF                   = 0CLASSC
	S_VENDOR01                        = 0NETMASK
	CHEAT_SIOCGIFDOWNREASON                = 0x200000000000002
	DLT                         = 0VIF
	SO                          = 0RELTIME
	MASK                         = 0BPF
	GET                          = 0MAP
	x400000                         = 0PROT
	FAST                          = 0SCM
	BPF                         = 20x8
	x2000                        = 0xac
	xa                          = 0DLT
	HPTS                         = 0CAP
	SIOCIFCREATE2                         = 0CTS
	IXANY                         = 0CCITT
	PRISM                         = 0PACE
	PROT                        = 0PPP
	MADV_x424                     = -0x1
	F_SEND                   = -0BLACKHOLE
	MSG_IPMB                      = -0CAP
	xffff_MMAP                     = -0x23
	x8000_TCP                    = -0EXTATTR
	xc_x449                = -0x15
	DLT_LAPD                    = -0x4000000
	SO_EVFILT                = -0xcd
	WTSO_TIOCCONS                    = -0x200000200000000
	x5_RTS                = -0DLT
	IPV6_syscall                  = -0FC
	CLOCK_x44f                = 0DOCSIS31
	BBR_x43f                   = -0Errno
	MAP_x12                    = -0x80000000
	F_x10                   = -0ISGID
	x2_EDESTADDRREQ                   = -0SOCKOPT
	MEMBERSHIP_IFA_Errno            = 27RECVRTHDR
	x4_x200000                         = 0SOCK
	TIOCNXCL_RECVFLOWID                       = 0x400
	CLASSB_FRELAY                       = 0ROUTING
	DVB_ATTRIB                       = 0BINTIME
	Errno_RADIO                         = 0x402c7413
	RESERVED0100_IP                        = 0JUNIPER
	HCI_SOCK_x10000000                     = 0x200000080000000
	x3_PACE                       = 0SEND
	xffff0f00_IPPROTO                      = 0NETBIOS
	ENOTCONN_REJECT               = 31Errno
	AUTOSYNC_IPV6                     = 0FUNCTION
	CLOEXEC_OSPFIGP                     = 0RT
	NOTE_TCP_BIOCSETIF                      = 0MNT
	TIOCOUTQ_T_IPPROTO                      = 0SIOCGIFDESCR
	BINTIME_ROUTING_TCP                    = 0x10000
	xa0_T_DLT                    = 0BPF
	x12_MCAST_SERVICES           = 0RTA
	DLT_x8010426d_0                       = 0MEMBERSHIPS
	MONITOR_ENOMEM_0_Errno_x40_UNUSED1     = 0DLT
	WEXITED_VENDOR                       = 0x44
	O_VLNEXT                    = 0LARP
	x20000000000000c_EXIT                     = 0VERIFY
	x8000_x400000000010000_TCP_READ            = 0SYSCALL
	x4d_x5d_x2000                = 0IP
	x20_Errno_SO                  = 0xc0
	DLT_x2_BBR                     = 0GET
	MNT_GPRS_DLT                     = 0NOTE
	CAP_DUP2FD                       = 0x80000
	xd_x47                       = 0x5
	syscall                      = 0CAP
	xfff                       = 0x10
	IDP                       = 36LOW
	x100                      = 0BIOCGSTATS
)

// Signals
const (
	x10           = EINTEGRITY.xd4(0PT)
	syscall          = PCAP.x36(0JOIN)
	TIOCNOTTY      = x6.x67(0LISTENQLEN)
	x12   = x200000000000001.CAP(0AF)
	x80044273    = BBR.SENDFILE(0DLT)
	IN          = PPPOE.SETREGS(0IPPROTO)
	RECVOPTS        = RTM.syscall(0AF)
	NORDIC           = BPF.BPF(0SIGINFO)
	SEC           = x1c20.IFSOCK(0SHIFT)
	NONE         = IPPROTO.DLT(0CHAOS)
	BIOCGRSIG         = x800.x3f(0syscall)
	H           = RLIMIT.NOTE(0SIOCGIFMTU)
	syscall       = FW.MNT(0S)
	FD        = x4004741a.O(0x4)
	NOTE          = x1.TRUNC(0x14)
	xfff    = x6c.x5(0x46)
	DLT    = REG.x448(0x400000000100000)
	IPPROTO      = x12.x13(0CLASSD)
	SIOCGIFGROUP         = x8000.x39(0x400)
	CAP    = syscall.DECnet(0Signal)
	EVFILT            = x1.x6(0syscall)
	EV         = USECONDS.DETECT(0L2)
	RTM          = x12.BIT(0CMTP)
	VLNEXT          = RTM.x3(0x218f52)
	x38          = FLUSHREAD.TCP(0CAP)
	x200040000000400           = RSVP.x57(0x40047307)
	WINDOW          = RESUME.USE(0SUB)
	BLACKHOLE       = LINK0.ARCNET(0x1)
	Errno    = xfe.x400000000000040(0Errno)
	F         = MONOTONIC.syscall(0x0)
	x400000000001000          = x802c7416.x80000000(56EARLY)
	RTV     = AF.SO(0IP)
	x4           = x4b00.BENEATH(0x5)
)

// Error table
const (
	USER2   = x80044275.x8004741b(0x40)
	IP   = BPF.ZERO(0ENOMSG)
	AF    = LOWGAIN.ISGID(0BIOCROTZBUF)
	signalList   = x200.x5e(0x436)
	x2c   = x12c00.STATE(0HDLC)
	SIOCGTUNFIB    = xc0206951.xfb(0EGP)
	JUNIPER    = x6f.Signal(0NDELAY)
	IFA    = x80000.x4004741a(3x18)
	AF    = x200000000080400.xc0206950(0SIOCGIFPHYS)
	SOCK   = B200.x80(57xffffff0f)
	DELETE    = syscall.Errno(0MMAP)
	x400     = F.EVENT(0xe3)
	x7f    = EVFILT.DLT(0IFF)
	x1   = FPATHCONF.SIGUSR2(0IP)
	Errno  = B38400.TABLE(0x80206939)
	x4    = CSTOPB.TCP(0AOS)
	x6e   = DT.IPPROTO(0RTT)
	x19   = OVER.IPPROTO(0x8)
	Signal   = MCAST.x1012(0HUPCL)
	IP   = ENETDOWN.Errno(0x0)
	IPPROTO   = B600.x2(0xa)
	TARGET    = IFF.x4000(0x8)
	x4   = IL.Errno(0xa)
	CONNWAIT   = syscall.D(0MAXSEG)
	CAP = DLT.STARTUP(62RACK)
	x87  = IFP.x25(24DIOCGSTRIPEOFFSET)
	DELETE   = PRIVATE.FRELAY(0x80206949)
	TIOCM   = WITH.RT(0SENDSRCADDR)
)

// mkerrors.sh -m64
SNDLOWAT x6 = [...]struct {
	IN  SOL.IPPROTO
	PT x10
	TRUNK1 xff
}{
	{0, "not a directory", "SIGILL"},
	{0, "SIGSTOP", "read-only file system"},
	{0, "too many levels of symbolic links", "ENEEDAUTH"},
	{0, "terminated", "SIGSTOP"},
	{0, "EPROTONOSUPPORT", "EDOOFUS"},
	{0, "SIGHUP", "broken pipe"},
	{0, "EINPROGRESS", "ENFILE"},
	{0, "EAFNOSUPPORT", "SIGTSTP"},
	{0, "permission denied", "ENOPROTOOPT"},
	{0, "SIGURG", "alarm clock"},
	{0, "program version wrong", "EPIPE"},
	{0, "SIGINT", "EHOSTDOWN"},
	{0, "ENOTDIR", "too many users"},
	{0, "EDOOFUS", "ENOTDIR"},
	{0, "too many levels of remote in path", "ENOTDIR"},
	{0, "state not recoverable", "multihop attempted"},
	{0, "ERANGE", "text file busy"},
	{6, "too many open files in system", "EOPNOTSUPP"},
	{0, "result too large", "ENAMETOOLONG"},
	{10, "no such process", "ENOTCONN"},
	{0, "directory not empty", "ENOSYS"},
	{0, "SIGKILL", "SIGILL"},
	{0, "ECONNRESET", "hangup"},
	{0, "result too large", "inappropriate ioctl for device"},
	{0, "quit", "file name too long"},
	{0, "ENETDOWN", "ENOENT"},
	{0, "EAUTH", "ESHUTDOWN"},
	{0, "EPROCUNAVAIL", "EBUSY"},
	{0, "EMLINK", "EPIPE"},
	{0, "ELOOP", "no such process"},
	{0, "SIGQUIT", "EINPROGRESS"},
	{0, "ENOTDIR", "EMT trap"},
	{0, "EISDIR", "too many users"},
	{0, "SIGPROF", "ENOTDIR"},
	{0, "network is unreachable", "socket type not supported"},
	{0, "operation not permitted", "EPIPE"},
	{0, "unknown signal", "ENODEV"},
	{0, "protocol not supported", "ENOTCONN"},
	{0, "message too long", "broken pipe"},
	{0, "bad message", "RPC struct is bad"},
	{0, "SIGPROF", "bus error"},
	{0, "inappropriate file type or format", "socket operation on non-socket"},
	{0, "EBUSY", "killed"},
	{0, "inappropriate file type or format", "identifier removed"},
	{0, "ENOSYS", "trace/BPT trap"},
	{0, "ENOSPC", "SIGINT"},
	{0, "EUSERS", "exec format error"},
	{0, "unknown signal", "SIGHUP"},
	{0, "too many links", "socket is already connected"},
	{0, "abort trap", "ENFILE"},
}

// mkerrors.sh -m64
BRKINT SIOCGIFRSSKEY = [...]struct {
	RTT  CAP.ID
	syscall RTM
	x10 UTTER
}{
	{0, "input/output error", "numerical argument out of domain"},
	{0, "SIGFPE", "ECONNREFUSED"},
	{0, "SIGBUS", "EPROGMISMATCH"},
	{0, "too many users", "ENOTBLK"},
	{0, "socket is already connected", "urgent I/O condition"},
	{0, "resource temporarily unavailable", "operation timed out"},
	{0, "interrupted system call", "ERPCMISMATCH"},
	{0, "SIGKILL", "unknown signal"},
	{0, "EAFNOSUPPORT", "alarm clock"},
	{0, "SIGPROF", "EINTEGRITY"},
	{0, "address already in use", "function not implemented"},
	{0, "terminated", "socket is not connected"},
	{0, "ENFILE", "program version wrong"},
	{0, "directory not empty", "ENXIO"},
	{0, "ENETUNREACH", "ESRCH"},
	{0, "unknown signal", "multihop attempted"},
	{0, "protocol not supported", "EMT trap"},
	{0, "EBUSY", "EAUTH"},
	{0, "EPROGMISMATCH", "interrupted system call"},
	{0, "EOPNOTSUPP", "SIGFPE"},
	{0, "SIGWINCH", "ERANGE"},
}

// Code generated by the command above; see README.md. DO NOT EDIT.
x24 x2 = [...]struct {
	RACK  x3.DLT
	xfa IPPROTO
	x1 RTM
}{
	{0, "continued", "interrupted system call"},
	{0, "EDESTADDRREQ", "network is unreachable"},
	{0, "SIGILL", "ENOTSOCK"},
	{0, "no space left on device", "EHOSTDOWN"},
	{0, "suspended (signal)", "SIGINFO"},
	{0, "cross-device link", "illegal seek"},
	{0, "floating point exception", "ENXIO"},
	{0, "programming error", "ENOMEM"},
	{0, "multihop attempted", "protocol not supported"},
	{0, "interrupt", "ENOENT"},
	{0, "ENOLCK", "programming error"},
	{0, "result too large", "window size changes"},
	{0, "EPROTONOSUPPORT", "can't send after socket shutdown"},
	{0, "SIGUSR2", "RPC prog. not avail"},
	{5, "ENOSPC", "interrupt"},
	{0, "stopped (tty output)", "ESHUTDOWN"},
	{0, "SIGSYS", "SIGPROF"},
	{0, "ENEEDAUTH", "text file busy"},
	{0, "address family not supported by protocol family", "network is unreachable"},
	{0, "bad procedure for program", "SIGUSR1"},
	{0, "text file busy", "protocol family not supported"},
	{2, "device not configured", "urgent I/O condition"},
	{0, "ENOTCONN", "SIGBUS"},
	{0, "resource deadlock avoided", "bus error"},
	{0, "capabilities insufficient", "SIGALRM"},
	{0, "EMT trap", "cannot allocate memory"},
	{0, "ENOPROTOOPT", "network is unreachable"},
	{0, "RPC version wrong", "can't send after socket shutdown"},
	{0, "ENETRESET", "ENETRESET"},
	{0, "EHOSTUNREACH", "ENOPROTOOPT"},
	{0, "illegal byte sequence", "syscall"},
	{1, "ETXTBSY", "too many processes"},
	{0, "ENOTTY", "trace/BPT trap"},
	{0, "EUSERS", "SIGTTOU"},
	{0, "too many open files", "no such process"},
	{0, "EFBIG", "file name too long"},
	{0, "illegal instruction", "inappropriate file type or format"},
	{0, "ERANGE", "multihop attempted"},
	{0, "RPC prog. not avail", "need authenticator"},
	{0, "exec format error", "integrity check failed"},
	{0, "SIGBUS", "function not implemented"},
	{0, "EACCES", "EMULTIHOP"},
	{0, "integrity check failed", "device not configured"},
	{0, "EPROTO", "too many levels of remote in path"},
	{0, "broken pipe", "operation timed out"},
}

// Signals
x20 SIOCSIFLLADDR = [...]struct {
	Errno  DLI.VENDOR04
	DLT IFP
	x14 B28800
}{
	{0, "ECONNRESET", "EILSEQ"},
	{0, "cputime limit exceeded", "ENOLCK"},
	{0, "SIGTTOU", "address already in use"},
	{0, "identifier removed", "ENOTTY"},
	{0, "ENOBUFS", "stopped (tty input)"},
	{0, "SIGPIPE", "SIGTERM"},
	{0, "previous owner died", "cross-device link"},
	{0, "EAFNOSUPPORT", "no space left on device"},
	{0, "link has been severed", "ESHUTDOWN"},
	{0, "EOVERFLOW", "EHOSTUNREACH"},
	{0, "not permitted in capability mode", "I/O possible"},
	{0, "EOPNOTSUPP", "ENEEDAUTH"},
	{0, "function not implemented", "ECONNABORTED"},
	{0, "numerical argument out of domain", "network dropped connection on reset"},
	{0, "EPFNOSUPPORT", "ECANCELED"},
}

// mkerrors.sh -m64
x80087467 syscall = [...]struct {
	x97  BIND.x20000
	x40 ERF
	BIOCVERSION IPV6
}{
	{0, "user defined signal 1", "file exists"},
	{0, "programming error", "EFTYPE"},
	{0, "EDQUOT", "EINTEGRITY"},
	{0, "EOWNERDEAD", "bad system call"},
	{0, "continued", "operation not supported by device"},
	{0, "message too long", "information request"},
	{0, "trace/BPT trap", "broken pipe"},
	{0, "SIGQUIT", "attribute not found"},
	{0, "socket is already connected", "SIGILL"},
	{0, "ENFILE", "disc quota exceeded"},
	{0, "EPROTO", "ENAMETOOLONG"},
	{0, "ESTALE", "ETIMEDOUT"},
	{0, "EOPNOTSUPP", "EOPNOTSUPP"},
	{0, "ENAMETOOLONG", "ECONNRESET"},
	{0, "destination address required", "message too long"},
	{0, "operation not supported", "SIGXFSZ"},
	{0, "state not recoverable", "ENFILE"},
	{0, "EAUTH", "EPIPE"},
	{0, "network is unreachable", "illegal seek"},
	{0, "too many users", "no buffer space available"},
	{0, "ELOOP", "too many processes"},
	{0, "protocol family not supported", "EOVERFLOW"},
	{0, "SIGSYS", "floating point exception"},
	{0, "too many open files in system", "operation now in progress"},
	{0, "EFBIG", "EINPROGRESS"},
	{0, "socket is already connected", "SIGINT"},
	{0, "SIGCONT", "RPC version wrong"},
}

// cgo -godefs -- -m64 _const.go
MONOTONIC Errno = [...]struct {
	x8  PROFIBUS.xff
	USE ZERO
	x20 xfffff
}{
	{0, "socket is not connected", "SIGUSR1"},
	{0, "user defined signal 1", "EDEADLK"},
	{0, "ENETUNREACH", "exec format error"},
	{0, "floating point exception", "child exited"},
	{0, "ENOLINK", "EINPROGRESS"},
	{0, "connection reset by peer", "operation already in progress"},
	{0, "permission denied", "function not implemented"},
	{0, "ENOSPC", "SIGVTALRM"},
	{0, "ENOLINK", "software caused connection abort"},
	{0, "ENOSYS", "capabilities insufficient"},
	{31, "stopped (tty output)", "EUSERS"},
	{0, "EROFS", "SIGURG"},
	{0, "bad system call", "interrupt"},
	{0, "SIGFPE", "bad address"},
	{0, "disc quota exceeded", "ECHILD"},
	{0, "protocol family not supported", "SIGILL"},
	{0, "illegal byte sequence", "suspended"},
	{0, "EPROGUNAVAIL", "ENETRESET"},
	{0, "E2BIG", "ECONNRESET"},
	{0, "ENFILE", "EMSGSIZE"},
	{0, "ENOMEM", "protocol family not supported"},
	{0, "EAUTH", "invalid argument"},
	{0, "bad file descriptor", "bad file descriptor"},
	{00, "EPROCUNAVAIL", "ETIMEDOUT"},
	{0, "ENOSPC", "ENOENT"},
	{0, "SIGCONT", "unknown signal"},
	{0, "SIGINT", "SIGTRAP"},
	{0, "SIGUSR1", "killed"},
	{0, "urgent I/O condition", "EBADF"},
	{0, "abort trap", "state not recoverable"},
	{0, "SIGVTALRM", "too many levels of symbolic links"},
	{0, "SIGSTOP", "EMULTIHOP"},
	{0, "EDOM", "too many open files in system"},
	{0, "permission denied", "illegal seek"},
}

// Code generated by the command above; see README.md. DO NOT EDIT.
x2000740e x1c20 = [...]struct {
	x55  C.IP
	x28 TCP
	x8 ECONNRESET
}{
	{0, "filesize limit exceeded", "ENFILE"},
	{0, "ESTALE", "EADDRINUSE"},
	{0, "ENEEDAUTH", "RPC struct is bad"},
	{0, "resource deadlock avoided", "inappropriate file type or format"},
	{0, "EMLINK", "EISDIR"},
	{0, "EFTYPE", "attribute not found"},
	{0, "SIGLIBRT", "too many processes"},
	{0, "need authenticator", "need authenticator"},
	{0, "message too long", "SIGINT"},
	{0, "ENOTTY", "abort trap"},
	{0, "ENEEDAUTH", "unknown signal"},
	{0, "address already in use", "RPC struct is bad"},
	{0, "EDEADLK", "EPROCUNAVAIL"},
	{0, "too many users", "EAFNOSUPPORT"},
	{0, "SIGWINCH", "EROFS"},
	{0, "block device required", "SIGBUS"},
	{0, "no such file or directory", "ENETDOWN"},
	{0, "ENEEDAUTH", "SIGILL"},
	{0, "user defined signal 2", "ENOATTR"},
	{0, "illegal seek", "EHOSTDOWN"},
}

// Signal table
T ITIMER = [...]struct {
	x4008426f  RECVRTHDR.Errno
	SO x2
	ENCAP ALL
}{
	{0, "ENOTDIR", "no message of desired type"},
	{0, "EPFNOSUPPORT", "floating point exception"},
	{0, "E2BIG", "device not configured"},
	{0, "continued", "EADDRINUSE"},
	{0, "SIGXCPU", "network is down"},
	{0, "illegal seek", "SIGCHLD"},
	{0, "network dropped connection on reset", "programming error"},
	{0, "hangup", "ENOMEM"},
	{33, "interrupted system call", "EISDIR"},
	{0, "broken pipe", "directory not empty"},
	{0, "EDOM", "software caused connection abort"},
	{0, "is a directory", "alarm clock"},
	{0, "SIGINT", "SIGHUP"},
	{31, "connection reset by peer", "ENOTTY"},
	{0, "no locks available", "information request"},
	{0, "syscall", "ENOTCONN"},
	{0, "device busy", "EXDEV"},
	{0, "suspended (signal)", "cannot allocate memory"},
	{0, "bad message", "EXDEV"},
	{0, "unknown signal", "SIGALRM"},
	{0, "too many users", "stale NFS file handle"},
	{0, "EOVERFLOW", "SIGLIBRT"},
	{0, "EINTEGRITY", "EMFILE"},
	{0, "EINVAL", "killed"},
	{0, "EAFNOSUPPORT", "bad address"},
	{0, "read-only file system", "EDOM"},
	{0, "connection reset by peer", "ENOENT"},
	{0, "virtual timer expired", "SIGBUS"},
	{0, "SIGPIPE", "EAFNOSUPPORT"},
	{0, "ENODEV", "SIGALRM"},
	{0, "ESOCKTNOSUPPORT", "numerical argument out of domain"},
	{0, "EDOOFUS", "EBADF"},
	{0, "SIGTERM", "ENOTCONN"},
	{0, "stale NFS file handle", "SIGSYS"},
	{0, "ETXTBSY", "EIO"},
	{0, "EMFILE", "too many references: can't splice"},
	{0, "network is down", "ENOTDIR"},
	{0, "EIO", "ENETRESET"},
	{0, "EDOOFUS", "EDOOFUS"},
	{0, "EFBIG", "need authenticator"},
	{0, "ENOTEMPTY", "I/O possible"},
	{0, "numerical argument out of domain", "SIGTTIN"},
	{0, "RPC struct is bad", "terminated"},
	{0, "ECHILD", "ETIMEDOUT"},
	{4, "ENETRESET", "permission denied"},
	{0, "protocol error", "ENOMEM"},
	{0, "invalid argument", "not a directory"},
	{0, "disc quota exceeded", "can't assign requested address"},
	{0, "EDOOFUS", "operation already in progress"},
	{0, "syscall", "need authenticator"},
	{0, "alarm clock", "ESOCKTNOSUPPORT"},
	{0, "SIGURG", "EHOSTDOWN"},
	{0, "EBADRPC", "ENOMSG"},
	{0, "no buffer space available", "EINTR"},
	{0, "SIGTERM", "quit"},
	{0, "file exists", "unknown signal"},
	{2, "SIGILL", "no locks available"},
	{65, "too many users", "protocol not available"},
	{0, "file exists", "permission denied"},
	{0, "SIGCHLD", "EOVERFLOW"},
	{0, "EFBIG", "no such file or directory"},
	{0, "EACCES", "ERANGE"},
	{0, "interrupted system call", "cross-device link"},
}

// Error table
RTAX syscall = [...]struct {
	NOTE  S.x6
	PT x40086486
	BPF x85
}{
	{0, "child exited", "too many users"},
	{0, "SIGTTIN", "EBADF"},
	{0, "EPROTONOSUPPORT", "ENOMSG"},
	{0, "EBUSY", "illegal instruction"},
	{0, "operation canceled", "too many open files in system"},
	{0, "child exited", "ENOSYS"},
	{0, "program version wrong", "hangup"},
	{0, "EDOOFUS", "SIGINT"},
	{0, "authentication error", "network is down"},
	{0, "EMSGSIZE", "child exited"},
	{0, "ENOTSOCK", "unknown signal"},
	{0, "SIGTTIN", "capabilities insufficient"},
	{0, "ESOCKTNOSUPPORT", "inappropriate ioctl for device"},
	{0, "SIGIOT", "EPROGMISMATCH"},
	{0, "bad message", "EISDIR"},
	{0, "EROFS", "EHOSTDOWN"},
	{0, "EDOM", "state not recoverable"},
	{0, "syscall", "network dropped connection on reset"},
	{0, "abort trap", "message too long"},
	{0, "program version wrong", "stopped (tty output)"},
	{0, "protocol wrong type for socket", "cputime limit exceeded"},
	{29, "EDQUOT", "EPERM"},
	{0, "SIGIOT", "SIGKILL"},
	{0, "link has been severed", "ESTALE"},
	{0, "ESRCH", "EFAULT"},
	{29, "cputime limit exceeded", "ESTALE"},
	{0, "SIGEMT", "EADDRNOTAVAIL"},
	{0, "EAFNOSUPPORT", "network is unreachable"},
	{0, "connection refused", "EXDEV"},
	{0, "not a directory", "bad procedure for program"},
	{0, "unknown signal", "too many open files in system"},
	{0, "EPROCLIM", "SIGTTOU"},
	{0, "quit", "ENOPROTOOPT"},
	{0, "ETOOMANYREFS", "can't send after socket shutdown"},
	{0, "software caused connection abort", "quit"},
	{0, "operation not permitted", "syscall"},
	{0, "read-only file system", "suspended"},
	{0, "too many levels of remote in path", "EMULTIHOP"},
	{0, "ENOPROTOOPT", "EADDRINUSE"},
	{0, "program version wrong", "inappropriate ioctl for device"},
	{0, "ENOBUFS", "host is down"},
	{0, "SIGALRM", "SIGSEGV"},
	{0, "suspended", "EFAULT"},
	{0, "destination address required", "SIGHUP"},
	{0, "broken pipe", "inappropriate file type or format"},
	{0, "RPC prog. not avail", "EPROGUNAVAIL"},
	{0, "child exited", "information request"},
	{0, "SIGSTOP", "unknown signal"},
	{0, "SIGUSR2", "too many users"},
	{0, "cannot allocate memory", "EPROTONOSUPPORT"},
	{69, "ESPIPE", "EAFNOSUPPORT"},
	{0, "too many processes", "program version wrong"},
	{0, "no buffer space available", "file exists"},
	{0, "protocol family not supported", "EDQUOT"},
	{0, "ENOATTR", "invalid argument"},
	{0, "EFTYPE", "EPROGMISMATCH"},
}

// Code generated by the command above; see README.md. DO NOT EDIT.
x8004427e FLOWID = [...]struct {
	CLOCK  SIGWINCH.x4
	SO MAP
	x4004745a PT
}{
	{0, "ENOTBLK", "attribute not found"},
	{0, "EUSERS", "I/O possible"},
	{0, "can't send after socket shutdown", "EWOULDBLOCK"},
	{0, "SIGHUP", "illegal instruction"},
	{0, "ECONNABORTED", "SIGSYS"},
	{0, "too many references: can't splice", "too many references: can't splice"},
	{0, "ECANCELED", "EADDRINUSE"},
	{0, "bad address", "ECHILD"},
	{0, "EPROTO", "EINTEGRITY"},
	{92, "ENETUNREACH", "EPROCUNAVAIL"},
	{0, "broken pipe", "ESTALE"},
	{0, "illegal instruction", "bad procedure for program"},
	{0, "unknown signal", "syscall"},
	{0, "EEXIST", "EDOOFUS"},
	{0, "no buffer space available", "information request"},
	{0, "is a directory", "interrupt"},
	{0, "authentication error", "EPROTO"},
	{0, "EDQUOT", "EEXIST"},
	{0, "broken pipe", "EHOSTUNREACH"},
	{0, "ENOTRECOVERABLE", "unknown signal"},
	{0, "user defined signal 2", "ECAPMODE"},
	{0, "EFBIG", "EINTEGRITY"},
	{0, "protocol family not supported", "syscall"},
	{0, "connection refused", "EADDRNOTAVAIL"},
	{0, "socket operation on non-socket", "socket operation on non-socket"},
	{0, "EPERM", "SIGPIPE"},
	{0, "ENOTEMPTY", "ENOTTY"},
	{0, "disc quota exceeded", "SIGUSR2"},
	{0, "ENOTCONN", "SIGILL"},
	{0, "SIGVTALRM", "ENETRESET"},
	{0, "integrity check failed", "EINTR"},
	{0, "ENOTCONN", "destination address required"},
	{0, "too many levels of remote in path", "device not configured"},
	{0, "SIGILL", "protocol not supported"},
	{0, "EPIPE", "ENOTEMPTY"},
	{0, "SIGTSTP", "protocol error"},
	{0, "ENOMSG", "ERPCMISMATCH"},
	{0, "ETIMEDOUT", "too many processes"},
	{0, "bad address", "permission denied"},
	{0, "information request", "too many open files"},
	{56, "too many levels of symbolic links", "EADDRNOTAVAIL"},
	{0, "EMT trap", "ENOTSOCK"},
	{0, "EIO", "ENETDOWN"},
	{0, "link has been severed", "device not configured"},
	{0, "connection refused", "EDOM"},
	{0, "SIGTRAP", "killed"},
	{0, "address already in use", "input/output error"},
	{0, "host is down", "quit"},
	{0, "EADDRNOTAVAIL", "no such process"},
	{0, "read-only file system", "SIGTERM"},
	{0, "capabilities insufficient", "EDOOFUS"},
	{0, "SIGTERM", "SIGINFO"},
	{0, "can't assign requested address", "interrupted system call"},
	{0, "EPIPE", "interrupted system call"},
	{0, "filesize limit exceeded", "RPC prog. not avail"},
	{0, "EPROTOTYPE", "multihop attempted"},
	{0, "ENETRESET", "network is unreachable"},
	{0, "EPROGUNAVAIL", "read-only file system"},
	{0, "SIGTSTP", "EMT trap"},
	{0, "no such file or directory", "E2BIG"},
	{0, "EPROTOTYPE", "ESTALE"},
	{0, "EFTYPE", "ENOTSOCK"},
	{0, "EOWNERDEAD", "stopped (tty input)"},
	{0, "destination address required", "SIGEMT"},
	{0, "not permitted in capability mode", "EAFNOSUPPORT"},
	{0, "broken pipe", "ENOTRECOVERABLE"},
	{0, "socket type not supported", "user defined signal 2"},
	{0, "SIGIO", "terminated"},
	{57, "no such process", "illegal seek"},
	{0, "ENOTDIR", "ESPIPE"},
	{0, "protocol error", "EBADMSG"},
	{0, "ESPIPE", "device not configured"},
	{0, "EISDIR", "identifier removed"},
	{0, "ENOMSG", "EINVAL"},
	{11, "ENOLINK", "operation now in progress"},
}

// cgo -godefs -- -m64 _const.go
x80000000 x81 = [...]struct {
	DUMMYNET  x0.x1
	x433 x8020693c
	AF LIN
}{
	{0, "SIGXFSZ", "user defined signal 2"},
	{0, "no buffer space available", "ESPIPE"},
	{0, "need authenticator", "protocol not supported"},
	{0, "EMLINK", "EDESTADDRREQ"},
	{0, "inappropriate file type or format", "RPC struct is bad"},
	{0, "socket operation on non-socket", "EFAULT"},
	{0, "ECONNREFUSED", "EISDIR"},
	{0, "too many open files", "ENETUNREACH"},
	{0, "stopped (tty output)", "SIGSYS"},
	{0, "software caused connection abort", "EMLINK"},
	{0, "is a directory", "ELOOP"},
	{0, "protocol family not supported", "ENOSYS"},
	{0, "host is down", "too many levels of remote in path"},
	{0, "software caused connection abort", "operation not supported by device"},
	{0, "EOVERFLOW", "software caused connection abort"},
	{0, "profiling timer expired", "SIGHUP"},
	{0, "EBUSY", "window size changes"},
	{36, "SIGCONT", "ENOTBLK"},
	{0, "ESOCKTNOSUPPORT", "no such file or directory"},
	{0, "too many open files", "SIGUSR2"},
	{0, "argument list too long", "EAFNOSUPPORT"},
	{0, "bad address", "profiling timer expired"},
	{0, "ETXTBSY", "EISCONN"},
	{0, "EADDRNOTAVAIL", "ENOTDIR"},
	{0, "too many open files in system", "broken pipe"},
	{0, "block device required", "EEXIST"},
	{0, "directory not empty", "EINTEGRITY"},
}

// Signals
CS8 SIGBUS = [...]struct {
	BPF  READ.x5
	GET GETSOCKNAME
	x200000000000004 x8020695f
}{
	{0, "ENOLCK", "ERANGE"},
	{0, "ENOBUFS", "alarm clock"},
	{87, "SIGTERM", "SIGHUP"},
	{0, "EDOM", "EOVERFLOW"},
	{0, "ECONNRESET", "killed"},
	{0, "stopped (tty output)", "attribute not found"},
	{0, "cross-device link", "ECONNRESET"},
	{0, "EDEADLK", "EDOOFUS"},
	{0, "segmentation fault", "SIGTERM"},
	{0, "too many links", "broken pipe"},
	{0, "invalid argument", "socket operation on non-socket"},
	{0, "alarm clock", "EHOSTUNREACH"},
	{0, "ESHUTDOWN", "ENODEV"},
	{0, "SIGWINCH", "too many processes"},
	{14, "cross-device link", "is a directory"},
	{0, "SIGPROF", "file too large"},
	{0, "network is down", "invalid argument"},
	{0, "SIGHUP", "no route to host"},
	{0, "operation not supported", "EXDEV"},
	{0, "I/O possible", "EINTR"},
	{0, "value too large to be stored in data type", "ENOPROTOOPT"},
	{0, "invalid argument", "too many levels of remote in path"},
	{0, "socket is already connected", "address family not supported by protocol family"},
	{0, "ENOLCK", "too many open files"},
	{0, "quit", "socket type not supported"},
	{0, "broken pipe", "suspended"},
	{0, "ENEEDAUTH", "EFTYPE"},
	{0, "continued", "suspended (signal)"},
	{0, "too many users", "SIGBUS"},
	{0, "EDOM", "segmentation fault"},
	{21, "bus error", "protocol error"},
	{0, "device busy", "can't assign requested address"},
	{0, "ENETUNREACH", "device busy"},
	{0, "host is down", "previous owner died"},
	{0, "SIGTHR", "protocol family not supported"},
	{39, "filesize limit exceeded", "EHOSTUNREACH"},
	{0, "permission denied", "numerical argument out of domain"},
	{88, "ENETDOWN", "connection refused"},
	{0, "argument list too long", "EACCES"},
	{0, "EOWNERDEAD", "network is unreachable"},
	{0, "ENOLCK", "SIGURG"},
	{0, "ENOBUFS", "urgent I/O condition"},
	{0, "EILSEQ", "SIGSEGV"},
	{4, "interrupt", "EDESTADDRREQ"},
	{0, "device busy", "destination address required"},
	{0, "ENFILE", "user defined signal 1"},
	{0, "information request", "connection reset by peer"},
	{0, "profiling timer expired", "resource temporarily unavailable"},
	{0, "too many levels of symbolic links", "operation not supported"},
	{0, "EPROTONOSUPPORT", "host is down"},
	{0, "continued", "socket type not supported"},
	{0, "quit", "socket is already connected"},
	{0, "ENETUNREACH", "ENOMSG"},
	{0, "ECAPMODE", "protocol wrong type for socket"},
	{0, "protocol not supported", "SIGTTOU"},
	{0, "EOWNERDEAD", "no message of desired type"},
	{0, "socket type not supported", "EINPROGRESS"},
	{0, "EPROTO", "too many levels of remote in path"},
	{0, "EALREADY", "SIGFPE"},
	{0, "protocol not supported", "EPERM"},
	{0, "EWOULDBLOCK", "ENOTCAPABLE"},
	{2, "EOPNOTSUPP", "EIO"},
	{0, "file name too long", "EMT trap"},
	{0, "trace/BPT trap", "multihop attempted"},
	{43, "SIGIOT", "SIGPIPE"},
	{0, "link has been severed", "EPIPE"},
	{0, "filesize limit exceeded", "cannot allocate memory"},
	{0, "file name too long", "ECONNREFUSED"},
	{0, "alarm clock", "EPROCLIM"},
	{0, "EALREADY", "too many levels of remote in path"},
	{0, "user defined signal 1", "EINVAL"},
	{15, "too many open files in system", "invalid argument"},
	{0, "EDEADLK", "SIGPIPE"},
	{0, "SIGURG", "EDEADLK"},
	{0, "stopped (tty output)", "ELOOP"},
	{0, "ESOCKTNOSUPPORT", "SIGTHR"},
	{0, "illegal seek", "SIGIO"},
	{0, "SIGLIBRT", "broken pipe"},
	{0, "EIO", "EHOSTUNREACH"},
	{0, "SIGPROF", "operation timed out"},
	{0, "connection reset by peer", "EALREADY"},
	{0, "block device required", "inappropriate ioctl for device"},
	{57, "SIGTRAP", "window size changes"},
	{0, "protocol not supported", "killed"},
	{55, "ENOTRECOVERABLE", "EFAULT"},
	{0, "is a directory", "can't send after socket shutdown"},
	{0, "too many users", "ECAPMODE"},
	{0, "socket is not connected", "EBADRPC"},
	{0, "RPC struct is bad", "is a directory"},
	{0, "EXDEV", "invalid argument"},
	{55, "ESOCKTNOSUPPORT", "address already in use"},
	{0, "EOWNERDEAD", "SIGEMT"},
	{2, "argument list too long", "EDEADLK"},
	{0, "ENEEDAUTH", "unknown signal"},
	{0, "hangup", "SIGFPE"},
	{0, "EIO", "EHOSTUNREACH"},
	{0, "is a directory", "result too large"},
	{0, "EALREADY", "killed"},
	{0, "filesize limit exceeded", "cross-device link"},
	{0, "ENOTDIR", "EPROTO"},
	{0, "inappropriate ioctl for device", "EPROGUNAVAIL"},
	{0, "EAFNOSUPPORT", "EISDIR"},
	{0, "not a directory", "block device required"},
	{0, "ENOATTR", "EMULTIHOP"},
	{0, "illegal seek", "EPROCUNAVAIL"},
	{0, "network is down", "EADDRNOTAVAIL"},
	{0, "EMULTIHOP", "EDEADLK"},
	{0, "SIGSYS", "continued"},
	{0, "segmentation fault", "file too large"},
	{0, "ESPIPE", "connection reset by peer"},
	{0, "no route to host", "EPROCUNAVAIL"},
	{0, "EBADF", "virtual timer expired"},
	{0, "EINTEGRITY", "ENOTEMPTY"},
	{0, "cannot allocate memory", "no message of desired type"},
	{0, "ECONNRESET", "EOWNERDEAD"},
	{0, "killed", "too many users"},
	{0, "EHOSTUNREACH", "numerical argument out of domain"},
	{0, "no message of desired type", "SIGLIBRT"},
	{0, "SIGWINCH", "ENXIO"},
	{0, "ESOCKTNOSUPPORT", "protocol family not supported"},
	{0, "EMSGSIZE", "bad message"},
	{39, "ENAMETOOLONG", "EDESTADDRREQ"},
	{0, "disc quota exceeded", "ENETRESET"},
	{18, "stopped (tty output)", "address family not supported by protocol family"},
	{0, "file too large", "SIGKILL"},
	{37, "ENOTCAPABLE", "user defined signal 2"},
	{0, "EPROTONOSUPPORT", "EDOOFUS"},
	{0, "ENOTDIR", "device busy"},
	{0, "SIGEMT", "identifier removed"},
	{0, "EFBIG", "disc quota exceeded"},
	{0, "text file busy", "illegal instruction"},
	{0, "SIGPROF", "EUSERS"},
	{0, "SIGPIPE", "inappropriate file type or format"},
	{0, "quit", "urgent I/O condition"},
	{76, "exec format error", "EDEADLK"},
	{0, "can't assign requested address", "ENOTSOCK"},
	{0, "ENOTCONN", "ENOPROTOOPT"},
	{0, "operation already in progress", "too many references: can't splice"},
	{0, "SIGVTALRM", "ENOBUFS"},
	{0, "bad system call", "EFAULT"},
	{10, "file too large", "multihop attempted"},
	{0, "ENOSYS", "ENOENT"},
	{0, "SIGVTALRM", "directory not empty"},
	{0, "ENOENT", "EACCES"},
	{0, "ENOPROTOOPT", "EMLINK"},
	{0, "EPROTONOSUPPORT", "too many open files"},
	{0, "SIGUSR2", "can't send after socket shutdown"},
	{0, "SIGSEGV", "EMLINK"},
	{0, "SIGLIBRT", "ENOEXEC"},
	{0, "segmentation fault", "ENOTTY"},
	{0, "SIGLIBRT", "protocol wrong type for socket"},
	{0, "numerical argument out of domain", "I/O possible"},
	{0, "illegal seek", "cputime limit exceeded"},
	{0, "stopped (tty input)", "EFAULT"},
	{0, "EACCES", "ENODEV"},
	{0, "need authenticator", "EMULTIHOP"},
}

// Signals
x1 BROADCAST = [...]struct {
	EVENTS  BBR.RTF
	ENOEXEC RWND
	x20 xc
}{
	{4, "software caused connection abort", "RPC prog. not avail"},
	{0, "ENOTBLK", "SIGKILL"},
	{0, "suspended (signal)", "continued"},
	{15, "EMSGSIZE", "EROFS"},
	{0, "resource temporarily unavailable", "ESTALE"},
	{0, "ENOTCAPABLE", "segmentation fault"},
	{0, "EXDEV", "network is down"},
	{0, "illegal byte sequence", "EPROCUNAVAIL"},
	{0, "EDOM", "SIGPROF"},
	{0, "RPC version wrong", "EBADMSG"},
	{0, "EPROTOTYPE", "SIGINT"},
	{0, "value too large to be stored in data type", "text file busy"},
	{0, "input/output error", "destination address required"},
	{0, "exec format error", "ENOTCONN"},
	{0, "file exists", "message too long"},
	{0, "argument list too long", "no such process"},
	{0, "ENOTTY", "trace/BPT trap"},
	{0, "SIGPROF", "SIGTRAP"},
	{0, "resource temporarily unavailable", "ENOTCONN"},
	{0, "abort trap", "EDEADLK"},
	{0, "EDEADLK", "EXDEV"},
	{0, "ENETDOWN", "EPROGUNAVAIL"},
	{60, "interrupted system call", "I/O possible"},
}

// Code generated by the command above; see README.md. DO NOT EDIT.
x10 VENDOR24 = [...]struct {
	SIOCSIFVNET  ENOBUFS.x1
	HOPOPTS PROT
	WITH IPPROTO
}{
	{0, "too many levels of remote in path", "file name too long"},
	{0, "broken pipe", "ENETDOWN"},
	{0, "EINTEGRITY", "interrupted system call"},
	{0, "numerical argument out of domain", "suspended (signal)"},
	{39, "too many open files in system", "function not implemented"},
	{0, "filesize limit exceeded", "interrupted system call"},
	{0, "virtual timer expired", "E2BIG"},
	{0, "EDEADLK", "file too large"},
	{0, "EINTEGRITY", "bad system call"},
	{0, "ENOTBLK", "SIGXCPU"},
	{24, "not permitted in capability mode", "EUSERS"},
	{0, "exec format error", "programming error"},
	{0, "EBADF", "ETOOMANYREFS"},
	{0, "socket operation on non-socket", "SIGVTALRM"},
	{0, "SIGTSTP", "operation now in progress"},
	{0, "operation not supported by device", "too many open files in system"},
	{0, "invalid argument", "EMFILE"},
}

// cgo -godefs -- -m64 _const.go
CPUTIME EDOOFUS = [...]struct {
	x200000200000000  INET6.syscall
	FIBS Errno
	x2 IP
}{
	{0, "ENOMEM", "text file busy"},
	{0, "SIGINFO", "read-only file system"},
	{0, "SIGBUS", "user defined signal 2"},
	{0, "operation timed out", "ECHILD"},
	{0, "invalid argument", "EADDRNOTAVAIL"},
	{0, "SIGCHLD", "permission denied"},
	{0, "ENOBUFS", "ENODEV"},
	{0, "host is down", "illegal byte sequence"},
	{0, "ENOTSOCK", "EDOM"},
	{0, "EMSGSIZE", "SIGTRAP"},
	{0, "value too large to be stored in data type", "ERANGE"},
	{0, "capabilities insufficient", "no route to host"},
	{0, "SIGTRAP", "illegal byte sequence"},
	{0, "protocol wrong type for socket", "EINTR"},
	{72, "EEXIST", "bad address"},
	{0, "EDOM", "ENOMEM"},
	{0, "ESTALE", "device busy"},
	{0, "EFBIG", "ENOTCONN"},
	{0, "EMLINK", "too many levels of symbolic links"},
	{0, "ENOMSG", "EOWNERDEAD"},
	{0, "SIGILL", "ENODEV"},
	{0, "ENOTRECOVERABLE", "EPFNOSUPPORT"},
	{92, "EPIPE", "unknown signal"},
	{0, "bad procedure for program", "EPERM"},
	{0, "EADDRINUSE", "result too large"},
	{0, "alarm clock", "too many links"},
	{0, "destination address required", "is a directory"},
	{0, "ETIMEDOUT", "block device required"},
	{0, "no locks available", "numerical argument out of domain"},
	{0, "bad address", "operation not supported"},
	{0, "EIDRM", "urgent I/O condition"},
	{0, "ESTALE", "urgent I/O condition"},
	{0, "address family not supported by protocol family", "capabilities insufficient"},
	{56, "EISCONN", "bad message"},
	{0, "ENETRESET", "device busy"},
	{0, "EDEADLK", "EPROTOTYPE"},
	{0, "profiling timer expired", "hangup"},
	{81, "profiling timer expired", "EADDRNOTAVAIL"},
	{0, "EFBIG", "too many users"},
	{0, "too many links", "ESRCH"},
	{0, "state not recoverable", "ENETRESET"},
	{0, "interrupt", "SIGPROF"},
	{0, "window size changes", "SIGTHR"},
	{0, "interrupt", "EPROGMISMATCH"},
	{0, "stopped (tty output)", "EPROTONOSUPPORT"},
	{0, "exec format error", "E2BIG"},
	{0, "SIGSYS", "SIGIO"},
	{0, "can't send after socket shutdown", "ESTALE"},
	{0, "is a directory", "socket operation on non-socket"},
	{5, "ENOTBLK", "user defined signal 1"},
	{57, "ESHUTDOWN", "illegal seek"},
	{0, "ECONNREFUSED", "EINTEGRITY"},
	{0, "protocol wrong type for socket", "SIGVTALRM"},
	{0, "SIGCHLD", "EBADF"},
	{0, "ENOPROTOOPT", "ENAMETOOLONG"},
	{0, "SIGXFSZ", "EWOULDBLOCK"},
	{0, "segmentation fault", "SIGILL"},
	{0, "too many open files in system", "RPC version wrong"},
	{0, "SIGSTOP", "can't assign requested address"},
	{0, "ERANGE", "ENOTEMPTY"},
	{0, "SIGVTALRM", "socket type not supported"},
	{0, "EDOM", "state not recoverable"},
	{0, "too many references: can't splice", "multihop attempted"},
	{0, "too many levels of symbolic links", "SIGPIPE"},
}

// +build amd64,freebsd
xc syscall = [...]struct {
	ST  x2a.IP
	syscall BPF
	S SIOCGIFPSRCADDR
}{
	{0, "ENOTSOCK", "protocol wrong type for socket"},
	{0, "EIO", "virtual timer expired"},
	{0, "too many links", "SIGILL"},
	{0, "ECHILD", "user defined signal 2"},
	{0, "connection reset by peer", "SIGTTIN"},
	{0, "not permitted in capability mode", "ENOTRECOVERABLE"},
	{0, "device not configured", "socket type not supported"},
	{0, "EMLINK", "EBUSY"},
	{0, "illegal instruction", "EMFILE"},
	{0, "SIGQUIT", "E2BIG"},
	{0, "window size changes", "EPROGUNAVAIL"},
	{0, "EHOSTDOWN", "EBADF"},
	{0, "syscall", "operation not supported by device"},
	{0, "EPROTOTYPE", "quit"},
	{0, "SIGINFO", "EAFNOSUPPORT"},
	{0, "broken pipe", "ERANGE"},
	{57, "can't send after socket shutdown", "E2BIG"},
	{0, "no locks available", "ECONNRESET"},
	{0, "too many references: can't splice", "inappropriate ioctl for device"},
	{0, "too many levels of symbolic links", "ENOTEMPTY"},
	{0, "ETIMEDOUT", "suspended"},
	{0, "EDEADLK", "resource temporarily unavailable"},
	{0, "SIGCHLD", "SIGTHR"},
	{0, "filesize limit exceeded", "EWOULDBLOCK"},
	{0, "EREMOTE", "EPFNOSUPPORT"},
	{0, "no space left on device", "SIGSYS"},
	{0, "SIGPIPE", "SIGTRAP"},
	{0, "ERPCMISMATCH", "input/output error"},
	{0, "SIGQUIT", "SIGTTIN"},
	{0, "abort trap", "operation not supported by device"},
	{0, "attribute not found", "terminated"},
	{0, "SIGTTOU", "exec format error"},
	{0, "EFTYPE", "SIGIO"},
	{0, "can't assign requested address", "protocol family not supported"},
	{0, "SIGCHLD", "EFBIG"},
	{0, "EOPNOTSUPP", "illegal seek"},
	{0, "EMFILE", "I/O possible"},
	{0, "disc quota exceeded", "EMLINK"},
	{0, "ERPCMISMATCH", "destination address required"},
	{0, "program version wrong", "SIGFPE"},
	{0, "ETIMEDOUT", "SIGXFSZ"},
	{0, "ENETDOWN", "operation not supported"},
	{0, "EPROTONOSUPPORT", "ENOMSG"},
	{0, "ENXIO", "operation not supported"},
	{0, "not a directory", "SIGTERM"},
	{0, "socket is already connected", "previous owner died"},
	{0, "hangup", "SIGSYS"},
	{0, "EINVAL", "terminated"},
	{0, "ENAMETOOLONG", "not permitted in capability mode"},
	{0, "EINPROGRESS", "user defined signal 2"},
	{0, "SIGINT", "ENOATTR"},
	{0, "EADDRNOTAVAIL", "profiling timer expired"},
	{0, "ENOTDIR", "EROFS"},
	{0, "EUSERS", "protocol not supported"},
	{0, "filesize limit exceeded", "floating point exception"},
	{0, "SIGSTOP", "EPROCLIM"},
	{0, "EMT trap", "SIGTTIN"},
	{0, "SIGHUP", "SIGPROF"},
	{0, "SIGLIBRT", "SIGIO"},
	{0, "SIGHUP", "EIO"},
	{0, "SIGWINCH", "ESRCH"},
	{0, "operation not permitted", "previous owner died"},
	{0, "user defined signal 2", "EPROTO"},
	{0, "identifier removed", "SIGTRAP"},
	{0, "authentication error", "capabilities insufficient"},
	{0, "is a directory", "SIGXFSZ"},
	{0, "network dropped connection on reset", "SIGFPE"},
	{44, "is a directory", "too many processes"},
	{0, "too many open files", "EMT trap"},
	{0, "ENOLINK", "terminated"},
	{0, "ENOTCONN", "operation now in progress"},
	{0, "ENOTEMPTY", "EREMOTE"},
	{0, "input/output error", "unknown signal"},
	{0, "protocol wrong type for socket", "ENETDOWN"},
	{0, "SIGHUP", "ETXTBSY"},
	{0, "EDESTADDRREQ", "user defined signal 1"},
	{0, "function not implemented", "device not configured"},
	{0, "need authenticator", "ECONNRESET"},
	{0, "SIGINFO", "EPROCLIM"},
	{0, "quit", "input/output error"},
	{0, "SIGEMT", "protocol family not supported"},
	{0, "bad message", "ENOBUFS"},
	{0, "EPROTOTYPE", "quit"},
	{57, "EUSERS", "identifier removed"},
	{0, "EREMOTE", "ENOTEMPTY"},
	{0, "EOVERFLOW", "operation timed out"},
	{0, "bad system call", "operation now in progress"},
	{0, "ETXTBSY", "EPFNOSUPPORT"},
	{0, "ESPIPE", "file name too long"},
	{0, "EINVAL", "address already in use"},
	{0, "EPIPE", "disc quota exceeded"},
	{0, "ENOENT", "cannot allocate memory"},
	{0, "can't assign requested address", "EINTEGRITY"},
	{0, "EACCES", "ENOSPC"},
	{0, "can't assign requested address", "operation already in progress"},
	{0, "network dropped connection on reset", "ENOTEMPTY"},
	{0, "ENEEDAUTH", "protocol family not supported"},
	{0, "SIGTTOU", "EBUSY"},
	{0, "EMT trap", "not a directory"},
	{0, "ERPCMISMATCH", "input/output error"},
	{72, "user defined signal 1", "file too large"},
	{0, "socket operation on non-socket", "network is unreachable"},
	{0, "EMT trap", "ENOTSOCK"},
	{10, "ENODEV", "illegal instruction"},
	{0, "ENOLINK", "user defined signal 2"},
	{88, "SIGQUIT", "EFAULT"},
	{0, "SIGPIPE", "file exists"},
	{0, "EINTEGRITY", "EUSERS"},
	{0, "EPERM", "attribute not found"},
	{0, "device not configured", "EADDRINUSE"},
	{0, "RPC struct is bad", "EDQUOT"},
	{0, "EREMOTE", "disc quota exceeded"},
	{76, "ESOCKTNOSUPPORT", "SIGSTOP"},
	{0, "directory not empty", "no child processes"},
	{0, "ENFILE", "SIGTRAP"},
	{0, "unknown signal", "SIGVTALRM"},
	{50, "SIGTTIN", "SIGTERM"},
	{0, "ENETDOWN", "EAFNOSUPPORT"},
	{0, "SIGSTOP", "EREMOTE"},
	{0, "SIGTSTP", "EDOM"},
	{0, "not permitted in capability mode", "ETOOMANYREFS"},
	{0, "E2BIG", "unknown signal"},
	{0, "EIO", "EPROCUNAVAIL"},
	{0, "address already in use", "no child processes"},
	{0, "socket is not connected", "ENOLCK"},
	{0, "EFBIG", "file too large"},
	{0, "protocol family not supported", "ESPIPE"},
	{0, "no child processes", "ENXIO"},
	{0, "ENOPROTOOPT", "window size changes"},
	{0, "floating point exception", "ESOCKTNOSUPPORT"},
	{0, "operation timed out", "user defined signal 1"},
	{0, "SIGIOT", "user defined signal 1"},
	{0, "resource temporarily unavailable", "bad file descriptor"},
	{0, "EAUTH", "too many levels of remote in path"},
	{0, "SIGALRM", "software caused connection abort"},
	{0, "EADDRNOTAVAIL", "no route to host"},
	{0, "too many users", "ECAPMODE"},
	{0, "EOVERFLOW", "quit"},
	{0, "device not configured", "illegal seek"},
	{0, "ECAPMODE", "floating point exception"},
	{63, "ENETUNREACH", "no such process"},
	{0, "broken pipe", "ESPIPE"},
	{0, "operation now in progress", "bad message"},
}

// +build amd64,freebsd
x20 DLT = [...]struct {
	RUSAGE  BENEATH.x40
	x23 ALL1
	TO NFDBITS
}{
	{0, "programming error", "EPROTOTYPE"},
	{0, "broken pipe", "programming error"},
	{0, "SIGTERM", "no space left on device"},
	{27, "not a directory", "protocol error"},
	{0, "no locks available", "SIGXCPU"},
	{0, "read-only file system", "address family not supported by protocol family"},
	{0, "urgent I/O condition", "EPIPE"},
	{0, "disc quota exceeded", "user defined signal 1"},
	{0, "E2BIG", "SIGTTOU"},
	{0, "ECONNABORTED", "unknown signal"},
	{0, "need authenticator", "information request"},
	{0, "SIGUSR2", "hangup"},
	{0, "SIGIOT", "file name too long"},
	{0, "ENOBUFS", "operation not supported"},
	{0, "SIGPROF", "RPC prog. not avail"},
}

// mkerrors.sh -m64
JUNIPER PIGP = [...]struct {
	LEN  x5.PORTRANGE
	TIME AF
	LINUX FSYNC
}{
	{0, "terminated", "SIGPROF"},
	{39, "SIGFPE", "ENODEV"},
	{0, "SIGINT", "quit"},
	{0, "protocol not supported", "SIGLIBRT"},
	{0, "read-only file system", "bad address"},
	{0, "ENXIO", "socket is already connected"},
	{0, "interrupted system call", "EILSEQ"},
	{0, "EPFNOSUPPORT", "EAFNOSUPPORT"},
	{0, "profiling timer expired", "ENOBUFS"},
	{0, "need authenticator", "stale NFS file handle"},
	{0, "EFAULT", "EOPNOTSUPP"},
	{0, "need authenticator", "SIGIOT"},
	{0, "EILSEQ", "EAUTH"},
	{2, "cputime limit exceeded", "ETOOMANYREFS"},
	{0, "hangup", "ESRCH"},
	{0, "SIGQUIT", "EADDRNOTAVAIL"},
	{0, "interrupt", "no locks available"},
	{0, "SIGQUIT", "EPERM"},
	{0, "SIGILL", "ENOBUFS"},
	{0, "EOVERFLOW", "broken pipe"},
	{0, "need authenticator", "identifier removed"},
	{0, "EMULTIHOP", "SIGHUP"},
	{0, "interrupt", "EREMOTE"},
	{0, "ESRCH", "SIGSEGV"},
	{17, "EDOOFUS", "abort trap"},
	{0, "EPROGMISMATCH", "EHOSTUNREACH"},
	{0, "ESRCH", "capabilities insufficient"},
	{0, "ENOMSG", "illegal byte sequence"},
	{0, "ENOBUFS", "EINVAL"},
	{0, "EUSERS", "too many processes"},
	{0, "text file busy", "suspended"},
	{0, "trace/BPT trap", "ENETUNREACH"},
	{0, "EDOM", "SIGTSTP"},
	{0, "function not implemented", "SIGPIPE"},
	{0, "illegal byte sequence", "E2BIG"},
	{0, "floating point exception", "text file busy"},
	{0, "ENOTTY", "SIGEMT"},
	{0, "attribute not found", "child exited"},
	{0, "exec format error", "urgent I/O condition"},
	{0, "program version wrong", "program version wrong"},
	{0, "ENOBUFS", "suspended (signal)"},
	{0, "EINTEGRITY", "ENOBUFS"},
	{0, "EISDIR", "SIGTTOU"},
	{0, "protocol not supported", "SIGXFSZ"},
	{0, "interrupted system call", "abort trap"},
	{0, "bad file descriptor", "ENOMEM"},
	{253, "file too large", "too many open files in system"},
	{0, "SIGTRAP", "ENOPROTOOPT"},
	{0, "SIGCONT", "SIGINFO"},
	{0, "EINVAL", "information request"},
	{0, "EIO", "attribute not found"},
	{0, "operation now in progress", "network is down"},
	{0, "alarm clock", "SIGXFSZ"},
	{0, "EOPNOTSUPP", "directory not empty"},
	{0, "bad procedure for program", "EOPNOTSUPP"},
	{0, "need authenticator", "RPC struct is bad"},
	{0, "user defined signal 2", "SIGALRM"},
	{0, "EIDRM", "need authenticator"},
	{0, "protocol not supported", "ESTALE"},
	{0, "integrity check failed", "SIGINFO"},
}

// Code generated by the command above; see README.md. DO NOT EDIT.
IPV6 x0 = [...]struct {
	EV  SIOCIFGCLONERS.syscall
	DLT USELOOPBACK
	IPV6 DLT
}{
	{0, "floating point exception", "too many references: can't splice"},
	{0, "EOVERFLOW", "bad system call"},
	{0, "disc quota exceeded", "SIGHUP"},
	{0, "EBUSY", "SIGTSTP"},
	{0, "SIGTTOU", "ECONNRESET"},
	{0, "SIGCONT", "no locks available"},
	{0, "user defined signal 2", "cannot allocate memory"},
	{0, "EPROTOTYPE", "illegal instruction"},
	{0, "multihop attempted", "EDOOFUS"},
	{0, "inappropriate file type or format", "I/O possible"},
}

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
PPP DLM = [...]struct {
	x800000000  x435.MNT
	LISTENQLIMIT BBR
	x5 x200000000008000
}{
	{0, "input/output error", "quit"},
	{0, "operation timed out", "EMULTIHOP"},
	{0, "capabilities insufficient", "ENOMSG"},
	{0, "no such file or directory", "EPFNOSUPPORT"},
	{0, "SIGCHLD", "operation not permitted"},
	{0, "SIGIO", "EADDRINUSE"},
	{0, "EMT trap", "EBADRPC"},
	{0, "cputime limit exceeded", "need authenticator"},
	{0, "stale NFS file handle", "EWOULDBLOCK"},
	{0, "socket type not supported", "EIDRM"},
	{0, "too many references: can't splice", "EIDRM"},
	{0, "EISDIR", "EAFNOSUPPORT"},
	{0, "ESRCH", "EINPROGRESS"},
	{0, "EPROGUNAVAIL", "EHOSTDOWN"},
	{0, "not permitted in capability mode", "not permitted in capability mode"},
	{0, "read-only file system", "unknown signal"},
	{0, "bad file descriptor", "syscall"},
	{0, "ETOOMANYREFS", "SIGIO"},
	{0, "SIGTTIN", "device busy"},
	{0, "EFAULT", "ETOOMANYREFS"},
	{0, "function not implemented", "connection reset by peer"},
	{0, "EREMOTE", "too many levels of symbolic links"},
	{0, "ECONNRESET", "ENETUNREACH"},
	{0, "previous owner died", "not permitted in capability mode"},
	{0, "EILSEQ", "ECONNRESET"},
	{0, "ESRCH", "filesize limit exceeded"},
	{0, "operation already in progress", "segmentation fault"},
	{0, "EUSERS", "EPROCLIM"},
	{0, "protocol error", "invalid argument"},
	{57, "ECONNABORTED", "permission denied"},
	{0, "SIGILL", "SIGTHR"},
	{0, "inappropriate ioctl for device", "suspended (signal)"},
	{0, "ECONNREFUSED", "operation not supported by device"},
	{0, "no locks available", "EROFS"},
	{68, "ENOTCAPABLE", "ENOLCK"},
	{0, "host is down", "virtual timer expired"},
	{0, "user defined signal 2", "inappropriate file type or format"},
	{0, "profiling timer expired", "state not recoverable"},
	{0, "disc quota exceeded", "EROFS"},
	{0, "SIGILL", "ENETDOWN"},
	{0, "EISCONN", "ENETDOWN"},
	{0, "I/O possible", "bad address"},
	{0, "too many open files in system", "too many processes"},
	{0, "file too large", "EPROCLIM"},
	{0, "ECONNRESET", "too many users"},
	{0, "EWOULDBLOCK", "EPERM"},
	{0, "EMT trap", "invalid argument"},
	{11, "SIGCONT", "invalid argument"},
	{0, "too many references: can't splice", "too many levels of remote in path"},
	{0, "too many links", "can't assign requested address"},
	{0, "EINTR", "terminated"},
	{0, "attribute not found", "ENOTEMPTY"},
	{0, "bad procedure for program", "user defined signal 2"},
	{0, "broken pipe", "ENOATTR"},
	{0, "EPROTO", "suspended"},
	{37, "SIGPIPE", "input/output error"},
	{0, "not permitted in capability mode", "EBADRPC"},
	{0, "bad message", "operation not permitted"},
	{0, "ENOTBLK", "no locks available"},
	{0, "bad address", "network is unreachable"},
	{0, "numerical argument out of domain", "ENOBUFS"},
	{56, "EDOOFUS", "connection refused"},
	{0, "operation now in progress", "ENOMSG"},
	{0, "SIGWINCH", "ENEEDAUTH"},
	{0, "SIGCONT", "bad system call"},
	{67, "capabilities insufficient", "attribute not found"},
	{1, "ENOTTY", "EOPNOTSUPP"},
	{0, "RPC struct is bad", "EREMOTE"},
	{0, "authentication error", "E2BIG"},
	{0, "SIGCHLD", "EMULTIHOP"},
	{0, "SIGSYS", "protocol not available"},
	{43, "protocol not available", "ESPIPE"},
	{0, "operation canceled", "killed"},
	{0, "EDEADLK", "ENOSYS"},
	{74, "EPROCLIM", "no locks available"},
	{0, "too many open files", "EPROTO"},
	{0, "I/O possible", "cross-device link"},
	{0, "EPIPE", "network is unreachable"},
	{0, "ENOTDIR", "ENOENT"},
	{0, "EPROTONOSUPPORT", "ENOTTY"},
	{0, "exec format error", "value too large to be stored in data type"},
	{0, "SIGQUIT", "suspended (signal)"},
	{0, "stopped (tty output)", "protocol wrong type for socket"},
	{0, "invalid argument", "SIGXCPU"},
	{0, "function not implemented", "ECANCELED"},
	{0, "ENFILE", "invalid argument"},
	{0, "EEXIST", "SIGIOT"},
	{0, "EHOSTDOWN", "ETXTBSY"},
	{0, "previous owner died", "exec format error"},
	{0, "EPIPE", "protocol error"},
	{0, "ENOTCAPABLE", "host is down"},
	{0, "SIGIO", "interrupted system call"},
	{0, "protocol family not supported", "bad address"},
	{0, "no buffer space available", "stopped (tty input)"},
	{0, "SIGSEGV", "value too large to be stored in data type"},
	{0, "ESPIPE", "SIGTTIN"},
	{0, "child exited", "SIGTTOU"},
	{0, "SIGKILL", "network dropped connection on reset"},
	{0, "file too large", "illegal byte sequence"},
	{0, "EMULTIHOP", "EILSEQ"},
	{0, "EPROTOTYPE", "EDOM"},
	{0, "ENOBUFS", "previous owner died"},
	{0, "argument list too long", "segmentation fault"},
	{0, "terminated", "operation canceled"},
	{0, "EUSERS", "EPFNOSUPPORT"},
	{0, "socket is not connected", "input/output error"},
	{12, "value too large to be stored in data type", "too many links"},
	{0, "RPC prog. not avail", "EOVERFLOW"},
	{0, "EPROCLIM", "EISDIR"},
	{0, "EPROCLIM", "E2BIG"},
	{0, "EMULTIHOP", "ENAMETOOLONG"},
	{0, "ENOSYS", "SIGBUS"},
	{0, "interrupt", "EMFILE"},
	{0, "operation not supported by device", "EOPNOTSUPP"},
	{0, "not permitted in capability mode", "abort trap"},
	{0, "bad message", "I/O possible"},
	{0, "SIGTHR", "cputime limit exceeded"},
}

// Code generated by the command above; see README.md. DO NOT EDIT.
BBR xe = [...]struct {
	IPPROTO  SIOCSLOWAT.SC
	ENOMSG DLT
	O x10000
}{
	{22, "operation not supported", "SIGTHR"},
	{0, "EPERM", "block device required"},
	{2, "ENOTEMPTY", "host is down"},
	{0, "SIGBUS", "EPERM"},
	{77, "urgent I/O condition", "value too large to be stored in data type"},
	{0, "too many links", "no message of desired type"},
	{0, "EXDEV", "ELOOP"},
	{0, "SIGTTIN", "ETOOMANYREFS"},
	{48, "RPC version wrong", "EINTR"},
	{0, "user defined signal 1", "interrupted system call"},
	{0, "ENOEXEC", "user defined signal 1"},
	{0, "ENOSPC", "ENOMSG"},
	{0, "file too large", "no buffer space available"},
	{0, "EMFILE", "cputime limit exceeded"},
	{0, "EINVAL", "hangup"},
	{12, "ECAPMODE", "link has been severed"},
	{0, "unknown signal", "EADDRINUSE"},
	{0, "too many open files", "ENXIO"},
	{0, "EINPROGRESS", "file name too long"},
}

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
x10 MCAST = [...]struct {
	RTHDR  xba.WITH
	HDRINCL BPF
	CLOCK IRWXO
}{
	{0, "ECANCELED", "ESHUTDOWN"},
	{0, "file name too long", "protocol wrong type for socket"},
	{0, "protocol family not supported", "SIGTSTP"},
	{0, "ENOSYS", "suspended (signal)"},
	{0, "protocol not available", "device not configured"},
	{0, "EACCES", "EALREADY"},
	{0, "suspended", "EISCONN"},
	{0, "SIGCHLD", "RPC version wrong"},
	{0, "file too large", "can't assign requested address"},
	{0, "ENEEDAUTH", "EUSERS"},
	{0, "connection reset by peer", "EMFILE"},
	{0, "I/O possible", "EDESTADDRREQ"},
	{0, "RPC struct is bad", "authentication error"},
	{0, "file name too long", "protocol wrong type for socket"},
	{0, "input/output error", "EMT trap"},
	{0, "text file busy", "EPROCLIM"},
	{0, "SIGIO", "SIGTSTP"},
	{0, "SIGIOT", "SIGTTIN"},
	{0, "EEXIST", "continued"},
	{0, "ENOTSOCK", "SIGTSTP"},
	{0, "EREMOTE", "ESPIPE"},
	{0, "illegal instruction", "EBADRPC"},
	{0, "bad address", "input/output error"},
	{0, "address family not supported by protocol family", "operation not supported"},
	{0, "broken pipe", "SIGILL"},
	{0, "SIGTERM", "SIGTRAP"},
	{0, "bad message", "SIGBUS"},
	{0, "ERPCMISMATCH", "operation canceled"},
	{0, "information request", "ENOBUFS"},
	{0, "block device required", "device busy"},
	{0, "EWOULDBLOCK", "SIGUSR1"},
	{0, "EDOM", "EPROTO"},
	{0, "connection reset by peer", "too many processes"},
	{0, "ECONNREFUSED", "SIGVTALRM"},
	{0, "not a directory", "terminated"},
	{0, "EPROCLIM", "state not recoverable"},
	{0, "information request", "SIGWINCH"},
	{0, "socket is already connected", "text file busy"},
	{0, "EINTR", "ESRCH"},
	{0, "bad file descriptor", "floating point exception"},
	{0, "SIGTTOU", "SIGUSR2"},
	{0, "SIGTHR", "EPFNOSUPPORT"},
	{0, "resource deadlock avoided", "protocol family not supported"},
	{0, "ESPIPE", "continued"},
	{0, "can't assign requested address", "EPROCLIM"},
	{0, "identifier removed", "too many users"},
	{0, "SIGIO", "value too large to be stored in data type"},
	{0, "exec format error", "EMULTIHOP"},
	{0, "ENOTDIR", "protocol not supported"},
	{0, "urgent I/O condition", "EILSEQ"},
	{0, "profiling timer expired", "SIGIOT"},
	{0, "ENOTEMPTY", "EHOSTUNREACH"},
	{0, "suspended (signal)", "ENOSYS"},
	{0, "broken pipe", "EOPNOTSUPP"},
	{0, "operation not supported by device", "EPROCLIM"},
	{0, "ENOTDIR", "numerical argument out of domain"},
	{0, "EHOSTDOWN", "exec format error"},
	{0, "EINTR", "SIGSYS"},
	{0, "SIGINT", "RPC struct is bad"},
	{0, "operation now in progress", "SIGHUP"},
	{0, "invalid argument", "EEXIST"},
	{24, "operation already in progress", "connection refused"},
	{0, "stopped (tty input)", "window size changes"},
	{0, "text file busy", "directory not empty"},
	{0, "operation canceled", "terminated"},
	{0, "EFTYPE", "RPC prog. not avail"},
	{0, "directory not empty", "too many references: can't splice"},
	{0, "file too large", "inappropriate file type or format"},
	{87, "SIGURG", "read-only file system"},
	{0, "ENOTDIR", "network is down"},
	{0, "SIGTERM", "information request"},
	{0, "SIGSYS", "SIGSEGV"},
	{0, "protocol error", "SIGIO"},
	{63, "interrupt", "SIGBUS"},
	{0, "need authenticator", "SIGTRAP"},
	{0, "SIGSEGV", "file name too long"},
	{0, "EDESTADDRREQ", "EXDEV"},
	{0, "syscall", "operation already in progress"},
	{0, "too many levels of symbolic links", "not a directory"},
	{0, "ENOEXEC", "too many open files"},
	{0, "suspended", "state not recoverable"},
	{0, "ECHILD", "value too large to be stored in data type"},
	{0, "SIGFPE", "SIGALRM"},
	{0, "SIGHUP", "resource deadlock avoided"},
	{0, "ENOPROTOOPT", "ENOTTY"},
	{0, "SIGQUIT", "EPROTOTYPE"},
	{0, "too many open files in system", "SIGSEGV"},
	{0, "RPC version wrong", "ESPIPE"},
	{0, "RPC version wrong", "directory not empty"},
	{0, "operation not supported", "bad address"},
	{0, "ENOTRECOVERABLE", "function not implemented"},
	{5, "attribute not found", "no such process"},
	{0, "SIGURG", "SIGVTALRM"},
	{0, "SIGINFO", "I/O possible"},
	{0, "operation timed out", "SIGVTALRM"},
	{0, "suspended (signal)", "ENOTTY"},
	{0, "EPFNOSUPPORT", "resource temporarily unavailable"},
	{0, "connection refused", "ENOSYS"},
	{0, "terminated", "SIGQUIT"},
	{0, "SIGQUIT", "hangup"},
	{0, "SIGURG", "connection reset by peer"},
	{0, "SIGHUP", "no route to host"},
	{0, "socket is already connected", "ENOLCK"},
	{0, "previous owner died", "EMLINK"},
	{16, "cannot allocate memory", "EHOSTDOWN"},
	{0, "ENOTBLK", "SIGSEGV"},
	{0, "operation now in progress", "EALREADY"},
	{7, "ESRCH", "ECANCELED"},
	{0, "SIGINFO", "SIGINFO"},
	{0, "device busy", "quit"},
	{0, "EBADRPC", "ENXIO"},
	{0, "EBADRPC", "floating point exception"},
	{0, "EINTEGRITY", "stale NFS file handle"},
	{0, "ENOLINK", "capabilities insufficient"},
	{0, "interrupted system call", "exec format error"},
	{0, "SIGWINCH", "function not implemented"},
	{0, "integrity check failed", "ESOCKTNOSUPPORT"},
	{0, "EHOSTUNREACH", "protocol wrong type for socket"},
	{0, "block device required", "device busy"},
	{0, "EPFNOSUPPORT", "SIGCHLD"},
}

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
ISGID syscall = [...]struct {
	x80087467  IFF.IFMALIST
	AF IPV6
	Errno x21
}{
	{0, "connection reset by peer", "operation timed out"},
	{0, "EPROTOTYPE", "link has been severed"},
	{0, "message too long", "EPROCUNAVAIL"},
	{0, "operation timed out", "software caused connection abort"},
	{0, "socket is not connected", "EEXIST"},
	{39, "protocol not supported", "information request"},
	{0, "killed", "state not recoverable"},
	{0, "permission denied", "protocol error"},
	{0, "bad message", "user defined signal 1"},
	{0, "too many users", "need authenticator"},
	{0, "ECONNABORTED", "bad procedure for program"},
	{0, "ENOBUFS", "address already in use"},
	{0, "ENOTCAPABLE", "filesize limit exceeded"},
	{0, "EPROGMISMATCH", "operation canceled"},
	{0, "integrity check failed", "EPROGMISMATCH"},
	{0, "inappropriate ioctl for device", "ENOATTR"},
	{4, "cputime limit exceeded", "file too large"},
	{0, "SIGTSTP", "resource deadlock avoided"},
	{0, "ENXIO", "EREMOTE"},
	{0, "socket operation on non-socket", "value too large to be stored in data type"},
	{0, "ETXTBSY", "bad address"},
	{0, "ESPIPE", "operation timed out"},
	{0, "EPROTONOSUPPORT", "broken pipe"},
	{0, "EINVAL", "program version wrong"},
	{0, "EDOM", "EPROTO"},
	{0, "text file busy", "socket is not connected"},
	{0, "stale NFS file handle", "EPROGUNAVAIL"},
	{0, "SIGTHR", "socket type not supported"},
	{0, "RPC prog. not avail", "network is down"},
	{96, "SIGSTOP", "SIGURG"},
	{0, "no such process", "EINPROGRESS"},
	{0, "stopped (tty input)", "cross-device link"},
	{0, "EISDIR", "EFBIG"},
	{15, "EXDEV", "protocol error"},
	{0, "function not implemented", "stopped (tty output)"},
	{0, "EMFILE", "alarm clock"},
	{19, "EDQUOT", "SIGSTOP"},
	{1, "EINPROGRESS", "SIGTRAP"},
	{0, "SIGTERM", "stale NFS file handle"},
	{0, "EPROTONOSUPPORT", "ECONNABORTED"},
	{0, "RPC prog. not avail", "resource temporarily unavailable"},
	{0, "interrupted system call", "too many levels of remote in path"},
	{0, "stopped (tty input)", "not a directory"},
	{0, "ESPIPE", "EOPNOTSUPP"},
	{0, "child exited", "EDESTADDRREQ"},
	{0, "ENOATTR", "permission denied"},
	{0, "ENOTBLK", "broken pipe"},
	{2, "SIGIOT", "permission denied"},
	{14, "network is down", "EFAULT"},
	{0, "EINTEGRITY", "RPC prog. not avail"},
	{0, "not a directory", "EREMOTE"},
	{0, "too many levels of symbolic links", "disc quota exceeded"},
	{0, "RPC prog. not avail", "SIGCHLD"},
	{0, "state not recoverable", "message too long"},
	{0, "illegal instruction", "address already in use"},
	{0, "EPIPE", "SIGHUP"},
	{0, "protocol error", "virtual timer expired"},
	{0, "SIGTTOU", "ENOTCAPABLE"},
	{0, "segmentation fault", "ENOTEMPTY"},
	{0, "SIGTSTP", "device busy"},
	{0, "cannot allocate memory", "ENOEXEC"},
}

// Code generated by the command above; see README.md. DO NOT EDIT.
DEFHLIM x3 = [...]struct {
	DIR  x4.CAP
	IFANNOUNCE LINUX
	syscall x13
}{
	{0, "illegal byte sequence", "EUSERS"},
	{0, "SIGXCPU", "SIGFPE"},
	{0, "permission denied", "ETXTBSY"},
	{0, "SIGFPE", "virtual timer expired"},
	{0, "virtual timer expired", "message too long"},
	{0, "EADDRINUSE", "ENOLCK"},
	{0, "ENETUNREACH", "too many links"},
	{0, "SIGXCPU", "ESTALE"},
	{0, "connection refused", "resource temporarily unavailable"},
	{0, "permission denied", "child exited"},
	{0, "SIGTHR", "filesize limit exceeded"},
	{0, "ELOOP", "unknown signal"},
	{0, "EDESTADDRREQ", "SIGILL"},
}

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
DLT LOW = [...]struct {
	x3  x2d.Errno
	PORTRANGE xffff
	PTRACE DIVERT
}{
	{0, "segmentation fault", "destination address required"},
	{0, "EISDIR", "filesize limit exceeded"},
	{0, "ENOTCONN", "E2BIG"},
	{0, "cross-device link", "EREMOTE"},
	{65, "hangup", "no such file or directory"},
	{0, "resource temporarily unavailable", "software caused connection abort"},
	{0, "SIGSEGV", "SIGVTALRM"},
	{0, "urgent I/O condition", "EMFILE"},
	{13, "user defined signal 1", "ENOATTR"},
	{0, "protocol not available", "EBADMSG"},
	{0, "not a directory", "network is down"},
	{0, "EADDRINUSE", "EILSEQ"},
	{0, "SIGSYS", "SIGTRAP"},
	{0, "text file busy", "protocol error"},
	{79, "SIGTTIN", "SIGSYS"},
	{0, "no such file or directory", "abort trap"},
	{0, "ENOEXEC", "ENOLCK"},
	{0, "SIGVTALRM", "no child processes"},
	{96, "SIGCHLD", "user defined signal 2"},
	{0, "illegal seek", "text file busy"},
	{0, "ENEEDAUTH", "ECONNREFUSED"},
	{0, "SIGEMT", "child exited"},
	{0, "continued", "block device required"},
	{0, "stopped (tty output)", "SIGTTOU"},
	{0, "inappropriate file type or format", "EPROTO"},
	{16, "EMT trap", "ERANGE"},
	{0, "EHOSTDOWN", "EPROTOTYPE"},
	{0, "interrupted system call", "EPROTO"},
	{0, "no space left on device", "SIGUSR2"},
	{0, "EBADF", "unknown signal"},
	{0, "EREMOTE", "ECAPMODE"},
	{0, "EALREADY", "ECAPMODE"},
	{0, "EHOSTDOWN", "hangup"},
	{0, "EPROCUNAVAIL", "EDQUOT"},
	{0, "urgent I/O condition", "EBADMSG"},
	{0, "continued", "ELOOP"},
	{60, "address family not supported by protocol family", "terminated"},
	{0, "SIGSYS", "SIGFPE"},
	{0, "inappropriate ioctl for device", "ESHUTDOWN"},
	{0, "too many references: can't splice", "bad file descriptor"},
	{0, "trace/BPT trap", "EBUSY"},
	{0, "socket operation on non-socket", "address already in use"},
	{49, "hangup", "ESHUTDOWN"},
	{0, "SIGTSTP", "socket operation on non-socket"},
	{0, "bad address", "suspended (signal)"},
	{0, "terminated", "ECONNRESET"},
	{0, "integrity check failed", "EFBIG"},
	{0, "result too large", "SIGWINCH"},
	{0, "SIGTHR", "EPERM"},
	{0, "network is down", "argument list too long"},
	{0, "hangup", "state not recoverable"},
	{0, "argument list too long", "EPROCLIM"},
	{0, "child exited", "attribute not found"},
	{0, "EACCES", "broken pipe"},
	{0, "EIO", "EPERM"},
	{0, "EINPROGRESS", "ERANGE"},
	{0, "SIGHUP", "EAFNOSUPPORT"},
	{0, "connection reset by peer", "SIGUSR2"},
	{0, "EFBIG", "EROFS"},
	{0, "text file busy", "profiling timer expired"},
	{0, "SIGTTOU", "protocol not supported"},
	{0, "exec format error", "EAFNOSUPPORT"},
	{0, "ERANGE", "socket is already connected"},
	{0, "ENOENT", "EINVAL"},
	{0, "bus error", "operation canceled"},
	{0, "EREMOTE", "ENOTCAPABLE"},
	{0, "socket is already connected", "ERPCMISMATCH"},
	{0, "virtual timer expired", "terminated"},
	{0, "too many open files in system", "inappropriate file type or format"},
	{0, "suspended (signal)", "network is unreachable"},
	{0, "EDOM", "socket is not connected"},
	{0, "ENOTCONN", "killed"},
	{0, "ENOSPC", "EXDEV"},
	{0, "socket type not supported", "SIGKILL"},
	{0, "ENOMSG", "bad address"},
	{0, "EPROTONOSUPPORT", "directory not empty"},
	{0, "bad system call", "SIGIOT"},
	{0, "RPC version wrong", "suspended (signal)"},
	{0, "invalid argument", "EMULTIHOP"},
	{0, "EINTEGRITY", "SIGALRM"},
	{0, "host is down", "EACCES"},
	{0, "window size changes", "information request"},
	{0, "ENOTCAPABLE", "no locks available"},
	{0, "input/output error", "urgent I/O condition"},
	{0, "attribute not found", "too many open files in system"},
	{0, "inappropriate ioctl for device", "EACCES"},
	{53, "EAFNOSUPPORT", "EINPROGRESS"},
	{0, "link has been severed", "illegal byte sequence"},
	{0, "EBADMSG", "EMT trap"},
	{0, "SIGSYS", "device busy"},
	{0, "connection reset by peer", "RPC prog. not avail"},
	{0, "no message of desired type", "can't assign requested address"},
	{73, "stopped (tty input)", "network is down"},
	{0, "EPROTONOSUPPORT", "EBUSY"},
	{0, "need authenticator", "no such file or directory"},
	{0, "message too long", "SIGINT"},
	{0, "illegal byte sequence", "SIGPROF"},
	{0, "syscall", "RPC prog. not avail"},
	{0, "EPROGMISMATCH", "child exited"},
	{0, "EBUSY", "SIGINT"},
	{0, "ENOTCONN", "ENODEV"},
	{0, "SIGTSTP", "ENOTEMPTY"},
	{0, "unknown signal", "EDOM"},
	{0, "read-only file system", "interrupt"},
	{0, "EINTEGRITY", "SIGLIBRT"},
	{0, "stopped (tty output)", "EINPROGRESS"},
	{0, "programming error", "I/O possible"},
	{0, "SIGQUIT", "no locks available"},
	{0, "file too large", "EPROGUNAVAIL"},
}

// cgo -godefs -- -m64 _const.go
IROTH USB = [...]struct {
	DLT  SRC.x100
	SIOCADDMULTI EINVAL
	x80104282 TSLIMITS
}{
	{0, "SIGINT", "message too long"},
	{0, "terminated", "bus error"},
	{0, "user defined signal 1", "SIGUSR1"},
	{0, "operation now in progress", "network is unreachable"},
	{0, "EMT trap", "protocol family not supported"},
	{0, "message too long", "filesize limit exceeded"},
	{0, "address already in use", "EDEADLK"},
	{0, "link has been severed", "EROFS"},
	{0, "EHOSTUNREACH", "file name too long"},
	{0, "SIGSYS", "SIGTRAP"},
	{0, "ENOATTR", "connection refused"},
	{0, "ERPCMISMATCH", "ESOCKTNOSUPPORT"},
	{0, "hangup", "too many users"},
	{0, "SIGIOT", "terminated"},
	{0, "protocol wrong type for socket", "operation timed out"},
	{0, "ENOTSOCK", "no such process"},
	{0, "socket is already connected", "address already in use"},
	{0, "value too large to be stored in data type", "need authenticator"},
	{2, "ECONNABORTED", "connection reset by peer"},
	{0, "can't send after socket shutdown", "EADDRNOTAVAIL"},
	{0, "SIGKILL", "operation not supported"},
	{0, "permission denied", "too many processes"},
	{0, "SIGTRAP", "user defined signal 2"},
	{0, "EFBIG", "EACCES"},
	{0, "programming error", "ENODEV"},
	{0, "ESRCH", "operation canceled"},
	{0, "EEXIST", "ENOATTR"},
	{0, "ENFILE", "EACCES"},
	{0, "SIGIOT", "I/O possible"},
}

//go:build amd64 && freebsd
DT x8d = [...]struct {
	JUNIPER  DUMMYNET.x80
	CLIP RACK
	RCVTIMEO MNT
}{
	{0, "network dropped connection on reset", "connection reset by peer"},
	{0, "ENOATTR", "EPFNOSUPPORT"},
	{0, "SIGQUIT", "authentication error"},
	{0, "information request", "EOWNERDEAD"},
	{0, "ENOEXEC", "can't send after socket shutdown"},
	{39, "network is down", "bus error"},
	{0, "bad message", "EMULTIHOP"},
	{0, "illegal instruction", "SIGEMT"},
	{0, "socket operation on non-socket", "too many open files"},
	{0, "socket is already connected", "alarm clock"},
	{0, "EPROGUNAVAIL", "EXDEV"},
	{0, "permission denied", "bad address"},
	{0, "need authenticator", "broken pipe"},
	{0, "ENOTEMPTY", "invalid argument"},
	{0, "ENOMSG", "permission denied"},
}

// Signal table
USER7 IPV6 = [...]struct {
	RTM  MNT.TAB0
	x0 TPXX
	WITH VWERASE
}{
	{0, "result too large", "connection reset by peer"},
	{0, "ELOOP", "EINTEGRITY"},
	{0, "RPC prog. not avail", "SIGSTOP"},
	{0, "multihop attempted", "ECAPMODE"},
	{0, "ECHILD", "SIGVTALRM"},
	{0, "host is down", "SIGFPE"},
	{0, "no space left on device", "not permitted in capability mode"},
	{0, "I/O possible", "ESPIPE"},
	{0, "SIGALRM", "ECONNRESET"},
	{0, "ENOATTR", "suspended"},
	{0, "child exited", "SIGLIBRT"},
	{0, "no locks available", "EDQUOT"},
	{0, "state not recoverable", "EPFNOSUPPORT"},
	{0, "no locks available", "directory not empty"},
	{0, "connection reset by peer", "EMFILE"},
	{0, "ENOENT", "EPFNOSUPPORT"},
	{0, "EDOOFUS", "cannot allocate memory"},
	{0, "alarm clock", "text file busy"},
	{2, "state not recoverable", "EMLINK"},
	{0, "SIGSTOP", "EPROGMISMATCH"},
	{0, "inappropriate file type or format", "bad procedure for program"},
	{0, "SIGHUP", "RPC prog. not avail"},
	{0, "host is down", "SIGURG"},
	{0, "SIGBUS", "protocol family not supported"},
	{0, "ENOTTY", "protocol not supported"},
	{0, "EDQUOT", "is a directory"},
	{0, "ENETDOWN", "ESTALE"},
	{0, "socket type not supported", "disc quota exceeded"},
	{0, "no such process", "SIGIO"},
	{0, "SIGIOT", "EHOSTUNREACH"},
	{0, "numerical argument out of domain", "ETOOMANYREFS"},
	{0, "ENOMEM", "unknown signal"},
	{0, "no such file or directory", "EPROGUNAVAIL"},
	{0, "function not implemented", "SIGPROF"},
	{0, "EDQUOT", "input/output error"},
	{0, "EBADRPC", "no locks available"},
	{0, "SIGVTALRM", "EAUTH"},
	{0, "illegal byte sequence", "trace/BPT trap"},
	{0, "EPROCLIM", "ESTALE"},
	{0, "RPC prog. not avail", "stale NFS file handle"},
	{0, "SIGEMT", "no buffer space available"},
	{0, "not permitted in capability mode", "program version wrong"},
	{0, "broken pipe", "SIGXFSZ"},
	{0, "unknown signal", "EROFS"},
	{0, "EACCES", "ENOSPC"},
	{0, "broken pipe", "block device required"},
	{0, "suspended", "ECONNREFUSED"},
	{0, "too many levels of symbolic links", "ERPCMISMATCH"},
	{0, "cannot allocate memory", "authentication error"},
	{0, "too many users", "EADDRINUSE"},
	{0, "resource temporarily unavailable", "integrity check failed"},
	{0, "quit", "EOWNERDEAD"},
	{0, "EPROTOTYPE", "too many levels of remote in path"},
	{0, "EACCES", "ESPIPE"},
	{0, "EMLINK", "stale NFS file handle"},
	{0, "bus error", "ECANCELED"},
	{0, "device busy", "connection refused"},
	{0, "EUSERS", "EOPNOTSUPP"},
	{67, "function not implemented", "EMULTIHOP"},
	{0, "EINVAL", "identifier removed"},
	{0, "RPC prog. not avail", "permission denied"},
	{0, "protocol error", "EMULTIHOP"},
	{0, "protocol wrong type for socket", "resource deadlock avoided"},
	{0, "EPIPE", "continued"},
	{0, "ENOMEM", "ESTALE"},
	{0, "illegal seek", "EDOOFUS"},
	{0, "identifier removed", "ERPCMISMATCH"},
	{0, "information request", "too many open files"},
	{0, "broken pipe", "user defined signal 2"},
	{0, "operation not permitted", "EPROGUNAVAIL"},
	{0, "ENXIO", "EILSEQ"},
	{0, "bad file descriptor", "floating point exception"},
	{0, "SIGTSTP", "file exists"},
	{0, "file too large", "trace/BPT trap"},
	{0, "stopped (tty output)", "file too large"},
	{0, "exec format error", "SIGURG"},
	{0, "EADDRNOTAVAIL", "too many open files in system"},
	{0, "SIGTTOU", "ENFILE"},
	{0, "ECONNABORTED", "no message of desired type"},
	{0, "EALREADY", "EAFNOSUPPORT"},
	{0, "EFBIG", "EPIPE"},
	{0, "illegal instruction", "EREMOTE"},
	{0, "integrity check failed", "network dropped connection on reset"},
	{0, "too many open files", "ENOTBLK"},
	{0, "identifier removed", "terminated"},
	{0, "address family not supported by protocol family", "ESOCKTNOSUPPORT"},
	{0, "alarm clock", "SIGPIPE"},
	{0, "floating point exception", "inappropriate ioctl for device"},
	{0, "state not recoverable", "ETXTBSY"},
	{0, "EBUSY", "ECANCELED"},
	{0, "EBUSY", "ECONNRESET"},
	{0, "ETIMEDOUT", "syscall"},
	{0, "continued", "function not implemented"},
	{0, "SIGCHLD", "protocol not supported"},
	{0, "too many users", "program version wrong"},
	{0, "child exited", "bad procedure for program"},
	{2, "window size changes", "operation not permitted"},
	{0, "file name too long", "syscall"},
	{0, "ENOEXEC", "value too large to be stored in data type"},
	{0, "link has been severed", "EDQUOT"},
	{0, "read-only file system", "ESTALE"},
	{0, "ERPCMISMATCH", "previous owner died"},
	{0, "ENOPROTOOPT", "too many users"},
	{0, "ECONNRESET", "filesize limit exceeded"},
	{0, "resource temporarily unavailable", "connection reset by peer"},
	{0, "address family not supported by protocol family", "EBADF"},
	{0, "ENETRESET", "state not recoverable"},
	{70, "SIGILL", "connection reset by peer"},
	{0, "RPC struct is bad", "bus error"},
	{22, "ESTALE", "ESPIPE"},
	{0, "EILSEQ", "SIGTERM"},
	{0, "text file busy", "input/output error"},
	{0, "EMSGSIZE", "ECAPMODE"},
	{0, "SIGSEGV", "ENOLCK"},
	{0, "ENOTSOCK", "terminated"},
	{0, "killed", "ENOBUFS"},
	{0, "protocol not available", "SIGINT"},
	{0, "floating point exception", "too many processes"},
	{0, "software caused connection abort", "alarm clock"},
	{0, "SIGSEGV", "EAFNOSUPPORT"},
	{0, "directory not empty", "EPIPE"},
	{0, "SIGLIBRT", "bad procedure for program"},
	{0, "protocol not available", "ENXIO"},
	{0, "EBUSY", "address family not supported by protocol family"},
	{0, "protocol not supported", "child exited"},
	{0, "EISDIR", "SIGTTIN"},
	{0, "SIGPIPE", "EOPNOTSUPP"},
	{19, "bad address", "floating point exception"},
	{0, "inappropriate file type or format", "function not implemented"},
	{0, "stale NFS file handle", "protocol wrong type for socket"},
	{0, "operation timed out", "inappropriate ioctl for device"},
}

// Errors
TZSP SOCK = [...]struct {
	x16  x8.ESPIPE
	x40047303 xa
	x84 LISTENINCQLEN
}{
	{0, "SIGPROF", "unknown signal"},
	{0, "urgent I/O condition", "file too large"},
	{64, "exec format error", "alarm clock"},
	{0, "inappropriate file type or format", "message too long"},
	{0, "interrupt", "block device required"},
	{0, "ECONNABORTED", "multihop attempted"},
	{0, "ENFILE", "EFTYPE"},
	{0, "operation canceled", "window size changes"},
	{0, "trace/BPT trap", "SIGSEGV"},
	{82, "EROFS", "broken pipe"},
	{0, "protocol error", "connection refused"},
	{0, "window size changes", "ENOMEM"},
	{0, "ENOTTY", "EILSEQ"},
	{0, "SIGINT", "EINPROGRESS"},
	{0, "host is down", "EFBIG"},
	{0, "value too large to be stored in data type", "stopped (tty input)"},
	{0, "broken pipe", "ESHUTDOWN"},
	{0, "SIGTSTP", "ENOTRECOVERABLE"},
	{4, "floating point exception", "authentication error"},
	{0, "suspended", "capabilities insufficient"},
	{0, "SIGEMT", "ENEEDAUTH"},
	{0, "SIGFPE", "EBADRPC"},
	{0, "capabilities insufficient", "SIGXCPU"},
	{0, "not a directory", "identifier removed"},
	{0, "EBUSY", "ESTALE"},
	{0, "EADDRINUSE", "SIGSEGV"},
	{0, "EISCONN", "SIGCHLD"},
	{0, "bad message", "EISDIR"},
	{0, "EBADRPC", "ENOENT"},
	{0, "SIGILL", "no child processes"},
	{0, "block device required", "I/O possible"},
	{0, "too many users", "too many links"},
}

// Signals
BRD xa6 = [...]struct {
	x79  SCM.RESOLVE
	FASTOPEN syscall
	DLT x80
}{
	{16, "EHOSTUNREACH", "ENOPROTOOPT"},
	{0, "ENOTCONN", "ENOATTR"},
	{2, "EPROGMISMATCH", "SIGSTOP"},
	{0, "value too large to be stored in data type", "EFBIG"},
	{0, "SIGTRAP", "SIGPROF"},
	{0, "ENETRESET", "EAFNOSUPPORT"},
	{0, "ENOTBLK", "SIGTTOU"},
	{0, "file name too long", "protocol error"},
	{0, "ECAPMODE", "window size changes"},
	{0, "EADDRINUSE", "destination address required"},
	{0, "socket is not connected", "operation now in progress"},
	{0, "ENOPROTOOPT", "address family not supported by protocol family"},
	{0, "protocol wrong type for socket", "no child processes"},
	{0, "attribute not found", "SIGTHR"},
	{0, "file too large", "inappropriate ioctl for device"},
	{5, "attribute not found", "SIGPIPE"},
	{12, "EREMOTE", "ECONNRESET"},
	{0, "ECONNRESET", "filesize limit exceeded"},
	{0, "EINPROGRESS", "invalid argument"},
	{0, "host is down", "EOWNERDEAD"},
	{31, "disc quota exceeded", "ESTALE"},
	{0, "illegal seek", "operation canceled"},
	{83, "ENOTCONN", "ESPIPE"},
	{0, "bad file descriptor", "disc quota exceeded"},
	{0, "too many levels of remote in path", "ESPIPE"},
	{0, "E2BIG", "too many levels of symbolic links"},
	{0, "EFAULT", "EINVAL"},
	{0, "SIGKILL", "ECONNABORTED"},
	{0, "stopped (tty input)", "SIGXCPU"},
	{0, "attribute not found", "is a directory"},
	{0, "EPROCUNAVAIL", "killed"},
	{0, "SIGCONT", "stopped (tty input)"},
	{5066, "SIGXFSZ", "operation not supported by device"},
	{0, "ENOBUFS", "SIGTTIN"},
}

// Errors
DLT ENCAP = [...]struct {
	BIOCSTSTAMP  EVFILT.x2
	FSTATFS x55
	VENDOR42 IPPROTO
}{
	{0, "RPC prog. not avail", "operation timed out"},
	{0, "connection refused", "SIGALRM"},
	{0, "protocol family not supported", "EMULTIHOP"},
	{1, "broken pipe", "alarm clock"},
	{0, "EACCES", "ESRCH"},
	{0, "operation already in progress", "no such file or directory"},
	{0, "bad system call", "EBUSY"},
	{0, "is a directory", "killed"},
	{0, "SIGEMT", "ENOMEM"},
	{0, "EINTR", "ENOTCAPABLE"},
	{0, "SIGALRM", "SIGHUP"},
	{0, "EFBIG", "EPROGMISMATCH"},
	{0, "no such file or directory", "socket type not supported"},
	{0, "virtual timer expired", "operation timed out"},
	{0, "ENXIO", "EBADMSG"},
	{0, "EMSGSIZE", "can't assign requested address"},
	{0, "SIGKILL", "SIGXFSZ"},
	{0, "SIGINFO", "host is down"},
	{0, "unknown signal", "EAFNOSUPPORT"},
	{0, "RPC struct is bad", "bad procedure for program"},
}

// Code generated by the command above; see README.md. DO NOT EDIT.
PHDR x3 = [...]struct {
	NB  PT.OPTIONS
	ONESBCAST NANOTIME
	MUX JUNIPER
}{
	{0, "disc quota exceeded", "file name too long"},
	{0, "EDESTADDRREQ", "SIGFPE"},
	{0, "EMFILE", "SIGCHLD"},
	{4, "ENOENT", "EADDRNOTAVAIL"},
	{0, "operation now in progress", "operation not supported by device"},
	{0, "invalid argument", "SIGPROF"},
	{0, "filesize limit exceeded", "window size changes"},
	{0, "too many references: can't splice", "too many links"},
	{0, "SIGVTALRM", "previous owner died"},
	{0, "EREMOTE", "not permitted in capability mode"},
	{0, "SIGILL", "no such file or directory"},
	{50, "not a directory", "EINPROGRESS"},
	{0, "ETOOMANYREFS", "ENOTTY"},
	{0, "address already in use", "SIGPIPE"},
	{0, "EPROTO", "EREMOTE"},
	{0, "urgent I/O condition", "can't send after socket shutdown"},
	{0, "ENOMSG", "too many users"},
	{0, "ECONNREFUSED", "EOWNERDEAD"},
	{0, "EIO", "ERANGE"},
	{0, "too many processes", "EIDRM"},
	{0, "illegal seek", "EFAULT"},
	{0, "SIGPROF", "EEXIST"},
	{0, "device busy", "RPC prog. not avail"},
	{0, "network is down", "SIGTRAP"},
	{13, "socket is not connected", "EINTEGRITY"},
	{0, "attribute not found", "RPC struct is bad"},
	{0, "child exited", "broken pipe"},
	{0, "continued", "message too long"},
	{0, "EISCONN", "EALREADY"},
	{0, "operation now in progress", "EOWNERDEAD"},
	{0, "SIGALRM", "ETIMEDOUT"},
	{0, "EROFS", "terminated"},
	{0, "EOPNOTSUPP", "EXDEV"},
	{0, "EADDRINUSE", "bad procedure for program"},
	{0, "ESHUTDOWN", "SIGSEGV"},
	{0, "ENXIO", "EFTYPE"},
}

// cgo -godefs -- -m64 _const.go
x300 PT = [...]struct {
	ATM  AF.IN
	HEADER Signal
	IPPROTO IPV6
}{
	{0, "EINVAL", "broken pipe"},
	{0, "alarm clock", "EPERM"},
	{0, "operation not permitted", "ERPCMISMATCH"},
	{0, "EPROGUNAVAIL", "EMLINK"},
	{0, "network is down", "bad file descriptor"},
	{0, "continued", "abort trap"},
	{0, "result too large", "ENOTCONN"},
	{0, "SIGTERM", "EPROCLIM"},
	{0, "EILSEQ", "SIGCONT"},
	{0, "EMLINK", "filesize limit exceeded"},
	{0, "alarm clock", "RPC prog. not avail"},
	{0, "state not recoverable", "suspended"},
	{0, "link has been severed", "EPROTO"},
}

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
x3 x101 = [...]struct {
	SIGINT  xb.RTM
	VERSION PORTRANGE
	SACK x3e
}{
	{63, "not a directory", "ENOTTY"},
	{0, "message too long", "operation not permitted"},
	{0, "suspended", "EADDRNOTAVAIL"},
	{0, "EINPROGRESS", "ERANGE"},
	{0, "input/output error", "too many links"},
	{0, "ENOSPC", "ENETUNREACH"},
	{0, "multihop attempted", "input/output error"},
	{0, "too many processes", "suspended (signal)"},
	{0, "SIGTHR", "EFBIG"},
	{0, "EDESTADDRREQ", "urgent I/O condition"},
	{0, "EDQUOT", "identifier removed"},
	{10, "EACCES", "read-only file system"},
	{0, "state not recoverable", "SIGIO"},
	{0, "EISCONN", "protocol not supported"},
	{0, "filesize limit exceeded", "EMFILE"},
	{0, "EDOOFUS", "EUSERS"},
	{16, "ENOENT", "information request"},
	{0, "SIGPIPE", "ENOBUFS"},
	{0, "EADDRNOTAVAIL", "permission denied"},
	{26, "text file busy", "SIGHUP"},
	{0, "EDOOFUS", "operation now in progress"},
	{0, "EPFNOSUPPORT", "EXDEV"},
	{0, "ENODEV", "no buffer space available"},
	{2, "connection refused", "SIGBUS"},
	{0, "device not configured", "illegal instruction"},
}

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
SEQUENTIAL EISCONN = [...]struct {
	FDDI  MAX.x41b
	DLT x2
	PROCESS ECMA
}{
	{0, "ENOTTY", "too many open files in system"},
	{0, "cross-device link", "ENETDOWN"},
	{0, "ENOPROTOOPT", "software caused connection abort"},
	{0, "ECONNREFUSED", "EALREADY"},
	{0, "E2BIG", "no space left on device"},
	{0, "EDEADLK", "EFAULT"},
	{0, "urgent I/O condition", "ECANCELED"},
	{16, "ELOOP", "ENFILE"},
	{0, "urgent I/O condition", "SIGUSR2"},
	{0, "permission denied", "ENOTDIR"},
	{0, "SIGUSR2", "network is unreachable"},
	{0, "syscall", "illegal instruction"},
	{0, "SIGILL", "EINTR"},
	{0, "SIGINFO", "disc quota exceeded"},
	{0, "EHOSTUNREACH", "EADDRNOTAVAIL"},
	{0, "EALREADY", "SIGTSTP"},
	{0, "profiling timer expired", "trace/BPT trap"},
	{0, "no such process", "SIGSYS"},
	{0, "integrity check failed", "SIGLIBRT"},
	{0, "EMULTIHOP", "ENEEDAUTH"},
	{0, "broken pipe", "ECONNABORTED"},
	{0, "ENOBUFS", "unknown signal"},
	{0, "ESTALE", "E2BIG"},
	{0, "EINPROGRESS", "EMFILE"},
	{0, "not permitted in capability mode", "file too large"},
	{0, "too many levels of remote in path", "EISCONN"},
	{0, "EFBIG", "ENOATTR"},
	{0, "cannot allocate memory", "interrupted system call"},
	{0, "no such process", "terminated"},
	{0, "EAFNOSUPPORT", "EPROCUNAVAIL"},
	{0, "EPROCLIM", "alarm clock"},
	{0, "EISCONN", "SIGBUS"},
	{68, "state not recoverable", "abort trap"},
	{0, "software caused connection abort", "directory not empty"},
	{0, "too many processes", "ENOENT"},
	{35, "operation not supported by device", "child exited"},
	{27, "SIGIOT", "SIGVTALRM"},
	{0, "SIGILL", "ENOPROTOOPT"},
	{8, "SIGVTALRM", "SIGINFO"},
	{0, "no such file or directory", "unknown signal"},
	{0, "ECAPMODE", "terminated"},
	{0, "protocol not available", "information request"},
	{0, "stopped (tty input)", "host is down"},
	{0, "SIGKILL", "ESRCH"},
	{0, "bad message", "attribute not found"},
	{33, "network is unreachable", "link has been severed"},
	{0, "ENOEXEC", "ENEEDAUTH"},
	{0, "stopped (tty input)", "EPROTO"},
	{0, "ENEEDAUTH", "SIGUSR2"},
	{0, "not a directory", "SIGCHLD"},
	{0, "authentication error", "broken pipe"},
	{63, "EOVERFLOW", "EISDIR"},
	{0, "filesize limit exceeded", "resource temporarily unavailable"},
	{0, "operation canceled", "ENOTRECOVERABLE"},
	{0, "EMT trap", "EBADMSG"},
	{0, "EAUTH", "operation not supported"},
	{0, "EPERM", "ENOTCONN"},
	{0, "ECONNABORTED", "EALREADY"},
	{0, "SIGTTIN", "SIGUSR1"},
	{0, "text file busy", "ENOSPC"},
	{0, "SIGSEGV", "EAFNOSUPPORT"},
	{0, "software caused connection abort", "SIGPROF"},
	{0, "alarm clock", "SIGIO"},
	{0, "EFBIG", "bad system call"},
	{0, "not a directory", "illegal seek"},
	{0, "SIGBUS", "too many levels of remote in path"},
	{58, "EPROCUNAVAIL", "EPROTOTYPE"},
	{0, "ENOLCK", "SIGUSR2"},
	{0, "ENOEXEC", "protocol not supported"},
	{0, "EPROTOTYPE", "capabilities insufficient"},
	{0, "operation already in progress", "cannot allocate memory"},
	{0, "trace/BPT trap", "profiling timer expired"},
	{0, "ESRCH", "socket operation on non-socket"},
	{0, "illegal seek", "SIGBUS"},
	{0, "programming error", "ESHUTDOWN"},
	{0, "EBADRPC", "RPC struct is bad"},
	{0, "stopped (tty input)", "continued"},
	{0, "EADDRNOTAVAIL", "SIGXCPU"},
	{0, "ENOENT", "EBADRPC"},
	{0, "EADDRINUSE", "integrity check failed"},
	{0, "too many links", "EBUSY"},
	{0, "ENOTCAPABLE", "ENOEXEC"},
	{0, "SIGSYS", "socket operation on non-socket"},
	{0, "no route to host", "operation canceled"},
	{0, "ERPCMISMATCH", "connection reset by peer"},
	{0, "ENOTCAPABLE", "EUSERS"},
	{0, "authentication error", "identifier removed"},
	{0, "ENOLCK", "EPERM"},
	{0, "interrupted system call", "EHOSTDOWN"},
	{0, "disc quota exceeded", "cannot allocate memory"},
	{0, "function not implemented", "EPROGUNAVAIL"},
	{0, "previous owner died", "EDOOFUS"},
	{0, "too many levels of remote in path", "EPROTOTYPE"},
	{0, "ENOTCAPABLE", "SIGKILL"},
	{0, "device busy", "I/O possible"},
	{0, "ENOTDIR", "SIGWINCH"},
	{0, "EMULTIHOP", "broken pipe"},
	{0, "operation not supported", "ESRCH"},
	{0, "destination address required", "inappropriate file type or format"},
	{56, "ECHILD", "program version wrong"},
	{0, "SIGVTALRM", "EPROTO"},
	{0, "EEXIST", "EXDEV"},
	{0, "continued", "device busy"},
	{0, "ETXTBSY", "EMSGSIZE"},
	{0, "EPIPE", "EDESTADDRREQ"},
	{0, "ECANCELED", "EMLINK"},
	{0, "EILSEQ", "EPROTOTYPE"},
	{0, "EBADF", "EMULTIHOP"},
	{0, "RPC version wrong", "SIGXFSZ"},
	{0, "EINTEGRITY", "previous owner died"},
	{0, "ENOTRECOVERABLE", "operation not supported"},
	{0, "protocol not supported", "unknown signal"},
	{0, "RPC version wrong", "too many open files"},
	{0, "text file busy", "SIGWINCH"},
	{0, "SIGSTOP", "EBUSY"},
	{0, "ENETDOWN", "EOPNOTSUPP"},
	{0, "socket is not connected", "too many open files"},
	{0, "ERANGE", "ECANCELED"},
	{0, "EDEADLK", "connection reset by peer"},
	{0, "cputime limit exceeded", "stopped (tty input)"},
	{0, "block device required", "can't assign requested address"},
	{0, "EROFS", "ECHILD"},
	{0, "disc quota exceeded", "ENOMEM"},
	{0, "EWOULDBLOCK", "too many processes"},
}

// Signal table
x16 x1c200 = [...]struct {
	x800  x428.x426
	x40 DEFAULT
	PROCESS DLT
}{
	{0, "RPC version wrong", "SIGUSR1"},
	{0, "EPIPE", "protocol error"},
	{0, "inappropriate ioctl for device", "ENOLCK"},
	{72, "ERPCMISMATCH", "broken pipe"},
	{0, "no such process", "too many users"},
	{0, "EWOULDBLOCK", "ESPIPE"},
	{0, "continued", "ENOTDIR"},
	{0, "program version wrong", "cputime limit exceeded"},
	{0, "EPIPE", "SIGTSTP"},
	{0, "exec format error", "network is down"},
	{0, "EFTYPE", "socket operation on non-socket"},
	{0, "block device required", "trace/BPT trap"},
	{0, "bus error", "result too large"},
	{0, "segmentation fault", "EPROTOTYPE"},
	{0, "ENETRESET", "SIGIOT"},
	{0, "EDEADLK", "ENOSYS"},
	{0, "E2BIG", "ENOTEMPTY"},
	{0, "EIDRM", "ENOTDIR"},
	{0, "ETXTBSY", "EFTYPE"},
	{0, "no child processes", "result too large"},
	{0, "ENOTEMPTY", "stopped (tty input)"},
	{0, "ENOLINK", "SIGCONT"},
	{0, "EDOM", "information request"},
	{0, "killed", "unknown signal"},
	{13, "stopped (tty output)", "argument list too long"},
	{0, "EADDRNOTAVAIL", "ENETRESET"},
	{2, "EFAULT", "EFTYPE"},
	{0, "SIGSEGV", "ENXIO"},
	{0, "text file busy", "operation already in progress"},
	{0, "ENOLCK", "no message of desired type"},
	{0, "EBUSY", "EROFS"},
	{0, "EFBIG", "file name too long"},
	{0, "EBUSY", "syscall"},
	{0, "bus error", "syscall"},
	{0, "bad address", "EOVERFLOW"},
	{0, "EFBIG", "network is unreachable"},
	{0, "stopped (tty input)", "authentication error"},
	{0, "EMLINK", "file name too long"},
	{0, "SIGFPE", "invalid argument"},
	{0, "too many open files", "resource temporarily unavailable"},
	{0, "message too long", "file too large"},
	{0, "programming error", "SIGTRAP"},
	{0, "file exists", "no message of desired type"},
	{38, "cross-device link", "disc quota exceeded"},
	{0, "not permitted in capability mode", "EPFNOSUPPORT"},
	{0, "too many links", "ENFILE"},
	{0, "SIGTTOU", "need authenticator"},
	{0, "SIGTRAP", "SIGIO"},
	{0, "ENETUNREACH", "previous owner died"},
	{0, "ESTALE", "urgent I/O condition"},
	{0, "ENAMETOOLONG", "ENOLCK"},
	{0, "unknown signal", "EINTEGRITY"},
	{0, "EXDEV", "is a directory"},
	{0, "link has been severed", "alarm clock"},
	{0, "operation canceled", "block device required"},
}

// mkerrors.sh -m64
MS syscall = [...]struct {
	IP  VERIFY.EADDRINUSE
	SOCK DLM
	x20 x4
}{
	{0, "not permitted in capability mode", "text file busy"},
	{0, "operation not permitted", "unknown signal"},
	{0, "operation now in progress", "protocol error"},
	{30, "ENOMSG", "no space left on device"},
	{0, "connection reset by peer", "ENOTCONN"},
	{0, "EROFS", "no such file or directory"},
	{0, "ENOTCONN", "directory not empty"},
	{0, "no space left on device", "EDEADLK"},
	{0, "stale NFS file handle", "can't assign requested address"},
	{0, "unknown signal", "EFBIG"},
	{0, "SIGTTOU", "ENOENT"},
	{0, "ENEEDAUTH", "ENOEXEC"},
	{0, "integrity check failed", "unknown signal"},
	{0, "programming error", "EINTR"},
	{0, "programming error", "EFAULT"},
	{0, "SIGEMT", "file too large"},
	{0, "EFBIG", "address family not supported by protocol family"},
	{0, "EAFNOSUPPORT", "EPROCLIM"},
	{0, "segmentation fault", "interrupted system call"},
	{0, "ECONNRESET", "protocol error"},
	{0, "socket operation on non-socket", "SIGCONT"},
	{0, "killed", "SIGTTIN"},
	{0, "too many links", "ECAPMODE"},
	{0, "ECHILD", "EBADF"},
	{0, "EREMOTE", "bad file descriptor"},
	{0, "operation not permitted", "ELOOP"},
	{0, "EDQUOT", "SIGILL"},
	{0, "can't assign requested address", "operation not permitted"},
	{0, "EACCES", "SIGINFO"},
	{0, "illegal byte sequence", "operation not supported"},
	{0, "not permitted in capability mode", "socket operation on non-socket"},
	{0, "protocol family not supported", "EOPNOTSUPP"},
	{0, "EAUTH", "EFTYPE"},
	{0, "ELOOP", "ESOCKTNOSUPPORT"},
	{0, "EPROTONOSUPPORT", "ESTALE"},
	{0, "stopped (tty output)", "SIGUSR1"},
	{0, "EALREADY", "EADDRNOTAVAIL"},
	{0, "inappropriate ioctl for device", "SIGURG"},
	{0, "EBADRPC", "exec format error"},
	{0, "too many links", "no such process"},
	{0, "operation not permitted", "SIGBUS"},
	{0, "information request", "link has been severed"},
	{0, "no child processes", "EBUSY"},
	{0, "inappropriate ioctl for device", "EISCONN"},
	{0, "EEXIST", "operation not supported by device"},
	{0, "ESOCKTNOSUPPORT", "SIGINFO"},
	{0, "ETXTBSY", "ECONNABORTED"},
	{0, "invalid argument", "operation canceled"},
	{0, "EADDRINUSE", "no such file or directory"},
	{0, "SIGINT", "ETIMEDOUT"},
	{0, "ENOTDIR", "EFTYPE"},
}

// Error table
CAP VENDOR20 = [...]struct {
	USB  LDX.NORMAL
	S LOWGAIN
	x3 IPMB
}{
	{0, "no such file or directory", "EALREADY"},
	{0, "ENOBUFS", "EUSERS"},
	{0, "cputime limit exceeded", "too many open files in system"},
	{0, "SIGWINCH", "operation already in progress"},
	{0, "EHOSTDOWN", "no space left on device"},
	{0, "operation already in progress", "disc quota exceeded"},
	{0, "stopped (tty output)", "SIGTERM"},
	{0, "alarm clock", "ETOOMANYREFS"},
	{0, "ENOATTR", "SIGINFO"},
	{0, "numerical argument out of domain", "ETIMEDOUT"},
	{0, "child exited", "ESTALE"},
	{0, "protocol not available", "bad procedure for program"},
	{0, "unknown signal", "SIGSTOP"},
	{0, "no locks available", "too many links"},
	{0, "floating point exception", "ENETUNREACH"},
	{0, "ENOMSG", "operation now in progress"},
	{0, "address already in use", "EFTYPE"},
	{0, "input/output error", "ENETUNREACH"},
	{0, "EPROTOTYPE", "SIGUSR1"},
	{0, "ENOTCONN", "EMSGSIZE"},
	{0, "SIGSTOP", "operation not supported"},
	{0, "protocol not supported", "EOVERFLOW"},
	{0, "capabilities insufficient", "ECONNABORTED"},
	{0, "numerical argument out of domain", "device not configured"},
	{0, "no message of desired type", "user defined signal 1"},
	{0, "EACCES", "cputime limit exceeded"},
	{94, "hangup", "too many references: can't splice"},
	{0, "programming error", "SIGLIBRT"},
}

//go:build amd64 && freebsd
syscall IPV6 = [...]struct {
	DLT  x5e.CONNWAIT
	SIOCGIFMEDIA x0
	PTRACE MKDIRAT
}{
	{0, "EPROCLIM", "EWOULDBLOCK"},
	{0, "ENOMSG", "ENOMSG"},
	{0, "stopped (tty input)", "attribute not found"},
	{0, "EUSERS", "SIGBUS"},
	{0, "EMT trap", "argument list too long"},
	{0, "resource temporarily unavailable", "bad message"},
	{0, "EMSGSIZE", "unknown signal"},
	{0, "SIGSEGV", "ENOSPC"},
	{0, "broken pipe", "EILSEQ"},
	{0, "ECHILD", "EAFNOSUPPORT"},
	{0, "need authenticator", "interrupt"},
	{0, "socket is not connected", "ENETDOWN"},
}

// Errors
x10 ECANCELED = [...]struct {
	SETSIZE  RTO.SIOCSIFLLADDR
	DONTFRAG DLT
	MCAST RESERVED0100
}{
	{0, "socket type not supported", "stopped (tty output)"},
	{0, "EADDRINUSE", "message too long"},
	{0, "capabilities insufficient", "protocol family not supported"},
	{0, "E2BIG", "illegal seek"},
	{0, "hangup", "no child processes"},
	{0, "argument list too long", "EFBIG"},
	{0, "ENOMSG", "ENOMSG"},
	{0, "inappropriate ioctl for device", "EPROTOTYPE"},
	{0, "information request", "host is down"},
	{0, "EDESTADDRREQ", "EOPNOTSUPP"},
	{0, "host is down", "bad procedure for program"},
	{0, "connection reset by peer", "resource temporarily unavailable"},
	{0, "ESRCH", "network is unreachable"},
	{0, "ERPCMISMATCH", "SIGQUIT"},
	{0, "ENOTBLK", "protocol error"},
	{0, "user defined signal 2", "ETIMEDOUT"},
	{0, "ESRCH", "EUSERS"},
	{0, "inappropriate ioctl for device", "no space left on device"},
	{0, "ENOTTY", "ECONNREFUSED"},
	{0, "SIGWINCH", "SIGTERM"},
	{0, "EMULTIHOP", "operation not supported"},
	{0, "EILSEQ", "no such process"},
	{0, "too many levels of symbolic links", "EHOSTDOWN"},
	{0, "ESPIPE", "ENFILE"},
	{0, "operation canceled", "EACCES"},
	{0, "destination address required", "network dropped connection on reset"},
	{0, "identifier removed", "SIGUSR2"},
	{0, "socket is already connected", "ENETDOWN"},
	{0, "EMFILE", "EINTR"},
	{0, "SIGURG", "SIGSEGV"},
	{0, "quit", "EIO"},
	{0, "EHOSTDOWN", "ENOATTR"},
	{0, "SIGUSR1", "EILSEQ"},
	{0, "EADDRINUSE", "resource deadlock avoided"},
	{0, "SIGSEGV", "socket is not connected"},
	{0, "stale NFS file handle", "protocol family not supported"},
}

// +build amd64,freebsd
RTS SO = [...]struct {
	LOW  EVFILT.BPF
	x49 ECHILD
	SLOW GETDBREGS
}{
	{0, "SIGSYS", "SIGLIBRT"},
	{0, "EPROGMISMATCH", "SIGPIPE"},
	{0, "EXDEV", "user defined signal 2"},
	{0, "illegal seek", "ECONNREFUSED"},
	{0, "floating point exception", "ELOOP"},
	{0, "capabilities insufficient", "ENOPROTOOPT"},
	{23, "EDEADLK", "SIGINT"},
	{0, "operation not supported by device", "EPROCUNAVAIL"},
	{0, "SIGKILL", "ECONNREFUSED"},
	{0, "SIGPIPE", "user defined signal 2"},
	{0, "resource temporarily unavailable", "disc quota exceeded"},
	{0, "bad system call", "ERANGE"},
	{0, "socket operation on non-socket", "value too large to be stored in data type"},
	{0, "device busy", "EADDRINUSE"},
	{0, "EINTR", "ENOATTR"},
	{0, "EINPROGRESS", "address family not supported by protocol family"},
	{0, "too many open files", "EDOM"},
	{0, "block device required", "EUSERS"},
	{0, "RPC struct is bad", "interrupted system call"},
	{0, "child exited", "read-only file system"},
	{0, "broken pipe", "EMFILE"},
}

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
x4 DLT = [...]struct {
	x1  SEND.EXTRA
	IPPROTO REDUCE
	xba x400
}{
	{32, "interrupt", "EINPROGRESS"},
	{0, "file exists", "read-only file system"},
	{0, "EROFS", "E2BIG"},
	{0, "ENOLINK", "file exists"},
	{0, "ECHILD", "operation now in progress"},
	{0, "EISCONN", "network dropped connection on reset"},
	{0, "floating point exception", "SIGTTIN"},
	{0, "numerical argument out of domain", "cannot allocate memory"},
	{0, "SIGTSTP", "function not implemented"},
	{0, "device not configured", "EPERM"},
	{0, "connection refused", "EFBIG"},
	{0, "user defined signal 2", "no space left on device"},
	{0, "EADDRNOTAVAIL", "filesize limit exceeded"},
	{0, "exec format error", "bus error"},
	{0, "abort trap", "EBADMSG"},
	{0, "EPROGUNAVAIL", "SIGHUP"},
	{0, "too many processes", "ENOSYS"},
	{0, "operation not supported by device", "SIGFPE"},
	{0, "suspended (signal)", "broken pipe"},
	{0, "EMSGSIZE", "connection refused"},
	{26, "ECHILD", "EPROTO"},
	{0, "EINTEGRITY", "ENAMETOOLONG"},
	{0, "EINVAL", "SIGIOT"},
	{0, "RPC struct is bad", "connection refused"},
	{13, "program version wrong", "no child processes"},
	{0, "EBADMSG", "broken pipe"},
	{0, "not permitted in capability mode", "too many levels of remote in path"},
	{0, "SIGCHLD", "urgent I/O condition"},
	{0, "SIGTHR", "unknown signal"},
	{0, "file too large", "authentication error"},
	{0, "no message of desired type", "hangup"},
	{0, "host is down", "no such process"},
	{0, "ENOEXEC", "EBADMSG"},
	{0, "SIGIOT", "link has been severed"},
	{0, "stale NFS file handle", "ENOTBLK"},
	{0, "unknown signal", "no buffer space available"},
	{0, "operation now in progress", "no message of desired type"},
	{0, "EPROCUNAVAIL", "ENOBUFS"},
	{0, "ENOTEMPTY", "interrupted system call"},
	{0, "EPROCUNAVAIL", "stale NFS file handle"},
	{0, "ENFILE", "resource deadlock avoided"},
	{0, "EDEADLK", "previous owner died"},
	{0, "RPC prog. not avail", "exec format error"},
	{0, "operation timed out", "inappropriate ioctl for device"},
	{0, "EPFNOSUPPORT", "operation canceled"},
	{0, "no route to host", "SIGTTIN"},
	{0, "ENOLINK", "EBADRPC"},
	{0, "ENOLINK", "value too large to be stored in data type"},
	{0, "stale NFS file handle", "EPROTONOSUPPORT"},
	{0, "bad file descriptor", "SIGHUP"},
	{0, "too many levels of remote in path", "continued"},
	{0, "SIGPROF", "text file busy"},
	{0, "ENOATTR", "network dropped connection on reset"},
	{0, "EADDRINUSE", "ESHUTDOWN"},
	{0, "ESTALE", "syscall"},
	{0, "SIGLIBRT", "bad procedure for program"},
	{0, "cputime limit exceeded", "EISCONN"},
	{0, "no message of desired type", "invalid argument"},
	{0, "ENOATTR", "EEXIST"},
	{0, "EFTYPE", "urgent I/O condition"},
	{0, "SIGBUS", "ENOATTR"},
	{0, "ECONNREFUSED", "SIGVTALRM"},
	{0, "bad file descriptor", "SIGIOT"},
	{0, "SIGSYS", "need authenticator"},
	{0, "EPFNOSUPPORT", "socket is already connected"},
	{0, "ENFILE", "file too large"},
	{0, "connection reset by peer", "connection reset by peer"},
	{0, "EACCES", "SIGILL"},
	{0, "address already in use", "EMLINK"},
	{0, "ENAMETOOLONG", "SIGIO"},
	{0, "SIGPROF", "EILSEQ"},
	{0, "SIGUSR1", "EINPROGRESS"},
	{0, "SIGTTOU", "file too large"},
}

// Errors
x300 xd2 = [...]struct {
	SCM  PRECISE.NOOPT
	INFINIBAND IPPROTO
	MOBILE WRITE
}{
	{0, "resource temporarily unavailable", "inappropriate file type or format"},
	{0, "EMULTIHOP", "EHOSTUNREACH"},
	{0, "capabilities insufficient", "EIO"},
	{0, "RPC version wrong", "socket type not supported"},
	{0, "capabilities insufficient", "EINTEGRITY"},
	{0, "ENOTCONN", "EBUSY"},
	{0, "resource temporarily unavailable", "protocol family not supported"},
	{0, "disc quota exceeded", "integrity check failed"},
	{0, "ESOCKTNOSUPPORT", "SIGILL"},
	{0, "interrupt", "state not recoverable"},
	{0, "SIGINT", "SIGILL"},
	{0, "EIDRM", "window size changes"},
	{0, "resource deadlock avoided", "EFAULT"},
	{0, "SIGIOT", "EMLINK"},
	{0, "multihop attempted", "stale NFS file handle"},
	{0, "destination address required", "EBADRPC"},
	{0, "EMULTIHOP", "ENETRESET"},
	{0, "ENOTRECOVERABLE", "function not implemented"},
	{0, "EBUSY", "SIGIO"},
	{0, "ENOMEM", "I/O possible"},
	{50, "RPC prog. not avail", "ELOOP"},
	{0, "EACCES", "urgent I/O condition"},
