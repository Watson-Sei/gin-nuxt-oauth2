package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main(context *gin.Context)  {
	context.JSON(http.StatusOK, gin.H{
		"message": "main json",
	})
}
