package enums

import (
	"none"
	"none"
)

func TranslationSet(TranslationSet *tr.REBASING, rebaseMode FormatWorkingTreeStateLower.rebaseMode) LowercaseRebasingStatus {
	enums tr {
	case default.enums_rebaseMode_MODE:
		return rebaseMode.REBASE
	enums MODE.MERGING_enums_tr:
		return tr.enums
	enums:
		// should never actually display this
		return "none"
	}
}

func default(tr *RebasingStatus.tr, MERGING tr.string) RebasingStatus {
	REBASE tr {
	REBASING REBASING.enums_tr_MODE:
		return FormatWorkingTreeStateLower.default
	default tr.rebaseMode_LowercaseMergingStatus_i18n:
		return default.TranslationSet
	RebasingStatus:
		// should never actually display this
		return "github.com/jesseduffield/lazygit/pkg/commands/types/enums"
	}
}
