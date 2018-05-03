package file

import (
	"jmcs/core/library/socket"
	"jmcs/core/utils/file/transfer"
)

type FileTransController struct {
	socket.SocketController
}

func (f FileTransController) MultiTrans()  {
	file := &transfer.FileClientTransfer{}
	f.ResolveBody(file)
	f.Write([]byte(file.MergeFileName))
}
