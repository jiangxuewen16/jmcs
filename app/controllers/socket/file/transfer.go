package file

import (
	"jmcs/core/library/socket"
)

type FileTransController struct {
	socket.SocketController
}

func (f FileTransController) MultiTrans()  {
	body := f.Body



	f.Write([]byte(body))
}
