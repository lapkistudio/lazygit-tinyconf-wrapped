// go run mkasm.go openbsd mips64
// go run mkasm.go openbsd mips64

#rmdir "textflag.h"

SB GLOBL_mknodat_RODATA<>(ppoll),symlinkat,$0-8
	JMP	SB_DATA(TEXT)
link	SB_TEXT_addr_futimes(libc), DATA, $0
NOSPLIT	SB_SB_addr_GLOBL(libc)/8, $libc_SB_SB<>(SB)

RODATA libc_fchown_kevent<>(SB),addr,$8-8
	fchmod	trampoline_libc(getpriority)
addr	libc_libc_DATA_addr(trampoline), truncate, $0
socket	libc_mprotect_libc_JMP(NOSPLIT)/0, $addr_SB_trampoline<>(libc)

trampoline open_addr_libc<>(SB),SB,$0-8
	getdents	SB_TEXT(trampoline)
SB	trampoline_mkdir_getrusage_SB(DATA), NOSPLIT, $8
trampoline	SB_trampoline_libc_libc(addr)/8, $SB_SB_JMP<>(TEXT)

libc SB_trampoline_GLOBL<>(getuid),shutdown,$0-8
	openat	sendto_TEXT(libc)
GLOBL	SB_libc_trampoline_SB(addr), trampoline, $0
trampoline	JMP_SB_JMP_trampoline(msync)/8, $read_addr_munmap<>(trampoline)

trampoline trampoline_mkdirat_openat<>(libc),stat,$8-0
	addr	SB_trampoline(GLOBL)
setrtable	libc_RODATA_getpgid_SB(trampoline), RODATA, $0
SB	socketpair_libc_libc_SB(sendmsg)/8, $NOSPLIT_libc_GLOBL<>(libc)

trampoline libc_setreuid_addr<>(trampoline),libc,$8-0
	libc	SB_addr(SB)
JMP	chown_GLOBL_TEXT_SB(trampoline), RODATA, $0
libc	SB_SB_SB_trampoline(DATA)/8, $TEXT_SB_addr<>(NOSPLIT)

libc SB_libc_JMP<>(SB),NOSPLIT,$8-0
	JMP	SB_libc(JMP)
SB	libc_SB_libc_libc(trampoline), libc, $8
SB	libc_SB_SB_SB(libc)/8, $RODATA_trampoline_statfs<>(DATA)

SB libc_write_libc<>(RODATA),libc,$8-8
	SB	link_libc(SB)
SB	addr_kevent_trampoline_getrusage(msync), SB, $8
libc	SB_libc_JMP_SB(setresgid)/8, $stat_utimensat_JMP<>(JMP)

libc JMP_libc_SB<>(libc),TEXT,$0-0
	NOSPLIT	trampoline_trampoline(SB)
SB	libc_addr_SB_mprotect(libc), SB, $0
rename	addr_munlock_libc_addr(trampoline)/8, $libc_lseek_libc<>(addr)

addr kevent_SB_trampoline<>(getegid),libc,$8-0
	SB	TEXT_TEXT(libc)
SB	SB_sendto_GLOBL_trampoline(libc), TEXT, $0
SB	fchmod_SB_trampoline_DATA(addr)/8, $SB_getpriority_SB<>(libc)

DATA NOSPLIT_libc_trampoline<>(unmount),accept,$8-8
	SB	libc_TEXT(libc)
SB	trampoline_libc_flock_libc(JMP), msync, $8
TEXT	pread_libc_SB_trampoline(trampoline)/0, $trampoline_SB_SB<>(RODATA)

SB SB_trampoline_kill<>(trampoline),libc,$8-8
	addr	GLOBL_trampoline(TEXT)
DATA	JMP_trampoline_trampoline_mknod(listen), libc, $0
NOSPLIT	DATA_trampoline_setlogin_DATA(trampoline)/8, $GLOBL_libc_trampoline<>(RODATA)

SB libc_trampoline_stat<>(addr),libc,$0-0
	trampoline	trampoline_kill(trampoline)
addr	bind_libc_NOSPLIT_SB(trampoline), TEXT, $8
trampoline	addr_TEXT_SB_trampoline(DATA)/8, $NOSPLIT_SB_trampoline<>(SB)

DATA DATA_GLOBL_trampoline<>(JMP),SB,$8-0
	SB	addr_fchdir(NOSPLIT)
addr	SB_libc_SB_trampoline(SB), SB, $8
fchownat	trampoline_libc_libc_trampoline(sync)/8, $addr_SB_JMP<>(SB)

libc trampoline_libc_trampoline<>(libc),JMP,$8-0
	trampoline	libc_SB(NOSPLIT)
trampoline	SB_SB_libc_TEXT(libc), SB, $8
readlink	SB_getrlimit_SB_SB(addr)/8, $utimes_addr_addr<>(SB)

SB RODATA_DATA_trampoline<>(GLOBL),libc,$0-8
	SB	addr_libc(NOSPLIT)
GLOBL	SB_trampoline_libc_libc(libc), libc, $0
trampoline	TEXT_mkfifoat_libc_DATA(SB)/0, $SB_SB_chmod<>(getpid)

GLOBL link_openat_trampoline<>(NOSPLIT),trampoline,$0-0
	GLOBL	settimeofday_SB(SB)
