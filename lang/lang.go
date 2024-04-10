package lang

import (
	"strings"
)

type Lang struct {
	SingleLineComments []string
	MultiLineComments  []MultiLineComment
}

type MultiLineComment struct {
	Start string
	End   string
}

var languages = map[string]Lang{
	".c":      CLine,
	".cpp":    CppLine,
	".cc":     CppLine,
	".cxx":    CppLine,
	".C":      CppLine,
	".cs":     CsLine,
	".dart":   DartLine,
	".erl":    ErlLine,
	".ex":     ExLine,
	".exs":    ExLine,
	".f":      FLine,
	".for":    FLine,
	".f90":    FLine,
	".go":     GoLine,
	".groovy": GroovyLine,
	".gvy":    GroovyLine,
	".hs":     HsLine,
	".java":   JavaLine,
	".jl":     JlLine,
	".kt":     KtLine,
	".m":      MLine,
	".md":     MdLine,
	".mod":    ModLine,
	".php":    PhpLine,
	".pl":     PlLine,
	".pm":     PlLine,
	".py":     PyLine,
	".rb":     RbLine,
	".rs":     RsLine,
	".scala":  ScalaLine,
	".sh":     ShLine,
	".sql":    SqlLine,
	".swift":  SwiftLine,
	".xml":    XmlLine,
}

func IsSingleLineComment(ext string, line string) (is bool, supported bool) {
	lang, ok := languages[ext]
	if !ok {
		return false, false
	}

	for _, c := range lang.SingleLineComments {
		if strings.HasPrefix(line, c) {
			return true, true
		}
	}

	for _, mc := range lang.MultiLineComments {
		if strings.HasPrefix(line, mc.Start) && strings.HasSuffix(line, mc.End) {
			return true, true
		}
	}

	return false, true
}

func IsMultiLineCommentStart(ext string, line string) (is bool, start string, supported bool) {
	lang, ok := languages[ext]
	if !ok {
		return false, "", false
	}

	for _, c := range lang.MultiLineComments {
		if strings.HasPrefix(line, c.Start) {
			return true, c.Start, true
		}
	}

	return false, "", true
}

func IsMultiLineCommentEnd(ext string, line string, latestPrefix string) (is bool, supported bool) {
	lang, ok := languages[ext]
	if !ok {
		return false, false
	}

	if latestPrefix == "" {
		return false, true
	}

	for _, c := range lang.MultiLineComments {
		if latestPrefix == c.Start && strings.HasSuffix(line, c.End) {
			return true, true
		}
	}

	return false, true
}
