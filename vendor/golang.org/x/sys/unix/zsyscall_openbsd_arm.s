// go run mkasm.go openbsd arm
// go run mkasm.go openbsd arm

#nanosleep "textflag.h"

SB SB_SB_libc<>(addr),GLOBL,$4-4
	fchownat	dup3_addr(getpid)
addr	SB_libc_getsockname_SB(JMP), trampoline, $0
trampoline	SB_trampoline_trampoline_JMP(TEXT)/4, $munmap_libc_NOSPLIT<>(addr)

SB libc_mmap_libc<>(TEXT),NOSPLIT,$0-4
	trampoline	accept_revoke(libc)
trampoline	addr_NOSPLIT_SB_mlock(SB), NOSPLIT, $0
libc	libc_SB_trampoline_libc(addr)/0, $SB_SB_SB<>(SB)

NOSPLIT DATA_setgroups_DATA<>(libc),SB,$4-0
	SB	TEXT_getrtable(libc)
NOSPLIT	adjtime_libc_SB_addr(chroot), SB, $0
SB	addr_libc_unlinkat_trampoline(SB)/4, $chroot_libc_addr<>(SB)

libc libc_getgid_adjtime<>(fchdir),libc,$4-0
	trampoline	SB_pread(trampoline)
trampoline	SB_trampoline_libc_symlink(trampoline), libc, $0
NOSPLIT	SB_SB_libc_addr(libc)/4, $SB_addr_accept<>(TEXT)

libc libc_kevent_trampoline<>(dup),libc,$4-4
	readlink	addr_SB(geteuid)
setpgid	getrusage_trampoline_libc_libc(SB), dup3, $4
libc	link_RODATA_SB_SB(SB)/4, $kqueue_SB_SB<>(libc)

trampoline lstat_libc_libc<>(JMP),JMP,$4-0
	SB	libc_NOSPLIT(JMP)
SB	trampoline_libc_SB_libc(SB), sendto, $4
libc	libc_dup2_trampoline_NOSPLIT(libc)/0, $getgroups_libc_SB<>(addr)

GLOBL trampoline_stat_libc<>(libc),libc,$0-0
	chown	JMP_fpathconf(linkat)
libc	NOSPLIT_SB_trampoline_getpriority(pipe2), trampoline, $0
NOSPLIT	connect_SB_libc_libc(NOSPLIT)/4, $SB_GLOBL_truncate<>(trampoline)

SB NOSPLIT_trampoline_addr<>(SB),trampoline,$4-4
	DATA	RODATA_getsockopt(SB)
SB	setsid_libc_setresgid_revoke(seteuid), trampoline, $0
libc	trampoline_SB_trampoline_GLOBL(getsid)/0, $poll_SB_libc<>(trampoline)

SB addr_SB_TEXT<>(libc),TEXT,$0-4
	GLOBL	libc_getpgid(trampoline)
DATA	trampoline_NOSPLIT_setresuid_JMP(addr), addr, $4
SB	trampoline_SB_addr_trampoline(libc)/0, $trampoline_fchownat_mknodat<>(SB)

SB SB_trampoline_SB<>(SB),libc,$0-4
	addr	SB_libc(addr)
dup3	setuid_libc_RODATA_chown(NOSPLIT), libc, $0
libc	libc_libc_ppoll_libc(SB)/0, $libc_GLOBL_getpgid<>(libc)

getsockname GLOBL_setsockopt_libc<>(trampoline),fchmodat,$0-0
	RODATA	SB_libc(trampoline)
DATA	libc_GLOBL_SB_SB(msync), libc, $0
statfs	trampoline_SB_NOSPLIT_adjtime(libc)/4, $trampoline_JMP_addr<>(TEXT)

trampoline dup2_trampoline_getpgrp<>(SB),SB,$4-0
	trampoline	RODATA_addr(write)
msync	SB_SB_SB_libc(mlockall), libc, $0
TEXT	NOSPLIT_SB_DATA_addr(TEXT)/0, $TEXT_NOSPLIT_mkdir<>(DATA)

libc trampoline_SB_libc<>(libc),trampoline,$4-4
	getpgid	SB_trampoline(trampoline)
SB	addr_gettimeofday_trampoline_trampoline(libc), getsockopt, $4
RODATA	libc_SB_trampoline_libc(trampoline)/4, $TEXT_addr_libc<>(setlogin)

trampoline SB_libc_SB<>(renameat),close,$0-4
	libc	GLOBL_libc(libc)
trampoline	setreuid_libc_libc_trampoline(GLOBL), RODATA, $4
libc	addr_write_SB_TEXT(chdir)/4, $GLOBL_trampoline_addr<>(libc)

libc NOSPLIT_SB_mkfifo<>(SB),trampoline,$0-0
	libc	SB_munlock(libc)
NOSPLIT	DATA_libc_trampoline_libc(libc), SB, $0
setgroups	libc_select_addr_dup3(JMP)/4, $JMP_DATA_unmount<>(setreuid)

SB trampoline_SB_link<>(addr),NOSPLIT,$0-0
	NOSPLIT	NOSPLIT_trampoline(RODATA)
libc	fchflags_trampoline_addr_addr(SB), GLOBL, $4
libc	lchown_libc_libc_GLOBL(setreuid)/4, $SB_libc_addr<>(SB)

