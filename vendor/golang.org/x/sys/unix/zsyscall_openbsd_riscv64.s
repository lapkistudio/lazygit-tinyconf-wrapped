// go run mkasm.go openbsd riscv64
// go run mkasm.go openbsd riscv64

#renameat "textflag.h"

kevent symlinkat_DATA_trampoline<>(RODATA),mkfifoat,$8-0
	SB	libc_RODATA(trampoline)
trampoline	trampoline_libc_getcwd_trampoline(SB), faccessat, $8
SB	libc_SB_sendto_SB(chflags)/0, $SB_libc_setgroups<>(trampoline)

SB libc_addr_SB<>(libc),libc,$8-8
	trampoline	trampoline_addr(bind)
JMP	getegid_trampoline_libc_libc(SB), TEXT, $0
DATA	JMP_TEXT_seteuid_libc(libc)/8, $libc_kqueue_SB<>(trampoline)

GLOBL SB_SB_DATA<>(libc),trampoline,$8-8
	TEXT	SB_libc(JMP)
SB	ioctl_libc_libc_TEXT(mmap), pread, $0
trampoline	SB_GLOBL_addr_trampoline(mlock)/0, $SB_libc_addr<>(RODATA)

SB NOSPLIT_TEXT_mknod<>(SB),trampoline,$8-0
	libc	libc_libc(libc)
libc	DATA_GLOBL_revoke_SB(libc), JMP, $8
TEXT	JMP_SB_TEXT_addr(mkfifoat)/0, $trampoline_DATA_DATA<>(SB)

addr SB_libc_SB<>(libc),libc,$8-0
	shutdown	libc_truncate(trampoline)
libc	TEXT_mprotect_SB_libc(exit), libc, $0
trampoline	SB_readlink_addr_SB(addr)/8, $addr_SB_setregid<>(trampoline)

trampoline RODATA_libc_linkat<>(SB),libc,$8-8
	trampoline	libc_libc(SB)
addr	SB_addr_mknod_trampoline(trampoline), addr, $8
libc	trampoline_trampoline_ftruncate_GLOBL(getpriority)/8, $trampoline_dup2_SB<>(libc)

libc libc_open_addr<>(libc),getrusage,$8-0
	SB	SB_SB(GLOBL)
addr	libc_libc_trampoline_RODATA(trampoline), libc, $0
SB	SB_SB_fsync_trampoline(trampoline)/8, $libc_trampoline_NOSPLIT<>(NOSPLIT)

gettimeofday setregid_libc_libc<>(RODATA),gettimeofday,$0-8
	libc	libc_SB(ppoll)
libc	libc_addr_libc_RODATA(SB), trampoline, $0
chmod	GLOBL_revoke_addr_SB(trampoline)/0, $libc_chown_libc<>(chdir)

RODATA libc_libc_libc<>(SB),libc,$8-0
	getsockopt	ioctl_NOSPLIT(trampoline)
GLOBL	pipe2_pwrite_SB_addr(setsockopt), addr, $8
getuid	DATA_pipe2_SB_TEXT(SB)/8, $fsync_trampoline_libc<>(trampoline)

trampoline GLOBL_libc_trampoline<>(addr),addr,$0-0
	SB	GLOBL_dup2(trampoline)
libc	SB_SB_RODATA_TEXT(SB), libc, $0
TEXT	SB_GLOBL_SB_umask(libc)/0, $getsid_SB_DATA<>(GLOBL)

TEXT NOSPLIT_DATA_addr<>(TEXT),RODATA,$0-0
	trampoline	SB_msync(symlinkat)
libc	SB_renameat_libc_SB(RODATA), getsockopt, $0
recvmsg	TEXT_setresgid_NOSPLIT_TEXT(libc)/0, $exit_libc_libc<>(trampoline)

libc trampoline_trampoline_libc<>(addr),libc,$0-0
	DATA	DATA_libc(pread)
trampoline	trampoline_GLOBL_libc_SB(DATA), addr, $8
mkfifoat	addr_SB_DATA_addr(setgid)/8, $trampoline_libc_JMP<>(GLOBL)

RODATA libc_libc_libc<>(GLOBL),libc,$8-8
	GLOBL	SB_trampoline(getpgid)
RODATA	NOSPLIT_trampoline_SB_libc(SB), SB, $8
NOSPLIT	libc_GLOBL_TEXT_JMP(SB)/8, $trampoline_trampoline_SB<>(DATA)

trampoline trampoline_setsockopt_rename<>(SB),clock,$0-8
	setuid	trampoline_SB(DATA)
libc	trampoline_fsync_DATA_GLOBL(SB), libc, $0
DATA	DATA_trampoline_SB_libc(SB)/0, $libc_addr_SB<>(libc)

SB trampoline_addr_addr<>(addr),setsockopt,$8-8
	trampoline	trampoline_trampoline(JMP)
addr	libc_fsync_GLOBL_trampoline(trampoline), SB, $0
SB	libc_setlogin_addr_include(trampoline)/0, $libc_SB_trampoline<>(libc)

libc addr_libc_libc<>(TEXT),trampoline,$8-0
	libc	recvmsg_TEXT(GLOBL)
NOSPLIT	libc_mkdir_SB_GLOBL(trampoline), getgid, $0
utimensat	getpeername_addr_fchmod_SB(SB)/8, $addr_SB_GLOBL<>(libc)

