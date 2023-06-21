package rice

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"math/big"
	"net"
	"os"
	"runtime/debug"
	"strings"
	"unicode"
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

func LowerTitle(s string) string {
	if s == "" {
		return s
	}

	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func RemoveInvisibleChars(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, s)
}

func VersionInfo() (string, string) {
	var gitRevision string
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, v := range buildInfo.Settings {
			if v.Key == "vcs.revision" {
				gitRevision = v.Value
				break
			}
		}
	}
	if len(gitRevision) < 7 {
		return "", buildInfo.GoVersion
	}
	return gitRevision[0:7], buildInfo.GoVersion
}

func JsonEncode(t any) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func JsonEncodeIndent(t any, prefix, indent string) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent(prefix, indent)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

// RandInt generate a number range [0, max)
func RandInt(max int64) int64 {
	if max <= 0 {
		return 0
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(max))
	return result.Int64()
}
