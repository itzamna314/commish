package leagues

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(public *gin.RouterGroup, private *gin.RouterGroup) {
	public.GET("/leagues", list)
	public.GET("/leagues/:id", fetch)
	public.POST("leagues/queries", find)
	private.POST("/leagues", create)
	private.PATCH("/leagues/:id", update)
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

func find(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "Not implemented",
	})
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

	req := League{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to league",
			"error":   err,
		})
		return
	}

	league, err := repo.UpdateLeague(publicId, &req)
	if err != nil {
		fmt.Printf("Failed to update league: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to update league",
		})
		return
	}

	c.JSON(200, gin.H{
		"leagues": []League{*league},
	})
}
