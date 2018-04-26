package http

import (
	"github.com/gin-gonic/gin"
	"fmt"
	 h "jmcs/core/library/http"
	"net/http"
)

func init()  {
	h.Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	h.Router.GET("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,
		})

	})

	h.Router.GET("/form_post", func(c *gin.Context) {
		message := c.Query("message")
		nick := c.DefaultQuery("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
}