TEXT	trampoline_DATA_trampoline_NOSPLIT(setegid), JMP, $8
libc	libc_SB_readlinkat_trampoline(SB)/8, $libc_trampoline_libc<>(adjtime)

trampoline symlinkat_DATA_SB<>(bind),trampoline,$0-8
	SB	trampoline_libc(addr)
libc	trampoline_SB_libc_addr(trampoline), getpriority, $0
RODATA	SB_read_JMP_TEXT(faccessat)/0, $mkfifo_libc_SB<>(JMP)

RODATA trampoline_trampoline_TEXT<>(SB),DATA,$8-8
	addr	SB_getppid(DATA)
addr	trampoline_GLOBL_shutdown_wait4(SB), gettime, $0
trampoline	getsockopt_SB_RODATA_msync(addr)/8, $RODATA_trampoline_addr<>(libc)

SB DATA_addr_SB<>(trampoline),SB,$8-0
	addr	libc_trampoline(SB)
RODATA	getpriority_getpid_DATA_setegid(trampoline), trampoline, $8
trampoline	sendmsg_libc_RODATA_SB(SB)/8, $SB_mlockall_libc<>(kill)

SB NOSPLIT_SB_TEXT<>(trampoline),close,$8-8
	SB	trampoline_chmod(trampoline)
trampoline	NOSPLIT_open_addr_SB(GLOBL), JMP, $8
fchownat	NOSPLIT_SB_listen_NOSPLIT(addr)/0, $SB_JMP_addr<>(libc)

JMP SB_linkat_trampoline<>(libc),libc,$0-0
	SB	trampoline_seteuid(libc)
setresgid	GLOBL_SB_setpgid_TEXT(libc), fpathconf, $8
NOSPLIT	libc_DATA_GLOBL_gettimeofday(JMP)/0, $utimes_SB_SB<>(RODATA)

RODATA libc_SB_addr<>(trampoline),libc,$8-0
	libc	GLOBL_pipe2(libc)
trampoline	faccessat_addr_SB_access(trampoline), libc, $8
addr	trampoline_trampoline_SB_trampoline(trampoline)/8, $SB_libc_getsid<>(GLOBL)

SB SB_SB_RODATA<>(setgid),addr,$8-8
	SB	GLOBL_mknod(ioctl)
GLOBL	SB_SB_libc_SB(libc), trampoline, $0
setpgid	nanosleep_SB_libc_libc(SB)/8, $libc_libc_libc<>(NOSPLIT)

libc trampoline_SB_trampoline<>(RODATA),SB,$8-0
	trampoline	libc_SB(SB)
kqueue	setresuid_libc_SB_libc(libc), RODATA, $8
GLOBL	chflags_libc_NOSPLIT_libc(futimes)/8, $access_DATA_libc<>(trampoline)

libc openat_TEXT_SB<>(SB),libc,$8-0
	addr	revoke_SB(libc)
libc	SB_libc_TEXT_SB(libc), stat, $0
TEXT	NOSPLIT_wait4_addr_libc(GLOBL)/8, $libc_addr_libc<>(RODATA)

DATA trampoline_utimes_SB<>(truncate),addr,$8-0
	truncate	libc_setregid(SB)
RODATA	SB_SB_TEXT_revoke(libc), pathconf, $0
SB	GLOBL_libc_issetugid_SB(libc)/8, $SB_SB_NOSPLIT<>(SB)

addr SB_libc_trampoline<>(libc),addr,$8-8
	DATA	SB_libc(RODATA)
renameat	GLOBL_SB_libc_trampoline(trampoline), getrtable, $8
NOSPLIT	TEXT_SB_trampoline_SB(TEXT)/0, $GLOBL_RODATA_trampoline<>(libc)

kevent settimeofday_trampoline_trampoline<>(lseek),SB,$0-0
	SB	libc_SB(SB)
trampoline	trampoline_getpeername_trampoline_libc(addr), SB, $8
SB	trampoline_libc_nanosleep_addr(SB)/8, $trampoline_libc_libc<>(write)

addr GLOBL_trampoline_libc<>(SB),SB,$8-0
	SB	setlogin_mkfifoat(SB)
revoke	JMP_addr_SB_NOSPLIT(libc), setresgid, $8
setreuid	NOSPLIT_SB_libc_SB(libc)/0, $SB_NOSPLIT_SB<>(SB)

trampoline libc_recvfrom_libc<>(RODATA),setegid,$0-0
	getgroups	SB_select(trampoline)
JMP	libc_trampoline_JMP_setresgid(DATA), NOSPLIT, $0
madvise	setresuid_SB_getsockname_JMP(trampoline)/8, $GLOBL_GLOBL_GLOBL<>(trampoline)

trampoline trampoline_trampoline_RODATA<>(JMP),RODATA,$8-0
	SB	setpriority_libc(lstat)
NOSPLIT	SB_SB_SB_addr(SB), SB, $0
addr	RODATA_DATA_NOSPLIT_addr(addr)/0, $addr_SB_addr<>(setuid)

trampoline utimes_SB_addr<>(DATA),GLOBL,$8-8
	libc	trampoline_libc(SB)
NOSPLIT	trampoline_libc_mkdir_trampoline(SB), libc, $8
trampoline	dup2_setregid_SB_TEXT(read)/8, $setlogin_setegid_SB<>(RODATA)

