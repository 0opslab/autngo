package autn

import (
	"runtime/debug"
	"time"
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

type Autn struct {
}

// 根据信息包装指定的错误消息
func (m *Autn) ErrorMsg(strMsg string, err error) error {
	if err != nil {
		strMsg += " " + err.Error()
	}
	return &BusError{
		time:  time.Now(),
		code:  0,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}

// 根据错误编码和消息生成错误消息
func (m *Autn) ErrorCode(i int, strMsg string) error {
	return &BusError{
		time:  time.Now(),
		code:  i,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}

// 根据错误编码和消息包装指定错误消息
func (m *Autn) ErrorCodeMsg(i int, strMsg string, err error) error {
	if err != nil {
		strMsg += " " + err.Error()
	}
	return &BusError{
		time:  time.Now(),
		code:  i,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}

// 根据消息生成错误消息
// func Error(strMsg string) error {
// 	return &BusError{
// 		time:  time.Now(),
// 		code:  0,
// 		msg:   strMsg,
// 		stack: string(debug.Stack()),
// 	}
// }

// 根据信息包装指定的错误消息
func ErrorMsg(strMsg string, err error) error {
	if err != nil {
		strMsg += " " + err.Error()
	}
	return &BusError{
		time:  time.Now(),
		code:  0,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}

// 根据错误编码和消息生成错误消息
func ErrorCode(i int, strMsg string) error {
	return &BusError{
		time:  time.Now(),
		code:  i,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}

// 根据错误编码和消息包装指定错误消息
func ErrorCodeMsg(i int, strMsg string, err error) error {
	if err != nil {
		strMsg += " " + err.Error()
	}
	return &BusError{
		time:  time.Now(),
		code:  i,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}
