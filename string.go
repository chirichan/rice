package rice

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unsafe"

	"github.com/mozillazg/go-pinyin"
)

// ByteStringUnsafe []byte to string
func ByteStringUnsafe(b []byte) string { return *(*string)(unsafe.Pointer(&b)) }

// StringByteUnsafe string to []byte
func StringByteUnsafe(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

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

func StringToInt[T ~int | ~int32 | ~int64](s string) T {
	i, _ := strconv.Atoi(s)
	return T(i)
}

func IntToString[T ~int | ~int32 | ~int64](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

// CreateBidirectionalMapping 高效地创建两个字符串字符之间的双向映射
// 它在一次循环中同时创建正向和反向两个 map
func CreateBidirectionalMapping(from, to string) (map[rune]rune, map[rune]rune, error) {
	fromRunes := []rune(from)
	toRunes := []rune(to)

	// 安全检查：确保可以一一对应
	if len(fromRunes) != len(toRunes) {
		return nil, nil, fmt.Errorf("无法创建映射：两个字符串长度不同 (%d vs %d)", len(fromRunes), len(toRunes))
	}

	// 创建两个 map，并预设容量以提高性能
	// 这样 Go 运行时就无需在循环中频繁地为 map 扩容
	mapFwd := make(map[rune]rune, len(fromRunes)) // 正向: from -> to
	mapRev := make(map[rune]rune, len(toRunes))   // 反向: to -> from

	for i, fromRune := range fromRunes {
		toRune := toRunes[i]
		mapFwd[fromRune] = toRune
		mapRev[toRune] = fromRune
	}

	return mapFwd, mapRev, nil
}
