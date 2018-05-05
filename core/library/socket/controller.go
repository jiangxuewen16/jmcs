package socket

import (
	"net"
	"encoding/json"
	"fmt"
	"bytes"
)

type SocketController struct {
	BaseUrl        string
	ActionName     string
	ControllerName string
	methodMapping  map[string]func()
	Body           []byte

	Head Head
	Conn *net.Conn

	//todo:socket业务相关的属性
}

type SocketControllerInterface interface {
	Init(conn *net.Conn, h Head)
	Write(b []byte)
}

func (sc *SocketController) Init(conn *net.Conn, h Head) {
	sc.BaseUrl = h.RequstRouter
	sc.Body = h.Body

	sc.Conn = conn
	sc.Head = h
}

func (sc SocketController) Handle() {

}

/*向socket通道写入数据*/
func (sc SocketController) Write(b []byte) {
	(*sc.Conn).Write(b)
}

func (c SocketController) setBashUrl(url string) {
	c.BaseUrl = url
}

func (c SocketController) setActionName(actionName string) {
	c.ActionName = actionName
}

func (c SocketController) setControllerName(cName string) {
	c.ControllerName = cName
}

/*获取body信息*/
func (s SocketController) getBody() []byte {
	return s.Body
}

/*解析body*/
func (s SocketController) ResolveBody(proto interface{}) {
	body := s.getBody()
	fmt.Println(string(body))
	err := json.Unmarshal(body, proto)
	if err != nil {
		panic(err)
	}
}

func (s SocketController) buildHead(router string) bytes.Buffer {
	var buf bytes.Buffer
	buf.WriteString("HEAD / SOCKET/1.0")
	buf.WriteString("\r\n\r\n")
	buf.WriteString("Requst-router:")
	buf.WriteString(router)
	buf.WriteString("\r\n\r\n")
	buf.WriteString("Status-code:")
	buf.WriteString("200")
	buf.WriteString("\r\n\r\n")
	buf.WriteString("Content-type:")
	buf.WriteString("application/json")
	buf.WriteString("\r\n\r\n")
	buf.WriteString("Authentication:")
	buf.WriteString("")
	buf.WriteString("body:")
	return buf
}

/*构造返回*/
func (s SocketController) Responser(data interface{}, router string, message string) {
	buf := s.buildHead(router)
	body := s.buildBody(&data)
	//buf.WriteString()
	s.Write(buf.Bytes())
}

func (s SocketController) buildBody(data interface{}) {

}
