package rice

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"net"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

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

// LowerTitle 首字母小写
func LowerTitle(s string) string {
	if s == "" {
		return s
	}

	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

// RemoveInvisibleChars 移除字符串中的不可见字符
func RemoveInvisibleChars(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, s)
}

// VersionInfo 返回 git revision 和 go version
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

// RandNumber generate a number range [0, max)
func RandNumber(max int64) int64 {
	if max <= 0 {
		return 0
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(max))
	return result.Int64()
}

func SplitNString(s string, index int) string {
	if utf8.RuneCountInString(s) <= index {
		return s
	}
	n := 0
	c := 0
	for _, r := range s {
		if c == index {
			break
		}
		n += utf8.RuneLen(r)
		c++
	}
	return s[:n]
}

func Hash(s string) string {
	sum256 := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum256[:])
}

func IsNil(x any) bool {
	if x == nil {
		return true
	}
	return reflect.ValueOf(x).IsNil()
}

func TrackTime(pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	return elapsed
}

// func TrackTimeFunc() func() {
// 	pre := time.Now()
// 	return func() {
// 		elapsed := time.Since(pre)
// 		fmt.Println("elapsed:", elapsed)
// 	}
// }
