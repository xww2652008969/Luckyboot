package utils

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
)

// Writefile 写入文件，返回ture
func Writefile(filename string, chadate []byte) bool {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	defer f.Close()
	if err != nil {
		return false

	} else {
		f.Write(chadate)
		return true
	}
}

// Readfile 读取文件返回true和数据
func Readfile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	return data, err
}

// 创建文件夹
func Createfolder(dirname string) bool {
	err := os.MkdirAll(dirname, 777)
	if err != nil {
		return true
	} else {
		log.Print(err)
		return false

	}
}

// Getfile 处理上传文件流
func Getfile(file *multipart.FileHeader) ([]byte, string, bool) {
	out, err := file.Open()
	outname := file.Filename
	if err == nil {
		outdate, err := ioutil.ReadAll(out)
		if err == nil {
			return outdate, outname, true
		} else {
			return nil, outname, false
		}
	} else {
		return nil, outname, false
	}

}
