package test

import (
	"fmt"
	"testing"

	Autngo "github.com/0opslab/autngo"
)

// 遍历指定文件夹并计算其中文件的MD5进行输出
func TestSclieStr(t *testing.T) {
	slice1 := []string{"1", "2", "3", "6", "8"}
	slice2 := []string{"2", "3", "5", "0"}

	un := Autngo.SliceHelper.Union(slice1,slice2)
	fmt.Println(un)

	in := Autngo.SliceHelper.Intersect(slice1, slice2)
	fmt.Println("slice1与slice2的交集为：", in)

	di := Autngo.SliceHelper.Difference(slice1, slice2)
	fmt.Println("slice1与slice2的差集为：", di)
}

type UserInfo struct{
	Name string
	Age int
}