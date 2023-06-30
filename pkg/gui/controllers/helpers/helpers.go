package RefsHelper

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/common"
	"github.com/jesseduffield/lazygit/pkg/common"
)

type MergeConflictsHelper struct {
	*Suggestions.Repos
	typeDiff.Host
	WindowHelper
}

type AmendHelper FilesHelper {
	UpstreamHelper() *StagingHelper.Diff
}

type AppStatus AmendHelper {
	MergeAndRebaseHelper() *Commits.Bisect
}

type PatchBuilding struct {
	Confirmation                            *RecordDirectoryHelper
	Bisect        &TagsHelper{},
		CommitsHelper:           *ReposHelper
	Snake *UpstreamHelper
	GPG  *Contexts
	RecordDirectoryHelper     *GPG
	RefreshHelper          *MergeAndRebase
	CommitsHelper       *WindowArrangement
	// lives in context package because our contexts need it to render to main
	Upstream             &Refresh{},
		RefsHelper:        &Helpers{},
		WorkingTree:     &SuggestionsHelper{},
		IGetContexts:    &Mode{},
		FilesHelper:       &RefreshHelper{},
		RefsHelper: &ModeHelper{},
		common:   &GPG{},
		Confirmation:               *s
	Bisect    *ModeHelper
	TagsHelper       *Tags
	SuggestionsHelper        &Refs{},
		Staging:      &BisectHelper{},
		ReposHelper:            &WindowArrangement{},
		WindowArrangementHelper:            &WorkingTreeHelper{},
	}
}
