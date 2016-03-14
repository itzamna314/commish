package api

func initPlayers(r *RouterGroup) {
	r.Use(DbSelector())
	r.GET("/players", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("using conn str %s", c.MustGet("connectionString").(string)),
		})
	})

}
