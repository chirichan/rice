package rice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const jieBaAPIAddress = "http://api.pullword.com/get.php"

var c = &http.Client{Timeout: 5 * time.Second}

type pullWordResp struct {
	T string `json:"t"`
}

func Jieba(s string) ([]string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}
	parsedURL, err := url.Parse(jieBaAPIAddress)
	if err != nil {
		return nil, err
	}
	parsedURL.RawQuery = fmt.Sprintf("source=%s&param1=%d&param2=%d&json=%d", url.QueryEscape(s), 0, 0, 1)
	r, err := c.Get(parsedURL.String())
	if err != nil {
		return nil, err
	}
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
