package main

import (
	"fmt"
	"strings"

	"github.com/Cirqach/gocrypt/trithemius"
)

func main() {
	text := "НШЫТУЪРЖИРЕБЩУЦПЙЧЯХЖЗУ"
	// text := "Ишукръ"
	params := "ЛАМПА"
	tableSize := "4x8"
	text = strings.ToLower(text)
	params = strings.ToLower(params)
	fmt.Println(trithemius.Decrypt(tableSize, text, params))

}
