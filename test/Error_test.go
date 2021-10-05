package test

import (
	"testing"

	Autngo "github.com/0opslab/autngo"
)

func Test_error(t *testing.T) {
	errors := Autngo.ComError.ComError("TestString测试")
	print("error",errors)
}