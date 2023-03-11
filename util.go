package rice

import (
	"errors"
	"net"
	"os"
	"strings"
)

func LocalHostname() (string, error) {
	return os.Hostname()
}

func LocalAddr() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	s := conn.LocalAddr().String()
	i := strings.LastIndex(s, ":")
	if i == -1 {
		return "", errors.New("can't get local addr")
	}
	return s[:i], nil
}
