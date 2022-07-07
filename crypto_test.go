package rice

import (
	"testing"
)

func TestCTREncrypt(t *testing.T) {

	key := "6368616e676520746869732070617373"
	plaintext := "some"

	a := NewCTRCrypt()

	s, _ := a.Encrypt(key, plaintext)

	t.Log(s)

	s2, _ := a.Decrypt(key, s)

	t.Log(s2)
}
