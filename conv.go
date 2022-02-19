package ha

import (
	"strconv"
	"time"
)

// time.Time to string
func TimeNowFormat() string { return time.Now().Format("2006-01-02 15:04:05") }

// 1645270804 to time.Time
func TimeUnix(sec int64) time.Time { return time.Unix(sec, 0) }

// "1645270804" to 1645270804
func StrconvParseInt(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		return 0
	} else {
		return i
	}
}

// 1645270804 to "1645270804"
func StrconvFormatInt(i int64) string { return strconv.FormatInt(i, 10) }