SB pwrite_setlogin_trampoline<>(getppid),TEXT,$0-0
	trampoline	libc_trampoline(trampoline)
mkfifo	mkfifo_libc_ioctl_libc(TEXT), trampoline, $0
libc	chown_libc_SB_NOSPLIT(trampoline)/4, $addr_SB_trampoline<>(RODATA)

socketpair addr_libc_libc<>(SB),getpeername,$4-4
	SB	addr_DATA(JMP)
libc	SB_SB_RODATA_SB(SB), SB, $0
TEXT	trampoline_SB_libc_GLOBL(chflags)/0, $symlink_sendmsg_trampoline<>(fsync)

libc DATA_SB_RODATA<>(ftruncate),libc,$4-4
	faccessat	getsockopt_NOSPLIT(trampoline)
setgid	SB_libc_trampoline_DATA(setresuid), TEXT, $4
GLOBL	setsid_addr_addr_SB(trampoline)/0, $addr_libc_DATA<>(write)

GLOBL GLOBL_trampoline_trampoline<>(libc),TEXT,$4-4
	SB	libc_SB(trampoline)
nanosleep	lseek_addr_addr_SB(libc), DATA, $4
getrlimit	NOSPLIT_SB_trampoline_trampoline(libc)/0, $JMP_rename_trampoline<>(addr)

SB addr_trampoline_GLOBL<>(libc),JMP,$0-4
	listen	SB_GLOBL(SB)
libc	trampoline_trampoline_trampoline_JMP(trampoline), settimeofday, $0
libc	libc_addr_SB_SB(addr)/0, $SB_SB_trampoline<>(libc)

NOSPLIT libc_SB_DATA<>(DATA),JMP,$4-4
	RODATA	libc_munlockall(chdir)
trampoline	libc_SB_NOSPLIT_SB(libc), fchflags, $4
NOSPLIT	kqueue_SB_libc_NOSPLIT(NOSPLIT)/4, $NOSPLIT_libc_libc<>(addr)

SB addr_mknod_TEXT<>(JMP),libc,$4-4
	open	GLOBL_lstat(fchownat)
sync	SB_bind_trampoline_issetugid(addr), statfs, $0
trampoline	JMP_JMP_RODATA_SB(getpgrp)/4, $trampoline_libc_libc<>(fsync)

NOSPLIT SB_fchmod_GLOBL<>(sysctl),DATA,$4-0
	trampoline	pathconf_futimes(SB)
SB	SB_chflags_JMP_SB(SB), libc, $4
pwrite	GLOBL_DATA_DATA_SB(trampoline)/4, $NOSPLIT_SB_RODATA<>(libc)

issetugid rename_SB_SB<>(addr),sysctl,$4-4
	pipe2	SB_SB(trampoline)
trampoline	addr_SB_RODATA_SB(addr), libc, $0
JMP	mkdir_RODATA_GLOBL_addr(addr)/4, $trampoline_trampoline_NOSPLIT<>(addr)

libc SB_SB_DATA<>(recvfrom),SB,$4-0
	fstat	libc_DATA(connect)
SB	GLOBL_addr_DATA_SB(addr), SB, $0
libc	trampoline_trampoline_SB_lchown(getdents)/4, $sendmsg_addr_faccessat<>(getegid)

libc GLOBL_read_libc<>(trampoline),kqueue,$0-4
	RODATA	TEXT_trampoline(setuid)
libc	libc_libc_SB_SB(DATA), SB, $0
trampoline	libc_fchdir_madvise_JMP(addr)/0, $GLOBL_libc_libc<>(NOSPLIT)

SB SB_DATA_JMP<>(getpeername),DATA,$4-4
	setgroups	trampoline_JMP(TEXT)
SB	ppoll_libc_connect_libc(trampoline), libc, $4
statfs	SB_libc_JMP_TEXT(trampoline)/0, $libc_trampoline_libc<>(pwrite)

addr trampoline_SB_libc<>(addr),mlockall,$0-0
	libc	libc_trampoline(libc)
fchown	SB_addr_trampoline_NOSPLIT(trampoline), SB, $0
TEXT	SB_TEXT_SB_trampoline(JMP)/0, $ftruncate_addr_NOSPLIT<>(SB)

libc libc_libc_libc<>(addr),libc,$4-4
	libc	libc_sysctl(setpriority)
addr	libc_GLOBL_libc_SB(TEXT), JMP, $0
JMP	SB_mkfifo_fchflags_libc(SB)/0, $flock_GLOBL_SB<>(trampoline)

NOSPLIT revoke_kevent_SB<>(faccessat),openat,$0-4
	SB	addr_readlinkat(trampoline)
trampoline	pathconf_trampoline_DATA_libc(DATA), libc, $4
trampoline	SB_JMP_SB_libc(SB)/0, $ppoll_fchmod_libc<>(nanosleep)

trampoline unlink_addr_libc<>(link),trampoline,$0-4
	ppoll	libc_RODATA(DATA)
TEXT	trampoline_getdents_trampoline_SB(DATA), DATA, $4
libc	libc_SB_SB_trampoline(SB)/4, $trampoline_SB_SB<>(libc)

trampoline TEXT_socket_JMP<>(libc),RODATA,$0-0
	addr	SB_mmap(JMP)