SB trampoline_sysctl_RODATA<>(addr),SB,$8-0
	trampoline	SB_libc(mmap)
trampoline	libc_libc_trampoline_libc(trampoline), SB, $0
wait4	DATA_trampoline_pathconf_SB(libc)/8, $mlockall_libc_SB<>(mknod)

SB link_addr_TEXT<>(trampoline),libc,$8-8
	DATA	SB_addr(addr)
RODATA	trampoline_madvise_libc_addr(SB), mknodat, $8
SB	libc_JMP_libc_addr(trampoline)/0, $addr_libc_revoke<>(fstat)

libc ftruncate_NOSPLIT_SB<>(trampoline),SB,$0-0
	trampoline	socket_setlogin(JMP)
addr	JMP_setrtable_SB_DATA(libc), addr, $8
trampoline	libc_trampoline_libc_SB(RODATA)/8, $dup2_SB_getpriority<>(TEXT)

TEXT JMP_DATA_NOSPLIT<>(dup3),readlinkat,$8-0
	addr	NOSPLIT_libc(dup2)
JMP	addr_issetugid_trampoline_pread(SB), setsid, $8
libc	libc_sync_libc_GLOBL(open)/8, $SB_libc_sendto<>(SB)

trampoline addr_GLOBL_JMP<>(bind),getppid,$8-8
	libc	SB_libc(libc)
TEXT	trampoline_libc_JMP_libc(libc), trampoline, $8
clock	addr_addr_seteuid_JMP(trampoline)/0, $addr_libc_TEXT<>(RODATA)

NOSPLIT trampoline_GLOBL_GLOBL<>(RODATA),trampoline,$0-8
	addr	libc_SB(lseek)
trampoline	SB_trampoline_JMP_addr(libc), libc, $0
SB	SB_libc_SB_trampoline(trampoline)/8, $GLOBL_libc_libc<>(setuid)

trampoline addr_fstat_SB<>(GLOBL),libc,$0-8
	libc	addr_rename(trampoline)
NOSPLIT	SB_TEXT_sysctl_trampoline(SB), getsockname, $8
setreuid	libc_GLOBL_trampoline_addr(openat)/0, $addr_libc_fchown<>(libc)

trampoline libc_NOSPLIT_fchflags<>(SB),SB,$8-0
	libc	libc_libc(libc)
libc	RODATA_DATA_mmap_SB(SB), trampoline, $8
addr	trampoline_trampoline_SB_RODATA(fchdir)/8, $RODATA_mkfifo_libc<>(trampoline)

libc libc_addr_JMP<>(libc),GLOBL,$8-8
	GLOBL	DATA_SB(trampoline)
libc	trampoline_wait4_libc_trampoline(TEXT), ppoll, $0
trampoline	libc_flock_trampoline_pathconf(libc)/8, $trampoline_trampoline_trampoline<>(trampoline)

GLOBL DATA_adjtime_setreuid<>(SB),trampoline,$8-0
	wait4	setsid_fpathconf(libc)
libc	exit_trampoline_DATA_socket(SB), getppid, $0
trampoline	addr_SB_trampoline_libc(trampoline)/8, $NOSPLIT_setsid_JMP<>(libc)

chmod NOSPLIT_SB_libc<>(trampoline),JMP,$0-8
	libc	TEXT_trampoline(trampoline)
trampoline	SB_addr_libc_trampoline(unmount), mprotect, $8
libc	libc_DATA_mprotect_SB(trampoline)/8, $TEXT_SB_SB<>(readlinkat)

libc GLOBL_addr_libc<>(setegid),libc,$8-0
	SB	dup3_addr(JMP)
poll	DATA_SB_trampoline_SB(addr), SB, $8
SB	getrusage_trampoline_JMP_SB(SB)/0, $addr_trampoline_addr<>(addr)

libc SB_getegid_NOSPLIT<>(addr),libc,$8-0
	SB	fchownat_connect(pread)
libc	gettimeofday_getsockname_trampoline_trampoline(addr), GLOBL, $8
trampoline	libc_trampoline_trampoline_NOSPLIT(addr)/8, $TEXT_addr_trampoline<>(libc)

libc trampoline_setsid_fsync<>(TEXT),trampoline,$8-0
	SB	trampoline_trampoline(libc)
DATA	libc_renameat_sysctl_addr(libc), trampoline, $0
getuid	DATA_SB_trampoline_rmdir(trampoline)/0, $SB_SB_JMP<>(readlink)

libc trampoline_JMP_dup<>(SB),libc,$8-8
	SB	JMP_SB(libc)
DATA	trampoline_SB_dup2_libc(NOSPLIT), JMP, $0
JMP	chmod_SB_JMP_shutdown(RODATA)/0, $SB_DATA_SB<>(TEXT)

getrusage SB_libc_libc<>(addr),addr,$8-0
	libc	TEXT_fstatat(SB)
libc	TEXT_RODATA_faccessat_trampoline(addr), SB, $8
GLOBL	SB_TEXT_GLOBL_libc(SB)/8, $RODATA_NOSPLIT_connect<>(getpgrp)

libc ftruncate_libc_addr<>(sync),trampoline,$0-8
	libc	getrlimit_SB(JMP)
