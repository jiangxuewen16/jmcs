package file

import "jmcs/core/library/controller"

type FileTransController struct {
	controller.SocketController
}

func (f FileTransController) MultiTrans()  {
	body := f.Head.Body
}