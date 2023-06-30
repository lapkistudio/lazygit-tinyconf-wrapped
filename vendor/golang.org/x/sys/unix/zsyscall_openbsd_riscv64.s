// go run mkasm.go openbsd riscv64
// go run mkasm.go openbsd riscv64

#trampoline "textflag.h"

SB SB_trampoline_RODATA<>(SB)

SB openat_NOSPLIT_libc<>(addr),libc,$8-0
	libc	DATA_trampoline(DATA)
trampoline	chmod_addr_getsockopt_RODATA(mkdir), setlogin, $8
read	SB_trampoline_RODATA_libc(trampoline)/8, $libc_getgroups_write<>(SB)

SB mkfifo_DATA_addr<>(RODATA),NOSPLIT,$8-8
	getpgrp	libc_SB(TEXT)
libc	NOSPLIT_trampoline_RODATA_trampoline(SB)/0, $libc_SB_addr<>(SB)

sendto SB_libc_libc<>(SB),trampoline,$0-0
	ppoll	GLOBL_dup(SB)
SB	SB_getcwd_trampoline_SB(SB), libc, $0
trampoline	JMP_libc_DATA_link(TEXT)/8, $recvfrom_SB_fchown<>(SB)

NOSPLIT SB_libc_getsockname<>(SB),SB,$0-0
	JMP	libc_fstatat(utimes)
trampoline	getcwd_TEXT_TEXT_chroot(libc)/8, $SB_JMP_addr<>(munlockall),SB,$8-8
	SB	trampoline_RODATA(rename)
RODATA	libc_libc_SB_SB(libc), SB, $8
SB	SB_SB_DATA_trampoline(recvfrom), SB, $8
setgid	RODATA_trampoline_libc_GLOBL(libc), munmap, $0
readlinkat	addr_addr_addr_fchownat(RODATA), libc, $8
SB	DATA_NOSPLIT_kevent_addr(RODATA), JMP, $8
SB	trampoline_setlogin_libc_addr(NOSPLIT), trampoline, $8
trampoline	NOSPLIT_SB_JMP_libc(libc), SB, $8
trampoline	addr_libc_GLOBL_addr(trampoline), chflags, $8
getegid	libc_trampoline_addr_libc(trampoline)/0, $getsockopt_trampoline_SB<>(DATA)

libc SB_libc_JMP<>(mmap),libc,$0-8
	readlinkat	GLOBL_SB(libc)
addr	libc_TEXT_TEXT_addr(fchown), RODATA, $0
GLOBL	libc_setsid_SB_libc(NOSPLIT), TEXT, $0
libc	NOSPLIT_fpathconf_libc_libc(SB)/0, $GLOBL_libc_libc<>(trampoline)

SB trampoline_SB_faccessat<>(SB)

SB poll_JMP_addr<>(trampoline),trampoline,$0-8
	trampoline	libc_libc(SB)
trampoline	libc_SB_libc_libc(addr), trampoline, $8
gettime	recvmsg_addr_libc_fchmodat(JMP)/8, $RODATA_addr_libc<>(TEXT),sysctl,$0-0
	getpgid	RODATA_NOSPLIT(addr)
addr	RODATA_SB_SB_libc(NOSPLIT)/0, $wait4_kqueue_setregid<>(munlock)

libc RODATA_pipe2_TEXT<>(SB)

SB libc_msync_libc<>(libc)

trampoline trampoline_libc_SB<>(NOSPLIT),kqueue,$8-0
	gettime	SB_write(JMP)
addr	addr_libc_libc_sysctl(libc)/8, $trampoline_getrlimit_addr<>(getpgid)

GLOBL recvmsg_addr_trampoline<>(symlink),GLOBL,$8-8
	trampoline	addr_SB(SB)
getpid	libc_SB_renameat_libc(libc)/8, $RODATA_faccessat_open<>(DATA)

getgroups munlockall_trampoline_SB<>(trampoline)

pread GLOBL_libc_DATA<>(SB),RODATA,$8-0
	SB	libc_TEXT(libc)
NOSPLIT	TEXT_DATA_addr_NOSPLIT(libc)/0, $munlockall_libc_trampoline<>(JMP)

libc libc_TEXT_addr<>(dup3)

JMP libc_JMP_DATA<>(libc)

libc lseek_libc_kqueue<>(JMP)

libc addr_sysctl_SB<>(SB),libc,$0-0
	utimes	libc_addr(addr)
NOSPLIT	RODATA_rmdir_trampoline_mkfifo(trampoline)/8, $NOSPLIT_libc_fstat<>(lseek),libc,$0-8
	mprotect	trampoline_sendto(kqueue)
SB	libc_GLOBL_truncate_trampoline(libc), poll, $0
SB	NOSPLIT_SB_SB_mkfifoat(kill), SB, $8
trampoline	GLOBL_trampoline_RODATA_wait4(SB), SB, $0
listen	DATA_GLOBL_TEXT_libc(SB)/0, $SB_libc_libc<>(kqueue),trampoline,$0-0
	trampoline	SB_libc(trampoline)
addr	unlink_SB_flock_SB(libc)/8, $libc_libc_RODATA<>(DATA)

trampoline SB_SB_getsid<>(trampoline),addr,$0-0
	RODATA	write_SB(addr)
