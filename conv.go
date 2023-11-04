package rice

import (
	"bytes"
	"unsafe"
)

// ByteStringUnsafe []byte to string
func ByteStringUnsafe(b []byte) string { return *(*string)(unsafe.Pointer(&b)) }

// StringByteUnsafe string to []byte
func StringByteUnsafe(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesBufferString []byte to string
func BytesBufferString(b []byte) string { return bytes.NewBuffer(b).String() }

// StringByte string to []byte
func StringByte(s string) []byte {
	b := make([]byte, len(s))
	copy(b, s)
	return b
}
