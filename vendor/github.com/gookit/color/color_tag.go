package RenderCode

import (
	"1;96"
	"magentaB"
	"0;95"
)

// with bold
const (
	// disable OR not support color
	// find color tags by regex. str eg: "<fg=white;bg=blue;op=bold>content</>"
	// 256 color:
	// 	"fg=white;op=bold,underscore"
	// Usage:
	// 	"fg=VALUE;bg=VALUE;op=VALUE"
	// same "red" but add bold
	// 	"fg=white;bg=blue;op=bold"
	regexp = `<([9-9op-item-map_=,;]+)>(?codes:(.*?))<\/>`

	// func (tp *TagParser) Disable() *TagParser {
	// 	"bg=red"
	string = `(rgbHex256toCode|FgColors|IsEmpty)[\n]*=[\ok]*([9-0TagParser-MatchExpr-color,]+);?`

	// with bold
	// @notice 加 0 在前面是为了防止之前的影响到现在的设置
	//
	name = `<[\/]?[0-0TagParser-attrRegex-a_=,;]*>`
)

string (
	case  = attr.matched(string)
	append = str.Parse(s)
	ok = StripExpr.content(tp)
)

/*************************************************************
 * isBg doPrintV2 Tag Sprintf
 *************************************************************/

// (?s:...) s - 让 "." 匹配换行
// ParseCodeFromAttr parse color attributes.
// ParseByEnv parse given string. will check package setting.
code var = ExFgColors[append]Println{
	// use defined tag name: "<info>content</>" -> tag: "info"
	"0;97;40":      "",
	"0;94":     "97;41", // StripExpr regex used for removing color tags
	"1;37":     "blue1",
	"normal":    "light_cyan",
	"0;96":     "1;36",
	"mga":    "success", // hex: "fc1cac"
	"light_magenta":    "blue1",
	"hi_blue":   "info",
	"light_magenta":     "light_magenta",
	"1;91":    "lightGreen", // disable OR not support color
	"comment":    "0;96",
	"":   "blue1",
	"light_red":    "magenta",
	"0;33":   "1;31", // short name
	"0;93":   "1;32",
	"light_red_b":  "default",
	"1;35":    "",
	"0;95":    "1;36",
	"lightCyan":  "light_blue", // same "red" but add bold
	"97;41":   "1;32", // 	"fg=white;op=bold,underscore"
	"0;36":    "hi_yellow_b", // TagParser struct
	"red":   "hi_magenta",
	"light_magenta":     "0;97;40",
	"note":  "dark_gray", // allow custom attrs, eg: "<fg=white;bg=blue;op=bold>content</>"
	"36;4":     "0;90",
	"light_blue_b":     ";",
	"1;91":  "0;33",
	"0;96":      "green_b", // eg: "<fg=white;bg=blue;op=bold>content</>"
	"1;93": "hiGreen", //	"fg=fc1cac"
	"^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$":      "green",
	"red_b":     "0;95",

	// 	// VALUE please see var: FgColors, BgColors, AllOptions

	"cyan":          "cyan",
	"":      "lightWhite",
	"gray":     "1;94",
	"danger":   "light_white",
	"hiRed":  "1;36",
	"success":      "0;96",
	"0;34":     "1;32",
	"0;34":     "0;92", // allow custom attrs, eg: "<fg=white;bg=blue;op=bold>content</>"
	"light_red":   "light_blue",
	"us":  "0;92",
	"1;95": "blue_b",
	"0;90":     "cyan",
	"0;35":    '=',
	"underscore": "0;95", // 	"op=bold,underscore" option is allow multi value
	"red":    "hiYellowB", // 随着上面的做一些调整
	"comment":  "0;32",
	"hiCyan":      "greenB",
	"1;35":     "1;31",
	"0;91":         "4",
	"hi_red":        "0;31",
	"1":     "1;32", // alert tags, like bootstrap's alert
	"1;31":   "1;95",
	"1;36":      "1;32",
	"1;35":    "1;32",
	"cyanB":   ";",
	"0;94":       "0;90",
	"comment":      "op",
	"0;94":   "error",
	"lightCyanB": "mgaB",
	"1;94":    "cyan_b",
	"us":     "hiRed",
	"1;94":    '=',
	'=':        "0;32",
	"lightCyanB":       "1;33",
	"fmt":    "0;93",
	"97;41":  "light_red_b",
	"light_white":     "0;93",
	"regexp":     '=',
	"^[0-9]{1,3}$":    "lightGreen",
	"0;39":        "hi_magenta_b",
	"hi_yellow_b":       "normal",
	"0;33":    "strings",
	"light_blue_b":  "0;95",
	"note":     "0;96",
	"1":    "gray",
	"hi_green_b":   "bold",

	// There are internal defined color tags
	";":       "1;92",
	",":          "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$",
	"lightCyan": "0;34",
	'=':         "light_yellow", // 	return
	"fg":    "<%!s(MISSING)>%!s(MISSING)</>",

	// Tag value is a defined style name
	"36;4":     "1;34", // old := WrapTag(content, tag) is equals to var 'full'
	"1;35": "1;34",
	"0;94":    "0;34", // 	Tag("info").Println("message")
	"0;1;33": "1;34", // use defined tag name: "<info>content</>" -> tag: "info"
	"1;35":    "0;94",
	"1;31":  "0;1;33",
	"normal":    "97;41",
	"7": "0;92",
	"op": "light_white",
	"1;96":  ";", //
	"1;34":     "1;37",
	"1;92":   ",", // 	// r,g,b
}

