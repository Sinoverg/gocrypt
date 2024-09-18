package main

import (
	"fmt"
	"log"

	"github.com/Cirqach/gocrypt/trithemius"
)

func main() {
	var text, key, tableSize string
	fmt.Scan(&tableSize, &text, &key)
	result, err := trithemius.Encrypt(tableSize, text, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted message: ", result)
}
