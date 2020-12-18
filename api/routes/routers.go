package routes

import (
	"github.com/Watson-Sei/gin-nuxt-oauth2/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000","http://localhost"},
		AllowMethods: []string{"GET","POST","OPTIONS","DELETE","PUT"},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return CorsChecker(origin, cors.DefaultConfig().AllowOrigins)
		},
	}))
	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello public!",
		})
	})
	r.GET("/private", middleware.FirebaseAuth, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello private!",
		})
	})
	return r
}

func CorsChecker(origin string, corsList []string) bool {
	for _, v := range corsList {
		if origin == v {
			return true
		}
	}
	return false
}