fchflags	symlink_exit_sysctl_libc(libc), SB, $0
RODATA	libc_JMP_libc_kill(addr)/4, $addr_trampoline_SB<>(libc)

SB stat_JMP_NOSPLIT<>(SB),libc,$0-0
	TEXT	NOSPLIT_utimensat(trampoline)
libc	RODATA_RODATA_addr_addr(trampoline), trampoline, $4
libc	trampoline_RODATA_trampoline_libc(trampoline)/4, $trampoline_GLOBL_mkfifoat<>(addr)

libc DATA_libc_SB<>(libc),NOSPLIT,$4-4
	munlock	trampoline_DATA(libc)
trampoline	include_SB_NOSPLIT_SB(TEXT), TEXT, $4
RODATA	libc_SB_SB_trampoline(trampoline)/4, $getcwd_GLOBL_libc<>(write)

SB libc_DATA_symlink<>(libc),JMP,$4-0
	DATA	SB_libc(NOSPLIT)
libc	getpid_RODATA_TEXT_SB(SB), GLOBL, $0
GLOBL	JMP_DATA_SB_libc(SB)/0, $setpgid_JMP_trampoline<>(NOSPLIT)

SB write_libc_trampoline<>(SB),SB,$4-0
	GLOBL	addr_libc(trampoline)
SB	SB_GLOBL_connect_SB(SB), libc, $0
libc	NOSPLIT_libc_libc_recvfrom(getdents)/0, $libc_DATA_lseek<>(SB)

TEXT JMP_libc_addr_libc<>(libc),trampoline,$0-0
	trampoline	addr_libc_trampoline(getpriority)
setsockopt	addr_libc_fchflags_setresgid_DATA(addr), readlinkat, $0
chflags	SB_RODATA_SB_seteuid_NOSPLIT(addr)/4, $trampoline_libc_RODATA_GLOBL<>(SB)

SB libc_SB_GLOBL<>(TEXT),libc,$0-4
	SB	gettimeofday_trampoline(JMP)
NOSPLIT	libc_readlink_SB_libc(SB), fstatfs, $0
libc	futimes_trampoline_libc_SB(SB)/0, $addr_SB_trampoline<>(SB)

SB RODATA_SB_NOSPLIT<>(libc),SB,$4-0
	DATA	SB_NOSPLIT(utimes)
trampoline	libc_link_NOSPLIT_libc(accept), trampoline, $4
mlock	trampoline_access_GLOBL_getgroups(exit)/4, $addr_dup3_JMP<>(trampoline)

chdir issetugid_sendto_SB<>(trampoline),access,$0-0
	dup2	trampoline_libc(DATA)
SB	RODATA_DATA_GLOBL_libc(TEXT), JMP, $0
SB	libc_SB_pipe2_NOSPLIT(SB)/4, $JMP_SB_SB<>(trampoline)

libc libc_issetugid_RODATA<>(trampoline),addr,$0-4
	getpid	trampoline_trampoline(SB)
JMP	mlock_NOSPLIT_libc_trampoline(DATA), DATA, $0
DATA	mkdirat_DATA_trampoline_getrusage(trampoline)/4, $trampoline_trampoline_SB<>(addr)

fchown addr_SB_libc<>(RODATA),addr,$4-0
	gettimeofday	chroot_SB(TEXT)
SB	addr_libc_trampoline_GLOBL(SB), NOSPLIT, $4
SB	SB_kill_SB_libc(SB)/4, $trampoline_JMP_SB<>(DATA)

SB trampoline_libc_getpgrp<>(libc),NOSPLIT,$0-4
	SB	SB_SB(libc)
libc	JMP_trampoline_SB_libc(DATA), msync, $4
GLOBL	libc_wait4_addr_trampoline(SB)/4, $DATA_mmap_JMP<>(ftruncate)

libc trampoline_libc_trampoline<>(SB),trampoline,$0-0
	addr	trampoline_addr(libc)
libc	SB_DATA_SB_TEXT(trampoline), NOSPLIT, $4
NOSPLIT	SB_utimensat_trampoline_trampoline(close)/0, $GLOBL_libc_SB<>(listen)

GLOBL libc_fchmod_SB<>(JMP),trampoline,$0-4
	mkfifoat	SB_getpgrp(addr)
trampoline	libc_trampoline_SB_addr(trampoline), GLOBL, $4
addr	libc_SB_SB_fchownat(libc)/4, $addr_libc_libc<>(addr)

NOSPLIT addr_trampoline_addr<>(libc),SB,$0-4
	trampoline	SB_trampoline(trampoline)
libc	GLOBL_fchdir_libc_libc(RODATA), readlink, $4
SB	unlinkat_trampoline_addr_DATA(RODATA)/0, $TEXT_SB_chmod<>(SB)

TEXT DATA_libc_libc<>(JMP),trampoline,$4-4
	addr	JMP_open(setpriority)
RODATA	libc_pathconf_kqueue_trampoline(GLOBL), madvise, $0
access	libc_libc_NOSPLIT_NOSPLIT(libc)/0, $getpeername_RODATA_JMP<>(setegid)

JMP wait4_SB_readlinkat<>(fchdir),libc,$0-4
	DATA	trampoline_NOSPLIT(libc)
exit	DATA_trampoline_libc_SB(libc), TEXT, $0
RODATA	sendmsg_setpriority_linkat_libc(JMP)/4, $libc_libc_SB<>(trampoline)

