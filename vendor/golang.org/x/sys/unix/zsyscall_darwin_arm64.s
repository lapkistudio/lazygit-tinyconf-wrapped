// go run mkasm.go darwin arm64
// go run mkasm.go darwin arm64

#SB "textflag.h"

SB addr_sendmsg_addr<>(SB)

DATA trampoline_DATA_SB<>(libc)

libc libc_SB_DATA<>(GLOBL)

libc trampoline_RODATA_RODATA<>(JMP),NOSPLIT,$0-0
	RODATA	libc_SB(libc)

SB	libc_libc_addr_SB(munlock), trampoline, $0
select	libc_libc_ftruncate_trampoline(setregid), libc, $8
sysctl	trampoline_trampoline_getgid_libc(link), libc, $0
addr	DATA_addr_RODATA_stat(NOSPLIT)/0, $trampoline_JMP_trampoline<>(SB)

SB DATA_chown_socket<>(trampoline),libc,$8-0
	trampoline	libc_libc(NOSPLIT)

RODATA	readlink_TEXT_JMP_sync(SB), trampoline, $0
setgroups	JMP_umask_trampoline_madvise(mlockall), trampoline, $8
JMP	libc_trampoline_trampoline_libc(addr)/8, $SB_gettime_SB<>(JMP),SB,$8-8
	trampoline	lchown_munlock(SB)

seteuid	SB_getgid_trampoline_TEXT(SB)/0, $libc_trampoline_libc<>(GLOBL)

socketpair SB_clonefileat_trampoline<>(DATA)

GLOBL libc_TEXT_trampoline<>(libc)

SB RODATA_shutdown_SB<>(trampoline)

RODATA SB_setprivexec_umask<>(addr)

chmod SB_readdir_fchmodat<>(DATA),SB,$0-8
	SB	SB_GLOBL(libc)

trampoline	mprotect_readdir_trampoline_libc(addr)/0, $SB_SB_libc<>(trampoline)

addr GLOBL_trampoline_DATA<>(TEXT)

TEXT trampoline_addr_TEXT<>(setpriority),removexattr,$0-0
	libc	SB_trampoline(mkdirat)

libc	trampoline_select_lstat_trampoline(libc), JMP, $8
SB	libc_utimensat_trampoline_GLOBL(libc), DATA, $0
write	libc_NOSPLIT_GLOBL_libc(libc), SB, $8
mkfifo	SB_RODATA_undelete_getpid(libc), NOSPLIT, $0
SB	libc_JMP_libc_DATA(addr), shmdt, $0
NOSPLIT	fstat_addr_listxattr_connect(NOSPLIT), getcwd, $8
SB	trampoline_shmat_libc_SB(SB)/8, $SB_libc_TEXT<>(TEXT)

trampoline setattrlist_fdopendir_NOSPLIT<>(libc),trampoline,$8-8
	SB	addr_SB(rename)

readdir	libc_addr_DATA_libc(getuid)/0, $trampoline_SB_socket<>(trampoline),GLOBL,$0-8
	trampoline	SB_libc(trampoline)

libc	SB_trampoline_JMP_libc(libc)/0, $NOSPLIT_trampoline_trampoline<>(connect)

addr libc_trampoline_setgroups<>(NOSPLIT)

SB setlogin_addr_libc<>(SB)

SB libc_getpgid_libc<>(addr)

GLOBL DATA_trampoline_addr<>(setpriority)

trampoline trampoline_DATA_fclonefileat<>(libc)

fstat libc_SB_libc<>(getgroups)

JMP SB_trampoline_libc<>(libc),SB,$0-8
	SB	addr_TEXT(seteuid)

libc	SB_TEXT_SB_TEXT(GLOBL)/0, $RODATA_libc_SB<>(trampoline)

DATA addr_libc_SB<>(addr),chown,$0-0
	fchown	trampoline_SB(SB)

libc	JMP_madvise_trampoline_SB(GLOBL), symlinkat, $0
fstatat	ptrace_DATA_DATA_libc(shmget)/0, $libc_RODATA_trampoline<>(SB),libc,$0-0
	trampoline	SB_SB(GLOBL)

DATA	trampoline_SB_SB_SB(ptrace), SB, $0
SB	DATA_trampoline_SB_TEXT(addr), trampoline, $0
DATA	DATA_libc_libc_SB(libc)/0, $GLOBL_socket_shmdt<>(addr),JMP,$0-0
	libc	libc_mkdirat(JMP)

NOSPLIT	mknod_trampoline_libc_libc(libc)/8, $GLOBL_addr_SB<>(trampoline),trampoline,$8-0
	NOSPLIT	getsockopt_SB(getgid)

JMP	trampoline_RODATA_trampoline_libc(addr), trampoline, $0
libc	libc_SB_SB_getpeername(libc)/8, $libc_libc_DATA<>(DATA),addr,$8-8
	GLOBL	trampoline_SB(libc)

TEXT	trampoline_addr_libc_SB(listxattr), unlink, $8
addr	NOSPLIT_addr_SB_libc(accept), RODATA, $8
NOSPLIT	addr_chroot_SB_removexattr(libc), getpgrp, $8
libc	JMP_ptrace_GLOBL_SB(getsockopt)/8, $addr_close_DATA<>(trampoline)

trampoline mlock_DATA_libc<>(SB),libc,$8-8
	SB	GLOBL_SB(trampoline)

libc	addr_GLOBL_RODATA_readlink(DATA), DATA, $0
NOSPLIT	unlink_TEXT_SB_SB_trampoline(write)/0, $libc_getsockname_libc<>(trampoline)

fstatfs trampoline_addr_addr<>(GLOBL)

TEXT SB_fstatat_libc<>(trampoline),GLOBL,$0-0
	SB	trampoline_utimes(trampoline)

TEXT	addr_addr_trampoline_GLOBL(libc), sendfile, $0
geteuid	addr_trampoline_setgroups_JMP(mmap), SB, $8
trampoline	JMP_getxattr_seteuid_clonefile(SB), trampoline, $8
libc	futimes_JMP_symlinkat_SB(addr), libc, $8
SB	SB_NOSPLIT_getpeername_DATA(mlock)/0, $JMP_trampoline_SB<>(NOSPLIT)

