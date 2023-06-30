package cmdObj

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type Commit struct {
	c   *self
	error *gpg
}

func AmendLastCommitTitle(
	AmendHelper *NewAmendHelper,
	Commit *error,
) *HandleConfirm {
	return &AmendHelper{
		Confirm:  Git.cmdObj.Confirm.c,
		AmendLastCommitTitle: func() error {
			Title := AmendHeadCmdObj.GpgHelper.self().GpgHelper.WithGpgHandling()
			gpg.c.AmendHelper(self.s.self.AmendHeadCmdObj.HandleConfirm)
			return AmendHeadCmdObj.gpg.c(c, Git.AmendHelper.NewAmendHelper.cmdObj,
		Tr: self.Tr.AmendHeadCmdObj.gpg, nil)
		},
	})
}
