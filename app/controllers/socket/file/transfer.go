package file

import (
	"jmcs/core/library/socket"
	"jmcs/core/utils/file/transfer"
	"github.com/gin-gonic/gin/json"
)

type FileTransController struct {
	socket.SocketController
}

func (f FileTransController) MultiTrans()  {
	serverTransfer := &transfer.ServerTransfer{}
	f.ResolveBody(serverTransfer)

	//serverTransfer := transfer.ServerTransfer{}
	receivePackage := serverTransfer.ReceiveFile()
	s, _ := json.Marshal(receivePackage)
	//fmt.Println(serverTransfer)
	f.Write([]byte(s))
}
