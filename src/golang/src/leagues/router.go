package leagues

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(public *gin.RouterGroup, private *gin.RouterGroup) {
	public.GET("/leagues", list)
	public.GET("/leagues/:id", fetch)
	private.POST("/leagues", create)
	private.PUT("/leagues/:id", replace)
}

func list(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := CreateRepo(db)
	if leagues, err := repo.ListLeagues(); err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("Failed to list leagues: %s", err),
		})
	} else {
		c.JSON(200, gin.H{
			"leagues": leagues,
		})
	}
}

func fetch(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	id := c.Param("id")
	repo := CreateRepo(db)
	if league, err := repo.FetchLeague(id); err != nil {
		fmt.Printf("Failed to fetch league: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Failed to fetch league",
		})
	} else {
		c.JSON(200, gin.H{
			"leagues": []League{*league},
		})
	}
}

func create(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := CreateRepo(db)

	req := League{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to league",
			"error":   err,
		})
		return
	}

	league, err := repo.CreateLeague(&req)
	if err != nil {
		fmt.Printf("Failed to create league: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to create league",
		})
		return
	}

	c.JSON(200, gin.H{
		"leagues": []League{*league},
	})
}

func replace(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := CreateRepo(db)

	publicId := c.Param("id")
	if publicId == "" {
		c.JSON(400, gin.H{
			"message": "Public id is required",
		})
	}

	req := League{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to league",
			"error":   err,
		})
		return
	}

	league, err := repo.ReplaceLeague(publicId, &req)
	if err != nil {
		fmt.Printf("Failed to replace league: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to replace league",
		})
		return
	}

	c.JSON(200, gin.H{
		"leagues": []League{*league},
	})
}