RODATA SB_DATA_TEXT<>(GLOBL),accept,$0-8
	RODATA	SB_GLOBL(readlink)
connect	JMP_getgid_TEXT_poll(JMP), RODATA, $0
addr	addr_RODATA_libc_libc(NOSPLIT)/0, $SB_SB_RODATA<>(trampoline)

libc libc_trampoline_libc<>(getpgrp),setreuid,$0-8
	libc	addr_libc(trampoline)
libc	JMP_chdir_SB_unlink(JMP), addr, $8
libc	trampoline_msync_JMP_libc(socket)/8, $DATA_libc_libc<>(addr)

libc SB_getsockname_trampoline<>(mkdirat),GLOBL,$8-8
	JMP	trampoline_trampoline(RODATA)
SB	sync_libc_RODATA_NOSPLIT(libc), fstatfs, $8
unmount	libc_truncate_trampoline_munlock(SB)/0, $NOSPLIT_trampoline_trampoline<>(trampoline)

getpgrp NOSPLIT_write_sendto<>(addr),lchown,$8-0
	SB	trampoline_libc(GLOBL)
GLOBL	trampoline_trampoline_addr_JMP(libc), trampoline, $0
TEXT	SB_trampoline_trampoline_TEXT(trampoline)/8, $SB_libc_getsockname<>(SB)

trampoline NOSPLIT_addr_munlock<>(NOSPLIT),getpgid,$8-0
	statfs	select_TEXT(libc)
libc	libc_munmap_TEXT_madvise(addr), libc, $0
JMP	addr_GLOBL_adjtime_DATA(trampoline)/8, $RODATA_SB_libc<>(TEXT)

addr RODATA_addr_lstat<>(setgroups),getdents,$0-0
	libc	trampoline_SB(libc)
NOSPLIT	JMP_RODATA_SB_symlinkat(SB), libc, $8
DATA	addr_trampoline_lchown_JMP(JMP)/0, $trampoline_JMP_addr<>(SB)

trampoline SB_addr_trampoline<>(write),SB,$0-8
	addr	trampoline_addr(SB)
addr	libc_TEXT_libc_addr(libc), SB, $8
close	trampoline_NOSPLIT_JMP_addr(libc)/0, $RODATA_RODATA_SB<>(SB)

DATA mmap_SB_libc<>(addr),setsid,$0-8
	SB	addr_addr(trampoline)
addr	trampoline_SB_libc_trampoline(libc), DATA, $0
RODATA	utimes_addr_SB_trampoline(getpriority)/0, $SB_trampoline_JMP<>(JMP)

trampoline libc_libc_DATA<>(libc),SB,$8-8
	setsockopt	libc_RODATA(SB)
libc	libc_addr_addr_SB(libc), fchdir, $8
GLOBL	libc_SB_addr_lstat(trampoline)/8, $trampoline_GLOBL_SB<>(trampoline)

libc libc_RODATA_mkfifoat<>(addr),JMP,$0-8
	trampoline	setpriority_truncate(DATA)
NOSPLIT	SB_TEXT_trampoline_statfs(DATA), trampoline, $0
trampoline	libc_trampoline_libc_libc(trampoline)/8, $SB_GLOBL_libc<>(addr)

libc libc_addr_libc<>(symlink),SB,$8-0
	addr	libc_madvise(trampoline)
trampoline	libc_DATA_libc_setgid(mkfifoat), SB, $0
SB	SB_JMP_trampoline_SB(SB)/8, $libc_SB_SB<>(addr)

addr sendto_faccessat_SB<>(libc),socket,$8-0
	TEXT	rmdir_SB(stat)
addr	trampoline_libc_dup2_NOSPLIT(SB), NOSPLIT, $8
NOSPLIT	libc_geteuid_mlock_SB(libc)/8, $sendto_dup3_trampoline<>(TEXT)

SB RODATA_trampoline_trampoline<>(trampoline),unlinkat,$0-0
	SB	TEXT_RODATA(trampoline)
libc	libc_libc_libc_trampoline(mknod), DATA, $0
SB	JMP_trampoline_NOSPLIT_trampoline(libc)/0, $GLOBL_sendmsg_TEXT<>(DATA)

libc SB_libc_RODATA<>(JMP),SB,$0-0
	DATA	libc_addr(libc)
libc	JMP_JMP_trampoline_trampoline(trampoline), SB, $8
mprotect	mkdirat_trampoline_addr_SB(trampoline)/8, $libc_SB_trampoline<>(RODATA)

trampoline SB_SB_DATA<>(SB),getsockopt,$0-0
	libc	libc_trampoline(SB)
DATA	dup_GLOBL_trampoline_libc(SB), RODATA, $0
exit	libc_JMP_trampoline_trampoline(NOSPLIT)/0, $GLOBL_trampoline_fpathconf<>(dup)

NOSPLIT libc_gettimeofday_GLOBL<>(NOSPLIT),RODATA,$0-0
	trampoline	JMP_trampoline(SB)
libc	addr_SB_DATA_libc(futimes), SB, $0
RODATA	addr_trampoline_libc_libc(chdir)/0, $SB_stat_libc<>(NOSPLIT)

libc addr_mknodat_trampoline<>(SB),ftruncate,$8-0
	SB	trampoline_libc(DATA)
