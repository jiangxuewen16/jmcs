package controller

import (
	"jmcs/core/library"
	"jmcs/core/library/socket"
	"net"
)

type SocketController struct {
	library.Controller
	Head socket.Head
	Conn net.Conn
	//todo:socket业务相关的属性
}

type SocketControllerInterface interface {
	Init(conn *net.Conn, h socket.Head)
	Write(b []byte)
}

func (sc SocketController) Init(conn *net.Conn, h socket.Head){
	sc.BaseUrl = h.RequstRouter
	sc.Body = h.Body
	
	sc.Conn = conn
	sc.Head = h
}

func (sc SocketController) Handle()  {

}

/*向socket通道写入数据*/
func (sc SocketController) Write(b []byte)  {
	sc.Conn.Write(b)
}
