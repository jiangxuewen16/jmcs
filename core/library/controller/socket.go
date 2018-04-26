package controller

import (
	"jmcs/core/library"
	"jmcs/core/library/socket"
	"net"
)

type SocketController struct {
	library.Controller
	Head socket.Head
	conn *net.Conn
	//todo:socket业务相关的属性
}

func (sc SocketController) HandleHead(conn *net.Conn, h socket.Head){

}

func (sc SocketController) Init(conn *net.Conn, h socket.Head){
	sc.conn = conn
	sc.Head = h

}
