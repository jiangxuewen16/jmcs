package utils

import (
	"os"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

/*判断文件是否存在*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/*获取文件夹中所有自定后缀文件的信息*/
func GetPathFileInfo(dir string, suffixs []string) []os.FileInfo {
	fileList, err := ioutil.ReadDir(dir)
	CheckErr(err)

	if len(suffixs) <= 0 {
		return fileList
	}

	var suffixFiles = make([]os.FileInfo, 0, len(fileList))

	for _, fileInfo := range fileList {
		suffix := path.Ext(fileInfo.Name())
		if inArray(suffix, suffixs) {
			suffixFiles = append(suffixFiles, fileInfo)
		}
	}
	return suffixFiles
}

/*获取文件夹中所有自定后缀文件的文件名*/
func GetPathFileName(dir string, suffixs []string) []string {
	fileInfos := GetPathFileInfo(dir, suffixs)
	var fileNames = []string{}

	for index, fileInfo := range fileInfos {
		fileNames[index] = fileInfo.Name()
	}

	return fileNames
}

/*获取文件夹中所有自定后缀文件的文件名（带路径）*/
func GetPathFilePath(dir string, suffixs []string) []string {
	fileInfos := GetPathFileInfo(dir, suffixs)

	filePaths := make([]string, 0, len(fileInfos))
	for _, fileInfo := range fileInfos {

		filePaths = append(filePaths, dir+"/"+fileInfo.Name())
	}
	return filePaths
}

/*创建临时文件*/
func MkTempDir(dir string) (string, error) {
	tempDir := os.TempDir() + "/" + dir
	if ok, err := PathExists(tempDir); !ok || err != nil {
		//err := os.MkdirAll(dir, 0777)  todo:权限？？？？
		if err := os.MkdirAll(tempDir, 0777); err != nil {
			return "",err
		}
	}
	return tempDir, nil
}

/*创建文件夹*/
func MkDirAll(path string) error {
	err := os.MkdirAll(path, 0777) //todo:0777的权限
	return err
}

/*
获取程序运行路径
*/
func getCurrentDirectory() (string,error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "",err
	}
	return strings.Replace(dir, "\\", "/", -1),nil
}