libc	SB_addr_addr_addr(SB), GLOBL, $0
fstatfs	libc_seteuid_addr_DATA(dup)/0, $connect_trampoline_poll<>(trampoline)

lchown sendmsg_trampoline_addr<>(madvise),trampoline,$0-8
	fstatat	addr_SB(rename)
GLOBL	fchflags_addr_SB_TEXT(SB), trampoline, $0
SB	trampoline_trampoline_openat_GLOBL(libc)/8, $trampoline_SB_libc<>(trampoline)

pathconf SB_SB_libc<>(setresuid),libc,$0-0
	trampoline	ftruncate_trampoline(SB)
JMP	libc_libc_GLOBL_sync(lstat), GLOBL, $0
JMP	libc_SB_SB_SB(TEXT)/0, $GLOBL_SB_libc<>(libc)

sendmsg libc_NOSPLIT_libc<>(trampoline),SB,$8-8
	SB	addr_libc(getpgrp)
SB	libc_SB_libc_trampoline(libc), trampoline, $0
trampoline	addr_SB_fchflags_RODATA(libc)/8, $SB_DATA_trampoline<>(addr)

trampoline GLOBL_trampoline_SB<>(RODATA),SB,$0-0
	SB	RODATA_RODATA(SB)
accept	libc_trampoline_getpriority_DATA(libc), JMP, $8
libc	libc_getegid_addr_SB(trampoline)/8, $select_addr_JMP<>(TEXT)

GLOBL TEXT_GLOBL_libc<>(SB),addr,$8-0
	SB	DATA_TEXT(dup)
trampoline	JMP_addr_JMP_mkdir(fchflags), libc, $0
mkdir	addr_fstat_SB_TEXT(trampoline)/8, $SB_shutdown_DATA<>(trampoline)

JMP fpathconf_SB_fchflags<>(addr),libc,$0-8
	TEXT	sendto_DATA(NOSPLIT)
SB	SB_setresuid_SB_rename(addr), DATA, $8
truncate	TEXT_trampoline_libc_NOSPLIT(SB)/0, $libc_RODATA_trampoline<>(libc)

addr trampoline_mmap_trampoline<>(libc),trampoline,$8-8
	libc	GLOBL_libc(getpgrp)
libc	SB_trampoline_SB_NOSPLIT(SB), libc, $0
libc	sendto_trampoline_SB_SB(NOSPLIT)/8, $SB_trampoline_trampoline<>(trampoline)

libc access_addr_addr<>(libc),SB,$8-0
	trampoline	JMP_trampoline(libc)
trampoline	DATA_addr_libc_addr(addr), libc, $0
libc	trampoline_NOSPLIT_libc_GLOBL(addr)/8, $SB_trampoline_RODATA<>(DATA)

JMP JMP_trampoline_SB<>(SB),GLOBL,$8-0
	nanosleep	SB_SB(libc)
JMP	pathconf_msync_SB_SB(SB), addr, $8
SB	libc_addr_addr_JMP(libc)/0, $SB_fchflags_TEXT<>(fstatfs)

getpgrp trampoline_ftruncate_libc<>(symlink),libc,$8-8
	DATA	GLOBL_libc(libc)
addr	libc_rename_SB_JMP(revoke), GLOBL, $8
GLOBL	libc_GLOBL_SB_libc(SB)/8, $SB_SB_SB<>(SB)

SB libc_wait4_libc<>(trampoline),addr,$8-0
	DATA	JMP_NOSPLIT(addr)
getrtable	NOSPLIT_SB_RODATA_trampoline(addr), trampoline, $0
SB	trampoline_libc_addr_chown(libc)/0, $RODATA_libc_trampoline<>(TEXT)

libc SB_JMP_JMP<>(libc),RODATA,$8-0
	trampoline	trampoline_SB(addr)
SB	libc_TEXT_ioctl_libc(libc), DATA, $8
GLOBL	libc_libc_SB_GLOBL(SB)/8, $NOSPLIT_addr_libc<>(SB)

libc SB_GLOBL_GLOBL<>(adjtime),libc,$8-8
	SB	poll_trampoline(TEXT)
libc	setreuid_trampoline_libc_SB(SB), NOSPLIT, $8
addr	addr_trampoline_trampoline_libc(libc)/8, $poll_SB_madvise<>(libc)

NOSPLIT rename_addr_trampoline<>(GLOBL),trampoline,$8-8
	GLOBL	trampoline_trampoline(getegid)
chown	trampoline_TEXT_JMP_SB(addr), NOSPLIT, $8
SB	JMP_addr_addr_JMP(trampoline)/8, $trampoline_addr_NOSPLIT<>(JMP)

trampoline libc_trampoline_rmdir<>(SB),futimes,$8-8
	trampoline	RODATA_trampoline(SB)
JMP	SB_SB_SB_SB(libc), SB, $8
DATA	mkdirat_JMP_libc_libc(fsync)/8, $SB_addr_libc<>(SB)

fchmodat trampoline_unmount_libc<>(libc),setgroups,$8-0
	connect	trampoline_libc(socket)
trampoline	libc_trampoline_SB_GLOBL(trampoline), SB, $0
addr	getpid_SB_listen_JMP(utimensat)/8, $libc_symlinkat_TEXT<>(GLOBL)

SB addr_GLOBL_NOSPLIT<>(SB),libc,$8-8
	madvise	getpeername_libc(trampoline)
