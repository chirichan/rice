package rice

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/mozillazg/go-pinyin"
)

const pullWordAPIAddress = "http://api.pullword.com/get.php"

type pullWordResp struct {
	T string `json:"t"`
}

func Jieba(s string) ([]string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}
	parsedURL, err := url.Parse(pullWordAPIAddress)
	if err != nil {
		return nil, err
	}
	parsedURL.RawQuery = fmt.Sprintf("source=%s&param1=%d&param2=%d&json=%d", url.QueryEscape(s), 0, 0, 1)
	resp, err := Get[[]pullWordResp](parsedURL.String())
	if err != nil {
		return nil, err
	}
	var s2 []string
	for _, v := range resp {
		s2 = append(s2, v.T)
	}
	return s2, nil
}

func Pinyin(s string) string {
	return strings.Join(pinyin.LazyConvert(s, nil), "")
}

func JiebaAndPinyin(s string) ([]string, error) {
	s2, err := Jieba(s)
	if err != nil {
		return nil, err
	}
	for _, v := range s2 {
		s3 := Pinyin(v)
		if s3 != "" {
			s2 = append(s2, s3)
		}
	}
	return s2, nil
}
