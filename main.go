package main

import (
	"fmt"

	"github.com/Cirqach/gocrypt/passgen"
)

func main() {
	fmt.Println("Generated password: ", passgen.Generate(12, "IvanovSergey"))
}
