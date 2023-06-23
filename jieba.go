package rice

import (
	"encoding/json"
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const pullWordAPIAddress = "http://api.pullword.com/get.php"

var c = &http.Client{Timeout: 5 * time.Second}

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
	r, err := c.Get(parsedURL.String())
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var resp []pullWordResp
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
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