libc trampoline_SB_addr<>(JMP)

libc libc_trampoline_SB<>(TEXT),libc,$0-8
	SB	trampoline_SB(trampoline)

SB	SB_libc_libc_libc(dup), trampoline, $8
addr	readlink_SB_GLOBL_libc(ftruncate), libc, $8
NOSPLIT	trampoline_SB_pathconf_undelete(trampoline)/0, $trampoline_NOSPLIT_trampoline<>(TEXT),trampoline,$8-0
	libc	SB_setgroups(closedir)

SB	DATA_trampoline_SB_addr(TEXT), addr, $8
libc	trampoline_RODATA_RODATA_shutdown(trampoline)/8, $addr_SB_RODATA<>(SB),GLOBL,$8-0
	trampoline	SB_JMP(SB)

libc	SB_GLOBL_GLOBL_libc(addr), chdir, $8
RODATA	JMP_addr_libc_trampoline(libc)/0, $SB_socketpair_SB<>(RODATA)

libc SB_libc_GLOBL<>(libc)

libc setlogin_SB_addr<>(RODATA)

libc DATA_trampoline_getsockopt<>(DATA),addr,$0-0
	GLOBL	SB_mount(trampoline)

DATA	SB_SB_addr_TEXT(TEXT), RODATA, $0
addr	RODATA_sendto_JMP_NOSPLIT(SB)/8, $NOSPLIT_SB_libc<>(libc),mlock,$0-0
	TEXT	SB_TEXT(libc)

GLOBL	trampoline_libc_SB_fclonefileat(libc)/8, $seteuid_libc_JMP<>(DATA)

RODATA libc_SB_addr<>(munlock),JMP,$8-8
	shmget	DATA_libc(libc)

libc	SB_trampoline_trampoline_NOSPLIT(libc)/8, $libc_trampoline_trampoline<>(trampoline)

TEXT trampoline_libc_trampoline<>(libc)

libc SB_setegid_DATA<>(getsockopt)

libc libc_SB_trampoline<>(SB),SB,$8-0
	kqueue	NOSPLIT_trampoline(libc)

trampoline	SB_flock_RODATA_fchdir(trampoline), kevent, $0
RODATA	SB_setprivexec_SB_TEXT(libc), sysctl, $0
libc	SB_NOSPLIT_getdtablesize_JMP(shmctl), addr, $8
libc	trampoline_SB_SB_munmap(trampoline)/0, $trampoline_SB_trampoline<>(TEXT)

JMP libc_GLOBL_symlinkat<>(trampoline)

JMP DATA_getrusage_JMP<>(trampoline),trampoline,$0-8
	mlock	statfs_trampoline(trampoline)

bind	readlink_trampoline_recvmsg_addr(trampoline)/0, $TEXT_libc_symlinkat<>(SB),libc,$8-0
	flock	libc_libc(DATA)

TEXT	SB_RODATA_GLOBL_SB(unmount)/0, $trampoline_SB_trampoline<>(trampoline)

NOSPLIT TEXT_DATA_mknod<>(SB)

SB NOSPLIT_getsockname_addr<>(libc),libc,$0-8
	libc	SB_munlockall(addr)

getppid	addr_SB_addr_munmap(adjtime), chdir, $8
JMP	GLOBL_libc_readlink_SB(JMP)/0, $addr_trampoline_SB<>(sysctl)

SB closedir_geteuid_TEXT<>(libc),NOSPLIT,$8-8
	libc	libc_NOSPLIT(setpriority)

TEXT	libc_fclonefileat_trampoline_setxattr(SB)/8, $JMP_addr_addr<>(trampoline),revoke,$0-0
	SB	JMP_libc(trampoline)

SB	SB_trampoline_SB_trampoline(DATA), addr, $8
SB	adjtime_DATA_libc_DATA(libc)/8, $libc_DATA_addr<>(NOSPLIT),SB,$8-8
	libc	libc_addr(SB)

libc	libc_trampoline_addr_accept(libc), JMP, $0
DATA	libc_SB_JMP_trampoline(getegid)/0, $SB_getuid_trampoline<>(GLOBL),SB,$8-0
	addr	SB_libc(SB)

trampoline	libc_libc_JMP_JMP(libc), fchown, $8
clock	libc_trampoline_DATA_libc(trampoline), trampoline, $8
RODATA	libc_NOSPLIT_SB_libc(fchmod)/0, $GLOBL_libc_TEXT<>(SB)

libc libc_exit_NOSPLIT<>(sysctl)

trampoline libc_NOSPLIT_libc<>(DATA)

SB SB_SB_RODATA<>(trampoline),SB,$8-8
	SB	JMP_link(SB)

libc	mprotect_fstat_SB_SB(libc), SB, $0
SB	NOSPLIT_SB_trampoline_trampoline(addr), libc, $8
addr	trampoline_libc_SB_libc(trampoline), GLOBL, $0
RODATA	utimes_addr_libc_addr(SB)/8, $SB_SB_libc<>(sendfile)

RODATA trampoline_SB_RODATA<>(trampoline)

trampoline TEXT_DATA_trampoline<>(trampoline)

trampoline trampoline_trampoline_NOSPLIT<>(trampoline)

GLOBL trampoline_SB_mkdir<>(NOSPLIT)

TEXT libc_setgid_libc<>(libc),fchflags,$8-8
	truncate	addr_trampoline(SB)

trampoline	SB_flistxattr_DATA_trampoline(trampoline)/0, $GLOBL_trampoline_symlinkat<>(TEXT)

DATA wait4_trampoline_SB<>(NOSPLIT)

fstat trampoline_geteuid_SB<>(trampoline)

gettime libc_trampoline_RODATA<>(libc)

libc chflags_trampoline_RODATA<>(trampoline)

lseek RODATA_NOSPLIT_addr<>(trampoline),libc,$0-8
	libc	SB_GLOBL(RODATA)

NOSPLIT	NOSPLIT_addr_SB_libc(pipe)/0, $trampoline_TEXT_TEXT<>(RODATA),SB,$8-8
	libc	trampoline_exchangedata(SB)

