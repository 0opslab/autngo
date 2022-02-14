package common

import (
	"fmt"
	"runtime/debug"
	"time"
)

type ComError struct {
	time  time.Time
	code  int
	msg   string
	stack string
}

const logFormat = "%v : %d - %s \n %s "

func (m *ComError) Error() string {
	return fmt.Sprintf(logFormat, m.time, m.code, m.msg, m.stack)
}

// 根据消息生成错误消息
func (m *ComError) ComError(strMsg string) *ComError {
	return &ComError{
		time:  time.Now(),
		code:  0,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}

// 根据信息包装指定的错误消息
func (m *ComError) ComWithError(strMsg string, err error) *ComError {
	if err != nil {
		strMsg += " " + err.Error()
	}
	return &ComError{
		time:  time.Now(),
		code:  0,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}

// 根据错误编码和消息生成错误消息
func (m *ComError) ComErrorCode(i int, strMsg string) *ComError {
	return &ComError{
		time:  time.Now(),
		code:  i,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}

// 根据错误编码和消息包装指定错误消息
func (m *ComError) ComWithErrorCode(i int, strMsg string, err error) *ComError {
	if err != nil {
		strMsg += " " + err.Error()
	}
	return &ComError{
		time:  time.Now(),
		code:  i,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}
