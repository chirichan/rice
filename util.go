package rice

import (
	"crypto/rand"
	"errors"
	"math/big"
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

func RandNumber(max int64) int64 {
	if max <= 0 {
		return 0
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(max))
	return result.Int64()
}
