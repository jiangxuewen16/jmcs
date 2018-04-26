package socket

import (
	"jmcs/core/library/socket"
	"net"
)

type SRouter map[string] HandleFunc

type HandleFunc func(conn *net.Conn, h socket.Head)

type Handle interface {
	HandleFunc()
}

var Router SRouter

/*添加路由*/
func Add(pattern string, h HandleFunc) {
	Router[pattern] = h
}

func