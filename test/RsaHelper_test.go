package test

import (
	"fmt"
	"strings"
	"testing"

	Autngo "github.com/0opslab/autngo"
)

func TestRsaEncryptWithPublic(t *testing.T) {
	_, e := Autngo.RsaHelper.RsaEncryptFileWithPublic("c:/ReadMe.md", "c:/1.data")
	if e != nil {
		fmt.Println(e)
	}

	_, e1 := Autngo.RsaHelper.RsaDecryptFileWithPrivte("c:/1.data", "c:/ReadMe1.md")
	if e1 != nil {
		fmt.Println(e1)
	}
}

func TestRsaDecryptWithPrivte(t *testing.T) {
	rsaInfo := "rsainfo"
	rsaPath := "D:\\test\\rsafile"
	fileMd5 := func(fileName string) {
		if Autngo.FileHelper.IsFile(fileName) {
			sha, _ := Autngo.FileHelper.Sha256File(fileName)

			rsaFile := rsaPath + "\\" + sha
			dstFile := strings.ReplaceAll(strings.ReplaceAll(fileName, "C:\\workspace\\", "D:\\test\\dst\\"), "\\", "/")
			rest := fmt.Sprintf("%s -> %s -> %s\n", fileName, rsaFile, dstFile)
			Autngo.FileHelper.WriteString(rsaPath+"\\"+rsaInfo, rest)
			fmt.Print(rest)
			go func() {
				_, e := Autngo.RsaHelper.RsaEncryptFileWithPublic(fileName, rsaFile)
				if e != nil {
					fmt.Println(e)
				}
				_, e1 := Autngo.RsaHelper.RsaDecryptFileWithPrivte(rsaFile, dstFile)
				if e1 != nil {
					fmt.Println(e1)
				}
			}()

		}

	}
	Autngo.FileHelper.WalkDirFilesHandler("C:\\workspace\\", fileMd5)
}
