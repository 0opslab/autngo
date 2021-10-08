package autngo

import (
	"github.com/0opslab/autngo/byte"
	"github.com/0opslab/autngo/date"
	"github.com/0opslab/autngo/encode"

	"github.com/0opslab/autngo/common"
	"github.com/0opslab/autngo/encrypt"
	"github.com/0opslab/autngo/file"
	"github.com/0opslab/autngo/random"
	sclie "github.com/0opslab/autngo/slice"
	"github.com/0opslab/autngo/string"
	"github.com/0opslab/autngo/sys"
)


var (
	ByteHelper byte.ByteHelper
	DateHelper date.DateHelper
	EncodeHelper encode.EncodeHelper
	EncryptHelper encrypt.EncryptHelper
	ComError common.ComError
	FileHepler file.FileHelper
	SysHelper sys.SysHelper
	StringHelper string.StringHelper
	Random random.Random
	RandomHelper random.RandomHelper
	SliceHelper sclie.SliceHelper
)