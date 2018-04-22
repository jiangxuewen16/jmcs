package utils

import (
	"os"
	"io/ioutil"
	"path"
)

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

func GetPathFileInfo(dir string, suffixs []string) []os.FileInfo {
	fileList, err := ioutil.ReadDir(dir)
	CheckErr(err)

	if len(suffixs) <= 0 {
		return fileList
	}

	var suffixFiles = []os.FileInfo{}
	for _, fileInfo := range fileList {
		suffix := path.Ext(fileInfo.Name())
		i := 0
		if inArray(suffix, suffixs) {
			suffixFiles[i] = fileInfo
			i++
		}
	}
	return suffixFiles
}

func GetPathFileName(dir string, suffixs []string) []string {
	fileInfos := GetPathFileInfo(dir, suffixs)
	var fileNames = []string{}

	for index, fileInfo := range fileInfos {
		fileNames[index] = fileInfo.Name()
	}

	return fileNames
}

func GetPathFilePath(dir string, suffixs []string) []string {
	fileInfos := GetPathFileInfo(dir, suffixs)

	filePaths := []string{}
	for key,fileInfo := range fileInfos {
		filePaths[key] = dir + fileInfo.Name()
	}
	return filePaths
}