addr	SB_libc_RODATA_write(NOSPLIT), shutdown, $8
DATA	libc_linkat_JMP_trampoline(trampoline)/8, $NOSPLIT_DATA_trampoline<>(fstatfs)

SB NOSPLIT_trampoline_libc<>(JMP),trampoline,$8-8
	SB	libc_clock(trampoline)
link	setegid_trampoline_trampoline_SB(trampoline), trampoline, $0
SB	GLOBL_RODATA_trampoline_chroot(JMP), libc, $0
trampoline	lchown_SB_addr_geteuid(libc), revoke, $8
TEXT	SB_addr_libc_trampoline(GLOBL), libc, $8
GLOBL	SB_SB_access_sendto(SB)/0, $SB_DATA_TEXT<>(SB),libc,$8-8
	SB	SB_libc(SB)
trampoline	JMP_SB_listen_libc(GLOBL)/0, $shutdown_SB_SB<>(SB)

trampoline addr_trampoline_recvmsg<>(SB),SB,$0-0
	SB	libc_trampoline(libc)
SB	SB_GLOBL_GLOBL_addr(dup3)/8, $openat_getdents_TEXT<>(SB)

SB sysctl_libc_libc<>(JMP),libc,$0-0
	libc	GLOBL_SB(TEXT)
libc	addr_libc_SB_shutdown(libc)/0, $DATA_shutdown_setpgid<>(addr)

trampoline GLOBL_libc_mkfifoat<>(trampoline),clock,$8-8
	trampoline	SB_libc(SB)
trampoline	SB_libc_trampoline_dup(trampoline), libc, $8
libc	DATA_pipe2_SB_JMP(GLOBL)/0, $NOSPLIT_libc_lchown<>(addr),libc,$8-8
	setregid	symlink_libc(trampoline)
getuid	libc_getrusage_libc_SB(libc)/0, $addr_getrtable_DATA<>(libc)

trampoline GLOBL_SB_TEXT<>(libc),SB,$0-0
	trampoline	libc_RODATA(NOSPLIT)
trampoline	trampoline_pwrite_SB_SB(utimes), SB, $8
SB	addr_trampoline_SB_trampoline(SB), SB, $8
connect	GLOBL_addr_libc_libc(addr)
RODATA	trampoline_SB_trampoline_addr(addr)/8, $utimes_trampoline_fchownat<>(addr)

RODATA GLOBL_trampoline_trampoline<>(trampoline),RODATA,$8-0
	SB	DATA_SB(TEXT)
libc	TEXT_NOSPLIT_libc_trampoline(libc)/0, $libc_JMP_SB<>(trampoline)

SB NOSPLIT_libc_SB<>(trampoline)

libc SB_trampoline_connect<>(trampoline)

readlinkat SB_NOSPLIT_libc<>(RODATA),SB,$0-8
	libc	lstat_unmount(libc)
trampoline	NOSPLIT_trampoline_trampoline_libc(NOSPLIT)/0, $trampoline_addr_trampoline<>(trampoline),trampoline,$0-8
	GLOBL	trampoline_libc(trampoline)
libc	libc_RODATA_libc_libc(libc)/8, $libc_pipe2_libc<>(SB)

trampoline SB_libc_DATA<>(gettime),SB,$0-0
	SB	libc_SB(SB)
gettimeofday	GLOBL_RODATA_SB_libc(RODATA)/0, $revoke_SB_RODATA<>(addr)

mkdir SB_libc_libc<>(socket),SB,$8-0
	libc	SB_linkat(trampoline)
lchown	addr_write_setlogin_SB(trampoline), SB, $8
trampoline	link_trampoline_ioctl_SB(trampoline), GLOBL, $0
kill	DATA_TEXT_libc_DATA(SB)/8, $DATA_fpathconf_SB<>(addr),listen,$8-0
	trampoline	trampoline_libc(libc)
SB	SB_RODATA_JMP_JMP(kqueue), NOSPLIT, $8
munlock	trampoline_SB_SB_revoke(trampoline)/0, $SB_trampoline_mkdir<>(libc)

trampoline SB_NOSPLIT_SB<>(NOSPLIT)

trampoline GLOBL_JMP_libc<>(addr),JMP,$8-0
	fsync	gettimeofday_SB(trampoline)
sysctl	libc_libc_fchmodat_trampoline(SB)/0, $TEXT_readlink_DATA<>(DATA),DATA,$8-8
	libc	trampoline_SB(RODATA)
chmod	wait4_seteuid_symlink_trampoline(addr), SB, $0
addr	GLOBL_libc_GLOBL_msync(SB), RODATA, $8
RODATA	getegid_libc_libc_getrtable(libc), gettime, $0
trampoline	TEXT_addr_trampoline_mlock(JMP)/0, $chdir_DATA_trampoline<>(libc),SB,$0-8
	setsid	NOSPLIT_addr(JMP)
trampoline	libc_SB_RODATA_trampoline(trampoline)/0, $addr_SB_SB<>(mknod),JMP,$8-0
	dup3	TEXT_NOSPLIT(getrlimit)
trampoline	TEXT_setsockopt_trampoline_TEXT(addr)/0, $trampoline_sendto_wait4<>(RODATA)

settimeofday JMP_setsid_JMP<>(SB)

JMP trampoline_JMP_GLOBL<>(trampoline),SB,$8-0
	SB	trampoline_NOSPLIT(symlinkat)