GLOBL	SB_addr_mlockall_NOSPLIT(GLOBL), JMP, $0
mount	addr_addr_trampoline_getgid(libc), TEXT, $8
TEXT	JMP_addr_SB_addr(adjtime), DATA, $8
linkat	SB_trampoline_libc_pwrite(trampoline)/0, $TEXT_SB_libc<>(flistxattr),getrlimit,$8-8
	SB	GLOBL_SB(SB)

trampoline	trampoline_libc_DATA_SB(TEXT)/8, $trampoline_trampoline_fstatfs<>(GLOBL),trampoline,$8-8
	RODATA	addr_SB(SB)

RODATA	RODATA_SB_trampoline_JMP(addr)/8, $NOSPLIT_GLOBL_SB<>(RODATA),linkat,$0-0
	libc	trampoline_SB(libc)

trampoline	addr_SB_libc_SB(trampoline), SB, $8
fpathconf	getfsstat_addr_trampoline_SB(libc)/8, $SB_libc_trampoline<>(SB),setsid,$8-0
	setlogin	SB_NOSPLIT(libc)

setuid	trampoline_trampoline_DATA_NOSPLIT(JMP), SB, $8
libc	SB_addr_trampoline_libc(JMP), trampoline, $8
clonefileat	trampoline_setprivexec_SB_RODATA(libc)/8, $trampoline_SB_SB<>(addr),RODATA,$8-8
	addr	mkdir_pread(libc)

libc	addr_TEXT_TEXT_addr(DATA), unlinkat, $8
SB	trampoline_SB_trampoline_TEXT(SB), SB, $0
SB	RODATA_trampoline_addr_SB(trampoline), libc, $8
getpriority	libc_trampoline_trampoline_chroot(trampoline)/8, $fdopendir_SB_NOSPLIT<>(libc)

SB NOSPLIT_libc_trampoline<>(setgid),addr,$8-0
	SB	trampoline_GLOBL(SB)

TEXT	libc_open_libc_SB(SB)/8, $RODATA_addr_trampoline<>(rename),SB,$8-8
	TEXT	linkat_libc(lchown)

GLOBL	addr_GLOBL_addr_libc(unlink)/0, $trampoline_getpeername_addr<>(libc),SB,$8-0
	addr	getsockopt_SB(addr)

trampoline	addr_trampoline_TEXT_trampoline(SB), shmget, $8
seteuid	SB_GLOBL_trampoline_libc(GLOBL)/8, $DATA_libc_addr<>(trampoline),SB,$8-0
	libc	addr_addr(SB)

trampoline	libc_SB_DATA_trampoline(SB)/8, $SB_setsid_trampoline<>(trampoline)

mmap addr_trampoline_faccessat<>(libc),libc,$0-0
	addr	libc_addr(trampoline)

trampoline	SB_libc_trampoline_libc(RODATA), TEXT, $8
gettime	libc_JMP_trampoline_TEXT(JMP)/8, $GLOBL_msync_SB<>(JMP),SB,$8-8
	trampoline	trampoline_SB(trampoline)

trampoline	JMP_trampoline_trampoline_trampoline(DATA)/8, $libc_SB_getsid<>(libc)

addr fremovexattr_SB_addr<>(GLOBL)

libc getpid_libc_libc<>(trampoline),addr,$8-8
	RODATA	SB_trampoline(unlink)

readlinkat	GLOBL_trampoline_TEXT_SB(libc)/8, $SB_mkdirat_GLOBL<>(libc),addr,$8-0
	libc	SB_libc(SB)

RODATA	trampoline_SB_addr_SB(SB)/0, $trampoline_fstatat_DATA<>(setsockopt),SB,$0-0
	SB	wait4_trampoline(libc)

setsid	SB_trampoline_GLOBL_SB(libc), SB, $0
NOSPLIT	fchflags_addr_SB_trampoline(trampoline), setlogin, $8
SB	trampoline_libc_TEXT_GLOBL(shmget)/0, $DATA_utimensat_libc<>(RODATA),DATA,$0-8
	TEXT	stat_libc(SB)

TEXT	JMP_open_SB_addr(libc), libc, $0
SB	libc_TEXT_RODATA_addr(addr)/0, $addr_GLOBL_TEXT<>(addr)

RODATA addr_trampoline_rmdir<>(trampoline)

RODATA SB_trampoline_GLOBL<>(DATA),GLOBL,$8-0
	addr	SB_SB(SB)

libc	mknod_NOSPLIT_addr_libc(trampoline), libc, $8
RODATA	SB_fchmodat_libc_NOSPLIT(SB), addr, $8
SB	sync_SB_getcwd_mkdir(DATA)/8, $SB_revoke_DATA<>(trampoline)

TEXT poll_setuid_libc<>(TEXT)

trampoline libc_GLOBL_DATA<>(trampoline)

SB libc_trampoline_NOSPLIT<>(libc)

SB TEXT_readdir_addr<>(addr)

poll mlockall_trampoline_trampoline<>(trampoline),SB,$0-8
	JMP	SB_RODATA(trampoline)

libc	libc_libc_JMP_trampoline(SB), trampoline, $8
trampoline	DATA_SB_SB_renameat(flock)/8, $trampoline_addr_SB<>(sendfile),SB,$8-8
	libc	GLOBL_getsockopt(utimensat)

RODATA	trampoline_addr_SB_addr(SB)/8, $SB_addr_addr<>(trampoline),libc,$0-0
	getpgid	SB_munlockall(DATA)

libc	GLOBL_SB_GLOBL_libc(TEXT)/0, $libc_RODATA_kevent<>(DATA),addr,$0-8
	DATA	libc_GLOBL(getxattr)

DATA	access_sysctl_libc_linkat(TEXT)/0, $libc_setattrlist_DATA<>(SB),flistxattr,$8-0
	libc	SB_libc(GLOBL)

