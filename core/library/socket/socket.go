package socket

import (
	"jmcs/core/utils/net/port"
	"jmcs/core/utils"
	"strings"
	"github.com/goinggo/mapstructure"
	"net"
	"fmt"
	"errors"
)

type socket struct {
	Enable      bool
	Port        port.Port
	HeartEnable bool
}

var conf socket

const (
	CONF_NAME = "net.socket" //配置名称，靠这个解析出该应用具体配置
)

func Run() {
	/*初始化socket应用配置*/
	initconf()

	/*启动socket服务端*/
	start()
}

func initconf() {
	/*socket配置*/
	confs := strings.Split(CONF_NAME, ".")
	baseconfig, ok := utils.Configs[confs[0]][confs[1]]
	if !ok {
		conf = socket{Enable: false}
	}

	conf = socket{}
	err := mapstructure.Decode(baseconfig, &conf) //解析socket配置
	utils.CheckErr(err)

	if !conf.Enable {
		return
	}

	if ok := conf.Port.CheckEnabled(nil); ok {
		err := errors.New("端口" + conf.Port.String() + "已被占用，请更换端口")
		utils.CheckErr(err)
	}
}

func start()  {
	server := ":" + conf.Port.String()
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	utils.CheckErr(err)

	fmt.Println("启动监听tcp:", conf.Port)
	listen := listenAddr(tcpAddr)
	fmt.Println("启动完成")

	for {
		conn, err := listen.Accept()
		fmt.Println("连接客户端：", conn.RemoteAddr().String())
		if err != nil { //如果其中一个连接出错，只需要处理掉这个连接即可，不要结束
			fmt.Println(err) //todo:这里需要日志处理，其他的错误处理
			continue
		}

		go handleTcp(&conn)
	}
}

/*监听tcp端口*/
func listenAddr(tcpAddr *net.TCPAddr) *net.TCPListener {
	listen, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckErr(err)

	return listen
}

/*socket业务具体处理 todo:这里调用路由合不合理*/
func handleTcp(conn *net.Conn) {

	for {

		buf := make([]byte, 512)
		(*conn).Read(buf) //todo：读取数据，需要按约定处理

		head := Head{}
		head.parse(buf)

		Handle(conn, head)
	}
}
