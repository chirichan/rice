package rice

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/xid"
	"github.com/yitter/idgenerator-go/idgen"
)

func InitYitterID(workID uint16) {
	options := idgen.NewIdGeneratorOptions(workID)
	options.BaseTime = time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local).UnixMilli()
	idgen.SetIdGenerator(options)
}

func NextYitterID() int64 { return idgen.NextId() }

func NextYitterStringID() string { return strconv.FormatInt(idgen.NextId(), 10) }

func NextUUID() string { return strings.ReplaceAll(uuid.NewString(), "-", "") }

func NextXID() string { return xid.New().String() }