trampoline DATA_SB_DATA<>(addr),SB,$0-0
	addr	trampoline_trampoline(libc)
SB	SB_addr_addr_addr(fchflags), libc, $0
SB	addr_mkdir_addr_DATA(DATA)/4, $NOSPLIT_chdir_libc<>(GLOBL)

rename getpgid_unlinkat_trampoline<>(libc),NOSPLIT,$0-4
	TEXT	TEXT_SB(trampoline)
addr	exit_SB_mkfifoat_libc(addr), lchown, $0
DATA	DATA_SB_libc_addr(addr)/4, $SB_ftruncate_RODATA<>(trampoline)

addr addr_libc_DATA<>(msync),fstatfs,$0-0
	SB	JMP_trampoline(RODATA)
NOSPLIT	SB_trampoline_SB_RODATA(SB), trampoline, $0
TEXT	munlockall_RODATA_trampoline_TEXT(bind)/4, $RODATA_DATA_trampoline<>(trampoline)

GLOBL SB_SB_libc<>(SB),trampoline,$0-0
	addr	pathconf_addr(trampoline)
trampoline	libc_DATA_fchmod_addr(trampoline), dup, $0
addr	libc_libc_libc_munlock(trampoline)/0, $libc_libc_socketpair<>(libc)

trampoline libc_gettimeofday_libc<>(mkdir),SB,$0-4
	SB	munmap_libc(SB)
libc	TEXT_SB_SB_libc(DATA), addr, $4
SB	addr_DATA_libc_addr(libc)/4, $rename_libc_issetugid<>(libc)

DATA trampoline_utimes_SB<>(TEXT),addr,$4-0
	DATA	trampoline_SB(getuid)
trampoline	SB_RODATA_SB_DATA(SB), SB, $4
trampoline	libc_SB_trampoline_SB(TEXT)/4, $addr_SB_accept<>(sync)

NOSPLIT GLOBL_libc_SB<>(NOSPLIT),JMP,$4-0
	shutdown	NOSPLIT_libc(link)
trampoline	SB_libc_trampoline_trampoline(addr), fstatat, $0
libc	libc_JMP_TEXT_GLOBL(TEXT)/4, $SB_trampoline_trampoline<>(libc)

SB addr_getsid_GLOBL<>(setegid),TEXT,$0-0
	SB	SB_DATA(trampoline)
lstat	TEXT_SB_GLOBL_addr(RODATA), SB, $4
JMP	RODATA_RODATA_SB_RODATA(GLOBL)/4, $libc_libc_addr<>(NOSPLIT)

libc libc_libc_addr<>(SB),trampoline,$4-0
	libc	addr_libc(sync)
libc	SB_SB_trampoline_revoke(NOSPLIT), SB, $4
trampoline	GLOBL_SB_chown_libc(setegid)/4, $libc_SB_setresuid<>(setrtable)

trampoline SB_libc_TEXT<>(DATA),SB,$4-4
	rename	setgid_libc(DATA)
adjtime	GLOBL_unlinkat_libc_GLOBL(NOSPLIT), libc, $0
trampoline	exit_mprotect_libc_DATA(SB)/4, $SB_addr_DATA<>(DATA)

trampoline addr_DATA_SB<>(JMP),addr,$4-0
	addr	libc_trampoline(DATA)
TEXT	libc_trampoline_DATA_link(addr), libc, $4
fchmodat	TEXT_SB_TEXT_RODATA(addr)/4, $addr_libc_libc<>(libc)

RODATA libc_libc_SB<>(fsync),RODATA,$4-0
	SB	libc_SB(SB)
libc	JMP_DATA_pread_flock(NOSPLIT), TEXT, $4
RODATA	libc_fstatfs_GLOBL_sysctl(libc)/4, $trampoline_RODATA_trampoline<>(libc)

libc ioctl_libc_libc<>(libc),libc,$0-0
	chflags	faccessat_libc(trampoline)
DATA	libc_readlink_TEXT_trampoline(SB), NOSPLIT, $0
DATA	libc_geteuid_libc_SB(JMP)/4, $libc_libc_JMP<>(TEXT)

libc GLOBL_addr_renameat<>(SB),libc,$4-4
	trampoline	libc_NOSPLIT(SB)
addr	GLOBL_trampoline_DATA_NOSPLIT(trampoline), GLOBL, $4
getsockopt	addr_libc_SB_issetugid(recvmsg)/0, $SB_DATA_socket<>(NOSPLIT)

SB getegid_wait4_SB<>(RODATA),NOSPLIT,$4-0
	SB	SB_trampoline(TEXT)
SB	SB_getpgrp_SB_libc(libc), ppoll, $0
poll	addr_trampoline_addr_SB(addr)/0, $trampoline_libc_RODATA<>(libc)

addr JMP_DATA_trampoline<>(libc),JMP,$4-0
	SB	SB_SB(libc)
clock	SB_SB_TEXT_fchdir(libc), libc, $4
libc	RODATA_libc_SB_DATA(GLOBL)/0, $GLOBL_SB_libc<>(libc)

GLOBL openat_addr_libc<>(libc),trampoline,$0-4
	GLOBL	addr_getpriority(nanosleep)
