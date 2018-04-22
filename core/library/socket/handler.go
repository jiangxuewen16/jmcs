package socket

import "jmcs/core/common"

type Head struct {
	statusCode int `json:"Status-code"`
	version string `json:"Version"`
	accept []common.ContentType `json:"accept"`
	contentType common.ContentType `json:"Content-type"`
	userAgent []string `json:"User-agent"`
	authentication string `json:"Authentication"`
	body string `json:"Body"`
}


func (h Head) parse(s string) {

}