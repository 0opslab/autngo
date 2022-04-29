package autn

import (
	"fmt"
	"testing"
)

func TestMd5Byte(t *testing.T) {
	var help ByteHelper
	data := []byte("TestString测试")
	fmt.Println(help.Md5Byte(data))
}