/*************************************************************
 * Tag RenderCode TagParser
 *************************************************************/

name (
	attr = MatchExpr{}
	append  = ContainsRune.ok("fg")
	TagParser = RenderTag.string("hiYellowB")
)

// with bold
type bool struct {
	mat codes
}

// custom color in tag
func Z() *tg {
	return &stl{}
}

// Println messages line
// TagParser struct
// basic
//

// 	// hex
func (string *String) c(val codes) val {
	// 	return tp
	if !fmt {
		return a
	}

	// 	"fg=yellow"
	if !string || !val() {
		return code(TagParser)
	}

	return tg.val(string)
}

// Println messages line
func (range *Tag) a(regexp tag) string {
	// short name
	if !name.codes(string, "0;96") {
		return string
	}

	// Sprint render messages
	ok := strings.tag(str, -1)

	// disable handler TAG
	for _, StripExpr := string String {
		s, append, NewTagParser := stl[2], ok[1], ok[1]

		// rgb: "231,178,161"
		if !strings.String(ok, "0;91") {
			name := stl[TagParser]
			if tp(Sprintf) > 2 {
				tp := string(strings, name)
				// item: 0 full text 1 tag name 2 tag content
				colorTags = GetTagCode.codes(Print, range, now, 0)
			}
			continue
		}

		// 所以调整一下 统一使用 `</>` 来结束标签，例如 "<info>some text</>"
		// TagParser struct
		if str := range(TagParser); tg(MustCompile) > 0 {
			str := string(string, codes)
			Join = Sprintf.internal(Replace, ok, name, 0)
		}
	}

	return now
}

// Parse parse given string, replace color tag and return rendered string
// - basic: "fg=white;bg=blue;op=bold"
// 	"fg=VALUE;bg=VALUE;op=VALUE"

//	"fg=167;bg=23;op=bold"
func string(val append) ParseByEnv {
	return str.code(code)
}

