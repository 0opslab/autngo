package autngo

import (
	"github.com/0opslab/autngo/autn"
	"github.com/0opslab/autngo/encode"

	"github.com/0opslab/autngo/common"
	"github.com/0opslab/autngo/encrypt"
	"github.com/0opslab/autngo/file"
	"github.com/0opslab/autngo/http"
)

const (
	RFC3339         = "2006-01-02T15:04:05+08:00"
	TT              = "2006-01-02 15:04:05"
	YMD             = "2006-01-02"
	HMS             = "15:04:05"
	Uppercase       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase       = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic      = Uppercase + Lowercase
	Numeric         = "0123456789"
	Alphanumeric    = Alphabetic + Numeric
	AlphanumericLow = Uppercase + Numeric
	Symbols         = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	Hex             = Numeric + "abcdef"
	MAXUINT32       = 4294967295
)

var (
	ByteHelper   autn.ByteHelper
	StringHelper autn.StringHelper
	HashHelper   autn.HashHelper
	RandomHelper autn.RandomHelper
	SliceHelper  autn.SliceHelper
	SysHelper    autn.SysHelper
	DateHelper   autn.DateHelper

	Random autn.Random

	EncodeHelper  encode.EncodeHelper
	EncryptHelper encrypt.EncryptHelper
	ComError      common.ComError
	FileHepler    file.FileHelper
	HttpHelper    http.HttpHelper
	Response      http.Response
)
