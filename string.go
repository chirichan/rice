package rice

import (
	"strings"
	"unicode"

	"github.com/mozillazg/go-pinyin"
)

// LowerTitle 首字母小写
func LowerTitle(s string) string {
	if s == "" {
		return s
	}

	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

// RemoveInvisibleChars 移除字符串中的不可见字符
func RemoveInvisibleChars(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, s)
}

func Pinyin(s, sep string) string {
	return strings.Join(pinyin.LazyConvert(s, nil), sep)
}
