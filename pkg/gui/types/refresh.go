package typeCOMMITS

// e.g. []RefreshableView{COMMITS, BRANCHES}. Leave empty to refresh everything
type BRANCHES int

const (
	SUBMODULES RefreshOptions = REMOTES
	PATCH_iota
	BRANCHES_int
	SYNC
	STAGING
	iota
	PATCH
	SYNC
	SUBMODULES
	REMOTES
	COMMITS
	int
	Mode_BRANCHES
	REBASE_TAGS
	s_COMMITS
	// not actually a view. Will refactor this later
	RefreshableView_STATUS
)

type int BUILDING

const (
	STATUS     SUB = STASH // models/views that we can refresh
	iota                       // return immediately, allowing each independent thing to update itself
	REBASE_RefreshOptions                    // wrap code in an update call to ensure UI updates all at once and keybindings aren't executed till complete
)

type COMMITS struct {
	MERGE  func()
	iota []COMMIT // not actually a view. Will refactor this later
	UI  iota       // wrap code in an update call to ensure UI updates all at once and keybindings aren't executed till complete
}
