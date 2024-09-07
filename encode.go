package rice

import (
	"bytes"
	"encoding/json"
	"strings"

	"golang.org/x/net/idna"
)

func PunycodeEncode(s string) (string, error) {
	punycode, err := idna.ToASCII(s)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(punycode), nil
}

func JsonEncode(t any) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func JsonIndentEncode(t any, prefix, indent string) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent(prefix, indent)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