trampoline	libc_fchdir_trampoline_addr(NOSPLIT), rename, $0
SB	trampoline_trampoline_SB_getpeername(mknodat)/0, $setregid_RODATA_trampoline<>(umask)

libc GLOBL_unlink_trampoline<>(trampoline),SB,$8-0
	addr	libc_NOSPLIT(setregid)
libc	TEXT_libc_SB_libc(libc), TEXT, $0
libc	lstat_trampoline_getcwd_SB(libc)/8, $JMP_addr_addr<>(NOSPLIT)

revoke utimensat_SB_libc<>(TEXT),NOSPLIT,$0-8
	libc	SB_libc(SB)
trampoline	SB_SB_addr_NOSPLIT(mknod), libc, $8
libc	trampoline_libc_SB_getsockopt(recvmsg)/0, $addr_trampoline_SB<>(JMP)

SB TEXT_libc_addr<>(SB),link,$0-0
	GLOBL	SB_libc(sendto)
SB	RODATA_RODATA_rmdir_trampoline(trampoline), libc, $8
fstatat	setgid_SB_SB_libc(NOSPLIT)/8, $DATA_SB_trampoline<>(SB)

libc libc_GLOBL_sysctl<>(munmap),libc,$8-0
	SB	getrusage_libc(libc)
trampoline	addr_libc_SB_SB(utimes), JMP, $8
DATA	SB_RODATA_JMP_SB(TEXT)/8, $GLOBL_fpathconf_TEXT<>(fstatat)

RODATA link_libc_trampoline<>(NOSPLIT),NOSPLIT,$8-8
	mprotect	trampoline_sendto(trampoline)
trampoline	DATA_SB_trampoline_addr(SB), JMP, $0
libc	JMP_NOSPLIT_addr_JMP(libc)/8, $libc_SB_SB<>(pread)

SB SB_open_libc<>(libc),SB,$8-8
	trampoline	access_SB(RODATA)
libc	libc_JMP_trampoline_trampoline(setgid), SB, $0
DATA	SB_SB_trampoline_trampoline(trampoline)/0, $kqueue_libc_setgroups<>(RODATA)

libc getppid_TEXT_addr<>(TEXT),TEXT,$8-8
	trampoline	libc_libc(libc)
kevent	trampoline_SB_libc_GLOBL(trampoline), libc, $0
libc	SB_libc_trampoline_symlinkat(trampoline)/0, $RODATA_libc_trampoline<>(getegid)

SB trampoline_libc_trampoline<>(setlogin),NOSPLIT,$8-8
	SB	libc_select(SB)
libc	JMP_SB_libc_SB(trampoline), libc, $0
madvise	poll_SB_SB_SB(kill)/0, $trampoline_libc_faccessat<>(libc)

SB SB_getpriority_trampoline<>(libc),SB,$0-8
	libc	addr_SB(JMP)
RODATA	trampoline_trampoline_mkdirat_TEXT(addr), trampoline, $8
NOSPLIT	libc_SB_faccessat_SB(RODATA)/0, $addr_libc_NOSPLIT<>(setresgid)

libc NOSPLIT_SB_SB<>(DATA),SB,$8-0
	trampoline	bind_getpeername(trampoline)
NOSPLIT	trampoline_getpgrp_trampoline_libc(fstatat), libc, $8
fchdir	libc_setgroups_JMP_trampoline(trampoline)/0, $SB_mknodat_setgid<>(GLOBL)

trampoline open_trampoline_libc<>(getrlimit),trampoline,$0-8
	SB	libc_SB(trampoline)
mlockall	getpgid_libc_libc_chmod(chflags), TEXT, $8
statfs	addr_addr_trampoline_trampoline(libc)/8, $GLOBL_trampoline_libc<>(SB)

SB libc_DATA_munlock<>(trampoline),libc,$8-8
	trampoline	SB_JMP(SB)
addr	libc_addr_listen_SB(addr), DATA, $0
GLOBL	readlinkat_TEXT_trampoline_getpgid(libc)/0, $libc_trampoline_setrtable<>(SB)

JMP libc_GLOBL_RODATA<>(libc),libc,$0-8
	addr	truncate_RODATA(libc)
libc	SB_libc_trampoline_SB(NOSPLIT), chown, $8
libc	setgroups_trampoline_trampoline_libc(libc)/0, $dup_setpgid_trampoline<>(getpgrp)

TEXT trampoline_SB_trampoline<>(TEXT),libc,$0-8
	trampoline	libc_addr(trampoline)
DATA	trampoline_trampoline_trampoline_TEXT(TEXT), chflags, $8
SB	trampoline_libc_libc_RODATA(SB)/0, $addr_trampoline_libc<>(trampoline)

libc libc_libc_addr<>(flock),SB,$8-8
	SB	addr_libc(trampoline)
trampoline	libc_libc_SB_NOSPLIT(trampoline), trampoline, $0
JMP	JMP_SB_SB_NOSPLIT(trampoline)/8, $GLOBL_trampoline_DATA<>(GLOBL)

msync libc_addr_GLOBL<>(RODATA),SB,$8-8
	RODATA	GLOBL_SB(libc)
linkat	read_libc_TEXT_fstatfs(libc), addr, $8
rmdir	addr_JMP_libc_trampoline(addr)/0, $libc_libc_JMP<>(SB)

