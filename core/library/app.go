package library

import (
	"jmcs/core/utils"
	"jmcs/core/library/socket"
	"jmcs/core/library/http"
	_ "jmcs/app/routers/socket" //初始化socket路由 todo:这里初始化路由
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

	/*启动socket*/
	go socket.Run() //todo:这里异步执行防止给短路，可以接着执行http应用

	/*启动web服务*/
	http.Run()

	//todo:websoket

	//todo: Hook::listen 切点，tp那种

}
