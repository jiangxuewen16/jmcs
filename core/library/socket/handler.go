package socket

import "jmcs/core/common"

type Head struct {
	protocol string `json:"Protocol"`		//协议
	requstRouter string `json:"Requst-router"`		//访问路由
	statusCode int `json:"Status-code"`
	// accept []common.ContentType `json:"accept"`
	contentType common.ContentType `json:"Content-type"`
	userAgent []string `json:"User-agent"`
	authentication string `json:"Authentication"`
	body string `json:"Body"`
}


func (h Head) parse(s string) {

}