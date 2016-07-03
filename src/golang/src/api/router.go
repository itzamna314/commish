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
