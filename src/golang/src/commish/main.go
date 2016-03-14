package main

import (
	"api"
	"flag"
	"fmt"
)

func main() {
	masterConn := flag.String("conn", "WebClient@tcp(localhost:3306)/auth", "MySql connection string")
	flag.Parse()

	api.Init(*masterConn)
	fmt.Println("Started!")
}
