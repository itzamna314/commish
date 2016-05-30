package dbselector

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	connectionHeader string = "X-COMMISH-CONNECTION"
)

type dbSelector struct {
	Map map[string]*sqlx.DB
}

func Create(masterConnStr string) (*dbSelector, error) {
	db, err := sqlx.Open("mysql", masterConnStr)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to connections db: %s", err)
	}

	d := dbSelector{
		Map: make(map[string]*sqlx.DB),
	}

	rows, err := db.Query(listConnectionsQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to list available connections: %s", err)
	}
	defer rows.Close()

	var publicId, connectionString string
	for rows.Next() {
		if err = rows.Scan(&publicId, &connectionString); err != nil {
			fmt.Printf("Warning: Failed to scan row with available connection: %s")
			continue
		}
		conn, err := sqlx.Open("mysql", connectionString)
		if err != nil {
			fmt.Printf("Warning: Failed to open connection %s: %s", publicId, err)
			continue
		}

		d.Map[publicId] = conn
	}

	return &d, nil
}

func (d *dbSelector) Public() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, db, err := d.loadConnection(c)
		if err != nil {
			return
		}

		c.Set("connectionName", name)
		c.Set("connectionDb", db)
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

		name, db, err := d.loadConnection(c)
		if err != nil {
			return
		}

		if name != connectionName {
			c.AbortWithError(401, fmt.Errorf("Insufficient privilege"))
			return
		}

		c.Set("connectionName", name)
		c.Set("connectionDb", db)
	}
}

func (d *dbSelector) loadConnection(c *gin.Context) (name string, db *sqlx.DB, err error) {
	hdr := c.Request.Header.Get(connectionHeader)
	if hdr == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Header %s is required", connectionHeader),
		})
		c.AbortWithError(400, fmt.Errorf("Missing connection header"))
		err = fmt.Errorf("Missing connection header")
		return
	}

	if d.Map[hdr] == nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Unknown connection %s", hdr),
		})
		c.AbortWithError(400, fmt.Errorf("Unknown connection header %s", hdr))
		err = fmt.Errorf("Unknown connection header")
		return
	}

	name = hdr
	db = d.Map[hdr]
	err = nil
	return
}

const listConnectionsQuery string = `
SELECT HEX(publicId) as publicId
     , connectionString
  FROM dbConnection
`