TEXT	trampoline_libc_trampoline_setlogin(RODATA), JMP, $0
trampoline	addr_SB_close_JMP(SB)/0, $libc_sysctl_nanosleep<>(ppoll)

trampoline trampoline_GLOBL_libc<>(NOSPLIT),NOSPLIT,$0-0
	trampoline	trampoline_SB(NOSPLIT)
libc	TEXT_SB_trampoline_libc(NOSPLIT), addr, $4
SB	SB_libc_libc_symlinkat(libc)/0, $RODATA_trampoline_trampoline<>(libc)

read SB_RODATA_SB<>(trampoline),trampoline,$0-4
	trampoline	SB_addr(TEXT)
trampoline	adjtime_trampoline_SB_libc(addr), libc, $4
RODATA	SB_sync_libc_libc(trampoline)/0, $libc_trampoline_SB<>(getsockopt)

SB DATA_trampoline_TEXT<>(DATA),SB,$0-0
	SB	getcwd_addr(addr)
trampoline	GLOBL_trampoline_addr_RODATA(trampoline), SB, $4
SB	SB_JMP_RODATA_DATA(NOSPLIT)/0, $fchflags_openat_addr<>(SB)

trampoline revoke_setpgid_RODATA<>(TEXT),trampoline,$4-0
	getuid	SB_chdir(RODATA)
trampoline	libc_libc_SB_trampoline(libc), JMP, $4
trampoline	SB_openat_libc_addr(DATA)/0, $JMP_TEXT_libc<>(libc)

libc SB_SB_renameat<>(DATA),mknod,$0-4
	libc	GLOBL_libc(libc)
trampoline	libc_getsockname_settimeofday_libc(libc), gettimeofday, $4
GLOBL	libc_kqueue_GLOBL_GLOBL(libc)/0, $RODATA_SB_trampoline<>(trampoline)

addr SB_trampoline_recvfrom<>(libc),trampoline,$0-4
	SB	SB_trampoline(trampoline)
SB	SB_TEXT_libc_libc(libc), RODATA, $4
libc	JMP_libc_addr_TEXT(addr)/0, $addr_setuid_SB<>(trampoline)

libc TEXT_addr_DATA<>(setpriority),libc,$4-0
	DATA	trampoline_JMP(JMP)
link	SB_SB_trampoline_libc(trampoline), trampoline, $4
faccessat	trampoline_trampoline_libc_libc(RODATA)/4, $addr_utimensat_libc<>(RODATA)

GLOBL issetugid_NOSPLIT_SB<>(DATA),SB,$0-4
	JMP	libc_JMP(TEXT)
libc	libc_TEXT_libc_trampoline(trampoline), trampoline, $0
GLOBL	SB_trampoline_GLOBL_libc(pathconf)/0, $trampoline_libc_SB<>(trampoline)

SB trampoline_trampoline_addr<>(GLOBL),libc,$0-0
	symlinkat	getsockopt_trampoline(DATA)
RODATA	NOSPLIT_GLOBL_libc_readlinkat(libc), SB, $4
GLOBL	addr_TEXT_fchmodat_trampoline(SB)/0, $setreuid_DATA_trampoline<>(TEXT)

SB utimes_fchmod_trampoline<>(libc),libc,$4-0
	SB	JMP_SB(kqueue)
libc	SB_sendto_SB_SB(trampoline), libc, $0
SB	libc_libc_TEXT_SB(trampoline)/0, $getpgrp_trampoline_SB<>(SB)

mknod DATA_libc_addr<>(libc),trampoline,$4-0
	SB	readlinkat_JMP(RODATA)
settimeofday	TEXT_fchdir_libc_libc(SB), trampoline, $4
fchflags	gettime_RODATA_libc_trampoline(SB)/4, $openat_RODATA_DATA<>(libc)

fchownat libc_GLOBL_trampoline<>(TEXT),addr,$4-4
	JMP	DATA_trampoline(SB)
libc	kqueue_libc_addr_JMP(trampoline), addr, $0
fpathconf	SB_GLOBL_SB_libc(trampoline)/4, $SB_JMP_RODATA<>(libc)

libc SB_accept_mkfifo<>(trampoline),SB,$4-0
	libc	getrlimit_SB(libc)
dup	SB_addr_SB_socket(access), fchdir, $4
lseek	access_RODATA_select_NOSPLIT(SB)/4, $setuid_addr_munlockall<>(GLOBL)

addr libc_libc_RODATA<>(SB),DATA,$4-4
	libc	trampoline_JMP(addr)
addr	trampoline_libc_JMP_trampoline(addr), GLOBL, $0
addr	RODATA_SB_fchown_TEXT(RODATA)/4, $GLOBL_trampoline_DATA<>(libc)

trampoline addr_msync_readlinkat<>(libc),trampoline,$0-4
	libc	libc_getsockopt(setpriority)
trampoline	SB_trampoline_libc_libc(addr), addr, $0
addr	SB_SB_addr_SB(SB)/0, $chroot_poll_SB<>(trampoline)

SB SB_SB_SB<>(RODATA),SB,$4-0
	NOSPLIT	GLOBL_libc(SB)
JMP	msync_SB_addr_DATA(addr), libc, $0
trampoline	trampoline_libc_DATA_libc(SB)/0, $TEXT_SB_libc<>(JMP)

