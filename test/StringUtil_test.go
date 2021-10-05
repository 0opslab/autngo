package test

import (
	"fmt"
	"testing"

	Autngo "github.com/0opslab/autngo"
)

func TestBase64(t *testing.T) {
	strs := "this is测试字符串";
	base64 := Autngo.StringHelper.Base64Encode(strs)
	strs1 := Autngo.StringHelper.Base64Decode(base64)

	fmt.Println(base64)
	fmt.Println(strs1)

	fmt.Println(Autngo.StringHelper.Reverse(strs))
}