package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type RandomHelper struct{

}
const (
	Uppercase              = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase              = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic             = Uppercase + Lowercase
	Numeric                = "0123456789"
	Alphanumeric           = Alphabetic + Numeric
	AlphanumericLow        = Uppercase + Numeric
	Symbols                = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	Hex                    = Numeric + "abcdef"
	MAXUINT32              = 4294967295
	DEFAULT_UUID_CNT_CACHE = 512
)

type Random struct {
	Prefix       string
	idGen        uint32
	internalChan chan uint32
}

func (ss *RandomHelper) NewRandom() *Random {
	rand.Seed(time.Now().UnixNano())
	random := &Random{
		Prefix:       "",
		idGen:        0,
		internalChan: make(chan uint32, DEFAULT_UUID_CNT_CACHE),
	}
	random.startGen()
	return random
}

func (ss *RandomHelper)  NewRandomGen(prefix string, startValue uint32) *Random {
	rand.Seed(time.Now().UnixNano())
	random := &Random{
		Prefix:       prefix,
		idGen:        startValue,
		internalChan: make(chan uint32, DEFAULT_UUID_CNT_CACHE),
	}
	random.startGen()
	return random
}

//开启 goroutine, 把生成的数字形式的UUID放入缓冲管道
func (this *Random) startGen() {
	go func() {
		for {
			if this.idGen == MAXUINT32 {
				this.idGen = 1
			} else {
				this.idGen += 1
			}
			this.internalChan <- this.idGen
		}
	}()
}

//获取带前缀的字符串形式的UUID
func (this *Random) GetUUID() string {
	idgen := <-this.internalChan
	return fmt.Sprintf("%s%d", this.Prefix, idgen)
}

//获取uint32形式的UUID
func (this *Random) GetUint32() uint32 {
	return <-this.internalChan
}

//从给定的字符串中随机指定长度的字符串
func (ss *RandomHelper)  String(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
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
		b[i] = AlphanumericLow[rand.Int63()%int64(len(AlphanumericLow))]
	}
	return string(b)
}

//随机指定范围的整数(伪随机数，其实一般用也没错)
func (ss *RandomHelper) RandomInt(length int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100)
}

//生成Guid字串
// func Guid() string {
// 	b := make([]byte, 48)

// 	if _, err := io.ReadFull(crand.Reader, b); err != nil {
// 		return ""
// 	}

// 	//return Md5String(base64.URLEncoding.EncodeToString(b))
// }
