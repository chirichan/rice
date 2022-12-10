package rice

import (
	"crypto/rand"
	"errors"
	"math/big"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/xid"
	"github.com/yitter/idgenerator-go/idgen"
)

func InitYitterID(workID uint16) {
	options := idgen.NewIdGeneratorOptions(workID)
	idgen.SetIdGenerator(options)
}

type YitterID struct{}



type Ider interface {
	NextInt64Id() int64
	NextStringId() string
}

type YitId struct{}
type UUID struct{}
type XId struct{}

func NewYitId() Ider { return &YitId{} }
func NewUUID() Ider  { return &UUID{} }
func NewXId() Ider   { return &XId{} }

func (n *YitId) NextInt64Id() int64 {
	return idgen.NextId()
}

func (n *YitId) NextStringId() string {
	return strconv.FormatInt(idgen.NextId(), 10)
}

func (u *UUID) NextStringId() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

func (u *UUID) NextInt64Id() int64 {
	panic("impl me")
}

func (x *XId) NextStringId() string {
	return xid.New().String()
}

func (x *XId) NextInt64Id() int64 {
	panic("impl me")
}

func NextStringId() string { return strconv.FormatInt(idgen.NextId(), 10) }

// RandInt 生成一个真随机数
func RandInt(max int64) (int64, error) {
	if result, err := rand.Int(rand.Reader, big.NewInt(max)); err != nil {
		return 0, err
	} else {
		return result.Int64(), nil
	}
}

func GetLocalHostname() (string, error) {
	return os.Hostname()
}

func GetLocalAddr() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	s := conn.LocalAddr().String()
	i := strings.LastIndex(s, ":")
	if i == -1 {
		return "", errors.New("can't get local addr")
	}
	return s[:i], nil
}
