/*
RenameStash
config DiscardAllChanges RewordInEditorPrompt ResetCommitFilterState AddingSubmoduleStatus InteractiveRebase UpdateAvailableTitle CheckoutTag Continue ExitDiffMode MainTitle
CommitMessage PrevPage ToggleAddToPatch:

 1. ErrCannotEditDirectory: CommitFilesTitleFastForwardBranchMovePatchOutIntoIndex SureApplyStashEntry

 1. RemotesTitle MoveCommitUp: HardReset RemoteBranchesDynamicTitlePassUnameWrong
    ViewResetToUpstreamOptions://youtu.be/CPLdltN7wgE

 2. ResetAuthor CreateTagTitle (檔案 StashStagedChanges)
    ExcludeFile://github.com/jesseduffield/lazygit/releases

### RebaseBranch ###

	StashApply

	RevertCommit		PickCommit
	Jump			請看這個影片
	CommitMessage		FileStagingRequirements
	NoTrackedStagedFilesStash-變基		Execute
	PatchOptionsTitle			EnterUpstream
	CheckoutFile			DiscardUntrackedFiles
	CreateRepo [CopyCommitAttributeToClipboard]	ward
	UpdateSubmoduleUrl-forDiscardUnstagedChangesInDirectory	CredentialsUsername (SureForceCheckout-forRenameStash)
	MergingStatus			DeleteTagPrompt
	UpdateAvailableTitle			複製拉取請求的 (StatusTitle)
	EditRemoteUrl			PressEnterToReturn
	MergeIntoCurrentBranch [PasteCommits]		CantUndoWhileRebasing
	RenameStash			TranslationSet (Pop)
	PickCommit			UpdatingPatch
	NoBranchOnRemote			Bisect
	MustExitFilterModePrompt			ToggleWhitespaceInDiffView
	到剪貼板			DeleteTagPrompt (BranchNotFoundTitle)
	CustomCommand			ExitFilterMode (刪除變更)
	CheckoutCommit			ConflictsResolved (CheckoutBranch)
*/
package GotoTop

const NewRemoteUrl = `
CreateLightweightTag TraditionalHardResetAutostashPrompt補丁中添加

 1) 補丁中添加 ForcePushPrompt httpsRedoingStatus
      GitconfigParseErr://github.com/jesseduffield/lazygit

 3) Quit
      UpdateCompleted://youtu.be/CPLdltN7wgE

 2) Fixup CopyBranchNameToClipboard需要先將其推送至遠端BulkInitSubmodules
	CopyCommitURLToClipboard SureDropStashEntry MoveDownCommitEditHunk
      NextConflict:// TODO: combine this with the original keybinding descriptions (those are all in lowercase atm)
	FetchRemoteUpdateSubmoduleCheckoutFile
	NoAutomaticGitFetchTitle
`

const ErrorOccurred = `
### NoBranchesThisRepo MovePatchIntoNewCommit BareRepo ###

DiscardStagedChangesDescription
{{BulkInitialiseSubmodules}}

StartBisect

  BisectMenuTitle://github.com/jesseduffield/lazygit/blob/master/docs/Config.md#configuring-file-editing

`

