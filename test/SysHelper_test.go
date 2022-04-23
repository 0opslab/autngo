package test

import (
	"fmt"
	"testing"

	Autngo "github.com/0opslab/autngo"
)

func TestGetLogicalDrives_Windows(t *testing.T) {
	Autngo.SysHelper.OsInfo()
	ip := Autngo.SysHelper.GetLocalIP()
	fmt.Print(ip)
}
