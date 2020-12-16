package routes

import (
	"github.com/Watson-Sei/gin-nuxt-oauth2/api/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controller.Main)
	return r
}