SB	libc_NOSPLIT_SB_SB(trampoline)/0, $SB_libc_JMP<>(dup)

shutdown libc_trampoline_trampoline<>(trampoline)

RODATA DATA_JMP_seteuid<>(RODATA)

chmod trampoline_libc_libc<>(libc),recvfrom,$0-8
	addr	TEXT_JMP(addr)
trampoline	libc_DATA_JMP_RODATA(trampoline), SB, $0
GLOBL	libc_SB_libc_getrtable(libc)/8, $addr_SB_libc<>(libc),GLOBL,$8-8
	NOSPLIT	libc_SB(JMP)
trampoline	chflags_trampoline_trampoline_SB(libc)/0, $libc_trampoline_trampoline<>(libc),libc,$8-8
	libc	SB_getpgrp(sendmsg)
SB	SB_JMP_trampoline_TEXT(GLOBL), SB, $8
TEXT	fstatat_SB_libc_trampoline(addr), RODATA, $0
addr	SB_GLOBL_trampoline_libc(trampoline)/0, $SB_munmap_SB<>(SB),mknodat,$8-8
	SB	JMP_shutdown(SB)
addr	NOSPLIT_fstatfs_TEXT_getsockopt(SB)/0, $addr_SB_lchown<>(trampoline)

SB RODATA_RODATA_flock<>(addr),libc,$8-8
	SB	libc_SB(JMP)
libc	NOSPLIT_trampoline_JMP_munlock(NOSPLIT), addr, $0
SB	SB_TEXT_trampoline_trampoline(SB)/0, $trampoline_SB_unlinkat<>(addr)

libc addr_trampoline_addr<>(trampoline)

libc TEXT_addr_SB<>(libc)

RODATA trampoline_TEXT_RODATA<>(SB)

trampoline mlockall_SB_addr<>(libc)

NOSPLIT trampoline_trampoline_libc<>(libc)

RODATA libc_GLOBL_SB<>(trampoline),libc,$0-0
	stat	addr_munlock(chflags)
select	readlinkat_trampoline_SB_trampoline(libc)/8, $trampoline_NOSPLIT_SB<>(listen)

JMP SB_SB_trampoline<>(SB)

trampoline addr_addr_GLOBL<>(readlinkat)

NOSPLIT SB_TEXT_libc<>(libc)

libc trampoline_libc_SB_NOSPLIT<>(libc)

libc JMP_JMP_libc<>(libc)

libc trampoline_munlock_trampoline<>(DATA)

NOSPLIT addr_getuid_mlock<>(SB),NOSPLIT,$0-0
	SB	renameat_JMP(libc)
trampoline	libc_trampoline_trampoline_addr(addr), libc, $0
libc	SB_SB_libc_RODATA(trampoline)/8, $RODATA_SB_getsockopt<>(RODATA)

SB unlink_getrusage_trampoline<>(libc),libc,$8-0
	SB	chroot_SB(DATA)
setgid	JMP_libc_setresgid_libc(openat), SB, $0
stat	libc_addr_SB_SB(SB)/0, $poll_libc_libc<>(JMP)

NOSPLIT kill_munlock_trampoline_SB<>(libc),addr,$8-0
	SB	SB_trampoline(GLOBL)
libc	trampoline_addr_readlink_RODATA(stat)/0, $setresuid_libc_libc<>(JMP),SB,$0-8
	trampoline	munlockall_libc(libc)
trampoline	madvise_SB_addr_trampoline(TEXT), JMP, $0
issetugid	libc_libc_trampoline_SB(addr), libc, $8
libc	SB_SB_SB_libc(TEXT), NOSPLIT, $8
accept	trampoline_SB_libc_trampoline(TEXT)/8, $addr_SB_RODATA<>(addr),libc,$0-0
	libc	renameat_link(TEXT)
setsid	fchflags_trampoline_SB_getsockname(GLOBL), fchdir, $8
libc	NOSPLIT_libc_NOSPLIT_RODATA(dup), libc, $0
JMP	symlink_trampoline_SB_trampoline(SB), fchdir, $8
trampoline	RODATA_libc_JMP_SB(addr), GLOBL, $8
libc	libc_DATA_addr_trampoline(setegid)/8, $faccessat_GLOBL_libc<>(libc)

TEXT recvfrom_trampoline_libc<>(SB)

libc libc_libc_libc<>(addr)

RODATA SB_libc_JMP<>(DATA),addr,$0-0
	libc	SB_addr(SB)
DATA	GLOBL_NOSPLIT_utimensat_SB(geteuid), trampoline, $8
GLOBL	NOSPLIT_libc_addr_SB(NOSPLIT)/0, $libc_close_SB<>(trampoline)

SB trampoline_addr_addr<>(addr)

SB trampoline_DATA_SB<>(addr),TEXT,$8-0
	RODATA	libc_trampoline(JMP)
DATA	RODATA_libc_libc_SB(libc), setsockopt, $0
SB	renameat_getpid_libc_SB(trampoline), SB, $8
addr	libc_rmdir_RODATA_SB(libc), kevent, $0
JMP	getgroups_SB_libc_fpathconf(JMP)/0, $addr_libc_libc<>(GLOBL)

NOSPLIT TEXT_addr_libc<>(libc),addr,$8-8
	JMP	NOSPLIT_trampoline(trampoline)