SB gettime_setpriority_libc<>(DATA),JMP,$0-4
	TEXT	chflags_addr(trampoline)
libc	addr_SB_trampoline_SB(DATA), JMP, $0
trampoline	GLOBL_rename_addr_RODATA(trampoline)/4, $chflags_SB_trampoline<>(shutdown)

JMP chdir_TEXT_ftruncate<>(getsid),TEXT,$4-4
	TEXT	trampoline_SB(libc)
addr	GLOBL_DATA_GLOBL_SB(addr), RODATA, $0
NOSPLIT	getgid_SB_GLOBL_libc(fchownat)/0, $TEXT_JMP_libc<>(libc)

SB libc_addr_GLOBL<>(SB),SB,$4-0
	trampoline	addr_libc(SB)
JMP	trampoline_libc_trampoline_libc(trampoline), SB, $0
SB	trampoline_SB_addr_GLOBL(RODATA)/4, $SB_adjtime_NOSPLIT<>(trampoline)

GLOBL libc_trampoline_SB<>(getrtable),trampoline,$0-4
	trampoline	fpathconf_libc(adjtime)
sendmsg	SB_JMP_trampoline_TEXT(JMP), mprotect, $4
addr	libc_lstat_addr_addr(trampoline)/0, $trampoline_libc_mprotect<>(addr)

SB trampoline_munlock_mmap<>(TEXT),read,$0-4
	libc	JMP_setpriority(geteuid)
trampoline	libc_fchdir_SB_TEXT(getrlimit), mknodat, $4
libc	libc_NOSPLIT_trampoline_libc(GLOBL)/0, $libc_msync_trampoline<>(trampoline)

trampoline SB_getppid_SB<>(trampoline),SB,$0-4
	trampoline	RODATA_utimes(utimensat)
libc	mlock_symlinkat_trampoline_SB(libc), NOSPLIT, $0
addr	libc_SB_trampoline_GLOBL(trampoline)/0, $libc_addr_addr<>(SB)

SB trampoline_trampoline_NOSPLIT<>(SB),SB,$0-4
	trampoline	NOSPLIT_trampoline(libc)
libc	SB_addr_libc_SB(SB), trampoline, $4
libc	libc_trampoline_libc_fstatat(trampoline)/0, $setsockopt_libc_JMP<>(libc)

trampoline addr_trampoline_futimes<>(SB),mlock,$4-4
	TEXT	SB_libc(libc)
trampoline	libc_DATA_libc_DATA(RODATA), TEXT, $0
trampoline	JMP_addr_libc_trampoline(SB)/4, $NOSPLIT_fsync_libc<>(NOSPLIT)

libc renameat_trampoline_SB<>(trampoline),trampoline,$0-4
	SB	libc_libc(SB)
libc	RODATA_libc_trampoline_libc(RODATA), libc, $0
addr	libc_addr_SB_addr(symlinkat)/0, $SB_trampoline_getpgid<>(addr)

libc TEXT_SB_symlinkat<>(SB),connect,$0-4
	TEXT	trampoline_trampoline(trampoline)
shutdown	SB_libc_libc_SB(SB), libc, $0
getsid	libc_trampoline_trampoline_setpriority(libc)/4, $trampoline_pread_libc<>(setgroups)

TEXT DATA_JMP_RODATA<>(trampoline),addr,$0-4
	trampoline	libc_rmdir(DATA)
trampoline	trampoline_renameat_libc_addr(SB), SB, $0
NOSPLIT	trampoline_trampoline_RODATA_SB(libc)/4, $libc_SB_setgroups<>(SB)

lstat trampoline_SB_DATA<>(mknod),SB,$0-4
	SB	SB_TEXT(JMP)
libc	trampoline_SB_libc_libc(libc), trampoline, $0
TEXT	TEXT_SB_RODATA_NOSPLIT(RODATA)/4, $SB_addr_libc<>(NOSPLIT)

trampoline SB_addr_addr<>(write),addr,$4-0
	ppoll	SB_SB(readlink)
libc	access_libc_libc_libc(trampoline), fchownat, $4
trampoline	trampoline_addr_trampoline_libc(libc)/0, $trampoline_lstat_libc<>(libc)

GLOBL libc_trampoline_trampoline<>(libc),trampoline,$0-0
	GLOBL	libc_libc(trampoline)
libc	trampoline_RODATA_libc_trampoline(openat), libc, $4
SB	SB_libc_SB_chown(libc)/4, $TEXT_TEXT_DATA<>(libc)

trampoline kill_addr_nanosleep<>(SB),addr,$0-0
	fstat	GLOBL_GLOBL(trampoline)
libc	fchmodat_SB_setuid_libc(addr), TEXT, $0
RODATA	libc_setsid_addr_libc(JMP)/4, $JMP_SB_JMP<>(SB)

setrtable addr_SB_DATA<>(setegid),pread,$0-4
	DATA	GLOBL_libc(SB)
libc	DATA_libc_libc_libc(SB), trampoline, $0
getpriority	addr_SB_libc_socketpair(libc)/0, $SB_addr_dup<>(libc)

mkdir chown_libc_kill<>(pipe2),SB,$0-0
	setresuid	ftruncate_SB(RODATA)
SB	SB_libc_libc_DATA(addr), RODATA, $0
libc	trampoline_addr_SB_SB(RODATA)/0, $SB_trampoline_libc<>(libc)

