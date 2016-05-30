package players

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(public *gin.RouterGroup, private *gin.RouterGroup) {
	public.GET("/players", list)
	private.POST("/players", create)
}

func list(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := createRepo(db)
	if players, err := repo.ListPlayers(); err != nil {
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
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := createRepo(db)

	req := Player{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to player",
			"error":   err,
		})
		return
	}

	privateId, err := repo.CreatePlayer(&req)
	if err != nil {
		fmt.Printf("Failed to create player: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to create player",
		})
		return
	}

	player, err := repo.fetchPlayerPrivate(privateId)
	if err != nil {
		fmt.Printf("Failed to fetch created player: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Failed to fetch newly-created player",
		})
		return
	}

	c.JSON(200, gin.H{
		"players": []Player{*player},
	})
}
