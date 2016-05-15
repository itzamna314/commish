package admin

import (
	"crypto/rsa"
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

type Router struct {
	ConnectionString string
	PrivateKey       *rsa.PrivateKey
	PublicKey        *rsa.PublicKey
}

func (r *Router) SetupRoutes(g *gin.RouterGroup) {
	g.GET("/health", r.health)
	g.POST("/logins", r.login)
}

// Use this endpoint to get a JWT that will
// allow you to perform admin operations
// against a particular connection
func (r *Router) login(c *gin.Context) {
	var body Login
	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{
			"message": "identifier and password are required",
		})
		return
	}

	account, err := r.findUserByIdentifier(body.Identifier)
	if err != nil {
		c.JSON(403, gin.H{
			"message": "invalid username and/or password",
		})
	}

	if err = bcrypt.CompareHashAndPassword([]byte(account.passwordHash), []byte(body.Password)); err != nil {
		c.JSON(403, gin.H{
			"message": "invalid username and/or password",
		})
	}

	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims["commish/connection"] = account.connection
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	tokenString, err := token.SignedString(r.PrivateKey)

	if err != nil {
		fmt.Printf("Failed to generate token %s", err)
		c.JSON(500, gin.H{
			"message": "error generating token",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": gin.H{
			"connection": account.connection,
			"identifier": body.Identifier,
			"token":      tokenString,
		},
	})
}

func (r *Router) health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("pong: %d", 10),
	})
}
