package file

import (
	"time"
	"os"
	"io/ioutil"
	"path/filepath"
)

type FileInfo struct {
	Name string
	Path string
	Size int64
	MTime time.Time
	CTime time.Duration
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

/*获取文件夹下所有文件*/
func GetFileList(dir string, level int) ([]FileInfo, error) {
	var fileList []FileInfo

	files,err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _,file := range files {
		if file.IsDir() {
			filepath.Join(dir,file.Name())
			files, err := GetFileList(filepath.Join(dir,file.Name()), level - 1)
			if err != nil {
				return nil,err
			}
			fileList = append(fileList, files...)
		} else {
			fileInfo := FileInfo{
				Name:file.Name(),
				Path:dir,
				Size:file.Size(),
				MTime:file.ModTime(),
				//CTime:,
			}
			fileList = append(fileList, fileInfo)
		}
	}
	return fileList,nil
}

