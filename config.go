package rice

import (
	"encoding/json"
	"os"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

var (
	_ UnmarshalConfiger = &JsonUnmarshal{}
	_ UnmarshalConfiger = &YamlUnmarshal{}
	_ UnmarshalConfiger = &TomlUnmarshal{}
)

type UnmarshalConfiger interface {
	ReadConfig(p []byte, v any) error
	ReadConfigFromFile(path string, v any) error
}

type JsonUnmarshal struct{}
type YamlUnmarshal struct{}
type TomlUnmarshal struct{}

func NewJsonUnmarshal() UnmarshalConfiger { return &JsonUnmarshal{} }
func NewYamlUnmarshal() UnmarshalConfiger { return &YamlUnmarshal{} }
func NewTomlUnmarshal() UnmarshalConfiger { return &TomlUnmarshal{} }

func (j *JsonUnmarshal) ReadConfig(p []byte, v any) error {
	return json.Unmarshal(p, v)
}

func (j *JsonUnmarshal) ReadConfigFromFile(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return j.ReadConfig(b, v)
}

func (y *YamlUnmarshal) ReadConfig(p []byte, v any) error {
	return yaml.Unmarshal(p, v)
}

func (y *YamlUnmarshal) ReadConfigFromFile(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return y.ReadConfig(b, v)
}

func (t *TomlUnmarshal) ReadConfig(p []byte, v any) error {
	return toml.Unmarshal(p, v)
}

func (t *TomlUnmarshal) ReadConfigFromFile(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return t.ReadConfig(b, v)
}
