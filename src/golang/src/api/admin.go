package api

import (
	"crypto/rsa"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Login struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type adminRouter struct {
	ConnectionString string
	PrivateKey       *rsa.PrivateKey
	PublicKey        *rsa.PublicKey
}

// Use this endpoint to get a JWT that will
// allow you to perform admin operations
// against a particular connection
func (a *adminRouter) LoginEndpoint(c *gin.Context) {
	var body Login
	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{
			"message": "identifier and password are required",
		})
		return
	}

	db, err := sql.Open("mysql", a.ConnectionString)
	if err != nil {
		fmt.Printf("Failed to connect to admin db: %s", err)
		c.JSON(500, gin.H{
			"message": "failed to connect to admin db",
		})
		return
	}

	rows, err := db.Query(findLoginsQuery, body.Identifier)
	if err != nil {
		fmt.Printf("Failed to query admin db: %s", err)
		c.JSON(500, gin.H{
			"message": "failed to connect to admin db",
		})
		return
	}

	var authorizedConnection *string = nil

	for rows.Next() {
		var passwordHash, connection string
		err := rows.Scan(&passwordHash, &connection)
		if err != nil {
			fmt.Printf("Failed to scan row from admin db: %s", err)
			c.JSON(500, gin.H{
				"message": "failed to connect to admin db",
			})
		}

		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(body.Password))

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
		return
	}

	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims["commish/connection"] = *authorizedConnection
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenString, err := token.SignedString(a.PrivateKey)

	if err != nil {
		fmt.Printf("Failed to generate token %s", err)
		c.JSON(500, gin.H{
			"message": "error generating token",
		})
		return
	}

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
SELECT pl.passwordHash
     , c.publicId as connection
  from principalLogin pl
  join principal p on p.id = pl.principalId
  join dbConnection c on c.id = p.dbConnectionId
 WHERE pl.identifier = ?
`
