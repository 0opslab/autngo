package test

import (
	"fmt"
	"strings"
	"testing"

	Autngo "github.com/0opslab/autngo"
)

func TestRsaEncryptWithPublic(t *testing.T) {
	_, e := Autngo.EncryptHelper.RsaEncryptFileWithPublic("c:/ReadMe.md", "c:/1.data")
	if e != nil{
		fmt.Println(e)
	}

	_, e1 := Autngo.EncryptHelper.RsaDecryptFileWithPrivte("c:/1.data","c:/ReadMe1.md")
	if e1 != nil{
		fmt.Println(e1)
	}
}

func TestRsaDecryptWithPrivte(t *testing.T) {
	rsaInfo := "rsainfo"
	rsaPath := "D:\\test\\rsafile"
	fileMd5 := func(fileName string) {
		if(Autngo.FileHepler.IsFile(fileName)){
			sha,_ := Autngo.FileHepler.Sha256File(fileName)

			rsaFile := rsaPath+"\\"+sha
			dstFile := strings.ReplaceAll(strings.ReplaceAll(fileName,"C:\\workspace\\","D:\\test\\dst\\"),"\\","/")
			rest := fmt.Sprintf("%s -> %s -> %s\n",fileName,rsaFile,dstFile)
			Autngo.FileHepler.WriteString(rsaPath+"\\"+rsaInfo,rest)
			fmt.Print(rest)
			go func(){
				_, e := Autngo.EncryptHelper.RsaEncryptFileWithPublic(fileName,rsaFile)
				if e != nil{
					fmt.Println(e)
				}
				_, e1 := Autngo.EncryptHelper.RsaDecryptFileWithPrivte(rsaFile,dstFile)
				if e1 != nil{
					fmt.Println(e1)
				}
			}()

		}

	}
	Autngo.FileHepler.WalkDirFilesHandler("C:\\workspace\\",fileMd5)
}
