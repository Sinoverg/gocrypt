package main

import (
	"fmt"

	"github.com/Cirqach/gocrypt/caesar"
)

func main() {
	c := caesar.NewCaesar("ru", 5, "ЗИМА")
	c.CreateTable()
	fmt.Println(string(c.Encrypt("привет")))
}
