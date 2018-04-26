package socket

import (
	"jmcs/core/library/socket"
	"net"
	"jmcs/core/library/controller"
)

type SRouter map[string] controller.SocketControllerInterface

type HandleFunc func(conn *net.Conn, h Head)

/*type Handle interface {
	HandleFunc()
}*/

var Router SRouter

/*添加路由*/
func Add(pattern string, h controller.SocketControllerInterface) {
	Router[pattern] = h
}

func Handle(pattern string, conn *net.Conn, h Head){
	HandleFunc := Router[pattern]
	HandleFunc.Init(conn, h)
	HandleFunc(conn, h)
}