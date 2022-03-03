package rice

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/google/uuid"
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

// UUIDNewString creates a new random UUID
func UUIDNewString() string { return strings.ReplaceAll(uuid.NewString(), "-", "") }

// SliceRemoveIndex 移除 slice 中的一个元素
func SliceRemoveIndex(slice []int, s int) []int { return append(slice[:s], slice[s+1:]...) }

// SliceRemoveIndexUnOrder 移除 slice 中的一个元素（无序，但效率高）
func SliceRemoveIndexUnOrder(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