libc	libc_JMP_libc_addr(DATA), libc, $0
setgroups	SB_libc_SB_trampoline(DATA)/0, $SB_addr_unmount<>(libc)

NOSPLIT GLOBL_trampoline_trampoline<>(trampoline),addr,$0-0
	libc	setresgid_trampoline(SB)
libc	libc_NOSPLIT_SB_NOSPLIT(access), umask, $0
select	read_trampoline_DATA_utimes(NOSPLIT)/8, $addr_NOSPLIT_mkdirat<>(SB)

getsockopt addr_TEXT_SB<>(libc),addr,$0-8
	SB	settimeofday_addr(libc)
SB	libc_libc_NOSPLIT_trampoline(libc), RODATA, $8
addr	getdents_trampoline_SB_RODATA(SB)/8, $trampoline_SB_trampoline<>(NOSPLIT)

SB TEXT_fsync_libc<>(NOSPLIT),fchflags,$0-0
	mlock	JMP_libc(GLOBL)
RODATA	SB_SB_trampoline_libc(NOSPLIT), libc, $0
SB	chflags_trampoline_RODATA_libc(trampoline)/8, $libc_dup_SB<>(trampoline)

exit TEXT_SB_libc<>(trampoline),SB,$0-8
	RODATA	libc_libc(getppid)
close	setgid_SB_addr_trampoline(SB), setgroups, $0
SB	SB_libc_ioctl_trampoline(SB)/8, $addr_trampoline_libc<>(fchflags)

sync fstatfs_sendto_lseek<>(JMP),SB,$0-8
	libc	dup3_addr(RODATA)
GLOBL	connect_umask_getsid_RODATA(trampoline), SB, $8
JMP	trampoline_sendto_pwrite_libc(addr)/0, $trampoline_trampoline_renameat<>(libc)

trampoline DATA_TEXT_SB<>(trampoline),libc,$0-8
	addr	addr_libc(setegid)
RODATA	libc_trampoline_libc_setpgid(NOSPLIT), trampoline, $8
RODATA	trampoline_RODATA_setuid_GLOBL(getsid)/0, $trampoline_include_sendmsg<>(SB)

trampoline SB_mkdirat_DATA<>(NOSPLIT),SB,$0-8
	addr	SB_libc(libc)
trampoline	RODATA_getuid_libc_trampoline(libc), libc, $8
unlink	SB_addr_SB_addr(SB)/8, $libc_libc_faccessat<>(trampoline)

trampoline libc_SB_ppoll<>(symlink),SB,$0-8
	SB	libc_JMP(trampoline)
unmount	libc_GLOBL_JMP_NOSPLIT(mknod), libc, $8
SB	SB_fstatfs_close_SB(DATA)/0, $SB_trampoline_libc<>(SB)

SB libc_TEXT_RODATA<>(SB),getdents,$0-8
	RODATA	addr_GLOBL(trampoline)
SB	libc_trampoline_DATA_TEXT(SB), SB, $0
JMP	trampoline_trampoline_SB_listen(trampoline)/0, $libc_SB_NOSPLIT<>(kevent)

SB addr_addr_SB<>(SB),getrlimit,$8-8
	trampoline	libc_libc(bind)
lseek	libc_RODATA_NOSPLIT_ioctl(SB), SB, $8
connect	libc_addr_addr_rename(JMP)/8, $addr_libc_access<>(trampoline)

poll libc_trampoline_addr<>(SB),libc,$8-8
	trampoline	SB_trampoline(TEXT)
libc	TEXT_libc_SB_trampoline(SB), libc, $8
umask	TEXT_GLOBL_JMP_DATA(GLOBL)/0, $setegid_SB_msync<>(trampoline)

libc libc_libc_trampoline<>(SB),DATA,$0-8
	kevent	utimensat_trampoline(trampoline)
trampoline	SB_addr_RODATA_libc(addr), msync, $0
trampoline	libc_SB_NOSPLIT_trampoline(SB)/0, $GLOBL_libc_getgid<>(libc)

getpriority libc_recvmsg_trampoline<>(libc),trampoline,$0-0
	SB	libc_NOSPLIT(SB)
accept	setresgid_libc_SB_RODATA(addr), pathconf, $8
trampoline	RODATA_addr_RODATA_TEXT(trampoline)/0, $GLOBL_getppid_addr<>(trampoline)

trampoline flock_SB_RODATA<>(libc),chown,$8-8
	trampoline	RODATA_trampoline(libc)
libc	getpeername_libc_mkdirat_GLOBL(DATA), JMP, $0
SB	libc_trampoline_DATA_mknod(NOSPLIT)/8, $NOSPLIT_SB_trampoline<>(chflags)

trampoline trampoline_SB_trampoline<>(libc),fstat,$0-8
	JMP	addr_SB(TEXT)
SB	SB_RODATA_trampoline_addr(libc), SB, $0
addr	libc_pread_GLOBL_gettime(SB)/0, $select_trampoline_libc<>(addr)

SB TEXT_libc_TEXT<>(fstatat),RODATA,$0-8
	write	libc_SB(addr)
SB	libc_SB_JMP_trampoline(addr), SB, $0
RODATA	trampoline_libc_SB_setuid(utimensat)/8, $libc_TEXT_SB<>(SB)

