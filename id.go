package rice

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jaevor/go-nanoid"
	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"
	"github.com/yitter/idgenerator-go/idgen"
)

var YitterBaseTime = time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local)

func InitYitterID(workID uint16) {
	options := idgen.NewIdGeneratorOptions(workID)
	options.BaseTime = YitterBaseTime.UnixMilli()
	idgen.SetIdGenerator(options)
}

func NextYitterID() int64 { return idgen.NextId() }

func NextYitterStringID() string { return strconv.FormatInt(idgen.NextId(), 10) }

func NextUUID() string { return strings.ReplaceAll(uuid.NewString(), "-", "") }

func NextXID() string { return xid.New().String() }

func NextNanoID() string {
	gen, _ := nanoid.Canonic()
	return gen()
}

func NextULID() string {
	return ulid.Make().String()
}
