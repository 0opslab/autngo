package test

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	Autngo "github.com/0opslab/autngo"
)

func TestWriteString(t *testing.T) {
	Autngo.FileHepler.WriteString("C:/Users/Administrator/Desktop/test1/test2/test.txt", "1111111111/n2222222222//33333")
	strs := "如果某个东西长得像鸭子，像鸭子一样游泳，像鸭子一样嘎嘎叫，那它就可以被看成是一只鸭子。"
	Autngo.FileHepler.WriteBytes("C:/Users/Administrator/Desktop/test1.txt", []byte(strs))
}

func TestReadByte(t *testing.T) {
	file := "C:/Windows/System32/drivers/etc/hosts"
	bytes, err := Autngo.FileHepler.ReadByteSize(file,8)
	if err != nil {
		fmt.Println("ReadError:" + err.Error())
	}
	fmt.Println("==============>")
	fmt.Println(bytes,"====>",Autngo.ByteHelper.BytesToHexString(bytes),"===>",string(bytes))
}


func TestReadFile(t *testing.T) {
	file := "C:/Windows/System32/drivers/etc/hosts"
	bytes, err := Autngo.FileHepler.ReadFile(file)
	if err != nil {
		fmt.Println("ReadError:" + err.Error())
	}
	fmt.Print(string(bytes))
}

func TestFileInfo(t *testing.T) {
	file := "C:/Windows/System32/drivers/etc/hosts"
	if Autngo.FileHepler.FileIsExist(file) {
		fileSize, _ := Autngo.FileHepler.FileSize(file)
		fileMd5, _ := Autngo.FileHepler.Md5File(file)
		fileSH1, _ := Autngo.FileHepler.Sha1File(file)
		fileSH2, _ := Autngo.FileHepler.Sha256File(file)
		fileSH5, _ := Autngo.FileHepler.Sha512File(file)

		fmt.Printf("FileInfo:fileSize=%v fileMd5=%s fileSH1=%s fileSH2=%s fileSH5=%s/n", fileSize, fileMd5, fileSH1, fileSH2, fileSH5)
	}

}

func TestWalkDirFiles(t *testing.T) {
	path := "c:/workspace/"
	files, dirs, _ := Autngo.FileHepler.WalkDirFiles(path, "java")
	for _, file := range files {
		fmt.Println(file)
	}
	for _, dir := range dirs {
		fmt.Println(dir)
	}
}

func TestWalkDirFilesFilter(t *testing.T) {
	files, _, _ := Autngo.FileHepler.WalkDirFilesFilter("c:/workspace/opslabJava", func(filename string) bool {
		fi, e := os.Stat(filename)
		if e != nil {
			return false
		}

		if (fi.IsDir()) {
			return false
		} else {
			if (path.Ext(filename) == ".java") {
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
		if(Autngo.FileHepler.IsFile(fileName)){
			md5,_ := Autngo.FileHepler.Md5File(fileName)
			fmt.Println(fileName,"==========>",md5)
		}
	}
	Autngo.FileHepler.WalkDirFilesHandler("C:/workspace/useful-documents",fileMd5)
	tts := time.Now().Unix()  - startnow1
	

	startnow := time.Now().Unix() 
	fileInfo := func(fileName string){
		isdir := 1
		mtime := int64(0)
		if(Autngo.FileHepler.IsFile(fileName)){
			isdir = 0
			mtime = Autngo.FileHepler.GetFileModTime(fileName)
		}
		fmt.Println("isdir:=",isdir," mtime:=",mtime," fileName:=",fileName)
	}
	Autngo.FileHepler.WalkDirFilesHandler("C:/workspace/useful-documents",fileInfo)
	
	fmt.Println("time2 ====>",time.Now().Unix()  - startnow)
	fmt.Println("time1 ====>",tts)

}

func TestCopy(t *testing.T) {
	if res,err := Autngo.FileHepler.CopyFile("C:/Windows/System32/drivers/etc/hosts","C:/Users/Administrator/Desktop/hosts"); err != nil{
		fmt.Println("CopyFile Error:",err)
	}else{
		fmt.Println("CopyFile status:",res)
	}

	if _, err := Autngo.FileHepler.CopyDir("C:/workspace/doc/", "C:/Users/Administrator/Desktop/doc/"); err != nil {
		fmt.Println(err)
	}
}


func TestZip(t *testing.T){
	if err := Autngo.FileHepler.ZipCompress("C:/workspace/doc","C:/Users/Administrator/Desktop/11.zip"); err != nil{
		fmt.Println("压缩文件错误:",err)
	}
	if err := Autngo.FileHepler.ZipUnCompress("C:/Users/Administrator/Desktop/11.zip","C:/Users/Administrator/Desktop/docs"); err != nil{
		fmt.Println("压缩文件错误:",err)
	}

}

func TestFileSat(t *testing.T){
	fileName := "C:\\temp\\dW9abvNBXK.txt"
	ll := Autngo.FileHepler.GetFileModTime(fileName)
	fmt.Println(ll)
}

