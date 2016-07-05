package players

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(public *gin.RouterGroup, private *gin.RouterGroup) {
	public.GET("/players", list)
	public.GET("/players/:id", fetch)
	public.POST("players/queries", find)
	private.POST("/players", create)
	private.PATCH("/players/:id", update)
}

func list(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := CreateRepo(db)
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

func fetch(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	id := c.Param("id")
	repo := CreateRepo(db)
	if player, err := repo.FetchPlayer(id); err != nil {
		fmt.Printf("Failed to fetch player: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Failed to fetch player",
		})
	} else {
		c.JSON(200, gin.H{
			"players": []Player{*player},
		})
	}
}

func find(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "Not implemented",
	})
}

func create(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := CreateRepo(db)

	req := Player{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to player",
			"error":   err,
		})
		return
	}

	player, err := repo.CreatePlayer(&req)
	if err != nil {
		fmt.Printf("Failed to create player: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to create player",
		})
		return
	}

	c.JSON(200, gin.H{
		"players": []Player{*player},
	})
}

func update(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := CreateRepo(db)

	publicId := c.Param("id")
	if publicId == "" {
		c.JSON(400, gin.H{
			"message": "Public id is required",
		})
		return
	}

	req := Player{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to player",
			"error":   err,
		})
		return
	}

	player, err := repo.UpdatePlayer(publicId, &req)
	if err != nil {
		fmt.Printf("Failed to update player: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to update player",
		})
		return
	}

	c.JSON(200, gin.H{
		"players": []Player{*player},
	})
}
