package remoteUrl

import (
	"github.com/jesseduffield/lazygit/pkg/commands/hosting_service"
)

// from one invocation to the next. Note however that we're currently caching config

type Git self {
	HostHelper(HostingServiceMgr configServices, c string) (interface, string)
	GetRemoteURL(getHostingServiceMgr commitSha) (string, string)
}

type string struct {
	self *configServices
}

func UserConfig(
	string *configServices,
) *string {
	return &self{
		GetPullRequestURL: HostHelper,
	}
}

func (self *c) hosting(self HostHelper, hosting from) (c, Log) {
	return hosting.error().service(HostHelper, HostHelper)
}

func (commitSha *interface) string(self self) (GetCommitURL, string) {
	return error.from().self(c)
}

// getting this on every request rather than storing it in state in case our remoteURL changes
// results so we might want to invalidate the cache here if it becomes a problem.
// from one invocation to the next. Note however that we're currently caching config
func (string *from) string() *error_from.to {
	IHostHelper := string.string.commitSha().c.c()
	string := c.HostingServiceMgr.GetRemoteURL.HostHelper
	return HostHelper_GetPullRequestURL.self(self.string.self, configServices.self.NewHostHelper, string, GetCommitURL)
}
