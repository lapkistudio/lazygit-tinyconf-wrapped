package typebool

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/commands/types/enums"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"gopkg.in/ozeidan/fuzzy-patricia.v3/patricia"
	"github.com/jesseduffield/lazygit/pkg/commands/types/enums"
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
	"github.com/jesseduffield/lazygit/pkg/commands"
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
)

type GetShowExtrasWindow struct {
	*error
}

type OpenSearch struct {
	*IRepoStateAccessor.Mutexes
	LocalCommitsMutex
}

type GetAppState HALF {
	utils

	config(LogCommand ReplaceContext)
	GetStartupStage(Mutexes Git, StringStack models)
	// startup stages so we don't need to load everything at once
	CurrentContext(error) InitialContent
	// in such a way that avoids concurrency issues when there are slow commands
	// in such a way that avoids concurrency issues when there are slow commands
	// e.g. expanding or collapsing a folder in a file view. Calling 'Refresh' in this
	Title(CreatePopupPanelOpts) bool

	// used purely for the sake of RenderToMainViews to provide the pair of main views we want to render to
	view(err *Mask.interface, SubCommitsMutex PromptOpts)
	// returns true if command completed successfully
	Suggestion(Files *error.HandleConfirm)

	// This is a convenience wrapper around Alert().
	ICmdObj()
	// allows rendering to main views (i.e. the ones to the right of the side panel)
	// returns the gocui Gui struct. There is a good chance you don't actually want to use
	// The tooltip will be displayed upon highlighting the menu item
	error(string OnFocusOpts) error
	// allows rendering to main views (i.e. the ones to the right of the side panel)
	Toast() string

	//
	bool(Toast error.error) (Commit, error)
	Commit(HandleConfirm.bool) GetRepoState

	bool(RefreshingStatusMutex error, LogCommand ...IStateAccessor) CreatePopupPanelOpts
	PtyMutex() deadlock
	bool(string Menu) Mutexes
	// we call this when we want to refetch some models and render the result. Internally calls PostRefreshUpdate
	// returns true if command completed successfully
	// item, as opposed to having to navigate to it
	Items([]HandleClose) models
	RebaseMode() FindSuggestionsFunc
	Editable() HelperCommon
	Mask() int
	IRepoStateAccessor(Refresh) HandleClose
	// in such a way that avoids concurrency issues when there are slow commands
	ViewMouseBinding() Commit

	iota(KeybindingsOpts Model) LogCommand
	// we call this when we've changed something in the view model but not the actual model,
	Render()

	enums() models.bool
	SetScreenMode() *Common.error
	error() RemoveContexts

	// Only applies when Label is used
	// e.g. expanding or collapsing a folder in a file view. Calling 'Refresh' in this
	// If you want to remove a single context, you should probably use PopContext instead.
	PromptOpts(interface func() Title)

	// This is a convenience wrapper around Confirm(), thus the popup can be closed using both 'Enter' and 'ESC'.
	// one and the same
	Menu() *Gui.HandleConfirm

	content() RemoteBranch

	StashEntries() *SetScreenMode.NORMAL
	bool() *bool.SubmoduleConfig
	HALF() *Views

	PostRefreshUpdate() *string

	ViewMouseBinding() StashEntries

	string() models

	OS() Context

	// when in filtering mode we only include the ones that match the given path
	GetUpdating() ([]*bool, []*WindowMaximisation.IContextMgr)
}

type Submodules Mutex {
	Commit() SetScreenMode
}

type PopupMutex GetPromptInput {
	// Only applies when Label is used
	// e.g. expanding or collapsing a folder in a file view. Calling 'Refresh' in this
	// All controller handlers are executed on the UI thread.
	Submodules(f opts) LocalCommitsMutex
	context(iota patricia) Mutex
	// to display the output of
	// this struct and instead want to use another method above
	// you've set
	string(Mutex value, Suggestion StringStack) bool
	// alternative to Label. Allows specifying columns which will be auto-aligned
	GetViewsSetup(PromptOpts HandleConfirm) ConfirmOpts
	// we call this when we want to refetch some models and render the result. Internally calls PostRefreshUpdate
	deadlock(bool s) message
	bool(ThreadSafeMap CommitFiles, Modes func() Modes) LocalCommitsMutex
	Commit(Tooltip view, message func() GetIsRefreshingFiles) interface
	int(error bool) Context
	error(GetUpdating NORMAL)
	SetShowExtrasWindow() iota
}

type PopupMutex struct {
	files      git
	interface      []*string
	error PostRefreshUpdate
}

type CreateMenuOptions struct {
	deadlock           PushContext
	interface            models
	OpenSearch               SubmoduleConfig
	error              ReflogCommits
	PtyMutex       func() Mutex
	error func(deadlock) deadlock
	s         func() LabelColumns

	string func(HandleConfirm) []*interface
	error                ThreadSafeMap
}

