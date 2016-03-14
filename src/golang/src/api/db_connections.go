package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	connectionHeader string = "X-COMMISH-CONNECTION"
)

type connectionMap struct {
	Map map[string]string
}

var dbConnections connectionMap

func (c *connectionMap) Init(masterConnection string) {
	db, err := sql.Open("mysql", masterConnection)
	if err != nil {
		fmt.Printf("Failed to connect to connections db: %s", err)
		os.Exit(2)
	}

	c.Map = make(map[string]string)

	rows, err := db.Query(listConnectionsQuery)
	if err != nil {
		fmt.Printf("Failed to list available connections: %s", err)
		os.Exit(2)
	}
	defer rows.Close()

	var publicId, connectionString string
	for rows.Next() {
		err := rows.Scan(&publicId, &connectionString)
		if err != nil {
			fmt.Printf("Failed to scan row with available connection: %s")
			continue
		}
		c.Map[publicId] = connectionString
	}
}

func DbSelector() gin.HandlerFunc {
	return func(c *gin.Context) {
		hdr := c.Request.Header.Get(connectionHeader)
		if hdr == "" {
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("Header %s is required", connectionHeader),
			})
			c.AbortWithError(400, fmt.Errorf("Missing connection header"))
			return
		}

		if dbConnections.Map[hdr] == "" {
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("Unknown connection %s", hdr),
			})
			c.AbortWithError(400, fmt.Errorf("Unknown connection header %s", hdr))
			return
		}

		c.Set("connectionName", hdr)
		c.Set("connectionString", dbConnections.Map[hdr])
	}
}

const listConnectionsQuery string = `
SELECT HEX(publicId) as publicId
     , connectionString
  FROM dbConnection
`
