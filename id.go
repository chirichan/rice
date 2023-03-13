package rice

import (
	"strconv"
	"time"

	"github.com/yitter/idgenerator-go/idgen"
)

func InitID(workID uint16) {
	options := idgen.NewIdGeneratorOptions(workID)
	options.BaseTime = time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local).UnixMilli()
	idgen.SetIdGenerator(options)
}

func NextID() int64 { return idgen.NextId() }

func NextStringID() string { return strconv.FormatInt(idgen.NextId(), 10) }
