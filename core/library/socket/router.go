package socket

import (
	"net"
	"reflect"
	"fmt"
)

type HandleFunc func(conn *net.Conn, h Head)

type ControllerRegister struct {
	controllerRouter *SocketControllerInterface
	method           string
}

var Router = make(map[string]ControllerRegister)

/*添加路由*/
func Add(pattern string, h SocketControllerInterface, method string) {

	checkPattern(pattern)

	controllerRegister := ControllerRegister{&h, method}
	Router[pattern] = controllerRegister

	fmt.Println("pattern")
}

/*检查是否能注册路由*/
func checkPattern(pattern string) {
	if _, ok := Router[pattern]; ok {
		panic("[" + pattern + "]" + "已存在，不能重复注册路由")
	}
}

/*处理路由*/
func Handle(conn net.Conn, h Head) {
	if len(h.RequstRouter) <= 0 {
		panic("没有自定路由") //todo:404
	}
	HandleFunc, ok := Router[h.RequstRouter]
	if !ok {
		panic("该路由不存在") //todo:404
	}

	handler := *HandleFunc.controllerRouter
	handler.Init(conn, h)

	reflectVal := reflect.ValueOf(handler)
	if val := reflectVal.MethodByName(HandleFunc.method); val.IsValid() {
		val.Call(nil)
	} else {
		panic("method doesn't exist in the controller " + HandleFunc.method)
	}

	//Mapper(handler, HandleFunc.method)
}

/*运行router 的 controller的方法*/
func Mapper(h SocketControllerInterface, method string) {
	reflectVal := reflect.ValueOf(h)
	if val := reflectVal.MethodByName(method); val.IsValid() {
		val.Call(nil)
	} else {
		panic("method doesn't exist in the controller " + method)
	}
}
