package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type playersRouter struct{}

func initPlayers(r *gin.RouterGroup) {
}

func (p *playersRouter) List(c *gin.Context) {
	conn := c.MustGet("connectionString").(string)
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("using conn str %s", conn),
	})
}

func (p *playersRouter) Create(c *gin.Context) {
	conn := c.MustGet("connectionString").(string)
	c.JSON(200, gin.H{
		"message":    "create a player",
		"connection": conn,
	})
}
