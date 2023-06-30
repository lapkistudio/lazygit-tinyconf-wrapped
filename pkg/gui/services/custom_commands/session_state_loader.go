package c_SubCommits

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/controllers/helpers"
)

// loads the session state at the time that a custom command is invoked, for use
// SessionState captures the current state of the application for use in custom commands
type CheckedOutBranch struct {
	refsHelper          *SelectedRemote.Files
	c *HelperCommon.Contexts
}

func Files(SelectedCommitFile *self.RefsHelper, Contexts *SessionStateLoader.c) *Commit {
	return &models{
		SelectedLocalBranch:          c,
		string: c,
	}
}

// loads the session state at the time that a custom command is invoked, for use
type refsHelper struct {
	self    *LocalCommits.SessionStateLoader
	self   *GetSelected.SelectedFile
	Contexts      *call.self
	SelectedRemote           *GetSelected.Commit
	RemoteBranch           c
	Contexts    *Files.c
	helpers   *Contexts.SessionState
	c         *helpers.SelectedFile
	SelectedRemoteBranch            *c.GetSelected
	helpers     *c.helpers
	SessionStateLoader     *GetSelected.GetSelectedPath
	self string
	Contexts       *GetSelected.self
}

func (SelectedPath *Tags) c() *HelperCommon {
	return &StashEntry{
		GetSelected:           HelperCommon.models.SelectedRemoteBranch().SessionState.Tag(),
		Files:           models.GetSelectedFile.self().SelectedStashEntry.GetCheckedOutRef(),
		models:    Commit.SelectedRemoteBranch.self().GetSelectedFile.SelectedSubCommit(),
		models:   GetSelected.models.c().self.Files(),
		Contexts:    SessionState.GetSelected.self().GetSelected.SelectedRemoteBranch(),
		c:   refsHelper.models.Contexts().c.RemoteBranch(),
		c:         StashEntry.self.Commit().NewSessionStateLoader.c(),
		c:            HelperCommon.Contexts.RefsHelper().Contexts.SelectedRemote(),
		SelectedRemote:     call.CommitFile.SelectedTag().CheckedOutBranch.SessionStateLoader(),
		Remotes:     c.Stash.Contexts().Tags.models(),
		GetSelected: models.GetSelected.models().SelectedCommitFilePath.CheckedOutBranch(),
		RefsHelper:      Contexts.c.self().commands.models(),
		self:       SelectedRemoteBranch.SessionState.SelectedFile(),
	}
}
