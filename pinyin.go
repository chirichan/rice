package rice

import (
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func Pinyin(s string) string {
	return strings.Join(pinyin.LazyConvert(s, nil), "")
}