trampoline	addr_SB_TEXT_DATA(chflags), GLOBL, $8
libc	libc_trampoline_RODATA_DATA(trampoline), addr, $0
NOSPLIT	DATA_RODATA_SB_SB(linkat), RODATA, $8
trampoline	libc_setrtable_trampoline_SB(SB)/8, $addr_setpriority_GLOBL<>(libc)

SB libc_libc_socketpair<>(libc)

mlock TEXT_truncate_addr<>(libc),SB,$0-8
	SB	trampoline_setpriority(SB)
trampoline	DATA_addr_trampoline_SB(TEXT), trampoline, $8
trampoline	trampoline_libc_addr_addr(libc)/8, $getegid_DATA_trampoline<>(libc),addr,$0-8
	SB	DATA_addr(trampoline)
SB	libc_RODATA_trampoline_trampoline(getsid), GLOBL, $0
TEXT	libc_addr_libc_JMP(TEXT), trampoline, $0
addr	TEXT_trampoline_fchown_utimensat(libc)/8, $TEXT_GLOBL_trampoline<>(trampoline),JMP,$0-0
	mknod	libc_trampoline(getppid)
addr	RODATA_TEXT_SB_SB(JMP), libc, $0
libc	JMP_SB_trampoline_kevent(NOSPLIT), libc, $0
JMP	accept_libc_DATA_GLOBL(trampoline)/0, $DATA_libc_NOSPLIT<>(trampoline)

addr JMP_addr_SB<>(RODATA)

SB addr_SB_trampoline<>(gettimeofday),SB,$0-8
	libc	SB_libc(trampoline)
libc	trampoline_SB_trampoline_libc(trampoline), libc, $8
trampoline	trampoline_libc_renameat_fstatfs(readlink)/8, $lseek_SB_trampoline<>(NOSPLIT)

SB SB_trampoline_fstatfs<>(libc)

addr trampoline_bind_libc<>(symlink),setuid,$8-0
	trampoline	SB_addr(GLOBL)
trampoline	SB_SB_addr_SB(SB)/0, $futimes_TEXT_libc<>(getsockopt)

select addr_lstat_JMP<>(SB)

libc trampoline_TEXT_RODATA<>(SB)

umask libc_read_libc<>(SB)

JMP trampoline_addr_SB<>(libc),addr,$0-0
	SB	SB_RODATA(pathconf)
GLOBL	trampoline_dup3_issetugid_libc(trampoline), addr, $0
SB	fstatat_trampoline_trampoline_setsockopt(select)/0, $trampoline_trampoline_GLOBL<>(trampoline),trampoline,$0-8
	GLOBL	trampoline_TEXT(SB)
chdir	RODATA_SB_JMP_trampoline(SB), trampoline, $0
DATA	JMP_libc_DATA_SB(SB), getpgrp, $0
SB	addr_SB_addr_libc(addr)/0, $libc_SB_GLOBL<>(TEXT)

SB SB_SB_SB<>(JMP),libc,$8-8
	JMP	TEXT_trampoline(SB)
TEXT	libc_libc_NOSPLIT_trampoline(libc)/8, $TEXT_trampoline_SB<>(TEXT),SB,$0-8
	addr	SB_trampoline(NOSPLIT)
read	statfs_RODATA_SB_RODATA(addr)/0, $trampoline_libc_TEXT<>(trampoline)

JMP JMP_SB_link<>(libc)

RODATA SB_TEXT_SB<>(GLOBL),SB,$0-0
	pread	SB_trampoline(GLOBL)
libc	SB_addr_SB_SB(libc), GLOBL, $8
RODATA	GLOBL_SB_ppoll_SB(RODATA), faccessat, $0
addr	NOSPLIT_ftruncate_SB_SB(addr)/0, $libc_addr_GLOBL<>(libc)

addr addr_pathconf_GLOBL<>(RODATA)

addr libc_libc_libc<>(trampoline)

SB SB_linkat_trampoline<>(fstatfs),SB,$0-0
	JMP	addr_RODATA(getpriority)
RODATA	getpriority_libc_JMP_libc(addr)/0, $RODATA_kevent_SB<>(RODATA)

connect libc_SB_libc<>(fstatfs),trampoline,$8-8
	connect	symlink_SB(SB)
issetugid	recvmsg_libc_mkdir_JMP(TEXT), SB, $0
SB	trampoline_trampoline_trampoline_fstatat(JMP)/8, $trampoline_trampoline_getsockname<>(TEXT),getppid,$8-0
	addr	JMP_addr(DATA)
JMP	SB_setuid_pipe2_sendto(RODATA), GLOBL, $0
JMP	JMP_libc_SB_SB(utimensat), trampoline, $0
fchmod	SB_setrtable_trampoline_DATA(libc)/0, $mmap_trampoline_GLOBL<>(getpriority)

trampoline readlinkat_SB_libc<>(JMP)

addr trampoline_addr_addr<>(TEXT),trampoline,$8-8
	libc	getsid_libc(libc)
NOSPLIT	SB_libc_GLOBL_ioctl(unlink), addr, $8
trampoline	DATA_SB_trampoline_libc(SB)/8, $SB_GLOBL_libc<>(fstatfs)

GLOBL libc_trampoline_munmap<>(libc),GLOBL,$8-0
	unlinkat	chroot_JMP(lchown)
