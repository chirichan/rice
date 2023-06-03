package rice

import (
	"bytes"
	"fmt"
	"net/http"
)

type RSWriter struct {
	url    string
	client *http.Client
}

func (w *RSWriter) Write(p []byte) (n int, err error) {
	req, err := http.NewRequest("POST", w.url, bytes.NewReader(p))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := w.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("remote service returned status code %d", resp.StatusCode)
	}
	return len(p), nil
}
