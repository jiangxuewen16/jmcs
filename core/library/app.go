package library

import (
	"jmcs/core/utils"
	"jmcs/core/library/socket"
	"jmcs/core/library/http"
	_ "jmcs/app/routers/socket" //初始化socket路由 todo:这里初始化路由
	"os"
	"fmt"
	"sync"
)

const (
	CONFIG_PATH = "D:/golang/src/jmcs/config"			//todo:这里一定要用户输入
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
	createTempDir(appConfig)
}

func Run() {
	var waitGroup sync.WaitGroup			//用于同步goroutine

	/*启动socket*/
	waitGroup.Add(1)
	go socket.Run(waitGroup)

	/*启动web服务*/
	waitGroup.Add(1)
	go http.Run(waitGroup)

	//todo:websoket

	//todo: Hook::listen 切点，tp那种

	waitGroup.Wait()		//阻塞住主线程，可同时监听多个端口
}

func createTempDir(appConfig map[string]interface{})  {
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
