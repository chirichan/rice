package rice

import (
	"bytes"
	"fmt"
	"strconv"
	"unsafe"
)

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
