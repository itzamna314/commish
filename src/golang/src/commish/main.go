package main

import (
	"api"
	"flag"
	"fmt"
)

func main() {
	masterConn := flag.String("conn", "WebClient@tcp(localhost:3306)/auth", "MySql connection string")
	certFile := flag.String("cert", "devkeys/public.pem", "public key file")
	keyFile := flag.String("key", "devkeys/private.pem", "private key file")
	flag.Parse()

	api.Init(*masterConn, *certFile, *keyFile)
	fmt.Println("Started!")
}
