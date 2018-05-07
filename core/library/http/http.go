package http

import (
	"strings"
	"github.com/goinggo/mapstructure"
	"jmcs/core/utils"
	"jmcs/core/utils/net/port"
	"errors"
	"fmt"
	R "jmcs/app/routers/http"			//初始化该包

)

type http struct {
	Enable      bool		`json:"enable"`
	Port        port.Port	`json:"port"`
}

var config http

const (
	CONF_NAME = "net.http"
)

func initConf() {
	/*socket配置*/
	confs := strings.Split(CONF_NAME, ".")
	fmt.Print(utils.Configs)
	baseConfig, ok := utils.Configs[confs[0]][confs[1]]
	if !ok {
		return
	}

	err := mapstructure.Decode(baseConfig, &config) //解析socket配置
	utils.CheckErr(err)

	if !config.Enable {
		return
	}

	if ok := config.Port.CheckEnabled(nil); ok {
		err := errors.New("端口" + config.Port.String() + "已被占用，请更换端口")
		utils.CheckErr(err)
	}
}

func Run() {
	/*初始化http应用配置*/
	initConf()

	/*启动http应用*/
	R.Router.Run(":" + config.Port.String())
}