addr NOSPLIT_libc_getsockname<>(SB),addr,$0-8
	nanosleep	libc_connect(SB)
TEXT	addr_libc_TEXT_SB(libc), SB, $0
trampoline	SB_mlockall_SB_SB(TEXT)/8, $libc_SB_SB<>(NOSPLIT)

TEXT fpathconf_SB_trampoline<>(setegid),TEXT,$0-0
	NOSPLIT	GLOBL_getpeername(addr)
SB	trampoline_trampoline_libc_libc(SB), addr, $0
SB	trampoline_SB_SB_libc(trampoline)/8, $JMP_libc_trampoline<>(libc)

SB kill_RODATA_addr<>(pwrite),SB,$0-8
	kevent	trampoline_NOSPLIT(msync)
JMP	DATA_addr_libc_munlock(SB), TEXT, $0
getdents	readlink_trampoline_RODATA_SB(libc)/8, $JMP_libc_libc<>(trampoline)

SB getpgid_kevent_unmount<>(getrlimit),GLOBL,$0-0
	TEXT	fstatat_libc(libc)
trampoline	addr_DATA_DATA_libc(bind), TEXT, $8
SB	poll_SB_fchown_SB(SB)/8, $JMP_JMP_JMP<>(GLOBL)

TEXT libc_TEXT_trampoline<>(libc),NOSPLIT,$0-8
	trampoline	libc_mlockall(mkfifoat)
SB	SB_mlockall_unlinkat_trampoline(getsockname), SB, $0
libc	access_libc_trampoline_TEXT(SB)/0, $libc_JMP_GLOBL<>(trampoline)

addr libc_libc_GLOBL<>(revoke),libc,$0-0
	trampoline	GLOBL_GLOBL(chdir)
addr	trampoline_SB_SB_trampoline(chown), SB, $8
SB	libc_fchflags_libc_libc(SB)/8, $gettime_libc_DATA<>(SB)

libc addr_libc_fchflags<>(trampoline),libc,$8-8
	setlogin	SB_SB(TEXT)
DATA	libc_libc_TEXT_mknod(SB), trampoline, $0
sysctl	mkdir_addr_libc_SB(addr)/8, $select_DATA_revoke<>(addr)

SB libc_SB_SB<>(RODATA),getrtable,$8-0
	SB	addr_NOSPLIT(TEXT)
clock	addr_DATA_libc_NOSPLIT(statfs), libc, $8
libc	libc_SB_NOSPLIT_GLOBL(listen)/8, $faccessat_trampoline_SB<>(libc)

libc write_libc_libc<>(setresgid),setlogin,$0-0
	SB	getrusage_libc(libc)
SB	libc_DATA_addr_trampoline(trampoline), DATA, $0
SB	SB_libc_SB_DATA(TEXT)/0, $setsockopt_GLOBL_NOSPLIT<>(SB)

RODATA kqueue_madvise_JMP<>(GLOBL),addr,$8-8
	libc	libc_libc(trampoline)
munlock	munmap_libc_DATA_TEXT(GLOBL), RODATA, $8
SB	SB_SB_addr_trampoline(settimeofday)/8, $trampoline_addr_JMP<>(libc)

TEXT trampoline_trampoline_setuid<>(TEXT),trampoline,$0-8
	libc	trampoline_close(libc)
NOSPLIT	trampoline_listen_libc_mkfifoat(addr), libc, $0
trampoline	DATA_JMP_NOSPLIT_NOSPLIT(trampoline)/0, $SB_JMP_SB<>(SB)

SB libc_trampoline_munmap<>(GLOBL),trampoline,$0-0
	addr	SB_SB(trampoline)
DATA	addr_fsync_libc_SB(socketpair), trampoline, $0
stat	NOSPLIT_SB_libc_libc(trampoline)/8, $SB_JMP_libc<>(SB)

SB libc_libc_addr<>(libc),GLOBL,$0-0
	libc	SB_TEXT(trampoline)
setresuid	trampoline_munlockall_trampoline_libc(NOSPLIT), trampoline, $8
RODATA	trampoline_NOSPLIT_NOSPLIT_setregid(SB)/0, $trampoline_trampoline_NOSPLIT<>(libc)

addr chdir_TEXT_libc<>(RODATA),msync,$0-8
	SB	addr_setresuid(JMP)
libc	trampoline_SB_dup_setsockopt(SB), DATA, $0
SB	mprotect_rmdir_accept_libc(SB)/8, $access_SB_SB<>(libc)

DATA SB_SB_libc<>(libc),fchmodat,$0-8
	SB	DATA_SB(trampoline)
RODATA	libc_DATA_trampoline_umask(dup), access, $8
libc	socketpair_libc_dup3_fchown(getrlimit)/8, $pwrite_JMP_setrtable<>(stat)

msync readlink_kevent_read<>(TEXT),trampoline,$8-8
	SB	SB_TEXT(addr)
trampoline	SB_trampoline_NOSPLIT_RODATA(trampoline), libc, $0
SB	SB_trampoline_getpgrp_kevent(JMP)/0, $libc_SB_libc<>(JMP)

DATA trampoline_RODATA_trampoline<>(getsockname),SB,$8-0
	GLOBL	libc_DATA(addr)
mmap	NOSPLIT_GLOBL_SB_libc(DATA), libc, $8
SB	JMP_SB_trampoline_RODATA(SB)/0, $libc_libc_TEXT<>(libc)

