package file

import (
	"archive/zip"
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/0opslab/autngo/common"
)

//@Document 文件相关的操作集合
type FileHelper struct{
	errors common.ComError
}

//创建文件，如果存在直接返回
func (ff *FileHelper) CreateFile(filepath string)(bool,error){
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		destdir, _ := path.Split(filepath)
		os.MkdirAll(destdir,os.ModePerm)
		file,err:=os.Create(filepath)
		if err!=nil{
			return false,err
		}
		defer file.Close()
		return true,nil
	}
	return true,nil
}

//打开一个文件支持多目录创建
func  (ff *FileHelper) CreateOpenFile(filepath string)(*os.File,error){
	if ff.FileIsExist(filepath){
		 return  os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	}
	destdir, _ := path.Split(strings.ReplaceAll(filepath,"\\","/"))
	if(len(destdir) > 0 && !ff.FileIsExist(destdir)){
		if err := os.MkdirAll(destdir,os.ModePerm);err != nil{
			return nil,ff.errors.ComWithError("MkdirAllError:"+destdir, err)
		}
	}
	return os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
}

// 读取文件内容
func  (ff *FileHelper) ReadFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

// 读取文件内容按照字节
func  (ff *FileHelper) ReadByteSize(filepath string,size int) ([]byte,error){
	w1, err0 := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err0 != nil {
		return nil, ff.errors.ComWithError("OpenFileError:"+filepath, err0)
	}
	defer w1.Close()

	data := make([]byte, size)
	n, err := w1.Read(data)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	return data[:n],nil
}

// 将字符串写入的文件中
func  (ff *FileHelper) WriteString(filepath, content string) (bool, error) {
	//打开文件，没有则创建，有则append内容
	w1, err0 := ff.CreateOpenFile(filepath)
	
	if err0 != nil {
		return false, ff.errors.ComWithError("OpenFileError:"+filepath, err0)
	}
	defer w1.Close()

	_, err1 := w1.Write([]byte(content))
	if err1 != nil {
		return false, ff.errors.ComWithError("WriteFileError:"+filepath, err1)
	}
	return true, nil
}

// 将字符串写入的文件中
func  (ff *FileHelper) WriteBytes(filepath string, content []byte) (bool, error) {
	//打开文件，没有此文件则创建文件，将写入的内容append进去
	w1, err0 := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	
	if err0 != nil {
		return false, ff.errors.ComWithError("OpenFileError:"+filepath, err0)
	}
	defer w1.Close()

	_, err1 := w1.Write(content)
	if err1 != nil {
		return false, ff.errors.ComWithError("WriteFileError:"+filepath, err1)
	}
	return true, nil
}

//  获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func  (ff *FileHelper) WalkDirFiles(dirPth, suffix string) (files []string, dirs []string, err error) {
	files = make([]string, 0, 30)
	dirs = make([]string, 0, 30)

	suffix = strings.ToUpper(suffix)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			dirs = append(dirs, filename)
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, dirs, err
}

// 遍历指定文件返回符合filter的文件和文件夹
func  (ff *FileHelper) WalkDirFilesFilter(dirPath string, filter func(string) bool) (files []string, dirs []string, err error) {
	files = make([]string, 0, 30)
	dirs = make([]string, 0, 30)

	err = filepath.Walk(dirPath, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			if (filter(filename)) {
				dirs = append(dirs, filename)
				return nil
			}
			return nil
		}
		if (filter(filename)) {
			files = append(files, filename)
		}
		return nil
	})
	return files, dirs, err
}

// 变量指定文件并使用指定函数进行处理
func  (ff *FileHelper) WalkDirFilesHandler(dirPath string, handler func(string)){
	filepath.Walk(dirPath, func(filename string, fi os.FileInfo, err error) error {
			handler(filename)
			return nil
		})
}

// 判断文件是否存在
func  (ff *FileHelper) FileIsExist(file string) bool {
	exists := true
	if _, err := os.Stat(file); os.IsNotExist(err) {
		exists = false
	}
	return exists
}

