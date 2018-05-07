package file

import (
	"jmcs/core/library/socket"
	"jmcs/core/utils/file/transfer"
	"github.com/gin-gonic/gin/json"
	"jmcs/core/utils/file"
	"net"
)

type FileTransController struct {
	socket.SocketController
}

func (f FileTransController) Receive() {
	serverTransfer := &transfer.ServerTransfer{}
	f.ResolveBody(serverTransfer)

	receivePackage := serverTransfer.ReceiveFile() //接收文件
	receiveByte, _ := json.Marshal(receivePackage)

	f.Responser(receiveByte, "/file/receive", "")

	//if serverTransfer.Finished {		//文件传输完成 todo:发送方断开???
	//	f.Close()
	//}
}

func (f FileTransController) Send() {
	filePath := "C:/temp/new"
	rootPath := "C:/temp/root"
	//conn := f.Conn

	fileInfos, err := file.GetFileList(filePath, 0)

	if err != nil {

	}

	conn, _ := net.Dial("tcp", "127.0.0.1:8002")

	head := f.BuildHead("/file/receive", 200)

	for _, fileInfo := range fileInfos {
		sp := transfer.SendPackage{}
		sp.Handle(conn, rootPath, head.Bytes(), fileInfo)
	}

	f.Close() //文件传输完成,断开连接
}