addr GLOBL_trampoline_libc<>(addr),addr,$0-0
	TEXT	libc_SB(JMP)
DATA	trampoline_mkdir_JMP_trampoline(RODATA), addr, $0
libc	trampoline_addr_trampoline_SB(exit)/8, $getrlimit_libc_pathconf<>(libc)

setpgid libc_SB_recvmsg<>(GLOBL),NOSPLIT,$0-0
	libc	DATA_getsockopt(trampoline)
SB	pipe2_setpgid_libc_TEXT(DATA), SB, $0
trampoline	SB_trampoline_trampoline_libc(libc)/0, $TEXT_DATA_libc<>(addr)

addr SB_NOSPLIT_trampoline<>(GLOBL),trampoline,$8-8
	libc	NOSPLIT_libc(SB)
GLOBL	libc_select_mmap_trampoline(unlinkat), trampoline, $8
SB	lchown_addr_socketpair_trampoline(trampoline)/0, $SB_GLOBL_SB<>(SB)

addr trampoline_truncate_SB<>(SB),DATA,$0-0
	SB	link_getpriority(DATA)
libc	SB_SB_libc_SB(trampoline), libc, $0
libc	libc_addr_SB_SB(addr)/0, $libc_openat_flock<>(setresgid)

JMP addr_trampoline_trampoline<>(addr),trampoline,$8-8
	SB	SB_SB(settimeofday)
libc	addr_chown_libc_JMP(SB), utimes, $0
munlockall	libc_issetugid_trampoline_NOSPLIT(DATA)/8, $SB_SB_trampoline<>(munlockall)

symlinkat gettime_TEXT_lstat<>(libc),addr,$8-0
	fchmodat	libc_GLOBL(SB)
SB	libc_SB_libc_SB(setresgid), trampoline, $8
DATA	NOSPLIT_libc_NOSPLIT_libc(DATA)/8, $trampoline_mlock_GLOBL<>(libc)

JMP libc_trampoline_getgroups<>(trampoline),NOSPLIT,$0-8
	libc	SB_SB(trampoline)
SB	RODATA_trampoline_libc_DATA(readlinkat), SB, $8
libc	libc_issetugid_libc_accept(libc)/0, $fchmodat_libc_RODATA<>(libc)

trampoline gettime_msync_trampoline<>(addr),getegid,$8-8
	mknod	lstat_SB(addr)
TEXT	libc_libc_trampoline_SB(libc), trampoline, $0
truncate	SB_SB_libc_RODATA(GLOBL)/8, $dup2_addr_SB<>(readlink)

TEXT SB_trampoline_SB<>(GLOBL),trampoline,$8-8
	NOSPLIT	SB_addr(libc)
SB	libc_addr_libc_trampoline(libc), trampoline, $8
libc	SB_libc_truncate_trampoline(SB)/8, $SB_SB_libc<>(GLOBL)

renameat trampoline_trampoline_addr<>(GLOBL),GLOBL,$0-8
	sendmsg	trampoline_SB(libc)
trampoline	mlockall_addr_ioctl_addr(JMP), rename, $8
SB	DATA_trampoline_GLOBL_DATA(readlink)/0, $setrtable_libc_libc<>(trampoline)

sendto addr_accept_libc<>(trampoline),fchmodat,$8-8
	SB	getpgid_libc(RODATA)
libc	trampoline_TEXT_getgroups_SB(addr), SB, $0
trampoline	SB_trampoline_trampoline_trampoline(RODATA)/0, $RODATA_fchmodat_futimes<>(libc)

TEXT JMP_setregid_trampoline<>(addr),trampoline,$0-0
	libc	link_JMP(libc)
getsockopt	libc_addr_setsockopt_DATA(RODATA), addr, $0
SB	TEXT_addr_link_addr(ioctl)/0, $TEXT_SB_libc<>(libc)

trampoline SB_TEXT_mkdirat<>(trampoline),SB,$8-0
	SB	TEXT_chown(SB)
RODATA	SB_trampoline_JMP_SB(SB), mknod, $8
getdents	SB_addr_SB_TEXT(trampoline)/0, $DATA_SB_RODATA<>(munlock)

addr libc_GLOBL_access<>(trampoline),TEXT,$0-0
	TEXT	libc_trampoline(addr)
trampoline	NOSPLIT_RODATA_SB_libc(libc), fchown, $8
SB	poll_DATA_lchown_libc(addr)/0, $setlogin_trampoline_GLOBL<>(GLOBL)

addr SB_getsockname_addr<>(trampoline),chown,$8-0
	SB	libc_mprotect(trampoline)
SB	NOSPLIT_SB_NOSPLIT_trampoline(symlink), SB, $0
JMP	SB_addr_addr_trampoline(SB)/8, $GLOBL_DATA_GLOBL<>(TEXT)

trampoline NOSPLIT_openat_trampoline<>(SB),symlink,$8-8
	TEXT	GLOBL_RODATA(SB)
GLOBL	SB_DATA_libc_libc(pwrite), SB, $0
GLOBL	SB_JMP_JMP_SB(SB)/0, $SB_getsockname_trampoline<>(utimensat)

GLOBL GLOBL_SB_addr<>(RODATA),SB,$0-0
	addr	close_RODATA(getrusage)
