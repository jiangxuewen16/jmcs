package controller

import (
	"jmcs/core/library"
	"jmcs/core/library/socket"
	"net"
)

type SocketController struct {
	library.Controller
	Head socket.Head
	Conn *net.Conn
	//todo:socket业务相关的属性
}

type SocketControllerInterface interface {
	Init(conn *net.Conn, h socket.Head)
}

func (sc SocketController) Init(conn *net.Conn, h socket.Head){
	sc.Conn = conn
	sc.Head = h
}

func (sc SocketController) Handle()  {

}
