package test

import (
	"reflect"
	"testing"

	Augo "github.com/0opslab/autngo"
)

func Test_error(t *testing.T) {
	resp := Augo.HttpHelper.HttpGet("https://studygolang.com/")
	print(reflect.TypeOf(resp).Name(), "==>", len(resp))
	print(Augo.Autn.ErrorMsg("res", nil))
	//print(Augo.)
}
