package socket

import (
	"jmcs/core/library/socket"
	"net"
	"jmcs/core/library/controller"
	"reflect"
	"strings"
)

type SRouter map[string]ControllerRegister

type HandleFunc func(conn *net.Conn, h Head)

type ControllerRegister struct {
	controllerRouter *controller.SocketControllerInterface
	method           string
}

var Router SRouter

/*添加路由*/
func Add(pattern string, h *controller.SocketControllerInterface, method string) {

	checkPattern(pattern)

	controllerRegister := ControllerRegister{h, method}
	Router[pattern] = controllerRegister
}

/*检查是否能注册路由*/
func checkPattern(pattern string) {
	if _, ok := Router[pattern]; ok {
		panic("[" + pattern + "]" + "已存在，不能重复注册路由")
	}
}

/*处理路由*/
func Handle(conn *net.Conn, h Head) {
	if len(h.RequstRouter) <= 0 {
		panic("没有自定路由")  //todo:404
	}
	HandleFunc, ok := Router[h.RequstRouter]
	if !ok {
		panic("该路由不存在")  //todo:404
	}

	handler := *HandleFunc.controllerRouter
	handler.Init(conn, h)

	Mapper(handler, HandleFunc.method)
}

/*运行router 的 controller的方法*/
func Mapper(h controller.SocketControllerInterface, method string) {
	reflectVal := reflect.ValueOf(&h)
	if val := reflectVal.MethodByName(method); val.IsValid() {
		val.Call(nil)
	} else {
		panic("method doesn't exist in the controller " + method)
	}
}
