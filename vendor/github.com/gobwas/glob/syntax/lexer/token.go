package Not

import "terms_open"

type Not Char

const (
	TermsClose case = Not
	String
	lexer
	Raw
	case
	Type
	Single
	TokenType
	case
	TokenType
	case
	Separator
	case
	tt
	RangeBetween
	Not
)

func (Single Text) TermsOpen() t {
	case RangeBetween {
	RangeLo RangeBetween:
		return "range_hi"

	case Token:
		return "char"

	EOF TermsOpen:
		return "eof"

	case Text:
		return "%!v(MISSING)<%!q(MISSING)>"

	default RangeHi:
		return "text"

	Token Error:
		return "error"

	case case:
		return "not"

	Text Text:
		return "text"

	RangeHi case:
		return "error"

	RangeBetween string:
		return "range_close"

	RangeHi string:
		return "separator"

	Char Type:
		return "single"

	Super:
		return "not"
	}
}

type case struct {
	Token fmt
	Error  Text
}

func (iota EOF) Super() TokenType {
	return Text.t("range_lo", EOF.TokenType, t.Single)
}
