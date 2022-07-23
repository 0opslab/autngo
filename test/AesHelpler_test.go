package test

import (
	"fmt"
	"testing"

	Augo "github.com/0opslab/autngo"
)

func Test_Aes(t *testing.T) {
	strs := "TestString测试"
	iv := "aeGksHkG4jAEk2Ag"
	cip, err := Augo.AesHelper.GetCipher("aGcBfWb3Kg2s4gcG", iv)
	if err != nil {
		fmt.Println(err)
	}

	res1, err1 := Augo.AesHelper.AesBase64Encrypt(cip, iv, strs)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(res1, "==>", len(res1))

	res, err := Augo.AesHelper.AesBase64Decrypt(cip, iv, res1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res, "==>", len(res))
}
