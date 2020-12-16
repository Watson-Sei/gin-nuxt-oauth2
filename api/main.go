package main

import (
	"github.com/Watson-Sei/gin-nuxt-oauth2/api/config"
	"github.com/Watson-Sei/gin-nuxt-oauth2/api/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main()  {
	gin.SetMode(gin.DebugMode)
	config.DB, err = gorm.Open(mysql.Open(config.DBUrl(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db, _ := config.DB.DB()
	defer db.Close()

	r := routes.SetupRouter()
	r.Run(":8080")
}