package api

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var adminConnStr string

type Login struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

func adminLogin(c *gin.Context) {
	var body Login
	if c.BindJson(&body) != nil {
		c.JSON(400, gin.H{
			"message": "identifier and password are required",
		})
	}

	db, err := sql.Open("mysql", adminConnStr)
	if err != nil {
		fmt.Printf("Failed to connect to admin db: %s", err)
		c.JSON(500, gin.H{
			"message": "failed to connect to admin db",
		})
	}

	rows, err := db.Query(findLoginsQuery, body.Identifier)
	if err != nil {
		fmt.Printf("Failed to query admin db: %s", err)
		c.JSON(500, gin.H{
			"message": "failed to connect to admin db",
		})
	}

	var authorizedConnection *string = nil

	for rows.Next() {
		var password, connection string
		err := rows.Scan(&password, &connection)
		if err != nil {
			fmt.Printf("Failed to scan row from admin db: %s", err)
			c.JSON(500, gin.H{
				"message": "failed to connect to admin db",
			})
		}
		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password))

		if err == nil {
			// We found a user
			authorizedConnection = &connection
			break
		}
	}

	if authorizedConnection == nil {
		c.JSON(404, gin.H{
			"message": "invalid username/password combination",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["commish/connection"] = *authorizedConnection
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenString, err := token.SignedString([]byte(superSecretKey))

	c.JSON(200, gin.H{
		"user": gin.H{
			"identifier": body.Identifier,
			"token":      tokenString,
		},
	})
}

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("pong: %d", 10),
	})
}

var findLoginsQuery string = `
SELECT pl.password
     , c.publicId as connection
  from principalLogin pl
  join principal p on p.id = pl.principalId
  join dbConnection c on c.id = p.dbConnectionId
 WHERE pl.identifier = ?
`
