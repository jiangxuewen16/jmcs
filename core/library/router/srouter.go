package router

import (
	"jmcs/core/library/socket"
	"net"
)

type HandleFunc func(conn *net.Conn, h socket.Head)


