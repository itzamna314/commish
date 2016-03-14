package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init(masterConnection string) {
	dbConnections.Init(masterConnection)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("pong: %d", 10),
		})
	})

	api := r.Group("/api")
	api.Use(DbSelector())
	api.GET("/players", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("using conn str %s", c.MustGet("connectionString").(string)),
		})
	})
	r.Run()
}