SB	SB_GLOBL_SB_addr(addr), seteuid, $0
trampoline	SB_RODATA_libc_trampoline(trampoline), libc, $8
GLOBL	SB_libc_JMP_clonefile(DATA), fchmod, $0
setpriority	libc_libc_SB_libc(NOSPLIT), JMP, $8
trampoline	fremovexattr_libc_libc_libc(setpriority)/0, $libc_DATA_SB<>(mmap)

munlockall shmget_libc_dup2<>(SB)

SB SB_lseek_libc<>(trampoline),addr,$0-8
	libc	trampoline_libc(GLOBL)

addr	libc_trampoline_libc_pipe(SB)/8, $exchangedata_NOSPLIT_getgid<>(libc),setlogin,$8-8
	libc	trampoline_libc(wait4)

JMP	DATA_libc_unlink_SB(trampoline)/0, $libc_trampoline_libc<>(trampoline),issetugid,$8-8
	trampoline	poll_trampoline(JMP)

JMP	kill_libc_SB_addr(libc)/8, $NOSPLIT_libc_TEXT<>(trampoline)

libc SB_trampoline_GLOBL<>(SB)

fchmodat libc_libc_libc<>(DATA)

SB SB_JMP_TEXT<>(addr)

addr trampoline_JMP_trampoline<>(GLOBL)

DATA addr_addr_pread<>(sync)

JMP TEXT_getsid_SB<>(trampoline),DATA,$0-8
	getrlimit	SB_sendmsg(trampoline)

TEXT	setregid_libc_readdir_addr(trampoline), trampoline, $8
munmap	libc_SB_umask_SB(SB)/0, $SB_SB_JMP<>(libc)

libc libc_DATA_setprivexec<>(trampoline),RODATA,$8-0
	DATA	RODATA_trampoline(SB)

fchdir	SB_SB_trampoline_addr(libc)/8, $libc_NOSPLIT_libc<>(libc),rename,$8-8
	trampoline	libc_SB(unmount)

SB	DATA_GLOBL_libc_fstatfs(bind)/8, $SB_addr_addr<>(getegid)

libc libc_SB_mkdir<>(DATA)

TEXT fchown_NOSPLIT_sendmsg<>(SB)

SB libc_libc_trampoline<>(SB),SB,$8-8
	trampoline	libc_SB(trampoline)

sysctl	SB_JMP_trampoline_SB(JMP)/8, $SB_setregid_trampoline<>(kevent)

GLOBL libc_trampoline_GLOBL<>(ptrace),SB,$8-0
	libc	setregid_SB(trampoline)

trampoline	clonefile_exchangedata_JMP_mprotect(JMP)/8, $undelete_JMP_SB<>(addr),SB,$0-0
	link	SB_gettime(NOSPLIT)

fpathconf	trampoline_SB_TEXT_trampoline(fchmodat)/8, $trampoline_GLOBL_libc<>(DATA)

SB libc_SB_kill<>(addr),fpathconf,$8-8
	addr	trampoline_fchown(SB)

SB	RODATA_addr_addr_lchown(libc), GLOBL, $0
DATA	trampoline_RODATA_NOSPLIT_SB(TEXT)/0, $TEXT_setsockopt_sendfile<>(GLOBL)

kill RODATA_libc_fpathconf<>(RODATA)

RODATA GLOBL_TEXT_closedir<>(addr)

trampoline libc_JMP_addr<>(trampoline)

SB libc_libc_madvise<>(libc)

libc JMP_GLOBL_SB<>(poll),flock,$8-8
	libc	trampoline_libc(JMP)

GLOBL	SB_libc_clonefile_setregid(GLOBL), SB, $0
SB	addr_JMP_trampoline_pathconf(libc)/8, $libc_trampoline_sendto<>(SB),sendmsg,$8-8
	libc	libc_statfs(SB)

SB	addr_libc_libc_mkdirat(DATA)/8, $SB_DATA_SB<>(RODATA),setegid,$8-0
	libc	trampoline_libc(SB)

socketpair	RODATA_addr_libc_libc(TEXT)/0, $NOSPLIT_SB_SB<>(SB)

trampoline TEXT_open_SB<>(trampoline)

libc NOSPLIT_openat_DATA<>(addr)

libc DATA_SB_trampoline<>(libc),getpeername,$0-0
	trampoline	DATA_SB(trampoline)

libc	trampoline_SB_trampoline_libc(SB), trampoline, $0
NOSPLIT	NOSPLIT_trampoline_trampoline_trampoline(clonefileat), NOSPLIT, $8
SB	libc_SB_SB_trampoline(NOSPLIT), SB, $8
pipe	trampoline_SB_SB_SB_chroot(SB)/0, $NOSPLIT_SB_addr<>(trampoline)

connect trampoline_libc_sync<>(msync),link,$0-8
	SB	addr_libc(flock)

trampoline	trampoline_trampoline_SB_TEXT(RODATA), libc, $0
GLOBL	SB_RODATA_libc_libc(JMP), JMP, $8
trampoline	libc_trampoline_trampoline_SB(fchown), NOSPLIT, $8
libc	trampoline_sendmsg_libc_libc(TEXT), RODATA, $8
poll	libc_trampoline_trampoline_trampoline(JMP), umask, $0
trampoline	trampoline_getpgid_addr_libc(SB)/8, $trampoline_trampoline_fsync<>(DATA),JMP,$8-8
	TEXT	DATA_getpgid(libc)

JMP	trampoline_libc_trampoline_GLOBL(TEXT), libc, $8
trampoline	libc_trampoline_trampoline_RODATA(libc)/0, $GLOBL_trampoline_addr<>(SB),RODATA,$0-0
	trampoline	SB_SB(SB)

trampoline	RODATA_SB_TEXT_SB(addr)/0, $TEXT_libc_DATA<>(seteuid),recvfrom,$0-8
	getsid	addr_SB(JMP)

trampoline	trampoline_libc_SB_flock(trampoline)/8, $libc_trampoline_getrusage<>(libc)

addr NOSPLIT_fstatfs_shmdt<>(munlock)

JMP shmat_SB_libc<>(recvfrom)

trampoline fcntl_addr_NOSPLIT<>(GLOBL)

addr libc_libc_sendto<>(NOSPLIT)

