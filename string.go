package rice

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/xid"
	"github.com/yitter/idgenerator-go/idgen"
)

func init() {
	options := idgen.NewIdGeneratorOptions(1)
	idgen.SetIdGenerator(options)
}

func NextId() string { return strconv.FormatUint(idgen.NextId(), 10) }

func XidNewString() string { return xid.New().String() }

// UUIDNewString creates a new random UUID
func UUIDNewString() string { return strings.ReplaceAll(uuid.NewString(), "-", "") }

// RandInt 生成一个真随机数
func RandInt(max int64) (int64, error) {
	if result, err := rand.Int(rand.Reader, big.NewInt(max)); err != nil {
		return 0, err
	} else {
		return result.Int64(), nil
	}
}
