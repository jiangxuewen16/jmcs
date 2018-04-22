package core

import (
	"jmcs/core/utils"
	"jmcs/app/routers"
	"os"
	"path"
	"strings"
)

const (
	CONFIG_PATH = "C:/golang/src/jmcs/config"
)



func init() {

	//初始化所有配置
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
	
	routers.Router.Run(":8000")
}