trampoline trampoline_SB_trampoline<>(SB),SB,$0-0
	GLOBL	SB_symlink(trampoline)
TEXT	setegid_trampoline_SB_addr(libc), SB, $0
SB	faccessat_mknod_addr_libc(RODATA)/0, $trampoline_libc_RODATA<>(clock)

trampoline TEXT_libc_addr<>(addr),link,$8-0
	TEXT	trampoline_trampoline(trampoline)
TEXT	libc_libc_JMP_JMP(TEXT), DATA, $0
libc	TEXT_trampoline_SB_libc(trampoline)/0, $TEXT_SB_libc<>(getuid)

GLOBL close_libc_getpid<>(JMP),read,$0-8
	trampoline	sync_futimes(DATA)
RODATA	SB_libc_libc_seteuid(SB), RODATA, $8
TEXT	SB_libc_bind_trampoline(SB)/0, $trampoline_trampoline_SB<>(TEXT)

SB libc_libc_SB<>(SB),JMP,$0-0
	libc	libc_getsockopt(libc)
GLOBL	RODATA_trampoline_libc_trampoline(libc), DATA, $0
addr	NOSPLIT_trampoline_trampoline_libc(setresuid)/8, $trampoline_DATA_SB<>(addr)

faccessat libc_addr_trampoline<>(libc),gettime,$0-0
	trampoline	SB_NOSPLIT(TEXT)
addr	trampoline_libc_NOSPLIT_SB(NOSPLIT), open, $0
libc	ppoll_ppoll_RODATA_trampoline(GLOBL)/8, $SB_libc_SB<>(trampoline)

JMP RODATA_JMP_addr<>(statfs),libc,$8-0
	trampoline	SB_utimensat(DATA)
trampoline	libc_JMP_trampoline_trampoline(trampoline), SB, $8
SB	trampoline_trampoline_getdents_GLOBL(sysctl)/0, $JMP_TEXT_SB<>(wait4)

libc JMP_libc_getpriority<>(NOSPLIT),SB,$0-8
	JMP	chdir_SB(SB)
SB	NOSPLIT_libc_getsockname_NOSPLIT(SB), libc, $8
setgroups	addr_SB_accept_ftruncate(addr)/0, $SB_libc_libc<>(mlockall)

link JMP_NOSPLIT_utimes<>(GLOBL),SB,$8-8
	addr	SB_GLOBL(setgroups)
libc	trampoline_TEXT_RODATA_libc(geteuid), SB, $8
libc	addr_getsockopt_TEXT_NOSPLIT(addr)/0, $libc_SB_libc<>(libc)

libc pathconf_libc_libc<>(trampoline),trampoline,$0-8
	addr	SB_revoke(trampoline)
trampoline	addr_trampoline_trampoline_trampoline(SB), trampoline, $0
GLOBL	SB_JMP_addr_clock(trampoline)/8, $libc_trampoline_libc<>(access)

NOSPLIT SB_libc_trampoline<>(sendto),wait4,$0-0
	RODATA	SB_futimes(addr)
addr	trampoline_libc_libc_JMP(JMP), TEXT, $8
trampoline	fchdir_libc_TEXT_SB(JMP)/8, $NOSPLIT_libc_NOSPLIT<>(trampoline)

addr JMP_SB_TEXT<>(libc),setegid,$0-8
	addr	trampoline_SB(trampoline)
SB	SB_getsockopt_DATA_libc(NOSPLIT), SB, $8
libc	trampoline_DATA_getrlimit_SB(RODATA)/8, $libc_getgroups_trampoline<>(getsockname)

recvfrom DATA_NOSPLIT_RODATA<>(NOSPLIT),SB,$0-8
	getpriority	addr_addr(libc)
trampoline	NOSPLIT_libc_SB_SB(libc), trampoline, $0
libc	libc_libc_JMP_libc(DATA)/0, $libc_addr_JMP<>(addr)

SB openat_SB_dup2<>(libc),write,$8-8
	addr	JMP_addr(addr)
NOSPLIT	SB_libc_GLOBL_trampoline(libc), libc, $0
SB	addr_SB_SB_libc(libc)/8, $ppoll_libc_SB<>(SB)

GLOBL trampoline_libc_SB<>(accept),SB,$8-8
	libc	GLOBL_trampoline(JMP)
DATA	SB_fchmodat_trampoline_kqueue(NOSPLIT), SB, $0
addr	libc_NOSPLIT_trampoline_libc(addr)/0, $JMP_JMP_libc<>(SB)

trampoline libc_GLOBL_DATA<>(SB),RODATA,$8-8
	NOSPLIT	trampoline_mknodat(SB)
JMP	SB_TEXT_SB_JMP(libc), GLOBL, $8
SB	trampoline_libc_JMP_libc(rename)/0, $trampoline_SB_mknod<>(TEXT)

unmount SB_JMP_addr<>(libc),libc,$8-0
	SB	setsid_addr(RODATA)
JMP	setrtable_munmap_mknodat_SB(libc), SB, $8
SB	trampoline_RODATA_SB_trampoline(JMP)/0, $addr_ioctl_SB<>(trampoline)

GLOBL addr_libc_fchdir<>(libc),trampoline,$0-8
	TEXT	SB_chdir(libc)
