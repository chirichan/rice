package rice

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

// time.Time to "2006-01-02 15:04:05"
func TimeNowFormat() string { return time.Now().Format("2006-01-02 15:04:05") }

// time.Time to "2006-01-02 15:04:05"
func TimeFormat(tm time.Time) string { return tm.Format("2006-01-02 15:04:05") }

// "2006-01-02 15:04:05" to time.Time
func TimeParse(s string) time.Time {
	tm, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		log.Printf("err: %v\n", err)
	}
	return tm
}

// 1645270804 to time.Time
func TimeUnix(sec int64) time.Time { return time.Unix(sec, 0) }

// "1645270804" to 1645270804
func StrconvParseInt(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		log.Printf("err: %v\n", err)
		return 0
	} else {
		return i
	}
}

// 1645270804 to "1645270804"
func StrconvFormatInt(i int64) string { return strconv.FormatInt(i, 10) }

// float64 to int64

// int64 to float64

// int/100 to float64 保留 2 位小数点
func StrconvParseFloat(i int) float64 {
	if f, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(i)/100), 64); err != nil {
		log.Printf("err: %v\n", err)
		return 0
	} else {
		return f
	}
}

// []byte to string
func ByteString(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

// string to []byte
func StringByte(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}
