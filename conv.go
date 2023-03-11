package rice

import (
	"bytes"
	"unsafe"
)

// ByteString []byte to string
func ByteString(b []byte) string { return *(*string)(unsafe.Pointer(&b)) }

// BytesBufferString []byte to string
func BytesBufferString(b []byte) string { return bytes.NewBuffer(b).String() }

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
