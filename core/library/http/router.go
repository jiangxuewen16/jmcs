package http

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func init()  {
	Router = gin.Default()
}