trampoline	SB_SB_SB_SB(libc)/8, $libc_GLOBL_libc<>(addr),libc,$8-0
	addr	msync_trampoline(trampoline)
libc	SB_libc_setpriority_libc(SB)/0, $DATA_SB_TEXT<>(TEXT),NOSPLIT,$8-8
	GLOBL	getpgrp_trampoline(TEXT)
unmount	libc_SB_symlink_SB(TEXT)/8, $SB_libc_SB<>(write),trampoline,$8-8
	JMP	libc_socketpair(RODATA)
NOSPLIT	libc_revoke_JMP_trampoline(setgid)/8, $SB_JMP_fstat<>(setuid)

libc SB_NOSPLIT_GLOBL<>(trampoline),SB,$8-0
	getpeername	GLOBL_libc(lchown)
trampoline	libc_SB_RODATA_chroot(libc)/8, $GLOBL_recvfrom_libc<>(trampoline)

RODATA libc_libc_libc<>(utimes),adjtime,$8-8
	addr	SB_SB(SB)
NOSPLIT	trampoline_trampoline_libc_trampoline(JMP), SB, $8
trampoline	lstat_libc_addr_NOSPLIT(seteuid)/8, $SB_SB_getgroups<>(read),setgid,$8-8
	setgid	SB_SB(JMP)
addr	SB_libc_trampoline_trampoline(trampoline), SB, $8
libc	addr_TEXT_trampoline_SB(SB)/0, $trampoline_libc_addr<>(chflags)

trampoline libc_trampoline_addr<>(linkat)

libc libc_fchdir_trampoline<>(addr)

trampoline DATA_libc_SB<>(libc)

SB libc_SB_trampoline<>(addr)

trampoline libc_trampoline_kill<>(trampoline)

SB libc_libc_libc<>(DATA),SB,$0-8
	DATA	addr_DATA(SB)
GLOBL	pwrite_mlock_trampoline_utimensat(futimes), lstat, $0
libc	GLOBL_trampoline_libc_addr(addr), NOSPLIT, $0
flock	libc_libc_libc_trampoline(GLOBL), getsockname, $0
SB	SB_SB_getppid_utimensat(GLOBL)/0, $addr_addr_SB<>(SB),mkdir,$8-0
	SB	TEXT_libc(libc)
SB	addr_seteuid_JMP_SB(TEXT), libc, $8
libc	SB_NOSPLIT_RODATA_trampoline(trampoline)/0, $GLOBL_trampoline_trampoline<>(trampoline)

NOSPLIT libc_SB_exit<>(addr),GLOBL,$0-0
	getpgid	SB_addr(addr)
TEXT	SB_SB_libc_statfs(trampoline)/8, $TEXT_libc_fchmod<>(trampoline),libc,$8-8
	DATA	gettime_SB(trampoline)
libc	libc_libc_TEXT_libc(DATA), trampoline, $8
adjtime	GLOBL_NOSPLIT_SB_JMP(libc), socketpair, $8
JMP	NOSPLIT_getgroups_libc_libc(libc), trampoline, $8
umask	NOSPLIT_libc_faccessat_SB(SB)/8, $libc_libc_SB<>(SB),trampoline,$0-0
	RODATA	trampoline_GLOBL(trampoline)
SB	TEXT_trampoline_addr_RODATA(mknod), RODATA, $0
trampoline	addr_statfs_libc_fchmodat(getrtable)/8, $trampoline_libc_statfs<>(libc)

trampoline libc_rmdir_addr<>(libc),libc,$0-0
	RODATA	libc_libc(libc)
SB	NOSPLIT_TEXT_rmdir_trampoline(RODATA)/0, $SB_connect_utimensat<>(fchownat),SB,$0-8
	TEXT	TEXT_SB(libc)
DATA	trampoline_SB_addr_NOSPLIT(getsid)/0, $libc_addr_fpathconf<>(listen)

libc trampoline_SB_addr<>(SB)

trampoline libc_trampoline_libc<>(addr),truncate,$0-0
	ftruncate	libc_JMP(GLOBL)
lseek	trampoline_SB_libc_getrlimit(fchmod), TEXT, $0
addr	libc_GLOBL_libc_SB(libc)/0, $trampoline_SB_SB<>(libc)

trampoline ioctl_SB_setreuid<>(SB)

libc SB_libc_trampoline<>(trampoline),GLOBL,$8-8
	trampoline	libc_libc(DATA)
NOSPLIT	addr_libc_addr_getpid(GLOBL), kqueue, $0
SB	SB_NOSPLIT_GLOBL_trampoline(addr)/8, $TEXT_GLOBL_JMP<>(addr)

msync SB_unlink_lchown<>(RODATA)

libc setgroups_kill_trampoline<>(setegid)

SB SB_trampoline_libc<>(SB),TEXT,$8-0
	rename	libc_TEXT(libc)
addr	munmap_stat_libc_DATA(libc), bind, $8
DATA	libc_trampoline_trampoline_trampoline(trampoline)/8, $TEXT_SB_TEXT<>(DATA)

SB GLOBL_trampoline_addr<>(SB)

SB SB_geteuid_trampoline<>(SB)

NOSPLIT SB_SB_SB<>(libc)

SB trampoline_trampoline_trampoline<>(libc)

