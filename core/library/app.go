package library

import (
	"strings"
	"path"
	"os"
	"jmcs/core/utils"
	"jmcs/core/library/socket"
	"jmcs/core/library/http"
)

const (
	CONFIG_PATH = "C:/golang/src/jmcs/config"
)

func init() {

	//初始化所有配置 todo：需不需要放到core.go里面去
	filePaths := utils.GetPathFilePath(CONFIG_PATH, utils.Suffix)
	for _, filePath := range filePaths {
		fileName := strings.Trim(path.Base(filePath), path.Ext(filePath))
		f, err := os.Open(filePath)
		utils.CheckErr(err)
		config := utils.Config{}
		config.Resolve(f, fileName)
	}
}

func Run() {
	http.Run()			//启动web服务
	socket.Run()		//启动socket
	//todo:websoket

	//todo: Hook::listen 切点，tp那种


}