package web

import (
	"github.com/gin-gonic/gin"
	"github.com/karim-w/go-clean-commments/pkg/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	controllers.HandleRequests(r)
	r.Run()
	return r
}