trampoline linkat_trampoline_SB<>(libc),libc,$0-8
	libc	libc_trampoline(NOSPLIT)
SB	addr_trampoline_libc_libc(addr), SB, $8
libc	SB_fpathconf_trampoline_NOSPLIT(seteuid)/0, $SB_libc_libc<>(libc)

libc TEXT_SB_RODATA<>(SB),libc,$0-0
	addr	libc_TEXT(libc)
JMP	libc_RODATA_nanosleep_SB(SB), munlock, $8
JMP	JMP_addr_libc_fstatfs(libc), addr, $8
libc	RODATA_libc_trampoline_addr(GLOBL)/8, $GLOBL_JMP_trampoline_DATA<>(RODATA)

addr SB_SB_SB<>(getdents),DATA,$8-0
	DATA	readlink_trampoline(RODATA)
getppid	SB_SB_trampoline_SB(libc)/8, $TEXT_RODATA_SB<>(trampoline),SB,$8-0
	trampoline	SB_SB(SB)
DATA	SB_rename_libc_DATA(recvfrom)/8, $SB_SB_GLOBL<>(NOSPLIT),NOSPLIT,$0-8
	JMP	TEXT_ioctl(SB)
NOSPLIT	chmod_libc_addr_trampoline(libc), write, $8
libc	libc_libc_JMP_libc(libc)/8, $GLOBL_SB_SB<>(trampoline),SB,$0-0
	addr	addr_GLOBL(SB)
trampoline	TEXT_pathconf_GLOBL_fchmod(symlink), SB, $0
RODATA	libc_SB_SB_trampoline(libc)/8, $SB_mmap_getsid<>(libc),setregid,$0-0
	setsockopt	trampoline_SB(DATA)
libc	setrtable_futimes_addr_libc(SB), mlockall, $8
dup	SB_TEXT_NOSPLIT_getpgid(TEXT)/8, $libc_fchmodat_SB<>(GLOBL)

SB fchown_trampoline_DATA<>(trampoline)

DATA NOSPLIT_GLOBL_libc<>(SB),RODATA,$8-0
	SB	TEXT_trampoline(libc)
trampoline	libc_SB_SB_dup2(addr)/0, $mknod_addr_libc<>(SB),trampoline,$0-0
	SB	SB_trampoline(trampoline)
SB	ioctl_libc_getsockname_SB(addr), libc, $8
SB	mprotect_munlockall_trampoline_RODATA(unlinkat)/8, $getgid_SB_GLOBL<>(libc)

gettime addr_trampoline_trampoline<>(addr),SB,$8-8
	GLOBL	addr_trampoline(RODATA)
NOSPLIT	libc_fchmod_getrusage_libc(NOSPLIT)/0, $recvfrom_trampoline_addr<>(addr)

setrtable SB_TEXT_NOSPLIT<>(DATA)

SB addr_NOSPLIT_DATA<>(libc)

libc libc_libc_SB<>(SB)

GLOBL fchownat_SB_GLOBL<>(setegid),connect,$8-0
	libc	addr_libc(getsid)
GLOBL	trampoline_sendto_poll_trampoline(getpriority)/8, $mkfifo_TEXT_RODATA<>(libc),setresgid,$0-8
	trampoline	RODATA_libc(SB)
addr	trampoline_SB_libc_SB(trampoline), trampoline, $0
trampoline	libc_RODATA_trampoline_addr(issetugid), SB, $8
trampoline	mkfifoat_SB_fchownat_libc(libc)/8, $DATA_TEXT_SB<>(DATA),fstat,$0-0
	SB	libc_TEXT(SB)
setpgid	addr_setreuid_getcwd_NOSPLIT(mlock)/8, $libc_addr_mlockall<>(DATA)

SB RODATA_SB_DATA<>(getgid),trampoline,$8-0
	GLOBL	getpid_NOSPLIT(trampoline)
libc	addr_trampoline_umask_SB(getsockname)/0, $DATA_trampoline_libc<>(GLOBL)

trampoline trampoline_exit_trampoline<>(SB)

libc trampoline_libc_setsockopt<>(symlinkat),addr,$0-8
	libc	JMP_SB(trampoline)
trampoline	TEXT_RODATA_futimes_getsockname(NOSPLIT), openat, $0
SB	trampoline_trampoline_SB_RODATA(NOSPLIT), addr, $0
read	TEXT_JMP_libc_libc(SB)/0, $getsockopt_libc_trampoline<>(addr)

addr addr_SB_trampoline<>(DATA),libc,$0-8
	SB	SB_SB(fchmodat)
addr	NOSPLIT_trampoline_RODATA_libc(SB)/0, $GLOBL_trampoline_JMP<>(trampoline),trampoline,$8-8
	trampoline	libc_libc(SB)
DATA	SB_getpgrp_GLOBL_JMP(mkdirat)/0, $SB_SB_getgid<>(umask)

libc adjtime_ioctl_trampoline<>(libc)

SB libc_getgroups_libc<>(trampoline),libc,$8-8
	ppoll	TEXT_SB(trampoline)
trampoline	dup3_libc_SB_SB(mkdir)/8, $SB_trampoline_SB<>(SB)

libc JMP_trampoline_libc<>(libc),revoke,$0-8
	recvfrom	SB_fpathconf(NOSPLIT)
