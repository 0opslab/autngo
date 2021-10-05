package test

import (
	"fmt"
	"testing"

	Autngo "github.com/0opslab/autngo"
)

func TestStrToGBK1(t *testing.T) {
	str := "hello 华夏10进制转36进制：";
	fmt.Println(Autngo.EncodeHelper.StrToGBK(str))
	fmt.Println(str)
	fmt.Println(Autngo.EncodeHelper.StrHex(str))
	hexSTT := "68656c6c6f20e58d8ee5a48f3130e8bf9be588b6e8bdac3336e8bf9be588b6efbc9a"
	str = Autngo.EncodeHelper.HexStr(hexSTT)
	fmt.Println(str)

}