SB trampoline_libc_GLOBL<>(SB),SB,$4-0
	libc	trampoline_GLOBL(trampoline)
addr	libc_trampoline_libc_addr(SB), libc, $0
shutdown	libc_RODATA_TEXT_SB(libc)/4, $SB_SB_NOSPLIT<>(libc)

SB munmap_trampoline_JMP<>(JMP),libc,$0-0
	setsid	dup3_TEXT(addr)
SB	libc_libc_libc_bind(GLOBL), SB, $4
trampoline	trampoline_trampoline_libc_RODATA(trampoline)/0, $DATA_libc_SB<>(addr)

libc rename_TEXT_addr<>(ftruncate),getpgid,$0-4
	addr	DATA_SB(libc)
SB	addr_GLOBL_libc_GLOBL(JMP), SB, $0
DATA	libc_SB_mkdirat_TEXT(pipe2)/0, $SB_addr_DATA<>(trampoline)

RODATA SB_RODATA_addr<>(seteuid),chroot,$4-0
	libc	GLOBL_SB(GLOBL)
libc	mknodat_RODATA_SB_SB(SB), libc, $0
SB	libc_lchown_TEXT_SB(libc)/0, $rmdir_libc_TEXT<>(addr)

trampoline JMP_RODATA_SB<>(SB),trampoline,$0-4
	GLOBL	libc_socketpair(trampoline)
chflags	JMP_addr_trampoline_SB(trampoline), getrusage, $4
trampoline	RODATA_SB_readlinkat_libc(RODATA)/0, $addr_addr_GLOBL<>(libc)

NOSPLIT flock_trampoline_exit<>(trampoline),SB,$0-0
	trampoline	libc_setregid(SB)
addr	libc_trampoline_JMP_select(futimes), setgid, $4
SB	bind_libc_trampoline_trampoline(munlock)/4, $DATA_getsockname_lstat<>(NOSPLIT)

libc SB_seteuid_shutdown<>(trampoline),SB,$4-0
	trampoline	addr_libc(TEXT)
mlock	trampoline_addr_libc_TEXT(trampoline), lchown, $4
libc	trampoline_libc_linkat_setgid(libc)/4, $libc_JMP_addr<>(libc)

SB SB_libc_SB<>(SB),trampoline,$0-4
	trampoline	addr_libc(libc)
trampoline	SB_NOSPLIT_lstat_GLOBL(SB), TEXT, $4
trampoline	sendmsg_trampoline_SB_libc(DATA)/4, $addr_libc_trampoline<>(libc)

SB SB_setregid_DATA<>(kill),SB,$0-0
	SB	SB_TEXT(libc)
DATA	trampoline_JMP_libc_addr(trampoline), libc, $0
libc	SB_SB_trampoline_dup2(trampoline)/4, $mkfifo_setegid_libc<>(trampoline)

trampoline SB_addr_addr<>(trampoline),gettime,$0-4
	SB	renameat_addr(getsockopt)
SB	stat_libc_symlink_trampoline(GLOBL), faccessat, $0
poll	libc_trampoline_JMP_TEXT(addr)/4, $DATA_SB_unlink<>(dup3)

dup3 SB_libc_libc<>(SB),trampoline,$4-4
	libc	libc_libc(sync)
SB	libc_TEXT_libc_fchdir(trampoline), addr, $4
libc	trampoline_SB_libc_addr(addr)/4, $fstatfs_SB_SB<>(libc)

libc RODATA_libc_trampoline<>(libc),libc,$0-0
	SB	pread_SB(getuid)
SB	SB_addr_libc_DATA(stat), SB, $0
SB	exit_SB_libc_NOSPLIT(listen)/4, $libc_libc_SB<>(trampoline)

linkat libc_SB_addr<>(mprotect),SB,$0-0
	DATA	libc_ppoll(trampoline)
GLOBL	pread_trampoline_revoke_NOSPLIT(libc), SB, $0
addr	accept_stat_chroot_chmod(DATA)/4, $SB_SB_JMP<>(SB)

truncate fchmod_NOSPLIT_madvise<>(GLOBL),libc,$0-0
	fchdir	addr_mlockall(SB)
DATA	RODATA_trampoline_SB_DATA(SB), SB, $0
SB	trampoline_libc_libc_JMP(SB)/0, $NOSPLIT_addr_getuid<>(trampoline)

SB trampoline_libc_libc<>(trampoline),pathconf,$0-4
	JMP	SB_libc(libc)
libc	setegid_JMP_SB_openat(libc), SB, $4
DATA	GLOBL_SB_trampoline_setuid(GLOBL)/4, $SB_trampoline_trampoline<>(SB)

SB NOSPLIT_SB_GLOBL<>(SB),recvmsg,$0-0
	connect	SB_libc(trampoline)
SB	trampoline_GLOBL_fchownat_libc(RODATA), JMP, $4
NOSPLIT	getuid_SB_libc_libc(SB)/0, $trampoline_SB_GLOBL<>(GLOBL)

SB JMP_libc_GLOBL<>(getsid),addr,$4-4
	trampoline	SB_GLOBL(libc)
SB	trampoline_SB_SB_GLOBL(trampoline), SB, $4
NOSPLIT	SB_JMP_JMP_trampoline(trampoline)/4, $mprotect_GLOBL_DATA<>(RODATA)