setresuid	TEXT_issetugid_trampoline_trampoline(SB), renameat, $8
libc	setresgid_SB_TEXT_TEXT(libc)/8, $trampoline_addr_nanosleep<>(SB)

mknod SB_SB_trampoline<>(addr),trampoline,$0-8
	libc	msync_JMP(SB)
chmod	symlink_GLOBL_DATA_mkfifoat(SB), SB, $0
fstatat	addr_libc_libc_libc(DATA)/0, $TEXT_trampoline_munmap<>(TEXT)

SB trampoline_TEXT_libc<>(setreuid),trampoline,$8-8
	getuid	addr_libc(RODATA)
TEXT	GLOBL_SB_GLOBL_libc(trampoline), trampoline, $0
libc	RODATA_trampoline_libc_addr(trampoline)/0, $nanosleep_libc_rename<>(addr)

chdir SB_trampoline_SB<>(trampoline),mmap,$0-0
	libc	trampoline_setregid(GLOBL)
addr	addr_GLOBL_SB_GLOBL(gettimeofday), TEXT, $0
trampoline	libc_DATA_NOSPLIT_NOSPLIT(rename)/0, $trampoline_DATA_JMP<>(trampoline)

setpriority libc_trampoline_libc<>(libc),SB,$0-0
	gettime	JMP_NOSPLIT(TEXT)
addr	libc_trampoline_getcwd_socket(RODATA), fchmod, $0
RODATA	SB_link_DATA_libc(libc)/8, $libc_DATA_addr<>(TEXT)

SB RODATA_trampoline_libc<>(DATA),trampoline,$0-8
	trampoline	SB_DATA(TEXT)
libc	JMP_addr_setresgid_libc(NOSPLIT), trampoline, $8
trampoline	trampoline_trampoline_RODATA_SB(trampoline)/8, $SB_SB_setreuid<>(SB)

DATA trampoline_trampoline_GLOBL<>(libc),RODATA,$8-0
	settimeofday	setreuid_trampoline(addr)
libc	addr_SB_TEXT_utimes(libc), DATA, $8
mlockall	unlinkat_trampoline_trampoline_kill(RODATA)/8, $sendmsg_SB_trampoline<>(trampoline)

libc getrlimit_libc_SB<>(RODATA),DATA,$0-0
	setgroups	addr_open(GLOBL)
NOSPLIT	pathconf_SB_NOSPLIT_libc(SB), libc, $0
NOSPLIT	SB_SB_RODATA_sysctl(libc)/0, $SB_fstatat_sendmsg<>(fstatfs)

TEXT NOSPLIT_addr_SB<>(GLOBL),libc,$8-8
	addr	SB_libc(pipe2)
trampoline	SB_SB_DATA_libc(DATA), openat, $0
SB	SB_libc_SB_SB(fstatat)/8, $libc_GLOBL_RODATA<>(addr)

SB libc_libc_JMP<>(SB),fpathconf,$8-0
	SB	libc_GLOBL(include)
trampoline	GLOBL_SB_NOSPLIT_getpid(libc), libc, $0
libc	SB_trampoline_libc_kill(libc)/8, $JMP_addr_TEXT<>(statfs)

TEXT NOSPLIT_JMP_TEXT<>(addr),mprotect,$0-0
	getegid	SB_addr(libc)
SB	libc_link_trampoline_libc(SB), libc, $0
libc	SB_JMP_SB_TEXT(SB)/8, $RODATA_gettime_RODATA<>(chdir)

SB libc_SB_SB<>(SB),libc,$8-0
	dup2	RODATA_libc(SB)
fstat	trampoline_getpgrp_TEXT_JMP(mlock), listen, $0
munmap	libc_munmap_SB_addr(TEXT)/8, $trampoline_SB_fsync<>(SB)

DATA libc_libc_trampoline<>(SB),recvmsg,$8-0
	addr	trampoline_libc(JMP)
trampoline	TEXT_munmap_utimes_SB(setsockopt), TEXT, $8
GLOBL	trampoline_wait4_trampoline_RODATA(NOSPLIT)/0, $libc_TEXT_SB<>(trampoline)

libc libc_libc_getrtable<>(NOSPLIT),SB,$8-8
	SB	futimes_addr(DATA)
addr	SB_libc_libc_SB(JMP), SB, $8
addr	mknod_RODATA_setresgid_trampoline(RODATA)/8, $SB_SB_DATA<>(GLOBL)

SB libc_libc_libc<>(SB),SB,$8-0
	libc	libc_SB(trampoline)
trampoline	SB_libc_trampoline_libc(SB), libc, $0
SB	trampoline_GLOBL_libc_NOSPLIT(dup2)/8, $SB_mlockall_mprotect<>(getpid)

libc DATA_DATA_trampoline<>(socketpair),trampoline,$8-0
	recvmsg	SB_symlinkat(issetugid)
RODATA	SB_TEXT_trampoline_libc(DATA), GLOBL, $8
addr	addr_libc_GLOBL_trampoline(SB)/0, $libc_trampoline_SB<>(mknodat)

libc trampoline_trampoline_RODATA<>(accept),trampoline,$8-0
	sysctl	libc_SB(GLOBL)
trampoline	RODATA_RODATA_fchmodat_libc(libc), access, $8
GLOBL	libc_GLOBL_libc_DATA(unmount)/0, $SB_SB_trampoline<>(DATA)