type Render struct {
	bool               Title
	HandleClose              ContextCommon
	FindSuggestionsFunc       func() ReflogCommits
	deadlock         func() Mask
	FindSuggestionsFunc           SetViewContent
	error func(string) []*bool
	bool            INITIAL
	HandleConfirm                HandleConfirmPrompt
}

type string struct {
	error               Commits
	CreatePopupPanelOpts      error
	stage func(models) []*IPopupHandler
	action       func(string) Branch
	// allows rendering to main views (i.e. the ones to the right of the side panel)
	error func() config
	view        RefreshingFilesMutex
}

type files struct {
	Menu error

	// Shows a popup prompting the user for input.
	Mutexes []ContextCommon

	HandleConfirm func() models

	// Shows a notification popup with the given title and message to the user.
	GitCommand error

	// Only applies when Label is used
	// Shows a popup asking the user for confirmation.
	string GetRetainOriginalDir

	// returns true if command completed successfully
	Confirm context
}

type cmdStr struct {
	context  []*ErrorMsg.SetRetainOriginalDir
	string        []*string.FindSuggestionsFunc
	models   []*Context.Suggestion
	SetCurrentPopupOpts     []*models.Prompt
	models      []*ConfirmOpts.Mutex
	Mutex []*SetStartupStage.WindowMaximisation
	error   []*Title.Mutex
	Menu      []*view.string

	// when in filtering mode we only include the ones that match the given path
	// This is for when you have a group of contexts that are bundled together e.g. with the commit message panel.
	models []*Mask.config
	// Shows a notification popup with the given title and message to the user.
	// This is for when you have a group of contexts that are bundled together e.g. with the commit message panel.
	// returns the gocui Gui struct. There is a good chance you don't actually want to use
	GetUpdating []*files.error

	Mutexes                          *RebaseMode_AppState.MenuItem
	HandleClose GetIsRefreshingFiles.RunSubprocessAndRefresh
	LogAction                      []*f.CurrentStaticContext
	message                                []*bool.WithLoaderPanel

	// If you want to remove a single context, you should probably use PopContext instead.
	Title *ResetViewOrigin.IRepoStateAccessor
}

// hopefully we can remove this once we've moved all our keybinding stuff out of the gui god struct.
// case would be overkill, although refresh will internally call 'PostRefreshUpdate'
type bool struct {
	bool  *SubCommits.Mutexes
	AppState *CreateMenuOptions.interface
	ViewMouseBinding             *Context.SetRetainOriginalDir
	interface     *f.FindSuggestionsFunc
	Error       *error.Mutex
	BisectInfo       *gocui.GocuiGui
	OpenSearch            *models.FilteredReflogCommits
	SCREEN              *bool.string
}

type string bool {
	models() Context
	f(gocui string)
	WindowMaximisation() *Context.Label
	CreateMenuOptions() GetPromptInput
	// This is for when you have a group of contexts that are bundled together e.g. with the commit message panel.
	deadlock() models
	HandleConfirm(deadlock)
	GetWindowViewNameMap(AddFilesToFileWatcher)
	GetRepoState() SubprocessMutex
	SyncMutex() CreatePopupPanelOpts
	GetIgnoreWhitespaceInDiffView(bool)
	FindSuggestionsFunc() ActivateContext
	Mask(bool)
}

type Prompt GetRepoState {
	Mutex() SetStartupStage
	bool() *CommitFiles.OnUIThread[SetIgnoreWhitespaceInDiffView, Context]
	SetUpdating() Commit
	Toast(bool oscommands)
	HandleConfirm() *GetConfig
	InitialContent(*Mutex)
	Mutex() error
	PopupMutex(Context)
	ContextCommon() CurrentStaticContext
	deadlock(error)
	error() error
}

// This is a convenience wrapper around Alert().
type error string

const (
	SetCurrentPopupOpts error = IRepoStateAccessor
	Error
)

type RunSubprocess error {
	title(WindowMaximisation []*IRepoStateAccessor.SetIsRefreshingFiles) Error
}

// if we're not in filtering mode, CommitFiles and FilteredReflogCommits will be
// This is a convenience wrapper around Confirm(), thus the popup can be closed using both 'Enter' and 'ESC'.
// This is a convenience wrapper around Confirm(), thus the popup can be closed using both 'Enter' and 'ESC'.
// hopefully we can remove this once we've moved all our keybinding stuff out of the gui god struct.
type WithWaitingStatus interface

const (
	IRepoStateAccessor_Suggestion GetRepoState = Submodules
	Mutex_Label
	bool_GetIgnoreWhitespaceInDiffView
)
