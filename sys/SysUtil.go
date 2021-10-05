package sys

import (
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

//封装与操作系统相关的一些操作
type SysHelper struct{

}

//获取系统主要信息(主要用于临时查看，不用过度封装)
func (ss *SysHelper) OsInfo(){
	println(`系统类型：`, runtime.GOOS)

	println(`系统架构：`, runtime.GOARCH)

	println(`CPU 核数：`, runtime.GOMAXPROCS(0))

	name, err := os.Hostname()
	if err != nil {
			panic(err)
	}
	println(`电脑名称：`, name)
}
// windows获取系统盘符
func  (ss *SysHelper) GetLogicalDrives() []string {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	n, _, _ := GetLogicalDrives.Call()
	s := strconv.FormatInt(int64(n), 2)
	var drives_all = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:",
		"J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：", "U：",
		"V：", "W：", "X：", "Y：", "Z："}
	temp := drives_all[0:len(s)]
	var d []string
	for i, v := range s {
		if v == 49 {
			l := len(s) - i - 1
			d = append(d, temp[l])
		}
	}
	var drives []string
	for i, v := range d {
		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
	}
	return drives
}

/**
获取程序当前所在目录
*/
func  (ss *SysHelper) GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}


// GetLocalIP returns the non loopback local IP of the host
func  (ss *SysHelper) GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
			return ""
	}
	for _, address := range addrs {
			// check the address type and if it is not a loopback the display it
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
							return ipnet.IP.String()
					}
			}
	}
	return ""
}