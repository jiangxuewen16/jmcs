package socket

import (
	"net"
	"encoding/json"
	"fmt"
)

type SocketController struct {
	BaseUrl string
	ActionName string
	ControllerName string
	methodMapping map[string]func()
	Body []byte

	Head Head
	Conn *net.Conn
	//todo:socket业务相关的属性
}

type SocketControllerInterface interface {
	Init(conn *net.Conn, h Head)
	Write(b []byte)
}

func (sc *SocketController) Init(conn *net.Conn, h Head){
	sc.BaseUrl = h.RequstRouter
	sc.Body = h.Body

	sc.Conn = conn
	sc.Head = h
}

func (sc SocketController) Handle()  {

}

/*向socket通道写入数据*/
func (sc SocketController) Write(b []byte)  {
	(*sc.Conn).Write(b)
}

func (c SocketController) setBashUrl(url string){
	c.BaseUrl = url
}

func (c SocketController) setActionName(actionName string){
	c.ActionName = actionName
}

func (c SocketController) setControllerName(cName string) {
	c.ControllerName = cName
}

func (s SocketController) getBody() []byte {
	return s.Body
}

func (s SocketController) ResolveBody(proto interface{})  {
	body := s.getBody()
	fmt.Println(string(body))
	err := json.Unmarshal(body, proto)
	if err != nil {
		panic(err)
	}
}