package library

import (
	"jmcs/core/utils"
	//"jmcs/core/library/socket"
	"jmcs/core/library/http"
)

const (
	CONFIG_PATH = "C:/golang/src/jmcs/config"
)

func init() {
	//初始化所有(总)配置
	filePaths := utils.GetPathFilePath(CONFIG_PATH, utils.Suffix)
	for _, filePath := range filePaths {
		config := utils.Config{}
		config.Resolve(filePath)
	}
}

func Run() {
	http.Run() //启动web服务
	//socket.Run()		//启动socket
	//todo:websoket

	//todo: Hook::listen 切点，tp那种

}