// 判断文件是否是文件夹
func  (ff *FileHelper) IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断文件是否是文件
func  (ff *FileHelper) IsFile(path string) bool {
	return !ff.IsDir(path)
}

// 创建文件夹如果不存在则创建
func  (ff *FileHelper) MakeDir(dir string) (bool, error) {
	dir = ff.TrimPathFile(dir)
	dirpath := filepath.Dir(dir)
	if !ff.FileIsExist(dirpath) {
		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return false, err
		}
	}
	return true, nil
}

//格式化文件路径
func   (ff *FileHelper) TrimPathFile(file_name string)string {
	re2, _ := regexp.Compile("\\\\{1,}")
	strs := re2.ReplaceAllString(file_name, "/")
	re3, _ := regexp.Compile("/{2,}")
	return re3.ReplaceAllString(strs, "/")
}

// 获取文件大小
func  (ff *FileHelper) FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, ff.errors.ComWithError("GetFileSize error:"+file, e)
	}
	return f.Size(), nil
}

// 获取文件md5值
func  (ff *FileHelper) Md5File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// 获取文件sha1值
func  (ff *FileHelper) Sha1File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// 文件sha256值
func  (ff *FileHelper) Sha256File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// 文件sha512值
func  (ff *FileHelper) Sha512File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha512.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// 复制文件
func  (ff *FileHelper) CopyFile(srcName string, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)

	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

//复制目录
func  (ff *FileHelper) CopyDir(srcPath string, destPath string) (bool, error) {
	if !ff.IsDir(srcPath) {
		return false, ff.errors.ComError("SrcPathIsNotExistOrIsNotDir:" + srcPath)
	}
	if _, err := ff.MakeDir(destPath); err != nil {
		return false, ff.errors.ComError("MakeDestPathIsNotExistOrMakeFail:" + destPath)
	}

	fmt.Println("CopyDir", srcPath, "==>", destPath)
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		path = strings.ReplaceAll(path, "\\", "/")
		destNewPath := strings.Replace(path, srcPath, destPath, -1)
		fmt.Println("Copy:", path, "==>", destNewPath)
		if ff.IsDir(path) {
			ff.MakeDir(destNewPath)
		} else {
			ff.CopyFile(path, destNewPath)
		}
		return nil
	})
	if err != nil {
		return false, ff.errors.ComWithError("CopyError:", err)
	}
	return true, nil
}

// 将指定的目录或文件压缩到指定的zip中
func  (ff *FileHelper) ZipCompress(srcFile string, destZip string) error {
	if ff.FileIsExist(destZip) {
		return ff.errors.ComError("FileISExist:" + destZip)
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
				if (header.Name == srcFile) {
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

			if ! info.IsDir() {
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
func  (ff *FileHelper) ZipUnCompress(zipFile string, destDir string) error {
	if ff.FileIsExist(destDir){
		return ff.errors.ComError("The file already exists:"+destDir)
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





// WriteLinesSlice writes the given slice of lines to the given file.
func  (ff *FileHelper) WriteLinesSlice(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// ReadLinesSlice reads a text file line by line into a slice of strings.
// Not recommended for use with very large files due to the memory needed.
//
//   lines, err := fileutil.ReadLinesSlice(filePath)
//   if err != nil {
//       log.Fatalf("readLines: %s\n", err)
//   }
//   for i, line := range lines {
//       fmt.Printf("  Line: %d %s\n", i, line)
//   }
//
// nil is returned if there is an error opening the file
//
func  (ff *FileHelper) ReadLinesSlice(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ReadLinesChannel reads a text file line by line into a channel.
//
//   c, err := fileutil.ReadLinesChannel(fileName)
//   if err != nil {
//      log.Fatalf("readLines: %s\n", err)
//   }
//   for line := range c {
//      fmt.Printf("  Line: %s\n", line)
//   }
//
// nil is returned (with the error) if there is an error opening the file
//
func  (ff *FileHelper) ReadLinesChannel(filePath string) (<-chan string, error) {
	c := make(chan string)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	go func() {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()
	return c, nil
}

//获取文件修改时间 返回unix时间戳
func  (ff *FileHelper)  GetFileModTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}
