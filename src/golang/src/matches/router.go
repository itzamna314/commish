package matches

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(public *gin.RouterGroup, private *gin.RouterGroup) {
	public.GET("/matches", list)
	public.GET("/matches/:id", fetch)
	private.POST("/matches", create)
	private.PUT("/matches/:id", replace)
}

func list(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := createRepo(db)
	if matches, err := repo.ListMatches(); err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("Failed to list matches: %s", err),
		})
	} else {
		c.JSON(200, gin.H{
			"matches": matches,
		})
	}
}

func fetch(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	id := c.Param("id")
	repo := createRepo(db)
	if match, err := repo.FetchMatch(id); err != nil {
		fmt.Printf("Failed to fetch match: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Failed to fetch match",
		})
	} else {
		c.JSON(200, gin.H{
			"matches": []Match{*match},
		})
	}
}

func create(c *gin.Context) {
	db := c.MustGet("connectionDb").(*sqlx.DB)
	repo := createRepo(db)

	req := Match{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to match",
			"error":   err,
		})
		return
	}

	req.State = "pending"
	match, err := repo.CreateMatch(&req)
	if err != nil {
		fmt.Printf("Failed to create match: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to create match",
		})
	} else {
		c.JSON(200, gin.H{
			"matches": []Match{*match},
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

	req := Match{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to bind to match",
			"error":   err,
		})
		return
	}

	match, err := repo.ReplaceMatch(publicId, &req)
	if err != nil {
		fmt.Printf("Failed to replace match: %s\n", err)
		c.JSON(500, gin.H{
			"message": "Server failed to replace player",
		})
	} else {
		c.JSON(200, gin.H{
			"matches": []Match{*match},
		})
	}
}
