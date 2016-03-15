package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: hashutil <password>")
		os.Exit(1)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(os.Args[1]), bcrypt.DefaultCost)

	fmt.Printf("%s", string(hashed))

	// sanity check
	err = bcrypt.CompareHashAndPassword(hashed, []byte(os.Args[1]))
	if err != nil {
		fmt.Printf("Failed to compare hash and password: %s", err)
		os.Exit(1)
	}
}
