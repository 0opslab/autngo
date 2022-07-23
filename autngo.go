package autngo

import (
	"github.com/0opslab/autngo/autn"
)

const (
	CST_TIME_RFC3339    = "2006-01-02T15:04:05+08:00"
	CST_TIME_TT         = "2006-01-02 15:04:05"
	CST_TIME_YMD        = "2006-01-02"
	CST_TIME_HMS        = "15:04:05"
	CST_UPPERCASE       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CST_LOWERCASE       = "abcdefghijklmnopqrstuvwxyz"
	CST_ALPHABETIC      = CST_UPPERCASE + CST_LOWERCASE
	CST_NUMERIC         = "0123456789"
	CST_ALPHANUMERIC    = CST_ALPHABETIC + CST_NUMERIC
	CST_ALPHANUMERICLOW = CST_UPPERCASE + CST_NUMERIC
	CST_SYMBOLS         = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	CST_Hex             = CST_NUMERIC + "abcdef"
	CST_MAXUINT32       = 4294967295
	CST_LOG_FORMAT      = "%v : %d - %s \n %s "
)

var (
	Autn         autn.Autn
	ByteHelper   autn.ByteHelper
	StringHelper autn.StringHelper
	HashHelper   autn.HashHelper
	RandomHelper autn.RandomHelper
	SliceHelper  autn.SliceHelper
	SysHelper    autn.SysHelper
	DateHelper   autn.DateHelper
	HttpHelper   autn.HttpHelper
	Response     autn.Response
	Random       autn.Random
	RsaHelper    autn.RsaHelper
	AesHelper    autn.AesHelper
	DesHelper    autn.DesHelper
	EncodeHelper autn.EncodeHelper
	FileHelper   autn.FileHelper
	ZipHelper    autn.ZipHelper
	BusError     autn.BusError
)
