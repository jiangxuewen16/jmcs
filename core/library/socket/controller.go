package socket

import (
	"net"
	"encoding/json"
	"fmt"
	"bytes"
	"net/http"
	"jmcs/core/common"
)

type SocketController struct {
	BaseUrl        string
	ActionName     string
	ControllerName string
	methodMapping  map[string]func()
	Body           []byte

	Head                Head
	Conn                net.Conn
	ResponseContentType common.ContentType
	//todo:socket业务相关的属性
}

type SocketControllerInterface interface {
	Init(conn net.Conn, h Head)
	Write(b []byte)
}

/*初始化*/
func (s *SocketController) Init(conn net.Conn, h Head) {
	s.checkAuthentication()
	s.BaseUrl = h.RequstRouter
	s.Body = h.Body

	s.Conn = conn
	s.Head = h


}

func (s SocketController) checkAuthentication() bool{
	//s.Head.Authentication
	return true
}

func (s SocketController) Close(){
	(s.Conn).Close()
}

/*向socket通道写入数据*/
func (s SocketController) Write(b []byte) {
	(s.Conn).Write(b)
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

func (s SocketController) buildHead(router string, status int) bytes.Buffer {
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
	buf.WriteString("\r\n\r\n")
	buf.WriteString("body:")
	return buf
}

/*返回*/
func (s SocketController) Responser(data interface{}, router string, message string) {
	status := http.StatusOK

	body, err := s.buildBody(&data)
	if err != nil {
		status = http.StatusInternalServerError //服务端错误
	}
	buf := s.buildHead(router, status)
	buf.Write(body)
	s.Write(buf.Bytes())
}

func (s SocketController) buildBody(data interface{}) ([]byte, error) {
	//todo:返回类型，不仅仅是json
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