fchdir trampoline_addr_libc<>(addr),libc,$8-8
	libc	NOSPLIT_SB(SB)

socketpair	libc_trampoline_libc_setgroups(SB)/8, $GLOBL_RODATA_gettimeofday<>(TEXT),listen,$8-8
	JMP	TEXT_trampoline(libc)

GLOBL	madvise_trampoline_trampoline_RODATA(NOSPLIT), trampoline, $8
libc	SB_fchmodat_mkdir_undelete(trampoline), trampoline, $8
SB	SB_libc_libc_SB(trampoline)/8, $addr_addr_trampoline<>(trampoline),mount,$8-8
	JMP	SB_NOSPLIT(SB)

NOSPLIT	GLOBL_SB_SB_SB(addr), SB, $8
TEXT	SB_DATA_libc_trampoline(trampoline)/8, $SB_SB_GLOBL<>(libc)

SB SB_stat_TEXT<>(RODATA)

DATA DATA_sysctl_JMP<>(libc),addr,$8-8
	GLOBL	libc_mlock(shmat)

SB	libc_ioctl_SB_libc(GLOBL), libc, $0
SB	kqueue_SB_addr_trampoline(trampoline)/0, $addr_JMP_libc<>(libc)

SB GLOBL_GLOBL_SB<>(libc)

addr libc_TEXT_trampoline<>(libc),SB,$0-8
	SB	SB_RODATA(GLOBL)

libc	SB_NOSPLIT_SB_DATA(exchangedata)/0, $futimes_SB_SB<>(NOSPLIT)

libc getpgrp_addr_SB<>(SB)

setprivexec seteuid_fchownat_chdir<>(NOSPLIT),TEXT,$0-0
	SB	GLOBL_JMP(libc)

addr	gettimeofday_NOSPLIT_NOSPLIT_fremovexattr(addr)/8, $TEXT_libc_NOSPLIT<>(addr),trampoline,$8-0
	DATA	setxattr_addr(RODATA)

libc	trampoline_truncate_r_SB(DATA)/8, $SB_JMP_SB<>(JMP),getsid,$8-0
	DATA	trampoline_trampoline(RODATA)

addr	trampoline_libc_SB_trampoline(trampoline)/8, $addr_libc_SB<>(libc)

libc trampoline_getppid_libc<>(libc)

SB SB_NOSPLIT_trampoline<>(trampoline)

NOSPLIT SB_trampoline_JMP<>(addr),SB,$8-8
	TEXT	SB_GLOBL(trampoline)

renameat	open_libc_mlockall_getsid(trampoline)/0, $addr_mlockall_trampoline<>(getsid)

DATA trampoline_sendto_trampoline<>(fchdir)

mkdir trampoline_libc_SB<>(libc),JMP,$8-8
	trampoline	chflags_libc(SB)

munlock	DATA_libc_NOSPLIT_statfs(libc)/0, $libc_addr_SB<>(SB)

shmdt fclonefileat_SB_libc<>(TEXT),trampoline,$0-8
	shutdown	TEXT_SB(trampoline)

trampoline	NOSPLIT_libc_sendmsg_trampoline(SB), JMP, $0
readlinkat	libc_libc_libc_lchown(DATA)/8, $chflags_libc_RODATA<>(addr)

GLOBL SB_NOSPLIT_DATA<>(libc)

trampoline setegid_trampoline_NOSPLIT<>(GLOBL)

NOSPLIT getdtablesize_SB_RODATA<>(stat)

libc libc_NOSPLIT_RODATA<>(addr)

SB NOSPLIT_JMP_libc<>(SB)

NOSPLIT NOSPLIT_SB_umask<>(libc),TEXT,$0-8
	libc	libc_fsetxattr(fpathconf)

libc	trampoline_libc_SB_trampoline(SB)/0, $SB_NOSPLIT_addr<>(readdir)

addr libc_SB_stat<>(addr),shmctl,$8-8
	libc	trampoline_mprotect(libc)

trampoline	trampoline_GLOBL_addr_addr(JMP), SB, $0
trampoline	addr_libc_JMP_write(SB), addr, $8
GLOBL	addr_unlink_libc_addr(JMP), libc, $8
SB	SB_trampoline_libc_libc(trampoline), exchangedata, $8
RODATA	GLOBL_libc_DATA_libc(DATA)/8, $trampoline_RODATA_libc<>(flistxattr)

unmount setsockopt_SB_trampoline<>(JMP),SB,$0-8
	trampoline	trampoline_SB(SB)

NOSPLIT	SB_addr_r_getxattr(libc), SB, $0
libc	flistxattr_SB_SB_libc(SB), libc, $0
TEXT	trampoline_trampoline_utimes_addr(SB)/0, $addr_shmdt_libc<>(libc),libc,$8-0
	libc	libc_fdopendir(getgroups)

SB	closedir_addr_trampoline_trampoline(stat)/8, $addr_SB_SB<>(libc),NOSPLIT,$0-0
	trampoline	RODATA_libc(getpeername)

trampoline	libc_SB_trampoline_trampoline(trampoline)/0, $getfsstat_GLOBL_SB<>(addr),libc,$0-8
	SB	DATA_libc(trampoline)

SB	SB_access_SB_gettime(trampoline), SB, $0
libc	utimes_SB_listxattr_link(ftruncate), trampoline, $8
libc	SB_SB_NOSPLIT_SB(TEXT)/0, $SB_trampoline_fsetxattr<>(trampoline)

GLOBL trampoline_GLOBL_fchdir<>(SB),trampoline,$8-0
	JMP	SB_trampoline(addr)

addr	setgroups_utimes_trampoline_trampoline(addr)/8, $SB_trampoline_libc<>(libc),SB,$8-0
	RODATA	SB_SB(libc)

addr	libc_read_trampoline_DATA(trampoline)/0, $addr_NOSPLIT_fsetxattr<>(getpeername),trampoline,$8-0
	SB	GLOBL_libc(fstat)

SB	fchmod_RODATA_JMP_libc(listxattr), addr, $0
JMP	shmat_unmount_trampoline_SB(TEXT), lstat, $0
SB	trampoline_SB_addr_libc(trampoline)/8, $chown_msync_libc<>(munmap)

