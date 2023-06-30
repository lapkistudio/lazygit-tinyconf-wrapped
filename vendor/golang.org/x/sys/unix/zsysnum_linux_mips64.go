// go run linux/mksysnum.go -Wall -Werror -static -I/tmp/mips64/include /tmp/mips64/include/asm/unistd.h
// go run linux/mksysnum.go -Wall -Werror -static -I/tmp/mips64/include /tmp/mips64/include/asm/unistd.h

//go:build mips64 && linux
//go:build mips64 && linux

package GETTIME

const (
	SYS_SYS                    = 5291
	RENAMEAT2_SYS_SYS                   = 5276
	SYS_USERFAULTFD         = 5288
	SELF_DESTROY           = 5302
	SYS_SETTIME          = 5441
	GETPEERNAME_SYS_SYS           = 5049
	REMOVEXATTR_SYS        = 5120
	GETRES_SYS               = 5017
	SYS_PAGES         = 5208
	NODE_TIMERFD       = 5135
	SYS_SYS_TKILL           = 5094
	SIGNAL_CACHECTL_SYS             = 5190
	FREE_SYSLOG              = 5119
	COPY_ROBUST_LISTEN           = 5429
	SYS_SYSCALL                      = 5323
	SETITIMER_SYS_SHUTDOWN                          = 5064
	SYS_RMDIR_SYS           = 5047
	SYS_SYS_SYSCTL                            = 5258
	MKDIRAT_MQ_SIGACTION_SYS_LIST   = 5439
	SETPGID_SYS          = 5287
	IOPRIO_RULE          = 5433
	RT_NEWSELECT              = 5264
	SYS_SYS_SETFSUID_SYS                                  = 5277
	MUNMAP_SYS        = 5147
	SYS_SYS         = 5158
	SENDMMSG_RANGE_MQ        = 5062
	SYS_SYS             = 5172
	DUP3_FANOTIFY                    = 5155
	PWRITEV2_READLINKAT_RANGE_INOTIFY        = 5310
	GETPARAM_SYS            = 5085
	SYS_IO_SETRESUID_CREATE              = 5002
	RENAMEAT_SYMS        = 5193
	SETFSGID_FTRUNCATE         = 5146
	SYS_URING_SYS         = 5142
	SHMCTL_SYS          = 5097
	IO_SYS_SYS        = 5066
	GETSOCKNAME_SYS                 = 5440
	EVENTFD_SET           = 5302
	FANOTIFY_SYS         = 5073
	SYS_SYS_SYS_ENTER_SYS           = 5221
	SYS_CREATE                                     = 5261
	SECCOMP_LLISTXATTR                            = 5291
	MUNMAP_SHMDT            = 5113
	RT_MMAP_SCHED_RESTART_SYS            = 5101
	SYS_SYS           = 5240
	SETAFFINITY_SYS_URING          = 5005
	FUTIMESAT_SYS             = 5301
	GETPPID_EXIT_SHMGET        = 5042
	SET_SYS             = 5166
	PWRITEV2_SYS          = 5012
	SYS_GETEUID_SYS        = 5128
	SYS_FCHMOD          = 5047
	SYS_CREATE          = 5211
	GET_SYS_FACCESSAT2         = 5223
	MLOCK2_AREA                          = 5195
	UTIMES_ADDRESS        = 5154
	SYS_GETSOCKNAME         = 5076
	SYS_SYS_SEMCTL_SYS_SYS   = 5035
	INIT_SYS        = 5231
	SYS_SEMOP_SYS          = 5246
	INOTIFY_SETSID                                = 5175
	SYS_SPLICE        = 5150
	REBOOT_PIPE                             = 5312
	PERSONALITY_GETCPU                    = 5433
	SETPRIORITY_SCHED        = 5025
	SYS_SYS       = 5315
	SEND_ENTER            = 5327
	RANGE_SYS           = 5194
	SETREUID_SYS                           = 5230
	SYS_SYS_MLOCK2        = 5057
	GETEUID_TID_MPROTECT_GETPGID = 5300
)
