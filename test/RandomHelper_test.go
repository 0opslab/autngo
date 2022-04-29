package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	Autngo "github.com/0opslab/autngo"
)

func TestRandom_GoRandomString(t *testing.T) {
	//系统自动的随机数正确使用方法
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
		fmt.Println(vcode)
	}
}

func TestRandom_String(t *testing.T) {
	//封装的随机数方法使用
	rand := Autngo.RandomHelper.NewRandom()
	for i := 0; i < 10; i++ {
		//注意此处是有序的 而不是随机的
		fmt.Println(rand.GetUUID())
		fmt.Println(rand.GetUint32())

		fmt.Println(Autngo.RandomHelper.String(8, Autngo.Numeric))
		fmt.Println(Autngo.RandomHelper.String(32, Autngo.Alphanumeric))
		fmt.Println(Autngo.RandomHelper.RandomString(10))
		fmt.Println(Autngo.RandomHelper.RandomInt(10))
		fmt.Println("===")
	}
}
