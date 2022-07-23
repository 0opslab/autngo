package autn

import (
	"math/rand"
	"strings"
	"time"
)

type RandomHelper struct {
}

//从给定的字符串中随机指定长度的字符串
func (ss *RandomHelper) String(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = CST_ALPHANUMERIC
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(len(charset))]
	}
	return string(b)
}

//随机指定长度的字符串(小写字母和数字)
func (ss *RandomHelper) RandomString(length uint8) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = CST_ALPHANUMERICLOW[rand.Int63()%int64(len(CST_ALPHANUMERICLOW))]
	}
	return string(b)
}

//随机指定范围的整数(伪随机数，其实一般用也没错)
func (ss *RandomHelper) RandomInt(length int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(length)
}