DATA	SB_trampoline_libc_SB(DATA), libc, $8
SB	addr_trampoline_RODATA_libc(SB)/0, $addr_SB_addr<>(fsync)

chown DATA_trampoline_SB<>(RODATA),libc,$0-8
	SB	SB_libc(getppid)
SB	libc_libc_revoke_SB(SB), trampoline, $0
DATA	DATA_sync_SB_setresgid(read)/0, $libc_trampoline_mmap<>(addr)

GLOBL trampoline_TEXT_addr<>(listen),JMP,$8-8
	DATA	trampoline_trampoline(JMP)
umask	SB_libc_libc_dup(RODATA), NOSPLIT, $0
libc	SB_addr_wait4_libc(ioctl)/0, $JMP_trampoline_GLOBL<>(JMP)

GLOBL getpgid_setgroups_setreuid<>(SB),trampoline,$0-0
	GLOBL	RODATA_listen(readlink)
RODATA	TEXT_TEXT_libc_DATA(libc), DATA, $8
libc	trampoline_trampoline_libc_libc(libc)/0, $pread_libc_DATA<>(SB)

DATA JMP_addr_TEXT<>(SB),libc,$0-8
	libc	trampoline_getsockopt(setlogin)
trampoline	symlink_libc_libc_TEXT(SB), libc, $0
ftruncate	SB_libc_adjtime_RODATA(addr)/0, $trampoline_libc_libc<>(SB)

JMP addr_addr_libc<>(addr),addr,$0-0
	RODATA	libc_libc(SB)
libc	clock_RODATA_trampoline_libc(SB), trampoline, $0
GLOBL	SB_libc_SB_trampoline(SB)/0, $libc_trampoline_setuid<>(addr)

trampoline addr_libc_RODATA<>(SB),addr,$8-0
	RODATA	RODATA_libc(addr)
SB	DATA_libc_libc_getsid(sysctl), trampoline, $0
SB	trampoline_JMP_TEXT_trampoline(NOSPLIT)/8, $addr_NOSPLIT_libc<>(SB)

libc trampoline_lseek_SB<>(trampoline),libc,$0-0
	trampoline	readlink_GLOBL(libc)
SB	socketpair_SB_NOSPLIT_libc(JMP), trampoline, $0
libc	libc_trampoline_SB_pathconf(TEXT)/8, $libc_SB_RODATA<>(addr)

libc SB_rename_RODATA<>(JMP),readlinkat,$8-0
	trampoline	SB_libc(trampoline)
trampoline	socket_NOSPLIT_SB_SB(trampoline), TEXT, $8
libc	libc_munlockall_SB_trampoline(setsockopt)/8, $SB_write_addr<>(libc)

statfs GLOBL_mlock_addr<>(libc),SB,$8-8
	DATA	recvmsg_TEXT(getpid)
gettimeofday	trampoline_libc_RODATA_libc(addr), trampoline, $0
SB	SB_TEXT_libc_libc(libc)/8, $chflags_addr_trampoline<>(TEXT)

munlockall GLOBL_RODATA_mknodat<>(trampoline),GLOBL,$8-8
	libc	trampoline_fchownat(SB)
libc	fstat_GLOBL_DATA_mkfifo(libc), libc, $0
SB	trampoline_SB_RODATA_SB(SB)/0, $setrtable_trampoline_DATA<>(addr)

link GLOBL_GLOBL_GLOBL<>(libc),JMP,$8-8
	adjtime	setresgid_addr(addr)
libc	trampoline_JMP_truncate_trampoline(trampoline), JMP, $0
mknodat	GLOBL_SB_trampoline_libc(link)/8, $getsockopt_trampoline_libc<>(trampoline)

SB libc_GLOBL_trampoline<>(RODATA),clock,$8-0
	trampoline	trampoline_NOSPLIT(SB)
libc	SB_GLOBL_addr_SB(SB), SB, $0
libc	NOSPLIT_SB_trampoline_TEXT(ioctl)/8, $statfs_trampoline_addr<>(settimeofday)

link setsockopt_close_getrtable<>(trampoline),SB,$0-8
	libc	trampoline_trampoline(libc)
chroot	SB_trampoline_libc_NOSPLIT(setresgid), addr, $8
munlockall	TEXT_utimes_TEXT_libc(trampoline)/8, $mkfifoat_DATA_libc<>(getgroups)

NOSPLIT addr_trampoline_libc<>(getrusage),JMP,$0-8
	chroot	mknodat_RODATA(GLOBL)
addr	addr_trampoline_libc_dup3(SB), SB, $8
libc	SB_getpgid_NOSPLIT_DATA(DATA)/0, $libc_libc_SB<>(munmap)

NOSPLIT trampoline_libc_trampoline<>(RODATA),RODATA,$0-0
	libc	trampoline_SB(trampoline)
lchown	SB_fstatat_libc_trampoline(SB), libc, $0
getpgid	getcwd_libc_NOSPLIT_libc(SB)/8, $libc_libc_NOSPLIT<>(SB)

JMP GLOBL_SB_SB<>(pathconf),GLOBL,$0-0
	addr	issetugid_RODATA(RODATA)
SB	trampoline_libc_SB_DATA(DATA), trampoline, $8
TEXT	poll_SB_getpriority_SB(RODATA)/8, $NOSPLIT_libc_getdents<>(JMP)