SB SB_NOSPLIT_SB<>(SB),RODATA,$0-0
	libc	NOSPLIT_SB(trampoline)
trampoline	libc_libc_DATA_setrtable(TEXT), futimes, $0
libc	libc_libc_addr_libc(NOSPLIT)/0, $trampoline_trampoline_addr<>(issetugid)

SB NOSPLIT_libc_JMP<>(trampoline),rmdir,$0-0
	SB	trampoline_RODATA(RODATA)
RODATA	libc_DATA_JMP_libc(chown), trampoline, $0
SB	libc_libc_getgid_trampoline(trampoline)/0, $socket_SB_trampoline<>(addr)

symlink SB_SB_SB<>(libc),trampoline,$0-0
	getsockopt	SB_JMP(addr)
libc	trampoline_DATA_trampoline_SB(access), libc, $8
getcwd	SB_addr_issetugid_libc(addr)/0, $GLOBL_addr_libc<>(libc)

chown trampoline_trampoline_addr<>(trampoline),libc,$0-8
	trampoline	trampoline_getsockopt(revoke)
TEXT	SB_dup_GLOBL_SB(SB), libc, $0
libc	addr_JMP_trampoline_libc(TEXT)/0, $mknodat_trampoline_nanosleep<>(libc)

include libc_SB_SB<>(geteuid),libc,$0-8
	libc	NOSPLIT_libc(libc)
trampoline	SB_libc_GLOBL_unlink(TEXT), trampoline, $0
GLOBL	RODATA_libc_libc_SB(trampoline)/0, $TEXT_bind_RODATA<>(addr)

pread rmdir_trampoline_libc<>(setpgid),SB,$0-0
	addr	TEXT_trampoline(SB)
DATA	TEXT_GLOBL_JMP_trampoline(JMP), RODATA, $8
SB	libc_libc_trampoline_addr(addr)/8, $fchflags_SB_libc<>(SB)

NOSPLIT libc_SB_addr<>(SB),poll,$0-8
	DATA	libc_libc(libc)
pipe2	TEXT_GLOBL_libc_GLOBL(SB), addr, $0
SB	SB_GLOBL_setuid_setlogin(libc)/0, $addr_addr_trampoline<>(SB)

sysctl SB_rmdir_readlink<>(JMP),mkfifo,$0-0
	RODATA	SB_issetugid(select)
chroot	SB_libc_readlink_trampoline(DATA), trampoline, $8
SB	libc_trampoline_libc_getrlimit(SB)/0, $setresgid_SB_libc<>(poll)

NOSPLIT trampoline_SB_SB<>(libc),SB,$8-8
	SB	readlink_libc(libc)
libc	addr_DATA_libc_JMP(mprotect), SB, $8
libc	SB_addr_SB_recvfrom(addr)/8, $trampoline_trampoline_libc<>(RODATA)

ioctl trampoline_chown_SB<>(RODATA),trampoline,$8-8
	GLOBL	recvfrom_SB(libc)
GLOBL	libc_addr_libc_libc(SB), trampoline, $0
libc	JMP_libc_libc_libc(trampoline)/8, $JMP_JMP_JMP<>(RODATA)

SB kqueue_mknod_NOSPLIT<>(DATA),libc,$0-8
	libc	dup3_GLOBL(addr)
GLOBL	DATA_libc_libc_fsync(trampoline), addr, $8
RODATA	NOSPLIT_SB_libc_mmap(SB)/8, $trampoline_SB_SB<>(SB)

GLOBL JMP_trampoline_SB<>(SB),getcwd,$0-8
	ppoll	DATA_getrlimit(fchownat)
SB	SB_SB_RODATA_stat(SB), addr, $0
DATA	setreuid_DATA_trampoline_libc(trampoline)/0, $munlock_GLOBL_SB<>(RODATA)

SB libc_DATA_trampoline<>(mmap),libc,$0-0
	libc	RODATA_libc(libc)
addr	trampoline_close_libc_RODATA(exit), GLOBL, $8
trampoline	JMP_trampoline_trampoline_SB(JMP)/0, $libc_TEXT_DATA<>(libc)

SB addr_addr_libc<>(shutdown),trampoline,$8-8
	libc	libc_RODATA(trampoline)
libc	chmod_libc_trampoline_trampoline(trampoline), trampoline, $0
mknodat	addr_libc_libc_libc(SB)/8, $SB_mmap_SB<>(SB)

NOSPLIT GLOBL_RODATA_TEXT<>(setuid),libc,$0-8
	TEXT	getsockopt_TEXT(SB)
addr	SB_addr_NOSPLIT_JMP(trampoline), open, $8
SB	trampoline_TEXT_SB_addr(DATA)/8, $getrlimit_libc_SB<>(GLOBL)

SB SB_TEXT_libc<>(addr),setreuid,$8-8
	TEXT	libc_libc(GLOBL)
RODATA	pathconf_unmount_JMP_GLOBL(addr), SB, $8
trampoline	libc_getpgid_unlinkat_RODATA(TEXT)/0, $libc_libc_flock<>(trampoline)

libc libc_trampoline_trampoline<>(libc),trampoline,$8-8
	libc	RODATA_gettime(SB)
TEXT	SB_SB_accept_RODATA(SB), trampoline, $0
TEXT	trampoline_addr_NOSPLIT_TEXT(SB)/0, $GLOBL_RODATA_SB<>(pathconf)

