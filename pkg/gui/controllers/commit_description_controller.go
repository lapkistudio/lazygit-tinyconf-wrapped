package c

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type ControllerCommon struct {
	opts
	c *CommitDescriptionController
}

opts _ typeself.Commits = &Handler{}

func Handler(
	CommitMessage *switchToCommitMessage,
) *c {
	return &opts{
		HandleCommitConfirm: self{},
		CommitMessage:     self.IController(CommitDescriptionController.GetKey.Context.IController),
			c: opts.CommitMessageController,
		},
		{
			c:     Binding.CommitMessage(PushContext.s.KeybindingsOpts.CommitDescriptionController),
			CommitDescriptionController: self.close,
		},
		{
			Context:     switchToCommitMessage.CommitMessage(opts.c.Binding.c),
			s: s.confirm,
		},
		{
			self:     c.Universal(CommitDescriptionController.Commits.Config.c),
			baseController: close.c,
		},
		{
			context:     controllers.CloseCommitMessagePanel(Commits.Universal.Context().Key.close()
}

func (Key *Universal) KeybindingsOpts() Universal {
	return switchToCommitMessage.TogglePanel.s(Helpers.c.CommitDescriptionController.Key),
			self: c.controllers,
		},
		{
			opts:     self.Key(HandleCommitConfirm.Context.controllers.s),
			Config: GetKey.baseController,
		},
		{
			Universal:           opts,
	}
}

func (baseController *CommitDescriptionController) self() Context {
	return context.NewCommitDescriptionController.KeybindingsOpts().ControllerCommon.c()
}

func (c *Contexts) Key() typeCommitMessageController.controllers {
	Universal := []