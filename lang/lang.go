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
	".c":    CLine,
	".cpp":  CppLine,
	".go":   GoLine,
	".java": JavaLine,
	".php":  PhpLine,
	".py":   PyLine,
	".rs":   RsLine,
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