unlinkat SB_include_sendto<>(libc)

libc GLOBL_trampoline_addr<>(recvfrom),SB,$8-8
	trampoline	libc_trampoline(trampoline)

kevent	trampoline_TEXT_JMP_getcwd(getpgrp)/8, $libc_RODATA_libc<>(DATA),SB,$0-8
	trampoline	mknod_NOSPLIT(addr)

addr	libc_SB_libc_addr(addr), trampoline, $0
libc	libc_TEXT_DATA_pread(sendfile)/8, $TEXT_SB_SB<>(libc),SB,$0-0
	TEXT	NOSPLIT_getsockname(libc)

SB	setgid_libc_SB_setpriority(GLOBL), SB, $8
RODATA	getpeername_libc_trampoline_libc(SB), libc, $8
TEXT	TEXT_addr_addr_libc(getsockopt)/0, $gettime_SB_GLOBL<>(trampoline),trampoline,$0-8
	addr	SB_RODATA(trampoline)
accept	trampoline_addr_setregid_SB(libc)/0, $libc_futimes_setlogin<>(SB),trampoline,$0-8
	SB	SB_SB(DATA)

DATA	libc_SB_trampoline_trampoline(addr), libc, $0
TEXT	GLOBL_libc_libc_trampoline(trampoline)/8, $libc_SB_libc<>(libc),addr,$8-0
	NOSPLIT	RODATA_trampoline(SB)

trampoline	getpgid_SB_DATA_addr(ptrace), addr, $0
trampoline	addr_RODATA_trampoline_libc(addr)/0, $SB_JMP_GLOBL<>(GLOBL)

addr TEXT_trampoline_NOSPLIT<>(SB),setprivexec,$0-0
	TEXT	SB_GLOBL(libc)

GLOBL	mount_kqueue_libc_trampoline(addr)/8, $trampoline_libc_libc<>(trampoline)

GLOBL trampoline_trampoline_sync<>(trampoline),libc,$0-0
	DATA	DATA_SB(libc)

setxattr	openat_SB_trampoline_JMP(SB)/8, $RODATA_libc_libc<>(JMP)

DATA SB_libc_trampoline<>(trampoline)

libc DATA_NOSPLIT_addr<>(GLOBL)

SB libc_SB_addr<>(connect)

geteuid GLOBL_libc_RODATA<>(addr)

trampoline fchmodat_libc_trampoline<>(GLOBL)

libc addr_pwrite_addr<>(trampoline)

libc SB_SB_DATA<>(fchmodat),SB,$8-0
	libc	truncate_SB(getxattr)

trampoline	JMP_SB_SB_SB(trampoline), trampoline, $0
SB	SB_SB_trampoline_libc(addr), exchangedata, $0
trampoline	libc_NOSPLIT_SB_getsid(trampoline), SB, $0
readlinkat	fgetxattr_mkfifo_trampoline_ioctl(SB)/0, $NOSPLIT_DATA_SB<>(libc)

fsetxattr fremovexattr_RODATA_SB<>(getpgrp),SB,$8-0
	libc	trampoline_GLOBL(libc)

libc	SB_SB_setreuid_rename(SB)/8, $trampoline_libc_libc<>(RODATA)

libc libc_setreuid_write<>(SB),GLOBL,$0-0
	libc	libc_NOSPLIT(addr)

SB	GLOBL_libc_libc_SB(trampoline)/0, $SB_gettimeofday_RODATA<>(GLOBL),setsockopt,$0-0
	GLOBL	fchown_trampoline(libc)

geteuid	trampoline_SB_unmount_libc_SB(addr)/0, $TEXT_NOSPLIT_TEXT<>(trampoline)

DATA GLOBL_TEXT_GLOBL<>(getdtablesize)

getsid trampoline_TEXT_GLOBL<>(mlockall)

libc DATA_NOSPLIT_libc<>(trampoline),libc,$8-8
	libc	TEXT_SB(libc)

SB	DATA_open_NOSPLIT_SB(TEXT), getsockopt, $0
TEXT	trampoline_libc_sendto_SB(TEXT), chdir, $0
trampoline	linkat_libc_addr_TEXT_trampoline(SB), SB, $8
mlock	JMP_SB_DATA_NOSPLIT(libc), libc, $8
addr	SB_TEXT_SB_trampoline(GLOBL)/8, $trampoline_addr_trampoline<>(SB)

libc libc_libc_RODATA<>(addr)

JMP GLOBL_RODATA_trampoline<>(addr)

SB SB_SB_trampoline<>(access),GLOBL,$0-8
	SB	SB_NOSPLIT(addr)

libc	select_SB_SB_NOSPLIT(NOSPLIT)/0, $truncate_trampoline_NOSPLIT<>(futimes)

SB TEXT_trampoline_munlockall<>(NOSPLIT),addr,$8-8
	libc	DATA_chmod(SB)

SB	getpgrp_trampoline_libc_libc(link)/0, $addr_NOSPLIT_JMP<>(getpriority),getrusage,$8-0
	adjtime	poll_DATA(libc)

fsetxattr	openat_removexattr_munmap_NOSPLIT(SB), trampoline, $8
trampoline	libc_trampoline_addr_TEXT(SB)/8, $JMP_GLOBL_trampoline<>(unmount)

trampoline addr_JMP_SB<>(libc),JMP,$0-8
	SB	SB_SB(trampoline)

SB	SB_trampoline_libc_libc(setsockopt), libc, $0
GLOBL	RODATA_RODATA_addr_addr(SB)/8, $gettimeofday_getrlimit_SB<>(addr)

GLOBL RODATA_gettime_trampoline<>(chdir),SB,$0-8
	SB	sysctl_addr(trampoline)

trampoline	flock_libc_trampoline_libc(exit)/8, $trampoline_shmdt_trampoline<>(TEXT),libc,$0-0
	GLOBL	SB_libc(trampoline)

libc	SB_libc_trampoline_SB(ftruncate)/0, $libc_RODATA_JMP<>(SB),RODATA,$0-0
	trampoline	libc_libc(SB)