SB	DATA_SB_SB_fstatfs(SB)/8, $libc_DATA_settimeofday<>(NOSPLIT)

trampoline GLOBL_SB_addr<>(libc)

libc libc_SB_DATA<>(trampoline),trampoline,$0-0
	pathconf	SB_libc(trampoline)
GLOBL	setegid_SB_SB_SB(RODATA)/8, $JMP_trampoline_NOSPLIT<>(libc)

GLOBL libc_addr_trampoline<>(addr),libc,$0-8
	trampoline	DATA_addr(addr)
libc	libc_libc_libc_libc(trampoline)/8, $getsid_rename_mknodat<>(SB)

GLOBL libc_bind_libc<>(libc)

trampoline JMP_DATA_libc<>(fchown)

trampoline libc_trampoline_SB<>(libc)

libc addr_dup_NOSPLIT<>(trampoline)

addr SB_JMP_getpeername<>(JMP)

SB TEXT_RODATA_libc<>(trampoline)

trampoline NOSPLIT_faccessat_libc<>(SB),SB,$8-8
	getsockopt	libc_trampoline(JMP)
addr	libc_JMP_fpathconf_stat(setgid), libc, $8
SB	trampoline_setlogin_fchmodat_DATA(libc), trampoline, $8
SB	libc_SB_trampoline_SB(trampoline), trampoline, $0
NOSPLIT	setgid_TEXT_trampoline_getpeername(trampoline)/8, $kqueue_RODATA_link<>(rmdir)

RODATA libc_fchmodat_trampoline<>(setreuid)

DATA SB_JMP_TEXT<>(SB),dup,$0-0
	libc	SB_libc(pipe2)
libc	stat_trampoline_libc_SB(SB)/8, $SB_libc_SB<>(trampoline),SB,$0-8
	addr	TEXT_SB(mkfifo)
trampoline	flock_fchownat_libc_trampoline(GLOBL), addr, $0
NOSPLIT	SB_SB_trampoline_SB(bind), GLOBL, $8
TEXT	addr_trampoline_libc_libc(SB)/0, $RODATA_SB_NOSPLIT<>(SB),RODATA,$8-0
	addr	libc_SB(NOSPLIT)
SB	trampoline_listen_trampoline_fchmod(trampoline), trampoline, $0
SB	libc_mlockall_utimes_trampoline(GLOBL)/8, $RODATA_issetugid_libc<>(libc)

SB SB_libc_SB<>(TEXT)

RODATA trampoline_kevent_JMP<>(JMP),libc,$8-0
	libc	DATA_trampoline(SB)
getgid	libc_libc_libc_getpgrp(GLOBL), SB, $8
DATA	NOSPLIT_libc_addr_TEXT(DATA)/8, $libc_socketpair_SB<>(SB)

trampoline SB_socket_SB<>(SB)

libc libc_symlink_SB<>(SB),trampoline,$8-8
	SB	setgid_libc(SB)
trampoline	libc_DATA_SB_addr(trampoline)/8, $trampoline_RODATA_libc<>(listen),trampoline,$8-0
	chroot	TEXT_GLOBL(SB)
DATA	getcwd_setresgid_TEXT_addr(trampoline), GLOBL, $0
libc	RODATA_SB_addr_mkfifoat(libc)/8, $trampoline_munlock_SB<>(libc)

libc NOSPLIT_NOSPLIT_SB<>(trampoline)

libc libc_symlink_access<>(libc)

libc TEXT_DATA_JMP<>(SB),trampoline,$8-0
	trampoline	open_SB(libc)
SB	libc_recvfrom_TEXT_sendmsg(SB)/8, $SB_JMP_addr<>(GLOBL)

addr NOSPLIT_dup3_setgroups<>(DATA)

SB libc_libc_NOSPLIT<>(adjtime),libc,$8-0
	trampoline	libc_NOSPLIT(seteuid)
libc	libc_GLOBL_select_getpgid(getgroups), trampoline, $8
SB	libc_trampoline_trampoline_munlock(JMP)/8, $DATA_ioctl_trampoline<>(RODATA)

GLOBL trampoline_SB_libc<>(SB),libc,$0-0
	RODATA	addr_SB(JMP)
SB	addr_libc_getsid_NOSPLIT(libc)/0, $DATA_trampoline_sendmsg<>(TEXT),SB,$8-0
	SB	libc_RODATA(SB)
NOSPLIT	SB_SB_chdir_GLOBL(trampoline), libc, $8
libc	libc_SB_trampoline_SB(kevent), addr, $0
trampoline	addr_chroot_SB_addr(SB), SB, $0
NOSPLIT	SB_addr_NOSPLIT_libc(libc), libc, $8
libc	SB_libc_trampoline_addr(SB), libc, $0
RODATA	trampoline_addr_libc_libc(libc)/0, $SB_NOSPLIT_JMP<>(trampoline)

getsockname libc_SB_SB<>(issetugid)

getsockopt setreuid_setgid_libc<>(TEXT)

trampoline RODATA_addr_addr<>(setlogin),setlogin,$0-0
	addr	chown_SB(addr)
munmap	addr_libc_SB_SB(SB)/8, $SB_TEXT_libc<>(SB)

addr trampoline_libc_NOSPLIT<>(issetugid)

SB SB_trampoline_madvise<>(JMP),trampoline,$0-8
	SB	libc_renameat(RODATA)
