package main

import (
	"api"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	masterConn := flag.String("conn", "WebClient@tcp(localhost:3306)/auth", "MySql connection string")
	certFile := flag.String("cert", "devkeys/public.pem", "public key file")
	keyFile := flag.String("key", "devkeys/private.pem", "private key file")
	flag.Parse()

	r := api.Init(*masterConn, *certFile, *keyFile)
	r.StaticFile("/", "./www/index.html")
	r.Static("/assets", "./www/assets")
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(301, "/")
	})
	r.Run()
	fmt.Println("Started!")
}
