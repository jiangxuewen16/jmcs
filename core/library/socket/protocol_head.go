package socket

import (
	"jmcs/core/common"
	"reflect"
	"fmt"
	"strings"
	"strconv"
)

//sph  -> socket protocol head
type Sph struct {
	Protocol string `json:"Protocol"`		//协议
	RequstRouter string `json:"Requst-router"`		//访问路由
	StatusCode int `json:"Status-code"`		 //传输状态码
	ContentType common.ContentType `json:"Content-type"`
	Authentication string `json:"Authentication"`	//对于socket来说没实际意义
	Body string `json:"Body"`		//请求body,用于存放具体传输数据

	// userAgent []string `json:"User-agent"`
	// accept []common.ContentType `json:"accept"`
}




func (h Sph) parse(s string) {
	 //strs := strings.Split(s, "\r\n\r\n")

	 /*for _,str := range strs {

	 }*/
}

func p(sph *Sph, s string, value string, i int)  {
	mutable := reflect.ValueOf(sph).Elem()
	elem := mutable.FieldByName(s)
	elem1 := reflect.TypeOf(sph).Elem().Field(i).Tag.Get("json")
	fmt.Println(elem1)
	if elem.CanSet() {
		elem.SetString(value)
	}
}



///////////////////////////////////////////////////////////////
func main() {

	a := Sph{}

	str := "HEAD / SOCKET/1.0 [没实际意义留着吧]\r\n\r\nRequst-router: /file/send [用于对应后台的路由]\r\n\r\nStatus-code:200\r\n\r\nContent-type: application/json [用于告诉对方body的数据类型,例如文件传输:multipart/file-data]\r\n\r\n" +
		"Authentication:\r\n\r\nBody: [请求体]"

	strs := strings.Split(str, "\r\n\r\n")
	for i,st := range strs {
		if i == 0 {
			a.Protocol = st
			continue
		}
		ff := strings.Split(st,":")
		p(&a, ff[1], i)
	}
	fmt.Println(a)
}

func p(sph *Sph, value string, i int) {
	mutable := reflect.ValueOf(sph).Elem()
	elem1 := reflect.TypeOf(sph).Elem().Field(i).Name
	elem := mutable.FieldByName(elem1)
	if elem.CanSet() {
		switch reflect.TypeOf(sph).Elem().Field(i).Type.String() {
		case "string":
			elem.SetString(value)
		case "int":
			if v,err := strconv.Atoi(value); err != nil {
				elem.SetInt(int64(v))

			}
		}
	}
}
