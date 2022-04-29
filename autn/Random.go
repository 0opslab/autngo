package autn

import (
	"fmt"
	"math/rand"
	"time"
)

const (
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

func (ss *RandomHelper) NewRandomGen(prefix string, startValue uint32) *Random {
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
func (tt *Random) startGen() {
	go func() {
		for {
			if tt.idGen == MAXUINT32 {
				tt.idGen = 1
			} else {
				tt.idGen += 1
			}
			tt.internalChan <- tt.idGen
		}
	}()
}

//获取带前缀的字符串形式的UUID
func (tt *Random) GetUUID() string {
	idgen := <-tt.internalChan
	return fmt.Sprintf("%s%d", tt.Prefix, idgen)
}

//获取uint32形式的UUID
func (tt *Random) GetUint32() uint32 {
	return <-tt.internalChan
}
