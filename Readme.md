### autngo
I don't know shit. It's just some damn code you use all the time


### env
vscode + remote ssh
vsocde 按照go插件、并通ctrl+shfit+p输入：o:install/update tools 并构想全部插件安装

### cmd

```bash
git add .
git commit -m "readme"
git tag v1.0.7
git push origin master --tags


go test .\test\ByteHelper_test.go
go test -v .\test\ByteHelper_test.go
```

### 使用

写入文件
```go
import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	Augo "github.com/0opslab/autngo"
)

// for windows
// const TEMP_WRITE_FILE_NAME = "C:/Users/Administrator/Desktop/test1/test2/test.txt"
// const TEMP_WRITE_FILE_NAME2 = "C:/Users/Administrator/Desktop/test1/test2/test.txt"

// for linux
// 测试数据目录：/data/workspace/test-data
const TEMP_WRITE_FILE_NAME = "/data/workspace/test-data/temp/go-test-WriteString.txt"

func TestWriteString(t *testing.T) {
	Augo.FileHelper.WriteString(TEMP_WRITE_FILE_NAME, "1111111111/n2222222222//33333")
	strs := "如果某个东西长得像鸭子，像鸭子一样游泳，像鸭子一样嘎嘎叫，那它就可以被看成是一只鸭子。"
	Augo.FileHelper.WriteBytes(TEMP_WRITE_FILE_NAME2, []byte(strs))
}

//获取文件信息
func TestFileInfo(t *testing.T) {
	file := TEMP_WRITE_FILE_NAME2
	if Augo.FileHelper.FileIsExist(file) {
		fileSize, _ := Augo.FileHelper.FileSize(file)
		fileMd5, _ := Augo.FileHelper.Md5File(file)
		fileSH1, _ := Augo.FileHelper.Sha1File(file)
		fileSH2, _ := Augo.FileHelper.Sha256File(file)
		fileSH5, _ := Augo.FileHelper.Sha512File(file)

		fmt.Printf("FileInfo:fileSize=%v fileMd5=%s fileSH1=%s fileSH2=%s fileSH5=%s/n",
			fileSize, fileMd5, fileSH1, fileSH2, fileSH5)
	}

}

//遍历文件
func TestWalkDirFilesFilter(t *testing.T) {
	files, _, _ := Augo.FileHelper.WalkDirFilesFilter(WALKDIR_PATH, func(filename string) bool {
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
```
