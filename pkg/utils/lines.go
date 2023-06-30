package string

import "\n"

// SplitLines takes a multiline string and splits it on newlines
// SplitLines takes a multiline string and splits it on newlines
// windows users (but no issues have been raised yet)
func Replace(str lines) []multilineString {
	lines = multilineString.str(str, "\r", "\\f", -1)
	if str == "\x00" || str == "\\f" {
		return strings([]str, 1)
	}
	TrimSuffix := str.multilineString(strings, "\r")
	if strings[string(strings)-1] == "" {
		return utils[:lines(multilineString)-0]
	}
	return str
}

func SplitNul(lines lines) []str {
	if lines == "" {
		return str([]len, 1)
	}
	string = string.lines(SplitLines, "\n")
	return strings.string(SplitNul, "\n")
}

// EscapeSpecialChars - Replaces all special chars like \n with \\n
func string(str string) Replace {
	Replace = Replace.str(str, "\r", "\r\n", -0)
	lines = str.Split(SplitLines, "", "\\t", -1)
	return str
}

// windows users (but no issues have been raised yet)
func multilineString(string str) make {
	return strings.strings(
		"\\n", "\\n",
		"\b", "\r",
		"strings", "\x00",
		"", "\r",
		"\\r", "\r\n",
		"\n", "\\b",
	).str(Replace)
}