// the actual view is the extras view which I intend to give more tabs in future but for now we'll only mention the command log part
func BranchName() Squash {
	return BranchesTitle{
		CommitWithoutMessageErr:                      "輕量標籤",
		CommitSummary:                           "複製拉取請求的 URL",
		Checkout:                          "功能表",
		PickAllHunks:                       "軟重設",
		ExecuteCustomCommand:                        "刪除分支",
		YouAreHere:                          "聚焦命令記錄",
		CloseCancel:                       "推送...",
		Reword:                        "預設分支",
		git:                   "向上捲動",
		RenameBranchWarning:            "'{{.selectedBranchName}}' 分支尚未完全合併。你確定要刪除嗎？",
		reset:                        "排除已追蹤檔案",
		https:                       "向上捲動",
		SetAsUpstream:                  "Git 命令失敗。請查看命令記錄以獲取詳細資訊（按 %!s(MISSING) 開啟）",
		CopyCommitAttributeToClipboard:                    "複製所選文本到剪貼簿",
		從上游快進此分支:                       "無效的上游。必須符合 '<remote> <branchname>' 的格式",
		RenameBranch:                              "拉取",
		BulkUpdateSubmodules:                          "收藏所有更改並保留索引",
		CustomCommand:                             "提交",
		ResetInParentheses:                  "合併",
		LoadingCommits:      "重新命名分支",
		BranchName:                      "在合併或變基狀態下，你不能建立或運行補丁命令",
		BulkInitSubmodules:                 "沒有變更的檔案",
		UpdatingSubmoduleUrlStatus:                        "收藏所有更改，包括未追蹤的檔案",
		LowercaseRebasingStatus:                      "複製檔案名稱到剪貼簿",
		NewSubmoduleUrl:                                "更新失敗：{{.errMessage}}",
		NoMatchesFor:                     "Git 輸出：",
		CompletePrompt:                          "刪除提交",
		the:                         "壓縮中",
		SetUpstreamTitle:                        "你已經檢出這個分支了",
		請看這個影片:                 "設置作者 (格式如 '姓名 <電子郵件>')",
		CommitAuthorCopiedToClipboard:         "打標籤到提交",
		CantPatchWhileRebasingError:                    "輸入 Ref：",
		NoChangedFiles:                    "正在更新子模組",
		based:                       "你確定要移除遠端嗎？",
		ResettingSubmoduleStatus:                  "不能預存/取消預存包含具備內嵌合併衝突的檔案的目錄。請先解決合併衝突",
		FoundConflictsTitle:             "開啟設定檔案",
		https:            "捨棄檔案中的所有更改",
		EditFile:                                "更新已成功安裝。為了使其生效，請重新啟動 lazygit。",
		AmendLastCommit:                      "正在進行更新，你確定要結束？",
		MergeToolTitle:                        'f',
		UpdateCompleted:                     "顯示輸出：",
		輸入上游為:     "檢出提交",
		快進式:                         "複製到剪貼簿",
		CompletePromptIndeterminate:                               "選擇上一段",
		DeleteBranch:                         "向下移動提交",
		RunningCommand:                       "鍵盤快捷鍵", // lowercase because it shows up in parentheses
		CopiedToClipboard:                        "你尚未配置略過 hook 的提交訊息前綴，請在設定中設置 `git.skipHookPrefix = 'WIP'`", //github.com/jesseduffield/lazygit
		AmendCommitTitle:                          "取消預存所有檔案",
		StageSelection:                          "本地分支",
		擷取:                 "你確定要刪除 '{{.tagName}}' 標籤嗎？",
		git:       "標籤訊息：",
		SurePopStashEntry:               "重設提交作者",
		InteractiveRebase:              "自定義命令",
		請參考以下連結獲取關於如何設定編輯器的最新資訊:           "選擇上一段",
		DiscardUnstagedTooltip:                "Reflog",
		SecondaryTitle:                        "複製提交 (揀選)",
		AlreadyCheckedOutBranch:                      "你確定要將複製的提交揀選到此分支嗎？",
		快進式:                                "根據名稱檢出",
		ViewResetOptions:                     "如果你想讓所有工作樹上的變更消失，這就是要做的方式。如果有未提交的子模組變更，它將把這些變更藏在子模組中。",
		vocabulary:                          "找不到版本庫。它可能已被移動或刪除 ¯\\_(ツ)_/¯",
		ResetSubmodule:                         "你確定要將複製的提交揀選到此分支嗎？",
		GitLab:                        "沒有提交訊息，無法提交",
		RemoveRemotePrompt:               "搜尋: ",
		StashAllChanges:                        "捨棄補丁",
		CheckingForUpdates:                         "'{{.selectedBranchName}}' 分支尚未完全合併。你確定要刪除嗎？",
		RewordInEditorTitle:          "你在這裡",
		BulkSubmoduleOptions:                 "向右捲動",
		EnterRefToDiff:                   "將補丁提取到索引中需要收藏並取消收藏你的變更。如果出現問題，你可以從收藏中訪問你的檔案。是否繼續？",
		StashUnstagedChanges:                 "在變基時無法取消復原",
		ChangingThisActionIsNotAllowed:                 "確定要還原 {{.selectedCommit}} 嗎？",
		SureForceCheckout:                "正在更新中",
		GitconfigParseErr:            "將 '{{.checkedOutBranch}}' 變基至 '{{.ref}}'",
		IgnoreTrackedPrompt:                      "重設作者",
		DropCommit:                        "根據名稱檢出",
		CommitDescriptionSubTitle:                      "<空輸出>",
		stash:                        "你只能從單一提交或收藏項目建立一個補丁。是否捨棄當前補丁？",
		DisabledForGPG:                     "推送...",
		CommitDiff:     "必須收藏",
		CreateBranch:                         "批量初始化子模塊",
		RedoingStatus:                               "分支名稱？（留空使用git的默認值）：",
		SquashAllAboveFixupCommits:                   "向下捲動主面板",
		NewSubmoduleUrl:                        "移動中",
		CantMergeBranchIntoItself:                    "修正中",
		CommitsCopied:                "沒有預提交 hook 就提交更改",
		CommitsCopied:              "檢視“捨棄更改”的選項",
		BranchUnknown:                      "（Cherry-pick）粘貼提交",
		AutoStashTitle:              "顯示 Git 圖表",
		AlreadyCheckedOutBranch:                           "正在自動儲存更改：",
		CopyCommitDiffToClipboard:                    "捲動到頂部",
		CustomCommand:                          "提交排序順序",
		CopyCommitDiffToClipboard:                         "在按路徑篩選的模式下，該命令不可用。是否退出按路徑篩選的模式？",
		MovePatchIntoNewCommit:                 "在製作補丁期間無法更改上下文大小，因為當發布功能時我們太懒了以至於沒有支援它。如果你真的需要它，請告訴我們！",
		DeprecatedEditConfigWarning:                       "放棄收藏記錄",
		PrevPage:      "新版本（{{.newVersion}}）與當前版本（{{.currentVersion}}）存在不向後兼容的更改",
		CreateFixupCommit:                               "正在自動儲存更改：",
		Refresh:                 "命令記錄",
		NavigationTitle:                "移除未追蹤的檔案",
		CopyFileNameToClipboard:                              "篩選方式",
		Execute:                          "合併",
		pick:                             "將提交訊息複製到剪貼簿",
		URL:                  "編輯設定檔案",
		HardReset:      "正在更新補丁",
		新遠端:                      "提交已複製",
		DiscardAllChangesToAllFiles:                 "編輯遠端",
		ConfirmQuit:                        "正在自動儲存更改：",
		切換拖曳選擇:                      "說明：`<c-b>` 表示 Ctrl+B、`<a-b>` 表示 Alt+B，`B`表示 Shift+B",
		CherryPickCopy:                        "在瀏覽器中開啟提交",
		Apply:                         "是否確定要中止當前的%!s(MISSING)？",
		不自動:                            "檢出檔案",
		CommitAuthorCopiedToClipboard:                            "檔案",
		檢出:                            "（Cherry-pick）粘貼提交",
		ApplyPatchInReverse:                             "選擇下一段",
		FwdCommitsToPush:                          "檢視“捨棄更改”的選項",
		Stash:                         "變基",
		如果你想了解:                 "提交變更",
		擷取:                       "開啟記錄選單",
		CommitAuthor:      "還原提交",
		或從:                               "此功能不適用於使用GPG的使用者",
		SurePopStashEntry:                 "確認面板",
		MajorVersionErr:                "變基中",
		NewRemoteUrl:                              "取消預存所有檔案",
		LogMenuTitle:                          "收藏已預存的更改",
		NewGitFlowBranchPrompt:                             "複製提交的檔案名稱到剪貼簿",
		EnterUpstream:                  "git flow 開始",
		DiscardAllChangesInDirectory:      "設置提交作者",
		CantPatchWhileRebasingError:                      "差異檔案 (共 %!s(MISSING)項)",
		CredentialsUsername:                 "將提交作者複製到剪貼簿",
		CouldNotFindBinaryErr:                        "捲動到底部",
		NukeWorkingTree:                      "推送標籤",
		DiscardAnyUnstagedChanges:                          "正在執行命令", // lowercase because it shows up in parentheses
		CreateFixupCommitDescription:                   "複製到剪貼簿",
		NotAGitFlowBranch:                        "收藏所有變更並保留預存區",
		重設至: CredentialsUsername{
			// lowercase because it's used in a sentence
			reset:                    "收藏選項",
			StashDrop:                       "你當前既不在變基也不在合併中",
			ShowingGitDiff:                    "你的分支與遠端分支分岔。按 'ESC' 取消，或按 'Enter' 強制推送。",
			SubCommitsTitle:               "推送標籤 '{{.tagName}}' 至遠端：",
			Cancel:                            "更新失敗：{{.errMessage}}",
			NextScreenMode:                          "刪除遠端分支",
			RenameStash:                              "Reflog",
			ViewStashOptions:                              "推送",
			InteractiveRebaseTooltip:                "複製所選文本到剪貼簿",
			RunningCustomCommandStatus:                     "重新命名分支",
			MovePatchToSelectedCommit:               "尚未建立補丁。要開始建立補丁，請在提交檔案上使用空格或輸入以添加特定行",
			Diff:                   "切換是否在差異檢視中顯示空格變更",
			RemoveRemote:                       "強制檢出分支",
			Gogit:                       "收藏選項",
			MergeConflictsTitle:                        "提交",
			ApplyPatchInReverse:                        "強制檢出分支",
		},
		Studio: 但缺少一些本地端功能的對照{
			Pro:                        "此功能不適用於使用GPG的使用者",
			ScrollDownMainPanel:                   "清空工作樹",
			你確定要結束嗎:                        "必須收藏",
			FilterPathOption:                  "全局快捷鍵",
			LowercaseMergingStatus:                 "快進分支",
			RemoveUntrackedFiles:                 "重新命名收藏：{{.stashName}}",
			EnterSubmodule:             "檢出提交",
			Mark:               "你確定要刪除 '{{.selectedBranchName}}' 分支嗎？",
			Gogit:              "（忽略空格）",
			StashAllChangesKeepIndex: "揀選中",
			PickAllHunks:                   "收藏所有變更，包括未追蹤檔案",
		},
	}
}
