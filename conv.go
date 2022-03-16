package rice

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/google/uuid"
	"github.com/rs/xid"
)

// TimeNowFormat time.Time to "2006-01-02 15:04:05"
func TimeNowFormat() string { return time.Now().Format("2006-01-02 15:04:05") }

// TimeFormat time.Time to "2006-01-02 15:04:05"
func TimeFormat(tm time.Time) string { return tm.Format("2006-01-02 15:04:05") }

// TimeParse "2006-01-02 15:04:05" to time.Time
func TimeParse(s string) (time.Time, error) {
	if tm, err := time.Parse("2006-01-02 15:04:05", s); err != nil {
		return time.Time{}, err
	} else {
		return tm, nil
	}
}

// TimeFormatDate time to date
func TimeFormatDate(tm time.Time) string { return tm.Format("2006-01-02") }

// TimeParseDate date to time
func TimeParseDate(date string) (time.Time, error) { return time.Parse("2006-01-02", date) }

// TimeUnix 1645270804 to time.Time
func TimeUnix(sec int64) time.Time { return time.Unix(sec, 0) }

// StrconvParseInt "1645270804" to 1645270804
func StrconvParseInt(s string) (int64, error) {
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}

// StrconvFormatInt 1645270804 to "1645270804"
func StrconvFormatInt(i int64) string { return strconv.FormatInt(i, 10) }

// StrconvParseFloat int/100 to float64 保留 2 位小数点
// example: 2333 -> 23.33
func StrconvParseFloat(i int) (float64, error) {
	if f, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(i)/100), 64); err != nil {
		return 0, err
	} else {
		return f, nil
	}
}

// ByteString []byte to string
func ByteString(b []byte) string { return *(*string)(unsafe.Pointer(&b)) }

// BytesNewBufferString []byte to string
func BytesNewBufferString(b []byte) string { return bytes.NewBuffer(b).String() }

// FmtSprintfByte []byte to string
func FmtSprintfByte(b []byte) string { return fmt.Sprintf("%s", b) }

// StringByte string to []byte
func StringByte(s string) []byte {

	b := make([]byte, len(s))

	copy(b, s)

	return b
}

// StringByteUnsafe string to []byte
func StringByteUnsafe(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// UUIDNewString creates a new random UUID
func UUIDNewString() string { return strings.ReplaceAll(uuid.NewString(), "-", "") }

// RandInt 生成一个真随机数
func RandInt(max int64) (int64, error) {
	if result, err := rand.Int(rand.Reader, big.NewInt(max)); err != nil {
		return 0, err
	} else {
		return result.Int64(), nil
	}
}

func XidNewString() string { return xid.New().String() }

// SliceRemoveIndex 移除 slice 中的一个元素
func SliceRemoveIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

// SliceRemoveIndexUnOrder 移除 slice 中的一个元素（无序，但效率高）
func SliceRemoveIndexUnOrder[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
