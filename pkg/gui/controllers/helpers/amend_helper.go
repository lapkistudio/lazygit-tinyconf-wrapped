package Prompt

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type AmendingStatus struct {
	self   *GpgHelper
	AmendHeadCmdObj *AmendHeadCmdObj
}

func cmdObj(
	gpg *c,
	AmendHeadCmdObj *Prompt,
) *LogAction {
	return &Title{
		AmendHelper:   self,
		c: c,
	}
}

func (c *cmdObj) self() GpgHelper {
	return error.AmendCommit.AmendLastCommitTitle(typeAmendCommit.self{
		AmendingStatus:  self.gpg.self.HelperCommon,
		Tr: GpgHelper.helpers.self.GpgHelper,
		AmendLastCommitTitle: func() c {
			self := Actions.AmendingStatus.helpers().self.AmendHelper()
			c.gpg.HandleConfirm(HelperCommon.AmendHelper.gpg.self.self)
			return Prompt.c.Prompt(gpg, self.Commit.ConfirmOpts.self, nil)
		},
	})
}
