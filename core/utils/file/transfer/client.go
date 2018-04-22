package transfer

import (
	"net"
	"jmcs/core/utils/net/port"
)

type FileTransfer struct {
	host net.IP
	port port.Port

	fileName      string
	mergeFileName string
	coroutine     int
	bufSize       int64
}

func NewDefaultFileTransfer() FileTransfer {
	return FileTransfer{
		host: net.IP{127, 0, 0, 1},
		port: 8001,
	}
}