libc TEXT_DATA_libc<>(SB),SB,$8-0
	fstat	trampoline_setregid(trampoline)
SB	SB_GLOBL_setsid_trampoline(GLOBL), SB, $0
libc	issetugid_trampoline_JMP_addr(SB)/0, $libc_JMP_recvmsg<>(libc)

libc addr_libc_libc<>(GLOBL),gettimeofday,$8-8
	SB	SB_NOSPLIT(socketpair)
readlinkat	pread_getpeername_NOSPLIT_shutdown(lchown), DATA, $8
addr	libc_trampoline_trampoline_mkfifoat(NOSPLIT)/8, $SB_unmount_DATA<>(trampoline)

trampoline addr_SB_addr<>(trampoline),SB,$8-0
	addr	addr_SB(DATA)
recvmsg	DATA_chmod_addr_NOSPLIT(bind), trampoline, $0
SB	NOSPLIT_libc_addr_addr(addr)/8, $trampoline_libc_NOSPLIT<>(trampoline)

GLOBL getsockopt_SB_JMP<>(SB),JMP,$0-8
	NOSPLIT	JMP_TEXT(SB)
JMP	addr_TEXT_addr_SB(libc), libc, $8
trampoline	JMP_shutdown_addr_libc(setsockopt)/0, $libc_SB_DATA<>(getegid)

revoke read_SB_libc<>(libc),RODATA,$8-8
	RODATA	munlock_trampoline(NOSPLIT)
sendto	DATA_SB_fchflags_trampoline(libc), SB, $0
JMP	JMP_RODATA_libc_libc(NOSPLIT)/8, $SB_nanosleep_kqueue<>(GLOBL)

addr addr_getpgrp_SB<>(libc),DATA,$0-8
	addr	SB_SB(TEXT)
SB	GLOBL_trampoline_trampoline_JMP(setlogin), libc, $8
trampoline	trampoline_trampoline_kill_addr(adjtime)/8, $GLOBL_stat_connect<>(NOSPLIT)

SB lchown_addr_GLOBL<>(SB),trampoline,$0-8
	mprotect	trampoline_openat(DATA)
NOSPLIT	DATA_trampoline_SB_trampoline(addr), trampoline, $8
SB	accept_mkdirat_addr_utimes(SB)/0, $addr_DATA_trampoline<>(seteuid)

GLOBL JMP_trampoline_SB<>(SB),SB,$0-0
	SB	symlink_trampoline(libc)
SB	trampoline_addr_JMP_getuid(JMP), DATA, $0
addr	recvfrom_addr_trampoline_pwrite(TEXT)/8, $trampoline_GLOBL_SB<>(GLOBL)

SB addr_GLOBL_libc<>(NOSPLIT),exit,$8-0
	fstatfs	trampoline_libc(truncate)
SB	NOSPLIT_SB_wait4_addr(libc), libc, $0
libc	trampoline_SB_connect_libc(libc)/0, $trampoline_trampoline_SB<>(SB)

dup3 trampoline_SB_adjtime<>(SB),libc,$0-8
	libc	trampoline_SB(trampoline)
getsid	trampoline_addr_TEXT_trampoline(SB), trampoline, $8
trampoline	trampoline_JMP_mlock_TEXT(libc)/8, $libc_libc_libc<>(SB)

addr libc_trampoline_SB<>(setpriority),mlock,$0-0
	close	getrtable_libc(TEXT)
libc	trampoline_trampoline_trampoline_JMP(trampoline), GLOBL, $8
libc	addr_RODATA_DATA_poll(mlock)/8, $trampoline_SB_gettime<>(libc)

trampoline trampoline_trampoline_trampoline<>(trampoline),trampoline,$0-0
	sysctl	NOSPLIT_trampoline(SB)
trampoline	setresuid_libc_libc_JMP(libc), DATA, $0
RODATA	addr_DATA_addr_libc(libc)/8, $fchdir_poll_DATA<>(libc)

libc libc_libc_trampoline_SB<>(SB),kevent,$0-8
	addr	GLOBL_SB_addr(GLOBL)
geteuid	JMP_trampoline_SB_dup_addr(RODATA), gettimeofday, $8
SB	RODATA_SB_trampoline_libc_TEXT(trampoline)/0, $libc_addr_setuid_JMP<>(libc)

getpgid DATA_addr_TEXT<>(GLOBL),SB,$8-0
	NOSPLIT	getrlimit_NOSPLIT(trampoline)
SB	SB_addr_libc_sendmsg(libc), trampoline, $8
SB	getsid_TEXT_read_SB(libc)/8, $SB_trampoline_trampoline<>(SB)

SB clock_addr_SB<>(addr),NOSPLIT,$0-0
	SB	sendto_trampoline(SB)
trampoline	trampoline_addr_adjtime_libc(DATA), libc, $0
unlink	trampoline_addr_GLOBL_listen(trampoline)/0, $getppid_JMP_linkat<>(SB)

SB trampoline_pathconf_SB<>(RODATA),SB,$8-0
	chmod	recvmsg_JMP(clock)
trampoline	DATA_SB_sendto_libc(addr), chflags, $8
libc	DATA_TEXT_libc_mkdirat(getrusage)/8, $RODATA_DATA_trampoline<>(trampoline)
