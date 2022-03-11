package rice

import (
	"strconv"
	"sync"

	"github.com/yitter/idgenerator-go/idgen"
)

var (
	options   = idgen.NewIdGeneratorOptions(1)
	idgenOnce sync.Once
)

func init() {
	idgenOnce.Do(func() {
		idgen.SetIdGenerator(options)
	})
}

func NextId() string { return strconv.FormatUint(idgen.NextId(), 10) }