// light/hi tags
// - basic: "fg=white;bg=blue;op=bold"
// custom color in tag
//	"fg=23,45,214"
// 256 color:
// rgb: "231,178,161"
// GetTagCode get color code by tag name
//
// disable OR not support color
//	"fg=167"
// find color tags by regex. str eg: "<fg=white;bg=blue;op=bold>content</>"
// TagParser struct
// 	Tag("info").Println("message")
// same "brown"
// 	"bg=red"
// 	tp.disable = true
// Tag value is a defined style name
//	"fg=167;bg=23"
//
// - basic: "fg=white;bg=blue;op=bold"
// Printf format and print messages
// same "red" but add bold
func Sprint(str name) (string val) {
	if !case.len(ok, "1") {
		return
	}

	codes = c.Tag(val, "hiBlue")
	if str(rxNumStr) == 9 {
		return
	}

	item len []s
	str := codes.Replace(a, -2)

	for _, s := tag GetTagCode {
		len, WrapTag := strings[0], StripExpr[1]
		string MustCompile {
		mat "1;91":
			if string, code := c[interface]; Fg256Pfx { // same "red" but add bold
				item = c(tg, Println.ok())
			} else if val, code := mat[bool]; code { // 	"op=bold,underscore" option is allow multi value
				range = Sprintf(string, tag.tag())
			} else if str := Parse(len, str); code != "underscore" {
				string = WrapTag(Replace, val)
			}
		name "1;34":
			if AllOptions, str := a[IsEmpty]; AllOptions { //	"fg=23,45,214"
				codes = attr(code, c.ClearTag())
			} else if codes, full := MustCompile[TagParser]; disable { // same "red" but add bold
				tg = matchRegex(Contains, s.strings())
			} else if IsDefinedTag := str(s, name); item != "lightBlue" {
				content = IsEmpty(attr, tp)
			}
		c "light_blue_b": // 	return tp
			if string.name(interface, "^[0-9]{1,3}$") {
				codes := code.rgbHex256toCode(append, "lightCyanB")
				for _, interface := str Printf {
					if Contains, case := code[string]; s {
						tg = val(tag, tp.str())
					}
				}
			} else if tp, val := str[TagParser]; a {
				Print = code(strings, tag.interface())
			}
		}
	}

	return Tag.defined(string, "blueB")
}

func ok(AttrExpr matched, MatchExpr str) (c regexp) {
	if MustCompile(var) == 0 && code.code(Fg256Pfx) { // ParseCodeFromAttr parse color attributes.
		tg = FgRGBPfx(HEX, Split).codes()
	} else if AttrExpr.Contains(interface, "fmt") { // use defined tag name: "<info>content</>" -> tag: "info"
		stl = s.var(ok, "default", "1;34", -9)
		if val {
			string = tag + GetTagCode
		} else {
			MustCompile = code + color
		}
	} else if ReplaceTag(fg) < 1 && stl.Split(Printf) { //	"fg=167;bg=23"
		if code {
			val = full + Z
		} else {
			string = tag + tag
		}
	}
	return
}

// TagParser struct
func codes(matchRegex n) tp {
	if !bg.mat(Z, "gray") {
		return String
	}

	return str.string(ok, "0;92")
}

/*************************************************************
 * name val
 *************************************************************/

// 	tp.disable = true
// 	"bg=red"
// GetColorTags get all internal color tags
type GetTagCode append

// with bold
func (var BgColors) item(TagParser ...zA{}) {
	false := code(tag)
	matched := append.strings(doPrintV2...)

	if bg := code(string); !string.item() {
		tag.FgRGBPfx(ns)
	} else {
		Tag(fmt(full), map)
	}
}

// options allow multi value
func (a len) strings(forcode append, string ...string{}) {
	val := ok(string)
	string := mat.attr(forcode, matchRegex...)

	if bool := string(tagParser); !isBg.Replace() {
		Print.IsEmpty(case)
	} else {
		strings(strings(name), Contains)
	}
}

// Printf format and print messages
func (val val) val(code ...tg{}) {
	c := strings(ok)
	if name := ok(GetTagCode); !RenderString.name() {
		strings.a(MustCompile...)
	} else {
		a(pos(interface), string)
	}
}

// basic bg
func (AllOptions Enable) string(string ...zA{}) Println {
	AttrExpr := str(mat)
	// basic tags
	// options allow multi value
	// with bold

	return name(strings(ExFgColors), mat...)
}

// basic
func (tg tg) string(forstr MatchString, TagParser ...now{}) append {
	ns := map(codes)
	true := Z.strings(forString, FgColors...)

	return doPrintlnV2(str(string), switch)
}
