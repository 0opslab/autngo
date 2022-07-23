package test

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	Autngo "github.com/0opslab/autngo"
)

// for windows
// const TEMP_WRITE_FILE_NAME = "C:/Users/Administrator/Desktop/test1/test2/test.txt"
// const TEMP_WRITE_FILE_NAME2 = "C:/Users/Administrator/Desktop/test1/test2/test.txt"

// for linux
// 测试数据目录：/data/workspace/test-data
const TEMP_WRITE_FILE_NAME = "/data/workspace/test-data/temp/go-test-WriteString.txt"
const TEMP_WRITE_FILE_NAME_temp = "/data/workspace/test-data/temp/go-test-WriteString1.txt"
const TEMP_WRITE_FILE_NAME2 = "/data/workspace/test-data/temp/go-test-WriteBytes.txt"
const TEMP_CREATEDIR_FILE_NAME2 = "/data/workspace/test-data/temp/testdir/go-test-WriteBytes.txt"
const WALKDIR_PATH = "/data/workspace/opslabPython"
const TEMP_FILEPATH = "/data/workspace/test-data/temp"

func TestWriteString(t *testing.T) {
	Autngo.FileHelper.WriteString(TEMP_WRITE_FILE_NAME, "1111111111/n2222222222//33333")
	strs := "如果某个东西长得像鸭子，像鸭子一样游泳，像鸭子一样嘎嘎叫，那它就可以被看成是一只鸭子。"
	Autngo.FileHelper.WriteBytes(TEMP_WRITE_FILE_NAME2, []byte(strs))
}

func TestReadByte(t *testing.T) {
	file := TEMP_WRITE_FILE_NAME2
	bytes, err := Autngo.FileHelper.ReadByteSize(file, 8)
	if err != nil {
		fmt.Println("ReadError:" + err.Error())
	}
	fmt.Println("==============>")
	fmt.Println(bytes, "====>", Autngo.ByteHelper.BytesToHexString(bytes), "===>", string(bytes))
	rs, _ := Autngo.FileHelper.MakeDir(TEMP_CREATEDIR_FILE_NAME2)
	fmt.Println("===rs:", rs)
}

func TestReadFile(t *testing.T) {
	file := TEMP_WRITE_FILE_NAME2
	bytes, err := Autngo.FileHelper.ReadFile(file)
	if err != nil {
		fmt.Println("ReadError:" + err.Error())
	}
	fmt.Print(string(bytes))
}

func TestFileSat(t *testing.T) {
	fileName := TEMP_WRITE_FILE_NAME
	ll := Autngo.FileHelper.GetFileModTime(fileName)
	fmt.Println(ll)
}

func TestFileInfo(t *testing.T) {
	file := TEMP_WRITE_FILE_NAME2
	if Autngo.FileHelper.FileIsExist(file) {
		fileSize, _ := Autngo.FileHelper.FileSize(file)
		fileMd5, _ := Autngo.FileHelper.Md5File(file)
		fileSH1, _ := Autngo.FileHelper.Sha1File(file)
		fileSH2, _ := Autngo.FileHelper.Sha256File(file)
		fileSH5, _ := Autngo.FileHelper.Sha512File(file)

		fmt.Printf("FileInfo:fileSize=%v fileMd5=%s fileSH1=%s fileSH2=%s fileSH5=%s/n",
			fileSize, fileMd5, fileSH1, fileSH2, fileSH5)
	}

}

func TestWalkDirFiles(t *testing.T) {
	path := WALKDIR_PATH
	files, dirs, _ := Autngo.FileHelper.WalkDirFiles(path, "py")
	for _, file := range files {
		fmt.Println(file)
	}
	for _, dir := range dirs {
		fmt.Println(dir)
	}
}

func TestWalkDirFilesFilter(t *testing.T) {
	files, _, _ := Autngo.FileHelper.WalkDirFilesFilter(WALKDIR_PATH, func(filename string) bool {
		fi, e := os.Stat(filename)
		if e != nil {
			return false
		}

		if fi.IsDir() {
			return false
		} else {
			if path.Ext(filename) == ".py" {
				return true
			}
		}
		return false
	})
	for _, file := range files {
		fmt.Println(file)
	}
}

// 遍历指定文件夹并计算其中文件的MD5进行输出
func TestWalkDirFilesHandler(t *testing.T) {
	startnow1 := time.Now().Unix()
	fileMd5 := func(fileName string) {
		if Autngo.FileHelper.IsFile(fileName) {
			md5, _ := Autngo.FileHelper.Md5File(fileName)
			fmt.Println(fileName, "==========>", md5)
		}
	}
	Autngo.FileHelper.WalkDirFilesHandler(WALKDIR_PATH, fileMd5)
	tts := time.Now().Unix() - startnow1

	startnow := time.Now().Unix()
	fileInfo := func(fileName string) {
		isdir := 1
		mtime := int64(0)
		if Autngo.FileHelper.IsFile(fileName) {
			isdir = 0
			mtime = Autngo.FileHelper.GetFileModTime(fileName)
		}
		fmt.Println("isdir:=", isdir, " mtime:=", mtime, " fileName:=", fileName)
	}
	Autngo.FileHelper.WalkDirFilesHandler(WALKDIR_PATH, fileInfo)

	fmt.Println("time2 ====>", time.Now().Unix()-startnow)
	fmt.Println("time1 ====>", tts)

}

func TestCopy(t *testing.T) {
	if res, err := Autngo.FileHelper.CopyFile(TEMP_WRITE_FILE_NAME, TEMP_WRITE_FILE_NAME_temp); err != nil {
		fmt.Println("CopyFile Error:", err)
	} else {
		fmt.Println("CopyFile status:", res)
	}

	if _, err := Autngo.FileHelper.CopyDir(WALKDIR_PATH, TEMP_FILEPATH+"/copydir/"); err != nil {
		fmt.Println(err)
	}
}
