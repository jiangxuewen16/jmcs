package controllers

import "github.com/gin-gonic/gin"

type BaseController struct {
	BaseUrl string
	ActionName string
	ControllerName string
	Context gin.Context
}
