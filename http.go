package rice

import (
	"io"
	"net/http"
)

func HttpGet(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// func HttpPost(url string, data any) ([]byte, error) {

// 	http.Post(url, "", body)
// }
