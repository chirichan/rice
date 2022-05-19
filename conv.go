package rice

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

// HMCompare 如果 1 > 2 return true
// 小时和分钟 比较大小
func HMCompare(h1, m1, h2, m2 int) bool {

	t := time.Date(0, 0, 0, h1, m1, 0, 0, time.Local)
	t2 := time.Date(0, 0, 0, h2, m2, 0, 0, time.Local)

	if t.Unix() > t2.Unix() {
		return true
	} else {
		return false
	}
}

// BetweenDays 两个时间之间隔了多少天 startTime >= endTime
func BetweenDays(startTime, endTime time.Time) int64 {

	var days int64

	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.Local)
	endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, time.Local)

	for i := startTime; i.Before(endTime); i = i.AddDate(0, 0, 1) {
		days += 1
	}
	return days
}

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

// TimeFormatUnixDate time to date
func TimeFormatUnixDate(stamp int64) string { return time.Unix(stamp, 0).Format("2006-01-02") }

// TimeFormatUnix time to date
func TimeFormatUnix(stamp int64) string { return time.Unix(stamp, 0).Format("2006-01-02 15:04:05") }

// TimeCompare 如果 t1>t2, return true, 如果 t1 <= t2, return false
func TimeCompare(t1, t2 time.Time) bool {

	if t1.After(t2) {
		return true
	} else {
		return false
	}
}

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
// func FmtSprintfByte(b []byte) string { return fmt.Sprintf("%s", b) }

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

func SliceToSlicePtr[T any](s []T) []*T {

	s2 := make([]*T, 0)

	for k := range s {
		s2 = append(s2, &s[k])
	}

	return s2
}

func SlicePtrToSlice[T any](s []*T) []T {

	s2 := make([]T, 0)

	for k := range s {
		s2 = append(s2, *s[k])
	}

	return s2
}
