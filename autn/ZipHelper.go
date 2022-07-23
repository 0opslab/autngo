package autn

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//@Document 文件相关的操作集合
type ZipHelper struct {
	fs FileHelper
}

// 将指定的目录或文件压缩到指定的zip中
func (ff *ZipHelper) ZipCompress(srcFile string, destZip string) error {

	if ff.fs.FileIsExist(destZip) {
		return ErrorMsg("FileISExist:"+destZip, nil)
	}

	if zipfile, err := os.Create(destZip); err != nil {
		return err
	} else {
		defer zipfile.Close()

		archive := zip.NewWriter(zipfile)
		defer archive.Close()

		err2 := filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}
			header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile))
			if info.IsDir() {
				if header.Name == srcFile {
					return nil
				}
				header.Name += "/"
			} else {
				header.Method = zip.Deflate
			}

			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}

			if !info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()
				_, err = io.Copy(writer, file)
				if err != nil {
					return err
				}
			}
			return err
		})

		return err2
	}

}

// 将指定的zip的文件压缩到指定的目录下
func (ff *ZipHelper) ZipUnCompress(zipFile string, destDir string) error {
	if ff.fs.FileIsExist(destDir) {
		return ErrorMsg("The file already exists:"+destDir, nil)
	}
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
