package rice

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"os"
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

func NextStringId() string { return strconv.FormatInt(idgen.NextId(), 10) }

func NextInt64Id() int64 { return idgen.NextId() }

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

func GetHostname() string {

	hostname, _ := os.Hostname()

	fmt.Printf("hostname: %v\n", hostname)

	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()

	fmt.Printf("conn.LocalAddr().String(): %v\n", conn.LocalAddr().String())

	return ""
}
