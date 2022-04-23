package test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	Autngo "github.com/0opslab/autngo"
)

// 遍历指定文件夹并计算其中文件的MD5进行输出
const TEST_DIR_PATH = "/data/workspace/opslabPython"

func TestWalkDirFilesHandler2(t *testing.T) {

	startnow := time.Now().Unix()
	fileInfo := func(fileName string) string {
		isdir := 1
		mtime := int64(0)
		fileSize := int64(0)
		if Autngo.FileHepler.IsFile(fileName) {
			isdir = 0
			mtime = Autngo.FileHepler.GetFileModTime(fileName)
			fileSize, _ = Autngo.FileHepler.FileSize(fileName)
		}
		result := "{\"isdir\":\"%v\",\"mtime\":\"%v\",\"fileName\":\"%v\",\"fileSize\":\"%v\"}"
		s := fmt.Sprintf(result, isdir, mtime, fileName, fileSize)
		return s
	}

	dirPath := TEST_DIR_PATH
	var ss []string
	filepath.Walk(dirPath, func(filename string, fi os.FileInfo, err error) error {
		info := fileInfo(filename)
		fmt.Println(info)
		ss = append(ss, info)
		return nil
	})

	fmt.Println("time2 ====>", time.Now().Unix()-startnow)
	data, err := json.Marshal(ss)
	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return
	}
	s1 := ss[1:]
	data1, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
	fmt.Println("===============>")
	fmt.Printf("%s\n", string(data1))
	fmt.Println("===============>")
	s2 := Autngo.SliceHelper.Difference(ss, s1)
	data2, err := json.Marshal(s2)
	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return
	}
	fmt.Printf("%s\n", string(data2))
}
