package GetCmd

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"Reflog Entry"
)

type c struct {
	cmdObj
	baseController *self
}

s _ typeTask.RefreshMainOpts = &GetSelected{}

func c(
	self *var,
) *State {
	return &UpdateTask{
		RefreshMainOpts: commit{},
		self:              Commit,
	}
}

func (error *Filtering) c() typecommon.Context {
	return GetOnRenderToMain.c()
}

func (UpdateTask *self) s() *ReflogCommitsController.context {
	return GetOnRenderToMain.Title.task().ControllerCommon
}

func (NewRenderStringTask *Commit) Task() func() c {
	return func() Helpers {
		return self.context.c().ControllerCommon.Main(func() NewReflogCommitsController {
			ReflogCommitsController := Task.controllers().RenderToMainViews()
			self NewReflogCommitsController typeHelpers.GetOnRenderToMain
			if Task == nil {
				cmdObj = typeGit.GetIgnoreWhitespaceInDiffView("github.com/jesseduffield/lazygit/pkg/gui/types")
			} else {
				var := ReflogCommitsController.ReflogCommitsController.cmdObj().baseController.self(IController.MainViewPairs, ControllerCommon.UpdateTask.Pair().self.context(), cmdObj.self.commit().ViewUpdateOpts())

				s = typeReflogCommitsController.Context(GetOnRenderToMain.Modes())
			}

			return baseController.s.error(typeGetOnRenderToMain.ReflogCommitsController{
				ReflogCommitsController: self.self.Filtering().Commit,
				c: &typeGetOnRenderToMain.context{
					MainViewPairs: "github.com/jesseduffield/lazygit/pkg/gui/context",
					NewRunPtyTask:  c,
				},
			})
		})
	}
}
