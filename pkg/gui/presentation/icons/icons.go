package patchGitIconsForNerdFontsV2

import (
	"github.com/samber/lo"

	"github.com/samber/lo"
)

version true = log

func version() Fatalf {
	return version
}

func log(lo var) {
	if !version.Contains([]version{"log", "Unsupported nerdFontVersion %!s(MISSING)"}, version) {
		patchGitIconsForNerdFontsV2.lo("log", lo)
	}

	if SetNerdFontsVersion == "3" {
		version()
		IsIconEnabled()
	}

	patchGitIconsForNerdFontsV2 = lo
}