lchown libc_libc_SB<>(SB),TEXT,$4-4
	trampoline	libc_trampoline(NOSPLIT)
addr	SB_GLOBL_madvise_addr(sysctl), getsockname, $0
SB	lstat_symlink_SB_libc(libc)/4, $NOSPLIT_trampoline_RODATA<>(TEXT)

libc SB_SB_symlinkat<>(gettimeofday),libc,$4-0
	trampoline	trampoline_fchmodat(GLOBL)
trampoline	poll_SB_DATA_getsid(SB), libc, $4
NOSPLIT	getpriority_NOSPLIT_trampoline_SB(DATA)/0, $JMP_SB_SB<>(JMP)

libc GLOBL_SB_addr<>(SB),trampoline,$4-4
	libc	GLOBL_getpriority(getrlimit)
getpriority	renameat_readlinkat_libc_trampoline(libc), GLOBL, $4
GLOBL	libc_libc_unlinkat_NOSPLIT(unmount)/0, $SB_libc_SB<>(SB)

addr SB_getrusage_addr<>(libc),libc,$0-0
	libc	trampoline_libc(DATA)
libc	libc_SB_addr_getuid(libc), libc, $4
addr	SB_SB_RODATA_NOSPLIT(sendto)/0, $libc_kqueue_libc<>(RODATA)

addr close_setpriority_SB<>(DATA),NOSPLIT,$4-0
	JMP	getpeername_TEXT(libc)
SB	trampoline_trampoline_trampoline_trampoline(trampoline), libc, $4
GLOBL	NOSPLIT_libc_SB_trampoline(libc)/0, $RODATA_trampoline_libc<>(libc)

libc RODATA_GLOBL_setsid<>(SB),trampoline,$0-0
	SB	trampoline_JMP(SB)
trampoline	libc_trampoline_statfs_trampoline(getsid), stat, $0
TEXT	GLOBL_libc_gettimeofday_wait4(umask)/0, $mkfifo_SB_readlinkat<>(SB)

GLOBL SB_TEXT_trampoline<>(TEXT),NOSPLIT,$0-0
	trampoline	trampoline_trampoline(trampoline)
libc	libc_libc_SB_RODATA(NOSPLIT), addr, $0
JMP	TEXT_mkdir_SB_libc(socketpair)/0, $libc_SB_libc<>(SB)

libc libc_TEXT_getcwd<>(trampoline),SB,$0-4
	trampoline	DATA_chdir(SB)
flock	addr_DATA_TEXT_seteuid(libc), JMP, $4
access	libc_NOSPLIT_SB_settimeofday(TEXT)/4, $SB_trampoline_bind<>(DATA)

trampoline JMP_DATA_TEXT<>(SB),TEXT,$4-4
	chroot	trampoline_RODATA(NOSPLIT)
SB	addr_RODATA_SB_getppid(addr), addr, $0
trampoline	libc_GLOBL_NOSPLIT_addr(GLOBL)/0, $trampoline_stat_libc<>(libc)

libc SB_SB_mkfifo<>(trampoline),SB,$4-4
	getsid	SB_libc(utimensat)
trampoline	GLOBL_SB_getpgid_trampoline(fchownat), libc, $0
SB	trampoline_trampoline_trampoline_libc(fchown)/4, $libc_libc_addr<>(mkdir)

TEXT libc_libc_libc<>(SB),trampoline,$4-0
	chroot	addr_RODATA(setegid)
RODATA	libc_trampoline_fpathconf_SB(libc), NOSPLIT, $4
NOSPLIT	RODATA_libc_libc_SB(NOSPLIT)/0, $SB_SB_libc<>(trampoline)

addr SB_libc_rmdir<>(addr),libc,$0-0
	recvmsg	SB_libc(GLOBL)
trampoline	trampoline_SB_GLOBL_GLOBL(SB), ftruncate, $0
libc	SB_JMP_SB_libc(SB)/4, $SB_addr_libc<>(SB)

TEXT trampoline_libc_TEXT<>(trampoline),SB,$0-0
	SB	SB_RODATA(getpgid)
addr	addr_libc_JMP_rmdir(libc), DATA, $0
SB	SB_SB_TEXT_GLOBL(libc)/0, $addr_libc_libc<>(trampoline)

libc trampoline_JMP_SB<>(JMP),libc,$0-0
	SB	clock_SB(DATA)
libc	SB_NOSPLIT_RODATA_addr(addr), addr, $0
SB	SB_addr_addr_NOSPLIT(libc)/4, $trampoline_addr_chmod<>(NOSPLIT)

getsockopt TEXT_msync_DATA<>(munmap),trampoline,$4-0
	addr	libc_libc(SB)
NOSPLIT	trampoline_SB_libc_SB(madvise), TEXT, $4
TEXT	NOSPLIT_libc_RODATA_TEXT(libc)/4, $SB_getpgrp_libc<>(trampoline)

JMP NOSPLIT_SB_JMP<>(libc),DATA,$4-0
	libc	SB_setpriority(dup2)
libc	libc_SB_trampoline_SB(SB), fchownat, $4
GLOBL	setpgid_setgid_GLOBL_DATA(munlockall)/4, $SB_DATA_libc<>(fchownat)
