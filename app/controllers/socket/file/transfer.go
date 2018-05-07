package file

import (
	"jmcs/core/library/socket"
	"jmcs/core/utils/file/transfer"
	"github.com/gin-gonic/gin/json"
	"jmcs/core/utils/strings"
	"jmcs/core/utils/file"
)

type FileTransController struct {
	socket.SocketController
}

func (f FileTransController) Receive() {
	serverTransfer := &transfer.ServerTransfer{}
	f.ResolveBody(serverTransfer)

	receivePackage := serverTransfer.ReceiveFile()		//接收文件
	receiveByte, _ := json.Marshal(receivePackage)

	f.Responser(receiveByte, "/file/receive", "")

	//if serverTransfer.Finished {		//文件传输完成 todo:发送方断开???
	//	f.Close()
	//}
}

func (f FileTransController) Send() {
	filePath := "C:/temp/new"
	rootPath := "C:/temp/root"
	conn := f.Conn

	fileInfos, err := file.GetFileList(filePath, 0)

	if err != nil {

	}

	for _, fileInfo := range fileInfos {
		sp := transfer.SendPackage{
			Size:          1024 * 1024,
			FileName:      fileInfo.Name,
			Path:          "",
			RootPath:      rootPath,
			MergeFileName: fileInfo.Name,
			Token:         strings.Rand().Hex(),		//UUID
			Coroutine:     10,
			BufSize:       1024 * 1024,
			//Position:i,
			//Data:          []byte("1234567890\r\n"),
		}

		sp.Handle(conn)
	}

	f.Close()		//文件传输完成,断开连接
}
