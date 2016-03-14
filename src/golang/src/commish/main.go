package main

import (
	"api"
	"flag"
	"fmt"
)

func main() {
	masterConn := flag.String("conn", "WebClient@tcp(localhost:3306)/auth", "MySql connection string")
	certFile := flag.String("cert", "../devkeys/dev_key.pub")
	keyFile := flag.String("key", "../devkeys/dev_key")
	flag.Parse()

	api.Init(*masterConn, *certFile, *keyFile)
	fmt.Println("Started!")
}
