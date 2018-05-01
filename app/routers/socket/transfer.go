package socket

import (
	"jmcs/core/library/socket"
	"jmcs/app/controllers/socket/file"
)

func init()  {

	socket.Add("/file/transfer", &file.FileTransController{}, "MultiTrans")
}