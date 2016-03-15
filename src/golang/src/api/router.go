package api

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itzamna314/gin-jwt"
	"io/ioutil"
)

// Set everything up.
// Everything related to the public API should be
// set up under the "api" router group,
// and should be defined in per-resource files
// Other administrative api endpoints may be
// described here
func Init(masterConnection, certFile, keyFile string) {
	publicKey, privateKey := initCert(certFile, keyFile)
	admin := adminRouter{
		ConnectionString: masterConnection,
		PrivateKey:       privateKey,
		PublicKey:        publicKey,
	}
	dbSelector := dbSelector{
		ConnectionString: masterConnection,
	}
	dbSelector.Init()

	validator := jwtauth.Validator{
		Key:    publicKey,
		Method: jwt.SigningMethodRS256,
	}

	r := gin.Default()
	r.GET("/admin/health", health)

	r.POST("/admin/logins", admin.LoginEndpoint)

	publicApi := r.Group("/api")
	publicApi.Use(dbSelector.Public())

	protectedApi := r.Group("/api")
	protectedApi.Use(validator.Middleware())
	protectedApi.Use(dbSelector.Protected())

	players := playersRouter{}
	publicApi.GET("/players", players.List)
	protectedApi.POST("/players", players.Create)

	r.Run()
}

func initCert(certFile, keyFile string) (*rsa.PublicKey, *rsa.PrivateKey) {
	cert, err := ioutil.ReadFile(certFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to read cert file %s: %s", certFile, err))
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(cert)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse cert file %s: %s", certFile, err))
	}

	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to read key file %s: %s", keyFile, err))
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse key file %s: %s", keyFile, err))
	}

	return pubKey, privKey
}
