package http

import (
	"strings"
	"github.com/goinggo/mapstructure"
	"jmcs/core/utils"
	"jmcs/core/utils/net/port"
	"errors"
	appRouter "jmcs/app/routers/http"
)

type http struct {
	enable      bool
	port        port.Port
	heartEnable bool
}

var Conf http
var ipAddr string

const (
	CONF_NAME = "net.http"
)

func init() {
	/*socket配置*/
	confs := strings.Split(CONF_NAME, ",")
	baseConfig, ok := utils.Configs[confs[0]][confs[1]]
	if !ok {
		Conf = http{enable: false}
	}

	Conf = http{}
	err := mapstructure.Decode(baseConfig, &Conf) //解析socket配置
	utils.CheckErr(err)

	if ok := Conf.port.CheckEnabled(nil); ok {
		err := errors.New("端口" + Conf.port.String() + "已被占用，请更换端口")
		utils.CheckErr(err)
	}
}

func Run() {
	if !Conf.enable {
		return
	}

	appRouter.Router.Run(":" + ipAddr)
}
