package test

import (
	"fmt"
	"testing"
	"github.com/0opslab/autngo"
)

func TestMd5Byte(t *testing.T) {
	data := []byte("TestString测试")
	fmt.Println(autngo.ByteHelper.Md5Byte(data))
}