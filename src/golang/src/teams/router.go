package teams

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(public *gin.RouterGroup, private *gin.RouterGroup) {
	public.GET("/teams", list)
	public.GET("/teams/:id", fetch)
	private.POST("/teams", create)
	private.PUT("/teams/:id", replace)
}

func list(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := CreateRepo(db)
	if teams, err := repo.ListTeams(); err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("Failed to list teams: %s", err),
		})
	} else {
		c.JSON(200, gin.H{
			"teams": teams,
		})
	}
}

func fetch(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	id := c.Param("id")
	repo := CreateRepo(db)
	if team, err := repo.FetchTeam(id); err != nil {
		fmt.Printf("Failed to fetch team: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Failed to fetch team",
		})
	} else {
		c.JSON(200, gin.H{
			"teams": []Team{*team},
		})
	}
}

func create(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := CreateRepo(db)

	req := Team{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to team",
			"error":   err,
		})
		return
	}

	team, err := repo.CreateTeam(&req)
	if err != nil {
		fmt.Printf("Failed to create team: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to create team",
		})
		return
	}

	c.JSON(200, gin.H{
		"teams": []Team{*team},
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

	req := Team{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to team",
			"error":   err,
		})
		return
	}

	team, err := repo.ReplaceTeam(publicId, &req)
	if err != nil {
		fmt.Printf("Failed to replace team: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to replace team",
		})
		return
	}

	c.JSON(200, gin.H{
		"teams": []Team{*team},
	})
}
