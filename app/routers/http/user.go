package http

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"

)

var Router = gin.Default()

func init()  {

	Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	Router.GET("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,
		})

	})

	Router.GET("/form_post", func(c *gin.Context) {
		message := c.Query("message")
		nick := c.DefaultQuery("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
}
