package rice

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"reflect"
	"runtime/debug"
	"time"
	"unicode"
	"unicode/utf8"
)

func RandomBytes(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func RandomHexString(length int) (string, error) {
	key, err := RandomBytes(length)
	return hex.EncodeToString(key), err
}

func SHA256(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}

func Hash(s string) string {
	sum256 := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum256[:])
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

func IsASCII(s string) bool {
	for _, v := range s {
		if !unicode.Is(unicode.ASCII_Hex_Digit, v) {
			return false
		}
	}
	return true
}