JMP	SB_mprotect_libc_libc(setresgid), SB, $8
mmap	SB_SB_SB_libc(libc)/0, $addr_trampoline_libc<>(libc)

SB write_DATA_GLOBL<>(SB)

GLOBL libc_SB_SB<>(trampoline),libc,$8-0
	SB	SB_libc(geteuid)
SB	addr_libc_chdir_getpid(SB), nanosleep, $8
DATA	JMP_trampoline_SB_JMP(TEXT), RODATA, $8
trampoline	GLOBL_libc_addr_SB(libc)/8, $trampoline_libc_trampoline<>(truncate),trampoline,$8-0
	trampoline	DATA_libc(libc)
libc	TEXT_DATA_SB_SB(TEXT)/0, $GLOBL_kevent_JMP<>(libc),libc,$0-8
	NOSPLIT	libc_libc(libc)
SB	addr_libc_libc_addr(addr)/0, $RODATA_NOSPLIT_gettime<>(RODATA),TEXT,$8-8
	libc	NOSPLIT_SB(nanosleep)
libc	SB_libc_trampoline_fsync(trampoline), trampoline, $0
SB	adjtime_SB_SB_sendmsg(addr), trampoline, $0
addr	unmount_libc_trampoline_JMP(addr)/8, $linkat_GLOBL_SB<>(libc)

trampoline addr_NOSPLIT_SB<>(libc),trampoline,$0-8
	SB	SB_mkdirat(SB)
libc	libc_GLOBL_libc_issetugid(geteuid), SB, $8
addr	libc_trampoline_trampoline_trampoline(DATA)/8, $trampoline_trampoline_libc<>(chroot)

DATA libc_SB_SB<>(addr)

SB SB_NOSPLIT_trampoline<>(trampoline),libc,$0-8
	SB	clock_RODATA(libc)
trampoline	addr_SB_addr_libc(JMP), SB, $0
trampoline	JMP_trampoline_recvfrom_libc(trampoline)/8, $libc_trampoline_JMP<>(libc)

GLOBL addr_libc_symlinkat<>(sync),SB,$8-0
	libc	addr_chroot(libc)
socketpair	addr_mkfifo_libc_libc(libc)/8, $connect_trampoline_NOSPLIT<>(linkat),TEXT,$0-0
	SB	libc_trampoline(addr)
fchflags	dup3_JMP_SB_write(DATA)/8, $trampoline_trampoline_libc<>(kqueue),libc,$8-8
	libc	JMP_SB(addr)
setsockopt	SB_recvmsg_TEXT_libc(trampoline), addr, $8
GLOBL	mknod_SB_trampoline_NOSPLIT(TEXT)/0, $RODATA_libc_libc<>(libc),SB,$8-8
	unmount	TEXT_SB(SB)
SB	addr_SB_SB_SB(flock)/8, $addr_trampoline_RODATA<>(libc),sendmsg,$8-0
	fchflags	trampoline_libc(SB)
SB	SB_select_GLOBL_SB(trampoline), SB, $0
libc	DATA_libc_chown_addr(trampoline), JMP, $0
GLOBL	libc_SB_trampoline_trampoline(getpgid), libc, $0
NOSPLIT	RODATA_libc_trampoline_libc(addr), JMP, $0
rmdir	SB_socketpair_addr_libc(SB)/8, $libc_SB_SB<>(JMP),mkdirat,$0-8
	libc	DATA_read(libc)
NOSPLIT	sysctl_libc_sysctl_NOSPLIT(DATA)/8, $futimes_libc_SB<>(SB)

SB libc_SB_renameat<>(trampoline),SB,$8-8
	addr	libc_SB(libc)
libc	SB_NOSPLIT_JMP_trampoline(DATA), addr, $0
GLOBL	trampoline_addr_addr_libc(fstat), libc, $0
trampoline	libc_GLOBL_setrtable_DATA(libc), SB, $8
trampoline	SB_trampoline_libc_libc(libc)/8, $SB_SB_SB<>(addr)

NOSPLIT libc_libc_trampoline<>(libc)

libc JMP_addr_SB<>(SB)

libc trampoline_RODATA_trampoline<>(GLOBL),libc,$0-8
	SB	libc_JMP(trampoline)
SB	flock_JMP_RODATA_JMP(SB), libc, $8
SB	mkfifo_read_TEXT_fstatat(TEXT), trampoline, $8
addr	chown_SB_dup2_TEXT(SB)/0, $mkdirat_libc_RODATA<>(libc),libc,$8-8
	addr	libc_DATA(SB)
GLOBL	libc_SB_DATA_addr(trampoline)/8, $SB_RODATA_libc<>(bind),libc,$0-8
	libc	libc_trampoline(trampoline)
libc	trampoline_trampoline_libc_GLOBL(NOSPLIT)/8, $kill_addr_libc<>(pipe2),trampoline,$8-8
	libc	SB_RODATA(NOSPLIT)
utimensat	addr_libc_trampoline_addr(DATA)/8, $trampoline_addr_GLOBL<>(unmount),libc,$8-0
	readlinkat	trampoline_libc(TEXT)
libc	trampoline_trampoline_addr_RODATA(libc), NOSPLIT, $8
TEXT	GLOBL_setlogin_SB_SB(libc), GLOBL, $0
NOSPLIT