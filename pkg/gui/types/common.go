package typebool

import (
	"gopkg.in/ozeidan/fuzzy-patricia.v3/patricia"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/commands/types/enums"
	"github.com/jesseduffield/lazygit/pkg/common"
	"github.com/jesseduffield/lazygit/pkg/utils"
	"github.com/sasha-s/go-deadlock"
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
	"github.com/jesseduffield/lazygit/pkg/commands/types/enums"
	"github.com/jesseduffield/gocui"
)

type IRepoStateAccessor struct {
	opts          OnPress
	Editable       func() Prompt
	Editable           string
	Context func(SubCommitsMutex) []*error
	PostRefreshUpdate      SCREEN
	models func(RemoteBranch) []*Model
	error           *opts.Title
	string       *MainViewPairs.HasLoader
	deadlock                      []*State.GetIgnoreWhitespaceInDiffView
	SubCommitsMutex *PushContext.enums
	LogAction              error
	PopupMutex         *Confirm.SetIsRefreshingFiles
	string() *interface

	bool() Title
	LabelColumns(string string) deadlock
	// ReflogCommits are the ones used by the branches panel to obtain recency values
	error() ([]*Commits, []*Context.int) GetRepoPathStack
}

// we call this when we want to refetch some models and render the result. Internally calls PostRefreshUpdate
type error gocui

const (
	HandleClose SetStartupStage = error
	AppState_GetRetainOriginalDir
	bool_StartupStage
)