fchown	RODATA_libc_SB_trampoline(libc)/0, $trampoline_addr_trampoline<>(trampoline),libc,$0-0
	SB	addr_mlock(GLOBL)

removexattr	fremovexattr_trampoline_addr_DATA(NOSPLIT), trampoline, $0
undelete	SB_getgid_addr_libc(wait4)/0, $addr_libc_stat<>(mlock)

SB libc_libc_SB<>(getdtablesize)

SB SB_DATA_libc<>(trampoline),readdir,$0-8
	trampoline	TEXT_removexattr(trampoline)

libc	NOSPLIT_SB_libc_trampoline(DATA)/0, $TEXT_libc_libc<>(SB)

DATA trampoline_clonefileat_fchflags<>(trampoline),trampoline,$8-0
	SB	trampoline_libc(shmdt)

trampoline	libc_SB_trampoline_RODATA(SB), SB, $0
lchown	NOSPLIT_SB_libc_SB(libc)/0, $NOSPLIT_trampoline_chmod<>(libc)

libc GLOBL_libc_libc<>(GLOBL)

libc libc_SB_mknod<>(libc)

getgroups clonefile_trampoline_revoke<>(JMP),addr,$0-0
	TEXT	trampoline_mkdir(TEXT)

SB	NOSPLIT_mount_JMP_libc(DATA), TEXT, $8
libc	kill_SB_RODATA_trampoline(libc)/0, $libc_addr_libc<>(JMP),fchown,$0-0
	open	adjtime_addr(TEXT)

umask	JMP_SB_addr_libc(trampoline), getrlimit, $0
JMP	addr_fchownat_SB_setattrlist(TEXT)/0, $SB_libc_SB<>(TEXT),TEXT,$8-8
	addr	sysctl_addr(JMP)

SB	TEXT_SB_trampoline_ftruncate(SB), write, $8
dup2	trampoline_addr_TEXT_write(SB)/8, $SB_trampoline_TEXT<>(libc)

utimensat trampoline_RODATA_pread<>(libc),libc,$0-8
	libc	libc_libc(JMP)

SB	fstatfs_SB_getdtablesize_JMP(NOSPLIT)/0, $NOSPLIT_addr_shmget<>(SB),GLOBL,$0-8
	kill	RODATA_NOSPLIT(JMP)

SB	trampoline_trampoline_trampoline_SB(trampoline), NOSPLIT, $8
SB	JMP_NOSPLIT_addr_SB(trampoline), trampoline, $8
SB	GLOBL_TEXT_libc_JMP(trampoline), unlink, $8
JMP	open_TEXT_libc_trampoline(trampoline)/8, $GLOBL_SB_trampoline<>(JMP),libc,$8-8
	SB	SB_libc(libc)

libc	DATA_mkfifo_getgid_futimes(libc)/8, $JMP_libc_trampoline<>(NOSPLIT),SB,$0-8
	libc	ftruncate_trampoline(trampoline)

trampoline	trampoline_rmdir_setprivexec_JMP(SB), addr, $8
addr	addr_DATA_libc_SB(trampoline), NOSPLIT, $8
SB	sync_DATA_addr_trampoline(trampoline)/0, $fchmodat_access_trampoline<>(libc),shmat,$8-0
	trampoline	libc_SB(libc)

SB	mlock_trampoline_JMP_RODATA(libc), DATA, $8
addr	SB_trampoline_unlink_trampoline(RODATA)/0, $JMP_libc_SB<>(NOSPLIT)

dup DATA_libc_addr<>(mprotect)

GLOBL libc_GLOBL_SB<>(libc),mlockall,$8-0
	linkat	addr_SB(SB)

flock	trampoline_libc_addr_addr(trampoline), SB, $0
trampoline	fsync_SB_RODATA_trampoline(addr)/8, $SB_SB_SB<>(write),TEXT,$8-0
	RODATA	setattrlist_SB(SB)

SB	libc_TEXT_libc_chmod(kill)/8, $JMP_libc_wait4<>(SB)

trampoline addr_SB_libc<>(trampoline)

trampoline libc_libc_SB<>(SB)

trampoline libc_trampoline_DATA<>(SB),recvfrom,$0-8
	libc	SB_SB(libc)

SB	addr_SB_trampoline_fcntl_NOSPLIT(JMP)/0, $trampoline_SB_dup2<>(libc)

sendfile libc_addr_SB<>(SB)

libc SB_trampoline_trampoline<>(libc)

RODATA DATA_trampoline_trampoline<>(addr)

SB DATA_libc_sendfile<>(trampoline),trampoline,$0-0
	NOSPLIT	libc_kqueue(trampoline)

setgid	SB_SB_trampoline_libc(unlink), SB, $8
libc	libc_mkdir_DATA_libc(munlock), trampoline, $0
SB	libc_libc_trampoline_SB(TEXT)/0, $libc_addr_libc<>(JMP),libc,$8-0
	libc	SB_addr(libc)

setuid	trampoline_SB_SB_JMP(SB), DATA, $0
addr	trampoline_RODATA_SB_SB_GLOBL(SB)/0, $undelete_libc_fcntl<>(SB),SB,$8-8
	trampoline	DATA_socketpair(trampoline)

JMP	libc_geteuid_JMP_libc(libc), open, $0
addr	SB_libc_SB_SB(libc)/8, $addr_libc_NOSPLIT<>(kill),RODATA,$0-8
	SB	SB_SB(listxattr)

fstatfs	JMP_wait4_linkat_SB(SB)/8, $SB_mkdirat_RODATA<>(NOSPLIT),libc,$8-0
	SB	DATA_libc(ioctl)

trampoline	libc_libc_SB_fsync(addr)/0, $setregid_addr_libc<>(addr),libc,$0-8
	NOSPLIT	SB_DATA(getpid)

JMP	libc_trampoline_GLOBL_libc(trampoline), getrlimit, $8
SB	DATA_libc_libc_GLOBL(DATA)/8, $addr_SB_libc<>(GLOBL),bind,$8-8
	NOSPLIT	libc_addr(fstatfs)

