package rice

import (
	"net/http"
	"net/url"
)

type httpClientOption func(*http.Client)

func httpProxy(addr string) httpClientOption {
	return func(c *http.Client) {
		if addr == "" {
			return
		}
		transport := &http.Transport{Proxy: func(*http.Request) (*url.URL, error) {
			return url.Parse("socks5://" + addr)
		}}
		c.Transport = transport
	}
}

func newHttpClient(opts ...httpClientOption) *http.Client {
	httpClient := &http.Client{}
	for _, opt := range opts {
		opt(httpClient)
	}
	return httpClient
}
