package socket

import (
	"jmcs/core/utils/net/port"
	"jmcs/core/utils"
	"strings"
	"github.com/goinggo/mapstructure"
	"net"
	"fmt"
)

type config struct {
	enable      bool
	port        port.Port
	heartEnable bool
}

var conf config

const (
	CONF_NAME = "net.socket"
)

func init() {
	/*socket配置*/
	confs := strings.Split(CONF_NAME, ",")
	baseConfig, ok := utils.Configs[confs[0]][confs[1]]
	if !ok {
		conf = config{enable: false}
	}

	conf = config{}
	err := mapstructure.Decode(baseConfig, &conf) //解析socket配置
	utils.CheckErr(err)
}

func Run() {

	if !conf.enable {
		return
	}

	server := ":" + conf.port.ToString()
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	utils.CheckErr(err)

	fmt.Println("启动监听tcp:",conf.port)
	listen := listenAddr(tcpAddr)
	fmt.Println("启动完成")

	for {
		conn, err := listen.Accept()
		fmt.Println("连接客户端：", conn.RemoteAddr().String())
		if err != nil { //如果其中一个连接出错，只需要处理掉这个连接即可，不要结束
			fmt.Println(err) //todo:这里需要日志处理，其他的错误处理
			continue
		}

		go handleTcp(conn)
	}

}

/*监听tcp端口*/
func listenAddr(tcpAddr *net.TCPAddr) *net.TCPListener {
	listen, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckErr(err)

	return listen
}

func handleTcp(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		conn.Read(buf)

		conn.Write()
	}
}