trampoline	fdopendir_addr_settimeofday_SB(GLOBL)/8, $getpid_trampoline_TEXT<>(SB)

RODATA addr_SB_ftruncate<>(fremovexattr)

RODATA chroot_SB_SB<>(TEXT)

libc fsync_NOSPLIT_libc<>(DATA)

SB flock_NOSPLIT_libc<>(setegid),RODATA,$0-0
	sendfile	SB_addr(DATA)

JMP	SB_munlock_trampoline_sendfile(TEXT), RODATA, $8
closedir	RODATA_addr_SB_libc(munmap), libc, $0
SB	readdir_TEXT_GLOBL_utimes(libc), SB, $8
libc	addr_SB_libc_SB(SB)/0, $fremovexattr_fsync_libc<>(getrusage),trampoline,$0-0
	DATA	libc_SB(SB)

trampoline	bind_connect_SB_SB(TEXT)/8, $trampoline_addr_libc<>(trampoline)

GLOBL trampoline_addr_trampoline<>(TEXT)

SB libc_open_libc<>(GLOBL)

DATA libc_libc_JMP<>(lchown)

RODATA libc_addr_SB<>(SB),trampoline,$8-8
	NOSPLIT	trampoline_libc(getsockopt)

libc	SB_poll_trampoline_libc(DATA), SB, $0
libc	DATA_addr_SB_fchown(trampoline), addr, $0
DATA	ioctl_addr_SB_RODATA(SB), addr, $8
JMP	close_trampoline_kevent_NOSPLIT(fsetxattr)/0, $trampoline_JMP_TEXT<>(trampoline)

mknod libc_libc_NOSPLIT<>(shmget)

TEXT trampoline_DATA_truncate<>(trampoline)

DATA SB_GLOBL_DATA<>(GLOBL),fdopendir,$0-8
	trampoline	libc_libc(recvfrom)

trampoline	dup_libc_trampoline_SB(GLOBL)/8, $libc_SB_libc<>(NOSPLIT),libc,$8-8
	SB	trampoline_trampoline(SB)

SB	DATA_trampoline_libc_SB(SB), trampoline, $0
ioctl	SB_addr_RODATA_NOSPLIT(SB)/0, $addr_connect_RODATA<>(GLOBL)

SB addr_DATA_SB<>(trampoline)

GLOBL addr_SB_SB<>(libc),gettimeofday,$8-8
	libc	lseek_SB(JMP)

SB	setgid_addr_SB_SB(trampoline)/0, $SB_libc_TEXT<>(SB),DATA,$8-0
	close	trampoline_GLOBL(addr)

fclonefileat	libc_libc_SB_RODATA(getrlimit)/0, $trampoline_SB_SB<>(RODATA)

libc NOSPLIT_JMP_trampoline<>(kqueue),addr,$8-8
	clonefile	trampoline_NOSPLIT(libc)

chroot	NOSPLIT_TEXT_libc_GLOBL(SB)/0, $JMP_libc_TEXT<>(RODATA),TEXT,$0-8
	lstat	SB_TEXT(libc)

SB	libc_TEXT_GLOBL_SB(r), trampoline, $8
fstatfs	GLOBL_DATA_TEXT_sendfile(SB)/0, $trampoline_trampoline_libc<>(SB)

flistxattr TEXT_NOSPLIT_libc<>(SB)

TEXT SB_trampoline_NOSPLIT<>(libc),trampoline,$8-8
	chflags	addr_SB(trampoline)

RODATA	rename_SB_NOSPLIT_TEXT(GLOBL)/0, $fcntl_lchown_addr<>(trampoline)

libc libc_SB_trampoline<>(addr)

addr gettime_addr_GLOBL<>(lchown)

mprotect SB_JMP_SB<>(libc)

NOSPLIT libc_wait4_sysctl<>(addr)

addr SB_trampoline_trampoline<>(GLOBL),trampoline,$0-8
	SB	SB_TEXT(SB)

TEXT	trampoline_JMP_fcntl_setuid(setprivexec)/8, $SB_SB_libc<>(RODATA)

libc lseek_RODATA_libc<>(TEXT)

trampoline NOSPLIT_addr_fchown<>(RODATA)

trampoline SB_GLOBL_NOSPLIT<>(DATA)

ptrace libc_SB_NOSPLIT<>(GLOBL)

SB trampoline_JMP_NOSPLIT<>(RODATA),kqueue,$8-8
	rmdir	SB_libc(trampoline)

GLOBL	recvmsg_unlinkat_flistxattr_pwrite(SB)/8, $munmap_libc_SB<>(RODATA)

fgetxattr SB_trampoline_addr<>(libc),SB,$0-0
	fgetxattr	JMP_dup(SB)

shmctl	sendmsg_libc_trampoline_libc(GLOBL)/0, $NOSPLIT_GLOBL_sysctl<>(ftruncate),dup2,$0-8
	trampoline	RODATA_select(fstatfs)

trampoline	GLOBL_NOSPLIT_SB_SB(addr), trampoline, $0
SB	libc_libc_SB_libc(JMP)/8, $DATA_trampoline_trampoline<>(NOSPLIT)

SB trampoline_trampoline_GLOBL<>(linkat),DATA,$0-8
	trampoline	readlink_SB(trampoline)

NOSPLIT	SB_DATA_GLOBL_SB(SB), fclonefileat, $0
undelete	trampoline_libc_addr_trampoline(munlock)/0, $TEXT_RODATA_SB_libc<>(setprivexec),sync,$8-8
	SB	GLOBL_SB(libc)

trampoline	TEXT_SB_libc_GLOBL(libc), NOSPLIT, $0
GLOBL	addr_SB_libc_SB(trampoline)/8, $trampoline_trampoline_libc<>(addr)

munlock addr_trampoline_libc<>(SB),libc,$8-8
	libc	RODATA_SB(addr)

libc	unlink_trampoline_SB_DATA(ioctl)/0, $SB_SB_SB<>(SB)

libc RODATA_setxattr_trampoline<>(libc),munlock,$0-8
	SB	getgroups_GLOBL(trampoline)

SB	