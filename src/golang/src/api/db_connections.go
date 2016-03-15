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

type dbSelector struct {
	Map              map[string]string
	ConnectionString string
}

func (d *dbSelector) Init() {
	db, err := sql.Open("mysql", d.ConnectionString)
	if err != nil {
		fmt.Printf("Failed to connect to connections db: %s", err)
		os.Exit(2)
	}

	d.Map = make(map[string]string)

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
		d.Map[publicId] = connectionString
	}
}

func (d *dbSelector) Public() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, connStr, err := d.loadConnection(c)
		if err != nil {
			return
		}

		c.Set("connectionName", name)
		c.Set("connectionString", connStr)
	}
}

// Only use after a jwtauth.Validator
func (d *dbSelector) Protected() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(map[string]interface{})
		connectionName, ok := claims["commish/connection"].(string)
		if !ok {
			c.AbortWithError(401, fmt.Errorf("Insufficient privilege"))
			return
		}

		name, connStr, err := d.loadConnection(c)
		if err != nil {
			return
		}

		if name != connectionName {
			c.AbortWithError(401, fmt.Errorf("Insufficient privilege"))
			return
		}

		c.Set("connectionName", name)
		c.Set("connectionString", connStr)
	}
}

func (d *dbSelector) loadConnection(c *gin.Context) (name, connStr string, err error) {
	hdr := c.Request.Header.Get(connectionHeader)
	if hdr == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Header %s is required", connectionHeader),
		})
		c.AbortWithError(400, fmt.Errorf("Missing connection header"))
		err = fmt.Errorf("Missing connection header")
		return
	}

	if d.Map[hdr] == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Unknown connection %s", hdr),
		})
		c.AbortWithError(400, fmt.Errorf("Unknown connection header %s", hdr))
		err = fmt.Errorf("Unknown connection header")
		return
	}

	name = hdr
	connStr = d.Map[hdr]
	err = nil
	return
}

const listConnectionsQuery string = `
SELECT HEX(publicId) as publicId
     , connectionString
  FROM dbConnection
`
