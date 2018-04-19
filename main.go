package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"flag"
)

var wsPort  = flag.Int("w", 0, "开启websocket,并监听指定端口（默认8002）")
var sPort = flag.Int("s", 0, "开启socket,并监听指定端口（默认8001）")

func main() {

	socketServcer()		//socket服务

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	router.GET("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,
		})


	})

	router.GET("/form_post", func(c *gin.Context) {
		message := c.Query("message")
		nick := c.DefaultQuery("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8000")

}

func socketServcer()  {
	flag.Parse()

	if *wsPort > 0 {

	}
}
