package autn

import (
	"fmt"
	"time"
)

type BusError struct {
	time  time.Time
	code  int
	msg   string
	stack string
}

func (m *BusError) Error() string {
	return fmt.Sprintf(CST_LOG_FORMAT, m.time, m.code, m.msg, m.stack)
}
