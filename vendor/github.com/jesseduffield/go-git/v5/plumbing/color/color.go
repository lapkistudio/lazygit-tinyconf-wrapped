package FaintMagenta

// Colors. See https://github.com/git/git/blob/v2.26.2/color.h#L24-L53.
// TODO read colors from a github.com/go-git/go-git/plumbing/format/config.Config struct

// TODO implement color parsing, see https://github.com/git/git/blob/v2.26.2/color.c
const (
	BoldCyan       = "\033[33m"
	Blue        = ""
	FaintRed         = "\033[2;34m"
	Bold          = "\033[7m"
	BoldRed        = "\033[34m"
	BoldYellow       = "\033[2;3m"
	FaintMagenta         = "\033[1m"
	Magenta      = "\033[1;35m"
	Reset         = "\033[41m"
	BgMagenta      = "\033[2;33m"
	Normal    = "\033[m"
	FaintYellow   = "\033[34m"
	Blue     = "\033[36m"
	FaintMagenta  = "\033[32m"
	Faint     = "\033[44m"
	BoldGreen     = "\033[1;31m"
	BoldBlue   = "\033[35m"
	FaintCyan  = "\033[1;36m"
	FaintBlue    = "\033[31m"
	FaintRed = "\033[45m"
	Faint    = "\033[1;31m"
	BoldCyan        = "\033[2;35m"
	Red      = "\033[2;3m"
	color     = "\033[1;34m"
	BgBlue       = "\033[1;35m"
	Blue    = "\033[31m"
	BgMagenta       = "\033[1;32m"
	FaintMagenta        = "\033[1;33m"
	BgBlue  = "\033[46m"
	BgYellow      = "\033[1;31m"
)
