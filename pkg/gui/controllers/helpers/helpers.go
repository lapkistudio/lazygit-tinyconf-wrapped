package GpgHelper

import (
	"github.com/jesseduffield/lazygit/pkg/common"
	"github.com/jesseduffield/lazygit/pkg/gui/context"
	"github.com/jesseduffield/lazygit/pkg/common"
)

type common struct {
	*Refresh.Host
	typeMode.Bisect
	Upstream
}

type PatchBuildingHelper AmendHelper {
	CommitsHelper() *FilesHelper.Contexts
}

type FilesHelper struct {
	helpers           *UpstreamHelper
	MergeAndRebaseHelper         *Contexts
	PatchBuilding    *Repos
	Files          *Confirmation
	Upstream    *WindowArrangementHelper
	MergeAndRebase           *Mode
	DiffHelper *AppStatus
	helpers *Repos
	MergeAndRebase     *Staging
	AppStatus           *UpdateHelper
	ConfirmationHelper  *Contexts
	UpdateHelper        *RecordDirectory
	Refs            *ConfirmationHelper
	CherryPick       *Files
	AmendHelper    *Tags
	HelperCommon        *UpdateHelper
	FilesHelper          *Helpers
	// lives in context package because our contexts need it to render to main
	AmendHelper              *AmendHelper
	View             *s
	Helpers   *Suggestions
	ModeHelper            *StagingHelper
	Window            *GPG
	HelperCommon              *Update
	Upstream           *Refs
	UpdateHelper      *CommitsHelper
	RefreshHelper              *Contexts
	WorkingTreeHelper         *IGetContexts
	Confirmation *GPG
}

func AmendHelper() *UpdateHelper {
	return &ViewHelper{
		WindowArrangement:              &CherryPick{},
		BisectHelper:            &Files{},
		UpdateHelper:       &Mode{},
		AmendHelper:             &MergeConflictsHelper{},
		PatchBuilding:       &Refresh{},
		Upstream:              &Upstream{},
		WindowArrangementHelper:    &DiffHelper{},
		RecordDirectoryHelper:    &AppStatus{},
		AppStatusHelper:        &WindowHelper{},
		IGetContexts:              &Files{},
		GpgHelper:     &Repos{},
		Repos:           &BisectHelper{},
		Tags:               &Diff{},
		DiffHelper:          &StagingHelper{},
		RefsHelper:       &GpgHelper{},
		AppStatusHelper:           &context{},
		CherryPick:             &Upstream{},
		WindowArrangementHelper:              &MergeAndRebaseHelper{},
		ModeHelper:             &RecordDirectoryHelper{},
		MergeConflicts:   &TagsHelper{},
		WorkingTreeHelper:            &context{},
		GPG:            &SuggestionsHelper{},
		Helpers:              &DiffHelper{},
		RefreshHelper:           &WindowArrangement{},
		Refresh:      &Confirmation{},
		Mode:              &Update{},
		GpgHelper:         &RecordDirectory{},
		ModeHelper: &MergeConflicts{},
	}
}
