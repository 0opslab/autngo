package test

import (
	"fmt"
	"testing"

	Autngo "github.com/0opslab/autngo"
)

func Test_zip(t *testing.T) {
	if err := Autngo.ZipHelper.ZipCompress(WALKDIR_PATH, TEMP_FILEPATH+"/ZipCompress_test.zip"); err != nil {
		fmt.Println("压缩文件错误:", err)
	}
	if err := Autngo.ZipHelper.ZipUnCompress(TEMP_FILEPATH+"/ZipCompress_test.zip", TEMP_FILEPATH+"/unZipCompress"); err != nil {
		fmt.Println("压缩文件错误:", err)
	}
}
