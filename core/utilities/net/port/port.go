package port

import (
	"net"
)

type Port int16

const (
	HTTPS = Port(433)
	HTTP  = Port(80)
	MYSQL = Port(3306)
	REDIS = Port(6379)
	SHELL = Port(22)
	FTP   = Port(21)
)

const (
	MIN = 1
	MAX = 1 << 16 - 1
)

func (p Port) getPort() int {
	return int(p)
}

func (p Port) CheckEnabled(ip net.IP) bool {
	/*默认本机ip*/
	if ip == nil {
		ip = net.IP{127,0,0,1}
	}

	tcpAddr := net.TCPAddr{
		IP: ip,
		Port: p.getPort(),
	}

	tcpConn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if err != nil {
		return false
	}
	defer tcpConn.Close()
	return true
}

func (p Port) checkconnd(ip net.IP) {

}
