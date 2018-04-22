package transfer

import (
	"net"
	"jmcs/core/utils/net/port"
)

var (
	host = "192.168.1.8" //服务端IP
	//port   = "9090"            //服务端端口
	//remote = host + ":" + port //构造连接串

	fileName      = "node.exe" //待发送文件名称
	mergeFileName = "mm.exe"   //待合并文件名称
	coroutine     = 10         //协程数量或拆分文件的数量
	bufsize       = 1024       //单次发送数据的大小
)

type FileTransfer struct {
	host net.IP
	port port.Port

	fileName      string
	mergeFileName string
	coroutine     int
	bufSize       int
}

func NewDefaultFileTransfer() FileTransfer {
	return FileTransfer{
		host: net.IP{127, 0, 0, 1},
		port:
	}
}