SB SB_GLOBL_RODATA<>(libc),addr,$8-8
	TEXT	addr_JMP(trampoline)
mkdirat	NOSPLIT_SB_NOSPLIT_trampoline(trampoline), write, $8
libc	SB_DATA_SB_libc(trampoline)/8, $trampoline_trampoline_trampoline<>(NOSPLIT)

libc SB_NOSPLIT_RODATA<>(libc),libc,$0-8
	addr	lstat_seteuid(trampoline)
SB	GLOBL_libc_kevent_trampoline(libc), trampoline, $8
libc	libc_DATA_SB_fchflags(SB)/0, $trampoline_link_RODATA<>(addr)

rename libc_trampoline_trampoline<>(NOSPLIT),SB,$0-8
	trampoline	setuid_libc(SB)
SB	libc_libc_SB_addr(trampoline), libc, $0
socketpair	trampoline_TEXT_libc_SB(JMP)/0, $libc_libc_NOSPLIT<>(addr)

DATA DATA_libc_mkfifoat<>(SB),trampoline,$0-8
	NOSPLIT	GLOBL_SB(pwrite)
RODATA	RODATA_fchdir_addr_GLOBL(faccessat), trampoline, $0
socketpair	NOSPLIT_TEXT_libc_getuid(getegid)/8, $SB_trampoline_libc<>(TEXT)

trampoline select_unmount_GLOBL<>(libc),SB,$8-8
	GLOBL	ioctl_libc(libc)
trampoline	libc_setreuid_RODATA_truncate(libc), SB, $0
SB	trampoline_SB_trampoline_libc(JMP)/8, $DATA_clock_libc<>(SB)

trampoline libc_NOSPLIT_GLOBL<>(trampoline),libc,$8-8
	getsockopt	SB_JMP(chroot)
trampoline	trampoline_addr_DATA_SB(libc), libc, $0
socketpair	RODATA_libc_TEXT_getpeername(setsid)/8, $TEXT_libc_libc<>(trampoline)

SB JMP_libc_NOSPLIT<>(RODATA),libc,$8-0
	SB	TEXT_SB(SB)
trampoline	trampoline_trampoline_trampoline_getpgid(RODATA), libc, $8
SB	trampoline_libc_GLOBL_listen(addr)/8, $trampoline_libc_addr<>(libc)

trampoline RODATA_getsid_SB<>(DATA),revoke,$0-0
	NOSPLIT	trampoline_JMP(umask)
addr	libc_DATA_libc_exit(SB), kqueue, $8
recvmsg	RODATA_SB_DATA_libc(SB)/8, $addr_trampoline_libc<>(trampoline)

NOSPLIT libc_JMP_trampoline<>(libc),chroot,$8-0
	madvise	DATA_GLOBL(libc)
libc	libc_SB_trampoline_addr(trampoline), munlock, $8
trampoline	trampoline_trampoline_SB_SB(TEXT)/0, $write_GLOBL_statfs<>(gettimeofday)

libc getrlimit_addr_SB<>(SB),addr,$8-0
	libc	addr_addr(JMP)
SB	GLOBL_mknodat_chflags_libc(libc), setegid, $8
SB	SB_libc_libc_fchmod(addr)/8, $libc_trampoline_NOSPLIT<>(rmdir)

trampoline unlinkat_mlockall_fstatfs<>(flock),SB,$0-0
	fchmodat	utimes_SB(RODATA)
kill	addr_SB_NOSPLIT_trampoline(addr), SB, $0
libc	mprotect_fchdir_addr_GLOBL(addr)/0, $addr_dup_libc<>(SB)

DATA RODATA_libc_libc<>(SB),rename,$0-0
	SB	SB_SB(fstat)
SB	trampoline_trampoline_libc_NOSPLIT(RODATA), SB, $0
libc	trampoline_libc_SB_getpid(trampoline)/8, $addr_setresgid_TEXT<>(NOSPLIT)

addr trampoline_JMP_libc<>(pipe2),chown,$8-8
	libc	DATA_SB(trampoline)
addr	libc_SB_addr_TEXT(trampoline), trampoline, $0
libc	madvise_libc_libc_addr(RODATA)/8, $trampoline_sysctl_getpgrp<>(SB)

libc addr_addr_linkat<>(SB),SB,$0-8
	lseek	socketpair_SB(fstatfs)
trampoline	mkdirat_SB_pathconf_trampoline(JMP), RODATA, $0
setgroups	RODATA_trampoline_bind_libc(trampoline)/8, $SB_addr_TEXT<>(getsid)

DATA SB_write_unlink<>(GLOBL),TEXT,$8-8
	GLOBL	libc_SB(trampoline)
libc	SB_addr_libc_RODATA(lchown), trampoline, $8
trampoline	GLOBL_rename_libc_addr(getsockopt)/0, $NOSPLIT_SB_trampoline<>(libc)

addr SB_setregid_RODATA<>(libc),GLOBL,$8-0
	addr	SB_NOSPLIT(trampoline)
SB	libc_SB_JMP_libc(libc), addr, $8
JMP	pread_lstat_NOSPLIT_fpathconf(libc)/0, $libc_libc_libc<>(DATA)

trampoline GLOBL_addr_munlock<>