package test

import (
	"fmt"
	"testing"

	Autngo "github.com/0opslab/autngo"
)

func TestGetLogicalDrives_Windows(t *testing.T) {
	drivers  := Autngo.SysHelper.GetLogicalDrives()
	fmt.Println(drivers)
	print(Autngo.SysHelper.GetLocalIP())
}
