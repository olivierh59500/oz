// DO NOT EDIT. Autogenerated by syscalls_gen_args.go

package seccomp

type SystemCall struct {
	prefix      string
	name        string
	num         int
	args        SystemCallArgs
	captureArgs SystemCallArgs
}

var syscalls = []SystemCall{

	SystemCall{
		name: "read",
		num:  0,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "write",
		num:  1,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "open",
		num:  2,
		args: SystemCallArgs{STRINGARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "close",
		num:  3,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "stat",
		num:  4,
		args: SystemCallArgs{STRINGARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "fstat",
		num:  5,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "lstat",
		num:  6,
		args: SystemCallArgs{STRINGARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "poll",
		num:  7,
		args: SystemCallArgs{PTRARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "lseek",
		num:  8,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "mmap",
		num:  9,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, INTARG, INTARG},
		captureArgs: SystemCallArgs{0, 0, STRINGARG, STRINGARG, 0, 0},
	},
	SystemCall{
		name: "mprotect",
		num:  10,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
		captureArgs: SystemCallArgs{0, 0, STRINGARG, 0, 0, 0},
	},
	SystemCall{
		name: "munmap",
		num:  11,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "brk",
		num:  12,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "rt_sigaction",
		num:  13,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "rt_sigprocmask",
		num:  14,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "rt_sigreturn",
		num:  15,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name:        "ioctl",
		num:         16,
		args:        SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
		captureArgs: SystemCallArgs{0, STRINGARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "pread64",
		num:  17,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "pwrite64",
		num:  18,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "readv",
		num:  19,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "writev",
		num:  20,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "access",
		num:  21,
		args: SystemCallArgs{STRINGARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "pipe",
		num:  22,
		args: SystemCallArgs{PTRARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "select",
		num:  23,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, PTRARG, PTRARG, 0},
	},
	SystemCall{
		name: "sched_yield",
		num:  24,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "mremap",
		num:  25,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "msync",
		num:  26,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "mincore",
		num:  27,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "madvise",
		num:  28,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "shmget",
		num:  29,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "shmat",
		num:  30,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "shmctl",
		num:  31,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "dup",
		num:  32,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "dup2",
		num:  33,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "pause",
		num:  34,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "nanosleep",
		num:  35,
		args: SystemCallArgs{PTRARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getitimer",
		num:  36,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "alarm",
		num:  37,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setitimer",
		num:  38,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "getpid",
		num:  39,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sendfile",
		num:  40,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, INTARG, 0, 0},
	},
	SystemCall{
		name:        "socket",
		num:         41,
		args:        SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
		captureArgs: SystemCallArgs{STRINGARG, STRINGARG, STRINGARG, 0, 0, 0},
	},
	SystemCall{
		name: "connect",
		num:  42,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "accept",
		num:  43,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "sendto",
		num:  44,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, PTRARG, INTARG},
	},
	SystemCall{
		name: "recvfrom",
		num:  45,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, PTRARG, PTRARG},
	},
	SystemCall{
		name: "sendmsg",
		num:  46,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "recvmsg",
		num:  47,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "shutdown",
		num:  48,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "bind",
		num:  49,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "listen",
		num:  50,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getsockname",
		num:  51,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "getpeername",
		num:  52,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "socketpair",
		num:  53,
		args: SystemCallArgs{INTARG, INTARG, INTARG, PTRARG, 0, 0},
	},
	SystemCall{
		name:        "setsockopt",
		num:         54,
		args:        SystemCallArgs{INTARG, INTARG, INTARG, PTRARG, INTARG, 0},
		captureArgs: SystemCallArgs{0, STRINGARG, STRINGARG, 0, 0, 0},
	},
	SystemCall{
		name: "getsockopt",
		num:  55,
		args: SystemCallArgs{INTARG, INTARG, INTARG, PTRARG, PTRARG, 0},
	},
	SystemCall{
		name: "clone",
		num:  56,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, PTRARG, INTARG, INTARG},
	},
	SystemCall{
		name: "fork",
		num:  57,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "vfork",
		num:  58,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "execve",
		num:  59,
		args: SystemCallArgs{STRINGARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "exit",
		num:  60,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "wait4",
		num:  61,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "kill",
		num:  62,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "uname",
		num:  63,
		args: SystemCallArgs{PTRARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "semget",
		num:  64,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "semop",
		num:  65,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "semctl",
		num:  66,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "shmdt",
		num:  67,
		args: SystemCallArgs{STRINGARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "msgget",
		num:  68,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "msgsnd",
		num:  69,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "msgrcv",
		num:  70,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "msgctl",
		num:  71,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name:        "fcntl",
		num:         72,
		args:        SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
		captureArgs: SystemCallArgs{0, STRINGARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "flock",
		num:  73,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "fsync",
		num:  74,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "fdatasync",
		num:  75,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "truncate",
		num:  76,
		args: SystemCallArgs{STRINGARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "ftruncate",
		num:  77,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getdents",
		num:  78,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "getcwd",
		num:  79,
		args: SystemCallArgs{PTRARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "chdir",
		num:  80,
		args: SystemCallArgs{STRINGARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "fchdir",
		num:  81,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "rename",
		num:  82,
		args: SystemCallArgs{STRINGARG, STRINGARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "mkdir",
		num:  83,
		args: SystemCallArgs{STRINGARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "rmdir",
		num:  84,
		args: SystemCallArgs{STRINGARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "creat",
		num:  85,
		args: SystemCallArgs{STRINGARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "link",
		num:  86,
		args: SystemCallArgs{STRINGARG, STRINGARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "unlink",
		num:  87,
		args: SystemCallArgs{STRINGARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "symlink",
		num:  88,
		args: SystemCallArgs{STRINGARG, STRINGARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "readlink",
		num:  89,
		args: SystemCallArgs{STRINGARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "chmod",
		num:  90,
		args: SystemCallArgs{STRINGARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "fchmod",
		num:  91,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "chown",
		num:  92,
		args: SystemCallArgs{STRINGARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "fchown",
		num:  93,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "lchown",
		num:  94,
		args: SystemCallArgs{STRINGARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "umask",
		num:  95,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "gettimeofday",
		num:  96,
		args: SystemCallArgs{PTRARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getrlimit",
		num:  97,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getrusage",
		num:  98,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sysinfo",
		num:  99,
		args: SystemCallArgs{PTRARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "times",
		num:  100,
		args: SystemCallArgs{PTRARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "ptrace",
		num:  101,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "getuid",
		num:  102,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "syslog",
		num:  103,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "getgid",
		num:  104,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setuid",
		num:  105,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setgid",
		num:  106,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "geteuid",
		num:  107,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getegid",
		num:  108,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setpgid",
		num:  109,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getppid",
		num:  110,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getpgrp",
		num:  111,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setsid",
		num:  112,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setreuid",
		num:  113,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setregid",
		num:  114,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getgroups",
		num:  115,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setgroups",
		num:  116,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setresuid",
		num:  117,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "getresuid",
		num:  118,
		args: SystemCallArgs{PTRARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "setresgid",
		num:  119,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "getresgid",
		num:  120,
		args: SystemCallArgs{PTRARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "getpgid",
		num:  121,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setfsuid",
		num:  122,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setfsgid",
		num:  123,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getsid",
		num:  124,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "capget",
		num:  125,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "capset",
		num:  126,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "rt_sigpending",
		num:  127,
		args: SystemCallArgs{PTRARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "rt_sigtimedwait",
		num:  128,
		args: SystemCallArgs{PTRARG, PTRARG, PTRARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "rt_sigqueueinfo",
		num:  129,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "rt_sigsuspend",
		num:  130,
		args: SystemCallArgs{PTRARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sigaltstack",
		num:  131,
		args: SystemCallArgs{PTRARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "utime",
		num:  132,
		args: SystemCallArgs{STRINGARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "mknod",
		num:  133,
		args: SystemCallArgs{STRINGARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "uselib",
		num:  134,
		args: SystemCallArgs{STRINGARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "personality",
		num:  135,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "ustat",
		num:  136,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "statfs",
		num:  137,
		args: SystemCallArgs{STRINGARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "fstatfs",
		num:  138,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sysfs",
		num:  139,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "getpriority",
		num:  140,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setpriority",
		num:  141,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "sched_setparam",
		num:  142,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sched_getparam",
		num:  143,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sched_setscheduler",
		num:  144,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "sched_getscheduler",
		num:  145,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sched_get_priority_max",
		num:  146,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sched_get_priority_min",
		num:  147,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sched_rr_get_interval",
		num:  148,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "mlock",
		num:  149,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "munlock",
		num:  150,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "mlockall",
		num:  151,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "munlockall",
		num:  152,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "vhangup",
		num:  153,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "modify_ldt",
		num:  154,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "pivot_root",
		num:  155,
		args: SystemCallArgs{STRINGARG, STRINGARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "_sysctl",
		num:  156,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name:        "prctl",
		num:         157,
		args:        SystemCallArgs{INTARG, INTARG, INTARG, INTARG, INTARG, 0},
		captureArgs: SystemCallArgs{STRINGARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "arch_prctl",
		num:  158,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "adjtimex",
		num:  159,
		args: SystemCallArgs{PTRARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setrlimit",
		num:  160,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "chroot",
		num:  161,
		args: SystemCallArgs{STRINGARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sync",
		num:  162,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "acct",
		num:  163,
		args: SystemCallArgs{STRINGARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "settimeofday",
		num:  164,
		args: SystemCallArgs{PTRARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "mount",
		num:  165,
		args: SystemCallArgs{STRINGARG, PTRARG, STRINGARG, INTARG, PTRARG, 0},
	},
	SystemCall{
		name: "umount2",
		num:  166,
		args: SystemCallArgs{STRINGARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "swapon",
		num:  167,
		args: SystemCallArgs{STRINGARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "swapoff",
		num:  168,
		args: SystemCallArgs{STRINGARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "reboot",
		num:  169,
		args: SystemCallArgs{INTARG, INTARG, INTARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "sethostname",
		num:  170,
		args: SystemCallArgs{STRINGARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "setdomainname",
		num:  171,
		args: SystemCallArgs{STRINGARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "iopl",
		num:  172,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "ioperm",
		num:  173,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "create_module",
		num:  174,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "init_module",
		num:  175,
		args: SystemCallArgs{PTRARG, INTARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "delete_module",
		num:  176,
		args: SystemCallArgs{STRINGARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "get_kernel_syms",
		num:  177,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "query_module",
		num:  178,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "quotactl",
		num:  179,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "nfsservctl",
		num:  180,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getpmsg",
		num:  181,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "putpmsg",
		num:  182,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "afs_syscall",
		num:  183,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "tuxcall",
		num:  184,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "security",
		num:  185,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "gettid",
		num:  186,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "readahead",
		num:  187,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "setxattr",
		num:  188,
		args: SystemCallArgs{STRINGARG, STRINGARG, PTRARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "lsetxattr",
		num:  189,
		args: SystemCallArgs{STRINGARG, STRINGARG, PTRARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "fsetxattr",
		num:  190,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "getxattr",
		num:  191,
		args: SystemCallArgs{STRINGARG, STRINGARG, PTRARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "lgetxattr",
		num:  192,
		args: SystemCallArgs{STRINGARG, STRINGARG, PTRARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "fgetxattr",
		num:  193,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "listxattr",
		num:  194,
		args: SystemCallArgs{STRINGARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "llistxattr",
		num:  195,
		args: SystemCallArgs{STRINGARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "flistxattr",
		num:  196,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "removexattr",
		num:  197,
		args: SystemCallArgs{STRINGARG, STRINGARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "lremovexattr",
		num:  198,
		args: SystemCallArgs{STRINGARG, STRINGARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "fremovexattr",
		num:  199,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "tkill",
		num:  200,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "time",
		num:  201,
		args: SystemCallArgs{PTRARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name:        "futex",
		num:         202,
		args:        SystemCallArgs{PTRARG, INTARG, INTARG, PTRARG, PTRARG, INTARG},
		captureArgs: SystemCallArgs{0, STRINGARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sched_setaffinity",
		num:  203,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "sched_getaffinity",
		num:  204,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "set_thread_area",
		num:  205,
		args: SystemCallArgs{PTRARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "io_setup",
		num:  206,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "io_destroy",
		num:  207,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "io_getevents",
		num:  208,
		args: SystemCallArgs{INTARG, INTARG, INTARG, PTRARG, PTRARG, 0},
	},
	SystemCall{
		name: "io_submit",
		num:  209,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "io_cancel",
		num:  210,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "get_thread_area",
		num:  211,
		args: SystemCallArgs{PTRARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "lookup_dcookie",
		num:  212,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "epoll_create",
		num:  213,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "epoll_ctl_old",
		num:  214,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "epoll_wait_old",
		num:  215,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "remap_file_pages",
		num:  216,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "getdents64",
		num:  217,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "set_tid_address",
		num:  218,
		args: SystemCallArgs{PTRARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "restart_syscall",
		num:  219,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "semtimedop",
		num:  220,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "fadvise64",
		num:  221,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "timer_create",
		num:  222,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "timer_settime",
		num:  223,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "timer_gettime",
		num:  224,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "timer_getoverrun",
		num:  225,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "timer_delete",
		num:  226,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "clock_settime",
		num:  227,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "clock_gettime",
		num:  228,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "clock_getres",
		num:  229,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "clock_nanosleep",
		num:  230,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "exit_group",
		num:  231,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "epoll_wait",
		num:  232,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "epoll_ctl",
		num:  233,
		args: SystemCallArgs{INTARG, INTARG, INTARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "tgkill",
		num:  234,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "utimes",
		num:  235,
		args: SystemCallArgs{STRINGARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "vserver",
		num:  236,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "mbind",
		num:  237,
		args: SystemCallArgs{INTARG, INTARG, INTARG, PTRARG, INTARG, INTARG},
	},
	SystemCall{
		name: "set_mempolicy",
		num:  238,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "get_mempolicy",
		num:  239,
		args: SystemCallArgs{PTRARG, PTRARG, INTARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "mq_open",
		num:  240,
		args: SystemCallArgs{STRINGARG, INTARG, INTARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "mq_unlink",
		num:  241,
		args: SystemCallArgs{STRINGARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "mq_timedsend",
		num:  242,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, PTRARG, 0},
	},
	SystemCall{
		name: "mq_timedreceive",
		num:  243,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, PTRARG, PTRARG, 0},
	},
	SystemCall{
		name: "mq_notify",
		num:  244,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "mq_getsetattr",
		num:  245,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "kexec_load",
		num:  246,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "waitid",
		num:  247,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, INTARG, PTRARG, 0},
	},
	SystemCall{
		name: "add_key",
		num:  248,
		args: SystemCallArgs{STRINGARG, STRINGARG, PTRARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "request_key",
		num:  249,
		args: SystemCallArgs{STRINGARG, STRINGARG, STRINGARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "keyctl",
		num:  250,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "ioprio_set",
		num:  251,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "ioprio_get",
		num:  252,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "inotify_init",
		num:  253,
		args: SystemCallArgs{0, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "inotify_add_watch",
		num:  254,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "inotify_rm_watch",
		num:  255,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "migrate_pages",
		num:  256,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "openat",
		num:  257,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "mkdirat",
		num:  258,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "mknodat",
		num:  259,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "fchownat",
		num:  260,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "futimesat",
		num:  261,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "newfstatat",
		num:  262,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "unlinkat",
		num:  263,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "renameat",
		num:  264,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "linkat",
		num:  265,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, PTRARG, INTARG, 0},
	},
	SystemCall{
		name: "symlinkat",
		num:  266,
		args: SystemCallArgs{STRINGARG, INTARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "readlinkat",
		num:  267,
		args: SystemCallArgs{INTARG, PTRARG, STRINGARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "fchmodat",
		num:  268,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "faccessat",
		num:  269,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "pselect6",
		num:  270,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, PTRARG, PTRARG, PTRARG},
	},
	SystemCall{
		name: "ppoll",
		num:  271,
		args: SystemCallArgs{PTRARG, INTARG, PTRARG, PTRARG, INTARG, 0},
	},
	SystemCall{
		name: "unshare",
		num:  272,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "set_robust_list",
		num:  273,
		args: SystemCallArgs{PTRARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "get_robust_list",
		num:  274,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "splice",
		num:  275,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, PTRARG, INTARG, INTARG},
	},
	SystemCall{
		name: "tee",
		num:  276,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "sync_file_range",
		num:  277,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "vmsplice",
		num:  278,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "move_pages",
		num:  279,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, PTRARG, PTRARG, INTARG},
	},
	SystemCall{
		name: "utimensat",
		num:  280,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "epoll_pwait",
		num:  281,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, PTRARG, INTARG},
	},
	SystemCall{
		name: "signalfd",
		num:  282,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "timerfd_create",
		num:  283,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "eventfd",
		num:  284,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "fallocate",
		num:  285,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "timerfd_settime",
		num:  286,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "timerfd_gettime",
		num:  287,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "accept4",
		num:  288,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "signalfd4",
		num:  289,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "eventfd2",
		num:  290,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "epoll_create1",
		num:  291,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "dup3",
		num:  292,
		args: SystemCallArgs{INTARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "pipe2",
		num:  293,
		args: SystemCallArgs{PTRARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "inotify_init1",
		num:  294,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "preadv",
		num:  295,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "pwritev",
		num:  296,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "rt_tgsigqueueinfo",
		num:  297,
		args: SystemCallArgs{INTARG, INTARG, INTARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "perf_event_open",
		num:  298,
		args: SystemCallArgs{PTRARG, INTARG, INTARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "recvmmsg",
		num:  299,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, PTRARG, 0},
	},
	SystemCall{
		name: "fanotify_init",
		num:  300,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "fanotify_mark",
		num:  301,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, STRINGARG, 0},
	},
	SystemCall{
		name: "prlimit64",
		num:  302,
		args: SystemCallArgs{INTARG, INTARG, PTRARG, PTRARG, 0, 0},
	},
	SystemCall{
		name: "name_to_handle_at",
		num:  303,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, PTRARG, INTARG, 0},
	},
	SystemCall{
		name: "open_by_handle_at",
		num:  304,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "clock_adjtime",
		num:  305,
		args: SystemCallArgs{INTARG, PTRARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "syncfs",
		num:  306,
		args: SystemCallArgs{INTARG, 0, 0, 0, 0, 0},
	},
	SystemCall{
		name: "sendmmsg",
		num:  307,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "setns",
		num:  308,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getcpu",
		num:  309,
		args: SystemCallArgs{PTRARG, PTRARG, PTRARG, 0, 0, 0},
	},
	SystemCall{
		name: "process_vm_readv",
		num:  310,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, PTRARG, INTARG, INTARG},
	},
	SystemCall{
		name: "process_vm_writev",
		num:  311,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, PTRARG, INTARG, INTARG},
	},
	SystemCall{
		name: "kcmp",
		num:  312,
		args: SystemCallArgs{INTARG, INTARG, INTARG, INTARG, INTARG, 0},
	},
	SystemCall{
		name: "finit_module",
		num:  313,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "sched_setattr",
		num:  314,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "sched_getattr",
		num:  315,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, INTARG, 0, 0},
	},
	SystemCall{
		name: "renameat2",
		num:  316,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, PTRARG, INTARG, 0},
	},
	SystemCall{
		name: "seccomp",
		num:  317,
		args: SystemCallArgs{INTARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "getrandom",
		num:  318,
		args: SystemCallArgs{PTRARG, INTARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "memfd_create",
		num:  319,
		args: SystemCallArgs{STRINGARG, INTARG, 0, 0, 0, 0},
	},
	SystemCall{
		name: "kexec_file_load",
		num:  320,
		args: SystemCallArgs{INTARG, INTARG, INTARG, PTRARG, INTARG, 0},
	},
	SystemCall{
		name: "bpf",
		num:  321,
		args: SystemCallArgs{INTARG, PTRARG, INTARG, 0, 0, 0},
	},
	SystemCall{
		name: "execveat",
		num:  322,
		args: SystemCallArgs{INTARG, PTRARG, PTRARG, PTRARG, INTARG, 0},
	},
}
