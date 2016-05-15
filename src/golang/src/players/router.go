package players

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(public *gin.RouterGroup, private *gin.RouterGroup) {
	public.GET("/players", list)
	private.POST("/players", create)
}

func list(c *gin.Context) {
	conn := c.MustGet("connectionString").(string)
	svc := PlayerService{
		ConnectionString: conn,
	}
	if players, err := svc.ListPlayers(); err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("Failed to list players: %s", err),
		})
	} else {
		c.JSON(200, gin.H{
			"players": players,
		})
	}
}

func create(c *gin.Context) {
	conn := c.MustGet("connectionString").(string)
	c.JSON(200, gin.H{
		"message":    "create a player",
		"connection": conn,
	})
}
