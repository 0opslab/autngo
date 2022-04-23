package sys

//封装与操作系统相关的一些操作
type WinHelper struct {
}

// windows获取系统盘符
func (ss *WinHelper) GetLogicalDrives() []string {
	// if runtime.GOOS == "windows" {
	// 	// windows系统
	// 	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	// 	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	// 	n, _, _ := GetLogicalDrives.Call()
	// 	s := strconv.FormatInt(int64(n), 2)
	// 	var drives_all = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:",
	// 		"J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：", "U：",
	// 		"V：", "W：", "X：", "Y：", "Z："}
	// 	temp := drives_all[0:len(s)]
	// 	var d []string
	// 	for i, v := range s {
	// 		if v == 49 {
	// 			l := len(s) - i - 1
	// 			d = append(d, temp[l])
	// 		}
	// 	}
	// 	var drives []string
	// 	for i, v := range d {
	// 		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
	// 	}
	// 	return drives
	// }
	return nil

}
