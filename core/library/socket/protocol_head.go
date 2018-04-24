package socket

import (
	"jmcs/core/common"
	"reflect"
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

/*解析socket约定的数据到结构体*/
func (h Sph) parse(s string) {
	 headStrs := strings.Split(s, "\r\n\r\n")
	for i,headStr := range headStrs {
		if i == 0 {		//第一行数据是协议
			h.Protocol = headStr
			continue
		}
		keyAndValue := strings.Split(headStr,":")
		(&h).setData(keyAndValue[1], i)
	}

}

func (h *Sph) setData(value string, i int) {
	mutable := reflect.ValueOf(h).Elem()
	elem1 := reflect.TypeOf(h).Elem().Field(i).Name
	elem := mutable.FieldByName(elem1)
	if elem.CanSet() {
		switch reflect.TypeOf(h).Elem().Field(i).Type.String() {
		case "string":
			elem.SetString(value)
		case "int":
			if v,err := strconv.Atoi(value); err == nil {
				elem.SetInt(int64(v))
			}
		}
	}
}
