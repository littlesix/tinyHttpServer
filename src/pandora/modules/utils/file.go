package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func GetFileContent(file string) string {
	b, e := ioutil.ReadFile(file)
	if e != nil {

		fmt.Println("GetFileContent", e)
		return ""
	}
	return string(b)
}
func SaveToFile(filePath, str string) int64 {
	//fmt.Println("saving", filePath, str)
	userFile := filePath

	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println("SaveToFile", err)
		return -1
	}

	fout.WriteString(str)
	return int64(len(str))
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil
}

func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("copyFile",src,dst,err.Error())
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}
func CreateDir(dir string) bool{
	err:=os.MkdirAll(dir, 0777)
	return err==nil
}
func RemoveFile(file string) error{
	err := os.Remove(file)
	return err
}
func RemoveAll(path string) error{
	err := os.RemoveAll(path)
	return err
}
