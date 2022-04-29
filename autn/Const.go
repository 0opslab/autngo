package autn

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
