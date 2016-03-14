package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil && err != io.EOF {
		fmt.Printf("Something went wrong: %s", err)
		os.Exit(1)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)

	fmt.Printf("%s", string(hashed))

	// sanity check
	err = bcrypt.CompareHashAndPassword(hashed, []byte(text))
	if err != nil {
		fmt.Printf("Failed to compare hash and password: %s", err)
		os.Exit(1)
	}
}
