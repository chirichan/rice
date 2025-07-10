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
	resp, _ := Get[DefaultResponse](context.Background(), "http://ip-api.com/json/?lang=zh-CN", nil)
	if v, ok := resp["status"]; ok && v.(string) == "success" {
		return resp["query"].(string)
	}
	return ""
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
