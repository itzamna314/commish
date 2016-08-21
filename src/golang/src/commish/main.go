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

	fmt.Printf("Using public key: %s, private key: %s", *certFile, *keyFile)

	r := api.Init(*masterConn, *certFile, *keyFile)
	r.StaticFile("/", "./www/index.html")
	r.Static("/assets", "./www/assets")
	r.Static("/fonts", "./www/fonts")
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == "GET" {
			c.File("./www/index.html")
			return
		}

		c.JSON(404, gin.H{
			"message": "resource not found",
		})
	})
	r.Run()
	fmt.Println("Started!")
}
