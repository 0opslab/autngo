package test

import (
	"fmt"
	"testing"

	Autngo "github.com/0opslab/autngo"
)

func TestMd5Byte(t *testing.T) {
	data := []byte("TestString测试")
	fmt.Println(Autngo.ByteHelper.Md5Byte(data))
}
