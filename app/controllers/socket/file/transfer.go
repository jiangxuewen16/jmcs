package file

import (
	"jmcs/core/library/socket"
	"jmcs/core/utils/file/transfer"
	"fmt"
)

type FileTransController struct {
	socket.SocketController
}

func (f FileTransController) MultiTrans()  {
	file := &transfer.SendPackage{}
	f.ResolveBody(file)

	serverTransfer := transfer.ServerTransfer{}
	serverTransfer.ReceiveFile()
	fmt.Println(file)
	f.Write([]byte("aaaa"))
}
