package api

import (
	"admin"
	"crypto/rsa"
	"dbselector"
	"fmt"
	"games"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itzamna314/gin-jwt"
	"io/ioutil"
	"leagues"
	"matches"
	"os"
	"players"
	"teams"
)

// Set everything up.
// Everything related to the public API should be
// set up under the "api" router group,
// and should be defined in per-resource files
// Other administrative api endpoints may be
// described here
func Init(masterConnection, certFile, keyFile string) *gin.Engine {
	r := gin.Default()

	adminApi := r.Group("/api/admin")

	publicKey, privateKey := initCert(certFile, keyFile)
	adminRouter := admin.Router{
		ConnectionString: masterConnection,
		PrivateKey:       privateKey,
		PublicKey:        publicKey,
	}
	adminRouter.SetupRoutes(adminApi)

	dbSelector, err := dbselector.Create(masterConnection)
	if err != nil {
		fmt.Printf("Failed to initialize connection map: %s\n", err)
		os.Exit(2)
	}

	validator := jwtauth.Validator{
		Key:    publicKey,
		Method: jwt.SigningMethodRS256,
	}

	publicApi := r.Group("/api")
	publicApi.Use(dbSelector.Public())

	protectedApi := r.Group("/api")
	protectedApi.Use(validator.Middleware())
	protectedApi.Use(dbSelector.Protected())

	games.SetupRoutes(publicApi, protectedApi)
	leagues.SetupRoutes(publicApi, protectedApi)
	matches.SetupRoutes(publicApi, protectedApi)
	players.SetupRoutes(publicApi, protectedApi)
	teams.SetupRoutes(publicApi, protectedApi)

	return r
}

func initCert(certFile, keyFile string) (*rsa.PublicKey, *rsa.PrivateKey) {
	cert, err := ioutil.ReadFile(certFile)
	if err != nil {
		fmt.Printf("Failed to read cert file %s. Aborting startup.\n\nDetails:\n%s\n", certFile, err)
		os.Exit(1)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(cert)
	if err != nil {
		fmt.Printf("Failed to parse cert file %s. Aborting startup.\n\nDetails:\n %s\n", certFile, err)
		os.Exit(2)
	}

	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		fmt.Printf("Failed to read key file %s. Aborting startup.\n\nDetails:\n%s\n", keyFile, err)
		os.Exit(3)
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		fmt.Printf("Failed to parse key file %s. Aborting startup.\n\nDetails:%s\n", keyFile, err)
		os.Exit(4)
	}

	return pubKey, privKey
}
