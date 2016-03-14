package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itzamna314/gin-jwt"
)

// Set everything up.
// Everything related to the public API should be
// set up under the "api" router group,
// and should be defined in per-resource files
// Other administrative api endpoints may be
// described here
func Init(masterConnection string) {
	adminConnStr = masterConnection
	dbConnections.Init(adminConnStr)

	r := gin.Default()
	r.GET("/admin/health", health)

	// Use this endpoint to get a JWT that will
	// allow you to perform admin operations
	// against a particular connection
	r.POST("/admin/logins", adminLogin)

	api = r.Group("/api")
	initPlayers(api)
	r.Run()
}
