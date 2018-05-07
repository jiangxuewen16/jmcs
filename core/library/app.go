package library

import (
	"jmcs/core/utils"
	"jmcs/core/library/socket"
	"jmcs/core/library/http"
	_ "jmcs/app/routers/socket" //初始化socket路由 todo:这里初始化路由
	"os"
	"fmt"
)

const (

	CONFIG_PATH = "C:/golang/src/jmcs/config"			//todo:这里一定要用户输入
	CONF_NAME = "app" //配置名称
)

func init() {
	ok, err := utils.PathExists(CONFIG_PATH)
	if !ok {
		fmt.Println("配置目录不存在")
		os.Exit(0)
	}
	if err != nil {
		utils.CheckErr(err)
	}

	//初始化所有(总)配置
	filePaths := utils.GetPathFilePath(CONFIG_PATH, utils.Suffix)
	for _, filePath := range filePaths {
		conf := utils.Config{}
		conf.Resolve(filePath)
	}

	appConfig, ok := utils.Configs[CONF_NAME]
	if !ok {
		fmt.Println("app.yml 配置不存在，请配置该文件")
		os.Exit(0)
	}


	/*创建项目临时文件夹*/
	var tempDirName string
	if appName, ok := appConfig["name"]; ok {
		tempDirName = appName.(string)
	} else {
		tempDirName = "jmcs"
	}
	appTempDir,err := utils.MkTempDir(tempDirName)
	if err != nil {
		fmt.Println("创建项目临时文件失败")
		os.Exit(0)
	}
	appConfig["temDir"] = appTempDir
	utils.AppConfig = appConfig

}

func Run() {

	/*启动socket*/
	go socket.Run() //todo:这里异步执行防止给短路，可以接着执行http应用

	/*启动web服务*/
	http.Run()

	//todo:websoket

	//todo: Hook::listen 切点，tp那种

}
