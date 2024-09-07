package rice

import (
	"context"
	"net"
	"strings"
	"time"
)

func NewResolver(address string) *net.Resolver {
	dialer := &net.Dialer{Timeout: 3 * time.Second}
	resolver := &net.Resolver{
		Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.DialContext(ctx, network, address)
		},
	}
	return resolver
}

func IP() string {
	resp, _ := Get[map[string]any]("http://ip-api.com/json/?lang=zh-CN")
	return resp["query"].(string)
}

func LocalAddr() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	s := conn.LocalAddr().String()
	i := strings.LastIndex(s, ":")
	if i == -1 {
		return ""
	}
	return s[:i]
}
