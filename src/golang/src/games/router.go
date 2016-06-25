package games

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(public *gin.RouterGroup, private *gin.RouterGroup) {
	public.GET("/games", list)
	public.GET("/games/:id", fetch)
	private.POST("/games", create)
	private.PUT("/games/:id", replace)
}

func list(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := createRepo(db)
	if games, err := repo.ListGames(); err != nil {
		fmt.Printf("Failed to list games: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Failed to list games",
		})
	} else {
		c.JSON(200, gin.H{
			"games": games,
		})
	}
}

func fetch(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	id := c.Param("id")
	repo := createRepo(db)
	if game, err := repo.FetchGame(id); err != nil {
		fmt.Printf("Failed to fetch game: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Failed to fetch game",
		})
	} else {
		c.JSON(200, gin.H{
			"games": []Game{*game},
		})
	}
}

func create(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := createRepo(db)

	req := Game{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to game",
			"error":   err,
		})
		return
	}

	req.State = "pending"
	game, err := repo.CreateGame(&req)
	if err != nil {
		fmt.Printf("Failed to create game: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to create game",
		})
	} else {
		c.JSON(200, gin.H{
			"games": []Game{*game},
		})
	}
}

func replace(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := createRepo(db)

	publicId := c.Param("id")
	if publicId == "" {
		c.JSON(400, gin.H{
			"message": "Id is required",
		})
		return
	}

	req := Game{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to game",
			"error":   err,
		})
		return
	}

	game, err := repo.ReplaceGame(publicId, &req)
	if err != nil {
		fmt.Printf("Failed to replace game: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to replace player",
		})
	} else {
		c.JSON(200, gin.H{
			"games": []Game{*game},
		})
	}
}
