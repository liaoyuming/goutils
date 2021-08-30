package efile

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if !os.IsNotExist(err) {
		return false, err
	}
	return false, nil
}

func CreateFile(path string) error {
	dir := filepath.Dir(path)
	err := ExistsOrCreateDir(dir)
	if err != nil {
		return err
	}
	_, err = os.Create(path)
	return err
}

func Remove(path string) error {
	return os.Remove(path)
}

//  RemoveAllDir removes path and any children it contains.
func RemoveAll(path string) error {
	return os.RemoveAll(path)
}


func ExistsOrCreateFile(path string) error {
	ok, err := PathExists(path)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	return CreateFile(path)
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

func ExistsOrCreateDir(path string) error {
	if ok, _ := PathExists(path); !ok {
		err := CreateDir(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func FileNameWithoutExt(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}

func WriteToFileEnd(filePath string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	// 查找文件末尾的偏移量
	n, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	// 从末尾的偏移量开始写入内容
	_, err = f.WriteAt([]byte(content), n)

	return err
}

func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	// 查找文件末尾的偏移量
	n, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	// 从末尾的偏移量开始写入内容
	_, err = f.WriteAt([]byte(content), n)

	return err
}

func ReadFile(fileName string) (string, error) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(f), nil
}