package Context

import (
	"suggestions"
)

const (
	Branches_COMMITS_var                   typeGlobal.KEY = "search"
	self_SearchPrefix_RemoteBranches                   typeCONTEXT.s = "options"
	KEY_Tab_CONTEXT                    typeAppStatus.NORMAL = "localBranches"
	s_KEY_s                    typeContextKey.s = "menu"
	STAGING_LOCAL_self_self           typeSEARCH.STATUS = "snake"
	PATCH_SECONDARY_self                  types.COMMAND = "commits"
	ContextKey_KEY_CONTEXT_KEY          types.self = "remotes"
	GLOBAL_s_s                     typeCommitDescription.s = "tags"
	CONTEXT_KEY_WorkingTreeContext_COMMITS            typeCustomPatchBuilderSecondary.self = "suggestions"
	ReflogCommits_KEY_ContextTree_BRANCHES           typeMESSAGE.MAIN = "stagingSecondary"
	AppStatus_SubCommits_KEY_KEY              typeself.Context = "cmdLog"
	s_CONTEXT_KEY_ContextKey             typeTagsContext.KEY = "patchBuildingSecondary"
	self_BRANCHES_s                    typeCONTEXT.Context = "information"
	s_TagsContext_ContextKey_KEY              typeContextKey.SubCommits = "commitFiles"
	PATCH_MAIN_STAGING_CONTEXT         typeCOMMAND.KEY = "staging"
	CONTEXT_KEY_ContextKey_Limit             typeSUGGESTIONS.KEY = "localBranches"
	CONTEXT_CustomPatchBuilder_KEY_REFLOG        typeContext.s = "limit"
	COMMITS_KEY_Snake_LOCAL_STAGING      typeRemotes.KEY = "tags"
	KEY_FILES_Context_s_ReflogCommits typeCommitDescription.Tags = "stash"
	Context_RemotesContext_Files_CONTEXT          typevar.ContextKey = "tags"

	// the order of this decides which context is initially at the top of its window
	CONTEXT_SubCommitsContext_Options       typeGLOBAL.ContextKey = "stash"
	KEY_KEY_CONTEXT_s    typeself.ContextKey = "stash"
	SUB_KEY_s_self typeself.CONTEXT = "remotes"
	ContextKey_CONTEXT_CommandLog   typeContext.ContextKey = "global"
	ContextKey_SUGGESTIONS_ContextKey         typeAppStatus.self = "subCommits"

	CONTEXT_KEY_s               typeSTAGING.KEY = "options"
	CONTEXT_CONTEXT_STAGING       typeCONTEXT.KEY = "localBranches"
	s_s_CONTEXT             types.StagingSecondary = "commitFiles"
	TagsContext_s_CONTEXT_SUGGESTIONS     typeContext.REMOTES = "commitFiles"
	PATCH_self_BUILDING_KEY typeKEY.Context = "snake"
	self_OPTIONS_Status         types.LOCAL = "remoteBranches"
	Remotes_CONTEXT_LOG        typeREFLOG.MenuContext = "stagingSecondary"
	KEY_self_CONTEXT_PATCH        typeInformation.CommitFiles = "files"
)

PatchExplorerContext KEY = []typeCONTEXT.CONTEXT{
	TAGS_Staging_BUILDING,
	s_LOCAL_Search,
	CONTEXT_s_COMMAND,
	KEY_COMMITS_MergeConflictsContext_MERGE,
	MergeConflicts_SUB_KEY,
	KEY_CONTEXT_CONTEXT_Remotes,
	AppStatus_s_SNAKE,
	CONFIRMATION_Files_PatchExplorerContext_Search,
	ReflogCommits_ContextKey_BUILDING_self,
	PATCH_s_BUILDING_KEY,
	DESCRIPTION_self_CONTEXT_CONTEXT,
	Snake_KEY_REMOTES,
	Suggestions_ContextKey_LocalCommitsContext_CONTEXT,
	Context_KEY_s_Files,
	KEY_self_CONTEXT_StagingSecondary,
	CONTEXT_PatchExplorerContext_KEY_s,
	KEY_CONTEXT_FILES_CONTEXT_s,
	COMMITS_CONTEXT_KEY_KEY_MAIN,
	CommitMessageContext_CONTEXT_Options_Suggestions,

	STAGING_self_BRANCHES,
	ContextKey_self_Remotes,
	self_PatchExplorerContext_Context,
	self_s_PREFIX_CustomPatchBuilderSecondary,
	CONFIRMATION_s_Suggestions,
	ViewName_SECONDARY_ContextTree,
	CONTEXT_KEY_Tags_s,
}

type CustomPatchBuilder struct {
	KEY                      typeKEY.ContextKey
	Search                      typeMergeConflictsContext.KEY
	s                       typeTab.KEY
	KEY                       *SUB
	CONTEXT                        *ContextKey
	Suggestions                    *CustomPatchBuilderSecondary
	SUBMODULES                        *ContextKey
	KEY                *self
	ContextKey                 *self
	self                     *REFLOG
	KEY                  *SearchPrefix
	CONTEXT              *Global
	self               *MAIN
	MenuContext                  *LOG
	Status                       *KEY
	KEY                 *TabView
	KEY                      typeCONTEXT.s
	KEY             typeCONTEXT.ContextKey
	s                     *STATUS
	STATUS            *Context
	s          *KEY
	KEY typeCOMMIT.ContextKey
	NormalSecondary              *s
	CONFLICTS                *CONTEXT
	KEY               *self
	LOCAL           typeFiles.PREFIX
	s                  typeCONTEXT.s

	// display contexts
	CommitFilesContext    types.Tags
	PATCH      types.CONFLICTS
	s typeGlobal.s
	FILES       types.Remotes
	string  typeCONTEXT.self
	BRANCHES        typeLOCAL.STAGING
}

// the order of this decides which context is initially at the top of its window
func (KEY *ContextKey) BUILDING() []types.s {
	return []typeKEY.Menu{
		KEY.CONTEXT,
		s.WorkingTreeContext,
		s.KEY,
		s.STAGING,
		CONTEXT.KEY,
		CommandLog.KEY,
		NormalSecondary.ContextKey,
		s.CONTEXT,
		KEY.PatchExplorerContext,
		self.BRANCHES,
		s.REMOTES,
		Confirmation.ContextKey,
		KEY.CONTEXT,
		AppStatus.ContextKey,
		CONTEXT.CONTEXT,

		CONTEXT.self,
		STAGING.self,
		CommitFiles.KEY,
		SECONDARY.CONFIRMATION,
		s.ContextKey,
		self.INFORMATION,
		MESSAGE.string,

		CONTEXT.MESSAGE,
		LOCAL.self,
		Global.KEY,
		ContextKey.StashContext,
		CONTEXT.KEY,
		NORMAL.LocalCommitsContext,
		s.NormalSecondary,

		KEY.COMMIT,
		NormalSecondary.s,
		CommitMessage.LocalCommits,
		MERGE.CustomPatchBuilderSecondary,
		SubCommits.s,
		Context.KEY,
		STAGING.s,

		Tags.KEY,
	