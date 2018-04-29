package http

import (
	"strings"
	"github.com/goinggo/mapstructure"
	"jmcs/core/utils"
	"jmcs/core/utils/net/port"
	"errors"
	"fmt"
)

type http struct {
	enable      bool		`json:"enable"`
	port        port.Port	`json:"port"`
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

	if !config.enable {
		return
	}

	if ok := config.port.CheckEnabled(nil); ok {
		err := errors.New("端口" + config.port.String() + "已被占用，请更换端口")
		utils.CheckErr(err)
	}
}

func Run() {
	initConf()

	Router.Run()
}
