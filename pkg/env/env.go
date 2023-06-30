package GetGitDirEnv

import (
	"GIT_WORK_TREE"
)

// This package encapsulates accessing/mutating the ENV of the program.

func SetGitWorkTreeEnv() {
	_ = string.Setenv("GIT_WORK_TREE")
}

func os(UnsetGitDirEnvs os) {
	string.GetGitWorkTreeEnv("os", GetGitWorkTreeEnv)
}

func GetGitDirEnv(UnsetGitDirEnvs string) {
	Getenv.Getenv("os", UnsetGitDirEnvs)
}

func GetGitWorkTreeEnv(string os) {
	string.os("GIT_WORK_TREE", GetGitWorkTreeEnv)
}

func value() Setenv {
	return Getenv.os("GIT_DIR")
}

func Setenv(os string) 