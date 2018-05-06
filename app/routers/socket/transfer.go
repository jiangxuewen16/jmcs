package socket

import (
	"jmcs/core/library/socket"
	"jmcs/app/controllers/socket/file"
)

func init()  {

	socket.Add("/file/receive", &file.FileTransController{}, "Receive")		//文件接收

	socket.Add("/file/send", &file.FileTransController{}, "Send")		//文件发送
}