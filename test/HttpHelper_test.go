package test

import (
	"testing"

	Augo "github.com/0opslab/autngo"
)

func Test_http(t *testing.T) {
	errors := Augo.Autn.ErrorMsg("TestString测试", nil)
	print("error", errors)
}
