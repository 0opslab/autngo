package test

import (
	"encoding/base64"
	"fmt"
	"testing"

	Augo "github.com/0opslab/autngo"
)

func Test_Des(t *testing.T) {
	key := []byte("12345678")
	src := []byte("这是需要加密的明文哦！")

	// ECB 加密
	cipher := Augo.DesHelper.DesECBEncrypt(src, key)

	// 转base64
	bs64 := base64.StdEncoding.EncodeToString(cipher)
	fmt.Println(bs64)

	// 转回byte
	bt, err := base64.StdEncoding.DecodeString(bs64)
	if err != nil {
		fmt.Println("base64转换失败")
	}
	// ECB 解密
	str := Augo.DesHelper.DesECBDecrypter(bt, key)
	fmt.Println(string(str))
}
